resource "awx_organization" "org" {
  name = "CredentialTypes"
}

locals {
  ssh_key = <<-EOT
    -----BEGIN RSA PRIVATE KEY-----
    MIIEowIBAAKCAQEAxkYZvo75rTb+x9+QYzGNGU4JaIoJb3/aSv2ZrPQGuZrkXkLO
    BBdIpuHSWDVG0NElJR0wvIlE6QG3b3znMWoxfo4DoGMgJPc3B0EQCJNi3Ba0xMCU
    bz4Xsa7iBL0kFv7DCNg76WV7qmJv5XzezCRYraHOlwjAupJtoUUw7oIJidt+O+jn
    xM7l4lrdBvA16iBPriev+byH5/bQVXb0Yjgy0FKwzErZbaJl5EqbHDMItbQjzr4G
    RmXf5AO4/Hm2Y5NFAYZJ3g9rmcWyldONomemJKGjJ/xpe6+hAeLFCdFDhuVNzcg0
    6jnRbdrvAUh4lhwbQbGN5Thfqu1hLFZNntwhgwIDAQABAoIBADRsdIId64YQ9GyE
    f/i6MURvja1GUhhZBa6uPuOr4dtRSPBjWXVbcqotKyCHdaHgqqGVhj1TThUNjyK7
    M6WtHkkH442Up/seIj6YxVR/I7RtjH45LQX8tOcWBkyadeBem9LK2YxseLYkMNHM
    olv3gDrofWcRuWObM8FgOf/WAg+gsgaqgnMMFwhKxd2KRFLg/dmpmf55jyAcHjRH
    EuUmybf8CRbo9FuukiTK+1+u53SFUoDVcM7jKhtmEw5kgmOubYZrqdV83qxIjo3/
    u17r6ISl9jw9r9q3QqBHLS0oU7cTqG16Ws3Yc3cT9tFn73wVtplPUJZRZ0psUvZU
    rTqitgECgYEA7MgbQitd4s1Y/aXxcNj/Po/fdThlsDl3EchchjmCwlAmaixZarWr
    IiWmmZ8RtmLayNU5k57vT4Y/IS9e1R6G4PsDVx89Tt/WTg5V9Vz68r/8atuVIxXO
    8vIsBKED/QchG5yV4xoPqO28pFixv6WQ0IGUu7ml+awohsXGNz73s5ECgYEA1l3f
    AR1dAdXd+E/5TjWwifzXqI2E+N5RcE2s4VUBVZJKCz/xRBoFPkIcUixfYljOYMsR
    x3Nml/OhZSZYU8sQZDCwF9cRoxdvWcuX7lVkFMYOfBtqK1OSPatYACsGd3BK8LKx
    Ii31qZI1uXEOyl9UaX6MPw7wv+GkCO2o2tEekdMCgYEAtM18GwO4ViYNTOn4uf3I
    sYH01LJy26SVqit/kzK5CR1gp/QjdxsviQyM8YMIaGeFvpMquvFCtUvCStm8JSqB
    PawOeektzFuZbdL+ijQsn05ANWxkrfzKILMRL5uvyvc1rbrcrSRKTujyAeCEId4P
    /VblNo9lEIgIDhlY6PiY9jECgYAJsEBt+nfDp3jayjKIWGDcO7M7iPnjjZi8rKYN
    oRvoVC5Ih+FNpKdsAuzowdylw0HAmiz5iLuwTnCS88K2Ns1M2e57hVrktiBwPIVn
    XjChx7nL3ilUP/iyFNZrW4Z1S3v3mPvsgYtC2LrY4MU8XEaiZffMuc4jvfz98k/Z
    Y+4vxQKBgEE0nmwsnKtucw+6/quASyry85EVGmjA32I66yJ1oAGYtWFeNKCtKgTZ
    f2xOo53Ny5VNhhRjYzaEjkf/boQTsX9bxoOs47FA8QdCv4FRhExImvSX2RqbtbRQ
    rA57+qvnMVzlQMZCNCJpwFvrjwSeNY3uuj+XJEk0hgku2x8hHtmK
    -----END RSA PRIVATE KEY-----
  EOT

  cert = <<-EOT
    -----BEGIN CERTIFICATE-----
    MIIDBTCCAe2gAwIBAgIUL7NXyNueMi9i9FI0S0qIb5PVGeMwDQYJKoZIhvcNAQEL
    BQAwEjEQMA4GA1UEAwwHdGZwcm9iZTAeFw0yNjA4MTgxODM1NTNaFw0yNzA4MTgx
    ODM1NTNaMBIxEDAOBgNVBAMMB3RmcHJvYmUwggEiMA0GCSqGSIb3DQEBAQUAA4IB
    DwAwggEKAoIBAQDQpNGqGzVTyinawfk8H+dgbdXy8r3Nzdiq65+0wN48FwOc102x
    xllXCqmQijEKqH4bM5WCKFoF5k6dJiT/luLYBIdQymPDYz9WRcwXimU382co70Od
    FmWWcCV9dc+rak7ksEmmZCOrHUUhkY1M4y2cF/JRMVfHfRAGALubu9u8d/SeoH8S
    Q3Fmm2TPNOsN2HWedByEdgbj5oufYtZJlhvu1kB8HnEgFI1kOH3EqXEwDwPEhD5X
    5FDiOQ5CGlARCxgOq080Tw6U4Tq/8ITK6BYEdwZUDIch7Cv26czzvOlSZLmP/Npw
    Z8auVl4H1liLL4Hq194lqLGYFIc28hGqVnVVAgMBAAGjUzBRMB0GA1UdDgQWBBQe
    K/ZpyKNHsUQqS8sXvawpKw2G3jAfBgNVHSMEGDAWgBQeK/ZpyKNHsUQqS8sXvawp
    Kw2G3jAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCaBl+REknT
    KBwrAzh9Nju1XH7qZod1kX/OAyAxwi7Ky4fYAyJFrhwr7E75EQiZRKI+0JRGXEQ2
    OIZCOu2Kb4zU/Di6CEJC1yCEUuRl9JA7Y2FvFp3mh7NrH/VMAsABmMPWZYaKrWqt
    U0Icc0SiEkTjr49n4JIdJt+YSKpm8v6jsVcwZ7ogI7CTHbpTGAywQPHSzaCZsnOQ
    KEgqEHZlJZOvQA+UOSzH/mObQoH5O01778qXxY7lQDr1q2imZznLUVXS2QFnOhyy
    CdBmPD6u0UVwbnemqJvchIwPhWKFjKpu82WuFaeS4KQf7ZZmX9zMU0QCY0sYeTk7
    XQ1fgYWlHtrq
    -----END CERTIFICATE-----
  EOT

  cert_key = <<-EOT
    -----BEGIN PRIVATE KEY-----
    MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDQpNGqGzVTyina
    wfk8H+dgbdXy8r3Nzdiq65+0wN48FwOc102xxllXCqmQijEKqH4bM5WCKFoF5k6d
    JiT/luLYBIdQymPDYz9WRcwXimU382co70OdFmWWcCV9dc+rak7ksEmmZCOrHUUh
    kY1M4y2cF/JRMVfHfRAGALubu9u8d/SeoH8SQ3Fmm2TPNOsN2HWedByEdgbj5ouf
    YtZJlhvu1kB8HnEgFI1kOH3EqXEwDwPEhD5X5FDiOQ5CGlARCxgOq080Tw6U4Tq/
    8ITK6BYEdwZUDIch7Cv26czzvOlSZLmP/NpwZ8auVl4H1liLL4Hq194lqLGYFIc2
    8hGqVnVVAgMBAAECggEAEHfwpbA5RMMYYSWTljjqyDF7OaPf3UVXdLby95nhS2sX
    Saiz6lQ+kTCCAwA+uo7jSfiHyfWElOJ2x5AvIO59usZxZiwyH4YaBdBYuZieqyx+
    8safTAsBxelVsuIFnDaLV3aJEFq8oavGh7hc/RnnFnhMaTdkwetU7whdRpUCKo1v
    uPE9F56y+qKlpNKugoS/vqv3+FEQQYXn9IZr7/uf2HN79hwFpUahjoJspNSCrMfD
    6SelawWt0pUSSDZ5nfpC3D6cRdRxygo5GULr6Ttgzi3fX7fbR6+XBeuwniq63OnA
    /AxeTJH1FK6dfZmw0WIhN+gBfrzpyX6QxoO3oHFzkQKBgQD8z319AUFHV8+X4yn5
    MrAESPUnZ7kQrUlyfGL8NPmUL10U3a0UtSAfhrlxrPkU9rB0WB5i2Zg15amDEMyV
    KC4LIf2CYMoXH9hl3ieb3Zx8gmGbdgSl4lVB3ABe4fArOBrjU57qu6w1Vj+WTG2/
    CTH+ayM23zayCiZHbwjT1fdIBQKBgQDTRq6s4x9O4Z6EXGk0M7isFsmbptnG6Q+u
    g1KS5FHo+g43t+WxGCrcyWrQuglflDzFT+f6z640FsyLh3Eycoakr1beWBFZ9v++
    WmQ2evARmw7tfezPS1l20SSM4ykGW2r0MLYfy53fDPjifIdQwHnsOzTjHpY7Ghf0
    YZThn+CJEQKBgCTzgukBhPQTjqEpr9nfocCOlV6UF4Wrfl/AdItehtg08Ar7t6+e
    JbrV60cFPEbnI7Vtp8tU+J1wGw2wJ+bpP4mbBz3GPeORYQTKqslqY9QDDpc1ccp5
    QXLl4Bv+NCdtBOkTMNgoZCxLlfG27BePFYKVUElV6N2vBBdO0GB+Fq2tAoGBAIyN
    lkEf68EfacRmcfqveejKVB8/tduVSsBvTuy4BiH20KSSq7TP/hvZxzzKttemF3Ow
    gASkSyEOFc+xMEY+WIKQvzq13f06y9KH8ekoijM8M4cdYgBaAU+BPU6ZBL7c7F89
    VLv8Sq+Lwrtx2teG5NWaYcnCnHD/H/aNQG2FXevxAoGAXZEv5s7XNp7oobqv0XJN
    /ga5QMjuK9LU/F1TELjfJx8wl4E5X0Eil4rO5f2sGsXjQ59wRZh07Gj4nXwzMBZT
    KooReo49iENL8+qEEUU+kvplhezcLoKIiASL4KtWG4+P5y4W/0Tm1nntJCliTYtZ
    dYGbAzIcdlCKjgGnKZ6XtFo=
    -----END PRIVATE KEY-----
  EOT

  gpg_key = <<-EOT
    -----BEGIN PGP PUBLIC KEY BLOCK-----

    mDMEaoSl6xYJKwYBBAHaRw8BAQdAbbmGxsDzP8FJxHgOnpCHEkGxpMv1by5XZH6R
    5yBNZgC0HXRmcHJvYmUgPHRmcHJvYmVAZXhhbXBsZS5jb20+iK8EExYKAFcWIQRe
    33lH3Vkb+xWTtwxa1X9RTUjjoQUCaoSl6xsUgAAAAAAEAA5tYW51MiwyLjUrMS4x
    MiwwLDMCGwMFCwkIBwICIgIGFQoJCAsCBBYCAwECHgcCF4AACgkQWtV/UU1I46Fy
    WwD8DR/igN8NzoHr5TsDUMZoTPdFIB2cE8Uj9VX32Tmvg8MBAPgi785IORd9BuQr
    aCmtRBgyh4qhOghhz8tWYg9/MDgPuDgEaoSl6xIKKwYBBAGXVQEFAQEHQK4IBaIE
    HY4BUmCHiCWa72GYsqWP9hT6NoNdKhoCqCY/AwEIB4iUBBgWCgA8FiEEXt95R91Z
    G/sVk7cMWtV/UU1I46EFAmqEpesbFIAAAAAABAAObWFudTIsMi41KzEuMTIsMCwz
    AhsMAAoJEFrVf1FNSOOhmYYBAMbkKR6+ytAMIvEVG1SBJpzyBI+4u7v2hX1fUgp/
    kj9TAP9QeKezEAL/JZ5Es8fa8MqVgkV9SBSYsEw6z/mOkZNpAw==
    =QpNb
    -----END PGP PUBLIC KEY BLOCK-----
  EOT
}

