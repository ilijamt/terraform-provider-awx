package awx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

var (
	_ resource.Resource                = &workflowJobTemplateSurvey{}
	_ resource.ResourceWithConfigure   = &workflowJobTemplateSurvey{}
	_ resource.ResourceWithImportState = &workflowJobTemplateSurvey{}
)

type workflowJobTemplateSurveyTerraformModel struct {
	WorkflowJobTemplateID types.Int64  `tfsdk:"workflow_job_template_id"`
	Spec                  types.String `tfsdk:"spec"`
}

func (o workflowJobTemplateSurveyTerraformModel) Clone() workflowJobTemplateSurveyTerraformModel {
	return workflowJobTemplateSurveyTerraformModel{
		WorkflowJobTemplateID: types.Int64Value(o.WorkflowJobTemplateID.ValueInt64()),
		Spec:                  types.StringValue(o.Spec.ValueString()),
	}
}

func (o workflowJobTemplateSurveyTerraformModel) BodyRequest() workflowJobTemplateSurveyModel {
	return workflowJobTemplateSurveyModel{
		Spec: json.RawMessage(o.Spec.ValueString()),
	}
}

type workflowJobTemplateSurveyModel struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Spec        json.RawMessage `json:"spec"`
}

// NewWorkflowJobTemplateSurveyResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateSurveyResource() resource.Resource {
	return &workflowJobTemplateSurvey{ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template_survey_spec", Endpoint: "/api/v2/workflow_job_templates/%d/survey_spec/"}}}
}

type workflowJobTemplateSurvey struct {
	framework.ResourceBase
}

func (o *workflowJobTemplateSurvey) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"workflow_job_template_id": schema.Int64Attribute{
				Description: "Database ID for this WorkflowJobTemplate.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"spec": schema.StringAttribute{
				Description: "The survey spec for this WorkflowJobTemplate.",
				Required:    true,
			},
		},
	}
}

// ImportState imports the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the WorkflowJobTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("workflow_job_template_id"), id)...)
}

// Delete the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.WorkflowJobTemplateID.ValueInt64())) + "/"
	if framework.DiagnosticsHasError(&response.Diagnostics, framework.DeleteRequest(ctx, o.Client, endpoint, "WorkflowJobTemplate/Survey")...) {
		return
	}
}

// Read the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.WorkflowJobTemplateID.ValueInt64())) + "/"
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "WorkflowJobTemplate/Survey")
	if framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, false)
		if dg.HasError() {
			response.Diagnostics.Append(dg...)
		}
	}

	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// Create the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan, state workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.WorkflowJobTemplateID.ValueInt64())) + "/"
	var bodyRequest = plan.BodyRequest()
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, bodyRequest, "WorkflowJobTemplate/Survey", "create"); framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.WorkflowJobTemplateID = types.Int64Value(plan.WorkflowJobTemplateID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// Update the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan, state workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.WorkflowJobTemplateID.ValueInt64())) + "/"
	var bodyRequest = plan.BodyRequest()
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, bodyRequest, "WorkflowJobTemplate/Survey", "update"); framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.WorkflowJobTemplateID = types.Int64Value(plan.WorkflowJobTemplateID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}
