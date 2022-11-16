package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// settingsJobsTerraformModel maps the schema for SettingsJobs when using Data Source
type settingsJobsTerraformModel struct {
	// AD_HOC_COMMANDS "List of modules allowed to be used by ad-hoc jobs."
	AD_HOC_COMMANDS types.List `tfsdk:"ad_hoc_commands" json:"AD_HOC_COMMANDS"`
	// ALLOW_JINJA_IN_EXTRA_VARS "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\"."
	ALLOW_JINJA_IN_EXTRA_VARS types.String `tfsdk:"allow_jinja_in_extra_vars" json:"ALLOW_JINJA_IN_EXTRA_VARS"`
	// ANSIBLE_FACT_CACHE_TIMEOUT "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed."
	ANSIBLE_FACT_CACHE_TIMEOUT types.Int64 `tfsdk:"ansible_fact_cache_timeout" json:"ANSIBLE_FACT_CACHE_TIMEOUT"`
	// AWX_ANSIBLE_CALLBACK_PLUGINS "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line."
	AWX_ANSIBLE_CALLBACK_PLUGINS types.List `tfsdk:"awx_ansible_callback_plugins" json:"AWX_ANSIBLE_CALLBACK_PLUGINS"`
	// AWX_COLLECTIONS_ENABLED "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects."
	AWX_COLLECTIONS_ENABLED types.Bool `tfsdk:"awx_collections_enabled" json:"AWX_COLLECTIONS_ENABLED"`
	// AWX_ISOLATION_BASE_PATH "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files)."
	AWX_ISOLATION_BASE_PATH types.String `tfsdk:"awx_isolation_base_path" json:"AWX_ISOLATION_BASE_PATH"`
	// AWX_ISOLATION_SHOW_PATHS "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. "
	AWX_ISOLATION_SHOW_PATHS types.List `tfsdk:"awx_isolation_show_paths" json:"AWX_ISOLATION_SHOW_PATHS"`
	// AWX_MOUNT_ISOLATED_PATHS_ON_K8S "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. "
	AWX_MOUNT_ISOLATED_PATHS_ON_K8S types.Bool `tfsdk:"awx_mount_isolated_paths_on_k8s" json:"AWX_MOUNT_ISOLATED_PATHS_ON_K8S"`
	// AWX_ROLES_ENABLED "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects."
	AWX_ROLES_ENABLED types.Bool `tfsdk:"awx_roles_enabled" json:"AWX_ROLES_ENABLED"`
	// AWX_SHOW_PLAYBOOK_LINKS "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself."
	AWX_SHOW_PLAYBOOK_LINKS types.Bool `tfsdk:"awx_show_playbook_links" json:"AWX_SHOW_PLAYBOOK_LINKS"`
	// AWX_TASK_ENV "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending."
	AWX_TASK_ENV types.String `tfsdk:"awx_task_env" json:"AWX_TASK_ENV"`
	// DEFAULT_INVENTORY_UPDATE_TIMEOUT "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this."
	DEFAULT_INVENTORY_UPDATE_TIMEOUT types.Int64 `tfsdk:"default_inventory_update_timeout" json:"DEFAULT_INVENTORY_UPDATE_TIMEOUT"`
	// DEFAULT_JOB_IDLE_TIMEOUT "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed."
	DEFAULT_JOB_IDLE_TIMEOUT types.Int64 `tfsdk:"default_job_idle_timeout" json:"DEFAULT_JOB_IDLE_TIMEOUT"`
	// DEFAULT_JOB_TIMEOUT "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this."
	DEFAULT_JOB_TIMEOUT types.Int64 `tfsdk:"default_job_timeout" json:"DEFAULT_JOB_TIMEOUT"`
	// DEFAULT_PROJECT_UPDATE_TIMEOUT "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this."
	DEFAULT_PROJECT_UPDATE_TIMEOUT types.Int64 `tfsdk:"default_project_update_timeout" json:"DEFAULT_PROJECT_UPDATE_TIMEOUT"`
	// EVENT_STDOUT_MAX_BYTES_DISPLAY "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated."
	EVENT_STDOUT_MAX_BYTES_DISPLAY types.Int64 `tfsdk:"event_stdout_max_bytes_display" json:"EVENT_STDOUT_MAX_BYTES_DISPLAY"`
	// GALAXY_IGNORE_CERTS "If set to true, certificate validation will not be done when installing content from any Galaxy server."
	GALAXY_IGNORE_CERTS types.Bool `tfsdk:"galaxy_ignore_certs" json:"GALAXY_IGNORE_CERTS"`
	// GALAXY_TASK_ENV "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git."
	GALAXY_TASK_ENV types.String `tfsdk:"galaxy_task_env" json:"GALAXY_TASK_ENV"`
	// MAX_FORKS "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied."
	MAX_FORKS types.Int64 `tfsdk:"max_forks" json:"MAX_FORKS"`
	// MAX_WEBSOCKET_EVENT_RATE "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit."
	MAX_WEBSOCKET_EVENT_RATE types.Int64 `tfsdk:"max_websocket_event_rate" json:"MAX_WEBSOCKET_EVENT_RATE"`
	// PROJECT_UPDATE_VVV "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates."
	PROJECT_UPDATE_VVV types.Bool `tfsdk:"project_update_vvv" json:"PROJECT_UPDATE_VVV"`
	// SCHEDULE_MAX_JOBS "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created."
	SCHEDULE_MAX_JOBS types.Int64 `tfsdk:"schedule_max_jobs" json:"SCHEDULE_MAX_JOBS"`
	// STDOUT_MAX_BYTES_DISPLAY "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded."
	STDOUT_MAX_BYTES_DISPLAY types.Int64 `tfsdk:"stdout_max_bytes_display" json:"STDOUT_MAX_BYTES_DISPLAY"`
}

