package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type workflowJobTemplateNodeApprovalTerraformModel struct {
	Description               types.String `tfsdk:"description" json:"description"`
	ExecutionEnvironment      types.Int64  `tfsdk:"execution_environment" json:"execution_environment"`
	ID                        types.Int64  `tfsdk:"id" json:"id"`
	Name                      types.String `tfsdk:"name" json:"name"`
	Status                    types.String `tfsdk:"status" json:"status"`
	Timeout                   types.Int64  `tfsdk:"timeout" json:"timeout"`
	WorkflowJobTemplateNodeId types.Int64  `tfsdk:"workflow_job_template_node_id" json:"workflow_job_template_node_id"`
}

func (o *workflowJobTemplateNodeApprovalTerraformModel) Clone() workflowJobTemplateNodeApprovalTerraformModel {
	return *o
}

func (o *workflowJobTemplateNodeApprovalTerraformModel) BodyRequest() *workflowJobTemplateNodeApprovalBodyRequestModel {
	var req workflowJobTemplateNodeApprovalBodyRequestModel
	req.Description = o.Description.ValueString()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Timeout = o.Timeout.ValueInt64()
	return &req
}

func (o *workflowJobTemplateNodeApprovalTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Status, data["status"], false))
	collect(helpers.AttrValueSetInt64(&o.Timeout, data["timeout"]))
	return diags, nil
}

type workflowJobTemplateNodeApprovalBodyRequestModel struct {
	Description               string `json:"description,omitempty"`
	ExecutionEnvironment      int64  `json:"execution_environment,omitempty"`
	Name                      string `json:"name"`
	Timeout                   int64  `json:"timeout"`
	WorkflowJobTemplateNodeId int64  `json:"workflow_job_template_node_id"`
}

type workflowJobTemplateNodeApprovalResource = framework.GenericResource[workflowJobTemplateNodeApprovalTerraformModel, workflowJobTemplateNodeApprovalBodyRequestModel, *workflowJobTemplateNodeApprovalTerraformModel]

// NewWorkflowJobTemplateNodeApprovalResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateNodeApprovalResource() resource.Resource {
	return &workflowJobTemplateNodeApprovalResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template_node_approval", Endpoint: "/api/v2/workflow_approval_templates/"}},
		Cfg: framework.ResourceCfg[workflowJobTemplateNodeApprovalTerraformModel, workflowJobTemplateNodeApprovalBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this workflow approval template.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"name": schema.StringAttribute{
						Description: "Name of this workflow approval template.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) before the approval node expires and fails.",
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
					"workflow_job_template_node_id": schema.Int64Attribute{
						Description: "Database ID of the workflow job template node to convert into an approval.",
						Required:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this workflow approval template.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"status": schema.StringAttribute{
						Description: "Status",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"new",
								"pending",
								"waiting",
								"running",
								"successful",
								"failed",
								"error",
								"canceled",
								"never updated",
								"ok",
								"missing",
								"none",
								"updating",
							),
						},
					},
				},
			},
			IDAccessor: func(m *workflowJobTemplateNodeApprovalTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			CreateEndpoint: func(m *workflowJobTemplateNodeApprovalTerraformModel) string {
				return fmt.Sprintf("/api/v2/workflow_job_template_nodes/%d/create_approval_template/", m.WorkflowJobTemplateNodeId.ValueInt64())
			},
			ImportIDParts: []string{"workflow_job_template_node_id", "id"},
			WriteOnlyPlanToBody: func(plan *workflowJobTemplateNodeApprovalTerraformModel, body *workflowJobTemplateNodeApprovalBodyRequestModel) {
				body.WorkflowJobTemplateNodeId = plan.WorkflowJobTemplateNodeId.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *workflowJobTemplateNodeApprovalTerraformModel) {
				state.WorkflowJobTemplateNodeId = types.Int64Value(plan.WorkflowJobTemplateNodeId.ValueInt64())
			},
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplateNodeApproval",
		},
	}
}

type workflowJobTemplateNodeApprovalDataSource = framework.GenericDataSource[workflowJobTemplateNodeApprovalTerraformModel, *workflowJobTemplateNodeApprovalTerraformModel]

// NewWorkflowJobTemplateNodeApprovalDataSource is a helper function to instantiate the WorkflowJobTemplateNodeApproval data source.
func NewWorkflowJobTemplateNodeApprovalDataSource() datasource.DataSource {
	return &workflowJobTemplateNodeApprovalDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template_node_approval", Endpoint: "/api/v2/workflow_approval_templates/"}},
		Cfg: framework.DataSourceCfg[workflowJobTemplateNodeApprovalTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this workflow approval template.",
						Computed:    true,
					},
					"execution_environment": dschema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this workflow approval template.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"name": dschema.StringAttribute{
						Description: "Name of this workflow approval template.",
						Computed:    true,
					},
					"status": dschema.StringAttribute{
						Description: "Status",
						Computed:    true,
					},
					"timeout": dschema.Int64Attribute{
						Description: "The amount of time (in seconds) before the approval node expires and fails.",
						Computed:    true,
					},
					"workflow_job_template_node_id": dschema.Int64Attribute{
						Description: "Database ID of the workflow job template node to convert into an approval.",
						Optional:    true,
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplateNodeApproval",
		},
	}
}
