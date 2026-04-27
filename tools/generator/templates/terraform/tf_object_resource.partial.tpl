{{- /*
attrSchema renders a single Required/Optional/Computed schema attribute. Bool
fields are omitted when false (Go zero); PlanModifiers/Validators lists are
omitted when empty. Used for both regular and write-only attributes.
*/ -}}
{{- define "attrSchema" -}}
{{- $key := .Key }}{{ $value := .Value -}}
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
{{- if $value.IsSensitive }}
	Sensitive:   true,
{{- end }}
{{- if $value.IsRequired }}
	Required:    true,
{{- else }}
	Optional:    true,
{{- end }}
{{- if $value.IsComputed }}
	Computed:    true,
{{- end }}
{{- if $value.HasDefaultValue }}
	Default:     {{ $value.DefaultValue }},
{{- end }}
{{- if not $value.IsRequired }}
	PlanModifiers: []planmodifier.{{ $value.Generated.AttributeType }}{
		{{ $value.Generated.AttributeType | lowerCase }}planmodifier.UseStateForUnknown(),
	},
{{- end }}
{{- if and (eq $value.Generated.AwxGoValue "types.StringValue") (hasKey $value.ValidatorData "max_length") }}
	Validators: []validator.{{ $value.Generated.AttributeType }}{
		stringvalidator.LengthAtMost({{ $value.ValidatorData.max_length }}),
{{- range $value.Constraints }}
		// {{ .Id }}
		{{ .Constraint }}({{ range $k := .Fields }}path.MatchRoot("{{ $k }}"), {{ end }}),
{{- end }}
	},
{{- else if and (eq $value.Generated.AwxGoValue "types.Int64Value") (hasKey $value.ValidatorData "min_value") (hasKey $value.ValidatorData "max_value") }}
	Validators: []validator.{{ $value.Generated.AttributeType }}{
		int64validator.Between({{ format_number $value.ValidatorData.min_value }}, {{ format_number $value.ValidatorData.max_value }}),
{{- range $value.Constraints }}
		// {{ .Id }}
		{{ .Constraint }}({{ range $k := .Fields }}path.MatchRoot("{{ $k }}"), {{ end }}),
{{- end }}
	},
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq $value.Type "choice") }}
	Validators: []validator.{{ $value.Generated.AttributeType }}{
		stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
			{{ $item | quote }},
{{- end }}
		),
{{- range $value.Constraints }}
		// {{ .Id }}
		{{ .Constraint }}({{ range $k := .Fields }}path.MatchRoot("{{ $k }}"), {{ end }}),
{{- end }}
	},
{{- else if and (eq $value.Generated.AwxGoValue "types.ListValueMust(types.StringType, val.Elements())") (eq $value.Type "list") (or $value.Generated.ValidationAvailableChoiceData $value.Validators) }}
	Validators: []validator.{{ $value.Generated.AttributeType }}{
{{- range $item := $value.Validators }}
		{{ $item }},
{{- end }}
		listvalidator.ValueStringsAre(stringvalidator.OneOf(
{{- range $item := $value.Generated.ValidationAvailableChoiceData }}
			{{ $item | quote }},
{{- end }}
		)),
{{- range $value.Constraints }}
		// {{ .Id }}
		{{ .Constraint }}({{ range $k := .Fields }}path.MatchRoot("{{ $k }}"), {{ end }}),
{{- end }}
	},
{{- else if $value.Constraints }}
	Validators: []validator.{{ $value.Generated.AttributeType }}{
{{- range $value.Constraints }}
		// {{ .Id }}
		{{ .Constraint }}({{ range $k := .Fields }}path.MatchRoot("{{ $k }}"), {{ end }}),
{{- end }}
	},
{{- end }}
},
{{- end -}}
{{- /*
resource_section emits the Terraform resource binding (schema + GenericResource
config) for a regular generated resource. Caller supplies ModelConfig.ToMap()
data; rendered inside tf_object.go.tpl which provides the package + imports.
*/ -}}
{{- define "resource_section" -}}
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
{{- range $key, $value := .WriteProperties }}
					{{ template "attrSchema" (dict "Key" $key "Value" $value) }}
{{- end }}
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
						Description: {{ escape_quotes (or .Description .Label) }},
{{- if .IsSensitive }}
						Sensitive:   true,
{{- end }}
						Computed: true,
						PlanModifiers: []planmodifier.{{ $value.Generated.AttributeType }}{
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
					if m.{{ camelCase $.IdKey }}.IsNull() || m.{{ camelCase $.IdKey }}.IsUnknown() {
						return ""
					}
{{- if eq $.IdProperty.Generated.AwxGoValue "types.Int64Value" }}
					if m.{{ camelCase $.IdKey }}.ValueInt64() == 0 {
						return ""
					}
{{- end }}
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
{{- end -}}
