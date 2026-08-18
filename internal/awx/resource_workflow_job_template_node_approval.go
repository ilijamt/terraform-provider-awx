package awx

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// Approval templates have no collection endpoint, so this cannot be generated.
// Create posts to a sub-endpoint of the node while read, update and delete work
// against a different collection keyed by the returned id, and GenericResource
// assumes one endpoint for all four.
const (
	approvalCreateEndpoint = "/api/v2/workflow_job_template_nodes/%d/create_approval_template/"
	approvalManageEndpoint = "/api/v2/workflow_approval_templates/"
)

var (
	_ resource.Resource                = (*workflowJobTemplateNodeApprovalResource)(nil)
	_ resource.ResourceWithConfigure   = (*workflowJobTemplateNodeApprovalResource)(nil)
	_ resource.ResourceWithImportState = (*workflowJobTemplateNodeApprovalResource)(nil)
)

type workflowJobTemplateNodeApprovalResource struct {
	framework.ResourceBase
}

// NewWorkflowJobTemplateNodeApprovalResource turns a workflow node into an
// approval gate. AWX repoints the node's unified_job_template at the template
// this creates and clears it again on delete, leaving a plain node behind.
func NewWorkflowJobTemplateNodeApprovalResource() resource.Resource {
	return &workflowJobTemplateNodeApprovalResource{
		ResourceBase: framework.ResourceBase{
			ProviderBase: framework.ProviderBase{
				TypeName: "workflow_job_template_node_approval",
				Endpoint: approvalManageEndpoint,
			},
		},
	}
}

type workflowJobTemplateNodeApprovalModel struct {
	ID                        types.Int64  `tfsdk:"id"`
	WorkflowJobTemplateNodeID types.Int64  `tfsdk:"workflow_job_template_node_id"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	Timeout                   types.Int64  `tfsdk:"timeout"`
	ExecutionEnvironment      types.Int64  `tfsdk:"execution_environment"`
}

func (o *workflowJobTemplateNodeApprovalResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Database ID for this workflow approval template.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"workflow_job_template_node_id": schema.Int64Attribute{
				Description: "Database ID of the workflow job template node to convert into an approval.",
				Required:    true,
				// The create endpoint is scoped to one node and AWX cannot move a
				// template to another, so a change has to recreate.
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of this workflow approval template.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this workflow approval template.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
			},
			"timeout": schema.Int64Attribute{
				Description: "Seconds to wait before the approval is marked as timed out. 0 waits forever.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"execution_environment": schema.Int64Attribute{
				Description: "Database ID of the execution environment to use.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (o *workflowJobTemplateNodeApprovalResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	// The template holds no reference back to its node, so the node id cannot be
	// recovered from a read and has to come in on the import string.
	parts := strings.Split(request.ID, "/")
	if len(parts) != 2 {
		response.Diagnostics.AddError(
			"Unable to import state for the workflow approval template, invalid format.",
			fmt.Sprintf("requires the identifier to be set to <workflow_job_template_node_id>/<id>, currently set to %s", request.ID),
		)
		return
	}

	nodeID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse %q as an int64, please provide the workflow_job_template_node_id.", parts[0]),
			err.Error(),
		)
		return
	}
	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse %q as an int64, please provide the id of the approval template.", parts[1]),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("workflow_job_template_node_id"), types.Int64Value(nodeID))...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), types.Int64Value(id))...)
}

func (o *workflowJobTemplateNodeApprovalResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan workflowJobTemplateNodeApprovalModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	endpoint := fmt.Sprintf(approvalCreateEndpoint, plan.WorkflowJobTemplateNodeID.ValueInt64())
	body := map[string]any{
		"name":        plan.Name.ValueString(),
		"description": plan.Description.ValueString(),
		"timeout":     plan.Timeout.ValueInt64(),
	}

	data, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, body, "WorkflowJobTemplateNodeApproval", "create")
	if framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	state := plan
	if framework.DiagnosticsHasError(&response.Diagnostics, state.fromApiData(data)...) {
		return
	}

	// execution_environment is absent from the create sub-endpoint and only
	// accepted by the manage collection, so it needs a follow-up PATCH.
	if !plan.ExecutionEnvironment.IsNull() && !plan.ExecutionEnvironment.IsUnknown() {
		if framework.DiagnosticsHasError(&response.Diagnostics, o.patch(ctx, &state, &plan)...) {
			return
		}
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *workflowJobTemplateNodeApprovalResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state workflowJobTemplateNodeApprovalModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	data, found, d := framework.ReadRequestAllowMissing(ctx, o.Client,
		framework.EndpointWithID(approvalManageEndpoint, state.ID.ValueInt64()), "WorkflowJobTemplateNodeApproval")
	if framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}
	if !found {
		response.State.RemoveResource(ctx)
		return
	}
	if framework.DiagnosticsHasError(&response.Diagnostics, state.fromApiData(data)...) {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *workflowJobTemplateNodeApprovalResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan, state workflowJobTemplateNodeApprovalModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	plan.ID = state.ID
	if framework.DiagnosticsHasError(&response.Diagnostics, o.patch(ctx, &plan, &plan)...) {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (o *workflowJobTemplateNodeApprovalResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state workflowJobTemplateNodeApprovalModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	response.Diagnostics.Append(framework.DeleteRequest(ctx, o.Client,
		framework.EndpointWithID(approvalManageEndpoint, state.ID.ValueInt64()), "WorkflowJobTemplateNodeApproval")...)
}

func (o *workflowJobTemplateNodeApprovalResource) patch(ctx context.Context, target, desired *workflowJobTemplateNodeApprovalModel) diag.Diagnostics {
	var diags diag.Diagnostics

	body := map[string]any{
		"name":        desired.Name.ValueString(),
		"description": desired.Description.ValueString(),
		"timeout":     desired.Timeout.ValueInt64(),
	}
	if !desired.ExecutionEnvironment.IsNull() && !desired.ExecutionEnvironment.IsUnknown() {
		body["execution_environment"] = desired.ExecutionEnvironment.ValueInt64()
	}

	data, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPatch,
		framework.EndpointWithID(approvalManageEndpoint, target.ID.ValueInt64()), body,
		"WorkflowJobTemplateNodeApproval", "update")
	if diags.Append(d...); diags.HasError() {
		return diags
	}

	diags.Append(target.fromApiData(data)...)
	return diags
}

// fromApiData deliberately leaves workflow_job_template_node_id alone, since the
// response carries no reference back to the node.
func (m *workflowJobTemplateNodeApprovalModel) fromApiData(data map[string]any) diag.Diagnostics {
	var diags diag.Diagnostics
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }

	collect(helpers.AttrValueSetInt64(&m.ID, data["id"]))
	collect(helpers.AttrValueSetString(&m.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&m.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&m.Timeout, data["timeout"]))
	collect(helpers.AttrValueSetInt64(&m.ExecutionEnvironment, data["execution_environment"]))

	return diags
}
