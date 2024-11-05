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
	return settingsMiscSystemTerraformModel{
		ACTIVITY_STREAM_ENABLED:                    o.ACTIVITY_STREAM_ENABLED,
		ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC: o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC,
		AUTOMATION_ANALYTICS_GATHER_INTERVAL:       o.AUTOMATION_ANALYTICS_GATHER_INTERVAL,
		AUTOMATION_ANALYTICS_LAST_ENTRIES:          o.AUTOMATION_ANALYTICS_LAST_ENTRIES,
		AUTOMATION_ANALYTICS_LAST_GATHER:           o.AUTOMATION_ANALYTICS_LAST_GATHER,
		AUTOMATION_ANALYTICS_URL:                   o.AUTOMATION_ANALYTICS_URL,
		CLEANUP_HOST_METRICS_LAST_TS:               o.CLEANUP_HOST_METRICS_LAST_TS,
		CSRF_TRUSTED_ORIGINS:                       o.CSRF_TRUSTED_ORIGINS,
		CUSTOM_VENV_PATHS:                          o.CUSTOM_VENV_PATHS,
		DEFAULT_CONTROL_PLANE_QUEUE_NAME:           o.DEFAULT_CONTROL_PLANE_QUEUE_NAME,
		DEFAULT_EXECUTION_ENVIRONMENT:              o.DEFAULT_EXECUTION_ENVIRONMENT,
		DEFAULT_EXECUTION_QUEUE_NAME:               o.DEFAULT_EXECUTION_QUEUE_NAME,
		HOST_METRIC_SUMMARY_TASK_LAST_TS:           o.HOST_METRIC_SUMMARY_TASK_LAST_TS,
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
		SUBSCRIPTION_USAGE_MODEL:                   o.SUBSCRIPTION_USAGE_MODEL,
		TOWER_URL_BASE:                             o.TOWER_URL_BASE,
		UI_NEXT:                                    o.UI_NEXT,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsMiscSystem
func (o *settingsMiscSystemTerraformModel) BodyRequest() (req settingsMiscSystemBodyRequestModel) {
	req.ACTIVITY_STREAM_ENABLED = o.ACTIVITY_STREAM_ENABLED.ValueBool()
	req.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC = o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC.ValueBool()
	req.AUTOMATION_ANALYTICS_GATHER_INTERVAL = o.AUTOMATION_ANALYTICS_GATHER_INTERVAL.ValueInt64()
	req.AUTOMATION_ANALYTICS_LAST_ENTRIES = o.AUTOMATION_ANALYTICS_LAST_ENTRIES.ValueString()
	req.AUTOMATION_ANALYTICS_URL = o.AUTOMATION_ANALYTICS_URL.ValueString()
	req.CSRF_TRUSTED_ORIGINS = []string{}
	for _, val := range o.CSRF_TRUSTED_ORIGINS.Elements() {
		if _, ok := val.(types.String); ok {
			req.CSRF_TRUSTED_ORIGINS = append(req.CSRF_TRUSTED_ORIGINS, val.(types.String).ValueString())
		} else {
			req.CSRF_TRUSTED_ORIGINS = append(req.CSRF_TRUSTED_ORIGINS, val.String())
		}
	}
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
	req.SUBSCRIPTION_USAGE_MODEL = o.SUBSCRIPTION_USAGE_MODEL.ValueString()
	req.TOWER_URL_BASE = o.TOWER_URL_BASE.ValueString()
	req.UI_NEXT = o.UI_NEXT.ValueBool()
	return
}

func (o *settingsMiscSystemTerraformModel) setActivityStreamEnabled(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.ACTIVITY_STREAM_ENABLED, data)
}

func (o *settingsMiscSystemTerraformModel) setActivityStreamEnabledForInventorySync(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC, data)
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsGatherInterval(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.AUTOMATION_ANALYTICS_GATHER_INTERVAL, data)
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsLastEntries(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_LAST_ENTRIES, data, false)
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsLastGather(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_LAST_GATHER, data, false)
}

