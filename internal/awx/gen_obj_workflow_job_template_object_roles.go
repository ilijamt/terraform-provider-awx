package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewWorkflowJobTemplateObjectRolesDataSource returns the WorkflowJobTemplate object_roles data source.
func NewWorkflowJobTemplateObjectRolesDataSource() datasource.DataSource {
	return framework.NewObjectRolesDataSource(
		"workflow_job_template_object_roles",
		"/api/v2/workflow_job_templates/%d/object_roles/",
		"WorkflowJobTemplate",
		true,
	)
}
