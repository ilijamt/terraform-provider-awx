package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewJobTemplateAssociateDisassociateInstanceGroupResource returns the JobTemplate ↔ InstanceGroup association resource.
func NewJobTemplateAssociateDisassociateInstanceGroupResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "job_template_associate_instance_group",
		Endpoint:      "/api/v2/job_templates/%d/instance_groups/",
		ParentName:    "JobTemplate",
		ParentIDAttr:  "job_template_id",
		ChildName:     "InstanceGroup",
		ChildIDAttr:   "instance_group_id",
		AssociateType: "",
		Deprecated:    true,
	})
}
