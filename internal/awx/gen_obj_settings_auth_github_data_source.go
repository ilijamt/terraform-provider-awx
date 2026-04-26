package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthGithubDataSource = framework.GenericDataSource[settingsAuthGithubTerraformModel, *settingsAuthGithubTerraformModel]

// NewSettingsAuthGithubDataSource is a helper function to instantiate the SettingsAuthGithub data source.
func NewSettingsAuthGithubDataSource() datasource.DataSource {
	return &settingsAuthGithubDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github", Endpoint: "/api/v2/settings/github/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"social_auth_github_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub developer application.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub developer application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithub,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithub",
		},
	}
}
