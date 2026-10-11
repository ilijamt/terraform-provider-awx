resource "awx_instance" "execution" {
  hostname      = "execution-1.example.com"
  node_type     = "execution"
  listener_port = 27199
  enabled       = true
}
