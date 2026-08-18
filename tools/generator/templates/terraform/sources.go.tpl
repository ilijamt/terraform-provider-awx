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
		New{{ $org }}DataSource,
{{- end }}
	}
}

// Resources is a helper function to return all defined resources, generated
// ones plus the hand-written entries in additionalResources.
func Resources() []func() resource.Resource {
	return append([]func() resource.Resource{
{{- range $org := .Resources }}
		New{{ $org }}Resource,
{{- end }}
	}, additionalResources()...)
}
