package awx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// endpointFor returns the survey-spec URL for a given parent ID.
func (o *workflowJobTemplateSurvey) endpointFor(parentID int64) string {
	return p.Clean(fmt.Sprintf(o.Endpoint, parentID)) + "/"
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

func (o *workflowJobTemplateSurvey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	endpoint := o.endpointFor(state.WorkflowJobTemplateID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, framework.DeleteRequest(ctx, o.Client, endpoint, "WorkflowJobTemplate/Survey")...) {
		return
	}
}

func (o *workflowJobTemplateSurvey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	endpoint := o.endpointFor(state.WorkflowJobTemplateID.ValueInt64())
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "WorkflowJobTemplate/Survey")
	if framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, false)
		response.Diagnostics.Append(dg...)
	}

	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// applyMutation handles Create and Update — both POST a survey spec for the
// parent ID. AWX returns no useful body, so we mirror the plan into state.
func (o *workflowJobTemplateSurvey) applyMutation(ctx context.Context, plan workflowJobTemplateSurveyTerraformModel, operation string, diags *diag.Diagnostics) (workflowJobTemplateSurveyTerraformModel, bool) {
	endpoint := o.endpointFor(plan.WorkflowJobTemplateID.ValueInt64())
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, plan.BodyRequest(), "WorkflowJobTemplate/Survey", operation); framework.DiagnosticsHasError(diags, d...) {
		return workflowJobTemplateSurveyTerraformModel{}, false
	}
	return workflowJobTemplateSurveyTerraformModel{
		WorkflowJobTemplateID: types.Int64Value(plan.WorkflowJobTemplateID.ValueInt64()),
		Spec:                  types.StringValue(plan.Spec.ValueString()),
	}, true
}

func (o *workflowJobTemplateSurvey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	state, ok := o.applyMutation(ctx, plan, "create", &response.Diagnostics)
	if !ok {
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *workflowJobTemplateSurvey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan workflowJobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	state, ok := o.applyMutation(ctx, plan, "update", &response.Diagnostics)
	if !ok {
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}
