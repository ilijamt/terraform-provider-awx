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

# Disabling clears next_run and the new rrule moves dtstart, so this step also
# covers the computed values AWX recalculates itself.
resource "awx_schedule" "inventory_refresh" {
  name                 = "Refresh example inventory every 12 hours"
  description          = "Reload the EC2 inventory on a slower interval"
  rrule                = "DTSTART:20250201T000000Z RRULE:FREQ=HOURLY;INTERVAL=12"
  unified_job_template = awx_inventory_source.schedule.id
  enabled              = false
}
