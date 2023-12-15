package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &settingsMiscLoggingDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsMiscLoggingDataSource{}
)

// NewSettingsMiscLoggingDataSource is a helper function to instantiate the SettingsMiscLogging data source.
func NewSettingsMiscLoggingDataSource() datasource.DataSource {
	return &settingsMiscLoggingDataSource{}
}

// settingsMiscLoggingDataSource is the data source implementation.
type settingsMiscLoggingDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsMiscLoggingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/logging/"
}

// Metadata returns the data source type name.
func (o *settingsMiscLoggingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_misc_logging"
}

// Schema defines the schema for the data source.
func (o *settingsMiscLoggingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"api_400_error_log_format": schema.StringAttribute{
				Description: "The format of logged messages when an API 4XX error occurs, the following variables will be substituted: \nstatus_code - The HTTP status code of the error\nuser_name - The user name attempting to use the API\nurl_path - The URL path to the API endpoint called\nremote_addr - The remote address seen for the user\nerror - The error set by the api endpoint\nVariables need to be in the format {<variable name>}.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"log_aggregator_action_max_disk_usage_gb": schema.Int64Attribute{
				Description: "Amount of data to store (in gigabytes) if an rsyslog action takes time to process an incoming message (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting on the action (e.g. omhttp). It stores files in the directory specified by LOG_AGGREGATOR_MAX_DISK_USAGE_PATH.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"log_aggregator_action_queue_size": schema.Int64Attribute{
				Description: "Defines how large the rsyslog action queue can grow in number of messages stored. This can have an impact on memory utilization. When the queue reaches 75% of this number, the queue will start writing to disk (queue.highWatermark in rsyslog). When it reaches 90%, NOTICE, INFO, and DEBUG messages will start to be discarded (queue.discardMark with queue.discardSeverity=5).",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"log_aggregator_enabled": schema.BoolAttribute{
				Description: "Enable sending logs to external log aggregator.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"log_aggregator_host": schema.StringAttribute{
				Description: "Hostname/IP where external logs will be sent to.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"log_aggregator_individual_facts": schema.BoolAttribute{
				Description: "If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"log_aggregator_level": schema.StringAttribute{
				Description: "Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting)",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}...),
				},
			},
			"log_aggregator_loggers": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of loggers that will send HTTP logs to the collector, these can include any or all of: \nawx - service logs\nactivity_stream - activity stream records\njob_events - callback data from Ansible job events\nsystem_tracking - facts gathered from scan jobs\nbroadcast_websocket - errors pertaining to websockets broadcast metrics\n",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.List{},
			},
			"log_aggregator_max_disk_usage_path": schema.StringAttribute{
				Description: "Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"log_aggregator_password": schema.StringAttribute{
				Description: "Password or authentication token for external log aggregator (if required; HTTP/s only).",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"log_aggregator_port": schema.Int64Attribute{
				Description: "Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator).",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"log_aggregator_protocol": schema.StringAttribute{
				Description: "Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"https", "tcp", "udp"}...),
				},
			},
			"log_aggregator_rsyslogd_debug": schema.BoolAttribute{
				Description: "Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"log_aggregator_tcp_timeout": schema.Int64Attribute{
				Description: "Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"log_aggregator_tower_uuid": schema.StringAttribute{
				Description: "Useful to uniquely identify instances.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"log_aggregator_type": schema.StringAttribute{
				Description: "Format messages for the chosen log aggregator.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"logstash", "splunk", "loggly", "sumologic", "other"}...),
				},
			},
			"log_aggregator_username": schema.StringAttribute{
				Description: "Username for external log aggregator (if required; HTTP/s only).",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"log_aggregator_verify_cert": schema.BoolAttribute{
				Description: "Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is \"https\". If enabled, the log handler will verify certificate sent by external log aggregator before establishing connection.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			// Write only elements
		},
	}
}

func (o *settingsMiscLoggingDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsMiscLoggingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsMiscLoggingTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsMiscLogging
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscLogging on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsMiscLogging
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsMiscLogging on %s", endpoint),
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
