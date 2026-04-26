package {{ .PackageName }}

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
{{- if .WaitLifecycle }}
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
{{- end }}
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
{{- if .WaitLifecycle }}
    // {{ .WaitLifecycle.WaitAttribute | camelCase }} is a Terraform-only toggle (not synced to the AWX API).
    {{ .WaitLifecycle.WaitAttribute | camelCase }} types.Bool `tfsdk:"{{ .WaitLifecycle.WaitAttribute }}" json:"-"`
    // Timeouts holds the user-configured timeouts {} block.
    Timeouts timeouts.Value `tfsdk:"timeouts" json:"-"`
{{- end }}
}

// Clone the object
func (o *{{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
    return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for {{ .Name }}
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
{{- range $key, $value := .ReadProperties }}
    {
{{- if eq $value.Generated.AwxGoValue "types.Int64Value" }}
        dg, _ := helpers.AttrValueSetInt64(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"])
{{- else if eq $value.Generated.AwxGoValue "types.Float64Value" }}
        dg, _ := helpers.AttrValueSetFloat64(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"])
{{- else if eq $value.Generated.AwxGoValue "types.BoolValue" }}
        dg, _ := helpers.AttrValueSetBool(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"])
{{- else if eq $value.Generated.AwxGoValue "types.ListValueMust(types.StringType, val.Elements())" }}
        dg, _ := helpers.AttrValueSetListString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }})
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "json") }}
        dg, _ := helpers.AttrValueSetJsonString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }})
{{- else if and (eq $value.Generated.AwxGoValue "types.StringValue") (eq .Type "json-yaml") }}
        dg, _ := helpers.AttrValueSetJsonYamlString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }})
{{- else if eq $value.Generated.AwxGoValue "types.StringValue" }}
        dg, _ := helpers.AttrValueSetString(&o.{{ $value.Generated.PropertyName }}, data["{{ $key }}"], {{ or .Trim false }})
{{- end }}
        diags.Append(dg...)
    }
{{- end }}
    return diags, nil
}

// {{ .Name | lowerCamelCase }}BodyRequestModel maps the schema for {{ .Name }} for creating and updating the data
type {{ .Name | lowerCamelCase }}BodyRequestModel struct {
{{- range $key, $value := .WriteProperties }}
    // {{ $value.Generated.PropertyName }} {{ escape_quotes (or $value.Description "") }}
    {{ $value.Generated.PropertyName }} {{ $value.Generated.BodyRequestModelType }} `json:"{{ $key }}{{ if and (not $value.IsRequired) (not (eq $value.Generated.BodyRequestModelType "bool")) $value.OmitEmpty }},omitempty{{ end }}"`
{{- end }}
}

