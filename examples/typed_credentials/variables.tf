# Populate from a tfvars file your VCS ignores, TF_VAR_secrets, or a secret
# manager data source.
variable "secrets" {
  type      = map(string)
  sensitive = true
}

variable "ssh_private_key" {
  type      = string
  sensitive = true
}

variable "gpg_public_key" {
  type = string
}

variable "ca_certificate" {
  type = string
}

variable "client_certificate" {
  type = string
}

variable "client_certificate_key" {
  type      = string
  sensitive = true
}
