package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
	// CLEANUP_HOST_METRICS_LAST_TS ""
	CLEANUP_HOST_METRICS_LAST_TS types.String `tfsdk:"cleanup_host_metrics_last_ts" json:"CLEANUP_HOST_METRICS_LAST_TS"`
	// CSRF_TRUSTED_ORIGINS "If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values. "
	CSRF_TRUSTED_ORIGINS types.List `tfsdk:"csrf_trusted_origins" json:"CSRF_TRUSTED_ORIGINS"`
	// CUSTOM_VENV_PATHS "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line."
	CUSTOM_VENV_PATHS types.List `tfsdk:"custom_venv_paths" json:"CUSTOM_VENV_PATHS"`
	// DEFAULT_CONTROL_PLANE_QUEUE_NAME ""
	DEFAULT_CONTROL_PLANE_QUEUE_NAME types.String `tfsdk:"default_control_plane_queue_name" json:"DEFAULT_CONTROL_PLANE_QUEUE_NAME"`
	// DEFAULT_EXECUTION_ENVIRONMENT "The Execution Environment to be used when one has not been configured for a job template."
	DEFAULT_EXECUTION_ENVIRONMENT types.Int64 `tfsdk:"default_execution_environment" json:"DEFAULT_EXECUTION_ENVIRONMENT"`
	// DEFAULT_EXECUTION_QUEUE_NAME ""
	DEFAULT_EXECUTION_QUEUE_NAME types.String `tfsdk:"default_execution_queue_name" json:"DEFAULT_EXECUTION_QUEUE_NAME"`
	// HOST_METRIC_SUMMARY_TASK_LAST_TS ""
	HOST_METRIC_SUMMARY_TASK_LAST_TS types.String `tfsdk:"host_metric_summary_task_last_ts" json:"HOST_METRIC_SUMMARY_TASK_LAST_TS"`
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
	// SUBSCRIPTION_USAGE_MODEL ""
	SUBSCRIPTION_USAGE_MODEL types.String `tfsdk:"subscription_usage_model" json:"SUBSCRIPTION_USAGE_MODEL"`
	// TOWER_URL_BASE "This setting is used by services like notifications to render a valid url to the service."
	TOWER_URL_BASE types.String `tfsdk:"tower_url_base" json:"TOWER_URL_BASE"`
	// UI_NEXT "Enable preview of new user interface."
	UI_NEXT types.Bool `tfsdk:"ui_next" json:"UI_NEXT"`
}

// Clone the object
func (o *settingsMiscSystemTerraformModel) Clone() settingsMiscSystemTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsMiscSystem
func (o *settingsMiscSystemTerraformModel) BodyRequest() *settingsMiscSystemBodyRequestModel {
	var req settingsMiscSystemBodyRequestModel
	req.ACTIVITY_STREAM_ENABLED = o.ACTIVITY_STREAM_ENABLED.ValueBool()
	req.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC = o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC.ValueBool()
	req.AUTOMATION_ANALYTICS_GATHER_INTERVAL = o.AUTOMATION_ANALYTICS_GATHER_INTERVAL.ValueInt64()
	req.AUTOMATION_ANALYTICS_LAST_ENTRIES = o.AUTOMATION_ANALYTICS_LAST_ENTRIES.ValueString()
	req.AUTOMATION_ANALYTICS_URL = o.AUTOMATION_ANALYTICS_URL.ValueString()
	req.CSRF_TRUSTED_ORIGINS = helpers.ListAsStringSlice(o.CSRF_TRUSTED_ORIGINS, false)
	req.DEFAULT_EXECUTION_ENVIRONMENT = o.DEFAULT_EXECUTION_ENVIRONMENT.ValueInt64()
	req.INSIGHTS_TRACKING_STATE = o.INSIGHTS_TRACKING_STATE.ValueBool()
	req.MANAGE_ORGANIZATION_AUTH = o.MANAGE_ORGANIZATION_AUTH.ValueBool()
	req.ORG_ADMINS_CAN_SEE_ALL_USERS = o.ORG_ADMINS_CAN_SEE_ALL_USERS.ValueBool()
	req.PROXY_IP_ALLOWED_LIST = helpers.ListAsStringSlice(o.PROXY_IP_ALLOWED_LIST, false)
	req.REDHAT_PASSWORD = o.REDHAT_PASSWORD.ValueString()
	req.REDHAT_USERNAME = o.REDHAT_USERNAME.ValueString()
	req.REMOTE_HOST_HEADERS = helpers.ListAsStringSlice(o.REMOTE_HOST_HEADERS, false)
	req.SUBSCRIPTIONS_PASSWORD = o.SUBSCRIPTIONS_PASSWORD.ValueString()
	req.SUBSCRIPTIONS_USERNAME = o.SUBSCRIPTIONS_USERNAME.ValueString()
	req.SUBSCRIPTION_USAGE_MODEL = o.SUBSCRIPTION_USAGE_MODEL.ValueString()
	req.TOWER_URL_BASE = o.TOWER_URL_BASE.ValueString()
	req.UI_NEXT = o.UI_NEXT.ValueBool()
	return &req
}

