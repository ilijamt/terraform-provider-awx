resource "awx_inventory" "inventory" {
  name         = "Example inventory"
  organization = awx_organization.inventory.id
  variables = jsonencode({
    "ansible_connection"         = "local"
    "ansible_python_interpreter" = "{{ ansible_playbook_python }}"
    "ansible_debug"              = true
  })
  depends_on = [awx_organization.inventory]
}
