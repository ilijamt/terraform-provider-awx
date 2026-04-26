data "awx_credential_type" "machine" {
  name = "Machine"
}

data "awx_credential_type" "hvsl" {
  name = "HashiCorp Vault Secret Lookup"
}

data "awx_credential_type" "scm" {
  name = "Source Control"
}

resource "awx_organization" "org" {
  name = "Credentials"
}

resource "awx_team" "team" {
  name         = "Team"
  organization = awx_organization.org.id
}

resource "awx_user" "user" {
  username = "credentials_user"
  password = "credentials_user"
}

resource "awx_credential" "organization" {
  name            = "Organization"
  credential_type = data.awx_credential_type.machine.id
  organization    = awx_organization.org.id
  description     = "Assigned to Organization (updated)"
  inputs = jsonencode({
    "username" = "test"
    "password" = "password"
  })
}

resource "awx_credential" "team" {
  name            = "Team"
  credential_type = data.awx_credential_type.machine.id
  team            = awx_team.team.id
  description     = "Assigned to Team"
  inputs = jsonencode({
    "username" = "test"
    "password" = "password"
  })
}

resource "awx_credential" "user" {
  name            = "User"
  credential_type = data.awx_credential_type.machine.id
  user            = awx_user.user.id
  description     = "Assigned to User"
  inputs = jsonencode({
    "username" = "test"
    "password" = "password"
  })
}

resource "awx_credential" "vault" {
  name            = "Vault"
  credential_type = data.awx_credential_type.hvsl.id
  organization    = awx_organization.org.id
  description     = "Vault token lookup source for other credentials (updated)"
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
  organization    = awx_organization.org.id
  description     = "Gitlab all access"
  inputs = jsonencode({
    "username" = "git"
  })
}

resource "awx_credential_input_source" "gitlab" {
  input_field_name  = "ssh_key_data"
  target_credential = awx_credential.gitlab.id
  source_credential = awx_credential.vault.id
  metadata = jsonencode({
    "auth_path"      = ""
    "secret_key"     = "private_key"
    "secret_path"    = "secrets/keys/awx-gitlab"
    "secret_backend" = "secrets"
    "secret_version" = ""
  })
}
