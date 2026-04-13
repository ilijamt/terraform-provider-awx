package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthGithubEnterpriseOrgDataSource = framework.GenericDataSource[settingsAuthGithubEnterpriseOrgTerraformModel, *settingsAuthGithubEnterpriseOrgTerraformModel]

// NewSettingsAuthGithubEnterpriseOrgDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseOrg data source.
func NewSettingsAuthGithubEnterpriseOrgDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseOrgDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise_org", Endpoint: "/api/v2/settings/github-enterprise-org/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubEnterpriseOrgTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"social_auth_github_enterprise_org_api_url": schema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_name": schema.StringAttribute{
						Description: "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_url": schema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithubEnterpriseOrg,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterpriseOrg",
		},
	}
}
