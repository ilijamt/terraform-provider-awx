package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type workflowJobTemplateNodeTerraformModel struct {
	AllParentsMustConverge types.Bool   `tfsdk:"all_parents_must_converge" json:"all_parents_must_converge"`
	AlwaysNodes            types.List   `tfsdk:"always_nodes" json:"always_nodes"`
	DiffMode               types.Bool   `tfsdk:"diff_mode" json:"diff_mode"`
	ExecutionEnvironment   types.Int64  `tfsdk:"execution_environment" json:"execution_environment"`
	ExtraData              types.String `tfsdk:"extra_data" json:"extra_data"`
	FailureNodes           types.List   `tfsdk:"failure_nodes" json:"failure_nodes"`
	Forks                  types.Int64  `tfsdk:"forks" json:"forks"`
	ID                     types.Int64  `tfsdk:"id" json:"id"`
	Identifier             types.String `tfsdk:"identifier" json:"identifier"`
	Inventory              types.Int64  `tfsdk:"inventory" json:"inventory"`
	JobSliceCount          types.Int64  `tfsdk:"job_slice_count" json:"job_slice_count"`
	JobTags                types.String `tfsdk:"job_tags" json:"job_tags"`
	JobType                types.String `tfsdk:"job_type" json:"job_type"`
	Limit                  types.String `tfsdk:"limit" json:"limit"`
	ScmBranch              types.String `tfsdk:"scm_branch" json:"scm_branch"`
	SkipTags               types.String `tfsdk:"skip_tags" json:"skip_tags"`
	SuccessNodes           types.List   `tfsdk:"success_nodes" json:"success_nodes"`
	Timeout                types.Int64  `tfsdk:"timeout" json:"timeout"`
	UnifiedJobTemplate     types.Int64  `tfsdk:"unified_job_template" json:"unified_job_template"`
	Verbosity              types.String `tfsdk:"verbosity" json:"verbosity"`
	WorkflowJobTemplate    types.Int64  `tfsdk:"workflow_job_template" json:"workflow_job_template"`
}

func (o *workflowJobTemplateNodeTerraformModel) Clone() workflowJobTemplateNodeTerraformModel {
	return *o
}

func (o *workflowJobTemplateNodeTerraformModel) BodyRequest() *workflowJobTemplateNodeBodyRequestModel {
	var req workflowJobTemplateNodeBodyRequestModel
	req.AllParentsMustConverge = o.AllParentsMustConverge.ValueBool()
	req.DiffMode = helpers.AttrBoolPointer(o.DiffMode)
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.ExtraData = json.RawMessage(o.ExtraData.ValueString())
	req.Forks = o.Forks.ValueInt64()
	req.Identifier = o.Identifier.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.JobSliceCount = o.JobSliceCount.ValueInt64()
	req.JobTags = o.JobTags.ValueString()
	req.JobType = o.JobType.ValueString()
	req.Limit = o.Limit.ValueString()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.SkipTags = o.SkipTags.ValueString()
	req.Timeout = o.Timeout.ValueInt64()
	req.UnifiedJobTemplate = o.UnifiedJobTemplate.ValueInt64()
	req.Verbosity = o.Verbosity.ValueString()
	req.WorkflowJobTemplate = o.WorkflowJobTemplate.ValueInt64()
	return &req
}

func (o *workflowJobTemplateNodeTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.AllParentsMustConverge, data["all_parents_must_converge"]))
	collect(helpers.AttrValueSetListInt64(&o.AlwaysNodes, data["always_nodes"]))
	collect(helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"]))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetJsonString(&o.ExtraData, data["extra_data"], false))
	collect(helpers.AttrValueSetListInt64(&o.FailureNodes, data["failure_nodes"]))
	collect(helpers.AttrValueSetInt64(&o.Forks, data["forks"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Identifier, data["identifier"], false))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetInt64(&o.JobSliceCount, data["job_slice_count"]))
	collect(helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false))
	collect(helpers.AttrValueSetString(&o.JobType, data["job_type"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetString(&o.SkipTags, data["skip_tags"], false))
	collect(helpers.AttrValueSetListInt64(&o.SuccessNodes, data["success_nodes"]))
	collect(helpers.AttrValueSetInt64(&o.Timeout, data["timeout"]))
	collect(helpers.AttrValueSetInt64(&o.UnifiedJobTemplate, data["unified_job_template"]))
	collect(helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false))
	collect(helpers.AttrValueSetInt64(&o.WorkflowJobTemplate, data["workflow_job_template"]))
	return diags, nil
}

