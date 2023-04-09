package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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

// Schema defines the schema for the data source.
func (o *settingsMiscSystemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"activity_stream_enabled": schema.BoolAttribute{
				Description: "Enable capturing activity for the activity stream.",
				Sensitive:   false,
				Computed:    true,
			},
			"activity_stream_enabled_for_inventory_sync": schema.BoolAttribute{
				Description: "Enable capturing activity for the activity stream when running inventory sync.",
				Sensitive:   false,
				Computed:    true,
			},
			"automation_analytics_gather_interval": schema.Int64Attribute{
				Description: "Interval (in seconds) between data gathering.",
				Sensitive:   false,
				Computed:    true,
			},
			"automation_analytics_last_entries": schema.StringAttribute{
				Description: "Last gathered entries from the data collection service of Automation Analytics",
				Sensitive:   false,
				Computed:    true,
			},
			"automation_analytics_last_gather": schema.StringAttribute{
				Description: "Last gather date for Automation Analytics.",
				Sensitive:   false,
				Computed:    true,
			},
			"automation_analytics_url": schema.StringAttribute{
				Description: "This setting is used to to configure the upload URL for data collection for Automation Analytics.",
				Sensitive:   false,
				Computed:    true,
			},
			"custom_venv_paths": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line.",
				Sensitive:   false,
				Computed:    true,
			},
			"default_control_plane_queue_name": schema.StringAttribute{
				Description: "The instance group where control plane tasks run",
				Sensitive:   false,
				Computed:    true,
			},
			"default_execution_environment": schema.Int64Attribute{
				Description: "The Execution Environment to be used when one has not been configured for a job template.",
				Sensitive:   false,
				Computed:    true,
			},
			"default_execution_queue_name": schema.StringAttribute{
				Description: "The instance group where user jobs run (currently only on non-VM installs)",
				Sensitive:   false,
				Computed:    true,
			},
			"insights_tracking_state": schema.BoolAttribute{
				Description: "Enables the service to gather data on automation and send it to Automation Analytics.",
				Sensitive:   false,
				Computed:    true,
			},
			"install_uuid": schema.StringAttribute{
				Description: "Unique identifier for an installation",
				Sensitive:   false,
				Computed:    true,
			},
			"is_k8s": schema.BoolAttribute{
				Description: "Indicates whether the instance is part of a kubernetes-based deployment.",
				Sensitive:   false,
				Computed:    true,
			},
			"license": schema.StringAttribute{
				Description: "The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license.",
				Sensitive:   false,
				Computed:    true,
			},
			"manage_organization_auth": schema.BoolAttribute{
				Description: "Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration.",
				Sensitive:   false,
				Computed:    true,
			},
			"org_admins_can_see_all_users": schema.BoolAttribute{
				Description: "Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization.",
				Sensitive:   false,
				Computed:    true,
			},
			"proxy_ip_allowed_list": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally')",
				Sensitive:   false,
				Computed:    true,
			},
			"redhat_password": schema.StringAttribute{
				Description: "This password is used to send data to Automation Analytics",
				Sensitive:   false,
				Computed:    true,
			},
			"redhat_username": schema.StringAttribute{
				Description: "This username is used to send data to Automation Analytics",
				Sensitive:   false,
				Computed:    true,
			},
			"remote_host_headers": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \"HTTP_X_FORWARDED_FOR\", if behind a reverse proxy. See the \"Proxy Support\" section of the AAP Installation guide for more details.",
				Sensitive:   false,
				Computed:    true,
			},
			"subscriptions_password": schema.StringAttribute{
				Description: "This password is used to retrieve subscription and content information",
				Sensitive:   false,
				Computed:    true,
			},
			"subscriptions_username": schema.StringAttribute{
				Description: "This username is used to retrieve subscription and content information",
				Sensitive:   false,
				Computed:    true,
			},
			"tower_url_base": schema.StringAttribute{
				Description: "This value has been set manually in a settings file.\n\nThis setting is used by services like notifications to render a valid url to the service.",
				Sensitive:   false,
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *settingsMiscSystemDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsMiscSystemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsMiscSystemTerraformModel
	var err error
	var endpoint = o.endpoint

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
