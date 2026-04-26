package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsMiscLoggingTerraformModel struct {
	API_400_ERROR_LOG_FORMAT                types.String `tfsdk:"api_400_error_log_format" json:"API_400_ERROR_LOG_FORMAT"`
	LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB types.Int64  `tfsdk:"log_aggregator_action_max_disk_usage_gb" json:"LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB"`
	LOG_AGGREGATOR_ACTION_QUEUE_SIZE        types.Int64  `tfsdk:"log_aggregator_action_queue_size" json:"LOG_AGGREGATOR_ACTION_QUEUE_SIZE"`
	LOG_AGGREGATOR_ENABLED                  types.Bool   `tfsdk:"log_aggregator_enabled" json:"LOG_AGGREGATOR_ENABLED"`
	LOG_AGGREGATOR_HOST                     types.String `tfsdk:"log_aggregator_host" json:"LOG_AGGREGATOR_HOST"`
	LOG_AGGREGATOR_INDIVIDUAL_FACTS         types.Bool   `tfsdk:"log_aggregator_individual_facts" json:"LOG_AGGREGATOR_INDIVIDUAL_FACTS"`
	LOG_AGGREGATOR_LEVEL                    types.String `tfsdk:"log_aggregator_level" json:"LOG_AGGREGATOR_LEVEL"`
	LOG_AGGREGATOR_LOGGERS                  types.List   `tfsdk:"log_aggregator_loggers" json:"LOG_AGGREGATOR_LOGGERS"`
	LOG_AGGREGATOR_MAX_DISK_USAGE_PATH      types.String `tfsdk:"log_aggregator_max_disk_usage_path" json:"LOG_AGGREGATOR_MAX_DISK_USAGE_PATH"`
	LOG_AGGREGATOR_PASSWORD                 types.String `tfsdk:"log_aggregator_password" json:"LOG_AGGREGATOR_PASSWORD"`
	LOG_AGGREGATOR_PORT                     types.Int64  `tfsdk:"log_aggregator_port" json:"LOG_AGGREGATOR_PORT"`
	LOG_AGGREGATOR_PROTOCOL                 types.String `tfsdk:"log_aggregator_protocol" json:"LOG_AGGREGATOR_PROTOCOL"`
	LOG_AGGREGATOR_RSYSLOGD_DEBUG           types.Bool   `tfsdk:"log_aggregator_rsyslogd_debug" json:"LOG_AGGREGATOR_RSYSLOGD_DEBUG"`
	LOG_AGGREGATOR_TCP_TIMEOUT              types.Int64  `tfsdk:"log_aggregator_tcp_timeout" json:"LOG_AGGREGATOR_TCP_TIMEOUT"`
	LOG_AGGREGATOR_TOWER_UUID               types.String `tfsdk:"log_aggregator_tower_uuid" json:"LOG_AGGREGATOR_TOWER_UUID"`
	LOG_AGGREGATOR_TYPE                     types.String `tfsdk:"log_aggregator_type" json:"LOG_AGGREGATOR_TYPE"`
	LOG_AGGREGATOR_USERNAME                 types.String `tfsdk:"log_aggregator_username" json:"LOG_AGGREGATOR_USERNAME"`
	LOG_AGGREGATOR_VERIFY_CERT              types.Bool   `tfsdk:"log_aggregator_verify_cert" json:"LOG_AGGREGATOR_VERIFY_CERT"`
}

