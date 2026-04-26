package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewJobTemplateObjectRolesDataSource returns the JobTemplate object_roles data source.
func NewJobTemplateObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"job_template_object_roles",
		"/api/v2/job_templates/%d/object_roles/",
		"JobTemplate",
		true,
	)
}