resource "awx_credential_aim" "this" {
  name          = "aim"
  organization  = awx_organization.org.id
  url           = "https://aim.example.com"
  webservice_id = "ws"
  app_id        = "app-id"
  client_key    = local.cert_key
  client_cert   = local.cert
  verify        = true
}

resource "awx_credential_aws" "this" {
  name           = "aws"
  organization   = awx_organization.org.id
  username       = "AKIAEXAMPLE"
  password       = "secret-key"
  security_token = "sts-token"
}

resource "awx_credential_aws_secretsmanager_credential" "this" {
  name           = "aws_secretsmanager_credential"
  organization   = awx_organization.org.id
  aws_access_key = "AKIAEXAMPLE"
  aws_secret_key = "secret-key"
}

resource "awx_credential_azure_kv" "this" {
  name         = "azure_kv"
  organization = awx_organization.org.id
  url          = "https://kv.example.com"
  client       = "client-id"
  secret       = "client-secret"
  tenant       = "tenant-id"
  cloud_name   = "AzureCloud"
}

resource "awx_credential_azure_rm" "this" {
  name              = "azure_rm"
  organization      = awx_organization.org.id
  subscription      = "subscription-id"
  client            = "client-id"
  secret            = "client-secret"
  tenant            = "tenant-id"
  cloud_environment = "AzureCloud"
}

