package awx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type jobTemplateResource = framework.GenericResource[jobTemplateTerraformModel, jobTemplateBodyRequestModel, *jobTemplateTerraformModel]

// NewJobTemplateResource is a helper function to simplify the provider implementation.
func NewJobTemplateResource() resource.Resource {
	return &jobTemplateResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "job_template", Endpoint: "/api/v2/job_templates/"}},
		Cfg: framework.ResourceCfg[jobTemplateTerraformModel, jobTemplateBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"allow_simultaneous": schema.BoolAttribute{
						Description: "Allow simultaneous",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_credential_on_launch": schema.BoolAttribute{
						Description: "Ask credential on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_diff_mode_on_launch": schema.BoolAttribute{
						Description: "Ask diff mode on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_execution_environment_on_launch": schema.BoolAttribute{
						Description: "Ask execution environment on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_forks_on_launch": schema.BoolAttribute{
						Description: "Ask forks on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_instance_groups_on_launch": schema.BoolAttribute{
						Description: "Ask instance groups on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_inventory_on_launch": schema.BoolAttribute{
						Description: "Ask inventory on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_job_slice_count_on_launch": schema.BoolAttribute{
						Description: "Ask job slice count on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_job_type_on_launch": schema.BoolAttribute{
						Description: "Ask job type on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_labels_on_launch": schema.BoolAttribute{
						Description: "Ask labels on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_limit_on_launch": schema.BoolAttribute{
						Description: "Ask limit on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_scm_branch_on_launch": schema.BoolAttribute{
						Description: "Ask scm branch on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_skip_tags_on_launch": schema.BoolAttribute{
						Description: "Ask skip tags on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_tags_on_launch": schema.BoolAttribute{
						Description: "Ask tags on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_timeout_on_launch": schema.BoolAttribute{
						Description: "Ask timeout on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_variables_on_launch": schema.BoolAttribute{
						Description: "Ask variables on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_verbosity_on_launch": schema.BoolAttribute{
						Description: "Ask verbosity on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"become_enabled": schema.BoolAttribute{
						Description: "Become enabled",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this job template.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"diff_mode": schema.BoolAttribute{
						Description: "If enabled, textual changes made to any templated files on the host are shown in the standard output",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"force_handlers": schema.BoolAttribute{
						Description: "Force handlers",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 2147483647),
						},
					},
					"host_config_key": schema.StringAttribute{
						Description: "Host config key",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_slice_count": schema.Int64Attribute{
						Description: "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 2147483647),
						},
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`run`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"run",
								"check",
							),
						},
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this job template.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"playbook": schema.StringAttribute{
						Description: "Playbook",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"prevent_instance_group_fallback": schema.BoolAttribute{
						Description: "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"project": schema.Int64Attribute{
						Description: "Project",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"scm_branch": schema.StringAttribute{
						Description: "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"start_at_task": schema.StringAttribute{
						Description: "Start at task",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"survey_enabled": schema.BoolAttribute{
						Description: "Survey enabled",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(-2147483648, 2147483647),
						},
					},
					"use_fact_cache": schema.BoolAttribute{
						Description: "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`0`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"0",
								"1",
								"2",
								"3",
								"4",
								"5",
							),
						},
					},
					"webhook_credential": schema.Int64Attribute{
						Description: "Personal Access Token for posting back the status to the service API",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"webhook_service": schema.StringAttribute{
						Description: "Service that webhook requests will be accepted from",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"github",
								"gitlab",
								"bitbucket_dc",
							),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this job template.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *jobTemplateTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *jobTemplateTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "JobTemplate",
		},
	}
}
