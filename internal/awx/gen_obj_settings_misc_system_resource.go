package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsMiscSystemResource = framework.GenericResource[settingsMiscSystemTerraformModel, settingsMiscSystemBodyRequestModel, *settingsMiscSystemTerraformModel]

// NewSettingsMiscSystemResource is a helper function to simplify the provider implementation.
func NewSettingsMiscSystemResource() resource.Resource {
	return &settingsMiscSystemResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_system", Endpoint: "/api/v2/settings/system/"}},
		Cfg: framework.ResourceCfg[settingsMiscSystemTerraformModel, settingsMiscSystemBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"activity_stream_enabled": schema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"activity_stream_enabled_for_inventory_sync": schema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream when running inventory sync.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"automation_analytics_gather_interval": schema.Int64Attribute{
						Description: "Interval (in seconds) between data gathering.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(14400),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"automation_analytics_last_entries": schema.StringAttribute{
						Description: "Last gathered entries from the data collection service of Automation Analytics",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"automation_analytics_url": schema.StringAttribute{
						Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`https://example.com`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"csrf_trusted_origins": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values. ",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.List{},
					},
					"default_execution_environment": schema.Int64Attribute{
						Description: "The Execution Environment to be used when one has not been configured for a job template.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"insights_tracking_state": schema.BoolAttribute{
						Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"manage_organization_auth": schema.BoolAttribute{
						Description: "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"org_admins_can_see_all_users": schema.BoolAttribute{
						Description: "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"proxy_ip_allowed_list": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.List{},
					},
					"redhat_password": schema.StringAttribute{
						Description: "This password is used to send data to Automation Analytics",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"redhat_username": schema.StringAttribute{
						Description: "This username is used to send data to Automation Analytics",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"remote_host_headers": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.List{},
					},
					"subscriptions_password": schema.StringAttribute{
						Description: "This password is used to retrieve subscription and content information",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"subscriptions_username": schema.StringAttribute{
						Description: "This username is used to retrieve subscription and content information",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"subscription_usage_model": schema.StringAttribute{
						Description: "Defines subscription usage model and shows Host Metrics",
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`https://localhost:8043`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"ui_next": schema.BoolAttribute{
						Description: "Enable preview of new user interface.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					// Write only elements
					// Data only elements
					"automation_analytics_last_gather": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"cleanup_host_metrics_last_ts": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"custom_venv_paths": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"default_control_plane_queue_name": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"default_execution_queue_name": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"host_metric_summary_task_last_ts": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"install_uuid": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"is_k8s": schema.BoolAttribute{
						Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"license": schema.StringAttribute{
						Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
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
