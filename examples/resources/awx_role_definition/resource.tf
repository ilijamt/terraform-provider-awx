resource "awx_role_definition" "inventory_reader" {
  name         = "Inventory Reader"
  description  = "Read-only access to an inventory"
  content_type = "awx.inventory"
  permissions = [
    "awx.view_inventory",
    "awx.adhoc_inventory",
  ]
}
