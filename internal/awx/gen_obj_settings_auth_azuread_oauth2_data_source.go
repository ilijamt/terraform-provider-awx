package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsAuthAzureAdoauth2DataSource = framework.GenericDataSource[settingsAuthAzureAdoauth2TerraformModel, *settingsAuthAzureAdoauth2TerraformModel]

// NewSettingsAuthAzureADOauth2DataSource is a helper function to instantiate the SettingsAuthAzureADOauth2 data source.
func NewSettingsAuthAzureADOauth2DataSource() datasource.DataSource {
	return &settingsAuthAzureAdoauth2DataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_azuread_oauth2", Endpoint: "/api/v2/settings/azuread-oauth2/"}},
		Cfg: framework.DataSourceCfg[settingsAuthAzureAdoauth2TerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"social_auth_azuread_oauth2_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail. ",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_azuread_oauth2_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your Azure AD application.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_azuread_oauth2_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_azuread_oauth2_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your Azure AD application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_azuread_oauth2_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
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
