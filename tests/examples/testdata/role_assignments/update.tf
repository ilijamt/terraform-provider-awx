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

# Built-in role definitions ship with AWX, so they are looked up rather than
# created. object_id is a string because the role can point at any content type.
data "awx_role_definition" "inventory" {
  name = "Inventory Use"
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
