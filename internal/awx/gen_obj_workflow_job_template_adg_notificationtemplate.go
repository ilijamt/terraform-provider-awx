package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource returns the WorkflowJobTemplate ↔ NotificationTemplate association resource.
func NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "workflow_job_template_associate_notification_template",
		Endpoint:      "/api/v2/workflow_job_templates/%d/notification_templates_%s/",
		ParentName:    "WorkflowJobTemplate",
		ParentIDAttr:  "workflow_job_template_id",
		ChildName:     "NotificationTemplate",
		ChildIDAttr:   "notification_template_id",
		AssociateType: "notification_job_workflow_template",
		Deprecated:    true,
	})
}
