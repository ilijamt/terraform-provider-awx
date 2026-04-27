package awx

import (
	"encoding/json"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsMiscSystemTerraformModel struct {
	ACTIVITY_STREAM_ENABLED                    types.Bool   `tfsdk:"activity_stream_enabled" json:"ACTIVITY_STREAM_ENABLED"`
	ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC types.Bool   `tfsdk:"activity_stream_enabled_for_inventory_sync" json:"ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"`
	AUTOMATION_ANALYTICS_GATHER_INTERVAL       types.Int64  `tfsdk:"automation_analytics_gather_interval" json:"AUTOMATION_ANALYTICS_GATHER_INTERVAL"`
	AUTOMATION_ANALYTICS_LAST_ENTRIES          types.String `tfsdk:"automation_analytics_last_entries" json:"AUTOMATION_ANALYTICS_LAST_ENTRIES"`
	AUTOMATION_ANALYTICS_LAST_GATHER           types.String `tfsdk:"automation_analytics_last_gather" json:"AUTOMATION_ANALYTICS_LAST_GATHER"`
	AUTOMATION_ANALYTICS_URL                   types.String `tfsdk:"automation_analytics_url" json:"AUTOMATION_ANALYTICS_URL"`
	CLEANUP_HOST_METRICS_LAST_TS               types.String `tfsdk:"cleanup_host_metrics_last_ts" json:"CLEANUP_HOST_METRICS_LAST_TS"`
	CSRF_TRUSTED_ORIGINS                       types.List   `tfsdk:"csrf_trusted_origins" json:"CSRF_TRUSTED_ORIGINS"`
	CUSTOM_VENV_PATHS                          types.List   `tfsdk:"custom_venv_paths" json:"CUSTOM_VENV_PATHS"`
	DEFAULT_CONTROL_PLANE_QUEUE_NAME           types.String `tfsdk:"default_control_plane_queue_name" json:"DEFAULT_CONTROL_PLANE_QUEUE_NAME"`
	DEFAULT_EXECUTION_ENVIRONMENT              types.Int64  `tfsdk:"default_execution_environment" json:"DEFAULT_EXECUTION_ENVIRONMENT"`
	DEFAULT_EXECUTION_QUEUE_NAME               types.String `tfsdk:"default_execution_queue_name" json:"DEFAULT_EXECUTION_QUEUE_NAME"`
	HOST_METRIC_SUMMARY_TASK_LAST_TS           types.String `tfsdk:"host_metric_summary_task_last_ts" json:"HOST_METRIC_SUMMARY_TASK_LAST_TS"`
	INSIGHTS_TRACKING_STATE                    types.Bool   `tfsdk:"insights_tracking_state" json:"INSIGHTS_TRACKING_STATE"`
	INSTALL_UUID                               types.String `tfsdk:"install_uuid" json:"INSTALL_UUID"`
	IS_K8S                                     types.Bool   `tfsdk:"is_k8s" json:"IS_K8S"`
	LICENSE                                    types.String `tfsdk:"license" json:"LICENSE"`
	MANAGE_ORGANIZATION_AUTH                   types.Bool   `tfsdk:"manage_organization_auth" json:"MANAGE_ORGANIZATION_AUTH"`
	ORG_ADMINS_CAN_SEE_ALL_USERS               types.Bool   `tfsdk:"org_admins_can_see_all_users" json:"ORG_ADMINS_CAN_SEE_ALL_USERS"`
	PROXY_IP_ALLOWED_LIST                      types.List   `tfsdk:"proxy_ip_allowed_list" json:"PROXY_IP_ALLOWED_LIST"`
	REDHAT_PASSWORD                            types.String `tfsdk:"redhat_password" json:"REDHAT_PASSWORD"`
	REDHAT_USERNAME                            types.String `tfsdk:"redhat_username" json:"REDHAT_USERNAME"`
	REMOTE_HOST_HEADERS                        types.List   `tfsdk:"remote_host_headers" json:"REMOTE_HOST_HEADERS"`
	SUBSCRIPTIONS_PASSWORD                     types.String `tfsdk:"subscriptions_password" json:"SUBSCRIPTIONS_PASSWORD"`
	SUBSCRIPTIONS_USERNAME                     types.String `tfsdk:"subscriptions_username" json:"SUBSCRIPTIONS_USERNAME"`
	SUBSCRIPTION_USAGE_MODEL                   types.String `tfsdk:"subscription_usage_model" json:"SUBSCRIPTION_USAGE_MODEL"`
	TOWER_URL_BASE                             types.String `tfsdk:"tower_url_base" json:"TOWER_URL_BASE"`
	UI_NEXT                                    types.Bool   `tfsdk:"ui_next" json:"UI_NEXT"`
}

func (o *settingsMiscSystemTerraformModel) Clone() settingsMiscSystemTerraformModel {
	return *o
}

func (o *settingsMiscSystemTerraformModel) BodyRequest() *settingsMiscSystemBodyRequestModel {
	var req settingsMiscSystemBodyRequestModel
	req.ACTIVITY_STREAM_ENABLED = o.ACTIVITY_STREAM_ENABLED.ValueBool()
	req.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC = o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC.ValueBool()
	req.AUTOMATION_ANALYTICS_GATHER_INTERVAL = o.AUTOMATION_ANALYTICS_GATHER_INTERVAL.ValueInt64()
	req.AUTOMATION_ANALYTICS_LAST_ENTRIES = json.RawMessage(o.AUTOMATION_ANALYTICS_LAST_ENTRIES.ValueString())
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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.ACTIVITY_STREAM_ENABLED, data["ACTIVITY_STREAM_ENABLED"]))
	collect(helpers.AttrValueSetBool(&o.ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC, data["ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"]))
	collect(helpers.AttrValueSetInt64(&o.AUTOMATION_ANALYTICS_GATHER_INTERVAL, data["AUTOMATION_ANALYTICS_GATHER_INTERVAL"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTOMATION_ANALYTICS_LAST_ENTRIES, data["AUTOMATION_ANALYTICS_LAST_ENTRIES"], false))
	collect(helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_LAST_GATHER, data["AUTOMATION_ANALYTICS_LAST_GATHER"], false))
	collect(helpers.AttrValueSetString(&o.AUTOMATION_ANALYTICS_URL, data["AUTOMATION_ANALYTICS_URL"], false))
	collect(helpers.AttrValueSetString(&o.CLEANUP_HOST_METRICS_LAST_TS, data["CLEANUP_HOST_METRICS_LAST_TS"], false))
	collect(helpers.AttrValueSetListString(&o.CSRF_TRUSTED_ORIGINS, data["CSRF_TRUSTED_ORIGINS"], false))
	collect(helpers.AttrValueSetListString(&o.CUSTOM_VENV_PATHS, data["CUSTOM_VENV_PATHS"], false))
	collect(helpers.AttrValueSetString(&o.DEFAULT_CONTROL_PLANE_QUEUE_NAME, data["DEFAULT_CONTROL_PLANE_QUEUE_NAME"], false))
	collect(helpers.AttrValueSetInt64(&o.DEFAULT_EXECUTION_ENVIRONMENT, data["DEFAULT_EXECUTION_ENVIRONMENT"]))
	collect(helpers.AttrValueSetString(&o.DEFAULT_EXECUTION_QUEUE_NAME, data["DEFAULT_EXECUTION_QUEUE_NAME"], false))
	collect(helpers.AttrValueSetString(&o.HOST_METRIC_SUMMARY_TASK_LAST_TS, data["HOST_METRIC_SUMMARY_TASK_LAST_TS"], false))
	collect(helpers.AttrValueSetBool(&o.INSIGHTS_TRACKING_STATE, data["INSIGHTS_TRACKING_STATE"]))
	collect(helpers.AttrValueSetString(&o.INSTALL_UUID, data["INSTALL_UUID"], false))
	collect(helpers.AttrValueSetBool(&o.IS_K8S, data["IS_K8S"]))
	collect(helpers.AttrValueSetJsonString(&o.LICENSE, data["LICENSE"], false))
	collect(helpers.AttrValueSetBool(&o.MANAGE_ORGANIZATION_AUTH, data["MANAGE_ORGANIZATION_AUTH"]))
	collect(helpers.AttrValueSetBool(&o.ORG_ADMINS_CAN_SEE_ALL_USERS, data["ORG_ADMINS_CAN_SEE_ALL_USERS"]))
	collect(helpers.AttrValueSetListString(&o.PROXY_IP_ALLOWED_LIST, data["PROXY_IP_ALLOWED_LIST"], false))
	collect(helpers.AttrValueSetString(&o.REDHAT_PASSWORD, data["REDHAT_PASSWORD"], false))
	collect(helpers.AttrValueSetString(&o.REDHAT_USERNAME, data["REDHAT_USERNAME"], false))
	collect(helpers.AttrValueSetListString(&o.REMOTE_HOST_HEADERS, data["REMOTE_HOST_HEADERS"], false))
	collect(helpers.AttrValueSetString(&o.SUBSCRIPTIONS_PASSWORD, data["SUBSCRIPTIONS_PASSWORD"], false))
	collect(helpers.AttrValueSetString(&o.SUBSCRIPTIONS_USERNAME, data["SUBSCRIPTIONS_USERNAME"], false))
	collect(helpers.AttrValueSetString(&o.SUBSCRIPTION_USAGE_MODEL, data["SUBSCRIPTION_USAGE_MODEL"], false))
	collect(helpers.AttrValueSetString(&o.TOWER_URL_BASE, data["TOWER_URL_BASE"], false))
	collect(helpers.AttrValueSetBool(&o.UI_NEXT, data["UI_NEXT"]))
	return diags, nil
}

