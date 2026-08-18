package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewWorkflowJobTemplateNodeAssociateDisassociateSuccessNodeResource returns the WorkflowJobTemplateNode ↔ SuccessNode association resource.
func NewWorkflowJobTemplateNodeAssociateDisassociateSuccessNodeResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "workflow_job_template_node_associate_success_node",
		Endpoint:      "/api/v2/workflow_job_template_nodes/%d/success_nodes/",
		ParentName:    "WorkflowJobTemplateNode",
		ParentIDAttr:  "workflow_job_template_node_id",
		ChildName:     "SuccessNode",
		ChildIDAttr:   "success_node_id",
		AssociateType: "",
		Deprecated:    false,
	})
}