// Clone the object
func (o settingsJobsTerraformModel) Clone() settingsJobsTerraformModel {
	return settingsJobsTerraformModel{
		AD_HOC_COMMANDS:                  o.AD_HOC_COMMANDS,
		ALLOW_JINJA_IN_EXTRA_VARS:        o.ALLOW_JINJA_IN_EXTRA_VARS,
		ANSIBLE_FACT_CACHE_TIMEOUT:       o.ANSIBLE_FACT_CACHE_TIMEOUT,
		AWX_ANSIBLE_CALLBACK_PLUGINS:     o.AWX_ANSIBLE_CALLBACK_PLUGINS,
		AWX_COLLECTIONS_ENABLED:          o.AWX_COLLECTIONS_ENABLED,
		AWX_ISOLATION_BASE_PATH:          o.AWX_ISOLATION_BASE_PATH,
		AWX_ISOLATION_SHOW_PATHS:         o.AWX_ISOLATION_SHOW_PATHS,
		AWX_MOUNT_ISOLATED_PATHS_ON_K8S:  o.AWX_MOUNT_ISOLATED_PATHS_ON_K8S,
		AWX_ROLES_ENABLED:                o.AWX_ROLES_ENABLED,
		AWX_SHOW_PLAYBOOK_LINKS:          o.AWX_SHOW_PLAYBOOK_LINKS,
		AWX_TASK_ENV:                     o.AWX_TASK_ENV,
		DEFAULT_INVENTORY_UPDATE_TIMEOUT: o.DEFAULT_INVENTORY_UPDATE_TIMEOUT,
		DEFAULT_JOB_IDLE_TIMEOUT:         o.DEFAULT_JOB_IDLE_TIMEOUT,
		DEFAULT_JOB_TIMEOUT:              o.DEFAULT_JOB_TIMEOUT,
		DEFAULT_PROJECT_UPDATE_TIMEOUT:   o.DEFAULT_PROJECT_UPDATE_TIMEOUT,
		EVENT_STDOUT_MAX_BYTES_DISPLAY:   o.EVENT_STDOUT_MAX_BYTES_DISPLAY,
		GALAXY_IGNORE_CERTS:              o.GALAXY_IGNORE_CERTS,
		GALAXY_TASK_ENV:                  o.GALAXY_TASK_ENV,
		MAX_FORKS:                        o.MAX_FORKS,
		MAX_WEBSOCKET_EVENT_RATE:         o.MAX_WEBSOCKET_EVENT_RATE,
		PROJECT_UPDATE_VVV:               o.PROJECT_UPDATE_VVV,
		SCHEDULE_MAX_JOBS:                o.SCHEDULE_MAX_JOBS,
		STDOUT_MAX_BYTES_DISPLAY:         o.STDOUT_MAX_BYTES_DISPLAY,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsJobs
func (o settingsJobsTerraformModel) BodyRequest() (req settingsJobsBodyRequestModel) {
	req.AD_HOC_COMMANDS = []string{}
	for _, val := range o.AD_HOC_COMMANDS.Elements() {
		if _, ok := val.(types.String); ok {
			req.AD_HOC_COMMANDS = append(req.AD_HOC_COMMANDS, val.(types.String).ValueString())
		} else {
			req.AD_HOC_COMMANDS = append(req.AD_HOC_COMMANDS, val.String())
		}
	}
	req.ALLOW_JINJA_IN_EXTRA_VARS = o.ALLOW_JINJA_IN_EXTRA_VARS.ValueString()
	req.ANSIBLE_FACT_CACHE_TIMEOUT = o.ANSIBLE_FACT_CACHE_TIMEOUT.ValueInt64()
	req.AWX_ANSIBLE_CALLBACK_PLUGINS = []string{}
	for _, val := range o.AWX_ANSIBLE_CALLBACK_PLUGINS.Elements() {
		if _, ok := val.(types.String); ok {
			req.AWX_ANSIBLE_CALLBACK_PLUGINS = append(req.AWX_ANSIBLE_CALLBACK_PLUGINS, val.(types.String).ValueString())
		} else {
			req.AWX_ANSIBLE_CALLBACK_PLUGINS = append(req.AWX_ANSIBLE_CALLBACK_PLUGINS, val.String())
		}
	}
	req.AWX_COLLECTIONS_ENABLED = o.AWX_COLLECTIONS_ENABLED.ValueBool()
	req.AWX_ISOLATION_BASE_PATH = o.AWX_ISOLATION_BASE_PATH.ValueString()
	req.AWX_ISOLATION_SHOW_PATHS = []string{}
	for _, val := range o.AWX_ISOLATION_SHOW_PATHS.Elements() {
		if _, ok := val.(types.String); ok {
			req.AWX_ISOLATION_SHOW_PATHS = append(req.AWX_ISOLATION_SHOW_PATHS, val.(types.String).ValueString())
		} else {
			req.AWX_ISOLATION_SHOW_PATHS = append(req.AWX_ISOLATION_SHOW_PATHS, val.String())
		}
	}
	req.AWX_MOUNT_ISOLATED_PATHS_ON_K8S = o.AWX_MOUNT_ISOLATED_PATHS_ON_K8S.ValueBool()
	req.AWX_ROLES_ENABLED = o.AWX_ROLES_ENABLED.ValueBool()
	req.AWX_SHOW_PLAYBOOK_LINKS = o.AWX_SHOW_PLAYBOOK_LINKS.ValueBool()
	req.AWX_TASK_ENV = json.RawMessage(o.AWX_TASK_ENV.ValueString())
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
	return
}

func (o *settingsJobsTerraformModel) setAdHocCommands(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AD_HOC_COMMANDS, data, false)
}

func (o *settingsJobsTerraformModel) setAllowJinjaInExtraVars(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ALLOW_JINJA_IN_EXTRA_VARS, data, false)
}

