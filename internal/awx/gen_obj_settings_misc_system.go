package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// settingsMiscSystemTerraformModel maps the schema for SettingsMiscSystem when using Data Source
type settingsMiscSystemTerraformModel struct {
	// ACTIVITY_STREAM_ENABLED "Enable capturing activity for the activity stream."
	ACTIVITY_STREAM_ENABLED types.Bool `tfsdk:"activity_stream_enabled" json:"ACTIVITY_STREAM_ENABLED"`
	// ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC "Enable capturing activity for the activity stream when running inventory sync."
	ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC types.Bool `tfsdk:"activity_stream_enabled_for_inventory_sync" json:"ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"`
	// AUTOMATION_ANALYTICS_GATHER_INTERVAL "Interval (in seconds) between data gathering."
	AUTOMATION_ANALYTICS_GATHER_INTERVAL types.Int64 `tfsdk:"automation_analytics_gather_interval" json:"AUTOMATION_ANALYTICS_GATHER_INTERVAL"`
	// AUTOMATION_ANALYTICS_LAST_ENTRIES ""
	AUTOMATION_ANALYTICS_LAST_ENTRIES types.String `tfsdk:"automation_analytics_last_entries" json:"AUTOMATION_ANALYTICS_LAST_ENTRIES"`
	// AUTOMATION_ANALYTICS_LAST_GATHER ""
	AUTOMATION_ANALYTICS_LAST_GATHER types.String `tfsdk:"automation_analytics_last_gather" json:"AUTOMATION_ANALYTICS_LAST_GATHER"`
	// AUTOMATION_ANALYTICS_URL "This setting is used to to configure the upload URL for data collection for Automation Analytics."
	AUTOMATION_ANALYTICS_URL types.String `tfsdk:"automation_analytics_url" json:"AUTOMATION_ANALYTICS_URL"`
	// CUSTOM_VENV_PATHS "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line."
	CUSTOM_VENV_PATHS types.List `tfsdk:"custom_venv_paths" json:"CUSTOM_VENV_PATHS"`
	// DEFAULT_CONTROL_PLANE_QUEUE_NAME ""
	DEFAULT_CONTROL_PLANE_QUEUE_NAME types.String `tfsdk:"default_control_plane_queue_name" json:"DEFAULT_CONTROL_PLANE_QUEUE_NAME"`
	// DEFAULT_EXECUTION_ENVIRONMENT "The Execution Environment to be used when one has not been configured for a job template."
	DEFAULT_EXECUTION_ENVIRONMENT types.Int64 `tfsdk:"default_execution_environment" json:"DEFAULT_EXECUTION_ENVIRONMENT"`
	// DEFAULT_EXECUTION_QUEUE_NAME ""
	DEFAULT_EXECUTION_QUEUE_NAME types.String `tfsdk:"default_execution_queue_name" json:"DEFAULT_EXECUTION_QUEUE_NAME"`
	// INSIGHTS_TRACKING_STATE "Enables the service to gather data on automation and send it to Automation Analytics."
	INSIGHTS_TRACKING_STATE types.Bool `tfsdk:"insights_tracking_state" json:"INSIGHTS_TRACKING_STATE"`
	// INSTALL_UUID ""
	INSTALL_UUID types.String `tfsdk:"install_uuid" json:"INSTALL_UUID"`
	// IS_K8S "Indicates whether the instance is part of a kubernetes-based deployment."
	IS_K8S types.Bool `tfsdk:"is_k8s" json:"IS_K8S"`
	// LICENSE "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license."
	LICENSE types.String `tfsdk:"license" json:"LICENSE"`
	// MANAGE_ORGANIZATION_AUTH "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration."
	MANAGE_ORGANIZATION_AUTH types.Bool `tfsdk:"manage_organization_auth" json:"MANAGE_ORGANIZATION_AUTH"`
	// ORG_ADMINS_CAN_SEE_ALL_USERS "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization."
	ORG_ADMINS_CAN_SEE_ALL_USERS types.Bool `tfsdk:"org_admins_can_see_all_users" json:"ORG_ADMINS_CAN_SEE_ALL_USERS"`
	// PROXY_IP_ALLOWED_LIST "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')"
	PROXY_IP_ALLOWED_LIST types.List `tfsdk:"proxy_ip_allowed_list" json:"PROXY_IP_ALLOWED_LIST"`
	// REDHAT_PASSWORD "This password is used to send data to Automation Analytics"
	REDHAT_PASSWORD types.String `tfsdk:"redhat_password" json:"REDHAT_PASSWORD"`
	// REDHAT_USERNAME "This username is used to send data to Automation Analytics"
	REDHAT_USERNAME types.String `tfsdk:"redhat_username" json:"REDHAT_USERNAME"`
	// REMOTE_HOST_HEADERS "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details."
	REMOTE_HOST_HEADERS types.List `tfsdk:"remote_host_headers" json:"REMOTE_HOST_HEADERS"`
	// SUBSCRIPTIONS_PASSWORD "This password is used to retrieve subscription and content information"
	SUBSCRIPTIONS_PASSWORD types.String `tfsdk:"subscriptions_password" json:"SUBSCRIPTIONS_PASSWORD"`
	// SUBSCRIPTIONS_USERNAME "This username is used to retrieve subscription and content information"
	SUBSCRIPTIONS_USERNAME types.String `tfsdk:"subscriptions_username" json:"SUBSCRIPTIONS_USERNAME"`
	// TOWER_URL_BASE "This value has been set manually in a settings file.\n\nThis setting is used by services like notifications to render a valid url to the service."
	TOWER_URL_BASE types.String `tfsdk:"tower_url_base" json:"TOWER_URL_BASE"`
}

