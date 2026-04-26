package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsMiscLoggingTerraformModel maps the schema for SettingsMiscLogging when using Data Source
type settingsMiscLoggingTerraformModel struct {
	// API_400_ERROR_LOG_FORMAT "The format of logged messages when an API 4XX error occurs, the following variables will be substituted: \nstatus_code - The HTTP status code of the error\nuser_name - The user name attempting to use the API\nurl_path - The URL path to the API endpoint called\nremote_addr - The remote address seen for the user\nerror - The error set by the api endpoint\nVariables need to be in the format {<variable name>}."
	API_400_ERROR_LOG_FORMAT types.String `tfsdk:"api_400_error_log_format" json:"API_400_ERROR_LOG_FORMAT"`
	// LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB "Amount of data to store (in gigabytes) if an rsyslog action takes time to process an incoming message (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting on the action (e.g. omhttp). It stores files in the directory specified by LOG_AGGREGATOR_MAX_DISK_USAGE_PATH."
	LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB types.Int64 `tfsdk:"log_aggregator_action_max_disk_usage_gb" json:"LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB"`
	// LOG_AGGREGATOR_ACTION_QUEUE_SIZE "Defines how large the rsyslog action queue can grow in number of messages stored. This can have an impact on memory utilization. When the queue reaches 75% of this number, the queue will start writing to disk (queue.highWatermark in rsyslog). When it reaches 90%, NOTICE, INFO, and DEBUG messages will start to be discarded (queue.discardMark with queue.discardSeverity=5)."
	LOG_AGGREGATOR_ACTION_QUEUE_SIZE types.Int64 `tfsdk:"log_aggregator_action_queue_size" json:"LOG_AGGREGATOR_ACTION_QUEUE_SIZE"`
	// LOG_AGGREGATOR_ENABLED "Enable sending logs to external log aggregator."
	LOG_AGGREGATOR_ENABLED types.Bool `tfsdk:"log_aggregator_enabled" json:"LOG_AGGREGATOR_ENABLED"`
	// LOG_AGGREGATOR_HOST "Hostname/IP where external logs will be sent to."
	LOG_AGGREGATOR_HOST types.String `tfsdk:"log_aggregator_host" json:"LOG_AGGREGATOR_HOST"`
	// LOG_AGGREGATOR_INDIVIDUAL_FACTS "If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing."
	LOG_AGGREGATOR_INDIVIDUAL_FACTS types.Bool `tfsdk:"log_aggregator_individual_facts" json:"LOG_AGGREGATOR_INDIVIDUAL_FACTS"`
	// LOG_AGGREGATOR_LEVEL "Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting)"
	LOG_AGGREGATOR_LEVEL types.String `tfsdk:"log_aggregator_level" json:"LOG_AGGREGATOR_LEVEL"`
	// LOG_AGGREGATOR_LOGGERS "List of loggers that will send HTTP logs to the collector, these can include any or all of: \nawx - service logs\nactivity_stream - activity stream records\njob_events - callback data from Ansible job events\nsystem_tracking - facts gathered from scan jobs\nbroadcast_websocket - errors pertaining to websockets broadcast metrics\n"
	LOG_AGGREGATOR_LOGGERS types.List `tfsdk:"log_aggregator_loggers" json:"LOG_AGGREGATOR_LOGGERS"`
	// LOG_AGGREGATOR_MAX_DISK_USAGE_PATH "Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting."
	LOG_AGGREGATOR_MAX_DISK_USAGE_PATH types.String `tfsdk:"log_aggregator_max_disk_usage_path" json:"LOG_AGGREGATOR_MAX_DISK_USAGE_PATH"`
	// LOG_AGGREGATOR_PASSWORD "Password or authentication token for external log aggregator (if required; HTTP/s only)."
	LOG_AGGREGATOR_PASSWORD types.String `tfsdk:"log_aggregator_password" json:"LOG_AGGREGATOR_PASSWORD"`
	// LOG_AGGREGATOR_PORT "Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator)."
	LOG_AGGREGATOR_PORT types.Int64 `tfsdk:"log_aggregator_port" json:"LOG_AGGREGATOR_PORT"`
	// LOG_AGGREGATOR_PROTOCOL "Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname."
	LOG_AGGREGATOR_PROTOCOL types.String `tfsdk:"log_aggregator_protocol" json:"LOG_AGGREGATOR_PROTOCOL"`
	// LOG_AGGREGATOR_RSYSLOGD_DEBUG "Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation."
	LOG_AGGREGATOR_RSYSLOGD_DEBUG types.Bool `tfsdk:"log_aggregator_rsyslogd_debug" json:"LOG_AGGREGATOR_RSYSLOGD_DEBUG"`
	// LOG_AGGREGATOR_TCP_TIMEOUT "Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols."
	LOG_AGGREGATOR_TCP_TIMEOUT types.Int64 `tfsdk:"log_aggregator_tcp_timeout" json:"LOG_AGGREGATOR_TCP_TIMEOUT"`
	// LOG_AGGREGATOR_TOWER_UUID "Useful to uniquely identify instances."
	LOG_AGGREGATOR_TOWER_UUID types.String `tfsdk:"log_aggregator_tower_uuid" json:"LOG_AGGREGATOR_TOWER_UUID"`
	// LOG_AGGREGATOR_TYPE "Format messages for the chosen log aggregator."
	LOG_AGGREGATOR_TYPE types.String `tfsdk:"log_aggregator_type" json:"LOG_AGGREGATOR_TYPE"`
	// LOG_AGGREGATOR_USERNAME "Username for external log aggregator (if required; HTTP/s only)."
	LOG_AGGREGATOR_USERNAME types.String `tfsdk:"log_aggregator_username" json:"LOG_AGGREGATOR_USERNAME"`
	// LOG_AGGREGATOR_VERIFY_CERT "Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is \"https\". If enabled, the log handler will verify certificate sent by external log aggregator before establishing connection."
	LOG_AGGREGATOR_VERIFY_CERT types.Bool `tfsdk:"log_aggregator_verify_cert" json:"LOG_AGGREGATOR_VERIFY_CERT"`
}

