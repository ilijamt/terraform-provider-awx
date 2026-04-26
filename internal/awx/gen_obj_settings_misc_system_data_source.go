package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsMiscSystemDataSource = framework.GenericDataSource[settingsMiscSystemTerraformModel, *settingsMiscSystemTerraformModel]

// NewSettingsMiscSystemDataSource is a helper function to instantiate the SettingsMiscSystem data source.
func NewSettingsMiscSystemDataSource() datasource.DataSource {
	return &settingsMiscSystemDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_system", Endpoint: "/api/v2/settings/system/"}},
		Cfg: framework.DataSourceCfg[settingsMiscSystemTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"activity_stream_enabled": schema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream.",
						Computed:    true,
					},
					"activity_stream_enabled_for_inventory_sync": schema.BoolAttribute{
						Description: "Enable capturing activity for the activity stream when running inventory sync.",
						Computed:    true,
					},
					"automation_analytics_gather_interval": schema.Int64Attribute{
						Description: "Interval (in seconds) between data gathering.",
						Computed:    true,
					},
					"automation_analytics_last_entries": schema.StringAttribute{
						Description: "Last gathered entries from the data collection service of Automation Analytics",
						Computed:    true,
					},
					"automation_analytics_last_gather": schema.StringAttribute{
						Description: "Last gather date for Automation Analytics.",
						Computed:    true,
					},
					"automation_analytics_url": schema.StringAttribute{
						Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
						Computed:    true,
					},
					"cleanup_host_metrics_last_ts": schema.StringAttribute{
						Description: "Last cleanup date for HostMetrics",
						Computed:    true,
					},
					"csrf_trusted_origins": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values. ",
						Computed:    true,
					},
					"custom_venv_paths": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
						Computed:    true,
					},
					"default_control_plane_queue_name": schema.StringAttribute{
						Description: "The instance group where control plane tasks run",
						Computed:    true,
					},
					"default_execution_environment": schema.Int64Attribute{
						Description: "The Execution Environment to be used when one has not been configured for a job template.",
						Computed:    true,
					},
					"default_execution_queue_name": schema.StringAttribute{
						Description: "The instance group where user jobs run (currently only on non-VM installs)",
						Computed:    true,
					},
					"host_metric_summary_task_last_ts": schema.StringAttribute{
						Description: "Last computing date of HostMetricSummaryMonthly",
						Computed:    true,
					},
					"insights_tracking_state": schema.BoolAttribute{
						Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
						Computed:    true,
					},
					"install_uuid": schema.StringAttribute{
						Description: "Unique identifier for an installation",
						Computed:    true,
					},
					"is_k8s": schema.BoolAttribute{
						Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
						Computed:    true,
					},
					"license": schema.StringAttribute{
						Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
						Computed:    true,
					},
					"manage_organization_auth": schema.BoolAttribute{
						Description: "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
						Computed:    true,
					},
					"org_admins_can_see_all_users": schema.BoolAttribute{
						Description: "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
						Computed:    true,
					},
					"proxy_ip_allowed_list": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
						Computed:    true,
					},
					"redhat_password": schema.StringAttribute{
						Description: "This password is used to send data to Automation Analytics",
						Computed:    true,
					},
					"redhat_username": schema.StringAttribute{
						Description: "This username is used to send data to Automation Analytics",
						Computed:    true,
					},
					"remote_host_headers": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
						Computed:    true,
					},
					"subscriptions_password": schema.StringAttribute{
						Description: "This password is used to retrieve subscription and content information",
						Computed:    true,
					},
					"subscriptions_username": schema.StringAttribute{
						Description: "This username is used to retrieve subscription and content information",
						Computed:    true,
					},
					"subscription_usage_model": schema.StringAttribute{
						Description: "Defines subscription usage model and shows Host Metrics",
						Computed:    true,
					},
					"tower_url_base": schema.StringAttribute{
						Description: "This setting is used by services like notifications to render a valid url to the service.",
						Computed:    true,
					},
					"ui_next": schema.BoolAttribute{
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
