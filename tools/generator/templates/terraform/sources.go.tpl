package {{ .PackageName }}

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
    ApiVersion string = "{{ $.ApiVersion }}"
)

// DataSources is a helper function to return all defined data sources
func DataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
{{- range $org := .DataSources }}
		{{ $org }},
{{- end }}
	}
}

// Resources is a helper function to return all defined resources
func Resources() []func() resource.Resource {
	return []func() resource.Resource{
{{- range $org := .Resources }}
		{{ $org }},
{{- end }}
	}
}
