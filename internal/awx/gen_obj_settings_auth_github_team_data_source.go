package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthGithubTeamDataSource = framework.GenericDataSource[settingsAuthGithubTeamTerraformModel, *settingsAuthGithubTeamTerraformModel]

// NewSettingsAuthGithubTeamDataSource is a helper function to instantiate the SettingsAuthGithubTeam data source.
func NewSettingsAuthGithubTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubTeamDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_team", Endpoint: "/api/v2/settings/github-team/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubTeamTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"social_auth_github_team_callback_url": schema.StringAttribute{
						Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_team_id": schema.StringAttribute{
						Description: "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_team_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_team_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_team_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_team_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
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
