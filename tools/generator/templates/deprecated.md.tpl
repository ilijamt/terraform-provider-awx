# Deprecated

## Terraform resources
{{- range $item := .Resources }}
- {{ $item }}
{{- end }}

## Resource properties
{{ range $item := .Properties }}
{{ if or $item.ReadProperties $item.WriteProperties }}
### {{ $item.Resource }}

#### Read properties
{{ range $prop := $item.ReadProperties }}
- {{ $prop }}
{{ end }}

#### Write properties
{{ range $prop := $item.WriteProperties }}
- {{ $prop }}
{{ end }}

{{ end }}
{{- end }}