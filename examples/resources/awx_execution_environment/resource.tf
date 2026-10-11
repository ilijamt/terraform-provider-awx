resource "awx_execution_environment" "test" {
  name  = "Default"
  image = "quay.io/ansible/awx-ee:latest"
}
