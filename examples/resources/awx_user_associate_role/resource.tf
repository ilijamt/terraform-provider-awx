resource "awx_user_associate_role" "user_team_admin" {
  user_id = awx_user.user_role.id
  role_id = data.awx_team_object_roles.team_role.roles["Admin"]
}
