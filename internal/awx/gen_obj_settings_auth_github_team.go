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

type settingsAuthGithubTeamTerraformModel struct {
	SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL     types.String `tfsdk:"social_auth_github_team_callback_url" json:"SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_TEAM_ID               types.String `tfsdk:"social_auth_github_team_id" json:"SOCIAL_AUTH_GITHUB_TEAM_ID"`
	SOCIAL_AUTH_GITHUB_TEAM_KEY              types.String `tfsdk:"social_auth_github_team_key" json:"SOCIAL_AUTH_GITHUB_TEAM_KEY"`
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_team_organization_map" json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_TEAM_SECRET           types.String `tfsdk:"social_auth_github_team_secret" json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET"`
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP         types.String `tfsdk:"social_auth_github_team_team_map" json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"`
}

func (o *settingsAuthGithubTeamTerraformModel) Clone() settingsAuthGithubTeamTerraformModel {
	return *o
}

func (o *settingsAuthGithubTeamTerraformModel) BodyRequest() *settingsAuthGithubTeamBodyRequestModel {
	var req settingsAuthGithubTeamBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_TEAM_ID = o.SOCIAL_AUTH_GITHUB_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP.ValueString())
	return &req
}

func (o *settingsAuthGithubTeamTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_ID, data["SOCIAL_AUTH_GITHUB_TEAM_ID"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_KEY, data["SOCIAL_AUTH_GITHUB_TEAM_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_SECRET, data["SOCIAL_AUTH_GITHUB_TEAM_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"], false))
	return diags, nil
}

type settingsAuthGithubTeamBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_TEAM_ID               string          `json:"SOCIAL_AUTH_GITHUB_TEAM_ID,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_KEY              string          `json:"SOCIAL_AUTH_GITHUB_TEAM_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP,omitempty"`
}

type settingsAuthGithubTeamResource = framework.GenericResource[settingsAuthGithubTeamTerraformModel, settingsAuthGithubTeamBodyRequestModel, *settingsAuthGithubTeamTerraformModel]

// NewSettingsAuthGithubTeamResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubTeamResource() resource.Resource {
	return &settingsAuthGithubTeamResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_team", Endpoint: "/api/v2/settings/github-team/"}},
		Cfg: framework.ResourceCfg[settingsAuthGithubTeamTerraformModel, settingsAuthGithubTeamBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_team_id": schema.StringAttribute{
						Description: "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_team_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_team_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_team_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_team_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_team_callback_url": schema.StringAttribute{
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
			Hook:         hookSettingsAuthGithubTeam,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubTeam",
		},
	}
}

type settingsAuthGithubTeamDataSource = framework.GenericDataSource[settingsAuthGithubTeamTerraformModel, *settingsAuthGithubTeamTerraformModel]

// NewSettingsAuthGithubTeamDataSource is a helper function to instantiate the SettingsAuthGithubTeam data source.
func NewSettingsAuthGithubTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubTeamDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_team", Endpoint: "/api/v2/settings/github-team/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubTeamTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_github_team_callback_url": dschema.StringAttribute{
						Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
						Computed:    true,
					},
					"social_auth_github_team_id": dschema.StringAttribute{
						Description: "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
						Computed:    true,
					},
					"social_auth_github_team_key": dschema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
						Computed:    true,
					},
					"social_auth_github_team_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_team_secret": dschema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_team_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithubTeam,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubTeam",
		},
	}
}
