package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsJobsTerraformModel struct {
	AD_HOC_COMMANDS                  types.List   `tfsdk:"ad_hoc_commands" json:"AD_HOC_COMMANDS"`
	ALLOW_JINJA_IN_EXTRA_VARS        types.String `tfsdk:"allow_jinja_in_extra_vars" json:"ALLOW_JINJA_IN_EXTRA_VARS"`
	ANSIBLE_FACT_CACHE_TIMEOUT       types.Int64  `tfsdk:"ansible_fact_cache_timeout" json:"ANSIBLE_FACT_CACHE_TIMEOUT"`
	AWX_ANSIBLE_CALLBACK_PLUGINS     types.List   `tfsdk:"awx_ansible_callback_plugins" json:"AWX_ANSIBLE_CALLBACK_PLUGINS"`
	AWX_COLLECTIONS_ENABLED          types.Bool   `tfsdk:"awx_collections_enabled" json:"AWX_COLLECTIONS_ENABLED"`
	AWX_ISOLATION_BASE_PATH          types.String `tfsdk:"awx_isolation_base_path" json:"AWX_ISOLATION_BASE_PATH"`
	AWX_ISOLATION_SHOW_PATHS         types.List   `tfsdk:"awx_isolation_show_paths" json:"AWX_ISOLATION_SHOW_PATHS"`
	AWX_MOUNT_ISOLATED_PATHS_ON_K8S  types.Bool   `tfsdk:"awx_mount_isolated_paths_on_k8s" json:"AWX_MOUNT_ISOLATED_PATHS_ON_K8S"`
	AWX_ROLES_ENABLED                types.Bool   `tfsdk:"awx_roles_enabled" json:"AWX_ROLES_ENABLED"`
	AWX_RUNNER_KEEPALIVE_SECONDS     types.Int64  `tfsdk:"awx_runner_keepalive_seconds" json:"AWX_RUNNER_KEEPALIVE_SECONDS"`
	AWX_SHOW_PLAYBOOK_LINKS          types.Bool   `tfsdk:"awx_show_playbook_links" json:"AWX_SHOW_PLAYBOOK_LINKS"`
	AWX_TASK_ENV                     types.String `tfsdk:"awx_task_env" json:"AWX_TASK_ENV"`
	DEFAULT_CONTAINER_RUN_OPTIONS    types.List   `tfsdk:"default_container_run_options" json:"DEFAULT_CONTAINER_RUN_OPTIONS"`
	DEFAULT_INVENTORY_UPDATE_TIMEOUT types.Int64  `tfsdk:"default_inventory_update_timeout" json:"DEFAULT_INVENTORY_UPDATE_TIMEOUT"`
	DEFAULT_JOB_IDLE_TIMEOUT         types.Int64  `tfsdk:"default_job_idle_timeout" json:"DEFAULT_JOB_IDLE_TIMEOUT"`
	DEFAULT_JOB_TIMEOUT              types.Int64  `tfsdk:"default_job_timeout" json:"DEFAULT_JOB_TIMEOUT"`
	DEFAULT_PROJECT_UPDATE_TIMEOUT   types.Int64  `tfsdk:"default_project_update_timeout" json:"DEFAULT_PROJECT_UPDATE_TIMEOUT"`
	EVENT_STDOUT_MAX_BYTES_DISPLAY   types.Int64  `tfsdk:"event_stdout_max_bytes_display" json:"EVENT_STDOUT_MAX_BYTES_DISPLAY"`
	GALAXY_IGNORE_CERTS              types.Bool   `tfsdk:"galaxy_ignore_certs" json:"GALAXY_IGNORE_CERTS"`
	GALAXY_TASK_ENV                  types.String `tfsdk:"galaxy_task_env" json:"GALAXY_TASK_ENV"`
	MAX_FORKS                        types.Int64  `tfsdk:"max_forks" json:"MAX_FORKS"`
	MAX_WEBSOCKET_EVENT_RATE         types.Int64  `tfsdk:"max_websocket_event_rate" json:"MAX_WEBSOCKET_EVENT_RATE"`
	PROJECT_UPDATE_VVV               types.Bool   `tfsdk:"project_update_vvv" json:"PROJECT_UPDATE_VVV"`
	SCHEDULE_MAX_JOBS                types.Int64  `tfsdk:"schedule_max_jobs" json:"SCHEDULE_MAX_JOBS"`
	STDOUT_MAX_BYTES_DISPLAY         types.Int64  `tfsdk:"stdout_max_bytes_display" json:"STDOUT_MAX_BYTES_DISPLAY"`
}

