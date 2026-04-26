package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewInstanceGroupObjectRolesDataSource returns the InstanceGroup object_roles data source.
func NewInstanceGroupObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"instance_group_object_roles",
		"/api/v2/instance_groups/%d/object_roles/",
		"InstanceGroup",
		true,
	)
}
