data "awx_label" "test" {
  name         = "test-label"
  organization = awx_organization.test.id
  depends_on   = [awx_label.test]
}
