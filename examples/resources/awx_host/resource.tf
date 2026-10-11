resource "awx_host" "example_host_vars" {
  name      = "vars.example.com"
  inventory = awx_inventory.inventory.id
  variables = jsonencode({
    "host1" = "val1"
    "host2" = "val2"
  })
}
