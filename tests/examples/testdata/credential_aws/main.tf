resource "awx_organization" "org" {
  name = "CredentialAws"
}

resource "awx_credential_aws" "smoke" {
  name         = "AWS smoke"
  description  = "Created by TestIntegration_CredentialAws"
  organization = awx_organization.org.id
  username     = "AKIAEXAMPLE"
  password     = "secret-key"
}

data "awx_credential_aws" "smoke_by_name" {
  name = awx_credential_aws.smoke.name
}
