resource "awx_role_team_assignment" "readers_inventory" {
  role_definition = data.awx_role_definition.inventory.id
  team            = awx_team.readers.id
  object_id       = awx_inventory.roles.id
}
