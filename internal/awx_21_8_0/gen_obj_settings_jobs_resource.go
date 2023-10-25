package awx_21_8_0

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

func (o *settingsJobsResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_jobs"
}

func (o *settingsJobsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Request elements
			"ad_hoc_commands": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of modules allowed to be used by ad-hoc jobs.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"allow_jinja_in_extra_vars": schema.StringAttribute{
				Description: "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\".",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`template`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"always", "never", "template"}...),
				},
			},
			"ansible_fact_cache_timeout": schema.Int64Attribute{
				Description: "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"awx_ansible_callback_plugins": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"awx_collections_enabled": schema.BoolAttribute{
				Description: "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"awx_isolation_base_path": schema.StringAttribute{
				Description: "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files).",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`/tmp`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"awx_isolation_show_paths": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. ",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"awx_mount_isolated_paths_on_k8s": schema.BoolAttribute{
				Description: "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. ",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"awx_roles_enabled": schema.BoolAttribute{
				Description: "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"awx_show_playbook_links": schema.BoolAttribute{
				Description: "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"awx_task_env": schema.StringAttribute{
				Description: "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"default_inventory_update_timeout": schema.Int64Attribute{
				Description: "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"default_job_idle_timeout": schema.Int64Attribute{
				Description: "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"default_job_timeout": schema.Int64Attribute{
				Description: "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"default_project_update_timeout": schema.Int64Attribute{
				Description: "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"event_stdout_max_bytes_display": schema.Int64Attribute{
				Description: "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1024),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"galaxy_ignore_certs": schema.BoolAttribute{
				Description: "If set to true, certificate validation will not be done when installing content from any Galaxy server.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"galaxy_task_env": schema.StringAttribute{
				Description: "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"ANSIBLE_FORCE_COLOR":"false","GIT_SSH_COMMAND":"ssh -o StrictHostKeyChecking=no"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"max_forks": schema.Int64Attribute{
				Description: "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(200),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"max_websocket_event_rate": schema.Int64Attribute{
				Description: "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(30),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"project_update_vvv": schema.BoolAttribute{
				Description: "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"schedule_max_jobs": schema.Int64Attribute{
				Description: "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(10),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"stdout_max_bytes_display": schema.Int64Attribute{
				Description: "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1.048576e+06),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			// Write only elements
			// Data only elements
		},
	}
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
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsJobs/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for create", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsJobs resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsJobs on %s", endpoint),
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
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for read", endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsJobs from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsJobs on %s", endpoint),
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
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsJobs/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsJobs on %s for update", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsJobs resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsJobs on %s", endpoint),
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
}
