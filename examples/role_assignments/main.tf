terraform {
  required_providers {
    awx = {
      source = "ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "roles" {
  name = "Role assignments"
}

resource "awx_inventory" "roles" {
  name         = "Role assignment inventory"
  organization = awx_organization.roles.id
}

resource "awx_user" "reader" {
  username = "role-reader"
  email    = "role-reader@example.com"
  password = "role-reader-secret"
}

resource "awx_team" "readers" {
  name         = "Role readers"
  organization = awx_organization.roles.id
}

resource "awx_role_definition" "inventory_reader" {
  name         = "Inventory Reader"
  description  = "Read-only access to an inventory"
  content_type = "awx.inventory"
  permissions = [
    "awx.view_inventory",
    "awx.adhoc_inventory",
  ]
}

# object_id is a string because a role can point at any content type.
data "awx_role_definition" "inventory" {
  name = "Inventory Admin"
}

resource "awx_role_user_assignment" "reader_inventory" {
  role_definition = data.awx_role_definition.inventory.id
  user            = awx_user.reader.id
  object_id       = awx_inventory.roles.id
}

resource "awx_role_team_assignment" "readers_inventory" {
  role_definition = data.awx_role_definition.inventory.id
  team            = awx_team.readers.id
  object_id       = awx_inventory.roles.id
}
