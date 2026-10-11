resource "awx_credential_input_source" "gitlab" {
  input_field_name = "ssh_key_data"
  metadata = jsonencode({
    "auth_path"      = ""
    "secret_key"     = "private_key"
    "secret_path"    = "secrets/keys/awx-gitlab"
    "secret_backend" = "secrets"
    "secret_version" = ""
  })
  target_credential = awx_credential.gitlab.id
  source_credential = awx_credential.vault_default_lookup.id
}
