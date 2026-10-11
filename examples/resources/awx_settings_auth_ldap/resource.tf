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
