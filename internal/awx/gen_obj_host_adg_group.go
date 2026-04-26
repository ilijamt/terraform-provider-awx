package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewHostAssociateDisassociateGroupResource returns the Host ↔ Group association resource.
func NewHostAssociateDisassociateGroupResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "host_associate_group",
		Endpoint:      "/api/v2/hosts/%d/groups/",
		ParentName:    "Host",
		ParentIDAttr:  "host_id",
		ChildName:     "Group",
		ChildIDAttr:   "group_id",
		AssociateType: "",
		Deprecated:    true,
	})
}
