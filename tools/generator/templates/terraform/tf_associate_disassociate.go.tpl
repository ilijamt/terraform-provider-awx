package {{ .PackageName }}

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// New{{ .Name }}AssociateDisassociate{{ .Type }}Resource returns the {{ .Name }} ↔ {{ .Type }} association resource.
func New{{ .Name }}AssociateDisassociate{{ .Type }}Resource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "{{ .Name | snakeCase }}_associate_{{ .Type | snakeCase }}",
		Endpoint:      "{{ .Endpoint }}",
		ParentName:    "{{ .Name }}",
		ParentIDAttr:  "{{ .Name | snakeCase }}_id",
		ChildName:     "{{ .Type }}",
		ChildIDAttr:   "{{ .Type | snakeCase }}_id",
		AssociateType: "{{ .AssociateType }}",
		Deprecated:    {{ if .Deprecated }}true{{ else }}false{{ end }},
	})
}
