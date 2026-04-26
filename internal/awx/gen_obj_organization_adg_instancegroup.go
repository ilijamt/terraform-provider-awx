package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewOrganizationAssociateDisassociateInstanceGroupResource returns the Organization ↔ InstanceGroup association resource.
func NewOrganizationAssociateDisassociateInstanceGroupResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "organization_associate_instance_group",
		Endpoint:      "/api/v2/organizations/%d/instance_groups/",
		ParentName:    "Organization",
		ParentIDAttr:  "organization_id",
		ChildName:     "InstanceGroup",
		ChildIDAttr:   "instance_group_id",
		AssociateType: "",
		Deprecated:    true,
	})
}