func (o *settingsMiscSystemTerraformModel) setAutomationAnalyticsUrl(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_URL, data, false)
}

func (o *settingsMiscSystemTerraformModel) setCleanupHostMetricsLastTs(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.CLEANUP_HOST_METRICS_LAST_TS, data, false)
}

func (o *settingsMiscSystemTerraformModel) setCsrfTrustedOrigins(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetListString(&o.CSRF_TRUSTED_ORIGINS, data, false)
}

func (o *settingsMiscSystemTerraformModel) setCustomVenvPaths(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetListString(&o.CUSTOM_VENV_PATHS, data, false)
}

func (o *settingsMiscSystemTerraformModel) setDefaultControlPlaneQueueName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.DEFAULT_CONTROL_PLANE_QUEUE_NAME, data, false)
}

func (o *settingsMiscSystemTerraformModel) setDefaultExecutionEnvironment(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.DEFAULT_EXECUTION_ENVIRONMENT, data)
}

func (o *settingsMiscSystemTerraformModel) setDefaultExecutionQueueName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.DEFAULT_EXECUTION_QUEUE_NAME, data, false)
}

func (o *settingsMiscSystemTerraformModel) setHostMetricSummaryTaskLastTs(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.HOST_METRIC_SUMMARY_TASK_LAST_TS, data, false)
}

func (o *settingsMiscSystemTerraformModel) setInsightsTrackingState(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.INSIGHTS_TRACKING_STATE, data)
}

func (o *settingsMiscSystemTerraformModel) setInstallUuid(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.INSTALL_UUID, data, false)
}

func (o *settingsMiscSystemTerraformModel) setIsK8S(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.IS_K8S, data)
}

func (o *settingsMiscSystemTerraformModel) setLicense(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.LICENSE, data, false)
}

func (o *settingsMiscSystemTerraformModel) setManageOrganizationAuth(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.MANAGE_ORGANIZATION_AUTH, data)
}

func (o *settingsMiscSystemTerraformModel) setOrgAdminsCanSeeAllUsers(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.ORG_ADMINS_CAN_SEE_ALL_USERS, data)
}

func (o *settingsMiscSystemTerraformModel) setProxyIpAllowedList(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetListString(&o.PROXY_IP_ALLOWED_LIST, data, false)
}

func (o *settingsMiscSystemTerraformModel) setRedhatPassword(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.REDHAT_PASSWORD, data, false)
}

func (o *settingsMiscSystemTerraformModel) setRedhatUsername(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.REDHAT_USERNAME, data, false)
}

func (o *settingsMiscSystemTerraformModel) setRemoteHostHeaders(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetListString(&o.REMOTE_HOST_HEADERS, data, false)
}

func (o *settingsMiscSystemTerraformModel) setSubscriptionsPassword(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SUBSCRIPTIONS_PASSWORD, data, false)
}

func (o *settingsMiscSystemTerraformModel) setSubscriptionsUsername(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SUBSCRIPTIONS_USERNAME, data, false)
}

func (o *settingsMiscSystemTerraformModel) setSubscriptionUsageModel(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SUBSCRIPTION_USAGE_MODEL, data, false)
}

func (o *settingsMiscSystemTerraformModel) setTowerUrlBase(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.TOWER_URL_BASE, data, false)
}

func (o *settingsMiscSystemTerraformModel) setUiNext(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.UI_NEXT, data)
}

func (o *settingsMiscSystemTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
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
	if dg, _ := o.setCleanupHostMetricsLastTs(data["CLEANUP_HOST_METRICS_LAST_TS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCsrfTrustedOrigins(data["CSRF_TRUSTED_ORIGINS"]); dg.HasError() {
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
	if dg, _ := o.setHostMetricSummaryTaskLastTs(data["HOST_METRIC_SUMMARY_TASK_LAST_TS"]); dg.HasError() {
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
	if dg, _ := o.setSubscriptionUsageModel(data["SUBSCRIPTION_USAGE_MODEL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTowerUrlBase(data["TOWER_URL_BASE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUiNext(data["UI_NEXT"]); dg.HasError() {
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
