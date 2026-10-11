resource "awx_credential_aws_secretsmanager_credential" "example" {
  name           = "AWS Secrets Manager lookup"
  organization   = awx_organization.typed.id
  aws_access_key = "AKIAEXAMPLE"
  aws_secret_key = var.secrets["aws_secretsmanager_credential_aws_secret_key"]
}
