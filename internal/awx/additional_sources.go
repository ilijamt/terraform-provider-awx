package awx

import "github.com/hashicorp/terraform-plugin-framework/resource"

// additionalResources holds resources that cannot be generated because their
// API shape does not fit GenericResource, which assumes one endpoint for all of
// CRUD. Appended to the generated list by Resources() in gen_sources.go.
func additionalResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewWorkflowJobTemplateNodeApprovalResource,
	}
}
