resource "awx_inventory_source" "schedule" {
  name      = "Example schedule inventory source"
  inventory = awx_inventory.schedule.id
  source    = "ec2"
}
