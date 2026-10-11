resource "awx_workflow_job_template_node_associate_failure_node" "build_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  failure_node_id               = awx_workflow_job_template_node.notify.id
}
