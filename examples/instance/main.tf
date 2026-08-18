terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

# Registers an execution node in the receptor mesh. AWX creates the record
# without contacting the host, so the node stays unavailable until receptor runs
# there. hostname and node_type are rejected on update, so a change to either
# replaces the resource.
resource "awx_instance" "execution" {
  hostname      = "execution-1.example.com"
  node_type     = "execution"
  listener_port = 27199
  enabled       = true
}

# A hop node forwards traffic and runs nothing. Peering it at the execution
# node's receptor address builds the mesh edge; AWX creates that address when
# listener_port is set.
resource "awx_instance" "hop" {
  hostname  = "hop-1.example.com"
  node_type = "hop"
}

data "awx_instance" "execution" {
  hostname = awx_instance.execution.hostname
}

output "execution_capacity" {
  value = data.awx_instance.execution.capacity
}
