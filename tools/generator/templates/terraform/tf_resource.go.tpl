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
{{- if not .NoId }}
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
	response.TypeName = request.ProviderTypeName + "_{{ .TypeName }}"
}

func (o *{{ .Name | lowerCamelCase }}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
        // Request elements
{{- range $key, $value := .WriteProperties }}
{{- if not $value.IsWriteOnly }}
            "{{ $key | lowerCase }}": schema.{{ $value.Generated.AttributeType }}Attribute{
{{- if eq $value.Generated.AttributeType "List" }}
				ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
                Description: {{ escape_quotes (or $value.Description $value.Label) }},
                Sensitive:   {{ $value.IsSensitive }},
                Required:    {{ $value.IsRequired }},
                Optional:    {{ not $value.IsRequired }},
                Computed:    {{ $value.IsComputed }},
{{- if .HasDefaultValue }}
                Default:     {{ $value.DefaultValue }},
{{- end }}
		        PlanModifiers: []planmodifier.{{ $value.Generated.AttributeType }} {
{{- if not .IsRequired }}
                    {{  $value.Generated.AttributeType | lowerCase }}planmodifier.UseStateForUnknown(),
{{- end }}
                },
				Validators: []validator.{{ $value.Generated.AttributeType }}{
{{- if and (eq $value.Generated.AwxGoValue "types.StringValue") (hasKey $value.ValidatorData "max_length") }}
					stringvalidator.LengthAtMost({{ $value.ValidatorData.max_length }}),
{{- else if and (eq $value.Generated.AwxGoValue "types.Int64Value") (hasKey $value.ValidatorData "min_value") (hasKey $value.ValidatorData "max_value") }}
					int64validator.Between({{ format_number $value.ValidatorData.min_value }}, {{ format_number $value.ValidatorData.max_value }}),
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "choice") }}
					stringvalidator.OneOf([]string{ {{- range $item := $value.Generated.ValidationAvailableChoiceData }}{{ $item | quote }},{{- end }} }...),
{{- end }}
{{- range $value.Constraints }}
                    // {{ .Id }}
                    {{ .Constraint }}(
{{- range $k := .Fields }}
                        path.MatchRoot("{{ $k }}"),
{{- end }}
                    ),
{{- end }}
                },
            },
{{- end }}
{{- end }}
        // Write only elements
{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
            "{{ $key | lowerCase }}": schema.{{ $value.Generated.AttributeType }}Attribute{
{{- if eq $value.Generated.AttributeType "List" }}
				ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
                Description: {{ escape_quotes (or $value.Description $value.Label) }},
                Sensitive:   {{ $value.IsSensitive }},
                Required:    {{ $value.IsRequired }},
                Optional:    {{ not $value.IsRequired }},
                Computed:    {{ $value.IsComputed }},
{{- if .HasDefaultValue }}
                Default:     {{ $value.DefaultValue }},
{{- end }}
		        PlanModifiers: []planmodifier.{{ $value.Generated.AttributeType }} {
{{- if not $value.IsRequired }}
                    {{ $value.Generated.AttributeType | lowerCase }}planmodifier.UseStateForUnknown(),
{{- end }}
				},
				Validators: []validator.{{ $value.Generated.AttributeType }}{
{{- if and (eq $value.Generated.AwxGoValue "types.StringValue") (hasKey $value.ValidatorData "max_length") }}
					stringvalidator.LengthAtMost({{ $value.ValidatorData.max_length }}),
{{- else if and (eq $value.Generated.AwxGoValue "types.Int64Value") (hasKey $value.ValidatorData "min_value") (hasKey $value.ValidatorData "max_value") }}
					int64validator.Between({{ format_number $value.ValidatorData.min_value }}, {{ format_number $value.ValidatorData.max_value }}),
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "choice") }}
					stringvalidator.OneOf([]string{ {{- range $item := $value.Generated.ValidationAvailableChoiceData }}{{ $item | quote }},{{- end }} }...),
{{- end }}
{{- range $value.Constraints }}
                    // {{ .Id }}
                    {{ .Constraint }}(
{{- range $k := .Fields }}
                        path.MatchRoot("{{ $k }}"),
{{- end }}
                    ),
{{- end }}
                },
            },
{{- end }}
{{- end }}
        // Data only elements
{{- range $key, $value := .ReadProperties }}
{{- if not $value.IsInWriteProperty }}
            "{{ $key | lowerCase }}": schema.{{ $value.Generated.AttributeType }}Attribute{
{{- if eq $value.Generated.AttributeType "List" }}
				ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
                Description: {{ escape_quotes (or $value.Description "") }},
                Required:    false,
                Optional:    false,
                Computed:    true,
                Sensitive:   {{ .IsSensitive }},
		        PlanModifiers: []planmodifier.{{ $value.Generated.AttributeType }} {
                    {{ $value.Generated.AttributeType | lowerCase }}planmodifier.UseStateForUnknown(),
				},
{{- if eq .Type "choice" }}
				Validators: []validator.{{ $value.Generated.AttributeType }}{
					stringvalidator.OneOf([]string{ {{- range $item := $value.Generated.ValidationAvailableChoiceData }}{{ $item | quote }},{{- end }} }...),
				},
{{- end }}
            },
{{- end }}
{{- end }}
        },
    }
}

