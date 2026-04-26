package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsAuthGoogleOauth2DataSource = framework.GenericDataSource[settingsAuthGoogleOauth2TerraformModel, *settingsAuthGoogleOauth2TerraformModel]

// NewSettingsAuthGoogleOauth2DataSource is a helper function to instantiate the SettingsAuthGoogleOauth2 data source.
func NewSettingsAuthGoogleOauth2DataSource() datasource.DataSource {
	return &settingsAuthGoogleOauth2DataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_google_oauth2", Endpoint: "/api/v2/settings/google-oauth2/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGoogleOauth2TerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"social_auth_google_oauth2_auth_extra_arguments": schema.StringAttribute{
						Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_google_oauth2_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_google_oauth2_key": schema.StringAttribute{
						Description: "The OAuth2 key from your web application.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_google_oauth2_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_google_oauth2_secret": schema.StringAttribute{
						Description: "The OAuth2 secret from your web application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_google_oauth2_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
						Computed:    true,
					},
					"social_auth_google_oauth2_whitelisted_domains": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGoogleOauth2,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGoogleOauth2",
		},
	}
}
