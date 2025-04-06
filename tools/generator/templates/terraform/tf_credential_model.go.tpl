package {{ .PackageName }}

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

var (
    _ resource.Updater = &{{ $.Name | lowerCamelCase }}CredentialTerraformModel{}
    _ resource.Cloner[{{ $.Name | lowerCamelCase }}CredentialTerraformModel] = &{{ $.Name | lowerCamelCase }}CredentialTerraformModel{}
    _ resource.Body = &{{ .Name | lowerCamelCase }}CredentialBodyRequestModel{}
)

// {{ .Name | lowerCamelCase }}CredentialBodyRequestModel maps the schema for Credential {{ .Name }} for creating and updating the data
type {{ .Name | lowerCamelCase }}CredentialBodyRequestModel struct {
{{- range $key, $value := .Fields }}
    // {{ $value.Generated.Name }} {{ escape_quotes (or $value.HelpText $value.Label) }}
    {{ $value.Generated.Name }} {{ $value.Generated.GoType }} `json:"{{ $value.Id }}{{ if not $value.Generated.Required }},omitempty{{ end }}"`
{{- end }}
}

func (o {{ .Name | lowerCamelCase }}CredentialBodyRequestModel) MarshalJSON() ([]byte, error) { return json.Marshal(o) }

// {{ .Name | lowerCamelCase }}CredentialTerraformModel maps the schema for Credential {{ .Name }}
type {{ .Name | lowerCamelCase }}CredentialTerraformModel struct {
	// ID "Database ID for this credential."
	ID types.Int64 `tfsdk:"id" json:"id"`
{{- range $key, $value := .Fields }}
    // {{ $value.Generated.Name }} {{ escape_quotes (or $value.HelpText $value.Label) }}
    {{ $value.Generated.Name }} {{ $value.Generated.Type }}  `tfsdk:"{{ $value.Id | lowerCase }}" json:"{{ $value.Id }}"`
{{- end }}
}

// Clone the object
func (o *{{ .Name | lowerCamelCase }}CredentialTerraformModel) Clone() {{ .Name | lowerCamelCase }}CredentialTerraformModel {
    return {{ .Name | lowerCamelCase }}CredentialTerraformModel{
        ID: o.ID,
    {{- range $key, $value := .Fields }}
        {{ $value.Generated.Name }}: o.{{ $value.Generated.Name }},
    {{- end }}
    }
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for {{ .Name }}
func (o *{{ .Name | lowerCamelCase }}CredentialTerraformModel) BodyRequest() (req {{ .Name | lowerCamelCase }}CredentialBodyRequestModel) {
{{- range $key, $value := .Fields }}
    req.{{ $value.Generated.Name }} = o.{{ $value.Generated.Name }}.{{ $value.Generated.TerraformValue }}()
{{- end }}
    return
}

{{ range $key, $value := .Fields }}
func (o *{{ $.Name | lowerCamelCase }}CredentialTerraformModel) set{{ $value.Generated.Name }}(data any) (_ diag.Diagnostics, _ error) {
{{- if eq $value.Generated.Type "types.String" }}
    return helpers.AttrValueSetString(&o.{{ $value.Generated.Name }}, data, false)
{{- else if eq $value.Generated.Type "types.Bool" }}
    return helpers.AttrValueSetBool(&o.{{ $value.Generated.Name }}, data)
{{- else if eq $value.Generated.Type "types.Int64" }}
    return helpers.AttrValueSetInt64(&o.{{ $value.Generated.Name }}, data)
{{- else }}
    return nil, fmt.Errorf("invalid type: $value.Generated.Type")
{{- end }}
}
{{ end }}

func (o *{{ .Name | lowerCamelCase }}CredentialTerraformModel) UpdateWithApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
    if data == nil {
        return diags, fmt.Errorf("no data passed")
    }

    for field, setter := range map[string]func(any) (diag.Diagnostics, error) {
           "id":             func(v any) (diag.Diagnostics, error) { return helpers.AttrValueSetInt64(&o.ID, v) },
{{- range $key, $value := .Fields }}
           "{{ $value.Id }}":           o.set{{ $value.Generated.Name }},
{{- end }}
       } {
         d, _ := setter(data[field])
         diags.Append(d...)
    }

    return diags, nil
}
