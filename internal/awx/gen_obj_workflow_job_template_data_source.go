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

type workflowJobTemplateDataSource = framework.GenericDataSource[workflowJobTemplateTerraformModel, *workflowJobTemplateTerraformModel]

// NewWorkflowJobTemplateDataSource is a helper function to instantiate the WorkflowJobTemplate data source.
func NewWorkflowJobTemplateDataSource() datasource.DataSource {
	return &workflowJobTemplateDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template", Endpoint: "/api/v2/workflow_job_templates/"}},
		Cfg: framework.DataSourceCfg[workflowJobTemplateTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"allow_simultaneous": schema.BoolAttribute{
						Description: "Allow simultaneous",
						Computed:    true,
					},
					"ask_inventory_on_launch": schema.BoolAttribute{
						Description: "Ask inventory on launch",
						Computed:    true,
					},
					"ask_labels_on_launch": schema.BoolAttribute{
						Description: "Ask labels on launch",
						Computed:    true,
					},
					"ask_limit_on_launch": schema.BoolAttribute{
						Description: "Ask limit on launch",
						Computed:    true,
					},
					"ask_scm_branch_on_launch": schema.BoolAttribute{
						Description: "Ask scm branch on launch",
						Computed:    true,
					},
					"ask_skip_tags_on_launch": schema.BoolAttribute{
						Description: "Ask skip tags on launch",
						Computed:    true,
					},
					"ask_tags_on_launch": schema.BoolAttribute{
						Description: "Ask tags on launch",
						Computed:    true,
					},
					"ask_variables_on_launch": schema.BoolAttribute{
						Description: "Ask variables on launch",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this workflow job template.",
						Computed:    true,
					},
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this workflow job template.",
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
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Computed:    true,
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this workflow job template.",
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
						Computed:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Scm branch",
						Computed:    true,
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Computed:    true,
					},
					"survey_enabled": schema.BoolAttribute{
						Description: "Survey enabled",
						Computed:    true,
					},
					"webhook_credential": schema.Int64Attribute{
						Description: "Personal Access Token for posting back the status to the service API",
						Computed:    true,
					},
					"webhook_service": schema.StringAttribute{
						Description: "Service that webhook requests will be accepted from",
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
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *workflowJobTemplateTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplate",
		},
	}
}
