package {{ .PackageName }}

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
    "github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &{{ .Name | lowerCamelCase }}Resource{}
	_ resource.ResourceWithConfigure   = &{{ .Name | lowerCamelCase }}Resource{}
{{- if not $.Config.NoId }}
	_ resource.ResourceWithImportState = &{{ .Name | lowerCamelCase }}Resource{}
{{- end }}
)

// New{{ .Name }}Resource is a helper function to simplify the provider implementation.
func New{{ .Name }}Resource() resource.Resource {
	return &{{ .Name | lowerCamelCase }}Resource{}
}

type {{ .Name | lowerCamelCase }}Resource struct {
    client   c.Client
    endpoint string
}

func (o *{{ .Name | lowerCamelCase }}Resource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
    if request.ProviderData == nil {
        return
    }

    o.client = request.ProviderData.(c.Client)
    o.endpoint = "{{ $.Endpoint }}"
}

func (o *{{ .Name | lowerCamelCase }}Resource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_{{ $.Config.TypeName }}"
}

func (o *{{ .Name | lowerCamelCase }}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
        // Request elements
{{- range $key := .PropertyPostKeys }}
{{- with (index $.PropertyPostData $key) }}
            "{{ $key | lowerCase }}": schema.{{ tf_attribute_type . }}Attribute{
{{- if eq (tf_attribute_type .) "List" }}
				ElementType: types.{{ camelCase .element_type }}Type,
{{- end }}
                Description: {{ escape_quotes (default .help_text .label) }},
                Sensitive:   {{ .sensitive }},
                Required:    {{ .required }},
                Optional:    {{ not .required }},
                Computed:    {{ .computed }},
{{- if and (hasKey . "default") (hasKey . "default_value") (ne .default nil) }}
                Default:     {{ .default_value }},
{{- end }}
		        PlanModifiers: []planmodifier.{{ tf_attribute_type . }} {
{{- if not .required }}
                    {{ tf_attribute_type . | lowerCase }}planmodifier.UseStateForUnknown(),
{{- end }}
                },
				Validators: []validator.{{ tf_attribute_type . }}{
{{- if and (eq (awx2go_value .) "types.StringValue") (hasKey . "max_length") }}
					stringvalidator.LengthAtMost({{ .max_length }}),
{{- else if and (eq (awx2go_value .) "types.Int64Value") (hasKey . "min_value") (hasKey . "max_value") }}
					int64validator.Between({{ .min_value }}, {{ .max_value }}),
{{- else if eq .type "choice" }}
					stringvalidator.OneOf({{ awx_type_choice_data .choices }}...),
{{- end }}
                },
            },
{{- end }}
{{- end }}
        // Write only elements
{{- range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
            "{{ $key | lowerCase }}": schema.{{ tf_attribute_type . }}Attribute{
{{- if eq (tf_attribute_type .) "List" }}
				ElementType: types.{{ camelCase .element_type }}Type,
{{- end }}
                Description: {{ escape_quotes (default .help_text .label) }},
                Sensitive:   {{ .sensitive }},
                Required:    {{ .required }},
                Optional:    {{ not .required }},
                Computed:    {{ .computed }},
{{- if and (hasKey . "default") (hasKey . "default_value") (ne .default nil) }}
                Default:     {{ .default_value }},
{{- end }}
		        PlanModifiers: []planmodifier.{{ tf_attribute_type . }} {
{{- if not .required }}
                    {{ tf_attribute_type . | lowerCase }}planmodifier.UseStateForUnknown(),
{{- end }}
				},
				Validators: []validator.{{ tf_attribute_type . }}{
{{- if and (eq (awx2go_value .) "types.StringValue") (hasKey . "max_length") }}
					stringvalidator.LengthAtMost({{ .max_length }}),
{{- else if and (eq (awx2go_value .) "types.Int64Value") (hasKey . "min_value") (hasKey . "max_value") }}
					int64validator.Between({{ .min_value }}, {{ .max_value }}),
{{- else if eq .type "choice" }}
					stringvalidator.OneOf({{ awx_type_choice_data .choices }}...),
{{- end }}
                },
            },
{{- end }}
{{- end }}
        // Data only elements
{{- range $key := .PropertyGetKeys }}
{{- if not (hasKey $.PropertyPostData $key) }}
{{- with (index $.PropertyGetData $key) }}
            "{{ $key | lowerCase }}": schema.{{ tf_attribute_type . }}Attribute{
{{- if eq (tf_attribute_type .) "List" }}
				ElementType: types.{{ camelCase .element_type }}Type,
{{- end }}
                Description: {{ escape_quotes (default .help_text "") }},
                Required:    false,
                Optional:    false,
                Computed:    true,
                Sensitive:   {{ .sensitive }},
		        PlanModifiers: []planmodifier.{{ tf_attribute_type . }} {
                    {{ tf_attribute_type . | lowerCase }}planmodifier.UseStateForUnknown(),
				},
{{- if eq .type "choice" }}
				Validators: []validator.{{ tf_attribute_type . }}{
					stringvalidator.OneOf({{ awx_type_choice_data .choices }}...),
				},
{{- end }}
            },
{{- end }}
{{- end }}
{{- end }}
        },
    }
}