func (o *settingsMiscLoggingTerraformModel) Clone() settingsMiscLoggingTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.API_400_ERROR_LOG_FORMAT, data["API_400_ERROR_LOG_FORMAT"], false))
	collect(helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB, data["LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB"]))
	collect(helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_ACTION_QUEUE_SIZE, data["LOG_AGGREGATOR_ACTION_QUEUE_SIZE"]))
	collect(helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_ENABLED, data["LOG_AGGREGATOR_ENABLED"]))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_HOST, data["LOG_AGGREGATOR_HOST"], false))
	collect(helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_INDIVIDUAL_FACTS, data["LOG_AGGREGATOR_INDIVIDUAL_FACTS"]))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_LEVEL, data["LOG_AGGREGATOR_LEVEL"], false))
	collect(helpers.AttrValueSetListString(&o.LOG_AGGREGATOR_LOGGERS, data["LOG_AGGREGATOR_LOGGERS"], false))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_MAX_DISK_USAGE_PATH, data["LOG_AGGREGATOR_MAX_DISK_USAGE_PATH"], false))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_PASSWORD, data["LOG_AGGREGATOR_PASSWORD"], false))
	collect(helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_PORT, data["LOG_AGGREGATOR_PORT"]))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_PROTOCOL, data["LOG_AGGREGATOR_PROTOCOL"], false))
	collect(helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_RSYSLOGD_DEBUG, data["LOG_AGGREGATOR_RSYSLOGD_DEBUG"]))
	collect(helpers.AttrValueSetInt64(&o.LOG_AGGREGATOR_TCP_TIMEOUT, data["LOG_AGGREGATOR_TCP_TIMEOUT"]))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_TOWER_UUID, data["LOG_AGGREGATOR_TOWER_UUID"], false))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_TYPE, data["LOG_AGGREGATOR_TYPE"], false))
	collect(helpers.AttrValueSetString(&o.LOG_AGGREGATOR_USERNAME, data["LOG_AGGREGATOR_USERNAME"], false))
	collect(helpers.AttrValueSetBool(&o.LOG_AGGREGATOR_VERIFY_CERT, data["LOG_AGGREGATOR_VERIFY_CERT"]))
	return diags, nil
}

type settingsMiscLoggingBodyRequestModel struct {
	API_400_ERROR_LOG_FORMAT                string   `json:"API_400_ERROR_LOG_FORMAT,omitempty"`
	LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB int64    `json:"LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB,omitempty"`
	LOG_AGGREGATOR_ACTION_QUEUE_SIZE        int64    `json:"LOG_AGGREGATOR_ACTION_QUEUE_SIZE,omitempty"`
	LOG_AGGREGATOR_ENABLED                  bool     `json:"LOG_AGGREGATOR_ENABLED"`
	LOG_AGGREGATOR_HOST                     string   `json:"LOG_AGGREGATOR_HOST,omitempty"`
	LOG_AGGREGATOR_INDIVIDUAL_FACTS         bool     `json:"LOG_AGGREGATOR_INDIVIDUAL_FACTS"`
	LOG_AGGREGATOR_LEVEL                    string   `json:"LOG_AGGREGATOR_LEVEL,omitempty"`
	LOG_AGGREGATOR_LOGGERS                  []string `json:"LOG_AGGREGATOR_LOGGERS,omitempty"`
	LOG_AGGREGATOR_MAX_DISK_USAGE_PATH      string   `json:"LOG_AGGREGATOR_MAX_DISK_USAGE_PATH,omitempty"`
	LOG_AGGREGATOR_PASSWORD                 string   `json:"LOG_AGGREGATOR_PASSWORD,omitempty"`
	LOG_AGGREGATOR_PORT                     int64    `json:"LOG_AGGREGATOR_PORT,omitempty"`
	LOG_AGGREGATOR_PROTOCOL                 string   `json:"LOG_AGGREGATOR_PROTOCOL,omitempty"`
	LOG_AGGREGATOR_RSYSLOGD_DEBUG           bool     `json:"LOG_AGGREGATOR_RSYSLOGD_DEBUG"`
	LOG_AGGREGATOR_TCP_TIMEOUT              int64    `json:"LOG_AGGREGATOR_TCP_TIMEOUT,omitempty"`
	LOG_AGGREGATOR_TOWER_UUID               string   `json:"LOG_AGGREGATOR_TOWER_UUID,omitempty"`
	LOG_AGGREGATOR_TYPE                     string   `json:"LOG_AGGREGATOR_TYPE,omitempty"`
	LOG_AGGREGATOR_USERNAME                 string   `json:"LOG_AGGREGATOR_USERNAME,omitempty"`
	LOG_AGGREGATOR_VERIFY_CERT              bool     `json:"LOG_AGGREGATOR_VERIFY_CERT"`
}