resource "awx_credential_bitbucket_dc_token" "this" {
  name         = "bitbucket_dc_token"
  organization = awx_organization.org.id
  token        = "bitbucket-token"
}

resource "awx_credential_centrify_vault_kv" "this" {
  name                 = "centrify_vault_kv"
  organization         = awx_organization.org.id
  url                  = "https://centrify.example.com"
  client_id            = "client-id"
  client_password      = "client-secret"
  oauth_application_id = "awx"
  oauth_scope          = "awx"
}

resource "awx_credential_conjur" "this" {
  name         = "conjur"
  organization = awx_organization.org.id
  url          = "https://conjur.example.com"
  api_key      = "api-key"
  account      = "account"
  username     = "conjur-user"
  cacert       = local.cert
}

resource "awx_credential_controller" "this" {
  name         = "controller"
  organization = awx_organization.org.id
  host         = "https://controller.example.com"
  username     = "admin"
  password     = "controller-secret"
  verify_ssl   = true
}

resource "awx_credential_galaxy_api_token" "this" {
  name         = "galaxy_api_token"
  organization = awx_organization.org.id
  url          = "https://galaxy.ansible.com/"
  auth_url     = "https://sso.example.com/"
  token        = "galaxy-token"
}

resource "awx_credential_gce" "this" {
  name         = "gce"
  organization = awx_organization.org.id
  username     = "svc@example.com"
  project      = "gce-project"
  ssh_key_data = local.ssh_key
}

