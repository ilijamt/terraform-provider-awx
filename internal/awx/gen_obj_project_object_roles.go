package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewProjectObjectRolesDataSource returns the Project object_roles data source.
func NewProjectObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"project_object_roles",
		"/api/v2/projects/%d/object_roles/",
		"Project",
		true,
	)
}
