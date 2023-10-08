package awx

import (
	"bytes"
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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
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
	return &jobTemplateSurvey{}
}

type jobTemplateSurvey struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateSurvey) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/survey_spec/"
}

func (o *jobTemplateSurvey) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template_survey_spec"
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
	var err error

	var state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

// Read the survey spec for JobTemplate
func (o *jobTemplateSurvey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error
	var state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, false)
		if dg.HasError() {
			response.Diagnostics.Append(dg...)
		}
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Create the survey spec for JobTemplate
func (o *jobTemplateSurvey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[JobTemplate/Create/Survey] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.JobTemplateID = types.Int64Value(plan.JobTemplateID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Update the survey spec for JobTemplate
func (o *jobTemplateSurvey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[JobTemplate/Update/SurveySpec] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.JobTemplateID = types.Int64Value(plan.JobTemplateID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}