// Clone the object
func (o *settingsMiscLoggingTerraformModel) Clone() settingsMiscLoggingTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsMiscLogging
func (o *settingsMiscLoggingTerraformModel) BodyRequest() *settingsMiscLoggingBodyRequestModel {
	var req settingsMiscLoggingBodyRequestModel
	req.API_400_ERROR_LOG_FORMAT = o.API_400_ERROR_LOG_FORMAT.ValueString()
	req.LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB = o.LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB.ValueInt64()
	req.LOG_AGGREGATOR_ACTION_QUEUE_SIZE = o.LOG_AGGREGATOR_ACTION_QUEUE_SIZE.ValueInt64()
	req.LOG_AGGREGATOR_ENABLED = o.LOG_AGGREGATOR_ENABLED.ValueBool()
	req.LOG_AGGREGATOR_HOST = o.LOG_AGGREGATOR_HOST.ValueString()
	req.LOG_AGGREGATOR_INDIVIDUAL_FACTS = o.LOG_AGGREGATOR_INDIVIDUAL_FACTS.ValueBool()
	req.LOG_AGGREGATOR_LEVEL = o.LOG_AGGREGATOR_LEVEL.ValueString()
	req.LOG_AGGREGATOR_LOGGERS = helpers.ListAsStringSlice(o.LOG_AGGREGATOR_LOGGERS, false)
	req.LOG_AGGREGATOR_MAX_DISK_USAGE_PATH = o.LOG_AGGREGATOR_MAX_DISK_USAGE_PATH.ValueString()
	req.LOG_AGGREGATOR_PASSWORD = o.LOG_AGGREGATOR_PASSWORD.ValueString()
	req.LOG_AGGREGATOR_PORT = o.LOG_AGGREGATOR_PORT.ValueInt64()
	req.LOG_AGGREGATOR_PROTOCOL = o.LOG_AGGREGATOR_PROTOCOL.ValueString()
	req.LOG_AGGREGATOR_RSYSLOGD_DEBUG = o.LOG_AGGREGATOR_RSYSLOGD_DEBUG.ValueBool()
	req.LOG_AGGREGATOR_TCP_TIMEOUT = o.LOG_AGGREGATOR_TCP_TIMEOUT.ValueInt64()
	req.LOG_AGGREGATOR_TOWER_UUID = o.LOG_AGGREGATOR_TOWER_UUID.ValueString()
	req.LOG_AGGREGATOR_TYPE = o.LOG_AGGREGATOR_TYPE.ValueString()
	req.LOG_AGGREGATOR_USERNAME = o.LOG_AGGREGATOR_USERNAME.ValueString()
	req.LOG_AGGREGATOR_VERIFY_CERT = o.LOG_AGGREGATOR_VERIFY_CERT.ValueBool()
	return &req
}

