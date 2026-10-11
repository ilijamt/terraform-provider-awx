data "awx_instance_group_object_roles" "instance_group_groups" {
  id = data.awx_instance_group.instance_group.id
}
