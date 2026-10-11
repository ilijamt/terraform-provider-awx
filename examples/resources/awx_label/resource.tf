resource "awx_label" "test" {
  name         = "test-label"
  organization = awx_organization.test.id
}