resource "awx_credential_github_token" "this" {
  name         = "github_token"
  organization = awx_organization.org.id
  token        = "github-token"
}

resource "awx_credential_gitlab_token" "this" {
  name         = "gitlab_token"
  organization = awx_organization.org.id
  token        = "gitlab-token"
}

resource "awx_credential_gpg_public_key" "this" {
  name           = "gpg_public_key"
  organization   = awx_organization.org.id
  gpg_public_key = local.gpg_key
}

resource "awx_credential_hashivault_kv" "this" {
  name                = "hashivault_kv"
  organization        = awx_organization.org.id
  url                 = "https://vault.example.com"
  token               = "vault-token"
  cacert              = local.cert
  role_id             = "role-id"
  secret_id           = "secret-id"
  client_cert_public  = local.cert
  client_cert_private = local.cert_key
  client_cert_role    = "cert-role"
  namespace           = "ns"
  kubernetes_role     = "k8s-role"
  username            = "vault-user"
  password            = "vault-secret"
  default_auth_path   = "approle"
  api_version         = "v1"
}

resource "awx_credential_hashivault_ssh" "this" {
  name                = "hashivault_ssh"
  organization        = awx_organization.org.id
  url                 = "https://vault.example.com"
  token               = "vault-token"
  cacert              = local.cert
  role_id             = "role-id"
  secret_id           = "secret-id"
  client_cert_public  = local.cert
  client_cert_private = local.cert_key
  client_cert_role    = "cert-role"
  namespace           = "ns"
  kubernetes_role     = "k8s-role"
  username            = "vault-user"
  password            = "vault-secret"
  default_auth_path   = "approle"
}

