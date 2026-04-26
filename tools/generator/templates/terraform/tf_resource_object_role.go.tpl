package {{ .PackageName }}

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// New{{ .Name }}ObjectRolesDataSource returns the {{ .Name }} object_roles data source.
func New{{ .Name }}ObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"{{ $.TypeName }}_object_roles",
		"{{ .Endpoint }}%d/object_roles/",
		"{{ .Name }}",
		{{ if (index .DeprecatedParts "ObjectRoles") }}true{{ else }}false{{ end }},
	)
}