type settingsMiscSystemBodyRequestModel struct {
	ACTIVITY_STREAM_ENABLED                    bool            `json:"ACTIVITY_STREAM_ENABLED"`
	ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC bool            `json:"ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC"`
	AUTOMATION_ANALYTICS_GATHER_INTERVAL       int64           `json:"AUTOMATION_ANALYTICS_GATHER_INTERVAL,omitempty"`
	AUTOMATION_ANALYTICS_LAST_ENTRIES          json.RawMessage `json:"AUTOMATION_ANALYTICS_LAST_ENTRIES,omitempty"`
	AUTOMATION_ANALYTICS_URL                   string          `json:"AUTOMATION_ANALYTICS_URL,omitempty"`
	CSRF_TRUSTED_ORIGINS                       []string        `json:"CSRF_TRUSTED_ORIGINS,omitempty"`
	DEFAULT_EXECUTION_ENVIRONMENT              int64           `json:"DEFAULT_EXECUTION_ENVIRONMENT,omitempty"`
	INSIGHTS_TRACKING_STATE                    bool            `json:"INSIGHTS_TRACKING_STATE"`
	MANAGE_ORGANIZATION_AUTH                   bool            `json:"MANAGE_ORGANIZATION_AUTH"`
	ORG_ADMINS_CAN_SEE_ALL_USERS               bool            `json:"ORG_ADMINS_CAN_SEE_ALL_USERS"`
	PROXY_IP_ALLOWED_LIST                      []string        `json:"PROXY_IP_ALLOWED_LIST,omitempty"`
	REDHAT_PASSWORD                            string          `json:"REDHAT_PASSWORD,omitempty"`
	REDHAT_USERNAME                            string          `json:"REDHAT_USERNAME,omitempty"`
	REMOTE_HOST_HEADERS                        []string        `json:"REMOTE_HOST_HEADERS,omitempty"`
	SUBSCRIPTIONS_PASSWORD                     string          `json:"SUBSCRIPTIONS_PASSWORD,omitempty"`
	SUBSCRIPTIONS_USERNAME                     string          `json:"SUBSCRIPTIONS_USERNAME,omitempty"`
	SUBSCRIPTION_USAGE_MODEL                   string          `json:"SUBSCRIPTION_USAGE_MODEL,omitempty"`
	TOWER_URL_BASE                             string          `json:"TOWER_URL_BASE,omitempty"`
	UI_NEXT                                    bool            `json:"UI_NEXT"`
}

