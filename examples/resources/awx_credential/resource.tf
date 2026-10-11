resource "awx_credential" "vault_default_lookup" {
  name            = "Vault"
  credential_type = data.awx_credential_type.hvsl.id
  organization    = awx_organization.credentials.id
  description     = "We can use this to lookup tokens in vault from AWX"
  inputs = jsonencode({
    "url"               = "https://vault.example.com"
    "token"             = "hvs.CAESIPcZog9etceag4580KjP5Ki_W1lykqgkRv-Tx63fp4LGh4KHGhgFORms0Y1BLakJ5OFlsdmw3aGl6dFY"
    "api_version"       = "v2"
    "default_auth_path" = "approle"
  })
}
