package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewWorkflowJobTemplateNodeAssociateDisassociateAlwaysNodeResource returns the WorkflowJobTemplateNode ↔ AlwaysNode association resource.
func NewWorkflowJobTemplateNodeAssociateDisassociateAlwaysNodeResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "workflow_job_template_node_associate_always_node",
		Endpoint:      "/api/v2/workflow_job_template_nodes/%d/always_nodes/",
		ParentName:    "WorkflowJobTemplateNode",
		ParentIDAttr:  "workflow_job_template_node_id",
		ChildName:     "AlwaysNode",
		ChildIDAttr:   "always_node_id",
		AssociateType: "",
		Deprecated:    false,
	})
}
