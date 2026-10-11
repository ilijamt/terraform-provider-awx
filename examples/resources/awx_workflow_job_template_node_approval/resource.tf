resource "awx_workflow_job_template_node_approval" "demo_approval" {
  workflow_job_template_node_id = awx_workflow_job_template_node.demo_approval.id
  name                          = "Demo Approval"
  description                   = "Pauses the demo workflow until someone approves"
  timeout                       = 3600
}
