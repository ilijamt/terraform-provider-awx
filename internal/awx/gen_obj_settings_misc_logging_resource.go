package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

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
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsMiscLoggingResource{}
	_ resource.ResourceWithConfigure = &settingsMiscLoggingResource{}
)

// NewSettingsMiscLoggingResource is a helper function to simplify the provider implementation.
func NewSettingsMiscLoggingResource() resource.Resource {
	return &settingsMiscLoggingResource{}
}

type settingsMiscLoggingResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsMiscLoggingResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/logging/"
}

func (o *settingsMiscLoggingResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_misc_logging"
}

func (o *settingsMiscLoggingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Request elements
			"api_400_error_log_format": schema.StringAttribute{
				Description: "The format of logged messages when an API 4XX error occurs, the following variables will be substituted: \nstatus_code - The HTTP status code of the error\nuser_name - The user name attempting to use the API\nurl_path - The URL path to the API endpoint called\nremote_addr - The remote address seen for the user\nerror - The error set by the api endpoint\nVariables need to be in the format {<variable name>}.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`status {status_code} received by user {user_name} attempting to access {url_path} from {remote_addr}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"log_aggregator_enabled": schema.BoolAttribute{
				Description: "Enable sending logs to external log aggregator.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"log_aggregator_host": schema.StringAttribute{
				Description: "Hostname/IP where external logs will be sent to.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"log_aggregator_individual_facts": schema.BoolAttribute{
				Description: "If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"log_aggregator_level": schema.StringAttribute{
				Description: "Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting)",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`INFO`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}...),
				},
			},
			"log_aggregator_loggers": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of loggers that will send HTTP logs to the collector, these can include any or all of: \nawx - service logs\nactivity_stream - activity stream records\njob_events - callback data from Ansible job events\nsystem_tracking - facts gathered from scan jobs.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"log_aggregator_max_disk_usage_gb": schema.Int64Attribute{
				Description: "Amount of data to store (in gigabytes) during an outage of the external log aggregator (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"log_aggregator_max_disk_usage_path": schema.StringAttribute{
				Description: "Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`/var/lib/awx`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"log_aggregator_password": schema.StringAttribute{
				Description: "Password or authentication token for external log aggregator (if required; HTTP/s only).",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"log_aggregator_port": schema.Int64Attribute{
				Description: "Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator).",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"log_aggregator_protocol": schema.StringAttribute{
				Description: "Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`https`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"https", "tcp", "udp"}...),
				},
			},
			"log_aggregator_rsyslogd_debug": schema.BoolAttribute{
				Description: "Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"log_aggregator_tcp_timeout": schema.Int64Attribute{
				Description: "Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(5),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"log_aggregator_tower_uuid": schema.StringAttribute{
				Description: "Useful to uniquely identify instances.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"log_aggregator_type": schema.StringAttribute{
				Description: "Format messages for the chosen log aggregator.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"logstash", "splunk", "loggly", "sumologic", "other"}...),
				},
			},
			"log_aggregator_username": schema.StringAttribute{
				Description: "Username for external log aggregator (if required; HTTP/s only).",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"log_aggregator_verify_cert": schema.BoolAttribute{
				Description: "Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is \"https\". If enabled, the log handler will verify certificate sent by external log aggregator before establishing connection.",
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
		},
	}
}

func (o *settingsMiscLoggingResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsMiscLoggingTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsMiscLogging
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsMiscLogging/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscLogging on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsMiscLogging resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsMiscLogging on %s", o.endpoint),
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

func (o *settingsMiscLoggingResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsMiscLoggingTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsMiscLogging
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscLogging on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsMiscLogging from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsMiscLogging on %s", o.endpoint),
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

func (o *settingsMiscLoggingResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsMiscLoggingTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsMiscLogging
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsMiscLogging/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscLogging on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsMiscLogging resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsMiscLogging on %s", o.endpoint),
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

func (o *settingsMiscLoggingResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