{{ if not .NoId }}
func (o *{{ $.Name | lowerCamelCase }}Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
{{- with $.IdProperty }}
{{- if eq .Generated.AwxGoValue "types.Int64Value" }}
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the {{ $.Name }}.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("{{ $.IdKey }}"), id)...)
{{- else if eq .Generated.AwxGoValue "types.StringValue" }}
	resource.ImportStatePassthroughID(ctx, path.Root("{{ $.IdKey }}"), request, response)
{{- end }}
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
{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
	bodyRequest.{{ $value.Generated.PropertyName }} = plan.{{ $value.Generated.PropertyName }}.{{ $value.Generated.TfGoPrimitiveValue }}()
{{- end }}
{{- end }}
	tflog.Debug(ctx, "[{{.Name}}/Create] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, {{ if .NoId }}http.MethodPatch{{ else }}http.MethodPost{{ end }}, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for create", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new {{ .Name }} resource in AWX
    var data map[string]any
    if data, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create resource for {{ .Name }} on %s", endpoint),
            err.Error(),
        )
        return
    }

    var d diag.Diagnostics
    if d, err = state.updateFromApiData(data); err != nil {
        response.Diagnostics.Append(d...)
        return
    }

{{ range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
	state.{{ $value.Generated.PropertyName }} = {{ $value.Generated.AwxGoValue }}(plan.{{ $value.Generated.PropertyName }}.{{ $value.Generated.TfGoPrimitiveValue }}())
{{- end }}
{{- end }}

{{ if .PreStateSetHookFunction }}
    if err = {{ .PreStateSetHookFunction }}(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeCreate, &plan, &state); err != nil {
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
{{- if $.PreStateSetHookFunction }}
	var orig = state.Clone()
{{- end }}

	// Creates a new request for {{ .Name }}
	var r *http.Request
{{- if $.NoId }}
	var endpoint = p.Clean(o.endpoint) + "/"
{{- else }}
	var id = state.{{ camelCase $.IdKey }}.{{ $.IdProperty.Generated.TfGoPrimitiveValue }}()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
{{- end }}
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for read", endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for {{ .Name }} from AWX
    var data map[string]any
    if data, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to read resource for {{ .Name }} on %s", endpoint),
            err.Error(),
        )
        return
    }

    var d diag.Diagnostics
    if d, err = state.updateFromApiData(data); err != nil {
        response.Diagnostics.Append(d...)
        return
    }

{{ if $.PreStateSetHookFunction }}
    if err = {{ $.PreStateSetHookFunction }}(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeRead, &orig, &state); err != nil {
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
{{- if $.NoId }}
	var endpoint = p.Clean(o.endpoint) + "/"
{{- else }}
	var id = plan.{{ camelCase $.IdKey }}.{{ $.IdProperty.Generated.TfGoPrimitiveValue }}()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
{{- end }}
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[{{.Name}}/Update] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})

{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
	bodyRequest.{{ $value.Generated.PropertyName }} = plan.{{ $value.Generated.PropertyName }}.{{ $value.Generated.TfGoPrimitiveValue }}()
{{- end }}
{{- end }}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for update", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new {{ .Name }} resource in AWX
    var data map[string]any
    if data, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to update resource for {{ .Name }} on %s", endpoint),
            err.Error(),
        )
        return
    }

    var d diag.Diagnostics
    if d, err = state.updateFromApiData(data); err != nil {
        response.Diagnostics.Append(d...)
        return
    }

{{ range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
	state.{{ $value.Generated.PropertyName }} = {{ $value.Generated.AwxGoValue }}(plan.{{ $value.Generated.PropertyName }}.{{ $value.Generated.TfGoPrimitiveValue }}())
{{- end }}
{{- end }}

{{ if $.PreStateSetHookFunction }}
    if err = {{ $.PreStateSetHookFunction }}(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeUpdate, &plan, &state); err != nil {
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
{{- if $.UnDeletable }}
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
	var id = state.{{ camelCase $.IdKey }}
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.{{ $.IdProperty.Generated.TfGoPrimitiveValue }}())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

    // Delete existing {{ .Name }}
    if _, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to delete resource for {{ .Name }} on %s", endpoint),
            err.Error(),
        )
        return
    }
{{- end }}
}
