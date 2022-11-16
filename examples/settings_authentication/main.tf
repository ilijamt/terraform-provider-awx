terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_settings_auth_google_oauth2" "default" {
  social_auth_google_oauth2_auth_extra_arguments = jsonencode({})
  social_auth_google_oauth2_key                  = "oauth2_key"
  social_auth_google_oauth2_organization_map     = jsonencode({})
  social_auth_google_oauth2_secret               = "secret"
  social_auth_google_oauth2_team_map             = jsonencode({})
  social_auth_google_oauth2_whitelisted_domains = [
    "https://example.com"
  ]
}

resource "awx_settings_auth_saml" "default" {
  saml_auto_create_objects = true
  social_auth_saml_enabled_idps = jsonencode({
    "authentik" : {
      "attr_email" : "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
      "attr_first_name" : "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name",
      "attr_groups" : "http://schemas.xmlsoap.org/claims/Group",
      "attr_user_permanent_id" : "http://schemas.goauthentik.io/2021/02/saml/uid",
      "attr_username" : "http://schemas.goauthentik.io/2021/02/saml/username",
      "entity_id" : "https://awx.example.com/sso/metadata/saml/",
      "url" : "https://sso.example.com/application/saml/msp-awx/sso/binding/redirect/",
      "x509cert" : "MIICZjCCAc+gAwIBAgIUQ1t6poOhmqAiTkhxbXnpnmHtcSgwDQYJKoZIhvcNAQELBQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMjExMDkxNjIyMjJaFw0zMjExMDYxNjIyMjJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBALEwkxf4IyofkaGmh4L6J3uvLVMCQYmzWr7U4RcsVSaKjVCi9Yh7nhP4lvWwfsiQMZBylIAfBrJwIb+A4i4OpVXAZz0QNX7+676+wyHcEjODrLadKabkxrVZ/a78/DfQ6Er3USrho9dqU/8AsxUVidWIae1HBvPCeow4z/lJJ3ITAgMBAAGjUzBRMB0GA1UdDgQWBBRRWCOAhV5gjo8+ce439wPfELkKWTAfBgNVHSMEGDAWgBRRWCOAhV5gjo8+ce439wPfELkKWTAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4GBAB28PECTDT6JuuBawRPq7O/wsORYz3njPIrNT1Nbt08H1SpW888GK6XkIJJjkODCre7Yrry2K2ETcn2mrtK4zDBOgp+1hfdHpUBEFXALiZWFsgNU8vHvX5s5LTrZnSD1giUztNdCgfVsD1ML5KpjORj1OuDee0RF0X1SRE2gYU+S"
    }
  })
  social_auth_saml_extra_data = []
  social_auth_saml_org_info = jsonencode({
    "en-US" : {
      "displayname" : "authentik",
      "name" : "authentik",
      "url" : "https://authentik.example.com"
    }
  })
  social_auth_saml_organization_attr = jsonencode({})
  social_auth_saml_organization_map = jsonencode({
    "Default" : {
      "users" : true
    }
  })
  social_auth_saml_security_config = jsonencode(
    {
      requestedAuthnContext = false
    }
  )
  social_auth_saml_support_contact = jsonencode({
    "emailAddress" : "admin@example.com",
    "givenName" : "example.com"
  })
  social_auth_saml_team_attr = jsonencode({})
  social_auth_saml_technical_contact = jsonencode({
    "emailAddress" : "admin@example.com",
    "givenName" : "example.com"
  })
  social_auth_saml_user_flags_by_attr = jsonencode({
    "is_superuser_attr" : "awx_is_superuser",
    "is_superuser_value" : [
      "True"
    ]
  })

  social_auth_saml_sp_private_key = <<EOT
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALEwkxf4IyofkaGm
h4L6J3uvLVMCQYmzWr7U4RcsVSaKjVCi9Yh7nhP4lvWwfsiQMZBylIAfBrJwIb+A
4i4OpVXAZz0QNX7+676+wyHcEjODrLadKabkxrVZ/a78/DfQ6Er3USrho9dqU/8A
sxUVidWIae1HBvPCeow4z/lJJ3ITAgMBAAECgYEAg0rb44ng/Ihxz5bmmH2lnfc6
nWRxjYgReI+izhnyamGFvrdROjVm96NesGR8PT7jCwZr1NeojbWavBzS+4+sn7Z6
7bNjRQoGznQLI5d9ogHaDpxIfRJF3wHuTybJ2xR6au92lphR2UaeAeyBVVe5cPkA
QHDIOi9481LR0MXYExkCQQDikVY5f4Gfqy4LraXiYZqMRzEZGbUl6qVNaY8/5RuG
wfmZKEL0hofAPSsWGAkZpq9efLvahsfjMmlujO2ejJAXAkEAyDUjwHQvlWooFULl
lJUTlt2in+AXTsyJNyQL+XTqiz9d65+mscdpH9Y5rHohiM/jR3Ugl08Mtm+PE+Qq
1FDPZQJAR/aBvKGTOnPUnED7f3wg6o1yOta/gtuUxZHRvim3JIZYER2IpsJUO+sx
1EKuIUegTBKyWCaXNsK8WjDJCKL84QJAZMx8Z5UXr/52l93KgPhdmIOWMTA+C+pm
22BGtx3qSJlqzArhfniLsP/GodQLtjoUkBGkiwm9uMyKGNWzypm1EQJAJUEQbOQZ
eIZk99mDSLYKTyKOd3iUeaRoOY3NJLlb0cRzng9h4DgvDNF/QyYhhhPZYGwVjrO/
RH5QBBZdpsFK9g==
-----END PRIVATE KEY-----
EOT

  social_auth_saml_sp_public_cert = <<EOT
-----BEGIN CERTIFICATE-----
MIICZjCCAc+gAwIBAgIUQ1t6poOhmqAiTkhxbXnpnmHtcSgwDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMjExMDkxNjIyMjJaFw0zMjEx
MDYxNjIyMjJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwgZ8wDQYJKoZIhvcNAQEB
BQADgY0AMIGJAoGBALEwkxf4IyofkaGmh4L6J3uvLVMCQYmzWr7U4RcsVSaKjVCi
9Yh7nhP4lvWwfsiQMZBylIAfBrJwIb+A4i4OpVXAZz0QNX7+676+wyHcEjODrLad
KabkxrVZ/a78/DfQ6Er3USrho9dqU/8AsxUVidWIae1HBvPCeow4z/lJJ3ITAgMB
AAGjUzBRMB0GA1UdDgQWBBRRWCOAhV5gjo8+ce439wPfELkKWTAfBgNVHSMEGDAW
gBRRWCOAhV5gjo8+ce439wPfELkKWTAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
DQEBCwUAA4GBAB28PECTDT6JuuBawRPq7O/wsORYz3njPIrNT1Nbt08H1SpW888G
K6XkIJJjkODCre7Yrry2K2ETcn2mrtK4zDBOgp+1hfdHpUBEFXALiZWFsgNU8vHv
X5s5LTrZnSD1giUztNdCgfVsD1ML5KpjORj1OuDee0RF0X1SRE2gYU+S
-----END CERTIFICATE-----
EOT
}

