package {{ .PackageName }}

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/models"
	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

var (
    _ resource.Updater = &terraformModel{}
    _ resource.Cloner[terraformModel] = &terraformModel{}
    _ resource.RequestBody = &terraformModel{}
	_ resource.Credential = &terraformModel{}
	_ resource.Id = &terraformModel{}
)

// terraformModel maps the schema for Credential {{ $.TypeName }}
type terraformModel struct {
	// ID "Database ID for this credential."
	ID types.Int64 `tfsdk:"id" json:"id"`
{{- range $key, $value := .Fields }}
    // {{ $value.Generated.Name }} {{ escape_quotes (or $value.HelpText $value.Label) }}
    {{ $value.Generated.Name }} {{ $value.Generated.Type }}  `tfsdk:"{{ $value.Id | lowerCase }}" json:"{{ $value.Id }}"`
{{- end }}

    // internal variables that are required for the request to finish
    // successfully
	userId int64
	credentialTypeId int64
}

func (o *terraformModel) GetId() (string, error) {
	if o.ID.IsNull() || o.ID.IsUnknown() {
		return "", fmt.Errorf("id not set")
	}
	return o.ID.String(), nil
}

func (o *terraformModel) Data() models.Credential {
    var inputs = map[string]any{
{{- range $key, $value := .Fields }}
{{- if and $value.IsInput $value.Generated.Required }}
        "{{ $value.Id }}": o.{{ $value.Generated.Name }}.{{ $value.Generated.TerraformValue }}(),
{{- end }}
{{- end }}
    }

{{- range $key, $value := .Fields }}
{{- if and $value.IsInput (not $value.Generated.Required) }}
    if !o.{{ $value.Generated.Name }}.IsNull() && !o.{{ $value.Generated.Name }}.IsUnknown() {
        inputs["{{ $value.Id }}"] = o.{{ $value.Generated.Name }}.{{ $value.Generated.TerraformValue }}()
    }
{{- end }}
{{- end }}

    return models.Credential{
		CredentialType: o.credentialTypeId,
		Inputs: inputs,
		User: o.userId,
{{- range $key, $value := .Fields }}
{{- if not $value.IsInput }}
        {{ $value.Generated.Name }}: o.{{ $value.Generated.Name }}.{{ $value.Generated.TerraformValue }}{{ if $value.Generated.Pointer }}Pointer{{ end }}(),
{{- end }}
{{- end }}
    }
}


func (o *terraformModel) RequestBody() ([]byte, error) {
    return json.Marshal(o.Data())
}

// Clone the object
func (o *terraformModel) Clone() terraformModel {
    return terraformModel{
        ID: o.ID,
    {{- range $key, $value := .Fields }}
        {{ $value.Generated.Name }}: o.{{ $value.Generated.Name }},
    {{- end }}
    }
}

{{ range $key, $value := .Fields }}
func (o *terraformModel) set{{ $value.Generated.Name }}(data any) (_ diag.Diagnostics, _ error) {
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

func (o *terraformModel) setId(data any) (_ diag.Diagnostics, _ error) {
    return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *terraformModel) UpdateWithApiData(callee resource.Callee, source resource.Source, data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
    if data == nil {
        return diags, fmt.Errorf("data is empty")
    }

	var fieldUTO []helpers.FieldMapping
{{- range $key, $value := .Fields }}
{{- if and ($value.IsUTO) (not $value.Generated.WriteOnly) }}
    o.{{ $value.Generated.Name }} = {{ $value.Generated.Type }}Null()
	if val, ok := data["{{ $value.Id }}"]; ok && val != nil {
		fieldUTO = append(
            fieldUTO,
            helpers.FieldMapping{ APIField:  "{{ $value.Id }}", Setter: o.set{{ $value.Generated.Name }} },
		)
	}
{{- end }}
{{- end }}

    // Set the default items to the values in the API payload
	var fieldMappings = append(
	    []helpers.FieldMapping{
            { APIField: "id", Setter: o.setId },
{{- range $key, $value := .Fields }}
{{- if and (not $value.IsInput) (not $value.IsUTO) (not $value.Generated.WriteOnly) }}
            { APIField:  "{{ $value.Id }}", Setter:   o.set{{ $value.Generated.Name }} },
{{- end }}
{{- end }}
	    },
	    fieldUTO...,
	)

	// We need to process all the inputs that are not a secret
	// if an input is a secret, then the value will be $encrypted$ which is not useful
	// so we skip those fields
	if inputs, ok := data["inputs"].(map[string]any); ok {
        fieldMappings = append(
            fieldMappings,
{{- range $key, $value := .Fields }}
{{- if and $value.IsInput (not $value.Secret) }}
            helpers.FieldMapping{
                APIField:  "{{ $value.Id }}",
                Setter: o.set{{ $value.Generated.Name }},
                Data: inputs,
            },
{{- end }}
{{- end }}
        )
	}

	diags, _ = helpers.ApplyFieldMappings(data, fieldMappings...)
    return diags, nil
}