func (o *settingsMiscLoggingTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.API_400_ERROR_LOG_FORMAT, data["API_400_ERROR_LOG_FORMAT"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB, data["LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_ACTION_QUEUE_SIZE, data["LOG_AGGREGATOR_ACTION_QUEUE_SIZE"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_ENABLED, data["LOG_AGGREGATOR_ENABLED"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_HOST, data["LOG_AGGREGATOR_HOST"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_INDIVIDUAL_FACTS, data["LOG_AGGREGATOR_INDIVIDUAL_FACTS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_LEVEL, data["LOG_AGGREGATOR_LEVEL"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.LOG_AGGREGATOR_LOGGERS, data["LOG_AGGREGATOR_LOGGERS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_MAX_DISK_USAGE_PATH, data["LOG_AGGREGATOR_MAX_DISK_USAGE_PATH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_PASSWORD, data["LOG_AGGREGATOR_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_PORT, data["LOG_AGGREGATOR_PORT"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_PROTOCOL, data["LOG_AGGREGATOR_PROTOCOL"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_RSYSLOGD_DEBUG, data["LOG_AGGREGATOR_RSYSLOGD_DEBUG"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_TCP_TIMEOUT, data["LOG_AGGREGATOR_TCP_TIMEOUT"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_TOWER_UUID, data["LOG_AGGREGATOR_TOWER_UUID"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_TYPE, data["LOG_AGGREGATOR_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOG_AGGREGATOR_USERNAME, data["LOG_AGGREGATOR_USERNAME"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_VERIFY_CERT, data["LOG_AGGREGATOR_VERIFY_CERT"])
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsMiscLoggingBodyRequestModel maps the schema for SettingsMiscLogging for creating and updating the data
type settingsMiscLoggingBodyRequestModel struct {
	// API_400_ERROR_LOG_FORMAT "The format of logged messages when an API 4XX error occurs, the following variables will be substituted: \nstatus_code - The HTTP status code of the error\nuser_name - The user name attempting to use the API\nurl_path - The URL path to the API endpoint called\nremote_addr - The remote address seen for the user\nerror - The error set by the api endpoint\nVariables need to be in the format {<variable name>}."
	API_400_ERROR_LOG_FORMAT string `json:"API_400_ERROR_LOG_FORMAT,omitempty"`
	// LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB "Amount of data to store (in gigabytes) if an rsyslog action takes time to process an incoming message (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting on the action (e.g. omhttp). It stores files in the directory specified by LOG_AGGREGATOR_MAX_DISK_USAGE_PATH."
	LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB int64 `json:"LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB,omitempty"`
	// LOG_AGGREGATOR_ACTION_QUEUE_SIZE "Defines how large the rsyslog action queue can grow in number of messages stored. This can have an impact on memory utilization. When the queue reaches 75% of this number, the queue will start writing to disk (queue.highWatermark in rsyslog). When it reaches 90%, NOTICE, INFO, and DEBUG messages will start to be discarded (queue.discardMark with queue.discardSeverity=5)."
	LOG_AGGREGATOR_ACTION_QUEUE_SIZE int64 `json:"LOG_AGGREGATOR_ACTION_QUEUE_SIZE,omitempty"`
	// LOG_AGGREGATOR_ENABLED "Enable sending logs to external log aggregator."
	LOG_AGGREGATOR_ENABLED bool `json:"LOG_AGGREGATOR_ENABLED"`
	// LOG_AGGREGATOR_HOST "Hostname/IP where external logs will be sent to."
	LOG_AGGREGATOR_HOST string `json:"LOG_AGGREGATOR_HOST,omitempty"`
	// LOG_AGGREGATOR_INDIVIDUAL_FACTS "If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing."
	LOG_AGGREGATOR_INDIVIDUAL_FACTS bool `json:"LOG_AGGREGATOR_INDIVIDUAL_FACTS"`
	// LOG_AGGREGATOR_LEVEL "Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting)"
	LOG_AGGREGATOR_LEVEL string `json:"LOG_AGGREGATOR_LEVEL,omitempty"`
	// LOG_AGGREGATOR_LOGGERS "List of loggers that will send HTTP logs to the collector, these can include any or all of: \nawx - service logs\nactivity_stream - activity stream records\njob_events - callback data from Ansible job events\nsystem_tracking - facts gathered from scan jobs\nbroadcast_websocket - errors pertaining to websockets broadcast metrics\n"
	LOG_AGGREGATOR_LOGGERS []string `json:"LOG_AGGREGATOR_LOGGERS,omitempty"`
	// LOG_AGGREGATOR_MAX_DISK_USAGE_PATH "Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting."
	LOG_AGGREGATOR_MAX_DISK_USAGE_PATH string `json:"LOG_AGGREGATOR_MAX_DISK_USAGE_PATH,omitempty"`
	// LOG_AGGREGATOR_PASSWORD "Password or authentication token for external log aggregator (if required; HTTP/s only)."
	LOG_AGGREGATOR_PASSWORD string `json:"LOG_AGGREGATOR_PASSWORD,omitempty"`
	// LOG_AGGREGATOR_PORT "Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator)."
	LOG_AGGREGATOR_PORT int64 `json:"LOG_AGGREGATOR_PORT,omitempty"`
	// LOG_AGGREGATOR_PROTOCOL "Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname."
	LOG_AGGREGATOR_PROTOCOL string `json:"LOG_AGGREGATOR_PROTOCOL,omitempty"`
	// LOG_AGGREGATOR_RSYSLOGD_DEBUG "Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation."
	LOG_AGGREGATOR_RSYSLOGD_DEBUG bool `json:"LOG_AGGREGATOR_RSYSLOGD_DEBUG"`
	// LOG_AGGREGATOR_TCP_TIMEOUT "Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols."
	LOG_AGGREGATOR_TCP_TIMEOUT int64 `json:"LOG_AGGREGATOR_TCP_TIMEOUT,omitempty"`
	// LOG_AGGREGATOR_TOWER_UUID "Useful to uniquely identify instances."
	LOG_AGGREGATOR_TOWER_UUID string `json:"LOG_AGGREGATOR_TOWER_UUID,omitempty"`
	// LOG_AGGREGATOR_TYPE "Format messages for the chosen log aggregator."
	LOG_AGGREGATOR_TYPE string `json:"LOG_AGGREGATOR_TYPE,omitempty"`
	// LOG_AGGREGATOR_USERNAME "Username for external log aggregator (if required; HTTP/s only)."
	LOG_AGGREGATOR_USERNAME string `json:"LOG_AGGREGATOR_USERNAME,omitempty"`
	// LOG_AGGREGATOR_VERIFY_CERT "Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is \"https\". If enabled, the log handler will verify certificate sent by external log aggregator before establishing connection."
	LOG_AGGREGATOR_VERIFY_CERT bool `json:"LOG_AGGREGATOR_VERIFY_CERT"`
}
