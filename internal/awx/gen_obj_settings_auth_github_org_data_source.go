package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthGithubOrgDataSource = framework.GenericDataSource[settingsAuthGithubOrgTerraformModel, *settingsAuthGithubOrgTerraformModel]

// NewSettingsAuthGithubOrgDataSource is a helper function to instantiate the SettingsAuthGithubOrg data source.
func NewSettingsAuthGithubOrgDataSource() datasource.DataSource {
	return &settingsAuthGithubOrgDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_org", Endpoint: "/api/v2/settings/github-org/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubOrgTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_org_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
					},
					"social_auth_github_org_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
						Computed:    true,
					},
					"social_auth_github_org_name": schema.StringAttribute{
						Description: "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/.",
						Computed:    true,
					},
					"social_auth_github_org_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_org_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_org_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithubOrg,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubOrg",
		},
	}
}
