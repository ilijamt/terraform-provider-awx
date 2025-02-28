# Deprecated

{{- if .Resources }}
## Terraform resources
{{- range $item := .Resources }}
- {{ $item }}
{{- end }}
{{- end }}

{{- if .DataSources }}
## Terraform data sources
{{- range $item := .DataSources }}
- {{ $item }}
{{- end }}
{{- end }}

{{- if .Properties }}
## Resource properties
{{- range $item := .Properties }}
{{- if or $item.ReadProperties $item.WriteProperties }}
### {{ $item.Resource }}

{{- if $item.ReadProperties }}
#### Read properties
{{- range $prop := $item.ReadProperties }}
- {{ $prop }}
{{- end }}
{{- end }}

{{- if $item.WriteProperties }}
#### Write properties
{{- range $prop := $item.WriteProperties }}
- {{ $prop }}
{{- end }}
{{- end }}

{{- end }}
{{- end }}
{{- end }}