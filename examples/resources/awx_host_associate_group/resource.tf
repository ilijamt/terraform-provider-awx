resource "awx_host_associate_group" "example_association_vars" {
  group_id = awx_group.example_group_vars.id
  host_id  = awx_host.example_host_vars.id
}
