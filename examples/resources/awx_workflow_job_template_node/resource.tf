resource "awx_workflow_job_template_node" "build" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_job_template.build.id
  identifier            = "build"
}
