package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewConstructedInventoriesObjectRolesDataSource returns the ConstructedInventories object_roles data source.
func NewConstructedInventoriesObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"constructed_inventories_object_roles",
		"/api/v2/constructed_inventories/%d/object_roles/",
		"ConstructedInventories",
		true,
	)
}
