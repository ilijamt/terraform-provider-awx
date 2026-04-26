package awx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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

type workflowJobTemplateResource = framework.GenericResource[workflowJobTemplateTerraformModel, workflowJobTemplateBodyRequestModel, *workflowJobTemplateTerraformModel]

// NewWorkflowJobTemplateResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateResource() resource.Resource {
	return &workflowJobTemplateResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template", Endpoint: "/api/v2/workflow_job_templates/"}},
		Cfg: framework.ResourceCfg[workflowJobTemplateTerraformModel, workflowJobTemplateBodyRequestModel]{
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
					"ask_inventory_on_launch": schema.BoolAttribute{
						Description: "Ask inventory on launch",
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
					"ask_variables_on_launch": schema.BoolAttribute{
						Description: "Ask variables on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this workflow job template.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"inventory": schema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this workflow job template.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"scm_branch": schema.StringAttribute{
						Description: "Scm branch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
						Description: "Database ID for this workflow job template.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *workflowJobTemplateTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *workflowJobTemplateTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplate",
		},
	}
}
