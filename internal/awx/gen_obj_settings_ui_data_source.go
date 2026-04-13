package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type settingsUiDataSource = framework.GenericDataSource[settingsUiTerraformModel, *settingsUiTerraformModel]

// NewSettingsUIDataSource is a helper function to instantiate the SettingsUI data source.
func NewSettingsUIDataSource() datasource.DataSource {
	return &settingsUiDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_ui", Endpoint: "/api/v2/settings/ui/"}},
		Cfg: framework.DataSourceCfg[settingsUiTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"custom_login_info": schema.StringAttribute{
						Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
						Sensitive:   false,
						Computed:    true,
					},
					"custom_logo": schema.StringAttribute{
						Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
						Sensitive:   false,
						Computed:    true,
					},
					"max_ui_job_events": schema.Int64Attribute{
						Description: "Maximum number of job events for the UI to retrieve within a single request.",
						Sensitive:   false,
						Computed:    true,
					},
					"pendo_tracking_state": schema.StringAttribute{
						Description: "Enable or Disable User Analytics Tracking.",
						Sensitive:   false,
						Computed:    true,
					},
					"ui_live_updates_enabled": schema.BoolAttribute{
						Description: "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsUI",
		},
	}
}
