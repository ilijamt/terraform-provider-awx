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

// credentialTerraformTerraformModel exposes the typed AWX Terraform backend configuration
// credential (credential_terraform) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialTerraformTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Configuration  types.String `tfsdk:"configuration" json:"-"`
	GceCredentials types.String `tfsdk:"gce_credentials" json:"-"`
}

func (o *credentialTerraformTerraformModel) Clone() credentialTerraformTerraformModel {
	return *o
}

type credentialTerraformBodyRequestModel struct {
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
func (o *credentialTerraformTerraformModel) BodyRequest() *credentialTerraformBodyRequestModel {
	req := &credentialTerraformBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Configuration.IsNull() && !o.Configuration.IsUnknown() {
		inputs["configuration"] = o.Configuration.ValueString()
	}
	if !o.GceCredentials.IsNull() && !o.GceCredentials.IsUnknown() {
		inputs["gce_credentials"] = o.GceCredentials.ValueString()
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
func (o *credentialTerraformTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Configuration, inputs["configuration"], false))
		collect(helpers.AttrValueSetString(&o.GceCredentials, inputs["gce_credentials"], false))
	}
	return diags, nil
}

// hookCredentialTerraform reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialTerraform(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialTerraformTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Configuration.IsNull() || orig.Configuration.IsUnknown() {
			state.Configuration = types.StringNull()
		} else {
			state.Configuration = orig.Configuration
		}
		if orig.GceCredentials.IsNull() || orig.GceCredentials.IsUnknown() {
			state.GceCredentials = types.StringNull()
		} else {
			state.GceCredentials = orig.GceCredentials
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Configuration, state.Configuration); subbed {
			state.Configuration = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.GceCredentials, state.GceCredentials); subbed {
			state.GceCredentials = v
		}
	}
	return nil
}

// credentialTerraformDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialTerraformDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
}

func (o *credentialTerraformDataSourceTerraformModel) Clone() credentialTerraformDataSourceTerraformModel {
	return *o
}

func (o *credentialTerraformDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
	return diags, nil
}

// credentialTerraformTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialTerraformTypeLookup = framework.NewCredentialTypeLookup()

type credentialTerraformResource = framework.GenericResource[credentialTerraformTerraformModel, credentialTerraformBodyRequestModel, *credentialTerraformTerraformModel]

// NewCredentialTerraformResource constructs the typed Terraform backend configuration credential resource.
// The credential_type ID is resolved by namespace (terraform) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialTerraformResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["configuration"] = schema.StringAttribute{
		Description: "Terraform backend config as Hashicorp configuration language.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["gce_credentials"] = schema.StringAttribute{
		Description: "Google Cloud Platform account credentials in JSON format.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	return &credentialTerraformResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_terraform", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialTerraformTerraformModel, credentialTerraformBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Terraform backend configuration` (terraform) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.terraform.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialTerraformTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialTerraform,
			OnConfigure: credentialTerraformTypeLookup.OnConfigure("terraform"),
			MutateBody: func(plan *credentialTerraformTerraformModel, body *credentialTerraformBodyRequestModel) {
				body.CredentialType = credentialTerraformTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialTerraformTerraformModel, body *credentialTerraformBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialTerraformTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialTerraformTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialTerraform",
		},
	}
}

type credentialTerraformDataSource = framework.GenericDataSource[credentialTerraformDataSourceTerraformModel, *credentialTerraformDataSourceTerraformModel]

// NewCredentialTerraformDataSource constructs the typed Terraform backend configuration credential data source.
func NewCredentialTerraformDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	return &credentialTerraformDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_terraform", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialTerraformDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Terraform backend configuration` (terraform) credential by ID or name.",
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
			OnConfigure:  credentialTerraformTypeLookup.OnConfigure("terraform"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialTerraform",
		},
	}
}