type workflowJobTemplateNodeBodyRequestModel struct {
	AllParentsMustConverge bool            `json:"all_parents_must_converge"`
	DiffMode               *bool           `json:"diff_mode,omitempty"`
	ExecutionEnvironment   int64           `json:"execution_environment,omitempty"`
	ExtraData              json.RawMessage `json:"extra_data,omitempty"`
	Forks                  int64           `json:"forks,omitempty"`
	Identifier             string          `json:"identifier,omitempty"`
	Inventory              int64           `json:"inventory,omitempty"`
	JobSliceCount          int64           `json:"job_slice_count,omitempty"`
	JobTags                string          `json:"job_tags,omitempty"`
	JobType                string          `json:"job_type,omitempty"`
	Limit                  string          `json:"limit,omitempty"`
	ScmBranch              string          `json:"scm_branch,omitempty"`
	SkipTags               string          `json:"skip_tags,omitempty"`
	Timeout                int64           `json:"timeout,omitempty"`
	UnifiedJobTemplate     int64           `json:"unified_job_template,omitempty"`
	Verbosity              string          `json:"verbosity,omitempty"`
	WorkflowJobTemplate    int64           `json:"workflow_job_template"`
}

type workflowJobTemplateNodeResource = framework.GenericResource[workflowJobTemplateNodeTerraformModel, workflowJobTemplateNodeBodyRequestModel, *workflowJobTemplateNodeTerraformModel]

// NewWorkflowJobTemplateNodeResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateNodeResource() resource.Resource {
	return &workflowJobTemplateNodeResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template_node", Endpoint: "/api/v2/workflow_job_template_nodes/"}},
		Cfg: framework.ResourceCfg[workflowJobTemplateNodeTerraformModel, workflowJobTemplateNodeBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"all_parents_must_converge": schema.BoolAttribute{
						Description: "If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
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
					"extra_data": schema.StringAttribute{
						Description: "Extra data",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"identifier": schema.StringAttribute{
						Description: "An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
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
					"job_slice_count": schema.Int64Attribute{
						Description: "Job slice count",
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
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"run",
								"check",
							),
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
					"timeout": schema.Int64Attribute{
						Description: "Timeout",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"unified_job_template": schema.Int64Attribute{
						Description: "Unified job template",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
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
					"workflow_job_template": schema.Int64Attribute{
						Description: "Workflow job template",
						Required:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"always_nodes": schema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Always nodes",
						Computed:    true,
					},
					"failure_nodes": schema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Failure nodes",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this workflow job template node.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"success_nodes": schema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Success nodes",
						Computed:    true,
					},
				},
			},
			IDAccessor:   func(m *workflowJobTemplateNodeTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplateNode",
		},
	}
}

type workflowJobTemplateNodeDataSource = framework.GenericDataSource[workflowJobTemplateNodeTerraformModel, *workflowJobTemplateNodeTerraformModel]

// NewWorkflowJobTemplateNodeDataSource is a helper function to instantiate the WorkflowJobTemplateNode data source.
func NewWorkflowJobTemplateNodeDataSource() datasource.DataSource {
	return &workflowJobTemplateNodeDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template_node", Endpoint: "/api/v2/workflow_job_template_nodes/"}},
		Cfg: framework.DataSourceCfg[workflowJobTemplateNodeTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"all_parents_must_converge": dschema.BoolAttribute{
						Description: "If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node",
						Computed:    true,
					},
					"always_nodes": dschema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Always nodes",
						Computed:    true,
					},
					"diff_mode": dschema.BoolAttribute{
						Description: "Diff mode",
						Computed:    true,
					},
					"execution_environment": dschema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"extra_data": dschema.StringAttribute{
						Description: "Extra data",
						Computed:    true,
					},
					"failure_nodes": dschema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Failure nodes",
						Computed:    true,
					},
					"forks": dschema.Int64Attribute{
						Description: "Forks",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this workflow job template node.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"identifier": dschema.StringAttribute{
						Description: "An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node.",
						Computed:    true,
					},
					"inventory": dschema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Computed:    true,
					},
					"job_slice_count": dschema.Int64Attribute{
						Description: "Job slice count",
						Computed:    true,
					},
					"job_tags": dschema.StringAttribute{
						Description: "Job tags",
						Computed:    true,
					},
					"job_type": dschema.StringAttribute{
						Description: "Job type",
						Computed:    true,
					},
					"limit": dschema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"scm_branch": dschema.StringAttribute{
						Description: "Scm branch",
						Computed:    true,
					},
					"skip_tags": dschema.StringAttribute{
						Description: "Skip tags",
						Computed:    true,
					},
					"success_nodes": dschema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Success nodes",
						Computed:    true,
					},
					"timeout": dschema.Int64Attribute{
						Description: "Timeout",
						Computed:    true,
					},
					"unified_job_template": dschema.Int64Attribute{
						Description: "Unified job template",
						Computed:    true,
					},
					"verbosity": dschema.StringAttribute{
						Description: "Verbosity",
						Computed:    true,
					},
					"workflow_job_template": dschema.Int64Attribute{
						Description: "Workflow job template",
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
			ResourceName: "WorkflowJobTemplateNode",
		},
	}
}
