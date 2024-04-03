terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "org_role" {
  name = "org-role"
}

data "awx_organization_object_roles" "org_role" {
  id = awx_organization.org_role.id
}

resource "awx_team" "team_role" {
  name         = "team-role"
  organization = awx_organization.org_role.id
}

data "awx_team_object_roles" "team_role" {
  id = awx_team.team_role.id
}

resource "awx_user" "user_role" {
  username = "user_role"
  password = "test"
}

resource "awx_team_associate_role" "team_role_admin_org" {
  team_id = awx_team.team_role.id
  role_id = data.awx_organization_object_roles.org_role.roles["Execute"]
}

resource "awx_team_associate_role" "team_role_inventory_admin" {
  team_id = awx_team.team_role.id
  role_id = data.awx_organization_object_roles.org_role.roles["Inventory Admin"]
}

resource "awx_user_associate_role" "user_team_admin" {
  user_id = awx_user.user_role.id
  role_id = data.awx_team_object_roles.team_role.roles["Admin"]
}

data "awx_credential_type" "machine" {
  name = "Machine"
}

resource "awx_credential" "demo_credential_machine" {
  name            = "Demo Machine"
  credential_type = data.awx_credential_type.machine.id
  organization    = awx_organization.org_role.id
}

data "awx_credential_object_roles" "demo_credential_machine" {
  id = awx_credential.demo_credential_machine.id
}

resource "awx_team_associate_role" "team_access_demo_credential" {
  team_id = awx_team.team_role.id
  role_id = data.awx_credential_object_roles.demo_credential_machine.roles["Use"]
}


data "awx_instance_group" "instance_group" {
  name = "default"
}

data "awx_instance_group_object_roles" "instance_group_groups" {
  id = data.awx_instance_group.instance_group.id
}

resource "awx_team_associate_role" "team_access_instance_groups" {
  team_id = awx_team.team_role.id
  role_id = data.awx_instance_group_object_roles.instance_group_groups.roles["Use"]
}
