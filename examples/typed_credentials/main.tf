terraform {
  required_providers {
    awx = {
      source = "ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "typed" {
  name = "Typed Credentials"
}

resource "awx_credential_aim" "example" {
  name         = "CyberArk Central Credential Provider Lookup"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  app_id       = var.secrets["aim_app_id"]
  verify       = true
}

resource "awx_credential_aws" "example" {
  name         = "Amazon Web Services"
  organization = awx_organization.typed.id
  username     = "service-account"
  password     = var.secrets["aws_password"]
}

resource "awx_credential_aws_secretsmanager_credential" "example" {
  name           = "AWS Secrets Manager lookup"
  organization   = awx_organization.typed.id
  aws_access_key = "AKIAEXAMPLE"
  aws_secret_key = var.secrets["aws_secretsmanager_credential_aws_secret_key"]
}

resource "awx_credential_azure_kv" "example" {
  name         = "Microsoft Azure Key Vault"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  client       = "client-id"
  secret       = var.secrets["azure_kv_secret"]
  tenant       = "tenant-id"
  cloud_name   = "AzureCloud"
}

resource "awx_credential_azure_rm" "example" {
  name         = "Microsoft Azure Resource Manager"
  organization = awx_organization.typed.id
  subscription = "subscription-id"
}

resource "awx_credential_bitbucket_dc_token" "example" {
  name         = "Bitbucket Data Center HTTP Access Token"
  organization = awx_organization.typed.id
  token        = var.secrets["bitbucket_dc_token_token"]
}

resource "awx_credential_centrify_vault_kv" "example" {
  name            = "Centrify Vault Credential Provider Lookup"
  organization    = awx_organization.typed.id
  url             = "https://service.example.com"
  client_id       = "client-id"
  client_password = var.secrets["centrify_vault_kv_client_password"]
}

resource "awx_credential_conjur" "example" {
  name         = "CyberArk Conjur Secrets Manager Lookup"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  api_key      = var.secrets["conjur_api_key"]
  account      = "account"
  username     = "service-account"
}

resource "awx_credential_controller" "example" {
  name         = "Red Hat Ansible Automation Platform"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  verify_ssl   = false
}

resource "awx_credential_galaxy_api_token" "example" {
  name         = "Ansible Galaxy/Automation Hub API Token"
  organization = awx_organization.typed.id
  url          = "https://galaxy.ansible.com/"
}

resource "awx_credential_gce" "example" {
  name         = "Google Compute Engine"
  organization = awx_organization.typed.id
  username     = "service-account@my-project.iam.gserviceaccount.com"
  ssh_key_data = var.ssh_private_key
}

resource "awx_credential_github_token" "example" {
  name         = "GitHub Personal Access Token"
  organization = awx_organization.typed.id
  token        = var.secrets["github_token_token"]
}

resource "awx_credential_gitlab_token" "example" {
  name         = "GitLab Personal Access Token"
  organization = awx_organization.typed.id
  token        = var.secrets["gitlab_token_token"]
}

resource "awx_credential_gpg_public_key" "example" {
  name           = "GPG Public Key"
  organization   = awx_organization.typed.id
  gpg_public_key = var.gpg_public_key
}

resource "awx_credential_hashivault_kv" "example" {
  name         = "HashiCorp Vault Secret Lookup"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  api_version  = "v1"
}

resource "awx_credential_hashivault_ssh" "example" {
  name         = "HashiCorp Vault Signed SSH"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
}

resource "awx_credential_insights" "example" {
  name         = "Insights"
  organization = awx_organization.typed.id
  username     = "service-account"
  password     = var.secrets["insights_password"]
}

resource "awx_credential_kubernetes_bearer_token" "example" {
  name         = "OpenShift or Kubernetes API Bearer Token"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  bearer_token = var.secrets["kubernetes_bearer_token_bearer_token"]
  verify_ssl   = true
}

resource "awx_credential_net" "example" {
  name         = "Network"
  organization = awx_organization.typed.id
  username     = "service-account"
  authorize    = false
}

resource "awx_credential_openstack" "example" {
  name         = "OpenStack"
  organization = awx_organization.typed.id
  username     = "service-account"
  password     = var.secrets["openstack_password"]
  host         = "https://service.example.com"
  project      = "my-project"
  verify_ssl   = true
}

resource "awx_credential_registry" "example" {
  name         = "Container Registry"
  organization = awx_organization.typed.id
  host         = "quay.io"
  verify_ssl   = true
}

resource "awx_credential_rhv" "example" {
  name         = "Red Hat Virtualization"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  username     = "service-account"
  password     = var.secrets["rhv_password"]
}

resource "awx_credential_satellite6" "example" {
  name         = "Red Hat Satellite 6"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  username     = "service-account"
  password     = var.secrets["satellite6_password"]
}

resource "awx_credential_scm" "example" {
  name         = "Source Control"
  organization = awx_organization.typed.id
}

resource "awx_credential_ssh" "example" {
  name         = "Machine"
  organization = awx_organization.typed.id
}

resource "awx_credential_terraform" "example" {
  name          = "Terraform backend configuration"
  organization  = awx_organization.typed.id
  configuration = var.secrets["terraform_configuration"]
}

resource "awx_credential_thycotic_dsv" "example" {
  name          = "Thycotic DevOps Secrets Vault"
  organization  = awx_organization.typed.id
  tenant        = "tenant-id"
  tld           = "com"
  client_id     = "client-id"
  client_secret = var.secrets["thycotic_dsv_client_secret"]
}

resource "awx_credential_thycotic_tss" "example" {
  name         = "Thycotic Secret Server"
  organization = awx_organization.typed.id
  server_url   = "https://tss.example.com"
  username     = "service-account"
  password     = var.secrets["thycotic_tss_password"]
}

resource "awx_credential_vault" "example" {
  name           = "Vault"
  organization   = awx_organization.typed.id
  vault_password = var.secrets["vault_vault_password"]
}

resource "awx_credential_vmware" "example" {
  name         = "VMware vCenter"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  username     = "service-account"
  password     = var.secrets["vmware_password"]
}

# Data sources expose the readable inputs only: AWX answers every secret with
# the literal "$encrypted$".
data "awx_credential_registry" "example" {
  name = awx_credential_registry.example.name
}

output "registry_host" {
  value = data.awx_credential_registry.example.host
}

output "registry_verify_ssl" {
  value = data.awx_credential_registry.example.verify_ssl
}
