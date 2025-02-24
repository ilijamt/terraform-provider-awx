terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

data "awx_execution_environment" "latest" {
  name = "AWX EE (latest)"
}

resource "awx_organization" "inventory" {
  name                = "Inventory"
  default_environment = data.awx_execution_environment.latest.id
}

resource "awx_inventory" "inventory" {
  name         = "Example inventory"
  organization = awx_organization.inventory.id
  variables = jsonencode({
    "ansible_connection"         = "local"
    "ansible_python_interpreter" = "{{ ansible_playbook_python }}"
    "ansible_debug"              = true
  })
  depends_on = [awx_organization.inventory]
}

resource "awx_host" "example_host_vars" {
  name      = "vars.example.com"
  inventory = awx_inventory.inventory.id
  variables = jsonencode({
    "host1" = "val1"
    "host2" = "val2"
  })
}

resource "awx_host" "example_host_no_vars" {
  name      = "no-vars.example.com"
  inventory = awx_inventory.inventory.id
  variables = jsonencode({
    "host1" = "val1"
    "host2" = "val2"
  })
}

resource "awx_group" "example_group_vars" {
  name      = "group-vars"
  inventory = awx_inventory.inventory.id
  variables = jsonencode({
    "grp1" = "val1"
    "grp2" = "val2"
  })
}

resource "awx_group" "example_group_no_vars" {
  name      = "group-no-vars"
  inventory = awx_inventory.inventory.id
}

resource "awx_host_associate_group" "example_association_vars" {
  group_id = awx_group.example_group_vars.id
  host_id  = awx_host.example_host_vars.id
}

resource "awx_host_associate_group" "example_association_no_vars" {
  group_id = awx_group.example_group_no_vars.id
  host_id  = awx_host.example_host_no_vars.id
}
