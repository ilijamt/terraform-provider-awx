terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
    local = {
      source  = "hashicorp/local"
      version = "~> 2.5"
    }
  }
}

# Provider config is read entirely from the environment:
#   TOWER_HOST / AWX_HOST
#   TOWER_USERNAME / AWX_USERNAME
#   TOWER_PASSWORD / AWX_PASSWORD
provider "awx" {}

resource "awx_token" "bootstrap" {
  description = "terraform-provider-awx bootstrap token (test recording)"
  scope       = "write"
}

resource "local_sensitive_file" "token" {
  filename        = "${path.module}/../.bootstrap-token"
  content         = awx_token.bootstrap.token
  file_permission = "0600"
}

output "token_path" {
  value       = local_sensitive_file.token.filename
  description = "Path to the file containing the bootstrap token"
}
