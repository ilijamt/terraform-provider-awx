package awx

import (
	"context"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type jobTemplateDataSource = framework.GenericDataSource[jobTemplateTerraformModel, *jobTemplateTerraformModel]

// NewJobTemplateDataSource is a helper function to instantiate the JobTemplate data source.
func NewJobTemplateDataSource() datasource.DataSource {
	return &jobTemplateDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "job_template", Endpoint: "/api/v2/job_templates/"}},
		Cfg: framework.DataSourceCfg[jobTemplateTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"allow_simultaneous": schema.BoolAttribute{
						Description: "Allow simultaneous",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_credential_on_launch": schema.BoolAttribute{
						Description: "Ask credential on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_diff_mode_on_launch": schema.BoolAttribute{
						Description: "Ask diff mode on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_execution_environment_on_launch": schema.BoolAttribute{
						Description: "Ask execution environment on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_forks_on_launch": schema.BoolAttribute{
						Description: "Ask forks on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_instance_groups_on_launch": schema.BoolAttribute{
						Description: "Ask instance groups on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_inventory_on_launch": schema.BoolAttribute{
						Description: "Ask inventory on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_job_slice_count_on_launch": schema.BoolAttribute{
						Description: "Ask job slice count on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_job_type_on_launch": schema.BoolAttribute{
						Description: "Ask job type on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_labels_on_launch": schema.BoolAttribute{
						Description: "Ask labels on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_limit_on_launch": schema.BoolAttribute{
						Description: "Ask limit on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_scm_branch_on_launch": schema.BoolAttribute{
						Description: "Ask scm branch on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_skip_tags_on_launch": schema.BoolAttribute{
						Description: "Ask skip tags on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_tags_on_launch": schema.BoolAttribute{
						Description: "Ask tags on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_timeout_on_launch": schema.BoolAttribute{
						Description: "Ask timeout on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_variables_on_launch": schema.BoolAttribute{
						Description: "Ask variables on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"ask_verbosity_on_launch": schema.BoolAttribute{
						Description: "Ask verbosity on launch",
						Sensitive:   false,
						Computed:    true,
					},
					"become_enabled": schema.BoolAttribute{
						Description: "Become enabled",
						Sensitive:   false,
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this job template.",
						Sensitive:   false,
						Computed:    true,
					},
					"diff_mode": schema.BoolAttribute{
						Description: "If enabled, textual changes made to any templated files on the host are shown in the standard output",
						Sensitive:   false,
						Computed:    true,
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Sensitive:   false,
						Computed:    true,
					},
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Sensitive:   false,
						Computed:    true,
					},
					"force_handlers": schema.BoolAttribute{
						Description: "Force handlers",
						Sensitive:   false,
						Computed:    true,
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Sensitive:   false,
						Computed:    true,
					},
					"host_config_key": schema.StringAttribute{
						Description: "Host config key",
						Sensitive:   false,
						Computed:    true,
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
					},
					"job_slice_count": schema.Int64Attribute{
						Description: "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1.",
						Sensitive:   false,
						Computed:    true,
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Sensitive:   false,
						Computed:    true,
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Sensitive:   false,
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Sensitive:   false,
						Computed:    true,
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
					},
					"playbook": schema.StringAttribute{
						Description: "Playbook",
						Sensitive:   false,
						Computed:    true,
					},
					"prevent_instance_group_fallback": schema.BoolAttribute{
						Description: "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
						Sensitive:   false,
						Computed:    true,
					},
					"project": schema.Int64Attribute{
						Description: "Project",
						Sensitive:   false,
						Computed:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true.",
						Sensitive:   false,
						Computed:    true,
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Sensitive:   false,
						Computed:    true,
					},
					"start_at_task": schema.StringAttribute{
						Description: "Start at task",
						Sensitive:   false,
						Computed:    true,
					},
					"survey_enabled": schema.BoolAttribute{
						Description: "Survey enabled",
						Sensitive:   false,
						Computed:    true,
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Sensitive:   false,
						Computed:    true,
					},
					"use_fact_cache": schema.BoolAttribute{
						Description: "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible.",
						Sensitive:   false,
						Computed:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Sensitive:   false,
						Computed:    true,
					},
					"webhook_credential": schema.Int64Attribute{
						Description: "Personal Access Token for posting back the status to the service API",
						Sensitive:   false,
						Computed:    true,
					},
					"webhook_service": schema.StringAttribute{
						Description: "Service that webhook requests will be accepted from",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *jobTemplateTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "JobTemplate",
		},
	}
}
