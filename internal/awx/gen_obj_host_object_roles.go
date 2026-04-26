package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewHostObjectRolesDataSource returns the Host object_roles data source.
func NewHostObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"host_object_roles",
		"/api/v2/hosts/%d/object_roles/",
		"Host",
		true,
	)
}
