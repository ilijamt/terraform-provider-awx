terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

data "awx_organization" "default" {
  name                = "Default"
}

resource "awx_inventory" "inventory" {
  name         = "Example inventory"
  organization = data.awx_organization.default.id
  variables = jsonencode({
    "ansible_connection"         = "local"
    "ansible_python_interpreter" = "{{ ansible_playbook_python }}"
    "ansible_debug"              = true
  })
}
