resource "awx_workflow_job_template_node_associate_always_node" "test_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.test.id
  always_node_id                = awx_workflow_job_template_node.notify.id
}
