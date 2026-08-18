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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialAwsSecretsmanagerCredentialTerraformModel exposes the typed AWX AWS Secrets Manager lookup
// credential (credential_aws_secretsmanager_credential) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialAwsSecretsmanagerCredentialTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	AwsAccessKey   types.String `tfsdk:"aws_access_key" json:"-"`
	AwsSecretKey   types.String `tfsdk:"aws_secret_key" json:"-"`
}

func (o *credentialAwsSecretsmanagerCredentialTerraformModel) Clone() credentialAwsSecretsmanagerCredentialTerraformModel {
	return *o
}

type credentialAwsSecretsmanagerCredentialBodyRequestModel struct {
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
func (o *credentialAwsSecretsmanagerCredentialTerraformModel) BodyRequest() *credentialAwsSecretsmanagerCredentialBodyRequestModel {
	req := &credentialAwsSecretsmanagerCredentialBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.AwsAccessKey.IsNull() && !o.AwsAccessKey.IsUnknown() {
		inputs["aws_access_key"] = o.AwsAccessKey.ValueString()
	}
	if !o.AwsSecretKey.IsNull() && !o.AwsSecretKey.IsUnknown() {
		inputs["aws_secret_key"] = o.AwsSecretKey.ValueString()
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
func (o *credentialAwsSecretsmanagerCredentialTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AwsAccessKey, inputs["aws_access_key"], false))
		collect(helpers.AttrValueSetString(&o.AwsSecretKey, inputs["aws_secret_key"], false))
	}
	return diags, nil
}

// hookCredentialAwsSecretsmanagerCredential reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialAwsSecretsmanagerCredential(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialAwsSecretsmanagerCredentialTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.AwsSecretKey.IsNull() || orig.AwsSecretKey.IsUnknown() {
			state.AwsSecretKey = types.StringNull()
		} else {
			state.AwsSecretKey = orig.AwsSecretKey
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.AwsSecretKey, state.AwsSecretKey); subbed {
			state.AwsSecretKey = v
		}
	}
	return nil
}

// credentialAwsSecretsmanagerCredentialDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialAwsSecretsmanagerCredentialDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	AwsAccessKey   types.String `tfsdk:"aws_access_key" json:"-"`
}

func (o *credentialAwsSecretsmanagerCredentialDataSourceTerraformModel) Clone() credentialAwsSecretsmanagerCredentialDataSourceTerraformModel {
	return *o
}

func (o *credentialAwsSecretsmanagerCredentialDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AwsAccessKey, inputs["aws_access_key"], false))
	}
	return diags, nil
}

// credentialAwsSecretsmanagerCredentialTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialAwsSecretsmanagerCredentialTypeLookup = framework.NewCredentialTypeLookup()

type credentialAwsSecretsmanagerCredentialResource = framework.GenericResource[credentialAwsSecretsmanagerCredentialTerraformModel, credentialAwsSecretsmanagerCredentialBodyRequestModel, *credentialAwsSecretsmanagerCredentialTerraformModel]

// NewCredentialAwsSecretsmanagerCredentialResource constructs the typed AWS Secrets Manager lookup credential resource.
// The credential_type ID is resolved by namespace (aws_secretsmanager_credential) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialAwsSecretsmanagerCredentialResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["aws_access_key"] = schema.StringAttribute{
		Description: "AWS Access Key.",
		Required:    true,
	}
	attrs["aws_secret_key"] = schema.StringAttribute{
		Description: "AWS Secret Key.",
		Required:    true,
		Sensitive:   true,
	}
	return &credentialAwsSecretsmanagerCredentialResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_aws_secretsmanager_credential", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialAwsSecretsmanagerCredentialTerraformModel, credentialAwsSecretsmanagerCredentialBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `AWS Secrets Manager lookup` (aws_secretsmanager_credential) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.aws_secretsmanager_credential.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialAwsSecretsmanagerCredentialTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialAwsSecretsmanagerCredential,
			OnConfigure: credentialAwsSecretsmanagerCredentialTypeLookup.OnConfigure("aws_secretsmanager_credential"),
			MutateBody: func(plan *credentialAwsSecretsmanagerCredentialTerraformModel, body *credentialAwsSecretsmanagerCredentialBodyRequestModel) {
				body.CredentialType = credentialAwsSecretsmanagerCredentialTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialAwsSecretsmanagerCredentialTerraformModel, body *credentialAwsSecretsmanagerCredentialBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialAwsSecretsmanagerCredentialTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialAwsSecretsmanagerCredentialTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAwsSecretsmanagerCredential",
		},
	}
}

type credentialAwsSecretsmanagerCredentialDataSource = framework.GenericDataSource[credentialAwsSecretsmanagerCredentialDataSourceTerraformModel, *credentialAwsSecretsmanagerCredentialDataSourceTerraformModel]

// NewCredentialAwsSecretsmanagerCredentialDataSource constructs the typed AWS Secrets Manager lookup credential data source.
func NewCredentialAwsSecretsmanagerCredentialDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["aws_access_key"] = dschema.StringAttribute{
		Description: "AWS Access Key.",
		Computed:    true,
	}
	return &credentialAwsSecretsmanagerCredentialDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_aws_secretsmanager_credential", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialAwsSecretsmanagerCredentialDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `AWS Secrets Manager lookup` (aws_secretsmanager_credential) credential by ID or name.",
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
			OnConfigure:  credentialAwsSecretsmanagerCredentialTypeLookup.OnConfigure("aws_secretsmanager_credential"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAwsSecretsmanagerCredential",
		},
	}
}
