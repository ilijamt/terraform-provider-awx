package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

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
