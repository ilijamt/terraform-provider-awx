resource "awx_project" "demo_project" {
  name         = "Job Demo Project"
  organization = awx_organization.test.id
  scm_url      = "https://github.com/ansible/ansible-tower-samples"
  scm_type     = "git"
  scm_clean    = false
}
