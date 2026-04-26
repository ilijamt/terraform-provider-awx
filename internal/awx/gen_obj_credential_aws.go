package awx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialAwsTerraformModel exposes the typed AWX Amazon Web Services
// credential (credential_aws) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialAwsTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Password       types.String `tfsdk:"password" json:"-"`
	SecurityToken  types.String `tfsdk:"security_token" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialAwsTerraformModel) Clone() credentialAwsTerraformModel {
	return *o
}

type credentialAwsBodyRequestModel struct {
	CredentialType int64           `json:"credential_type"`
	Description    string          `json:"description,omitempty"`
	Inputs         json.RawMessage `json:"inputs,omitempty"`
	Name           string          `json:"name"`
	Organization   int64           `json:"organization,omitempty"`
	Team           int64           `json:"team,omitempty"`
	User           int64           `json:"user,omitempty"`
}

// BodyRequest folds typed input fields back into a single `inputs` JSON object;
// null/unknown values are dropped so the API doesn't receive empty strings for
// unset optionals.
func (o *credentialAwsTerraformModel) BodyRequest() *credentialAwsBodyRequestModel {
	req := &credentialAwsBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.SecurityToken.IsNull() && !o.SecurityToken.IsUnknown() {
		inputs["security_token"] = o.SecurityToken.ValueString()
	}
	if !o.Username.IsNull() && !o.Username.IsUnknown() {
		inputs["username"] = o.Username.ValueString()
	}
	if len(inputs) > 0 {
		payload, _ := json.Marshal(inputs)
		req.Inputs = payload
	}
	return req
}

// UpdateFromApiData unfolds the AWX response back into the typed model. Secret
// fields come back as `$encrypted$` placeholders; the per-credential-type
// pre-state-set hook reconciles them against prior plan state.
func (o *credentialAwsTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))

	if inputs, ok := data["inputs"].(map[string]any); ok {
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.SecurityToken, inputs["security_token"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialAws reconciles `$encrypted$` placeholders that AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan. Data-source reads have orig==nil and skip reconciliation.
func hookCredentialAws(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialAwsTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Password.IsNull() || orig.Password.IsUnknown() {
			state.Password = types.StringNull()
		} else {
			state.Password = orig.Password
		}
		if orig.SecurityToken.IsNull() || orig.SecurityToken.IsUnknown() {
			state.SecurityToken = types.StringNull()
		} else {
			state.SecurityToken = orig.SecurityToken
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.SecurityToken, state.SecurityToken); subbed {
			state.SecurityToken = v
		}
	}
	return nil
}

// credentialAwsTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialAwsTypeLookup = framework.NewCredentialTypeLookup()

type credentialAwsResource = framework.GenericResource[credentialAwsTerraformModel, credentialAwsBodyRequestModel, *credentialAwsTerraformModel]

// NewCredentialAwsResource constructs the typed Amazon Web Services credential resource.
// The credential_type ID is resolved by namespace (aws) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialAwsResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["password"] = schema.StringAttribute{
		Description: "Secret Key",
		Required:    true,
		Sensitive:   true,
	}
	attrs["security_token"] = schema.StringAttribute{
		Description: "Security Token Service (STS) is a web service that enables you to request temporary, limited-privilege credentials for AWS Identity and Access Management (IAM) users.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Access Key",
		Required:    true,
	}
	return &credentialAwsResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_aws", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialAwsTerraformModel, credentialAwsBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Amazon Web Services` (aws) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.aws.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialAwsTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialAws,
			OnConfigure: credentialAwsTypeLookup.OnConfigure("aws"),
			MutateBody: func(plan *credentialAwsTerraformModel, body *credentialAwsBodyRequestModel) {
				body.CredentialType = credentialAwsTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialAwsTerraformModel, body *credentialAwsBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialAwsTerraformModel) {
				state.Team = types.Int64Value(plan.Team.ValueInt64())
				state.User = types.Int64Value(plan.User.ValueInt64())
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialAwsTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAws",
		},
	}
}

type credentialAwsDataSource = framework.GenericDataSource[credentialAwsTerraformModel, *credentialAwsTerraformModel]

// NewCredentialAwsDataSource constructs the typed Amazon Web Services credential data source.
func NewCredentialAwsDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["password"] = dschema.StringAttribute{
		Description: "Secret Key",
		Computed:    true,
		Sensitive:   true,
	}
	attrs["security_token"] = dschema.StringAttribute{
		Description: "Security Token Service (STS) is a web service that enables you to request temporary, limited-privilege credentials for AWS Identity and Access Management (IAM) users.",
		Computed:    true,
		Sensitive:   true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Access Key",
		Computed:    true,
	}
	return &credentialAwsDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_aws", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialAwsTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Amazon Web Services` (aws) credential by ID or name.",
				Attributes:          attrs,
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "/?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			OnConfigure:  credentialAwsTypeLookup.OnConfigure("aws"),
			Hook:         hookCredentialAws,
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAws",
		},
	}
}
