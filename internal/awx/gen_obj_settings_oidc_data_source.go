package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsOpenIdconnectDataSource = framework.GenericDataSource[settingsOpenIdconnectTerraformModel, *settingsOpenIdconnectTerraformModel]

// NewSettingsOpenIDConnectDataSource is a helper function to instantiate the SettingsOpenIDConnect data source.
func NewSettingsOpenIDConnectDataSource() datasource.DataSource {
	return &settingsOpenIdconnectDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_oidc", Endpoint: "/api/v2/settings/oidc/"}},
		Cfg: framework.DataSourceCfg[settingsOpenIdconnectTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_oidc_key": schema.StringAttribute{
						Description: "The OIDC key (Client ID) from your IDP.",
						Computed:    true,
					},
					"social_auth_oidc_oidc_endpoint": schema.StringAttribute{
						Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
						Computed:    true,
					},
					"social_auth_oidc_secret": schema.StringAttribute{
						Description: "The OIDC secret (Client Secret) from your IDP.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_oidc_verify_ssl": schema.BoolAttribute{
						Description: "Verify the OIDC provider ssl certificate.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsOidc,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsOpenIDConnect",
		},
	}
}