func (o *settingsJobsTerraformModel) Clone() settingsJobsTerraformModel {
	return *o
}

func (o *settingsJobsTerraformModel) BodyRequest() *settingsJobsBodyRequestModel {
	var req settingsJobsBodyRequestModel
	req.AD_HOC_COMMANDS = helpers.ListAsStringSlice(o.AD_HOC_COMMANDS, false)
	req.ALLOW_JINJA_IN_EXTRA_VARS = o.ALLOW_JINJA_IN_EXTRA_VARS.ValueString()
	req.ANSIBLE_FACT_CACHE_TIMEOUT = o.ANSIBLE_FACT_CACHE_TIMEOUT.ValueInt64()
	req.AWX_ANSIBLE_CALLBACK_PLUGINS = helpers.ListAsStringSlice(o.AWX_ANSIBLE_CALLBACK_PLUGINS, false)
	req.AWX_COLLECTIONS_ENABLED = o.AWX_COLLECTIONS_ENABLED.ValueBool()
	req.AWX_ISOLATION_BASE_PATH = o.AWX_ISOLATION_BASE_PATH.ValueString()
	req.AWX_ISOLATION_SHOW_PATHS = helpers.ListAsStringSlice(o.AWX_ISOLATION_SHOW_PATHS, false)
	req.AWX_MOUNT_ISOLATED_PATHS_ON_K8S = o.AWX_MOUNT_ISOLATED_PATHS_ON_K8S.ValueBool()
	req.AWX_ROLES_ENABLED = o.AWX_ROLES_ENABLED.ValueBool()
	req.AWX_RUNNER_KEEPALIVE_SECONDS = o.AWX_RUNNER_KEEPALIVE_SECONDS.ValueInt64()
	req.AWX_SHOW_PLAYBOOK_LINKS = o.AWX_SHOW_PLAYBOOK_LINKS.ValueBool()
	req.AWX_TASK_ENV = json.RawMessage(o.AWX_TASK_ENV.ValueString())
	req.DEFAULT_CONTAINER_RUN_OPTIONS = helpers.ListAsStringSlice(o.DEFAULT_CONTAINER_RUN_OPTIONS, false)
	req.DEFAULT_INVENTORY_UPDATE_TIMEOUT = o.DEFAULT_INVENTORY_UPDATE_TIMEOUT.ValueInt64()
	req.DEFAULT_JOB_IDLE_TIMEOUT = o.DEFAULT_JOB_IDLE_TIMEOUT.ValueInt64()
	req.DEFAULT_JOB_TIMEOUT = o.DEFAULT_JOB_TIMEOUT.ValueInt64()
	req.DEFAULT_PROJECT_UPDATE_TIMEOUT = o.DEFAULT_PROJECT_UPDATE_TIMEOUT.ValueInt64()
	req.EVENT_STDOUT_MAX_BYTES_DISPLAY = o.EVENT_STDOUT_MAX_BYTES_DISPLAY.ValueInt64()
	req.GALAXY_IGNORE_CERTS = o.GALAXY_IGNORE_CERTS.ValueBool()
	req.GALAXY_TASK_ENV = json.RawMessage(o.GALAXY_TASK_ENV.ValueString())
	req.MAX_FORKS = o.MAX_FORKS.ValueInt64()
	req.MAX_WEBSOCKET_EVENT_RATE = o.MAX_WEBSOCKET_EVENT_RATE.ValueInt64()
	req.PROJECT_UPDATE_VVV = o.PROJECT_UPDATE_VVV.ValueBool()
	req.SCHEDULE_MAX_JOBS = o.SCHEDULE_MAX_JOBS.ValueInt64()
	req.STDOUT_MAX_BYTES_DISPLAY = o.STDOUT_MAX_BYTES_DISPLAY.ValueInt64()
	return &req
}

