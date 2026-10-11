terraform {
  required_providers {
    awx = {
      source = "ilijamt/awx"
    }
  }
}

provider "awx" {
  hostname = "https://awx.example.com"
  token    = var.awx_token

  # Or log in with a username and password instead of a token:
  # username = var.awx_username
  # password = var.awx_password
}
