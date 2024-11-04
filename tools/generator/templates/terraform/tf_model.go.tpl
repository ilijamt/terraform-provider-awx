package {{ .PackageName }}

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// {{ .Name | lowerCamelCase }}TerraformModel maps the schema for {{ .Name }} when using Data Source
type {{ .Name | lowerCamelCase }}TerraformModel struct {
{{- range $key, $value := .ReadProperties }}
    // {{ $value.Generated.PropertyName }} {{ escape_quotes (or $value.Description "") }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.AwxGoType }} `tfsdk:"{{ $key | lowerCase }}" json:"{{ $key }}"`
{{- end }}
{{- range $key, $value := .WriteProperties }}
{{- if $value.IsWriteOnly }}
    // {{ $value.Generated.PropertyName }} {{ escape_quotes (or $value.Description "") }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.AwxGoType }} `tfsdk:"{{ $key | lowerCase }}" json:"{{ $key }}"`
{{- end }}
{{- end }}
}

// Clone the object
func (o *{{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
    return {{ .Name | lowerCamelCase }}TerraformModel{
    {{- range $key, $value := .ReadProperties }}
        {{ $value.Generated.PropertyName }}: o.{{ $value.Generated.PropertyName }},
    {{- end }}
    {{- range $key, $value := .WriteProperties }}
    {{- if $value.IsWriteOnly }}
        {{ $value.Generated.PropertyName }}: o.{{ $value.Generated.PropertyName }},
    {{- end }}
    {{- end }}
    }
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}TerraformModel) BodyRequest() (req {{ .Name | lowerCamelCase }}BodyRequestModel) {
{{- range $key, $value := .WriteProperties }}
{{- if not $value.IsWriteOnly }}
{{- if eq $value.Generated.AwxGoType "types.List" }}
    req.{{ $value.Generated.PropertyName }} = []string{}
    for _, val := range o.{{ $value.Generated.PropertyName }}.Elements() {
        if _, ok := val.(types.String); ok {
            req.{{ $value.Generated.PropertyName }} = append(req.{{ $value.Generated.PropertyName }}, val.(types.String).ValueString())
        } else {
            req.{{ $value.Generated.PropertyName }} = append(req.{{ $value.Generated.PropertyName }}, val.String())
        }
    }
{{- else }}
    req.{{ $value.Generated.PropertyName }} = {{ $value.Generated.ModelBodyRequestValue }}
{{- end }}
{{- end }}
{{- end }}
    return
}

{{ range $key, $value := .ReadProperties }}
func (o *{{ $.Name | lowerCamelCase }}TerraformModel) set{{ $key | setPropertyCase }}(data any) (_ diag.Diagnostics, _ error) {
{{- if eq $value.Generated.AwxGoValue "types.Int64Value" }}
    return helpers.AttrValueSetInt64(&o.{{ $value.Generated.PropertyName }}, data)
{{- else if eq $value.Generated.AwxGoValue "types.Float64Value" }}
    return helpers.AttrValueSetFloat64(&o.{{ $value.Generated.PropertyName }}, data)
{{- else if eq $value.Generated.AwxGoValue "types.BoolValue" }}
    return helpers.AttrValueSetBool(&o.{{ $value.Generated.PropertyName }}, data)
{{- else if eq $value.Generated.AwxGoValue "types.ListValueMust(types.StringType, val.Elements())" }}
    return helpers.AttrValueSetListString(&o.{{ $value.Generated.PropertyName }}, data, {{ or .Trim false }})
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "json") }}
    return helpers.AttrValueSetJsonString(&o.{{ $value.Generated.PropertyName }}, data, {{ or .Trim false }})
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "json-yaml") }}
    return helpers.AttrValueSetJsonYamlString(&o.{{ $value.Generated.PropertyName }}, data, {{ or .Trim false }})
{{- else if eq $value.Generated.AwxGoValue "types.StringValue" }}
    return helpers.AttrValueSetString(&o.{{ $value.Generated.PropertyName }}, data, {{ or .Trim false }})
{{- end }}
}
{{ end }}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
    if data == nil {
        return diags, fmt.Errorf("no data passed")
    }
{{- range $key, $value := .ReadProperties }}
    if dg, _ := o.set{{ $key | setPropertyCase }}(data["{{ $key }}"]); dg.HasError() {
        diags.Append(dg...)
    }
{{- end }}
    return diags, nil
}

// {{ .Name | lowerCamelCase }}BodyRequestModel maps the schema for {{ .Name }} for creating and updating the data
type {{ .Name | lowerCamelCase }}BodyRequestModel struct {
{{- range $key, $value := .WriteProperties }}
    // {{ $value.Generated.PropertyName }} {{ escape_quotes (or $value.Description "") }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.BodyRequestModelType }} `json:"{{ $key }}{{ if and (not $value.IsRequired) (not (eq $value.Generated.BodyRequestModelType "bool")) }},omitempty{{ end }}"`
{{- end }}
}

{{ if .HasObjectRoles }}
type {{ .Name | lowerCamelCase }}ObjectRolesModel struct {
    ID     types.Int64 `tfsdk:"id"`
	Roles  types.Map   `tfsdk:"roles"`
}
{{- end }}