func (o *settingsJobsTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetListString(&o.AD_HOC_COMMANDS, data["AD_HOC_COMMANDS"], false))
	collect(helpers.AttrValueSetString(&o.ALLOW_JINJA_IN_EXTRA_VARS, data["ALLOW_JINJA_IN_EXTRA_VARS"], false))
	collect(helpers.AttrValueSetInt64(&o.ANSIBLE_FACT_CACHE_TIMEOUT, data["ANSIBLE_FACT_CACHE_TIMEOUT"]))
	collect(helpers.AttrValueSetListString(&o.AWX_ANSIBLE_CALLBACK_PLUGINS, data["AWX_ANSIBLE_CALLBACK_PLUGINS"], false))
	collect(helpers.AttrValueSetBool(&o.AWX_COLLECTIONS_ENABLED, data["AWX_COLLECTIONS_ENABLED"]))
	collect(helpers.AttrValueSetString(&o.AWX_ISOLATION_BASE_PATH, data["AWX_ISOLATION_BASE_PATH"], false))
	collect(helpers.AttrValueSetListString(&o.AWX_ISOLATION_SHOW_PATHS, data["AWX_ISOLATION_SHOW_PATHS"], false))
	collect(helpers.AttrValueSetBool(&o.AWX_MOUNT_ISOLATED_PATHS_ON_K8S, data["AWX_MOUNT_ISOLATED_PATHS_ON_K8S"]))
	collect(helpers.AttrValueSetBool(&o.AWX_ROLES_ENABLED, data["AWX_ROLES_ENABLED"]))
	collect(helpers.AttrValueSetInt64(&o.AWX_RUNNER_KEEPALIVE_SECONDS, data["AWX_RUNNER_KEEPALIVE_SECONDS"]))
	collect(helpers.AttrValueSetBool(&o.AWX_SHOW_PLAYBOOK_LINKS, data["AWX_SHOW_PLAYBOOK_LINKS"]))
	collect(helpers.AttrValueSetJsonString(&o.AWX_TASK_ENV, data["AWX_TASK_ENV"], false))
	collect(helpers.AttrValueSetListString(&o.DEFAULT_CONTAINER_RUN_OPTIONS, data["DEFAULT_CONTAINER_RUN_OPTIONS"], false))
	collect(helpers.AttrValueSetInt64(&o.DEFAULT_INVENTORY_UPDATE_TIMEOUT, data["DEFAULT_INVENTORY_UPDATE_TIMEOUT"]))
	collect(helpers.AttrValueSetInt64(&o.DEFAULT_JOB_IDLE_TIMEOUT, data["DEFAULT_JOB_IDLE_TIMEOUT"]))
	collect(helpers.AttrValueSetInt64(&o.DEFAULT_JOB_TIMEOUT, data["DEFAULT_JOB_TIMEOUT"]))
	collect(helpers.AttrValueSetInt64(&o.DEFAULT_PROJECT_UPDATE_TIMEOUT, data["DEFAULT_PROJECT_UPDATE_TIMEOUT"]))
	collect(helpers.AttrValueSetInt64(&o.EVENT_STDOUT_MAX_BYTES_DISPLAY, data["EVENT_STDOUT_MAX_BYTES_DISPLAY"]))
	collect(helpers.AttrValueSetBool(&o.GALAXY_IGNORE_CERTS, data["GALAXY_IGNORE_CERTS"]))
	collect(helpers.AttrValueSetJsonString(&o.GALAXY_TASK_ENV, data["GALAXY_TASK_ENV"], false))
	collect(helpers.AttrValueSetInt64(&o.MAX_FORKS, data["MAX_FORKS"]))
	collect(helpers.AttrValueSetInt64(&o.MAX_WEBSOCKET_EVENT_RATE, data["MAX_WEBSOCKET_EVENT_RATE"]))
	collect(helpers.AttrValueSetBool(&o.PROJECT_UPDATE_VVV, data["PROJECT_UPDATE_VVV"]))
	collect(helpers.AttrValueSetInt64(&o.SCHEDULE_MAX_JOBS, data["SCHEDULE_MAX_JOBS"]))
	collect(helpers.AttrValueSetInt64(&o.STDOUT_MAX_BYTES_DISPLAY, data["STDOUT_MAX_BYTES_DISPLAY"]))
	return diags, nil
}

