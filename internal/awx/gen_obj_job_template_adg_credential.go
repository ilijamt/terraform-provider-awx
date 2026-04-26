package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewJobTemplateAssociateDisassociateCredentialResource returns the JobTemplate ↔ Credential association resource.
func NewJobTemplateAssociateDisassociateCredentialResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "job_template_associate_credential",
		Endpoint:      "/api/v2/job_templates/%d/credentials/",
		ParentName:    "JobTemplate",
		ParentIDAttr:  "job_template_id",
		ChildName:     "Credential",
		ChildIDAttr:   "credential_id",
		AssociateType: "",
		Deprecated:    true,
	})
}
