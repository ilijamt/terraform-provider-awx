package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthGithubEnterpriseTeamDataSource = framework.GenericDataSource[settingsAuthGithubEnterpriseTeamTerraformModel, *settingsAuthGithubEnterpriseTeamTerraformModel]

// NewSettingsAuthGithubEnterpriseTeamDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseTeam data source.
func NewSettingsAuthGithubEnterpriseTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseTeamDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise_team", Endpoint: "/api/v2/settings/github-enterprise-team/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubEnterpriseTeamTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_enterprise_team_api_url": schema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_callback_url": schema.StringAttribute{
						Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_id": schema.StringAttribute{
						Description: "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_enterprise_team_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_team_url": schema.StringAttribute{
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
