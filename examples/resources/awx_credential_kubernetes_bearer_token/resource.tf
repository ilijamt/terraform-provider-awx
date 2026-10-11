resource "awx_credential_kubernetes_bearer_token" "example" {
  name         = "OpenShift or Kubernetes API Bearer Token"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  bearer_token = var.secrets["kubernetes_bearer_token_bearer_token"]
  verify_ssl   = true
}
