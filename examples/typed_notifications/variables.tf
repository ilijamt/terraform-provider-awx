# Populate from a tfvars file your VCS ignores, TF_VAR_secrets, or a secret
# manager data source.
variable "secrets" {
  type      = map(string)
  sensitive = true
}
