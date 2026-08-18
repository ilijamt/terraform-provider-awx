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

// credentialGalaxyApiTokenTerraformModel exposes the typed AWX Ansible Galaxy/Automation Hub API Token
// credential (credential_galaxy_api_token) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialGalaxyApiTokenTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	AuthUrl        types.String `tfsdk:"auth_url" json:"-"`
	Token          types.String `tfsdk:"token" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
}

func (o *credentialGalaxyApiTokenTerraformModel) Clone() credentialGalaxyApiTokenTerraformModel {
	return *o
}

type credentialGalaxyApiTokenBodyRequestModel struct {
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
func (o *credentialGalaxyApiTokenTerraformModel) BodyRequest() *credentialGalaxyApiTokenBodyRequestModel {
	req := &credentialGalaxyApiTokenBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.AuthUrl.IsNull() && !o.AuthUrl.IsUnknown() {
		inputs["auth_url"] = o.AuthUrl.ValueString()
	}
	if !o.Token.IsNull() && !o.Token.IsUnknown() {
		inputs["token"] = o.Token.ValueString()
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
func (o *credentialGalaxyApiTokenTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AuthUrl, inputs["auth_url"], false))
		collect(helpers.AttrValueSetString(&o.Token, inputs["token"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
	}
	return diags, nil
}

// hookCredentialGalaxyApiToken reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialGalaxyApiToken(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialGalaxyApiTokenTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Token.IsNull() || orig.Token.IsUnknown() {
			state.Token = types.StringNull()
		} else {
			state.Token = orig.Token
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Token, state.Token); subbed {
			state.Token = v
		}
	}
	return nil
}

// credentialGalaxyApiTokenDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialGalaxyApiTokenDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	AuthUrl        types.String `tfsdk:"auth_url" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
}

func (o *credentialGalaxyApiTokenDataSourceTerraformModel) Clone() credentialGalaxyApiTokenDataSourceTerraformModel {
	return *o
}

func (o *credentialGalaxyApiTokenDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AuthUrl, inputs["auth_url"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
	}
	return diags, nil
}

// credentialGalaxyApiTokenTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialGalaxyApiTokenTypeLookup = framework.NewCredentialTypeLookup()

type credentialGalaxyApiTokenResource = framework.GenericResource[credentialGalaxyApiTokenTerraformModel, credentialGalaxyApiTokenBodyRequestModel, *credentialGalaxyApiTokenTerraformModel]

// NewCredentialGalaxyApiTokenResource constructs the typed Ansible Galaxy/Automation Hub API Token credential resource.
// The credential_type ID is resolved by namespace (galaxy_api_token) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialGalaxyApiTokenResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["auth_url"] = schema.StringAttribute{
		Description: "The URL of a Keycloak server token_endpoint, if using SSO auth.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["token"] = schema.StringAttribute{
		Description: "A token to use for authentication against the Galaxy instance.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["url"] = schema.StringAttribute{
		Description: "The URL of the Galaxy instance to connect to.",
		Required:    true,
	}
	return &credentialGalaxyApiTokenResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_galaxy_api_token", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialGalaxyApiTokenTerraformModel, credentialGalaxyApiTokenBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Ansible Galaxy/Automation Hub API Token` (galaxy_api_token) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.galaxy_api_token.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialGalaxyApiTokenTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialGalaxyApiToken,
			OnConfigure: credentialGalaxyApiTokenTypeLookup.OnConfigure("galaxy_api_token"),
			MutateBody: func(plan *credentialGalaxyApiTokenTerraformModel, body *credentialGalaxyApiTokenBodyRequestModel) {
				body.CredentialType = credentialGalaxyApiTokenTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialGalaxyApiTokenTerraformModel, body *credentialGalaxyApiTokenBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialGalaxyApiTokenTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialGalaxyApiTokenTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialGalaxyApiToken",
		},
	}
}

type credentialGalaxyApiTokenDataSource = framework.GenericDataSource[credentialGalaxyApiTokenDataSourceTerraformModel, *credentialGalaxyApiTokenDataSourceTerraformModel]

// NewCredentialGalaxyApiTokenDataSource constructs the typed Ansible Galaxy/Automation Hub API Token credential data source.
func NewCredentialGalaxyApiTokenDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["auth_url"] = dschema.StringAttribute{
		Description: "The URL of a Keycloak server token_endpoint, if using SSO auth.",
		Computed:    true,
	}
	attrs["url"] = dschema.StringAttribute{
		Description: "The URL of the Galaxy instance to connect to.",
		Computed:    true,
	}
	return &credentialGalaxyApiTokenDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_galaxy_api_token", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialGalaxyApiTokenDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Ansible Galaxy/Automation Hub API Token` (galaxy_api_token) credential by ID or name.",
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
			OnConfigure:  credentialGalaxyApiTokenTypeLookup.OnConfigure("galaxy_api_token"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialGalaxyApiToken",
		},
	}
}
