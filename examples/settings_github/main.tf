terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

locals {
  team_map = {
    "My Team" : {
      "organization" : "Test Org",
      "users" : ["/^[^@]+?@test\\.example\\.com$/"],
      "remove" : true
    },
    "Other Team" : {
      "organization" : "Test Org 2",
      "users" : ["/^[^@]+?@test\\.example\\.com$/"],
      "remove" : false
    }
  }
  org_map = {
    "Default" : {
      "users" : true
    },
    "Test Org" : {
      "admins" : ["admin@example.com"],
      "users" : true
    },
    "Test Org 2" : {
      "admins" : ["admin@example.com", "/^tower-[^@]+?@.*$/i"],
      "users" : "/^[^@].*?@example\\.com$/"
    }

  }
}

resource "awx_settings_auth_github" "default" {
  social_auth_github_key              = "key"
  social_auth_github_organization_map = jsonencode(local.org_map)
  social_auth_github_secret           = "secret"
  social_auth_github_team_map         = jsonencode(local.team_map)
}

resource "awx_settings_auth_github_org" "default" {
  social_auth_github_org_key              = "key"
  social_auth_github_org_organization_map = jsonencode(local.org_map)
  social_auth_github_org_secret           = "secret"
  social_auth_github_org_team_map         = jsonencode(local.team_map)
}

resource "awx_settings_auth_github_team" "default" {
  social_auth_github_team_id               = "example-id"
  social_auth_github_team_key              = "key"
  social_auth_github_team_secret           = "secret"
  social_auth_github_team_organization_map = jsonencode(local.org_map)
  social_auth_github_team_team_map         = jsonencode(local.team_map)
}

resource "awx_settings_auth_github_enterprise" "default" {
  social_auth_github_enterprise_url              = "https://example.com"
  social_auth_github_enterprise_api_url          = "https://api.example.com"
  social_auth_github_enterprise_key              = "key"
  social_auth_github_enterprise_organization_map = jsonencode(local.org_map)
  social_auth_github_enterprise_secret           = "secret"
  social_auth_github_enterprise_team_map         = jsonencode(local.team_map)
}

resource "awx_settings_auth_github_enterprise_org" "default" {
  social_auth_github_enterprise_org_url              = "https://example.com"
  social_auth_github_enterprise_org_api_url          = "https://api.example.com"
  social_auth_github_enterprise_org_name             = "Example Org name"
  social_auth_github_enterprise_org_key              = "key"
  social_auth_github_enterprise_org_organization_map = jsonencode(local.org_map)
  social_auth_github_enterprise_org_secret           = "secret"
  social_auth_github_enterprise_org_team_map         = jsonencode(local.team_map)
}

resource "awx_settings_auth_github_enterprise_team" "default" {
  social_auth_github_enterprise_team_id               = "example-id"
  social_auth_github_enterprise_team_url              = "https://example.com"
  social_auth_github_enterprise_team_api_url          = "https://api.example.com"
  social_auth_github_enterprise_team_key              = "key"
  social_auth_github_enterprise_team_secret           = "secret"
  social_auth_github_enterprise_team_organization_map = jsonencode(local.org_map)
  social_auth_github_enterprise_team_team_map         = jsonencode(local.team_map)
}