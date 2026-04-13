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
	_ resource.Resource                = &jobTemplateSurvey{}
	_ resource.ResourceWithConfigure   = &jobTemplateSurvey{}
	_ resource.ResourceWithImportState = &jobTemplateSurvey{}
)

type jobTemplateSurveyTerraformModel struct {
	JobTemplateID types.Int64  `tfsdk:"job_template_id"`
	Spec          types.String `tfsdk:"spec"`
}

func (o jobTemplateSurveyTerraformModel) Clone() jobTemplateSurveyTerraformModel {
	return jobTemplateSurveyTerraformModel{
		JobTemplateID: types.Int64Value(o.JobTemplateID.ValueInt64()),
		Spec:          types.StringValue(o.Spec.ValueString()),
	}
}

func (o jobTemplateSurveyTerraformModel) BodyRequest() jobTemplateSurveyModel {
	return jobTemplateSurveyModel{
		Spec: json.RawMessage(o.Spec.ValueString()),
	}
}

type jobTemplateSurveyModel struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Spec        json.RawMessage `json:"spec"`
}

// NewJobTemplateSurveyResource is a helper function to simplify the provider implementation.
func NewJobTemplateSurveyResource() resource.Resource {
	return &jobTemplateSurvey{ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "job_template_survey_spec", Endpoint: "/api/v2/job_templates/%d/survey_spec/"}}}
}

type jobTemplateSurvey struct {
	framework.ResourceBase
}

func (o *jobTemplateSurvey) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"job_template_id": schema.Int64Attribute{
				Description: "Database ID for this JobTemplate.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"spec": schema.StringAttribute{
				Description: "The survey spec for this JobTemplate.",
				Required:    true,
			},
		},
	}
}

// ImportState imports the survey spec for JobTemplate
func (o *jobTemplateSurvey) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the JobTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("job_template_id"), id)...)
}

// Delete the survey spec for JobTemplate
func (o *jobTemplateSurvey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state jobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.JobTemplateID.ValueInt64())) + "/"
	if framework.DiagnosticsHasError(&response.Diagnostics, framework.DeleteRequest(ctx, o.Client, endpoint, "JobTemplate/Survey")...) {
		return
	}
}

// Read the survey spec for JobTemplate
func (o *jobTemplateSurvey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state jobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.JobTemplateID.ValueInt64())) + "/"
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "JobTemplate/Survey")
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

// Create the survey spec for JobTemplate
func (o *jobTemplateSurvey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan, state jobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var bodyRequest = plan.BodyRequest()
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, bodyRequest, "JobTemplate/Survey", "create"); framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.JobTemplateID = types.Int64Value(plan.JobTemplateID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// Update the survey spec for JobTemplate
func (o *jobTemplateSurvey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan, state jobTemplateSurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var bodyRequest = plan.BodyRequest()
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, bodyRequest, "JobTemplate/Survey", "update"); framework.DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.JobTemplateID = types.Int64Value(plan.JobTemplateID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}
