package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewTeamAssociateDisassociateRoleResource returns the Team ↔ Role association resource.
func NewTeamAssociateDisassociateRoleResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "team_associate_role",
		Endpoint:      "/api/v2/teams/%d/roles/",
		ParentName:    "Team",
		ParentIDAttr:  "team_id",
		ChildName:     "Role",
		ChildIDAttr:   "role_id",
		AssociateType: "",
		Deprecated:    true,
	})
}
