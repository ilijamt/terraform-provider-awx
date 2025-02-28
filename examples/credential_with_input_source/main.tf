terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

data "awx_credential_type" "hvsl" {
  name = "HashiCorp Vault Secret Lookup"
}

data "awx_credential_type" "scm" {
  name = "Source Control"
}

resource "awx_organization" "credentials" {
  name = "Credentials (with input)"
}

resource "awx_credential" "vault_default_lookup" {
  name            = "Vault"
  credential_type = data.awx_credential_type.hvsl.id
  organization    = awx_organization.credentials.id
  description     = "We can use this to lookup tokens in vault from AWX"
  inputs = jsonencode({
    "url"               = "https://vault.example.com"
    "token"             = "hvs.CAESIPcZog9etceag4580KjP5Ki_W1lykqgkRv-Tx63fp4LGh4KHGhgFORms0Y1BLakJ5OFlsdmw3aGl6dFY"
    "api_version"       = "v2"
    "default_auth_path" = "approle"
  })
}

resource "awx_credential" "gitlab" {
  name            = "Gitlab"
  credential_type = data.awx_credential_type.scm.id
  organization    = awx_organization.credentials.id
  description     = "Gitlab all access"
  inputs = jsonencode({
    "username" = "git"
  })
}

resource "awx_credential_input_source" "gitlab" {
  input_field_name = "ssh_key_data"
  metadata = jsonencode({
    "auth_path"      = ""
    "secret_key"     = "private_key"
    "secret_path"    = "secrets/keys/awx-gitlab"
    "secret_backend" = "secrets"
    "secret_version" = ""
  })
  target_credential = awx_credential.gitlab.id
  source_credential = awx_credential.vault_default_lookup.id
}

data "awx_credential" "gitlab" {
  name = awx_credential.gitlab.name
}
