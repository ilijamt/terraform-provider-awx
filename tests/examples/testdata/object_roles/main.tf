data "awx_credential_type" "machine" {
  name = "Machine"
}

resource "awx_organization" "org" {
  name = "Object Roles"
}

resource "awx_team" "team" {
  name         = "Team"
  organization = awx_organization.org.id
}

resource "awx_user" "user" {
  username = "object_roles_user"
  password = "object_roles_user"
}

resource "awx_credential" "machine" {
  name            = "Demo Machine"
  credential_type = data.awx_credential_type.machine.id
  organization    = awx_organization.org.id
}

data "awx_organization_object_roles" "org" {
  id = awx_organization.org.id
}

data "awx_team_object_roles" "team" {
  id = awx_team.team.id
}

data "awx_credential_object_roles" "machine" {
  id = awx_credential.machine.id
}

data "awx_instance_group" "default" {
  name = "default"
}

data "awx_instance_group_object_roles" "default" {
  id = data.awx_instance_group.default.id
}

resource "awx_team_associate_role" "team_org_execute" {
  team_id = awx_team.team.id
  role_id = data.awx_organization_object_roles.org.roles["Execute"]
}

resource "awx_team_associate_role" "team_org_inventory_admin" {
  team_id = awx_team.team.id
  role_id = data.awx_organization_object_roles.org.roles["Inventory Admin"]
}

resource "awx_team_associate_role" "team_credential_use" {
  team_id = awx_team.team.id
  role_id = data.awx_credential_object_roles.machine.roles["Use"]
}

resource "awx_team_associate_role" "team_instance_group_use" {
  team_id = awx_team.team.id
  role_id = data.awx_instance_group_object_roles.default.roles["Use"]
}

resource "awx_user_associate_role" "user_team_admin" {
  user_id = awx_user.user.id
  role_id = data.awx_team_object_roles.team.roles["Admin"]
}
