package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewTeamObjectRolesDataSource returns the Team object_roles data source.
func NewTeamObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"team_object_roles",
		"/api/v2/teams/%d/object_roles/",
		"Team",
		true,
	)
}