// Clone the object
func (o settingsMiscSystemTerraformModel) Clone() settingsMiscSystemTerraformModel {
	return settingsMiscSystemTerraformModel{
		ACTIVITY_STREAM_ENABLED:                    o.ACTIVITY_STREAM_ENABLED,
		ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC: o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC,
		AUTOMATION_ANALYTICS_GATHER_INTERVAL:       o.AUTOMATION_ANALYTICS_GATHER_INTERVAL,
		AUTOMATION_ANALYTICS_LAST_ENTRIES:          o.AUTOMATION_ANALYTICS_LAST_ENTRIES,
		AUTOMATION_ANALYTICS_LAST_GATHER:           o.AUTOMATION_ANALYTICS_LAST_GATHER,
		AUTOMATION_ANALYTICS_URL:                   o.AUTOMATION_ANALYTICS_URL,
		CUSTOM_VENV_PATHS:                          o.CUSTOM_VENV_PATHS,
		DEFAULT_CONTROL_PLANE_QUEUE_NAME:           o.DEFAULT_CONTROL_PLANE_QUEUE_NAME,
		DEFAULT_EXECUTION_ENVIRONMENT:              o.DEFAULT_EXECUTION_ENVIRONMENT,
		DEFAULT_EXECUTION_QUEUE_NAME:               o.DEFAULT_EXECUTION_QUEUE_NAME,
		INSIGHTS_TRACKING_STATE:                    o.INSIGHTS_TRACKING_STATE,
		INSTALL_UUID:                               o.INSTALL_UUID,
		IS_K8S:                                     o.IS_K8S,
		LICENSE:                                    o.LICENSE,
		MANAGE_ORGANIZATION_AUTH:                   o.MANAGE_ORGANIZATION_AUTH,
		ORG_ADMINS_CAN_SEE_ALL_USERS:               o.ORG_ADMINS_CAN_SEE_ALL_USERS,
		PROXY_IP_ALLOWED_LIST:                      o.PROXY_IP_ALLOWED_LIST,
		REDHAT_PASSWORD:                            o.REDHAT_PASSWORD,
		REDHAT_USERNAME:                            o.REDHAT_USERNAME,
		REMOTE_HOST_HEADERS:                        o.REMOTE_HOST_HEADERS,
		SUBSCRIPTIONS_PASSWORD:                     o.SUBSCRIPTIONS_PASSWORD,
		SUBSCRIPTIONS_USERNAME:                     o.SUBSCRIPTIONS_USERNAME,
		TOWER_URL_BASE:                             o.TOWER_URL_BASE,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsMiscSystem
func (o settingsMiscSystemTerraformModel) BodyRequest() (req settingsMiscSystemBodyRequestModel) {
	req.ACTIVITY_STREAM_ENABLED = o.ACTIVITY_STREAM_ENABLED.ValueBool()
	req.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC = o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC.ValueBool()
	req.AUTOMATION_ANALYTICS_GATHER_INTERVAL = o.AUTOMATION_ANALYTICS_GATHER_INTERVAL.ValueInt64()
	req.AUTOMATION_ANALYTICS_LAST_ENTRIES = o.AUTOMATION_ANALYTICS_LAST_ENTRIES.ValueString()
	req.AUTOMATION_ANALYTICS_URL = o.AUTOMATION_ANALYTICS_URL.ValueString()
	req.DEFAULT_EXECUTION_ENVIRONMENT = o.DEFAULT_EXECUTION_ENVIRONMENT.ValueInt64()
	req.INSIGHTS_TRACKING_STATE = o.INSIGHTS_TRACKING_STATE.ValueBool()
	req.MANAGE_ORGANIZATION_AUTH = o.MANAGE_ORGANIZATION_AUTH.ValueBool()
	req.ORG_ADMINS_CAN_SEE_ALL_USERS = o.ORG_ADMINS_CAN_SEE_ALL_USERS.ValueBool()
	req.PROXY_IP_ALLOWED_LIST = []string{}
	for _, val := range o.PROXY_IP_ALLOWED_LIST.Elements() {
		if _, ok := val.(types.String); ok {
			req.PROXY_IP_ALLOWED_LIST = append(req.PROXY_IP_ALLOWED_LIST, val.(types.String).ValueString())
		} else {
			req.PROXY_IP_ALLOWED_LIST = append(req.PROXY_IP_ALLOWED_LIST, val.String())
		}
	}
	req.REDHAT_PASSWORD = o.REDHAT_PASSWORD.ValueString()
	req.REDHAT_USERNAME = o.REDHAT_USERNAME.ValueString()
	req.REMOTE_HOST_HEADERS = []string{}
	for _, val := range o.REMOTE_HOST_HEADERS.Elements() {
		if _, ok := val.(types.String); ok {
			req.REMOTE_HOST_HEADERS = append(req.REMOTE_HOST_HEADERS, val.(types.String).ValueString())
		} else {
			req.REMOTE_HOST_HEADERS = append(req.REMOTE_HOST_HEADERS, val.String())
		}
	}
	req.SUBSCRIPTIONS_PASSWORD = o.SUBSCRIPTIONS_PASSWORD.ValueString()
	req.SUBSCRIPTIONS_USERNAME = o.SUBSCRIPTIONS_USERNAME.ValueString()
	return
}

func (o *settingsMiscSystemTerraformModel) setActivityStreamEnabled(data any) (d diag.Diagnostics, err error) {
	// Decode "ACTIVITY_STREAM_ENABLED"
	if val, ok := data.(bool); ok {
		o.ACTIVITY_STREAM_ENABLED = types.BoolValue(val)
	} else {
		o.ACTIVITY_STREAM_ENABLED = types.BoolNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setActivityStreamEnabledForInventorySync(data any) (d diag.Diagnostics, err error) {
	// Decode "ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"
	if val, ok := data.(bool); ok {
		o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC = types.BoolValue(val)
	} else {
		o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC = types.BoolNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsGatherInterval(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTOMATION_ANALYTICS_GATHER_INTERVAL"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.AUTOMATION_ANALYTICS_GATHER_INTERVAL = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.AUTOMATION_ANALYTICS_GATHER_INTERVAL = types.Int64Value(val)
	} else {
		o.AUTOMATION_ANALYTICS_GATHER_INTERVAL = types.Int64Null()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsLastEntries(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTOMATION_ANALYTICS_LAST_ENTRIES"
	if val, ok := data.(string); ok {
		o.AUTOMATION_ANALYTICS_LAST_ENTRIES = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTOMATION_ANALYTICS_LAST_ENTRIES = types.StringValue(val.String())
	} else {
		o.AUTOMATION_ANALYTICS_LAST_ENTRIES = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsLastGather(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTOMATION_ANALYTICS_LAST_GATHER"
	if val, ok := data.(string); ok {
		o.AUTOMATION_ANALYTICS_LAST_GATHER = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTOMATION_ANALYTICS_LAST_GATHER = types.StringValue(val.String())
	} else {
		o.AUTOMATION_ANALYTICS_LAST_GATHER = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsUrl(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTOMATION_ANALYTICS_URL"
	if val, ok := data.(string); ok {
		o.AUTOMATION_ANALYTICS_URL = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTOMATION_ANALYTICS_URL = types.StringValue(val.String())
	} else {
		o.AUTOMATION_ANALYTICS_URL = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setCustomVenvPaths(data any) (d diag.Diagnostics, err error) {
	// Decode "CUSTOM_VENV_PATHS"
	if val, ok := data.(types.List); ok {
		o.CUSTOM_VENV_PATHS = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.CUSTOM_VENV_PATHS = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.CUSTOM_VENV_PATHS = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setDefaultControlPlaneQueueName(data any) (d diag.Diagnostics, err error) {
	// Decode "DEFAULT_CONTROL_PLANE_QUEUE_NAME"
	if val, ok := data.(string); ok {
		o.DEFAULT_CONTROL_PLANE_QUEUE_NAME = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.DEFAULT_CONTROL_PLANE_QUEUE_NAME = types.StringValue(val.String())
	} else {
		o.DEFAULT_CONTROL_PLANE_QUEUE_NAME = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setDefaultExecutionEnvironment(data any) (d diag.Diagnostics, err error) {
	// Decode "DEFAULT_EXECUTION_ENVIRONMENT"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.DEFAULT_EXECUTION_ENVIRONMENT = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.DEFAULT_EXECUTION_ENVIRONMENT = types.Int64Value(val)
	} else {
		o.DEFAULT_EXECUTION_ENVIRONMENT = types.Int64Null()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setDefaultExecutionQueueName(data any) (d diag.Diagnostics, err error) {
	// Decode "DEFAULT_EXECUTION_QUEUE_NAME"
	if val, ok := data.(string); ok {
		o.DEFAULT_EXECUTION_QUEUE_NAME = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.DEFAULT_EXECUTION_QUEUE_NAME = types.StringValue(val.String())
	} else {
		o.DEFAULT_EXECUTION_QUEUE_NAME = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setInsightsTrackingState(data any) (d diag.Diagnostics, err error) {
	// Decode "INSIGHTS_TRACKING_STATE"
	if val, ok := data.(bool); ok {
		o.INSIGHTS_TRACKING_STATE = types.BoolValue(val)
	} else {
		o.INSIGHTS_TRACKING_STATE = types.BoolNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setInstallUuid(data any) (d diag.Diagnostics, err error) {
	// Decode "INSTALL_UUID"
	if val, ok := data.(string); ok {
		o.INSTALL_UUID = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.INSTALL_UUID = types.StringValue(val.String())
	} else {
		o.INSTALL_UUID = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setIsK8S(data any) (d diag.Diagnostics, err error) {
	// Decode "IS_K8S"
	if val, ok := data.(bool); ok {
		o.IS_K8S = types.BoolValue(val)
	} else {
		o.IS_K8S = types.BoolNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setLicense(data any) (d diag.Diagnostics, err error) {
	// Decode "LICENSE"
	if val, ok := data.(string); ok {
		o.LICENSE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.LICENSE = types.StringValue(val.String())
	} else {
		o.LICENSE = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setManageOrganizationAuth(data any) (d diag.Diagnostics, err error) {
	// Decode "MANAGE_ORGANIZATION_AUTH"
	if val, ok := data.(bool); ok {
		o.MANAGE_ORGANIZATION_AUTH = types.BoolValue(val)
	} else {
		o.MANAGE_ORGANIZATION_AUTH = types.BoolNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setOrgAdminsCanSeeAllUsers(data any) (d diag.Diagnostics, err error) {
	// Decode "ORG_ADMINS_CAN_SEE_ALL_USERS"
	if val, ok := data.(bool); ok {
		o.ORG_ADMINS_CAN_SEE_ALL_USERS = types.BoolValue(val)
	} else {
		o.ORG_ADMINS_CAN_SEE_ALL_USERS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setProxyIpAllowedList(data any) (d diag.Diagnostics, err error) {
	// Decode "PROXY_IP_ALLOWED_LIST"
	if val, ok := data.(types.List); ok {
		o.PROXY_IP_ALLOWED_LIST = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.PROXY_IP_ALLOWED_LIST = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.PROXY_IP_ALLOWED_LIST = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setRedhatPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "REDHAT_PASSWORD"
	if val, ok := data.(string); ok {
		o.REDHAT_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.REDHAT_PASSWORD = types.StringValue(val.String())
	} else {
		o.REDHAT_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setRedhatUsername(data any) (d diag.Diagnostics, err error) {
	// Decode "REDHAT_USERNAME"
	if val, ok := data.(string); ok {
		o.REDHAT_USERNAME = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.REDHAT_USERNAME = types.StringValue(val.String())
	} else {
		o.REDHAT_USERNAME = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setRemoteHostHeaders(data any) (d diag.Diagnostics, err error) {
	// Decode "REMOTE_HOST_HEADERS"
	if val, ok := data.(types.List); ok {
		o.REMOTE_HOST_HEADERS = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.REMOTE_HOST_HEADERS = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.REMOTE_HOST_HEADERS = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setSubscriptionsPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "SUBSCRIPTIONS_PASSWORD"
	if val, ok := data.(string); ok {
		o.SUBSCRIPTIONS_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.SUBSCRIPTIONS_PASSWORD = types.StringValue(val.String())
	} else {
		o.SUBSCRIPTIONS_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setSubscriptionsUsername(data any) (d diag.Diagnostics, err error) {
	// Decode "SUBSCRIPTIONS_USERNAME"
	if val, ok := data.(string); ok {
		o.SUBSCRIPTIONS_USERNAME = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.SUBSCRIPTIONS_USERNAME = types.StringValue(val.String())
	} else {
		o.SUBSCRIPTIONS_USERNAME = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) setTowerUrlBase(data any) (d diag.Diagnostics, err error) {
	// Decode "TOWER_URL_BASE"
	if val, ok := data.(string); ok {
		o.TOWER_URL_BASE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.TOWER_URL_BASE = types.StringValue(val.String())
	} else {
		o.TOWER_URL_BASE = types.StringNull()
	}
	return d, nil
}

func (o *settingsMiscSystemTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setActivityStreamEnabled(data["ACTIVITY_STREAM_ENABLED"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setActivityStreamEnabledForInventorySync(data["ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAutomationAnalyticsGatherInterval(data["AUTOMATION_ANALYTICS_GATHER_INTERVAL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAutomationAnalyticsLastEntries(data["AUTOMATION_ANALYTICS_LAST_ENTRIES"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAutomationAnalyticsLastGather(data["AUTOMATION_ANALYTICS_LAST_GATHER"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAutomationAnalyticsUrl(data["AUTOMATION_ANALYTICS_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCustomVenvPaths(data["CUSTOM_VENV_PATHS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultControlPlaneQueueName(data["DEFAULT_CONTROL_PLANE_QUEUE_NAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultExecutionEnvironment(data["DEFAULT_EXECUTION_ENVIRONMENT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultExecutionQueueName(data["DEFAULT_EXECUTION_QUEUE_NAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInsightsTrackingState(data["INSIGHTS_TRACKING_STATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInstallUuid(data["INSTALL_UUID"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsK8S(data["IS_K8S"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLicense(data["LICENSE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setManageOrganizationAuth(data["MANAGE_ORGANIZATION_AUTH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrgAdminsCanSeeAllUsers(data["ORG_ADMINS_CAN_SEE_ALL_USERS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setProxyIpAllowedList(data["PROXY_IP_ALLOWED_LIST"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRedhatPassword(data["REDHAT_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRedhatUsername(data["REDHAT_USERNAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRemoteHostHeaders(data["REMOTE_HOST_HEADERS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSubscriptionsPassword(data["SUBSCRIPTIONS_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSubscriptionsUsername(data["SUBSCRIPTIONS_USERNAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTowerUrlBase(data["TOWER_URL_BASE"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsMiscSystemBodyRequestModel maps the schema for SettingsMiscSystem for creating and updating the data
type settingsMiscSystemBodyRequestModel struct {
	// ACTIVITY_STREAM_ENABLED "Enable capturing activity for the activity stream."
	ACTIVITY_STREAM_ENABLED bool `json:"ACTIVITY_STREAM_ENABLED"`
	// ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC "Enable capturing activity for the activity stream when running inventory sync."
	ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC bool `json:"ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"`
	// AUTOMATION_ANALYTICS_GATHER_INTERVAL "Interval (in seconds) between data gathering."
	AUTOMATION_ANALYTICS_GATHER_INTERVAL int64 `json:"AUTOMATION_ANALYTICS_GATHER_INTERVAL,omitempty"`
	// AUTOMATION_ANALYTICS_LAST_ENTRIES ""
	AUTOMATION_ANALYTICS_LAST_ENTRIES string `json:"AUTOMATION_ANALYTICS_LAST_ENTRIES,omitempty"`
	// AUTOMATION_ANALYTICS_URL "This setting is used to to configure the upload URL for data collection for Automation Analytics."
	AUTOMATION_ANALYTICS_URL string `json:"AUTOMATION_ANALYTICS_URL,omitempty"`
	// DEFAULT_EXECUTION_ENVIRONMENT "The Execution Environment to be used when one has not been configured for a job template."
	DEFAULT_EXECUTION_ENVIRONMENT int64 `json:"DEFAULT_EXECUTION_ENVIRONMENT,omitempty"`
	// INSIGHTS_TRACKING_STATE "Enables the service to gather data on automation and send it to Automation Analytics."
	INSIGHTS_TRACKING_STATE bool `json:"INSIGHTS_TRACKING_STATE"`
	// MANAGE_ORGANIZATION_AUTH "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration."
	MANAGE_ORGANIZATION_AUTH bool `json:"MANAGE_ORGANIZATION_AUTH"`
	// ORG_ADMINS_CAN_SEE_ALL_USERS "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization."
	ORG_ADMINS_CAN_SEE_ALL_USERS bool `json:"ORG_ADMINS_CAN_SEE_ALL_USERS"`
	// PROXY_IP_ALLOWED_LIST "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')"
	PROXY_IP_ALLOWED_LIST []string `json:"PROXY_IP_ALLOWED_LIST"`
	// REDHAT_PASSWORD "This password is used to send data to Automation Analytics"
	REDHAT_PASSWORD string `json:"REDHAT_PASSWORD,omitempty"`
	// REDHAT_USERNAME "This username is used to send data to Automation Analytics"
	REDHAT_USERNAME string `json:"REDHAT_USERNAME,omitempty"`
	// REMOTE_HOST_HEADERS "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details."
	REMOTE_HOST_HEADERS []string `json:"REMOTE_HOST_HEADERS"`
	// SUBSCRIPTIONS_PASSWORD "This password is used to retrieve subscription and content information"
	SUBSCRIPTIONS_PASSWORD string `json:"SUBSCRIPTIONS_PASSWORD,omitempty"`
	// SUBSCRIPTIONS_USERNAME "This username is used to retrieve subscription and content information"
	SUBSCRIPTIONS_USERNAME string `json:"SUBSCRIPTIONS_USERNAME,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsMiscSystemDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsMiscSystemDataSource{}
)

// NewSettingsMiscSystemDataSource is a helper function to instantiate the SettingsMiscSystem data source.
func NewSettingsMiscSystemDataSource() datasource.DataSource {
	return &settingsMiscSystemDataSource{}
}

// settingsMiscSystemDataSource is the data source implementation.
type settingsMiscSystemDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsMiscSystemDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/system/"
}

// Metadata returns the data source type name.
func (o *settingsMiscSystemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_misc_system"
}

// GetSchema defines the schema for the data source.
func (o *settingsMiscSystemDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsMiscSystem",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"activity_stream_enabled": {
					Description: "Enable capturing activity for the activity stream.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"activity_stream_enabled_for_inventory_sync": {
					Description: "Enable capturing activity for the activity stream when running inventory sync.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"automation_analytics_gather_interval": {
					Description: "Interval (in seconds) between data gathering.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"automation_analytics_last_entries": {
					Description: "Last gathered entries from the data collection service of Automation Analytics",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"automation_analytics_last_gather": {
					Description: "Last gather date for Automation Analytics.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"automation_analytics_url": {
					Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"custom_venv_paths": {
					Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_control_plane_queue_name": {
					Description: "The instance group where control plane tasks run",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_execution_environment": {
					Description: "The Execution Environment to be used when one has not been configured for a job template.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_execution_queue_name": {
					Description: "The instance group where user jobs run (currently only on non-VM installs)",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"insights_tracking_state": {
					Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"install_uuid": {
					Description: "Unique identifier for an installation",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"is_k8s": {
					Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"license": {
					Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"manage_organization_auth": {
					Description: "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"org_admins_can_see_all_users": {
					Description: "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"proxy_ip_allowed_list": {
					Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"redhat_password": {
					Description: "This password is used to send data to Automation Analytics",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"redhat_username": {
					Description: "This username is used to send data to Automation Analytics",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"remote_host_headers": {
					Description: "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"subscriptions_password": {
					Description: "This password is used to retrieve subscription and content information",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"subscriptions_username": {
					Description: "This username is used to retrieve subscription and content information",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"tower_url_base": {
					Description: "This value has been set manually in a settings file.\n\nThis setting is used by services like notifications to render a valid url to the service.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsMiscSystemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsMiscSystemTerraformModel
	var err error
	var endpoint string
	endpoint = o.endpoint

	// Creates a new request for SettingsMiscSystem
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscSystem on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsMiscSystem
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsMiscSystem on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsMiscSystemResource{}
	_ resource.ResourceWithConfigure = &settingsMiscSystemResource{}
)

// NewSettingsMiscSystemResource is a helper function to simplify the provider implementation.
func NewSettingsMiscSystemResource() resource.Resource {
	return &settingsMiscSystemResource{}
}

type settingsMiscSystemResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsMiscSystemResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/system/"
}

func (o settingsMiscSystemResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_misc_system"
}

func (o settingsMiscSystemResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsMiscSystem",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"activity_stream_enabled": {
					Description:   "Enable capturing activity for the activity stream.",
					Type:          types.BoolType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"activity_stream_enabled_for_inventory_sync": {
					Description:   "Enable capturing activity for the activity stream when running inventory sync.",
					Type:          types.BoolType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"automation_analytics_gather_interval": {
					Description: "Interval (in seconds) between data gathering.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(14400)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"automation_analytics_last_entries": {
					Description: "Last gathered entries from the data collection service of Automation Analytics",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"automation_analytics_url": {
					Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`https://example.com`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"default_execution_environment": {
					Description: "The Execution Environment to be used when one has not been configured for a job template.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"insights_tracking_state": {
					Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"manage_organization_auth": {
					Description:   "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
					Type:          types.BoolType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"org_admins_can_see_all_users": {
					Description:   "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
					Type:          types.BoolType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"proxy_ip_allowed_list": {
					Description:   "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
					Type:          types.ListType{ElemType: types.StringType},
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"redhat_password": {
					Description: "This password is used to send data to Automation Analytics",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"redhat_username": {
					Description: "This username is used to send data to Automation Analytics",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"remote_host_headers": {
					Description:   "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
					Type:          types.ListType{ElemType: types.StringType},
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"subscriptions_password": {
					Description: "This password is used to retrieve subscription and content information",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"subscriptions_username": {
					Description: "This username is used to retrieve subscription and content information",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"automation_analytics_last_gather": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"custom_venv_paths": {
					Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
					Computed:    true,
					Type:        types.ListType{ElemType: types.StringType},
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"default_control_plane_queue_name": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"default_execution_queue_name": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"install_uuid": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"is_k8s": {
					Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
					Computed:    true,
					Type:        types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"license": {
					Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"tower_url_base": {
					Description: "This value has been set manually in a settings file.\n\nThis setting is used by services like notifications to render a valid url to the service.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *settingsMiscSystemResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsMiscSystemTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsMiscSystem
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscSystem on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsMiscSystem resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsMiscSystem on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsMiscSystemResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsMiscSystemTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsMiscSystem
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscSystem on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsMiscSystem from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsMiscSystem on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsMiscSystemResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsMiscSystemTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsMiscSystem
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscSystem on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsMiscSystem resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsMiscSystem on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsMiscSystemResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	return
}
