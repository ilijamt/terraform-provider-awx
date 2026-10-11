resource "awx_workflow_job_template" "demo_workflow_template" {
  name         = "Demo Workflow Job"
  organization = awx_organization.demo_organization.id
}
