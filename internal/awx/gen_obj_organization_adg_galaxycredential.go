package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewOrganizationAssociateDisassociateGalaxyCredentialResource returns the Organization ↔ GalaxyCredential association resource.
func NewOrganizationAssociateDisassociateGalaxyCredentialResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "organization_associate_galaxy_credential",
		Endpoint:      "/api/v2/organizations/%d/galaxy_credentials/",
		ParentName:    "Organization",
		ParentIDAttr:  "organization_id",
		ChildName:     "GalaxyCredential",
		ChildIDAttr:   "galaxy_credential_id",
		AssociateType: "",
		Deprecated:    true,
	})
}
