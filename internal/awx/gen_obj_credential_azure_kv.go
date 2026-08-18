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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialAzureKvTerraformModel exposes the typed AWX Microsoft Azure Key Vault
// credential (credential_azure_kv) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialAzureKvTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Client         types.String `tfsdk:"client" json:"-"`
	CloudName      types.String `tfsdk:"cloud_name" json:"-"`
	Secret         types.String `tfsdk:"secret" json:"-"`
	Tenant         types.String `tfsdk:"tenant" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
}

func (o *credentialAzureKvTerraformModel) Clone() credentialAzureKvTerraformModel {
	return *o
}

type credentialAzureKvBodyRequestModel struct {
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
func (o *credentialAzureKvTerraformModel) BodyRequest() *credentialAzureKvBodyRequestModel {
	req := &credentialAzureKvBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Client.IsNull() && !o.Client.IsUnknown() {
		inputs["client"] = o.Client.ValueString()
	}
	if !o.CloudName.IsNull() && !o.CloudName.IsUnknown() {
		inputs["cloud_name"] = o.CloudName.ValueString()
	}
	if !o.Secret.IsNull() && !o.Secret.IsUnknown() {
		inputs["secret"] = o.Secret.ValueString()
	}
	if !o.Tenant.IsNull() && !o.Tenant.IsUnknown() {
		inputs["tenant"] = o.Tenant.ValueString()
	}
	if !o.Url.IsNull() && !o.Url.IsUnknown() {
		inputs["url"] = o.Url.ValueString()
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
func (o *credentialAzureKvTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Client, inputs["client"], false))
		collect(helpers.AttrValueSetString(&o.CloudName, inputs["cloud_name"], false))
		collect(helpers.AttrValueSetString(&o.Secret, inputs["secret"], false))
		collect(helpers.AttrValueSetString(&o.Tenant, inputs["tenant"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
	}
	return diags, nil
}

// hookCredentialAzureKv reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialAzureKv(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialAzureKvTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Secret.IsNull() || orig.Secret.IsUnknown() {
			state.Secret = types.StringNull()
		} else {
			state.Secret = orig.Secret
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Secret, state.Secret); subbed {
			state.Secret = v
		}
	}
	return nil
}

// credentialAzureKvDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialAzureKvDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Client         types.String `tfsdk:"client" json:"-"`
	CloudName      types.String `tfsdk:"cloud_name" json:"-"`
	Tenant         types.String `tfsdk:"tenant" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
}

func (o *credentialAzureKvDataSourceTerraformModel) Clone() credentialAzureKvDataSourceTerraformModel {
	return *o
}

func (o *credentialAzureKvDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Client, inputs["client"], false))
		collect(helpers.AttrValueSetString(&o.CloudName, inputs["cloud_name"], false))
		collect(helpers.AttrValueSetString(&o.Tenant, inputs["tenant"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
	}
	return diags, nil
}

// credentialAzureKvTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialAzureKvTypeLookup = framework.NewCredentialTypeLookup()

type credentialAzureKvResource = framework.GenericResource[credentialAzureKvTerraformModel, credentialAzureKvBodyRequestModel, *credentialAzureKvTerraformModel]

// NewCredentialAzureKvResource constructs the typed Microsoft Azure Key Vault credential resource.
// The credential_type ID is resolved by namespace (azure_kv) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialAzureKvResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["client"] = schema.StringAttribute{
		Description: "Client ID.",
		Required:    true,
	}
	attrs["cloud_name"] = schema.StringAttribute{
		Description: "Specify which azure cloud environment to use. Allowed values: \"AzureChinaCloud\", \"AzureCloud\", \"AzureGermanCloud\", \"AzureUSGovernment\". AWX defaults this to \"AzureCloud\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Validators: []validator.String{
			stringvalidator.OneOf("AzureChinaCloud", "AzureCloud", "AzureGermanCloud", "AzureUSGovernment"),
		},
	}
	attrs["secret"] = schema.StringAttribute{
		Description: "Client Secret.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["tenant"] = schema.StringAttribute{
		Description: "Tenant ID.",
		Required:    true,
	}
	attrs["url"] = schema.StringAttribute{
		Description: "Vault URL (DNS Name).",
		Required:    true,
	}
	return &credentialAzureKvResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_azure_kv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialAzureKvTerraformModel, credentialAzureKvBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Microsoft Azure Key Vault` (azure_kv) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.azure_kv.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialAzureKvTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialAzureKv,
			OnConfigure: credentialAzureKvTypeLookup.OnConfigure("azure_kv"),
			MutateBody: func(plan *credentialAzureKvTerraformModel, body *credentialAzureKvBodyRequestModel) {
				body.CredentialType = credentialAzureKvTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialAzureKvTerraformModel, body *credentialAzureKvBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialAzureKvTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialAzureKvTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAzureKv",
		},
	}
}

type credentialAzureKvDataSource = framework.GenericDataSource[credentialAzureKvDataSourceTerraformModel, *credentialAzureKvDataSourceTerraformModel]

// NewCredentialAzureKvDataSource constructs the typed Microsoft Azure Key Vault credential data source.
func NewCredentialAzureKvDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["client"] = dschema.StringAttribute{
		Description: "Client ID.",
		Computed:    true,
	}
	attrs["cloud_name"] = dschema.StringAttribute{
		Description: "Specify which azure cloud environment to use. Allowed values: \"AzureChinaCloud\", \"AzureCloud\", \"AzureGermanCloud\", \"AzureUSGovernment\". AWX defaults this to \"AzureCloud\" when unset.",
		Computed:    true,
	}
	attrs["tenant"] = dschema.StringAttribute{
		Description: "Tenant ID.",
		Computed:    true,
	}
	attrs["url"] = dschema.StringAttribute{
		Description: "Vault URL (DNS Name).",
		Computed:    true,
	}
	return &credentialAzureKvDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_azure_kv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialAzureKvDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Microsoft Azure Key Vault` (azure_kv) credential by ID or name.",
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
			OnConfigure:  credentialAzureKvTypeLookup.OnConfigure("azure_kv"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAzureKv",
		},
	}
}
