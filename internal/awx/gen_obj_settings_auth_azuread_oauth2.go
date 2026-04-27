package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthAzureAdoauth2TerraformModel struct {
	SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL     types.String `tfsdk:"social_auth_azuread_oauth2_callback_url" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_KEY              types.String `tfsdk:"social_auth_azuread_oauth2_key" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP types.String `tfsdk:"social_auth_azuread_oauth2_organization_map" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET           types.String `tfsdk:"social_auth_azuread_oauth2_secret" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP         types.String `tfsdk:"social_auth_azuread_oauth2_team_map" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"`
}

func (o *settingsAuthAzureAdoauth2TerraformModel) Clone() settingsAuthAzureAdoauth2TerraformModel {
	return *o
}

func (o *settingsAuthAzureAdoauth2TerraformModel) BodyRequest() *settingsAuthAzureAdoauth2BodyRequestModel {
	var req settingsAuthAzureAdoauth2BodyRequestModel
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY = o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET = o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP.ValueString())
	return &req
}

func (o *settingsAuthAzureAdoauth2TerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL, data["SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY, data["SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP, data["SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET, data["SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP, data["SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"], false))
	return diags, nil
}

type settingsAuthAzureAdoauth2BodyRequestModel struct {
	SOCIAL_AUTH_AZUREAD_OAUTH2_KEY              string          `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY,omitempty"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET           string          `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET,omitempty"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP,omitempty"`
}

type settingsAuthAzureAdoauth2Resource = framework.GenericResource[settingsAuthAzureAdoauth2TerraformModel, settingsAuthAzureAdoauth2BodyRequestModel, *settingsAuthAzureAdoauth2TerraformModel]

// NewSettingsAuthAzureADOauth2Resource is a helper function to simplify the provider implementation.
func NewSettingsAuthAzureADOauth2Resource() resource.Resource {
	return &settingsAuthAzureAdoauth2Resource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_azuread_oauth2", Endpoint: "/api/v2/settings/azuread-oauth2/"}},
		Cfg: framework.ResourceCfg[settingsAuthAzureAdoauth2TerraformModel, settingsAuthAzureAdoauth2BodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_azuread_oauth2_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your Azure AD application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_azuread_oauth2_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_azuread_oauth2_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your Azure AD application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_azuread_oauth2_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_azuread_oauth2_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail. ",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsAuthAzureADOauth2,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthAzureADOauth2",
		},
	}
}

type settingsAuthAzureAdoauth2DataSource = framework.GenericDataSource[settingsAuthAzureAdoauth2TerraformModel, *settingsAuthAzureAdoauth2TerraformModel]

// NewSettingsAuthAzureADOauth2DataSource is a helper function to instantiate the SettingsAuthAzureADOauth2 data source.
func NewSettingsAuthAzureADOauth2DataSource() datasource.DataSource {
	return &settingsAuthAzureAdoauth2DataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_azuread_oauth2", Endpoint: "/api/v2/settings/azuread-oauth2/"}},
		Cfg: framework.DataSourceCfg[settingsAuthAzureAdoauth2TerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_azuread_oauth2_callback_url": dschema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail. ",
						Computed:    true,
					},
					"social_auth_azuread_oauth2_key": dschema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your Azure AD application.",
						Computed:    true,
					},
					"social_auth_azuread_oauth2_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_azuread_oauth2_secret": dschema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your Azure AD application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_azuread_oauth2_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthAzureADOauth2,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthAzureADOauth2",
		},
	}
}
