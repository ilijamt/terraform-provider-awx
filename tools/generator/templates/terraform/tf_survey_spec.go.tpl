package {{ .PackageName }}

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

// Delete the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) { return }

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.{{ .Name }}ID.ValueInt64())) + "/"
	if framework.DiagnosticsHasError(&response.Diagnostics, framework.DeleteRequest(ctx, o.Client, endpoint, "{{ .Name }}/Survey")...) { return }
}

// Read the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) { return }

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.{{ .Name }}ID.ValueInt64())) + "/"
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "{{ .Name }}/Survey")
	if framework.DiagnosticsHasError(&response.Diagnostics, d...) { return }

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, {{ or .Trim false }})
		if dg.HasError() {
			response.Diagnostics.Append(dg...)
		}
	}

	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) { return }
}

// Create the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan, state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) { return }

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.{{ .Name }}ID.ValueInt64())) + "/"
	var bodyRequest = plan.BodyRequest()
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, bodyRequest, "{{ .Name }}/Survey", "create"); framework.DiagnosticsHasError(&response.Diagnostics, d...) { return }

    state.Spec = types.StringValue(plan.Spec.ValueString())
	state.{{ .Name }}ID = types.Int64Value(plan.{{ .Name }}ID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) { return }
}

// Update the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan, state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) { return }

	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.{{ .Name }}ID.ValueInt64())) + "/"
	var bodyRequest = plan.BodyRequest()
	if _, d := framework.CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, bodyRequest, "{{ .Name }}/Survey", "update"); framework.DiagnosticsHasError(&response.Diagnostics, d...) { return }

    state.Spec = types.StringValue(plan.Spec.ValueString())
	state.{{ .Name }}ID = types.Int64Value(plan.{{ .Name }}ID.ValueInt64())
	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) { return }
}
