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

// credentialAzureRmTerraformModel exposes the typed AWX Microsoft Azure Resource Manager
// credential (credential_azure_rm) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialAzureRmTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	Team             types.Int64  `tfsdk:"team" json:"team"`
	User             types.Int64  `tfsdk:"user" json:"user"`
	Kind             types.String `tfsdk:"kind" json:"kind"`
	Managed          types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType   types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Client           types.String `tfsdk:"client" json:"-"`
	CloudEnvironment types.String `tfsdk:"cloud_environment" json:"-"`
	Password         types.String `tfsdk:"password" json:"-"`
	Secret           types.String `tfsdk:"secret" json:"-"`
	Subscription     types.String `tfsdk:"subscription" json:"-"`
	Tenant           types.String `tfsdk:"tenant" json:"-"`
	Username         types.String `tfsdk:"username" json:"-"`
}

func (o *credentialAzureRmTerraformModel) Clone() credentialAzureRmTerraformModel {
	return *o
}

type credentialAzureRmBodyRequestModel struct {
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
func (o *credentialAzureRmTerraformModel) BodyRequest() *credentialAzureRmBodyRequestModel {
	req := &credentialAzureRmBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Client.IsNull() && !o.Client.IsUnknown() {
		inputs["client"] = o.Client.ValueString()
	}
	if !o.CloudEnvironment.IsNull() && !o.CloudEnvironment.IsUnknown() {
		inputs["cloud_environment"] = o.CloudEnvironment.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.Secret.IsNull() && !o.Secret.IsUnknown() {
		inputs["secret"] = o.Secret.ValueString()
	}
	if !o.Subscription.IsNull() && !o.Subscription.IsUnknown() {
		inputs["subscription"] = o.Subscription.ValueString()
	}
	if !o.Tenant.IsNull() && !o.Tenant.IsUnknown() {
		inputs["tenant"] = o.Tenant.ValueString()
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
func (o *credentialAzureRmTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.CloudEnvironment, inputs["cloud_environment"], false))
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.Secret, inputs["secret"], false))
		collect(helpers.AttrValueSetString(&o.Subscription, inputs["subscription"], false))
		collect(helpers.AttrValueSetString(&o.Tenant, inputs["tenant"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialAzureRm reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialAzureRm(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialAzureRmTerraformModel) error {
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
		if orig.Secret.IsNull() || orig.Secret.IsUnknown() {
			state.Secret = types.StringNull()
		} else {
			state.Secret = orig.Secret
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.Secret, state.Secret); subbed {
			state.Secret = v
		}
	}
	return nil
}

// credentialAzureRmDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialAzureRmDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	Team             types.Int64  `tfsdk:"team" json:"team"`
	User             types.Int64  `tfsdk:"user" json:"user"`
	Kind             types.String `tfsdk:"kind" json:"kind"`
	Managed          types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType   types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Client           types.String `tfsdk:"client" json:"-"`
	CloudEnvironment types.String `tfsdk:"cloud_environment" json:"-"`
	Subscription     types.String `tfsdk:"subscription" json:"-"`
	Tenant           types.String `tfsdk:"tenant" json:"-"`
	Username         types.String `tfsdk:"username" json:"-"`
}

func (o *credentialAzureRmDataSourceTerraformModel) Clone() credentialAzureRmDataSourceTerraformModel {
	return *o
}

func (o *credentialAzureRmDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.CloudEnvironment, inputs["cloud_environment"], false))
		collect(helpers.AttrValueSetString(&o.Subscription, inputs["subscription"], false))
		collect(helpers.AttrValueSetString(&o.Tenant, inputs["tenant"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialAzureRmTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialAzureRmTypeLookup = framework.NewCredentialTypeLookup()

type credentialAzureRmResource = framework.GenericResource[credentialAzureRmTerraformModel, credentialAzureRmBodyRequestModel, *credentialAzureRmTerraformModel]

// NewCredentialAzureRmResource constructs the typed Microsoft Azure Resource Manager credential resource.
// The credential_type ID is resolved by namespace (azure_rm) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialAzureRmResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["client"] = schema.StringAttribute{
		Description: "Client ID.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["cloud_environment"] = schema.StringAttribute{
		Description: "Environment variable AZURE_CLOUD_ENVIRONMENT when using Azure GovCloud or Azure stack.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["password"] = schema.StringAttribute{
		Description: "Password.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["secret"] = schema.StringAttribute{
		Description: "Client Secret.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["subscription"] = schema.StringAttribute{
		Description: "Subscription ID is an Azure construct, which is mapped to a username.",
		Required:    true,
	}
	attrs["tenant"] = schema.StringAttribute{
		Description: "Tenant ID.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	return &credentialAzureRmResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_azure_rm", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialAzureRmTerraformModel, credentialAzureRmBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Microsoft Azure Resource Manager` (azure_rm) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.azure_rm.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialAzureRmTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialAzureRm,
			OnConfigure: credentialAzureRmTypeLookup.OnConfigure("azure_rm"),
			MutateBody: func(plan *credentialAzureRmTerraformModel, body *credentialAzureRmBodyRequestModel) {
				body.CredentialType = credentialAzureRmTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialAzureRmTerraformModel, body *credentialAzureRmBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialAzureRmTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialAzureRmTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAzureRm",
		},
	}
}

type credentialAzureRmDataSource = framework.GenericDataSource[credentialAzureRmDataSourceTerraformModel, *credentialAzureRmDataSourceTerraformModel]

// NewCredentialAzureRmDataSource constructs the typed Microsoft Azure Resource Manager credential data source.
func NewCredentialAzureRmDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["client"] = dschema.StringAttribute{
		Description: "Client ID.",
		Computed:    true,
	}
	attrs["cloud_environment"] = dschema.StringAttribute{
		Description: "Environment variable AZURE_CLOUD_ENVIRONMENT when using Azure GovCloud or Azure stack.",
		Computed:    true,
	}
	attrs["subscription"] = dschema.StringAttribute{
		Description: "Subscription ID is an Azure construct, which is mapped to a username.",
		Computed:    true,
	}
	attrs["tenant"] = dschema.StringAttribute{
		Description: "Tenant ID.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	return &credentialAzureRmDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_azure_rm", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialAzureRmDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Microsoft Azure Resource Manager` (azure_rm) credential by ID or name.",
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
			OnConfigure:  credentialAzureRmTypeLookup.OnConfigure("azure_rm"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAzureRm",
		},
	}
}
