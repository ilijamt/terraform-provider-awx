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
func (o {{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
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
func (o {{ .Name | lowerCamelCase }}TerraformModel) BodyRequest() (req {{ .Name | lowerCamelCase }}BodyRequestModel) {
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
    // Decode "{{ $key }}"
    if val, ok := data.({{ awx2go_value_cast . }}); ok {
{{- if eq (awx2go_value .) "types.Int64Value" }}
        v, err := val.Int64()
        if err != nil {
            d.AddError(
                fmt.Sprintf("failed to convert %v to int64", val),
                err.Error(),
            )
            return d, err
        }
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}(v)
    } else if val, ok := data.(int64); ok {
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}(val)
{{- else if eq (awx2go_value .) "types.Float64Value" }}
        v, err := val.Float64()
        if err != nil {
            d.AddError(
                fmt.Sprintf("failed to convert %v to float64", val),
                err.Error(),
            )
            return d, err
        }
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}(v)
    } else if val, ok := data.(float64); ok {
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}(val)
{{- else if and (eq (awx2go_value .) "types.StringValue") (eq .type "json") }}
		o.{{ property_case $key $.Config }} = types.StringValue(helpers.TrimString({{ default .trim_space false }}, {{ default .trim_new_line false }}, val))
	} else if val, ok := data.(map[string]any); ok {
	    var v []byte
		if v, err = json.Marshal(val); err != nil {
		    d.AddError(
		        fmt.Sprintf("failed to decode map"),
		        err.Error(),
		    )
		    return
		}
		o.{{ property_case $key $.Config }} = types.StringValue(helpers.TrimString({{ default .trim_space false }}, {{ default .trim_new_line false }}, string(v)))
{{- else if eq (awx2go_value_cast .) "types.List" }}
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}
    } else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
   			list = append(list, types.StringValue(helpers.TrimString({{ default .trim_space false }}, {{ default .trim_new_line false }}, v.(string))))
		}
		o.{{ property_case $key $.Config }} = types.ListValueMust(types.StringType, list)
    } else if data == nil {
		o.{{ property_case $key $.Config }} = types.ListValueMust(types.StringType, []attr.Value{})
	{{- else if eq (awx2go_type .) "types.Map" }}
        types.MapValueMust(types.StringType, val.Elements())
    } else if val, ok := data.(map[string]any); ok {
		var obj map[string]attr.Value
		for k, v := range val {
			obj[k] = types.StringValue(helpers.TrimString({{ default .trim_space false }}, {{ default .trim_new_line false }}, v.(string))))
		}
		o.{{ property_case $key $.Config }} = types.MapValueMust(types.StringType, obj)
{{- else }}
{{- if and (eq (awx2go_value .) "types.StringValue") }}
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}(helpers.TrimString({{ default .trim_space false }}, {{ default .trim_new_line false }}, val))
	} else if val, ok := data.(json.Number); ok {
		o.{{ property_case $key $.Config }} = types.StringValue(val.String())
{{- else }}
        o.{{ property_case $key $.Config }} = {{ awx2go_value . }}(val)
{{- end }}
{{- end }}
    } else {
{{- if eq (awx2go_value .) "types.Int64Value" }}
        o.{{ property_case $key $.Config }} = types.Int64Null()
{{- else if eq (awx2go_value .) "types.Float64Value" }}
        o.{{ property_case $key $.Config }} = types.Float64Null()
{{- else if eq (awx2go_value .) "types.StringValue" }}
        o.{{ property_case $key $.Config }} = types.StringNull()
{{- else if eq (awx2go_value .) "types.BoolValue" }}
        o.{{ property_case $key $.Config }} = types.BoolNull()
{{- else if or (eq (awx2go_value_cast .) "types.List") (eq (awx2go_value_cast .) "types.Map") }}
        err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
        d.AddError(
            fmt.Sprintf("failed to decode value of type %T for {{ awx2go_value_cast . }}", data),
            err.Error(),
        )
        return d, err
{{- end }}
    }
	return d, nil
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
{{ block "tf_model" . }}{{ end }}