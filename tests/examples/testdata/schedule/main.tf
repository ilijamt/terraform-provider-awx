resource "awx_organization" "schedule" {
  name = "Schedule"
}

resource "awx_inventory" "schedule" {
  name         = "Example schedule inventory"
  organization = awx_organization.schedule.id
}

resource "awx_inventory_source" "schedule" {
  name      = "Example schedule inventory source"
  inventory = awx_inventory.schedule.id
  source    = "ec2"
}

# Inventory sources and system job templates do not prompt for diff_mode, so
# leaving it unset keeps it out of the request rather than sending false.
resource "awx_schedule" "inventory_refresh" {
  name                 = "Refresh example inventory every 6 hours"
  description          = "Reload the EC2 inventory on a fixed interval"
  rrule                = "DTSTART:20250101T000000Z RRULE:FREQ=HOURLY;INTERVAL=6"
  unified_job_template = awx_inventory_source.schedule.id
  enabled              = true
}
