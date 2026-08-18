resource "awx_instance" "execution" {
  hostname            = "execution-1.example.com"
  node_type           = "execution"
  listener_port       = 27199
  enabled             = false
  capacity_adjustment = 0
}

resource "awx_instance" "hop" {
  hostname  = "hop-1.example.com"
  node_type = "hop"
}

data "awx_instance" "execution" {
  hostname = awx_instance.execution.hostname
}
