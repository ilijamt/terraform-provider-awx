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

resource "awx_credential_aws" "test" {
  name           = "Credential for AWS (test)"
  username       = "username"
  password       = "password"
  description    = "Description for our new shiny AWS credential"
  user = 1
}
