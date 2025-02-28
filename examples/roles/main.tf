terraform {
  required_providers {
    awx = {
      source = "ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "org" {
  name = "roles"
}

resource "awx_team" "team" {
  name         = "roles"
  organization = awx_organization.org.id
}

resource "awx_user" "user" {
  username = "roles"
  password = "test"
}
