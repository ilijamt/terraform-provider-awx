package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewOrganizationObjectRolesDataSource returns the Organization object_roles data source.
func NewOrganizationObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"organization_object_roles",
		"/api/v2/organizations/%d/object_roles/",
		"Organization",
		true,
	)
}
