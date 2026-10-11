resource "awx_team_associate_role" "team_role_admin_org" {
  team_id = awx_team.team_role.id
  role_id = data.awx_organization_object_roles.org_role.roles["Execute"]
}