func (o *settingsJobsTerraformModel) setAnsibleFactCacheTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ANSIBLE_FACT_CACHE_TIMEOUT, data)
}

func (o *settingsJobsTerraformModel) setAwxAnsibleCallbackPlugins(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AWX_ANSIBLE_CALLBACK_PLUGINS, data, false)
}

func (o *settingsJobsTerraformModel) setAwxCollectionsEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AWX_COLLECTIONS_ENABLED, data)
}

func (o *settingsJobsTerraformModel) setAwxIsolationBasePath(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AWX_ISOLATION_BASE_PATH, data, false)
}

func (o *settingsJobsTerraformModel) setAwxIsolationShowPaths(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AWX_ISOLATION_SHOW_PATHS, data, false)
}

func (o *settingsJobsTerraformModel) setAwxMountIsolatedPathsOnK8S(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AWX_MOUNT_ISOLATED_PATHS_ON_K8S, data)
}

func (o *settingsJobsTerraformModel) setAwxRolesEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AWX_ROLES_ENABLED, data)
}

func (o *settingsJobsTerraformModel) setAwxShowPlaybookLinks(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AWX_SHOW_PLAYBOOK_LINKS, data)
}

func (o *settingsJobsTerraformModel) setAwxTaskEnv(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AWX_TASK_ENV, data, false)
}

func (o *settingsJobsTerraformModel) setDefaultInventoryUpdateTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.DEFAULT_INVENTORY_UPDATE_TIMEOUT, data)
}