type settingsMiscSystemResource = framework.GenericResource[settingsMiscSystemTerraformModel, settingsMiscSystemBodyRequestModel, *settingsMiscSystemTerraformModel]

// NewSettingsMiscSystemResource is a helper function to simplify the provider implementation.
func NewSettingsMiscSystemResource() resource.Resource {
	return &settingsMiscSystemResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_system", Endpoint: "/api/v2/settings/system/"}},
		Cfg: framework.ResourceCfg[settingsMiscSystemTerraformModel, settingsMiscSystemBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"activity_stream_enabled": schema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"activity_stream_enabled_for_inventory_sync": schema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream when running inventory sync.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"automation_analytics_gather_interval": schema.Int64Attribute{
						Description: "Interval (in seconds) between data gathering.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(14400),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"automation_analytics_last_entries": schema.StringAttribute{
						Description: "Last gathered entries from the data collection service of Automation Analytics",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"automation_analytics_url": schema.StringAttribute{
						Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`https://example.com`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"csrf_trusted_origins": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values. ",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"default_execution_environment": schema.Int64Attribute{
						Description: "The Execution Environment to be used when one has not been configured for a job template.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"insights_tracking_state": schema.BoolAttribute{
						Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"manage_organization_auth": schema.BoolAttribute{
						Description: "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"org_admins_can_see_all_users": schema.BoolAttribute{
						Description: "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"proxy_ip_allowed_list": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"redhat_password": schema.StringAttribute{
						Description: "This password is used to send data to Automation Analytics",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"redhat_username": schema.StringAttribute{
						Description: "This username is used to send data to Automation Analytics",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"remote_host_headers": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"subscriptions_password": schema.StringAttribute{
						Description: "This password is used to retrieve subscription and content information",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"subscriptions_username": schema.StringAttribute{
						Description: "This username is used to retrieve subscription and content information",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"subscription_usage_model": schema.StringAttribute{
						Description: "Defines subscription usage model and shows Host Metrics",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"unique_managed_hosts",
							),
						},
					},
					"tower_url_base": schema.StringAttribute{
						Description: "This setting is used by services like notifications to render a valid url to the service.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`https://localhost:8043`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"ui_next": schema.BoolAttribute{
						Description: "Enable preview of new user interface.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"automation_analytics_last_gather": schema.StringAttribute{
						Description: "Last gather date for Automation Analytics.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"cleanup_host_metrics_last_ts": schema.StringAttribute{
						Description: "Last cleanup date for HostMetrics",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"custom_venv_paths": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"default_control_plane_queue_name": schema.StringAttribute{
						Description: "The instance group where control plane tasks run",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"default_execution_queue_name": schema.StringAttribute{
						Description: "The instance group where user jobs run (currently only on non-VM installs)",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"host_metric_summary_task_last_ts": schema.StringAttribute{
						Description: "Last computing date of HostMetricSummaryMonthly",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"install_uuid": schema.StringAttribute{
						Description: "Unique identifier for an installation",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"is_k8s": schema.BoolAttribute{
						Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"license": schema.StringAttribute{
						Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsMiscSystem",
		},
	}
}

type settingsMiscSystemDataSource = framework.GenericDataSource[settingsMiscSystemTerraformModel, *settingsMiscSystemTerraformModel]

// NewSettingsMiscSystemDataSource is a helper function to instantiate the SettingsMiscSystem data source.
func NewSettingsMiscSystemDataSource() datasource.DataSource {
	return &settingsMiscSystemDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_system", Endpoint: "/api/v2/settings/system/"}},
		Cfg: framework.DataSourceCfg[settingsMiscSystemTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"activity_stream_enabled": dschema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream.",
						Computed:    true,
					},
					"activity_stream_enabled_for_inventory_sync": dschema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream when running inventory sync.",
						Computed:    true,
					},
					"automation_analytics_gather_interval": dschema.Int64Attribute{
						Description: "Interval (in seconds) between data gathering.",
						Computed:    true,
					},
					"automation_analytics_last_entries": dschema.StringAttribute{
						Description: "Last gathered entries from the data collection service of Automation Analytics",
						Computed:    true,
					},
					"automation_analytics_last_gather": dschema.StringAttribute{
						Description: "Last gather date for Automation Analytics.",
						Computed:    true,
					},
					"automation_analytics_url": dschema.StringAttribute{
						Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
						Computed:    true,
					},
					"cleanup_host_metrics_last_ts": dschema.StringAttribute{
						Description: "Last cleanup date for HostMetrics",
						Computed:    true,
					},
					"csrf_trusted_origins": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values. ",
						Computed:    true,
					},
					"custom_venv_paths": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
						Computed:    true,
					},
					"default_control_plane_queue_name": dschema.StringAttribute{
						Description: "The instance group where control plane tasks run",
						Computed:    true,
					},
					"default_execution_environment": dschema.Int64Attribute{
						Description: "The Execution Environment to be used when one has not been configured for a job template.",
						Computed:    true,
					},
					"default_execution_queue_name": dschema.StringAttribute{
						Description: "The instance group where user jobs run (currently only on non-VM installs)",
						Computed:    true,
					},
					"host_metric_summary_task_last_ts": dschema.StringAttribute{
						Description: "Last computing date of HostMetricSummaryMonthly",
						Computed:    true,
					},
					"insights_tracking_state": dschema.BoolAttribute{
						Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
						Computed:    true,
					},
					"install_uuid": dschema.StringAttribute{
						Description: "Unique identifier for an installation",
						Computed:    true,
					},
					"is_k8s": dschema.BoolAttribute{
						Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
						Computed:    true,
					},
					"license": dschema.StringAttribute{
						Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
						Computed:    true,
					},
					"manage_organization_auth": dschema.BoolAttribute{
						Description: "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
						Computed:    true,
					},
					"org_admins_can_see_all_users": dschema.BoolAttribute{
						Description: "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
						Computed:    true,
					},
					"proxy_ip_allowed_list": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
						Computed:    true,
					},
					"redhat_password": dschema.StringAttribute{
						Description: "This password is used to send data to Automation Analytics",
						Computed:    true,
					},
					"redhat_username": dschema.StringAttribute{
						Description: "This username is used to send data to Automation Analytics",
						Computed:    true,
					},
					"remote_host_headers": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
						Computed:    true,
					},
					"subscriptions_password": dschema.StringAttribute{
						Description: "This password is used to retrieve subscription and content information",
						Computed:    true,
					},
					"subscriptions_username": dschema.StringAttribute{
						Description: "This username is used to retrieve subscription and content information",
						Computed:    true,
					},
					"subscription_usage_model": dschema.StringAttribute{
						Description: "Defines subscription usage model and shows Host Metrics",
						Computed:    true,
					},
					"tower_url_base": dschema.StringAttribute{
						Description: "This setting is used by services like notifications to render a valid url to the service.",
						Computed:    true,
					},
					"ui_next": dschema.BoolAttribute{
						Description: "Enable preview of new user interface.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsMiscSystem",
		},
	}
}
