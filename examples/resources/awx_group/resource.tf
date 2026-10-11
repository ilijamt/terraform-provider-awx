resource "awx_group" "example_group_vars" {
  name      = "group-vars"
  inventory = awx_inventory.inventory.id
  variables = jsonencode({
    "grp1" = "val1"
    "grp2" = "val2"
  })
}
