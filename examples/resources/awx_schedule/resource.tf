resource "awx_schedule" "inventory_refresh" {
  name                 = "Refresh example inventory every 6 hours"
  description          = "Reload the EC2 inventory on a fixed interval"
  rrule                = "DTSTART:20250101T000000Z RRULE:FREQ=HOURLY;INTERVAL=6"
  unified_job_template = awx_inventory_source.schedule.id
  enabled              = true
}