resource "awx_settings_auth_azuread_oauth2" "default" {
  social_auth_azuread_oauth2_key              = "key"
  social_auth_azuread_oauth2_secret           = "access"
  social_auth_azuread_oauth2_organization_map = jsonencode({})
  social_auth_azuread_oauth2_team_map         = jsonencode({})
}

resource "awx_settings_auth_ldap" "default" {
  auth_ldap_group_search = []
  auth_ldap_group_type   = "MemberDNGroupType"
  auth_ldap_group_type_params = jsonencode(
    {
      member_attr = "member"
      name_attr   = "cn"
    }
  )
  auth_ldap_organization_map    = jsonencode({})
  auth_ldap_start_tls           = false
  auth_ldap_team_map            = jsonencode({})
  auth_ldap_user_attr_map       = jsonencode({})
  auth_ldap_user_flags_by_group = jsonencode({})
  auth_ldap_user_search         = []

  auth_ldap_1_connection_options = jsonencode(
    {
      OPT_NETWORK_TIMEOUT = 30
      OPT_REFERRALS       = 0
    }
  )
  auth_ldap_1_group_search = []
  auth_ldap_1_group_type   = "MemberDNGroupType"
  auth_ldap_1_group_type_params = jsonencode(
    {
      member_attr = "member"
      name_attr   = "cn"
    }
  )
  auth_ldap_1_organization_map    = jsonencode({})
  auth_ldap_1_start_tls           = false
  auth_ldap_1_team_map            = jsonencode({})
  auth_ldap_1_user_attr_map       = jsonencode({})
  auth_ldap_1_user_flags_by_group = jsonencode({})
  auth_ldap_1_user_search         = []

  auth_ldap_2_connection_options = jsonencode(
    {
      OPT_NETWORK_TIMEOUT = 30
      OPT_REFERRALS       = 0
    }
  )
  auth_ldap_2_group_search = []
  auth_ldap_2_group_type   = "MemberDNGroupType"
  auth_ldap_2_group_type_params = jsonencode(
    {
      member_attr = "member"
      name_attr   = "cn"
    }
  )
  auth_ldap_2_organization_map    = jsonencode({})
  auth_ldap_2_start_tls           = false
  auth_ldap_2_team_map            = jsonencode({})
  auth_ldap_2_user_attr_map       = jsonencode({})
  auth_ldap_2_user_flags_by_group = jsonencode({})
  auth_ldap_2_user_search         = []

  auth_ldap_3_connection_options = jsonencode(
    {
      OPT_NETWORK_TIMEOUT = 30
      OPT_REFERRALS       = 0
    }
  )
  auth_ldap_3_group_search = []
  auth_ldap_3_group_type   = "MemberDNGroupType"
  auth_ldap_3_group_type_params = jsonencode(
    {
      member_attr = "member"
      name_attr   = "cn"
    }
  )
  auth_ldap_3_organization_map    = jsonencode({})
  auth_ldap_3_start_tls           = false
  auth_ldap_3_team_map            = jsonencode({})
  auth_ldap_3_user_attr_map       = jsonencode({})
  auth_ldap_3_user_flags_by_group = jsonencode({})
  auth_ldap_3_user_search         = []

  auth_ldap_4_connection_options = jsonencode(
    {
      OPT_NETWORK_TIMEOUT = 30
      OPT_REFERRALS       = 0
    }
  )
  auth_ldap_4_group_search = []
  auth_ldap_4_group_type   = "MemberDNGroupType"
  auth_ldap_4_group_type_params = jsonencode(
    {
      member_attr = "member"
      name_attr   = "cn"
    }
  )
  auth_ldap_4_organization_map    = jsonencode({})
  auth_ldap_4_start_tls           = false
  auth_ldap_4_team_map            = jsonencode({})
  auth_ldap_4_user_attr_map       = jsonencode({})
  auth_ldap_4_user_flags_by_group = jsonencode({})
  auth_ldap_4_user_search         = []

  auth_ldap_5_connection_options = jsonencode(
    {
      OPT_NETWORK_TIMEOUT = 30
      OPT_REFERRALS       = 0
    }
  )
  auth_ldap_5_group_search = []
  auth_ldap_5_group_type   = "MemberDNGroupType"
  auth_ldap_5_group_type_params = jsonencode(
    {
      member_attr = "member"
      name_attr   = "cn"
    }
  )
  auth_ldap_5_organization_map    = jsonencode({})
  auth_ldap_5_start_tls           = false
  auth_ldap_5_team_map            = jsonencode({})
  auth_ldap_5_user_attr_map       = jsonencode({})
  auth_ldap_5_user_flags_by_group = jsonencode({})
  auth_ldap_5_user_search         = []
  auth_ldap_connection_options = jsonencode(
    {
      OPT_NETWORK_TIMEOUT = 30
      OPT_REFERRALS       = 0
    }
  )
}
