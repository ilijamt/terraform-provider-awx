package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewUserAssociateDisassociateRoleResource returns the User ↔ Role association resource.
func NewUserAssociateDisassociateRoleResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "user_associate_role",
		Endpoint:      "/api/v2/users/%d/roles/",
		ParentName:    "User",
		ParentIDAttr:  "user_id",
		ChildName:     "Role",
		ChildIDAttr:   "role_id",
		AssociateType: "",
		Deprecated:    false,
	})
}
