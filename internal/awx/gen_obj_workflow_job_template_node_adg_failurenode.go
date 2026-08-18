package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// NewWorkflowJobTemplateNodeAssociateDisassociateFailureNodeResource returns the WorkflowJobTemplateNode ↔ FailureNode association resource.
func NewWorkflowJobTemplateNodeAssociateDisassociateFailureNodeResource() resource.Resource {
	return framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:      "workflow_job_template_node_associate_failure_node",
		Endpoint:      "/api/v2/workflow_job_template_nodes/%d/failure_nodes/",
		ParentName:    "WorkflowJobTemplateNode",
		ParentIDAttr:  "workflow_job_template_node_id",
		ChildName:     "FailureNode",
		ChildIDAttr:   "failure_node_id",
		AssociateType: "",
		Deprecated:    false,
	})
}
