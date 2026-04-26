package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthGithubEnterpriseDataSource = framework.GenericDataSource[settingsAuthGithubEnterpriseTerraformModel, *settingsAuthGithubEnterpriseTerraformModel]

// NewSettingsAuthGithubEnterpriseDataSource is a helper function to instantiate the SettingsAuthGithubEnterprise data source.
func NewSettingsAuthGithubEnterpriseDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise", Endpoint: "/api/v2/settings/github-enterprise/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubEnterpriseTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_enterprise_api_url": schema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
					"social_auth_github_enterprise_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
					},
					"social_auth_github_enterprise_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise developer application.",
						Computed:    true,
					},
					"social_auth_github_enterprise_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_enterprise_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_url": schema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithubEnterprise,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterprise",
		},
	}
}
