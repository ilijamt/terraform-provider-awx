package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewCredentialObjectRolesDataSource returns the Credential object_roles data source.
func NewCredentialObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"credential_object_roles",
		"/api/v2/credentials/%d/object_roles/",
		"Credential",
		true,
	)
}
