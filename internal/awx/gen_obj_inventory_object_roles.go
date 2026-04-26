package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewInventoryObjectRolesDataSource returns the Inventory object_roles data source.
func NewInventoryObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"inventory_object_roles",
		"/api/v2/inventories/%d/object_roles/",
		"Inventory",
		true,
	)
}
