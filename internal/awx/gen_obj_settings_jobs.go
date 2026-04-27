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

type settingsJobsResource = framework.GenericResource[settingsJobsTerraformModel, settingsJobsBodyRequestModel, *settingsJobsTerraformModel]

// NewSettingsJobsResource is a helper function to simplify the provider implementation.
func NewSettingsJobsResource() resource.Resource {
	return &settingsJobsResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_jobs", Endpoint: "/api/v2/settings/jobs/"}},
		Cfg: framework.ResourceCfg[settingsJobsTerraformModel, settingsJobsBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"ad_hoc_commands": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of modules allowed to be used by ad-hoc jobs.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"allow_jinja_in_extra_vars": schema.StringAttribute{
						Description: "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\".",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`template`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"always",
								"never",
								"template",
							),
						},
					},
					"ansible_fact_cache_timeout": schema.Int64Attribute{
						Description: "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"awx_ansible_callback_plugins": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_collections_enabled": schema.BoolAttribute{
						Description: "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_isolation_base_path": schema.StringAttribute{
						Description: "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files).",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`/tmp`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_isolation_show_paths": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. ",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_mount_isolated_paths_on_k8s": schema.BoolAttribute{
						Description: "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. ",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_roles_enabled": schema.BoolAttribute{
						Description: "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_runner_keepalive_seconds": schema.Int64Attribute{
						Description: "Only applies to jobs running in a Container Group. If not 0, send a message every so-many seconds to keep connection open.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"awx_show_playbook_links": schema.BoolAttribute{
						Description: "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"awx_task_env": schema.StringAttribute{
						Description: "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"default_container_run_options": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of options to pass to podman run example: ['--network', 'slirp4netns:enable_ipv6=true', '--log-level', 'debug']",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"default_inventory_update_timeout": schema.Int64Attribute{
						Description: "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"default_job_idle_timeout": schema.Int64Attribute{
						Description: "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"default_job_timeout": schema.Int64Attribute{
						Description: "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"default_project_update_timeout": schema.Int64Attribute{
						Description: "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"event_stdout_max_bytes_display": schema.Int64Attribute{
						Description: "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1024),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"galaxy_ignore_certs": schema.BoolAttribute{
						Description: "If set to true, certificate validation will not be done when installing content from any Galaxy server.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"galaxy_task_env": schema.StringAttribute{
						Description: "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"ANSIBLE_FORCE_COLOR":"false","GIT_SSH_COMMAND":"ssh -o StrictHostKeyChecking=no"}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"max_forks": schema.Int64Attribute{
						Description: "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(200),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"max_websocket_event_rate": schema.Int64Attribute{
						Description: "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(30),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"project_update_vvv": schema.BoolAttribute{
						Description: "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"schedule_max_jobs": schema.Int64Attribute{
						Description: "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(10),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"stdout_max_bytes_display": schema.Int64Attribute{
						Description: "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1.048576e+06),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsJobs",
		},
	}
}

type settingsJobsDataSource = framework.GenericDataSource[settingsJobsTerraformModel, *settingsJobsTerraformModel]

// NewSettingsJobsDataSource is a helper function to instantiate the SettingsJobs data source.
func NewSettingsJobsDataSource() datasource.DataSource {
	return &settingsJobsDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_jobs", Endpoint: "/api/v2/settings/jobs/"}},
		Cfg: framework.DataSourceCfg[settingsJobsTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"ad_hoc_commands": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of modules allowed to be used by ad-hoc jobs.",
						Computed:    true,
					},
					"allow_jinja_in_extra_vars": dschema.StringAttribute{
						Description: "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\".",
						Computed:    true,
					},
					"ansible_fact_cache_timeout": dschema.Int64Attribute{
						Description: "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed.",
						Computed:    true,
					},
					"awx_ansible_callback_plugins": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line.",
						Computed:    true,
					},
					"awx_collections_enabled": dschema.BoolAttribute{
						Description: "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects.",
						Computed:    true,
					},
					"awx_isolation_base_path": dschema.StringAttribute{
						Description: "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files).",
						Computed:    true,
					},
					"awx_isolation_show_paths": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. ",
						Computed:    true,
					},
					"awx_mount_isolated_paths_on_k8s": dschema.BoolAttribute{
						Description: "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. ",
						Computed:    true,
					},
					"awx_roles_enabled": dschema.BoolAttribute{
						Description: "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects.",
						Computed:    true,
					},
					"awx_runner_keepalive_seconds": dschema.Int64Attribute{
						Description: "Only applies to jobs running in a Container Group. If not 0, send a message every so-many seconds to keep connection open.",
						Computed:    true,
					},
					"awx_show_playbook_links": dschema.BoolAttribute{
						Description: "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself.",
						Computed:    true,
					},
					"awx_task_env": dschema.StringAttribute{
						Description: "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending.",
						Computed:    true,
					},
					"default_container_run_options": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of options to pass to podman run example: ['--network', 'slirp4netns:enable_ipv6=true', '--log-level', 'debug']",
						Computed:    true,
					},
					"default_inventory_update_timeout": dschema.Int64Attribute{
						Description: "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this.",
						Computed:    true,
					},
					"default_job_idle_timeout": dschema.Int64Attribute{
						Description: "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed.",
						Computed:    true,
					},
					"default_job_timeout": dschema.Int64Attribute{
						Description: "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this.",
						Computed:    true,
					},
					"default_project_update_timeout": dschema.Int64Attribute{
						Description: "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this.",
						Computed:    true,
					},
					"event_stdout_max_bytes_display": dschema.Int64Attribute{
						Description: "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated.",
						Computed:    true,
					},
					"galaxy_ignore_certs": dschema.BoolAttribute{
						Description: "If set to true, certificate validation will not be done when installing content from any Galaxy server.",
						Computed:    true,
					},
					"galaxy_task_env": dschema.StringAttribute{
						Description: "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git.",
						Computed:    true,
					},
					"max_forks": dschema.Int64Attribute{
						Description: "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied.",
						Computed:    true,
					},
					"max_websocket_event_rate": dschema.Int64Attribute{
						Description: "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit.",
						Computed:    true,
					},
					"project_update_vvv": dschema.BoolAttribute{
						Description: "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates.",
						Computed:    true,
					},
					"schedule_max_jobs": dschema.Int64Attribute{
						Description: "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created.",
						Computed:    true,
					},
					"stdout_max_bytes_display": dschema.Int64Attribute{
						Description: "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsJobs",
		},
	}
}
