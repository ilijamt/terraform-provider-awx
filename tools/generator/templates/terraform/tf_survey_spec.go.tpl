package {{ .PackageName }}

import (
	"bytes"
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
	_ resource.Resource                  = &{{ .Name | lowerCamelCase }}Survey{}
	_ resource.ResourceWithConfigure     = &{{ .Name | lowerCamelCase }}Survey{}
	_ resource.ResourceWithImportState   = &{{ .Name | lowerCamelCase }}Survey{}
)

type {{ .Name | lowerCamelCase }}SurveyTerraformModel struct {
	{{ .Name }}ID       types.Int64  `tfsdk:"{{ .Name | snakeCase }}_id"`
    Spec                types.String `tfsdk:"spec"`
}

func (o {{ .Name | lowerCamelCase }}SurveyTerraformModel) Clone() {{ .Name | lowerCamelCase }}SurveyTerraformModel {
    return {{ .Name | lowerCamelCase }}SurveyTerraformModel{
        {{ .Name }}ID: types.Int64Value(o.{{ .Name }}ID.ValueInt64()),
		Spec:          types.StringValue(o.Spec.ValueString()),
    }
}

func (o {{ .Name | lowerCamelCase }}SurveyTerraformModel) BodyRequest() {{ .Name | lowerCamelCase }}SurveyModel {
    return {{ .Name | lowerCamelCase }}SurveyModel{
		Spec: json.RawMessage(o.Spec.ValueString()),
    }
}

type {{ .Name | lowerCamelCase }}SurveyModel struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Spec        json.RawMessage `json:"spec"`
}

// New{{ .Name }}SurveyResource is a helper function to simplify the provider implementation.
func New{{ .Name }}SurveyResource() resource.Resource {
	return &{{ .Name | lowerCamelCase }}Survey{ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .Name | snakeCase }}_survey_spec", Endpoint: "{{ .Endpoint | url_path_clean }}/%d/survey_spec/"}}}
}

type {{ .Name | lowerCamelCase }}Survey struct {
    framework.ResourceBase
}

// endpointFor returns the survey-spec URL for a given parent ID.
func (o *{{ .Name | lowerCamelCase }}Survey) endpointFor(parentID int64) string {
    return p.Clean(fmt.Sprintf(o.Endpoint, parentID)) + "/"
}

func (o *{{ .Name | lowerCamelCase }}Survey) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
            Attributes: map[string]schema.Attribute{
				"{{ .Name | snakeCase }}_id": schema.Int64Attribute{
					Description: "Database ID for this {{ .Name }}.",
					Required:    true,
                    PlanModifiers: []planmodifier.Int64{
                        int64planmodifier.RequiresReplace(),
                    },
				},
				"spec": schema.StringAttribute{
					Description: "The survey spec for this {{ .Name }}.",
					Required:    true,
				},
            },
	    }
}

// ImportState imports the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the {{ .Name }}.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("{{ .Name | snakeCase }}_id"), id)...)
}

func (o *{{ .Name | lowerCamelCase }}Survey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) { return }

	endpoint := o.endpointFor(state.{{ .Name }}ID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, framework.DeleteRequest(ctx, o.Client, endpoint, "{{ .Name }}/Survey")...) { return }
}

func (o *{{ .Name | lowerCamelCase }}Survey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) { return }

	endpoint := o.endpointFor(state.{{ .Name }}ID.ValueInt64())
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "{{ .Name }}/Survey")
	if framework.DiagnosticsHasError(&response.Diagnostics, d...) { return }

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, {{ or .Trim false }})
		response.Diagnostics.Append(dg...)
	}

	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) { return }
}

// applyMutation handles Create and Update — both POST a survey spec for the
// parent ID. AWX returns no useful body, so we mirror the plan into state.
func (o *{{ .Name | lowerCamelCase }}Survey) applyMutation(ctx context.Context, plan {{ .Name | lowerCamelCase }}SurveyTerraformModel, operation string, diags *diag.Diagnostics) ({{ .Name | lowerCamelCase }}SurveyTerraformModel, bool) {
	endpoint := o.endpointFor(plan.{{ .Name }}ID.ValueInt64())
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, plan.BodyRequest(), "{{ .Name }}/Survey", operation); framework.DiagnosticsHasError(diags, d...) {
		return {{ .Name | lowerCamelCase }}SurveyTerraformModel{}, false
	}
	return {{ .Name | lowerCamelCase }}SurveyTerraformModel{
		{{ .Name }}ID: types.Int64Value(plan.{{ .Name }}ID.ValueInt64()),
		Spec:          types.StringValue(plan.Spec.ValueString()),
	}, true
}

func (o *{{ .Name | lowerCamelCase }}Survey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) { return }

	state, ok := o.applyMutation(ctx, plan, "create", &response.Diagnostics)
	if !ok { return }
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *{{ .Name | lowerCamelCase }}Survey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) { return }

	state, ok := o.applyMutation(ctx, plan, "update", &response.Diagnostics)
	if !ok { return }
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}
