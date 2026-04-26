package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