type settingsJobsBodyRequestModel struct {
	AD_HOC_COMMANDS                  []string        `json:"AD_HOC_COMMANDS,omitempty"`
	ALLOW_JINJA_IN_EXTRA_VARS        string          `json:"ALLOW_JINJA_IN_EXTRA_VARS,omitempty"`
	ANSIBLE_FACT_CACHE_TIMEOUT       int64           `json:"ANSIBLE_FACT_CACHE_TIMEOUT,omitempty"`
	AWX_ANSIBLE_CALLBACK_PLUGINS     []string        `json:"AWX_ANSIBLE_CALLBACK_PLUGINS,omitempty"`
	AWX_COLLECTIONS_ENABLED          bool            `json:"AWX_COLLECTIONS_ENABLED"`
	AWX_ISOLATION_BASE_PATH          string          `json:"AWX_ISOLATION_BASE_PATH,omitempty"`
	AWX_ISOLATION_SHOW_PATHS         []string        `json:"AWX_ISOLATION_SHOW_PATHS,omitempty"`
	AWX_MOUNT_ISOLATED_PATHS_ON_K8S  bool            `json:"AWX_MOUNT_ISOLATED_PATHS_ON_K8S"`
	AWX_ROLES_ENABLED                bool            `json:"AWX_ROLES_ENABLED"`
	AWX_RUNNER_KEEPALIVE_SECONDS     int64           `json:"AWX_RUNNER_KEEPALIVE_SECONDS,omitempty"`
	AWX_SHOW_PLAYBOOK_LINKS          bool            `json:"AWX_SHOW_PLAYBOOK_LINKS"`
	AWX_TASK_ENV                     json.RawMessage `json:"AWX_TASK_ENV,omitempty"`
	DEFAULT_CONTAINER_RUN_OPTIONS    []string        `json:"DEFAULT_CONTAINER_RUN_OPTIONS,omitempty"`
	DEFAULT_INVENTORY_UPDATE_TIMEOUT int64           `json:"DEFAULT_INVENTORY_UPDATE_TIMEOUT,omitempty"`
	DEFAULT_JOB_IDLE_TIMEOUT         int64           `json:"DEFAULT_JOB_IDLE_TIMEOUT,omitempty"`
	DEFAULT_JOB_TIMEOUT              int64           `json:"DEFAULT_JOB_TIMEOUT,omitempty"`
	DEFAULT_PROJECT_UPDATE_TIMEOUT   int64           `json:"DEFAULT_PROJECT_UPDATE_TIMEOUT,omitempty"`
	EVENT_STDOUT_MAX_BYTES_DISPLAY   int64           `json:"EVENT_STDOUT_MAX_BYTES_DISPLAY"`
	GALAXY_IGNORE_CERTS              bool            `json:"GALAXY_IGNORE_CERTS"`
	GALAXY_TASK_ENV                  json.RawMessage `json:"GALAXY_TASK_ENV,omitempty"`
	MAX_FORKS                        int64           `json:"MAX_FORKS,omitempty"`
	MAX_WEBSOCKET_EVENT_RATE         int64           `json:"MAX_WEBSOCKET_EVENT_RATE"`
	PROJECT_UPDATE_VVV               bool            `json:"PROJECT_UPDATE_VVV"`
	SCHEDULE_MAX_JOBS                int64           `json:"SCHEDULE_MAX_JOBS,omitempty"`
	STDOUT_MAX_BYTES_DISPLAY         int64           `json:"STDOUT_MAX_BYTES_DISPLAY"`
}
