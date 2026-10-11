data "awx_workflow_job_template_node" "build" {
  id = awx_workflow_job_template_node.build.id

  depends_on = [
    awx_workflow_job_template_node_associate_success_node.build_to_test,
    awx_workflow_job_template_node_associate_failure_node.build_to_notify,
  ]
}