func (o *settingsMiscSystemTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ACTIVITY_STREAM_ENABLED, data["ACTIVITY_STREAM_ENABLED"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC, data["ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.AUTOMATION_ANALYTICS_GATHER_INTERVAL, data["AUTOMATION_ANALYTICS_GATHER_INTERVAL"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_LAST_ENTRIES, data["AUTOMATION_ANALYTICS_LAST_ENTRIES"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_LAST_GATHER, data["AUTOMATION_ANALYTICS_LAST_GATHER"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_URL, data["AUTOMATION_ANALYTICS_URL"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.CLEANUP_HOST_METRICS_LAST_TS, data["CLEANUP_HOST_METRICS_LAST_TS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.CSRF_TRUSTED_ORIGINS, data["CSRF_TRUSTED_ORIGINS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.CUSTOM_VENV_PATHS, data["CUSTOM_VENV_PATHS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.DEFAULT_CONTROL_PLANE_QUEUE_NAME, data["DEFAULT_CONTROL_PLANE_QUEUE_NAME"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.DEFAULT_EXECUTION_ENVIRONMENT, data["DEFAULT_EXECUTION_ENVIRONMENT"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.DEFAULT_EXECUTION_QUEUE_NAME, data["DEFAULT_EXECUTION_QUEUE_NAME"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.HOST_METRIC_SUMMARY_TASK_LAST_TS, data["HOST_METRIC_SUMMARY_TASK_LAST_TS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.INSIGHTS_TRACKING_STATE, data["INSIGHTS_TRACKING_STATE"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.INSTALL_UUID, data["INSTALL_UUID"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.IS_K8S, data["IS_K8S"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LICENSE, data["LICENSE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.MANAGE_ORGANIZATION_AUTH, data["MANAGE_ORGANIZATION_AUTH"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ORG_ADMINS_CAN_SEE_ALL_USERS, data["ORG_ADMINS_CAN_SEE_ALL_USERS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.PROXY_IP_ALLOWED_LIST, data["PROXY_IP_ALLOWED_LIST"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.REDHAT_PASSWORD, data["REDHAT_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.REDHAT_USERNAME, data["REDHAT_USERNAME"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.REMOTE_HOST_HEADERS, data["REMOTE_HOST_HEADERS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SUBSCRIPTIONS_PASSWORD, data["SUBSCRIPTIONS_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SUBSCRIPTIONS_USERNAME, data["SUBSCRIPTIONS_USERNAME"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SUBSCRIPTION_USAGE_MODEL, data["SUBSCRIPTION_USAGE_MODEL"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.TOWER_URL_BASE, data["TOWER_URL_BASE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.UI_NEXT, data["UI_NEXT"])
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
	// CSRF_TRUSTED_ORIGINS "If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values. "
	CSRF_TRUSTED_ORIGINS []string `json:"CSRF_TRUSTED_ORIGINS,omitempty"`
	// DEFAULT_EXECUTION_ENVIRONMENT "The Execution Environment to be used when one has not been configured for a job template."
	DEFAULT_EXECUTION_ENVIRONMENT int64 `json:"DEFAULT_EXECUTION_ENVIRONMENT,omitempty"`
	// INSIGHTS_TRACKING_STATE "Enables the service to gather data on automation and send it to Automation Analytics."
	INSIGHTS_TRACKING_STATE bool `json:"INSIGHTS_TRACKING_STATE"`
	// MANAGE_ORGANIZATION_AUTH "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration."
	MANAGE_ORGANIZATION_AUTH bool `json:"MANAGE_ORGANIZATION_AUTH"`
	// ORG_ADMINS_CAN_SEE_ALL_USERS "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization."
	ORG_ADMINS_CAN_SEE_ALL_USERS bool `json:"ORG_ADMINS_CAN_SEE_ALL_USERS"`
	// PROXY_IP_ALLOWED_LIST "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')"
	PROXY_IP_ALLOWED_LIST []string `json:"PROXY_IP_ALLOWED_LIST,omitempty"`
	// REDHAT_PASSWORD "This password is used to send data to Automation Analytics"
	REDHAT_PASSWORD string `json:"REDHAT_PASSWORD,omitempty"`
	// REDHAT_USERNAME "This username is used to send data to Automation Analytics"
	REDHAT_USERNAME string `json:"REDHAT_USERNAME,omitempty"`
	// REMOTE_HOST_HEADERS "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details."
	REMOTE_HOST_HEADERS []string `json:"REMOTE_HOST_HEADERS,omitempty"`
	// SUBSCRIPTIONS_PASSWORD "This password is used to retrieve subscription and content information"
	SUBSCRIPTIONS_PASSWORD string `json:"SUBSCRIPTIONS_PASSWORD,omitempty"`
	// SUBSCRIPTIONS_USERNAME "This username is used to retrieve subscription and content information"
	SUBSCRIPTIONS_USERNAME string `json:"SUBSCRIPTIONS_USERNAME,omitempty"`
	// SUBSCRIPTION_USAGE_MODEL ""
	SUBSCRIPTION_USAGE_MODEL string `json:"SUBSCRIPTION_USAGE_MODEL,omitempty"`
	// TOWER_URL_BASE "This setting is used by services like notifications to render a valid url to the service."
	TOWER_URL_BASE string `json:"TOWER_URL_BASE,omitempty"`
	// UI_NEXT "Enable preview of new user interface."
	UI_NEXT bool `json:"UI_NEXT"`
}
