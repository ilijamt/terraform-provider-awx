package awx_21_8_0

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

// Schema defines the schema for the data source.
func (o *settingsJobsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"ad_hoc_commands": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of modules allowed to be used by ad-hoc jobs.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.List{},
			},
			"allow_jinja_in_extra_vars": schema.StringAttribute{
				Description: "Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \"template\" or \"never\".",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"always", "never", "template"}...),
				},
			},
			"ansible_fact_cache_timeout": schema.Int64Attribute{
				Description: "Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"awx_ansible_callback_plugins": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.List{},
			},
			"awx_collections_enabled": schema.BoolAttribute{
				Description: "Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"awx_isolation_base_path": schema.StringAttribute{
				Description: "The directory in which the service will create new temporary directories for job execution and isolation (such as credential files).",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"awx_isolation_show_paths": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]]. ",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.List{},
			},
			"awx_mount_isolated_paths_on_k8s": schema.BoolAttribute{
				Description: "Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible. ",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"awx_roles_enabled": schema.BoolAttribute{
				Description: "Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"awx_show_playbook_links": schema.BoolAttribute{
				Description: "Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"awx_task_env": schema.StringAttribute{
				Description: "Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"default_inventory_update_timeout": schema.Int64Attribute{
				Description: "Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"default_job_idle_timeout": schema.Int64Attribute{
				Description: "If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"default_job_timeout": schema.Int64Attribute{
				Description: "Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"default_project_update_timeout": schema.Int64Attribute{
				Description: "Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"event_stdout_max_bytes_display": schema.Int64Attribute{
				Description: "Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"galaxy_ignore_certs": schema.BoolAttribute{
				Description: "If set to true, certificate validation will not be done when installing content from any Galaxy server.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"galaxy_task_env": schema.StringAttribute{
				Description: "Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"max_forks": schema.Int64Attribute{
				Description: "Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"max_websocket_event_rate": schema.Int64Attribute{
				Description: "Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"project_update_vvv": schema.BoolAttribute{
				Description: "Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"schedule_max_jobs": schema.Int64Attribute{
				Description: "Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"stdout_max_bytes_display": schema.Int64Attribute{
				Description: "Maximum Size of Standard Output in bytes to display before requiring the output be downloaded.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			// Write only elements
		},
	}
}

func (o *settingsJobsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsJobsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsJobsTerraformModel
	var err error
	var endpoint = o.endpoint

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
			fmt.Sprintf("Unable to read resource for SettingsJobs on %s", endpoint),
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
