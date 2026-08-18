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

// credentialCentrifyVaultKvTerraformModel exposes the typed AWX Centrify Vault Credential Provider Lookup
// credential (credential_centrify_vault_kv) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialCentrifyVaultKvTerraformModel struct {
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	Name               types.String `tfsdk:"name" json:"name"`
	Description        types.String `tfsdk:"description" json:"description"`
	Organization       types.Int64  `tfsdk:"organization" json:"organization"`
	Team               types.Int64  `tfsdk:"team" json:"team"`
	User               types.Int64  `tfsdk:"user" json:"user"`
	Kind               types.String `tfsdk:"kind" json:"kind"`
	Managed            types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType     types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	ClientId           types.String `tfsdk:"client_id" json:"-"`
	ClientPassword     types.String `tfsdk:"client_password" json:"-"`
	OauthApplicationId types.String `tfsdk:"oauth_application_id" json:"-"`
	OauthScope         types.String `tfsdk:"oauth_scope" json:"-"`
	Url                types.String `tfsdk:"url" json:"-"`
}

func (o *credentialCentrifyVaultKvTerraformModel) Clone() credentialCentrifyVaultKvTerraformModel {
	return *o
}

type credentialCentrifyVaultKvBodyRequestModel struct {
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
func (o *credentialCentrifyVaultKvTerraformModel) BodyRequest() *credentialCentrifyVaultKvBodyRequestModel {
	req := &credentialCentrifyVaultKvBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.ClientId.IsNull() && !o.ClientId.IsUnknown() {
		inputs["client_id"] = o.ClientId.ValueString()
	}
	if !o.ClientPassword.IsNull() && !o.ClientPassword.IsUnknown() {
		inputs["client_password"] = o.ClientPassword.ValueString()
	}
	if !o.OauthApplicationId.IsNull() && !o.OauthApplicationId.IsUnknown() {
		inputs["oauth_application_id"] = o.OauthApplicationId.ValueString()
	}
	if !o.OauthScope.IsNull() && !o.OauthScope.IsUnknown() {
		inputs["oauth_scope"] = o.OauthScope.ValueString()
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
func (o *credentialCentrifyVaultKvTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.ClientId, inputs["client_id"], false))
		collect(helpers.AttrValueSetString(&o.ClientPassword, inputs["client_password"], false))
		collect(helpers.AttrValueSetString(&o.OauthApplicationId, inputs["oauth_application_id"], false))
		collect(helpers.AttrValueSetString(&o.OauthScope, inputs["oauth_scope"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
	}
	return diags, nil
}

// hookCredentialCentrifyVaultKv reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialCentrifyVaultKv(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialCentrifyVaultKvTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.ClientPassword.IsNull() || orig.ClientPassword.IsUnknown() {
			state.ClientPassword = types.StringNull()
		} else {
			state.ClientPassword = orig.ClientPassword
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.ClientPassword, state.ClientPassword); subbed {
			state.ClientPassword = v
		}
	}
	return nil
}

// credentialCentrifyVaultKvDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialCentrifyVaultKvDataSourceTerraformModel struct {
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	Name               types.String `tfsdk:"name" json:"name"`
	Description        types.String `tfsdk:"description" json:"description"`
	Organization       types.Int64  `tfsdk:"organization" json:"organization"`
	Team               types.Int64  `tfsdk:"team" json:"team"`
	User               types.Int64  `tfsdk:"user" json:"user"`
	Kind               types.String `tfsdk:"kind" json:"kind"`
	Managed            types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType     types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	ClientId           types.String `tfsdk:"client_id" json:"-"`
	OauthApplicationId types.String `tfsdk:"oauth_application_id" json:"-"`
	OauthScope         types.String `tfsdk:"oauth_scope" json:"-"`
	Url                types.String `tfsdk:"url" json:"-"`
}

func (o *credentialCentrifyVaultKvDataSourceTerraformModel) Clone() credentialCentrifyVaultKvDataSourceTerraformModel {
	return *o
}

func (o *credentialCentrifyVaultKvDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.ClientId, inputs["client_id"], false))
		collect(helpers.AttrValueSetString(&o.OauthApplicationId, inputs["oauth_application_id"], false))
		collect(helpers.AttrValueSetString(&o.OauthScope, inputs["oauth_scope"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
	}
	return diags, nil
}

// credentialCentrifyVaultKvTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialCentrifyVaultKvTypeLookup = framework.NewCredentialTypeLookup()

type credentialCentrifyVaultKvResource = framework.GenericResource[credentialCentrifyVaultKvTerraformModel, credentialCentrifyVaultKvBodyRequestModel, *credentialCentrifyVaultKvTerraformModel]

// NewCredentialCentrifyVaultKvResource constructs the typed Centrify Vault Credential Provider Lookup credential resource.
// The credential_type ID is resolved by namespace (centrify_vault_kv) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialCentrifyVaultKvResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["client_id"] = schema.StringAttribute{
		Description: "Centrify API User, having necessary permissions as mentioned in support doc.",
		Required:    true,
	}
	attrs["client_password"] = schema.StringAttribute{
		Description: "Password of Centrify API User with necessary permissions.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["oauth_application_id"] = schema.StringAttribute{
		Description: "Application ID of the configured OAuth2 Client (defaults to 'awx'). AWX defaults this to \"awx\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["oauth_scope"] = schema.StringAttribute{
		Description: "Scope of the configured OAuth2 Client (defaults to 'awx'). AWX defaults this to \"awx\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["url"] = schema.StringAttribute{
		Description: "Centrify Tenant URL.",
		Required:    true,
	}
	return &credentialCentrifyVaultKvResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_centrify_vault_kv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialCentrifyVaultKvTerraformModel, credentialCentrifyVaultKvBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Centrify Vault Credential Provider Lookup` (centrify_vault_kv) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.centrify_vault_kv.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialCentrifyVaultKvTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialCentrifyVaultKv,
			OnConfigure: credentialCentrifyVaultKvTypeLookup.OnConfigure("centrify_vault_kv"),
			MutateBody: func(plan *credentialCentrifyVaultKvTerraformModel, body *credentialCentrifyVaultKvBodyRequestModel) {
				body.CredentialType = credentialCentrifyVaultKvTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialCentrifyVaultKvTerraformModel, body *credentialCentrifyVaultKvBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialCentrifyVaultKvTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialCentrifyVaultKvTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialCentrifyVaultKv",
		},
	}
}

type credentialCentrifyVaultKvDataSource = framework.GenericDataSource[credentialCentrifyVaultKvDataSourceTerraformModel, *credentialCentrifyVaultKvDataSourceTerraformModel]

// NewCredentialCentrifyVaultKvDataSource constructs the typed Centrify Vault Credential Provider Lookup credential data source.
func NewCredentialCentrifyVaultKvDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["client_id"] = dschema.StringAttribute{
		Description: "Centrify API User, having necessary permissions as mentioned in support doc.",
		Computed:    true,
	}
	attrs["oauth_application_id"] = dschema.StringAttribute{
		Description: "Application ID of the configured OAuth2 Client (defaults to 'awx'). AWX defaults this to \"awx\" when unset.",
		Computed:    true,
	}
	attrs["oauth_scope"] = dschema.StringAttribute{
		Description: "Scope of the configured OAuth2 Client (defaults to 'awx'). AWX defaults this to \"awx\" when unset.",
		Computed:    true,
	}
	attrs["url"] = dschema.StringAttribute{
		Description: "Centrify Tenant URL.",
		Computed:    true,
	}
	return &credentialCentrifyVaultKvDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_centrify_vault_kv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialCentrifyVaultKvDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Centrify Vault Credential Provider Lookup` (centrify_vault_kv) credential by ID or name.",
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
			OnConfigure:  credentialCentrifyVaultKvTypeLookup.OnConfigure("centrify_vault_kv"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialCentrifyVaultKv",
		},
	}
}