{{ if not $.Config.NoId }}
func (o *{{ .Name | lowerCamelCase }}Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
{{- if eq (awx2go_value (index $.PropertyGetData $.Config.IdKey)) "types.Int64Value" }}
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the {{ .Name }}.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("{{ $.Config.IdKey }}"), id)...)
{{- else if eq (awx2go_value (index $.PropertyGetData $.Config.IdKey "type")) "types.StringValue" }}
	resource.ImportStatePassthroughID(ctx, path.Root("{{ $.Config.IdKey }}"), request, response)
{{- end }}
}
{{- end }}

func (o *{{ .Name | lowerCamelCase }}Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
    var err error
	var plan, state {{ .Name | lowerCamelCase }}TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for {{ .Name }}
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[{{.Name}}/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
{{- range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
	bodyRequest.{{ property_case $key $.Config }} = plan.{{ property_case $key $.Config }}.{{ tf2go_primitive_value . }}()
{{- end }}
{{- end }}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, {{ if $.Config.NoId }}http.MethodPatch{{ else }}http.MethodPost{{ end }}, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new {{ .Name }} resource in AWX
    var data map[string]any
    if data, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create resource for {{ .Name }} on %s", o.endpoint),
            err.Error(),
        )
        return
    }

    var d diag.Diagnostics
    if d, err = state.updateFromApiData(data); err != nil {
        response.Diagnostics.Append(d...)
        return
    }

{{ range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
	state.{{ property_case $key $.Config }} = {{ awx2go_value . }}(plan.{{ property_case $key $.Config }}.{{ tf2go_primitive_value . }}())
{{- end }}
{{- end }}

{{ if $.Config.PreStateSetHookFunction }}
    if err = {{ $.Config.PreStateSetHookFunction }}(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on {{ .Name }}",
			err.Error(),
		)
		return
    }
{{ end }}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *{{ .Name | lowerCamelCase }}Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state {{ .Name | lowerCamelCase }}TerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
{{- if $.Config.PreStateSetHookFunction }}
	var orig = state.Clone()
{{- end }}

	// Creates a new request for {{ .Name }}
	var r *http.Request
{{- if $.Config.NoId }}
	var endpoint = p.Clean(o.endpoint) + "/"
{{- else }}
	var id = state.{{ camelCase $.Config.IdKey }}.{{ tf2go_primitive_value (index $.PropertyGetData $.Config.IdKey) }}()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
{{- end }}
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for {{ .Name }} from AWX
    var data map[string]any
    if data, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to read resource for {{ .Name }} on %s", o.endpoint),
            err.Error(),
        )
        return
    }

    var d diag.Diagnostics
    if d, err = state.updateFromApiData(data); err != nil {
        response.Diagnostics.Append(d...)
        return
    }

{{ if $.Config.PreStateSetHookFunction }}
    if err = {{ $.Config.PreStateSetHookFunction }}(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on {{ .Name }}",
			err.Error(),
		)
		return
    }
{{ end }}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *{{ .Name | lowerCamelCase }}Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
    var err error
	var plan, state {{ .Name | lowerCamelCase }}TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for {{ .Name }}
	var r *http.Request
{{- if $.Config.NoId }}
	var endpoint = p.Clean(o.endpoint) + "/"
{{- else }}
	var id = plan.{{ camelCase $.Config.IdKey }}.{{ tf2go_primitive_value (index $.PropertyGetData $.Config.IdKey) }}()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
{{- end }}
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[{{.Name}}/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
{{- range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
	bodyRequest.{{ property_case $key $.Config }} = plan.{{ property_case $key $.Config }}.{{ tf2go_primitive_value . }}()
{{- end }}
{{- end }}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new {{ .Name }} resource in AWX
    var data map[string]any
    if data, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to update resource for {{ .Name }} on %s", o.endpoint),
            err.Error(),
        )
        return
    }

    var d diag.Diagnostics
    if d, err = state.updateFromApiData(data); err != nil {
        response.Diagnostics.Append(d...)
        return
    }

{{ range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
	state.{{ property_case $key $.Config }} = {{ awx2go_value . }}(plan.{{ property_case $key $.Config }}.{{ tf2go_primitive_value . }}())
{{- end }}
{{- end }}

{{ if $.Config.PreStateSetHookFunction }}
    if err = {{ $.Config.PreStateSetHookFunction }}(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on {{ .Name }}",
			err.Error(),
		)
		return
    }
{{ end }}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *{{ .Name | lowerCamelCase }}Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
{{- if $.Config.Undeletable }}
{{- else }}
	var err error

	// Retrieve values from state
	var state {{ .Name | lowerCamelCase }}TerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for {{ .Name }}
	var r *http.Request
	var id = state.{{ camelCase $.Config.IdKey }}
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.{{ tf2go_primitive_value (index $.PropertyGetData $.Config.IdKey) }}())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

    // Delete existing {{ .Name }}
    if _, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to delete resource for {{ .Name }} on %s", o.endpoint),
            err.Error(),
        )
        return
    }
{{- end }}
}