resource "awx_credential_insights" "this" {
  name         = "insights"
  organization = awx_organization.org.id
  username     = "insights-user"
  password     = "insights-secret"
}

resource "awx_credential_kubernetes_bearer_token" "this" {
  name         = "kubernetes_bearer_token"
  organization = awx_organization.org.id
  host         = "https://k8s.example.com"
  bearer_token = "k8s-token"
  verify_ssl   = true
  ssl_ca_cert  = local.cert
}

resource "awx_credential_net" "this" {
  name               = "net"
  organization       = awx_organization.org.id
  username           = "net-user"
  password           = "net-secret"
  ssh_key_data       = local.ssh_key
  authorize          = true
  authorize_password = "enable-secret"
}

resource "awx_credential_openstack" "this" {
  name                = "openstack"
  organization        = awx_organization.org.id
  username            = "os-user"
  password            = "os-secret"
  host                = "https://openstack.example.com"
  project             = "os-project"
  project_domain_name = "os-project-domain"
  domain              = "os-domain"
  region              = "os-region"
  verify_ssl          = true
}

resource "awx_credential_registry" "this" {
  name         = "registry"
  organization = awx_organization.org.id
  host         = "quay.io"
  username     = "robot"
  password     = "registry-secret"
  verify_ssl   = true
}

resource "awx_credential_rhv" "this" {
  name         = "rhv"
  organization = awx_organization.org.id
  host         = "https://rhv.example.com"
  username     = "rhv-user"
  password     = "rhv-secret"
  ca_file      = "/etc/pki/ca.crt"
}

resource "awx_credential_satellite6" "this" {
  name         = "satellite6"
  organization = awx_organization.org.id
  host         = "https://satellite.example.com"
  username     = "sat-user"
  password     = "sat-secret"
}

resource "awx_credential_scm" "this" {
  name         = "scm"
  organization = awx_organization.org.id
  username     = "git"
  password     = "scm-secret"
  ssh_key_data = local.ssh_key
}

resource "awx_credential_ssh" "this" {
  name            = "ssh"
  organization    = awx_organization.org.id
  username        = "deploy"
  ssh_key_data    = local.ssh_key
  become_method   = "sudo"
  become_username = "root"
  become_password = "become-secret"
}

resource "awx_credential_terraform" "this" {
  name            = "terraform"
  organization    = awx_organization.org.id
  configuration   = "backend \"local\" {\n  path = \"terraform.tfstate\"\n}\n"
  gce_credentials = "{\"type\":\"service_account\"}"
}

resource "awx_credential_thycotic_dsv" "this" {
  name          = "thycotic_dsv"
  organization  = awx_organization.org.id
  tenant        = "dsv-tenant"
  tld           = "com"
  client_id     = "client-id"
  client_secret = "client-secret"
}

resource "awx_credential_thycotic_tss" "this" {
  name         = "thycotic_tss"
  organization = awx_organization.org.id
  server_url   = "https://tss.example.com"
  username     = "tss-user"
  domain       = "tss-domain"
  password     = "tss-secret"
}

resource "awx_credential_vault" "this" {
  name           = "vault"
  organization   = awx_organization.org.id
  vault_password = "vault-secret"
  vault_id       = "primary"
}

resource "awx_credential_vmware" "this" {
  name         = "vmware"
  organization = awx_organization.org.id
  host         = "https://vcenter.example.com"
  username     = "vc-user"
  password     = "vc-secret"
}

data "awx_credential_registry" "by_name" {
  name = awx_credential_registry.this.name
}

data "awx_credential_vault" "by_id" {
  id = awx_credential_vault.this.id
}
