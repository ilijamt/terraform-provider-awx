package awx

import "github.com/hashicorp/terraform-plugin-framework/resource"

// additionalResources holds resources whose API shape does not fit
// GenericResource, which assumes one endpoint for all of CRUD.
func additionalResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewWorkflowJobTemplateNodeApprovalResource,
	}
}
