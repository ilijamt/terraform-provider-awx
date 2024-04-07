terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

data "awx_credential_type" "machine" {
  name = "Machine"
}

resource "awx_organization" "credentials" {
  name = "Credentials"
}

resource "awx_team" "team" {
  name         = "Team"
  organization = awx_organization.credentials.id
}

resource "awx_user" "user" {
  username = "user"
  password = "user"
}

resource "awx_credential" "organization" {
  name            = "Organization"
  credential_type = data.awx_credential_type.machine.id
  organization    = awx_organization.credentials.id
  description     = "Assigned to Organization"
  inputs = jsonencode({
    "username": "test",
    "password": "password"
  })
}

resource "awx_credential" "team" {
  name            = "Team"
  credential_type = data.awx_credential_type.machine.id
  team            = awx_team.team.id
  description     = "Assigned to Team"
  inputs = jsonencode({
    "username": "test",
    "password": "password"
  })
}

resource "awx_credential" "user" {
  name            = "User"
  credential_type = data.awx_credential_type.machine.id
  user            = awx_user.user.id
  description     = "Assigned to User"
  inputs = jsonencode({
    "username": "test",
    "password": "password"
  })
}
