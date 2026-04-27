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

type settingsAuthGithubEnterpriseTeamTerraformModel struct {
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL          types.String `tfsdk:"social_auth_github_enterprise_team_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL     types.String `tfsdk:"social_auth_github_enterprise_team_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID               types.String `tfsdk:"social_auth_github_enterprise_team_id" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY              types.String `tfsdk:"social_auth_github_enterprise_team_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_team_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET           types.String `tfsdk:"social_auth_github_enterprise_team_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP         types.String `tfsdk:"social_auth_github_enterprise_team_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL              types.String `tfsdk:"social_auth_github_enterprise_team_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"`
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) Clone() settingsAuthGithubEnterpriseTeamTerraformModel {
	return *o
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) BodyRequest() *settingsAuthGithubEnterpriseTeamBodyRequestModel {
	var req settingsAuthGithubEnterpriseTeamBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL.ValueString()
	return &req
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"], false))
	return diags, nil
}

type settingsAuthGithubEnterpriseTeamBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL          string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID               string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY              string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL              string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL,omitempty"`
}

type settingsAuthGithubEnterpriseTeamResource = framework.GenericResource[settingsAuthGithubEnterpriseTeamTerraformModel, settingsAuthGithubEnterpriseTeamBodyRequestModel, *settingsAuthGithubEnterpriseTeamTerraformModel]

// NewSettingsAuthGithubEnterpriseTeamResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubEnterpriseTeamResource() resource.Resource {
	return &settingsAuthGithubEnterpriseTeamResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise_team", Endpoint: "/api/v2/settings/github-enterprise-team/"}},
		Cfg: framework.ResourceCfg[settingsAuthGithubEnterpriseTeamTerraformModel, settingsAuthGithubEnterpriseTeamBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_enterprise_team_api_url": schema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_id": schema.StringAttribute{
						Description: "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_url": schema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_callback_url": schema.StringAttribute{
						Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsAuthGithubEnterpriseTeam,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterpriseTeam",
		},
	}
}

type settingsAuthGithubEnterpriseTeamDataSource = framework.GenericDataSource[settingsAuthGithubEnterpriseTeamTerraformModel, *settingsAuthGithubEnterpriseTeamTerraformModel]

// NewSettingsAuthGithubEnterpriseTeamDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseTeam data source.
func NewSettingsAuthGithubEnterpriseTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseTeamDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise_team", Endpoint: "/api/v2/settings/github-enterprise-team/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubEnterpriseTeamTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_github_enterprise_team_api_url": dschema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_callback_url": dschema.StringAttribute{
						Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_id": dschema.StringAttribute{
						Description: "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_key": dschema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_secret": dschema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_enterprise_team_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_url": dschema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithubEnterpriseTeam,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterpriseTeam",
		},
	}
}
