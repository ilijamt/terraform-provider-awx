{{- $hasConstraints := false }}
{{- range $key, $value := $.WriteProperties }}
{{- if $value.Constraints }}{{ $hasConstraints = true }}{{ end }}
{{- end }}
package {{ .PackageName }}

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
{{- if $hasConstraints }}
	"github.com/hashicorp/terraform-plugin-framework/path"
{{- end }}
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

    "github.com/ilijamt/terraform-provider-awx/internal/hooks"
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type {{ .Name | lowerCamelCase }}Resource = framework.GenericResource[{{ .Name | lowerCamelCase }}TerraformModel, {{ .Name | lowerCamelCase }}BodyRequestModel, *{{ .Name | lowerCamelCase }}TerraformModel]

// New{{ .Name }}Resource is a helper function to simplify the provider implementation.
func New{{ .Name }}Resource() resource.Resource {
	return &{{ .Name | lowerCamelCase }}Resource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .TypeName }}", Endpoint: "{{ $.Endpoint }}"}},
		Cfg: framework.ResourceCfg[{{ .Name | lowerCamelCase }}TerraformModel, {{ .Name | lowerCamelCase }}BodyRequestModel]{
			Schema: schema.Schema{
{{- if .Deprecated }}
				DeprecationMessage: "This resource has been deprecated and will be removed in a future release.",
{{- end }}
				Attributes: map[string]schema.Attribute{
				// Request elements
{{- range $key, $value := .WriteProperties }}
{{- if not $value.IsWriteOnly }}
					"{{ $key | lowerCase }}": schema.{{ $value.Generated.AttributeType }}Attribute{
{{- if and (eq $value.Generated.AttributeType "List") (eq $value.ElementType "choice") }}
						ElementType: types.ListType{ElemType: types.StringType},
{{- else if eq $value.Generated.AttributeType "List" }}
						ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
{{- if $value.Deprecated }}
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
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
							stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
								{{ $item | quote }},
{{- end }}
							),
{{- else if and (eq $value.Generated.AwxGoValue "types.ListValueMust(types.StringType, val.Elements())") (eq .Type "list") (or $value.Generated.ValidationAvailableChoiceData $value.Validators) }}
{{- range $item := $value.Validators }}
							{{ $item }},
{{- end }}
							listvalidator.ValueStringsAre(stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
								{{ $item | quote }},
{{- end }}
							)),
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
{{- if and (eq $value.Generated.AttributeType "List") (eq $value.ElementType "choice") }}
						ElementType: types.ListType{ElemType: types.StringType},
{{- else if eq $value.Generated.AttributeType "List" }}
						ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
{{- if $value.Deprecated }}
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
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
							stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
								{{ $item | quote }},
{{- end }}
							),
{{- else if and (eq $value.Generated.AwxGoValue "types.ListValueMust(types.StringType, val.Elements())") (eq .Type "list") (or $value.Generated.ValidationAvailableChoiceData $value.Validators) }}
{{- range $item := $value.Validators }}
							{{ $item }},
{{- end }}
							listvalidator.ValueStringsAre(stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
								{{ $item | quote }},
{{- end }}
							)),
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
{{- if and (eq $value.Generated.AttributeType "List") (eq $value.ElementType "choice") }}
						ElementType: types.ListType{ElemType: types.StringType},
{{- else if eq $value.Generated.AttributeType "List" }}
						ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
{{- if $value.Deprecated }}
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
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
							stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
								{{ $item | quote }},
{{- end }}
							),
						},
{{- end }}
					},
{{- end }}
{{- end }}
{{- if .WaitLifecycle }}
				// Terraform-only lifecycle controls
					"{{ .WaitLifecycle.WaitAttribute }}": schema.BoolAttribute{
						Description: {{ escape_quotes .WaitLifecycle.WaitDescription }},
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
{{- end }}
				},
			},
{{- if not .NoId }}
			IDAccessor: func(m *{{ .Name | lowerCamelCase }}TerraformModel) any { return m.{{ camelCase $.IdKey }}.{{ $.IdProperty.Generated.TfGoPrimitiveValue }}() },
{{- if not .NoImport }}
			IDKey: "{{ $.IdKey }}",
{{- if eq $.IdProperty.Generated.AwxGoValue "types.StringValue" }}
			IDIsString: true,
{{- end }}
{{- end }}
{{- end }}
{{- if .NoId }}
			NoId: true,
{{- end }}
{{- if .NoImport }}
			NoImport: true,
{{- end }}
{{- if .UnDeletable }}
			UnDeletable: true,
{{- end }}
{{- if .PreStateSetHookFunction }}
{{- if eq .PreStateSetHookFunction "hooks.RequireResourceStateOrOrig" }}
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *{{ .Name | lowerCamelCase }}TerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
{{- else }}
			Hook: {{ .PreStateSetHookFunction }},
{{- end }}
{{- end }}
{{- $hasWriteOnly := false }}
{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
{{- $hasWriteOnly = true }}
{{- end }}
{{- end }}
{{- if $hasWriteOnly }}
			WriteOnlyPlanToBody: func(plan *{{ .Name | lowerCamelCase }}TerraformModel, body *{{ .Name | lowerCamelCase }}BodyRequestModel) {
{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
				body.{{ $value.Generated.PropertyName }} = plan.{{ $value.Generated.PropertyName }}.{{ $value.Generated.TfGoPrimitiveValue }}()
{{- end }}
{{- end }}
			},
			WriteOnlyPlanToState: func(plan, state *{{ .Name | lowerCamelCase }}TerraformModel) {
{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
				state.{{ $value.Generated.PropertyName }} = {{ $value.Generated.AwxGoValue }}(plan.{{ $value.Generated.PropertyName }}.{{ $value.Generated.TfGoPrimitiveValue }}())
{{- end }}
{{- end }}
			},
{{- end }}
{{- if .WaitLifecycle }}
			EmitTimeouts: true,
			CopyExtraAttributes: func(plan, state *{{ .Name | lowerCamelCase }}TerraformModel) {
				state.{{ .WaitLifecycle.WaitAttribute | camelCase }} = plan.{{ .WaitLifecycle.WaitAttribute | camelCase }}
				state.Timeouts = plan.Timeouts
			},
			WaitLifecycle: &framework.WaitLifecycleCfg[{{ .Name | lowerCamelCase }}TerraformModel]{
				ShouldWait: func(plan *{{ .Name | lowerCamelCase }}TerraformModel) bool {
					return !plan.{{ .WaitLifecycle.WaitAttribute | camelCase }}.IsNull() && plan.{{ .WaitLifecycle.WaitAttribute | camelCase }}.ValueBool()
				},
				EndpointForModel: func(m *{{ .Name | lowerCamelCase }}TerraformModel) string {
					return framework.EndpointWithID("{{ $.Endpoint }}", m.{{ camelCase $.IdKey }}.{{ $.IdProperty.Generated.TfGoPrimitiveValue }}())
				},
				Field:          {{ .WaitLifecycle.StatusField | quote }},
				SuccessValues:  {{ go_string_slice .WaitLifecycle.SuccessValues }},
				FailureValues:  {{ go_string_slice .WaitLifecycle.FailureValues }},
				PollInterval:   {{ go_duration .WaitLifecycle.PollInterval }},
				DefaultTimeout: {{ go_duration .WaitLifecycle.DefaultTimeout }},
				ResolveTimeout: func(ctx context.Context, plan *{{ .Name | lowerCamelCase }}TerraformModel, callee hooks.Callee) (time.Duration, diag.Diagnostics) {
					if callee == hooks.CalleeUpdate {
						return plan.Timeouts.Update(ctx, {{ go_duration .WaitLifecycle.DefaultTimeout }})
					}
					return plan.Timeouts.Create(ctx, {{ go_duration .WaitLifecycle.DefaultTimeout }})
				},
			},
{{- end }}
			ApiVersion: ApiVersion,
			ResourceName: "{{ .Name }}",
		},
	}
}

