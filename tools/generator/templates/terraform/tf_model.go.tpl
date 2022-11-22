{{ define "tf_model" }}
// {{ .Name | lowerCamelCase }}TerraformModel maps the schema for {{ .Name }} when using Data Source
type {{ .Name | lowerCamelCase }}TerraformModel struct {
{{- range $key := .PropertyGetKeys }}
{{- with (index $.PropertyGetData $key) }}
    // {{ property_case $key $.Config }} {{ escape_quotes (default .help_text "") }}
    {{ property_case $key $.Config }} {{ awx2go_type . }} `tfsdk:"{{ $key | lowerCase }}" json:"{{ $key }}"`
{{- end }}
{{- end }}
{{- range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
    // {{ property_case $key $.Config }} {{ escape_quotes (default .help_text "") }}
    {{ property_case $key $.Config }} {{ awx2go_type . }} `tfsdk:"{{ $key | lowerCase }}" json:"{{ $key }}"`
{{- end }}
{{- end }}
}

// Clone the object
func (o *{{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
    return {{ .Name | lowerCamelCase }}TerraformModel{
    {{- range $key := .PropertyGetKeys }}
    {{- with (index $.PropertyGetData $key) }}
        {{ property_case $key $.Config }}: o.{{ property_case $key $.Config }},
    {{- end }}
    {{- end }}
    {{- range $key := .PropertyWriteOnlyKeys }}
    {{- with (index $.PropertyWriteOnlyData $key) }}
        {{ property_case $key $.Config }}: o.{{ property_case $key $.Config }},
    {{- end }}
    {{- end }}
    }
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}TerraformModel) BodyRequest() (req {{ .Name | lowerCamelCase }}BodyRequestModel) {
{{- range $key := .PropertyPostKeys }}
{{- with (index $.PropertyPostData $key) }}
{{- if eq (awx2go_type .) "types.List" }}
    req.{{ property_case $key $.Config }} = []string{}
    for _, val := range o.{{ property_case $key $.Config }}.Elements() {
        if _, ok := val.(types.String); ok {
            req.{{ property_case $key $.Config }} = append(req.{{ property_case $key $.Config }}, val.(types.String).ValueString())
        } else {
            req.{{ property_case $key $.Config }} = append(req.{{ property_case $key $.Config }}, val.String())
        }
    }
{{- else }}
    req.{{ property_case $key $.Config }} = {{ if eq .type "json" }}json.RawMessage(o.{{ property_case $key $.Config }}.{{ tf2go_primitive_value . }}()){{ else }}o.{{ property_case $key $.Config }}.{{ tf2go_primitive_value . }}(){{ end }}
{{- end }}
{{- end }}
{{- end }}
    return
}

{{ range $key := .PropertyGetKeys }}
{{- with (index $.PropertyGetData $key) }}
func (o *{{ $.Name | lowerCamelCase }}TerraformModel) set{{ $key | setPropertyCase }}(data any) (d diag.Diagnostics, err error) {
{{- if eq (awx2go_value .) "types.Int64Value" }}
    return helpers.AttrValueSetInt64(&o.{{ property_case $key $.Config }}, data)
{{- else if eq (awx2go_value .) "types.Float64Value" }}
    return helpers.AttrValueSetFloat64(&o.{{ property_case $key $.Config }}, data)
{{- else if eq (awx2go_value .) "types.BoolValue" }}
    return helpers.AttrValueSetBool(&o.{{ property_case $key $.Config }}, data)
{{- else if eq (awx2go_value .) "types.ListValueMust(types.StringType, val.Elements())" }}
    return helpers.AttrValueSetListString(&o.{{ property_case $key $.Config }}, data, {{ default .trim false }})
{{- else if and (eq (awx2go_value .) "types.StringValue") (eq .type "json") }}
    return helpers.AttrValueSetJsonString(&o.{{ property_case $key $.Config }}, data, {{ default .trim false }})
{{- else if eq (awx2go_value .) "types.StringValue" }}
    return helpers.AttrValueSetString(&o.{{ property_case $key $.Config }}, data, {{ default .trim false }})
{{- end }}
}
{{- end }}
{{ end }}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
    if data == nil {
        return diags, fmt.Errorf("no data passed")
    }
{{- range $key := .PropertyGetKeys }}
{{- with (index $.PropertyGetData $key) }}
    if dg, _ := o.set{{ $key | setPropertyCase }}(data["{{ $key }}"]); dg.HasError() {
        diags.Append(dg...)
    }
{{- end }}
{{- end }}
    return diags, nil
}

// {{ .Name | lowerCamelCase }}BodyRequestModel maps the schema for {{ .Name }} for creating and updating the data
type {{ .Name | lowerCamelCase }}BodyRequestModel struct {
{{- range $key := .PropertyPostKeys }}
{{- with (index $.PropertyPostData $key) }}
    // {{ property_case $key $.Config }} {{ escape_quotes (default .help_text "") }}
    {{ property_case $key $.Config }} {{ if eq .type "json" }}json.RawMessage{{ else }}{{ awx2go_primitive_type . }}{{end}} `json:"{{ $key }}{{if and (not .required) (not (eq (awx2go_primitive_type .) "bool"))}},omitempty{{end}}"`
{{- end }}
{{- end }}
{{- range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
    // {{ property_case $key $.Config }} {{ escape_quotes (default .help_text "") }}
    {{ property_case $key $.Config }} {{ if eq .type "json" }}json.RawMessage{{ else }}{{ awx2go_primitive_type . }}{{end}} `json:"{{ $key }},omitempty"`
{{- end }}
{{- end }}
}

{{ if $.Config.HasObjectRoles }}
type {{ .Name | lowerCamelCase }}ObjectRolesModel struct {
    ID     types.Int64 `tfsdk:"id"`
	Roles  types.Map   `tfsdk:"roles"`
}
{{- end }}
{{ end }}
