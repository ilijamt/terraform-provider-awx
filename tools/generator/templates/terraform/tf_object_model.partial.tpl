{{- /*
model_section emits the typed terraform model, BodyRequest builder, body
request struct, and UpdateFromApiData reader for a regular (non-credential)
generated resource. Caller provides ModelConfig.ToMap() data and renders this
inside tf_object.go.tpl, which carries the package+imports.
*/ -}}
{{- define "model_section" -}}
type {{ .Name | lowerCamelCase }}TerraformModel struct {
{{- range $key, $value := .ReadProperties }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.AwxGoType }} `tfsdk:"{{ $key | lowerCase }}" json:"{{ $key }}"`
{{- end }}
{{- range $key, $value := .WriteProperties }}
{{- if $value.IsWriteOnly }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.AwxGoType }} `tfsdk:"{{ $key | lowerCase }}" json:"{{ $key }}"`
{{- end }}
{{- end }}
{{- if .WaitLifecycle }}
    // {{ .WaitLifecycle.WaitAttribute | camelCase }} is a Terraform-only toggle, not synced to the AWX API.
    {{ .WaitLifecycle.WaitAttribute | camelCase }} types.Bool `tfsdk:"{{ .WaitLifecycle.WaitAttribute }}" json:"-"`
    Timeouts timeouts.Value `tfsdk:"timeouts" json:"-"`
{{- end }}
}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
    return *o
}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) BodyRequest() *{{ .Name | lowerCamelCase }}BodyRequestModel {
    var req {{ .Name | lowerCamelCase }}BodyRequestModel
{{- range $key, $value := .WriteProperties }}
{{- if not $value.IsWriteOnly }}
{{- if eq $value.Generated.AwxGoType "types.List" }}
    req.{{ $value.Generated.PropertyName }} = helpers.ListAsStringSlice(o.{{ $value.Generated.PropertyName }}, {{ or .Trim false }})
{{- else }}
    req.{{ $value.Generated.PropertyName }} = {{ $value.Generated.ModelBodyRequestValue }}
{{- end }}
{{- end }}
{{- end }}
    return &req
}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
    diags = make(diag.Diagnostics, 0)
    if data == nil {
        return diags, fmt.Errorf("no data passed")
    }
    collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
{{- range $key, $value := .ReadProperties }}
{{- if eq $value.Generated.AwxGoValue "types.Int64Value" }}
    collect(helpers.AttrValueSetInt64(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"]))
{{- else if eq $value.Generated.AwxGoValue "types.Float64Value" }}
    collect(helpers.AttrValueSetFloat64(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"]))
{{- else if eq $value.Generated.AwxGoValue "types.BoolValue" }}
    collect(helpers.AttrValueSetBool(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"]))
{{- else if eq $value.Generated.AwxGoValue "types.ListValueMust(types.StringType, val.Elements())" }}
    collect(helpers.AttrValueSetListString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }}))
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "json") }}
    collect(helpers.AttrValueSetJsonString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }}))
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "json-yaml") }}
    collect(helpers.AttrValueSetJsonYamlString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }}))
{{- else if eq $value.Generated.AwxGoValue "types.StringValue" }}
    collect(helpers.AttrValueSetString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }}))
{{- end }}
{{- end }}
    return diags, nil
}

type {{ .Name | lowerCamelCase }}BodyRequestModel struct {
{{- range $key, $value := .WriteProperties }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.BodyRequestModelType }} `json:"{{ $key }}{{ if and (not $value.IsRequired) (not (eq $value.Generated.BodyRequestModelType "bool")) $value.OmitEmpty }},omitempty{{ end }}"`
{{- end }}
}
{{- end -}}
