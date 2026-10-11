resource "awx_job_template_associate_credential" "demo_machine" {
  credential_id   = awx_credential.demo_credential.id
  job_template_id = awx_job_template.demo_job_template.id
}
