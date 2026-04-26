package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewJobTemplateAssociateDisassociateNotificationTemplateResource returns the JobTemplate ↔ NotificationTemplate association resource.
func NewJobTemplateAssociateDisassociateNotificationTemplateResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "job_template_associate_notification_template",
		Endpoint:      "/api/v2/job_templates/%d/notification_templates_%s/",
		ParentName:    "JobTemplate",
		ParentIDAttr:  "job_template_id",
		ChildName:     "NotificationTemplate",
		ChildIDAttr:   "notification_template_id",
		AssociateType: "notification_job_template",
		Deprecated:    true,
	})
}
