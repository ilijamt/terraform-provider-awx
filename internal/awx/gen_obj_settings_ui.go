package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsUiTerraformModel struct {
	CUSTOM_LOGIN_INFO       types.String `tfsdk:"custom_login_info" json:"CUSTOM_LOGIN_INFO"`
	CUSTOM_LOGO             types.String `tfsdk:"custom_logo" json:"CUSTOM_LOGO"`
	MAX_UI_JOB_EVENTS       types.Int64  `tfsdk:"max_ui_job_events" json:"MAX_UI_JOB_EVENTS"`
	PENDO_TRACKING_STATE    types.String `tfsdk:"pendo_tracking_state" json:"PENDO_TRACKING_STATE"`
	UI_LIVE_UPDATES_ENABLED types.Bool   `tfsdk:"ui_live_updates_enabled" json:"UI_LIVE_UPDATES_ENABLED"`
}

func (o *settingsUiTerraformModel) Clone() settingsUiTerraformModel {
	return *o
}

func (o *settingsUiTerraformModel) BodyRequest() *settingsUiBodyRequestModel {
	var req settingsUiBodyRequestModel
	req.CUSTOM_LOGIN_INFO = o.CUSTOM_LOGIN_INFO.ValueString()
	req.CUSTOM_LOGO = o.CUSTOM_LOGO.ValueString()
	req.MAX_UI_JOB_EVENTS = o.MAX_UI_JOB_EVENTS.ValueInt64()
	req.UI_LIVE_UPDATES_ENABLED = o.UI_LIVE_UPDATES_ENABLED.ValueBool()
	return &req
}

func (o *settingsUiTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.CUSTOM_LOGIN_INFO, data["CUSTOM_LOGIN_INFO"], false))
	collect(helpers.AttrValueSetString(&o.CUSTOM_LOGO, data["CUSTOM_LOGO"], false))
	collect(helpers.AttrValueSetInt64(&o.MAX_UI_JOB_EVENTS, data["MAX_UI_JOB_EVENTS"]))
	collect(helpers.AttrValueSetString(&o.PENDO_TRACKING_STATE, data["PENDO_TRACKING_STATE"], false))
	collect(helpers.AttrValueSetBool(&o.UI_LIVE_UPDATES_ENABLED, data["UI_LIVE_UPDATES_ENABLED"]))
	return diags, nil
}

type settingsUiBodyRequestModel struct {
	CUSTOM_LOGIN_INFO       string `json:"CUSTOM_LOGIN_INFO,omitempty"`
	CUSTOM_LOGO             string `json:"CUSTOM_LOGO,omitempty"`
	MAX_UI_JOB_EVENTS       int64  `json:"MAX_UI_JOB_EVENTS,omitempty"`
	UI_LIVE_UPDATES_ENABLED bool   `json:"UI_LIVE_UPDATES_ENABLED"`
}

type settingsUiResource = framework.GenericResource[settingsUiTerraformModel, settingsUiBodyRequestModel, *settingsUiTerraformModel]

// NewSettingsUIResource is a helper function to simplify the provider implementation.
func NewSettingsUIResource() resource.Resource {
	return &settingsUiResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_ui", Endpoint: "/api/v2/settings/ui/"}},
		Cfg: framework.ResourceCfg[settingsUiTerraformModel, settingsUiBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"custom_login_info": schema.StringAttribute{
						Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"custom_logo": schema.StringAttribute{
						Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"max_ui_job_events": schema.Int64Attribute{
						Description: "Maximum number of job events for the UI to retrieve within a single request.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(4000),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"ui_live_updates_enabled": schema.BoolAttribute{
						Description: "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"pendo_tracking_state": schema.StringAttribute{
						Description: "Enable or Disable User Analytics Tracking.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"off",
								"anonymous",
								"detailed",
							),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsUI",
		},
	}
}

type settingsUiDataSource = framework.GenericDataSource[settingsUiTerraformModel, *settingsUiTerraformModel]

// NewSettingsUIDataSource is a helper function to instantiate the SettingsUI data source.
func NewSettingsUIDataSource() datasource.DataSource {
	return &settingsUiDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_ui", Endpoint: "/api/v2/settings/ui/"}},
		Cfg: framework.DataSourceCfg[settingsUiTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"custom_login_info": dschema.StringAttribute{
						Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
						Computed:    true,
					},
					"custom_logo": dschema.StringAttribute{
						Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
						Computed:    true,
					},
					"max_ui_job_events": dschema.Int64Attribute{
						Description: "Maximum number of job events for the UI to retrieve within a single request.",
						Computed:    true,
					},
					"pendo_tracking_state": dschema.StringAttribute{
						Description: "Enable or Disable User Analytics Tracking.",
						Computed:    true,
					},
					"ui_live_updates_enabled": dschema.BoolAttribute{
						Description: "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsUI",
		},
	}
}
