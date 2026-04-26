resource "awx_organization" "org" {
  name = "CredentialAws"
}

resource "awx_credential_aws" "smoke" {
  name           = "AWS smoke (updated)"
  description    = "Updated by TestIntegration_CredentialAws"
  organization   = awx_organization.org.id
  username       = "AKIAUPDATED"
  password       = "secret-key"
  security_token = "sts-token-value"
}

data "awx_credential_aws" "smoke_by_name" {
  name = awx_credential_aws.smoke.name
}
