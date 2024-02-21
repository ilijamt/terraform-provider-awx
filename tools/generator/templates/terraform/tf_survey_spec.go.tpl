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
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
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
	return &{{ .Name | lowerCamelCase }}Survey{}
}

type {{ .Name | lowerCamelCase }}Survey struct {
    client   c.Client
    endpoint string
}

func (o *{{ .Name | lowerCamelCase }}Survey) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
    if request.ProviderData == nil {
        return
    }

    o.client = request.ProviderData.(c.Client)
    o.endpoint = "{{ .Endpoint | url_path_clean }}/%d/survey_spec/"
}

func (o *{{ .Name | lowerCamelCase }}Survey) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
    response.TypeName = request.ProviderTypeName + "_{{ .Name | snakeCase }}_survey_spec"
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
	var err error

	var state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.{{ .Name }}ID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for {{ .Name }}/Survey on %s", endpoint),
			err.Error(),
		)
		return
	}
}

// Read the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error
	var state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.{{ .Name }}ID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

    var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for {{ .Name }}/Survey on %s", endpoint),
			err.Error(),
		)
		return
	}

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, {{ default .trim false }})
		if dg.HasError() {
			response.Diagnostics.Append(dg...)
		}
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Create the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
    var err error
	var plan, state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.{{ .Name }}ID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[{{.Name}}/Create/Survey] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for {{ .Name }}/Survey on %s", endpoint),
			err.Error(),
		)
		return
	}

    state.Spec = types.StringValue(plan.Spec.ValueString())
	state.{{ .Name }}ID = types.Int64Value(plan.{{ .Name }}ID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Update the survey spec for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}Survey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
    var err error
	var plan, state {{ .Name | lowerCamelCase }}SurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.{{ .Name }}ID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[{{.Name}}/Update/SurveySpec] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for {{ .Name }}/Survey on %s", endpoint),
			err.Error(),
		)
		return
	}

    state.Spec = types.StringValue(plan.Spec.ValueString())
	state.{{ .Name }}ID = types.Int64Value(plan.{{ .Name }}ID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}
