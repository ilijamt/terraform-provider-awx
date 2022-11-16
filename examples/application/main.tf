terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "application" {
  name = "Application"
}

resource "awx_application" "application" {
  organization             = awx_organization.application.id
  authorization_grant_type = "authorization-code"
  name                     = "Application"
  client_type              = "confidential"
  redirect_uris            = "https://localhost"
}

data "awx_application" "application_by_id" {
  id = awx_application.application.id
}

data "awx_application" "application_by_name_org" {
  name         = awx_application.application.name
  organization = awx_application.application.organization
}