resource "awx_role_user_assignment" "reader_inventory" {
  role_definition = data.awx_role_definition.inventory.id
  user            = awx_user.reader.id
  object_id       = awx_inventory.roles.id
}