func (o *settingsJobsTerraformModel) setDefaultJobIdleTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.DEFAULT_JOB_IDLE_TIMEOUT, data)
}

func (o *settingsJobsTerraformModel) setDefaultJobTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.DEFAULT_JOB_TIMEOUT, data)
}

func (o *settingsJobsTerraformModel) setDefaultProjectUpdateTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.DEFAULT_PROJECT_UPDATE_TIMEOUT, data)
}

func (o *settingsJobsTerraformModel) setEventStdoutMaxBytesDisplay(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.EVENT_STDOUT_MAX_BYTES_DISPLAY, data)
}

func (o *settingsJobsTerraformModel) setGalaxyIgnoreCerts(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.GALAXY_IGNORE_CERTS, data)
}

func (o *settingsJobsTerraformModel) setGalaxyTaskEnv(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.GALAXY_TASK_ENV, data, false)
}

func (o *settingsJobsTerraformModel) setMaxForks(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.MAX_FORKS, data)
}

func (o *settingsJobsTerraformModel) setMaxWebsocketEventRate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.MAX_WEBSOCKET_EVENT_RATE, data)
}

func (o *settingsJobsTerraformModel) setProjectUpdateVvv(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.PROJECT_UPDATE_VVV, data)
}

func (o *settingsJobsTerraformModel) setScheduleMaxJobs(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SCHEDULE_MAX_JOBS, data)
}

func (o *settingsJobsTerraformModel) setStdoutMaxBytesDisplay(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.STDOUT_MAX_BYTES_DISPLAY, data)
}

