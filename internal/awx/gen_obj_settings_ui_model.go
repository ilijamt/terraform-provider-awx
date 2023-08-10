package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsUiTerraformModel maps the schema for SettingsUI when using Data Source
type settingsUiTerraformModel struct {
	// CUSTOM_LOGIN_INFO "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported."
	CUSTOM_LOGIN_INFO types.String `tfsdk:"custom_login_info" json:"CUSTOM_LOGIN_INFO"`
	// CUSTOM_LOGO "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported."
	CUSTOM_LOGO types.String `tfsdk:"custom_logo" json:"CUSTOM_LOGO"`
	// MAX_UI_JOB_EVENTS "Maximum number of job events for the UI to retrieve within a single request."
	MAX_UI_JOB_EVENTS types.Int64 `tfsdk:"max_ui_job_events" json:"MAX_UI_JOB_EVENTS"`
	// PENDO_TRACKING_STATE "Enable or Disable User Analytics Tracking."
	PENDO_TRACKING_STATE types.String `tfsdk:"pendo_tracking_state" json:"PENDO_TRACKING_STATE"`
	// UI_LIVE_UPDATES_ENABLED "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details."
	UI_LIVE_UPDATES_ENABLED types.Bool `tfsdk:"ui_live_updates_enabled" json:"UI_LIVE_UPDATES_ENABLED"`
}

// Clone the object
func (o *settingsUiTerraformModel) Clone() settingsUiTerraformModel {
	return settingsUiTerraformModel{
		CUSTOM_LOGIN_INFO:       o.CUSTOM_LOGIN_INFO,
		CUSTOM_LOGO:             o.CUSTOM_LOGO,
		MAX_UI_JOB_EVENTS:       o.MAX_UI_JOB_EVENTS,
		PENDO_TRACKING_STATE:    o.PENDO_TRACKING_STATE,
		UI_LIVE_UPDATES_ENABLED: o.UI_LIVE_UPDATES_ENABLED,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsUI
func (o *settingsUiTerraformModel) BodyRequest() (req settingsUiBodyRequestModel) {
	req.CUSTOM_LOGIN_INFO = o.CUSTOM_LOGIN_INFO.ValueString()
	req.CUSTOM_LOGO = o.CUSTOM_LOGO.ValueString()
	req.MAX_UI_JOB_EVENTS = o.MAX_UI_JOB_EVENTS.ValueInt64()
	req.UI_LIVE_UPDATES_ENABLED = o.UI_LIVE_UPDATES_ENABLED.ValueBool()
	return
}

func (o *settingsUiTerraformModel) setCustomLoginInfo(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.CUSTOM_LOGIN_INFO, data, false)
}

func (o *settingsUiTerraformModel) setCustomLogo(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.CUSTOM_LOGO, data, false)
}

func (o *settingsUiTerraformModel) setMaxUiJobEvents(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.MAX_UI_JOB_EVENTS, data)
}

func (o *settingsUiTerraformModel) setPendoTrackingState(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.PENDO_TRACKING_STATE, data, false)
}

func (o *settingsUiTerraformModel) setUiLiveUpdatesEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.UI_LIVE_UPDATES_ENABLED, data)
}

func (o *settingsUiTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setCustomLoginInfo(data["CUSTOM_LOGIN_INFO"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCustomLogo(data["CUSTOM_LOGO"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMaxUiJobEvents(data["MAX_UI_JOB_EVENTS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPendoTrackingState(data["PENDO_TRACKING_STATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUiLiveUpdatesEnabled(data["UI_LIVE_UPDATES_ENABLED"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsUiBodyRequestModel maps the schema for SettingsUI for creating and updating the data
type settingsUiBodyRequestModel struct {
	// CUSTOM_LOGIN_INFO "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported."
	CUSTOM_LOGIN_INFO string `json:"CUSTOM_LOGIN_INFO,omitempty"`
	// CUSTOM_LOGO "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported."
	CUSTOM_LOGO string `json:"CUSTOM_LOGO,omitempty"`
	// MAX_UI_JOB_EVENTS "Maximum number of job events for the UI to retrieve within a single request."
	MAX_UI_JOB_EVENTS int64 `json:"MAX_UI_JOB_EVENTS,omitempty"`
	// UI_LIVE_UPDATES_ENABLED "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details."
	UI_LIVE_UPDATES_ENABLED bool `json:"UI_LIVE_UPDATES_ENABLED"`
}
