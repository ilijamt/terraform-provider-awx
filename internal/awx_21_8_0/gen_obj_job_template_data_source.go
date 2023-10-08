package awx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &jobTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &jobTemplateDataSource{}
)

// NewJobTemplateDataSource is a helper function to instantiate the JobTemplate data source.
func NewJobTemplateDataSource() datasource.DataSource {
	return &jobTemplateDataSource{}
}

// jobTemplateDataSource is the data source implementation.
type jobTemplateDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *jobTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/"
}

// Metadata returns the data source type name.
func (o *jobTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_job_template"
}

// Schema defines the schema for the data source.
func (o *jobTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"allow_simultaneous": schema.BoolAttribute{
				Description: "Allow simultaneous",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_credential_on_launch": schema.BoolAttribute{
				Description: "Ask credential on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_diff_mode_on_launch": schema.BoolAttribute{
				Description: "Ask diff mode on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_execution_environment_on_launch": schema.BoolAttribute{
				Description: "Ask execution environment on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_forks_on_launch": schema.BoolAttribute{
				Description: "Ask forks on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_instance_groups_on_launch": schema.BoolAttribute{
				Description: "Ask instance groups on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_inventory_on_launch": schema.BoolAttribute{
				Description: "Ask inventory on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_job_slice_count_on_launch": schema.BoolAttribute{
				Description: "Ask job slice count on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_job_type_on_launch": schema.BoolAttribute{
				Description: "Ask job type on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_labels_on_launch": schema.BoolAttribute{
				Description: "Ask labels on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_limit_on_launch": schema.BoolAttribute{
				Description: "Ask limit on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_scm_branch_on_launch": schema.BoolAttribute{
				Description: "Ask scm branch on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_skip_tags_on_launch": schema.BoolAttribute{
				Description: "Ask skip tags on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_tags_on_launch": schema.BoolAttribute{
				Description: "Ask tags on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_timeout_on_launch": schema.BoolAttribute{
				Description: "Ask timeout on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_variables_on_launch": schema.BoolAttribute{
				Description: "Ask variables on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_verbosity_on_launch": schema.BoolAttribute{
				Description: "Ask verbosity on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"become_enabled": schema.BoolAttribute{
				Description: "Become enabled",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this job template.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"diff_mode": schema.BoolAttribute{
				Description: "If enabled, textual changes made to any templated files on the host are shown in the standard output",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"execution_environment": schema.Int64Attribute{
				Description: "The container image to be used for execution.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"extra_vars": schema.StringAttribute{
				Description: "Extra vars",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"force_handlers": schema.BoolAttribute{
				Description: "Force handlers",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"forks": schema.Int64Attribute{
				Description: "Forks",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"host_config_key": schema.StringAttribute{
				Description: "Host config key",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this job template.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("name"),
					),
				},
			},
			"inventory": schema.Int64Attribute{
				Description: "Inventory",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"job_slice_count": schema.Int64Attribute{
				Description: "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"job_tags": schema.StringAttribute{
				Description: "Job tags",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"job_type": schema.StringAttribute{
				Description: "Job type",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"run", "check"}...),
				},
			},
			"limit": schema.StringAttribute{
				Description: "Limit",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this job template.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("name"),
					),
				},
			},
			"organization": schema.Int64Attribute{
				Description: "The organization used to determine access to this template.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"playbook": schema.StringAttribute{
				Description: "Playbook",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"prevent_instance_group_fallback": schema.BoolAttribute{
				Description: "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"project": schema.Int64Attribute{
				Description: "Project",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"scm_branch": schema.StringAttribute{
				Description: "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"skip_tags": schema.StringAttribute{
				Description: "Skip tags",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"start_at_task": schema.StringAttribute{
				Description: "Start at task",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"survey_enabled": schema.BoolAttribute{
				Description: "Survey enabled",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"timeout": schema.Int64Attribute{
				Description: "The amount of time (in seconds) to run before the task is canceled.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			"use_fact_cache": schema.BoolAttribute{
				Description: "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"verbosity": schema.StringAttribute{
				Description: "Verbosity",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
				},
			},
			"webhook_credential": schema.Int64Attribute{
				Description: "Personal Access Token for posting back the status to the service API",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"webhook_service": schema.StringAttribute{
				Description: "Service that webhook requests will be accepted from",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"", "github", "gitlab"}...),
				},
			},
			// Write only elements
		},
	}
}

func (o *jobTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *jobTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state jobTemplateTerraformModel
	var err error
	var endpoint string
	var searchDefined bool

	// Only one group should evaluate to True, terraform should prevent from being able to set
	// the conflicting groups

	// Evaluate group 'by_id' based on the schema definition
	var groupByIdExists = func() bool {
		var groupByIdExists = true
		var paramsById = []any{o.endpoint}
		var attrID types.Int64
		req.Config.GetAttribute(ctx, path.Root("id"), &attrID)
		groupByIdExists = groupByIdExists && (!attrID.IsNull() && !attrID.IsUnknown())
		paramsById = append(paramsById, attrID.ValueInt64())
		if groupByIdExists {
			endpoint = p.Clean(fmt.Sprintf("%s/%d/", paramsById...))
		}
		return groupByIdExists
	}()
	searchDefined = searchDefined || groupByIdExists

	// Evaluate group 'by_name' based on the schema definition
	var groupByNameExists = func() bool {
		var groupByNameExists = true
		var paramsByName = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameExists = groupByNameExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByName = append(paramsByName, url.PathEscape(attrName.ValueString()))
		if groupByNameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s", paramsByName...))
		}
		return groupByNameExists
	}()
	searchDefined = searchDefined || groupByNameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for JobTemplate
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for JobTemplate
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = helpers.ExtractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	if err = hooks.RequireResourceStateOrOrig(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on JobTemplate",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
