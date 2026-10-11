resource "awx_workflow_job_template_node_associate_success_node" "build_to_test" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  success_node_id               = awx_workflow_job_template_node.test.id
}