func (o *settingsJobsTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAdHocCommands(data["AD_HOC_COMMANDS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAllowJinjaInExtraVars(data["ALLOW_JINJA_IN_EXTRA_VARS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAnsibleFactCacheTimeout(data["ANSIBLE_FACT_CACHE_TIMEOUT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxAnsibleCallbackPlugins(data["AWX_ANSIBLE_CALLBACK_PLUGINS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxCollectionsEnabled(data["AWX_COLLECTIONS_ENABLED"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxIsolationBasePath(data["AWX_ISOLATION_BASE_PATH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxIsolationShowPaths(data["AWX_ISOLATION_SHOW_PATHS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxMountIsolatedPathsOnK8S(data["AWX_MOUNT_ISOLATED_PATHS_ON_K8S"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxRolesEnabled(data["AWX_ROLES_ENABLED"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxShowPlaybookLinks(data["AWX_SHOW_PLAYBOOK_LINKS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAwxTaskEnv(data["AWX_TASK_ENV"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultInventoryUpdateTimeout(data["DEFAULT_INVENTORY_UPDATE_TIMEOUT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultJobIdleTimeout(data["DEFAULT_JOB_IDLE_TIMEOUT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultJobTimeout(data["DEFAULT_JOB_TIMEOUT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultProjectUpdateTimeout(data["DEFAULT_PROJECT_UPDATE_TIMEOUT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setEventStdoutMaxBytesDisplay(data["EVENT_STDOUT_MAX_BYTES_DISPLAY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setGalaxyIgnoreCerts(data["GALAXY_IGNORE_CERTS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setGalaxyTaskEnv(data["GALAXY_TASK_ENV"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMaxForks(data["MAX_FORKS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMaxWebsocketEventRate(data["MAX_WEBSOCKET_EVENT_RATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setProjectUpdateVvv(data["PROJECT_UPDATE_VVV"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScheduleMaxJobs(data["SCHEDULE_MAX_JOBS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setStdoutMaxBytesDisplay(data["STDOUT_MAX_BYTES_DISPLAY"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsJobsBodyRequestModel maps the schema for SettingsJobs for creating and updating the data
type settingsJobsBodyRequestModel struct {
	// AD_HOC_COMMANDS "List of modules allowed to be used by ad-hoc jobs."
	AD_HOC_COMMANDS []string `json:"AD_HOC_COMMANDS,omitempty"`
	// ALLOW_JINJA_IN_EXTRA_VARS "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\"."
	ALLOW_JINJA_IN_EXTRA_VARS string `json:"ALLOW_JINJA_IN_EXTRA_VARS"`
	// ANSIBLE_FACT_CACHE_TIMEOUT "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed."
	ANSIBLE_FACT_CACHE_TIMEOUT int64 `json:"ANSIBLE_FACT_CACHE_TIMEOUT,omitempty"`
	// AWX_ANSIBLE_CALLBACK_PLUGINS "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line."
	AWX_ANSIBLE_CALLBACK_PLUGINS []string `json:"AWX_ANSIBLE_CALLBACK_PLUGINS,omitempty"`
	// AWX_COLLECTIONS_ENABLED "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects."
	AWX_COLLECTIONS_ENABLED bool `json:"AWX_COLLECTIONS_ENABLED"`
	// AWX_ISOLATION_BASE_PATH "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files)."
	AWX_ISOLATION_BASE_PATH string `json:"AWX_ISOLATION_BASE_PATH"`
	// AWX_ISOLATION_SHOW_PATHS "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. "
	AWX_ISOLATION_SHOW_PATHS []string `json:"AWX_ISOLATION_SHOW_PATHS,omitempty"`
	// AWX_MOUNT_ISOLATED_PATHS_ON_K8S "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. "
	AWX_MOUNT_ISOLATED_PATHS_ON_K8S bool `json:"AWX_MOUNT_ISOLATED_PATHS_ON_K8S"`
	// AWX_ROLES_ENABLED "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects."
	AWX_ROLES_ENABLED bool `json:"AWX_ROLES_ENABLED"`
	// AWX_SHOW_PLAYBOOK_LINKS "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself."
	AWX_SHOW_PLAYBOOK_LINKS bool `json:"AWX_SHOW_PLAYBOOK_LINKS"`
	// AWX_TASK_ENV "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending."
	AWX_TASK_ENV json.RawMessage `json:"AWX_TASK_ENV,omitempty"`
	// DEFAULT_INVENTORY_UPDATE_TIMEOUT "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this."
	DEFAULT_INVENTORY_UPDATE_TIMEOUT int64 `json:"DEFAULT_INVENTORY_UPDATE_TIMEOUT,omitempty"`
	// DEFAULT_JOB_IDLE_TIMEOUT "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed."
	DEFAULT_JOB_IDLE_TIMEOUT int64 `json:"DEFAULT_JOB_IDLE_TIMEOUT,omitempty"`
	// DEFAULT_JOB_TIMEOUT "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this."
	DEFAULT_JOB_TIMEOUT int64 `json:"DEFAULT_JOB_TIMEOUT,omitempty"`
	// DEFAULT_PROJECT_UPDATE_TIMEOUT "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this."
	DEFAULT_PROJECT_UPDATE_TIMEOUT int64 `json:"DEFAULT_PROJECT_UPDATE_TIMEOUT,omitempty"`
	// EVENT_STDOUT_MAX_BYTES_DISPLAY "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated."
	EVENT_STDOUT_MAX_BYTES_DISPLAY int64 `json:"EVENT_STDOUT_MAX_BYTES_DISPLAY"`
	// GALAXY_IGNORE_CERTS "If set to true, certificate validation will not be done when installing content from any Galaxy server."
	GALAXY_IGNORE_CERTS bool `json:"GALAXY_IGNORE_CERTS"`
	// GALAXY_TASK_ENV "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git."
	GALAXY_TASK_ENV json.RawMessage `json:"GALAXY_TASK_ENV"`
	// MAX_FORKS "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied."
	MAX_FORKS int64 `json:"MAX_FORKS,omitempty"`
	// MAX_WEBSOCKET_EVENT_RATE "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit."
	MAX_WEBSOCKET_EVENT_RATE int64 `json:"MAX_WEBSOCKET_EVENT_RATE,omitempty"`
	// PROJECT_UPDATE_VVV "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates."
	PROJECT_UPDATE_VVV bool `json:"PROJECT_UPDATE_VVV"`
	// SCHEDULE_MAX_JOBS "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created."
	SCHEDULE_MAX_JOBS int64 `json:"SCHEDULE_MAX_JOBS"`
	// STDOUT_MAX_BYTES_DISPLAY "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded."
	STDOUT_MAX_BYTES_DISPLAY int64 `json:"STDOUT_MAX_BYTES_DISPLAY"`
}

var (
	_ datasource.DataSource              = &settingsJobsDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsJobsDataSource{}
)

// NewSettingsJobsDataSource is a helper function to instantiate the SettingsJobs data source.
func NewSettingsJobsDataSource() datasource.DataSource {
	return &settingsJobsDataSource{}
}

// settingsJobsDataSource is the data source implementation.
type settingsJobsDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsJobsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/jobs/"
}

// Metadata returns the data source type name.
func (o *settingsJobsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_jobs"
}

// GetSchema defines the schema for the data source.
func (o *settingsJobsDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsJobs",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"ad_hoc_commands": {
					Description: "List of modules allowed to be used by ad-hoc jobs.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"allow_jinja_in_extra_vars": {
					Description: "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\".",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"always", "never", "template"}...),
					},
				},
				"ansible_fact_cache_timeout": {
					Description: "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_ansible_callback_plugins": {
					Description: "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_collections_enabled": {
					Description: "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_isolation_base_path": {
					Description: "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files).",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_isolation_show_paths": {
					Description: "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. ",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_mount_isolated_paths_on_k8s": {
					Description: "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. ",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_roles_enabled": {
					Description: "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_show_playbook_links": {
					Description: "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"awx_task_env": {
					Description: "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_inventory_update_timeout": {
					Description: "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_job_idle_timeout": {
					Description: "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_job_timeout": {
					Description: "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_project_update_timeout": {
					Description: "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"event_stdout_max_bytes_display": {
					Description: "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"galaxy_ignore_certs": {
					Description: "If set to true, certificate validation will not be done when installing content from any Galaxy server.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"galaxy_task_env": {
					Description: "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"max_forks": {
					Description: "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"max_websocket_event_rate": {
					Description: "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"project_update_vvv": {
					Description: "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"schedule_max_jobs": {
					Description: "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"stdout_max_bytes_display": {
					Description: "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsJobsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsJobsTerraformModel
	var err error
	var endpoint string
	endpoint = o.endpoint

	// Creates a new request for SettingsJobs
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsJobs
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsJobs on %s", o.endpoint),
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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsJobsResource{}
	_ resource.ResourceWithConfigure = &settingsJobsResource{}
)

// NewSettingsJobsResource is a helper function to simplify the provider implementation.
func NewSettingsJobsResource() resource.Resource {
	return &settingsJobsResource{}
}

type settingsJobsResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsJobsResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/jobs/"
}

func (o settingsJobsResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_jobs"
}

func (o settingsJobsResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsJobs",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"ad_hoc_commands": {
					Description: "List of modules allowed to be used by ad-hoc jobs.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"allow_jinja_in_extra_vars": {
					Description: "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\".",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`template`)),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"always", "never", "template"}...),
					},
				},
				"ansible_fact_cache_timeout": {
					Description: "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_ansible_callback_plugins": {
					Description: "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_collections_enabled": {
					Description: "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_isolation_base_path": {
					Description: "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files).",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`/tmp`)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_isolation_show_paths": {
					Description: "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. ",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_mount_isolated_paths_on_k8s": {
					Description: "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. ",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_roles_enabled": {
					Description: "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_show_playbook_links": {
					Description: "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"awx_task_env": {
					Description: "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"default_inventory_update_timeout": {
					Description: "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"default_job_idle_timeout": {
					Description: "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"default_job_timeout": {
					Description: "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"default_project_update_timeout": {
					Description: "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"event_stdout_max_bytes_display": {
					Description: "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated.",
					Type:        types.Int64Type,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(1024)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"galaxy_ignore_certs": {
					Description: "If set to true, certificate validation will not be done when installing content from any Galaxy server.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"galaxy_task_env": {
					Description: "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git.",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"ANSIBLE_FORCE_COLOR":"false","GIT_SSH_COMMAND":"ssh -o StrictHostKeyChecking=no"}`)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"max_forks": {
					Description: "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(200)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"max_websocket_event_rate": {
					Description: "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(30)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"project_update_vvv": {
					Description:   "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates.",
					Type:          types.BoolType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"schedule_max_jobs": {
					Description: "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created.",
					Type:        types.Int64Type,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(10)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"stdout_max_bytes_display": {
					Description: "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded.",
					Type:        types.Int64Type,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(1.048576e+06)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
			},
		}), nil
}

func (o *settingsJobsResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsJobsTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsJobs
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsJobs resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsJobs on %s", o.endpoint),
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

func (o *settingsJobsResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsJobsTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsJobs
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsJobs from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsJobs on %s", o.endpoint),
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

func (o *settingsJobsResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsJobsTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsJobs
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsJobs resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsJobs on %s", o.endpoint),
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

func (o *settingsJobsResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	return
}
