terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {
  hostname = "http://awx.local"
  username = "admin"
  password = "admin"
}

resource "awx_credential_net" "test" {
  name        = "Credential for Net"
  username    = "username"
  password    = "password"
  description = "description"
}
