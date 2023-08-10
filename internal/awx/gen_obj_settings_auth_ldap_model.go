package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthLdapTerraformModel maps the schema for SettingsAuthLDAP when using Data Source
type settingsAuthLdapTerraformModel struct {
	// AUTH_LDAP_1_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_1_BIND_DN types.String `tfsdk:"auth_ldap_1_bind_dn" json:"AUTH_LDAP_1_BIND_DN"`
	// AUTH_LDAP_1_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_1_BIND_PASSWORD types.String `tfsdk:"auth_ldap_1_bind_password" json:"AUTH_LDAP_1_BIND_PASSWORD"`
	// AUTH_LDAP_1_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_1_CONNECTION_OPTIONS types.String `tfsdk:"auth_ldap_1_connection_options" json:"AUTH_LDAP_1_CONNECTION_OPTIONS"`
	// AUTH_LDAP_1_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_1_DENY_GROUP types.String `tfsdk:"auth_ldap_1_deny_group" json:"AUTH_LDAP_1_DENY_GROUP"`
	// AUTH_LDAP_1_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_1_GROUP_SEARCH types.List `tfsdk:"auth_ldap_1_group_search" json:"AUTH_LDAP_1_GROUP_SEARCH"`
	// AUTH_LDAP_1_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_1_GROUP_TYPE types.String `tfsdk:"auth_ldap_1_group_type" json:"AUTH_LDAP_1_GROUP_TYPE"`
	// AUTH_LDAP_1_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_1_GROUP_TYPE_PARAMS types.String `tfsdk:"auth_ldap_1_group_type_params" json:"AUTH_LDAP_1_GROUP_TYPE_PARAMS"`
	// AUTH_LDAP_1_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_1_ORGANIZATION_MAP types.String `tfsdk:"auth_ldap_1_organization_map" json:"AUTH_LDAP_1_ORGANIZATION_MAP"`
	// AUTH_LDAP_1_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_1_REQUIRE_GROUP types.String `tfsdk:"auth_ldap_1_require_group" json:"AUTH_LDAP_1_REQUIRE_GROUP"`
	// AUTH_LDAP_1_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_1_SERVER_URI types.String `tfsdk:"auth_ldap_1_server_uri" json:"AUTH_LDAP_1_SERVER_URI"`
	// AUTH_LDAP_1_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_1_START_TLS types.Bool `tfsdk:"auth_ldap_1_start_tls" json:"AUTH_LDAP_1_START_TLS"`
	// AUTH_LDAP_1_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_1_TEAM_MAP types.String `tfsdk:"auth_ldap_1_team_map" json:"AUTH_LDAP_1_TEAM_MAP"`
	// AUTH_LDAP_1_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_1_USER_ATTR_MAP types.String `tfsdk:"auth_ldap_1_user_attr_map" json:"AUTH_LDAP_1_USER_ATTR_MAP"`
	// AUTH_LDAP_1_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_1_USER_DN_TEMPLATE types.String `tfsdk:"auth_ldap_1_user_dn_template" json:"AUTH_LDAP_1_USER_DN_TEMPLATE"`
	// AUTH_LDAP_1_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_1_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_1_user_flags_by_group" json:"AUTH_LDAP_1_USER_FLAGS_BY_GROUP"`
	// AUTH_LDAP_1_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_1_USER_SEARCH types.List `tfsdk:"auth_ldap_1_user_search" json:"AUTH_LDAP_1_USER_SEARCH"`
	// AUTH_LDAP_2_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_2_BIND_DN types.String `tfsdk:"auth_ldap_2_bind_dn" json:"AUTH_LDAP_2_BIND_DN"`
	// AUTH_LDAP_2_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_2_BIND_PASSWORD types.String `tfsdk:"auth_ldap_2_bind_password" json:"AUTH_LDAP_2_BIND_PASSWORD"`
	// AUTH_LDAP_2_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_2_CONNECTION_OPTIONS types.String `tfsdk:"auth_ldap_2_connection_options" json:"AUTH_LDAP_2_CONNECTION_OPTIONS"`
	// AUTH_LDAP_2_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_2_DENY_GROUP types.String `tfsdk:"auth_ldap_2_deny_group" json:"AUTH_LDAP_2_DENY_GROUP"`
	// AUTH_LDAP_2_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_2_GROUP_SEARCH types.List `tfsdk:"auth_ldap_2_group_search" json:"AUTH_LDAP_2_GROUP_SEARCH"`
	// AUTH_LDAP_2_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_2_GROUP_TYPE types.String `tfsdk:"auth_ldap_2_group_type" json:"AUTH_LDAP_2_GROUP_TYPE"`
	// AUTH_LDAP_2_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_2_GROUP_TYPE_PARAMS types.String `tfsdk:"auth_ldap_2_group_type_params" json:"AUTH_LDAP_2_GROUP_TYPE_PARAMS"`
	// AUTH_LDAP_2_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_2_ORGANIZATION_MAP types.String `tfsdk:"auth_ldap_2_organization_map" json:"AUTH_LDAP_2_ORGANIZATION_MAP"`
	// AUTH_LDAP_2_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_2_REQUIRE_GROUP types.String `tfsdk:"auth_ldap_2_require_group" json:"AUTH_LDAP_2_REQUIRE_GROUP"`
	// AUTH_LDAP_2_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_2_SERVER_URI types.String `tfsdk:"auth_ldap_2_server_uri" json:"AUTH_LDAP_2_SERVER_URI"`
	// AUTH_LDAP_2_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_2_START_TLS types.Bool `tfsdk:"auth_ldap_2_start_tls" json:"AUTH_LDAP_2_START_TLS"`
	// AUTH_LDAP_2_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_2_TEAM_MAP types.String `tfsdk:"auth_ldap_2_team_map" json:"AUTH_LDAP_2_TEAM_MAP"`
	// AUTH_LDAP_2_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_2_USER_ATTR_MAP types.String `tfsdk:"auth_ldap_2_user_attr_map" json:"AUTH_LDAP_2_USER_ATTR_MAP"`
	// AUTH_LDAP_2_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_2_USER_DN_TEMPLATE types.String `tfsdk:"auth_ldap_2_user_dn_template" json:"AUTH_LDAP_2_USER_DN_TEMPLATE"`
	// AUTH_LDAP_2_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_2_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_2_user_flags_by_group" json:"AUTH_LDAP_2_USER_FLAGS_BY_GROUP"`
	// AUTH_LDAP_2_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_2_USER_SEARCH types.List `tfsdk:"auth_ldap_2_user_search" json:"AUTH_LDAP_2_USER_SEARCH"`
	// AUTH_LDAP_3_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_3_BIND_DN types.String `tfsdk:"auth_ldap_3_bind_dn" json:"AUTH_LDAP_3_BIND_DN"`
	// AUTH_LDAP_3_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_3_BIND_PASSWORD types.String `tfsdk:"auth_ldap_3_bind_password" json:"AUTH_LDAP_3_BIND_PASSWORD"`
	// AUTH_LDAP_3_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_3_CONNECTION_OPTIONS types.String `tfsdk:"auth_ldap_3_connection_options" json:"AUTH_LDAP_3_CONNECTION_OPTIONS"`
	// AUTH_LDAP_3_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_3_DENY_GROUP types.String `tfsdk:"auth_ldap_3_deny_group" json:"AUTH_LDAP_3_DENY_GROUP"`
	// AUTH_LDAP_3_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_3_GROUP_SEARCH types.List `tfsdk:"auth_ldap_3_group_search" json:"AUTH_LDAP_3_GROUP_SEARCH"`
	// AUTH_LDAP_3_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_3_GROUP_TYPE types.String `tfsdk:"auth_ldap_3_group_type" json:"AUTH_LDAP_3_GROUP_TYPE"`
	// AUTH_LDAP_3_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_3_GROUP_TYPE_PARAMS types.String `tfsdk:"auth_ldap_3_group_type_params" json:"AUTH_LDAP_3_GROUP_TYPE_PARAMS"`
	// AUTH_LDAP_3_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_3_ORGANIZATION_MAP types.String `tfsdk:"auth_ldap_3_organization_map" json:"AUTH_LDAP_3_ORGANIZATION_MAP"`
	// AUTH_LDAP_3_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_3_REQUIRE_GROUP types.String `tfsdk:"auth_ldap_3_require_group" json:"AUTH_LDAP_3_REQUIRE_GROUP"`
	// AUTH_LDAP_3_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_3_SERVER_URI types.String `tfsdk:"auth_ldap_3_server_uri" json:"AUTH_LDAP_3_SERVER_URI"`
	// AUTH_LDAP_3_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_3_START_TLS types.Bool `tfsdk:"auth_ldap_3_start_tls" json:"AUTH_LDAP_3_START_TLS"`
	// AUTH_LDAP_3_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_3_TEAM_MAP types.String `tfsdk:"auth_ldap_3_team_map" json:"AUTH_LDAP_3_TEAM_MAP"`
	// AUTH_LDAP_3_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_3_USER_ATTR_MAP types.String `tfsdk:"auth_ldap_3_user_attr_map" json:"AUTH_LDAP_3_USER_ATTR_MAP"`
	// AUTH_LDAP_3_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_3_USER_DN_TEMPLATE types.String `tfsdk:"auth_ldap_3_user_dn_template" json:"AUTH_LDAP_3_USER_DN_TEMPLATE"`
	// AUTH_LDAP_3_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_3_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_3_user_flags_by_group" json:"AUTH_LDAP_3_USER_FLAGS_BY_GROUP"`
	// AUTH_LDAP_3_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_3_USER_SEARCH types.List `tfsdk:"auth_ldap_3_user_search" json:"AUTH_LDAP_3_USER_SEARCH"`
	// AUTH_LDAP_4_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_4_BIND_DN types.String `tfsdk:"auth_ldap_4_bind_dn" json:"AUTH_LDAP_4_BIND_DN"`
	// AUTH_LDAP_4_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_4_BIND_PASSWORD types.String `tfsdk:"auth_ldap_4_bind_password" json:"AUTH_LDAP_4_BIND_PASSWORD"`
	// AUTH_LDAP_4_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_4_CONNECTION_OPTIONS types.String `tfsdk:"auth_ldap_4_connection_options" json:"AUTH_LDAP_4_CONNECTION_OPTIONS"`
	// AUTH_LDAP_4_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_4_DENY_GROUP types.String `tfsdk:"auth_ldap_4_deny_group" json:"AUTH_LDAP_4_DENY_GROUP"`
	// AUTH_LDAP_4_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_4_GROUP_SEARCH types.List `tfsdk:"auth_ldap_4_group_search" json:"AUTH_LDAP_4_GROUP_SEARCH"`
	// AUTH_LDAP_4_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_4_GROUP_TYPE types.String `tfsdk:"auth_ldap_4_group_type" json:"AUTH_LDAP_4_GROUP_TYPE"`
	// AUTH_LDAP_4_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_4_GROUP_TYPE_PARAMS types.String `tfsdk:"auth_ldap_4_group_type_params" json:"AUTH_LDAP_4_GROUP_TYPE_PARAMS"`
	// AUTH_LDAP_4_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_4_ORGANIZATION_MAP types.String `tfsdk:"auth_ldap_4_organization_map" json:"AUTH_LDAP_4_ORGANIZATION_MAP"`
	// AUTH_LDAP_4_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_4_REQUIRE_GROUP types.String `tfsdk:"auth_ldap_4_require_group" json:"AUTH_LDAP_4_REQUIRE_GROUP"`
	// AUTH_LDAP_4_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_4_SERVER_URI types.String `tfsdk:"auth_ldap_4_server_uri" json:"AUTH_LDAP_4_SERVER_URI"`
	// AUTH_LDAP_4_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_4_START_TLS types.Bool `tfsdk:"auth_ldap_4_start_tls" json:"AUTH_LDAP_4_START_TLS"`
	// AUTH_LDAP_4_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_4_TEAM_MAP types.String `tfsdk:"auth_ldap_4_team_map" json:"AUTH_LDAP_4_TEAM_MAP"`
	// AUTH_LDAP_4_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_4_USER_ATTR_MAP types.String `tfsdk:"auth_ldap_4_user_attr_map" json:"AUTH_LDAP_4_USER_ATTR_MAP"`
	// AUTH_LDAP_4_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_4_USER_DN_TEMPLATE types.String `tfsdk:"auth_ldap_4_user_dn_template" json:"AUTH_LDAP_4_USER_DN_TEMPLATE"`
	// AUTH_LDAP_4_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_4_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_4_user_flags_by_group" json:"AUTH_LDAP_4_USER_FLAGS_BY_GROUP"`
	// AUTH_LDAP_4_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_4_USER_SEARCH types.List `tfsdk:"auth_ldap_4_user_search" json:"AUTH_LDAP_4_USER_SEARCH"`
	// AUTH_LDAP_5_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_5_BIND_DN types.String `tfsdk:"auth_ldap_5_bind_dn" json:"AUTH_LDAP_5_BIND_DN"`
	// AUTH_LDAP_5_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_5_BIND_PASSWORD types.String `tfsdk:"auth_ldap_5_bind_password" json:"AUTH_LDAP_5_BIND_PASSWORD"`
	// AUTH_LDAP_5_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_5_CONNECTION_OPTIONS types.String `tfsdk:"auth_ldap_5_connection_options" json:"AUTH_LDAP_5_CONNECTION_OPTIONS"`
	// AUTH_LDAP_5_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_5_DENY_GROUP types.String `tfsdk:"auth_ldap_5_deny_group" json:"AUTH_LDAP_5_DENY_GROUP"`
	// AUTH_LDAP_5_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_5_GROUP_SEARCH types.List `tfsdk:"auth_ldap_5_group_search" json:"AUTH_LDAP_5_GROUP_SEARCH"`
	// AUTH_LDAP_5_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_5_GROUP_TYPE types.String `tfsdk:"auth_ldap_5_group_type" json:"AUTH_LDAP_5_GROUP_TYPE"`
	// AUTH_LDAP_5_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_5_GROUP_TYPE_PARAMS types.String `tfsdk:"auth_ldap_5_group_type_params" json:"AUTH_LDAP_5_GROUP_TYPE_PARAMS"`
	// AUTH_LDAP_5_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_5_ORGANIZATION_MAP types.String `tfsdk:"auth_ldap_5_organization_map" json:"AUTH_LDAP_5_ORGANIZATION_MAP"`
	// AUTH_LDAP_5_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_5_REQUIRE_GROUP types.String `tfsdk:"auth_ldap_5_require_group" json:"AUTH_LDAP_5_REQUIRE_GROUP"`
	// AUTH_LDAP_5_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_5_SERVER_URI types.String `tfsdk:"auth_ldap_5_server_uri" json:"AUTH_LDAP_5_SERVER_URI"`
	// AUTH_LDAP_5_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_5_START_TLS types.Bool `tfsdk:"auth_ldap_5_start_tls" json:"AUTH_LDAP_5_START_TLS"`
	// AUTH_LDAP_5_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_5_TEAM_MAP types.String `tfsdk:"auth_ldap_5_team_map" json:"AUTH_LDAP_5_TEAM_MAP"`
	// AUTH_LDAP_5_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_5_USER_ATTR_MAP types.String `tfsdk:"auth_ldap_5_user_attr_map" json:"AUTH_LDAP_5_USER_ATTR_MAP"`
	// AUTH_LDAP_5_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_5_USER_DN_TEMPLATE types.String `tfsdk:"auth_ldap_5_user_dn_template" json:"AUTH_LDAP_5_USER_DN_TEMPLATE"`
	// AUTH_LDAP_5_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_5_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_5_user_flags_by_group" json:"AUTH_LDAP_5_USER_FLAGS_BY_GROUP"`
	// AUTH_LDAP_5_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_5_USER_SEARCH types.List `tfsdk:"auth_ldap_5_user_search" json:"AUTH_LDAP_5_USER_SEARCH"`
	// AUTH_LDAP_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_BIND_DN types.String `tfsdk:"auth_ldap_bind_dn" json:"AUTH_LDAP_BIND_DN"`
	// AUTH_LDAP_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_BIND_PASSWORD types.String `tfsdk:"auth_ldap_bind_password" json:"AUTH_LDAP_BIND_PASSWORD"`
	// AUTH_LDAP_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_CONNECTION_OPTIONS types.String `tfsdk:"auth_ldap_connection_options" json:"AUTH_LDAP_CONNECTION_OPTIONS"`
	// AUTH_LDAP_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_DENY_GROUP types.String `tfsdk:"auth_ldap_deny_group" json:"AUTH_LDAP_DENY_GROUP"`
	// AUTH_LDAP_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_GROUP_SEARCH types.List `tfsdk:"auth_ldap_group_search" json:"AUTH_LDAP_GROUP_SEARCH"`
	// AUTH_LDAP_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_GROUP_TYPE types.String `tfsdk:"auth_ldap_group_type" json:"AUTH_LDAP_GROUP_TYPE"`
	// AUTH_LDAP_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_GROUP_TYPE_PARAMS types.String `tfsdk:"auth_ldap_group_type_params" json:"AUTH_LDAP_GROUP_TYPE_PARAMS"`
	// AUTH_LDAP_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_ORGANIZATION_MAP types.String `tfsdk:"auth_ldap_organization_map" json:"AUTH_LDAP_ORGANIZATION_MAP"`
	// AUTH_LDAP_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_REQUIRE_GROUP types.String `tfsdk:"auth_ldap_require_group" json:"AUTH_LDAP_REQUIRE_GROUP"`
	// AUTH_LDAP_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_SERVER_URI types.String `tfsdk:"auth_ldap_server_uri" json:"AUTH_LDAP_SERVER_URI"`
	// AUTH_LDAP_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_START_TLS types.Bool `tfsdk:"auth_ldap_start_tls" json:"AUTH_LDAP_START_TLS"`
	// AUTH_LDAP_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_TEAM_MAP types.String `tfsdk:"auth_ldap_team_map" json:"AUTH_LDAP_TEAM_MAP"`
	// AUTH_LDAP_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_USER_ATTR_MAP types.String `tfsdk:"auth_ldap_user_attr_map" json:"AUTH_LDAP_USER_ATTR_MAP"`
	// AUTH_LDAP_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_USER_DN_TEMPLATE types.String `tfsdk:"auth_ldap_user_dn_template" json:"AUTH_LDAP_USER_DN_TEMPLATE"`
	// AUTH_LDAP_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_user_flags_by_group" json:"AUTH_LDAP_USER_FLAGS_BY_GROUP"`
	// AUTH_LDAP_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_USER_SEARCH types.List `tfsdk:"auth_ldap_user_search" json:"AUTH_LDAP_USER_SEARCH"`
}

// Clone the object
func (o *settingsAuthLdapTerraformModel) Clone() settingsAuthLdapTerraformModel {
	return settingsAuthLdapTerraformModel{
		AUTH_LDAP_1_BIND_DN:             o.AUTH_LDAP_1_BIND_DN,
		AUTH_LDAP_1_BIND_PASSWORD:       o.AUTH_LDAP_1_BIND_PASSWORD,
		AUTH_LDAP_1_CONNECTION_OPTIONS:  o.AUTH_LDAP_1_CONNECTION_OPTIONS,
		AUTH_LDAP_1_DENY_GROUP:          o.AUTH_LDAP_1_DENY_GROUP,
		AUTH_LDAP_1_GROUP_SEARCH:        o.AUTH_LDAP_1_GROUP_SEARCH,
		AUTH_LDAP_1_GROUP_TYPE:          o.AUTH_LDAP_1_GROUP_TYPE,
		AUTH_LDAP_1_GROUP_TYPE_PARAMS:   o.AUTH_LDAP_1_GROUP_TYPE_PARAMS,
		AUTH_LDAP_1_ORGANIZATION_MAP:    o.AUTH_LDAP_1_ORGANIZATION_MAP,
		AUTH_LDAP_1_REQUIRE_GROUP:       o.AUTH_LDAP_1_REQUIRE_GROUP,
		AUTH_LDAP_1_SERVER_URI:          o.AUTH_LDAP_1_SERVER_URI,
		AUTH_LDAP_1_START_TLS:           o.AUTH_LDAP_1_START_TLS,
		AUTH_LDAP_1_TEAM_MAP:            o.AUTH_LDAP_1_TEAM_MAP,
		AUTH_LDAP_1_USER_ATTR_MAP:       o.AUTH_LDAP_1_USER_ATTR_MAP,
		AUTH_LDAP_1_USER_DN_TEMPLATE:    o.AUTH_LDAP_1_USER_DN_TEMPLATE,
		AUTH_LDAP_1_USER_FLAGS_BY_GROUP: o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP,
		AUTH_LDAP_1_USER_SEARCH:         o.AUTH_LDAP_1_USER_SEARCH,
		AUTH_LDAP_2_BIND_DN:             o.AUTH_LDAP_2_BIND_DN,
		AUTH_LDAP_2_BIND_PASSWORD:       o.AUTH_LDAP_2_BIND_PASSWORD,
		AUTH_LDAP_2_CONNECTION_OPTIONS:  o.AUTH_LDAP_2_CONNECTION_OPTIONS,
		AUTH_LDAP_2_DENY_GROUP:          o.AUTH_LDAP_2_DENY_GROUP,
		AUTH_LDAP_2_GROUP_SEARCH:        o.AUTH_LDAP_2_GROUP_SEARCH,
		AUTH_LDAP_2_GROUP_TYPE:          o.AUTH_LDAP_2_GROUP_TYPE,
		AUTH_LDAP_2_GROUP_TYPE_PARAMS:   o.AUTH_LDAP_2_GROUP_TYPE_PARAMS,
		AUTH_LDAP_2_ORGANIZATION_MAP:    o.AUTH_LDAP_2_ORGANIZATION_MAP,
		AUTH_LDAP_2_REQUIRE_GROUP:       o.AUTH_LDAP_2_REQUIRE_GROUP,
		AUTH_LDAP_2_SERVER_URI:          o.AUTH_LDAP_2_SERVER_URI,
		AUTH_LDAP_2_START_TLS:           o.AUTH_LDAP_2_START_TLS,
		AUTH_LDAP_2_TEAM_MAP:            o.AUTH_LDAP_2_TEAM_MAP,
		AUTH_LDAP_2_USER_ATTR_MAP:       o.AUTH_LDAP_2_USER_ATTR_MAP,
		AUTH_LDAP_2_USER_DN_TEMPLATE:    o.AUTH_LDAP_2_USER_DN_TEMPLATE,
		AUTH_LDAP_2_USER_FLAGS_BY_GROUP: o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP,
		AUTH_LDAP_2_USER_SEARCH:         o.AUTH_LDAP_2_USER_SEARCH,
		AUTH_LDAP_3_BIND_DN:             o.AUTH_LDAP_3_BIND_DN,
		AUTH_LDAP_3_BIND_PASSWORD:       o.AUTH_LDAP_3_BIND_PASSWORD,
		AUTH_LDAP_3_CONNECTION_OPTIONS:  o.AUTH_LDAP_3_CONNECTION_OPTIONS,
		AUTH_LDAP_3_DENY_GROUP:          o.AUTH_LDAP_3_DENY_GROUP,
		AUTH_LDAP_3_GROUP_SEARCH:        o.AUTH_LDAP_3_GROUP_SEARCH,
		AUTH_LDAP_3_GROUP_TYPE:          o.AUTH_LDAP_3_GROUP_TYPE,
		AUTH_LDAP_3_GROUP_TYPE_PARAMS:   o.AUTH_LDAP_3_GROUP_TYPE_PARAMS,
		AUTH_LDAP_3_ORGANIZATION_MAP:    o.AUTH_LDAP_3_ORGANIZATION_MAP,
		AUTH_LDAP_3_REQUIRE_GROUP:       o.AUTH_LDAP_3_REQUIRE_GROUP,
		AUTH_LDAP_3_SERVER_URI:          o.AUTH_LDAP_3_SERVER_URI,
		AUTH_LDAP_3_START_TLS:           o.AUTH_LDAP_3_START_TLS,
		AUTH_LDAP_3_TEAM_MAP:            o.AUTH_LDAP_3_TEAM_MAP,
		AUTH_LDAP_3_USER_ATTR_MAP:       o.AUTH_LDAP_3_USER_ATTR_MAP,
		AUTH_LDAP_3_USER_DN_TEMPLATE:    o.AUTH_LDAP_3_USER_DN_TEMPLATE,
		AUTH_LDAP_3_USER_FLAGS_BY_GROUP: o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP,
		AUTH_LDAP_3_USER_SEARCH:         o.AUTH_LDAP_3_USER_SEARCH,
		AUTH_LDAP_4_BIND_DN:             o.AUTH_LDAP_4_BIND_DN,
		AUTH_LDAP_4_BIND_PASSWORD:       o.AUTH_LDAP_4_BIND_PASSWORD,
		AUTH_LDAP_4_CONNECTION_OPTIONS:  o.AUTH_LDAP_4_CONNECTION_OPTIONS,
		AUTH_LDAP_4_DENY_GROUP:          o.AUTH_LDAP_4_DENY_GROUP,
		AUTH_LDAP_4_GROUP_SEARCH:        o.AUTH_LDAP_4_GROUP_SEARCH,
		AUTH_LDAP_4_GROUP_TYPE:          o.AUTH_LDAP_4_GROUP_TYPE,
		AUTH_LDAP_4_GROUP_TYPE_PARAMS:   o.AUTH_LDAP_4_GROUP_TYPE_PARAMS,
		AUTH_LDAP_4_ORGANIZATION_MAP:    o.AUTH_LDAP_4_ORGANIZATION_MAP,
		AUTH_LDAP_4_REQUIRE_GROUP:       o.AUTH_LDAP_4_REQUIRE_GROUP,
		AUTH_LDAP_4_SERVER_URI:          o.AUTH_LDAP_4_SERVER_URI,
		AUTH_LDAP_4_START_TLS:           o.AUTH_LDAP_4_START_TLS,
		AUTH_LDAP_4_TEAM_MAP:            o.AUTH_LDAP_4_TEAM_MAP,
		AUTH_LDAP_4_USER_ATTR_MAP:       o.AUTH_LDAP_4_USER_ATTR_MAP,
		AUTH_LDAP_4_USER_DN_TEMPLATE:    o.AUTH_LDAP_4_USER_DN_TEMPLATE,
		AUTH_LDAP_4_USER_FLAGS_BY_GROUP: o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP,
		AUTH_LDAP_4_USER_SEARCH:         o.AUTH_LDAP_4_USER_SEARCH,
		AUTH_LDAP_5_BIND_DN:             o.AUTH_LDAP_5_BIND_DN,
		AUTH_LDAP_5_BIND_PASSWORD:       o.AUTH_LDAP_5_BIND_PASSWORD,
		AUTH_LDAP_5_CONNECTION_OPTIONS:  o.AUTH_LDAP_5_CONNECTION_OPTIONS,
		AUTH_LDAP_5_DENY_GROUP:          o.AUTH_LDAP_5_DENY_GROUP,
		AUTH_LDAP_5_GROUP_SEARCH:        o.AUTH_LDAP_5_GROUP_SEARCH,
		AUTH_LDAP_5_GROUP_TYPE:          o.AUTH_LDAP_5_GROUP_TYPE,
		AUTH_LDAP_5_GROUP_TYPE_PARAMS:   o.AUTH_LDAP_5_GROUP_TYPE_PARAMS,
		AUTH_LDAP_5_ORGANIZATION_MAP:    o.AUTH_LDAP_5_ORGANIZATION_MAP,
		AUTH_LDAP_5_REQUIRE_GROUP:       o.AUTH_LDAP_5_REQUIRE_GROUP,
		AUTH_LDAP_5_SERVER_URI:          o.AUTH_LDAP_5_SERVER_URI,
		AUTH_LDAP_5_START_TLS:           o.AUTH_LDAP_5_START_TLS,
		AUTH_LDAP_5_TEAM_MAP:            o.AUTH_LDAP_5_TEAM_MAP,
		AUTH_LDAP_5_USER_ATTR_MAP:       o.AUTH_LDAP_5_USER_ATTR_MAP,
		AUTH_LDAP_5_USER_DN_TEMPLATE:    o.AUTH_LDAP_5_USER_DN_TEMPLATE,
		AUTH_LDAP_5_USER_FLAGS_BY_GROUP: o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP,
		AUTH_LDAP_5_USER_SEARCH:         o.AUTH_LDAP_5_USER_SEARCH,
		AUTH_LDAP_BIND_DN:               o.AUTH_LDAP_BIND_DN,
		AUTH_LDAP_BIND_PASSWORD:         o.AUTH_LDAP_BIND_PASSWORD,
		AUTH_LDAP_CONNECTION_OPTIONS:    o.AUTH_LDAP_CONNECTION_OPTIONS,
		AUTH_LDAP_DENY_GROUP:            o.AUTH_LDAP_DENY_GROUP,
		AUTH_LDAP_GROUP_SEARCH:          o.AUTH_LDAP_GROUP_SEARCH,
		AUTH_LDAP_GROUP_TYPE:            o.AUTH_LDAP_GROUP_TYPE,
		AUTH_LDAP_GROUP_TYPE_PARAMS:     o.AUTH_LDAP_GROUP_TYPE_PARAMS,
		AUTH_LDAP_ORGANIZATION_MAP:      o.AUTH_LDAP_ORGANIZATION_MAP,
		AUTH_LDAP_REQUIRE_GROUP:         o.AUTH_LDAP_REQUIRE_GROUP,
		AUTH_LDAP_SERVER_URI:            o.AUTH_LDAP_SERVER_URI,
		AUTH_LDAP_START_TLS:             o.AUTH_LDAP_START_TLS,
		AUTH_LDAP_TEAM_MAP:              o.AUTH_LDAP_TEAM_MAP,
		AUTH_LDAP_USER_ATTR_MAP:         o.AUTH_LDAP_USER_ATTR_MAP,
		AUTH_LDAP_USER_DN_TEMPLATE:      o.AUTH_LDAP_USER_DN_TEMPLATE,
		AUTH_LDAP_USER_FLAGS_BY_GROUP:   o.AUTH_LDAP_USER_FLAGS_BY_GROUP,
		AUTH_LDAP_USER_SEARCH:           o.AUTH_LDAP_USER_SEARCH,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthLDAP
func (o *settingsAuthLdapTerraformModel) BodyRequest() (req settingsAuthLdapBodyRequestModel) {
	req.AUTH_LDAP_1_BIND_DN = o.AUTH_LDAP_1_BIND_DN.ValueString()
	req.AUTH_LDAP_1_BIND_PASSWORD = o.AUTH_LDAP_1_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_1_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_1_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_1_DENY_GROUP = o.AUTH_LDAP_1_DENY_GROUP.ValueString()
	req.AUTH_LDAP_1_GROUP_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_1_GROUP_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_1_GROUP_SEARCH = append(req.AUTH_LDAP_1_GROUP_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_1_GROUP_SEARCH = append(req.AUTH_LDAP_1_GROUP_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_1_GROUP_TYPE = o.AUTH_LDAP_1_GROUP_TYPE.ValueString()
	req.AUTH_LDAP_1_GROUP_TYPE_PARAMS = json.RawMessage(o.AUTH_LDAP_1_GROUP_TYPE_PARAMS.ValueString())
	req.AUTH_LDAP_1_ORGANIZATION_MAP = json.RawMessage(o.AUTH_LDAP_1_ORGANIZATION_MAP.ValueString())
	req.AUTH_LDAP_1_REQUIRE_GROUP = o.AUTH_LDAP_1_REQUIRE_GROUP.ValueString()
	req.AUTH_LDAP_1_SERVER_URI = o.AUTH_LDAP_1_SERVER_URI.ValueString()
	req.AUTH_LDAP_1_START_TLS = o.AUTH_LDAP_1_START_TLS.ValueBool()
	req.AUTH_LDAP_1_TEAM_MAP = json.RawMessage(o.AUTH_LDAP_1_TEAM_MAP.ValueString())
	req.AUTH_LDAP_1_USER_ATTR_MAP = json.RawMessage(o.AUTH_LDAP_1_USER_ATTR_MAP.ValueString())
	req.AUTH_LDAP_1_USER_DN_TEMPLATE = o.AUTH_LDAP_1_USER_DN_TEMPLATE.ValueString()
	req.AUTH_LDAP_1_USER_FLAGS_BY_GROUP = json.RawMessage(o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP.ValueString())
	req.AUTH_LDAP_1_USER_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_1_USER_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_1_USER_SEARCH = append(req.AUTH_LDAP_1_USER_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_1_USER_SEARCH = append(req.AUTH_LDAP_1_USER_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_2_BIND_DN = o.AUTH_LDAP_2_BIND_DN.ValueString()
	req.AUTH_LDAP_2_BIND_PASSWORD = o.AUTH_LDAP_2_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_2_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_2_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_2_DENY_GROUP = o.AUTH_LDAP_2_DENY_GROUP.ValueString()
	req.AUTH_LDAP_2_GROUP_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_2_GROUP_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_2_GROUP_SEARCH = append(req.AUTH_LDAP_2_GROUP_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_2_GROUP_SEARCH = append(req.AUTH_LDAP_2_GROUP_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_2_GROUP_TYPE = o.AUTH_LDAP_2_GROUP_TYPE.ValueString()
	req.AUTH_LDAP_2_GROUP_TYPE_PARAMS = json.RawMessage(o.AUTH_LDAP_2_GROUP_TYPE_PARAMS.ValueString())
	req.AUTH_LDAP_2_ORGANIZATION_MAP = json.RawMessage(o.AUTH_LDAP_2_ORGANIZATION_MAP.ValueString())
	req.AUTH_LDAP_2_REQUIRE_GROUP = o.AUTH_LDAP_2_REQUIRE_GROUP.ValueString()
	req.AUTH_LDAP_2_SERVER_URI = o.AUTH_LDAP_2_SERVER_URI.ValueString()
	req.AUTH_LDAP_2_START_TLS = o.AUTH_LDAP_2_START_TLS.ValueBool()
	req.AUTH_LDAP_2_TEAM_MAP = json.RawMessage(o.AUTH_LDAP_2_TEAM_MAP.ValueString())
	req.AUTH_LDAP_2_USER_ATTR_MAP = json.RawMessage(o.AUTH_LDAP_2_USER_ATTR_MAP.ValueString())
	req.AUTH_LDAP_2_USER_DN_TEMPLATE = o.AUTH_LDAP_2_USER_DN_TEMPLATE.ValueString()
	req.AUTH_LDAP_2_USER_FLAGS_BY_GROUP = json.RawMessage(o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP.ValueString())
	req.AUTH_LDAP_2_USER_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_2_USER_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_2_USER_SEARCH = append(req.AUTH_LDAP_2_USER_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_2_USER_SEARCH = append(req.AUTH_LDAP_2_USER_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_3_BIND_DN = o.AUTH_LDAP_3_BIND_DN.ValueString()
	req.AUTH_LDAP_3_BIND_PASSWORD = o.AUTH_LDAP_3_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_3_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_3_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_3_DENY_GROUP = o.AUTH_LDAP_3_DENY_GROUP.ValueString()
	req.AUTH_LDAP_3_GROUP_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_3_GROUP_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_3_GROUP_SEARCH = append(req.AUTH_LDAP_3_GROUP_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_3_GROUP_SEARCH = append(req.AUTH_LDAP_3_GROUP_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_3_GROUP_TYPE = o.AUTH_LDAP_3_GROUP_TYPE.ValueString()
	req.AUTH_LDAP_3_GROUP_TYPE_PARAMS = json.RawMessage(o.AUTH_LDAP_3_GROUP_TYPE_PARAMS.ValueString())
	req.AUTH_LDAP_3_ORGANIZATION_MAP = json.RawMessage(o.AUTH_LDAP_3_ORGANIZATION_MAP.ValueString())
	req.AUTH_LDAP_3_REQUIRE_GROUP = o.AUTH_LDAP_3_REQUIRE_GROUP.ValueString()
	req.AUTH_LDAP_3_SERVER_URI = o.AUTH_LDAP_3_SERVER_URI.ValueString()
	req.AUTH_LDAP_3_START_TLS = o.AUTH_LDAP_3_START_TLS.ValueBool()
	req.AUTH_LDAP_3_TEAM_MAP = json.RawMessage(o.AUTH_LDAP_3_TEAM_MAP.ValueString())
	req.AUTH_LDAP_3_USER_ATTR_MAP = json.RawMessage(o.AUTH_LDAP_3_USER_ATTR_MAP.ValueString())
	req.AUTH_LDAP_3_USER_DN_TEMPLATE = o.AUTH_LDAP_3_USER_DN_TEMPLATE.ValueString()
	req.AUTH_LDAP_3_USER_FLAGS_BY_GROUP = json.RawMessage(o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP.ValueString())
	req.AUTH_LDAP_3_USER_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_3_USER_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_3_USER_SEARCH = append(req.AUTH_LDAP_3_USER_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_3_USER_SEARCH = append(req.AUTH_LDAP_3_USER_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_4_BIND_DN = o.AUTH_LDAP_4_BIND_DN.ValueString()
	req.AUTH_LDAP_4_BIND_PASSWORD = o.AUTH_LDAP_4_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_4_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_4_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_4_DENY_GROUP = o.AUTH_LDAP_4_DENY_GROUP.ValueString()
	req.AUTH_LDAP_4_GROUP_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_4_GROUP_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_4_GROUP_SEARCH = append(req.AUTH_LDAP_4_GROUP_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_4_GROUP_SEARCH = append(req.AUTH_LDAP_4_GROUP_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_4_GROUP_TYPE = o.AUTH_LDAP_4_GROUP_TYPE.ValueString()
	req.AUTH_LDAP_4_GROUP_TYPE_PARAMS = json.RawMessage(o.AUTH_LDAP_4_GROUP_TYPE_PARAMS.ValueString())
	req.AUTH_LDAP_4_ORGANIZATION_MAP = json.RawMessage(o.AUTH_LDAP_4_ORGANIZATION_MAP.ValueString())
	req.AUTH_LDAP_4_REQUIRE_GROUP = o.AUTH_LDAP_4_REQUIRE_GROUP.ValueString()
	req.AUTH_LDAP_4_SERVER_URI = o.AUTH_LDAP_4_SERVER_URI.ValueString()
	req.AUTH_LDAP_4_START_TLS = o.AUTH_LDAP_4_START_TLS.ValueBool()
	req.AUTH_LDAP_4_TEAM_MAP = json.RawMessage(o.AUTH_LDAP_4_TEAM_MAP.ValueString())
	req.AUTH_LDAP_4_USER_ATTR_MAP = json.RawMessage(o.AUTH_LDAP_4_USER_ATTR_MAP.ValueString())
	req.AUTH_LDAP_4_USER_DN_TEMPLATE = o.AUTH_LDAP_4_USER_DN_TEMPLATE.ValueString()
	req.AUTH_LDAP_4_USER_FLAGS_BY_GROUP = json.RawMessage(o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP.ValueString())
	req.AUTH_LDAP_4_USER_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_4_USER_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_4_USER_SEARCH = append(req.AUTH_LDAP_4_USER_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_4_USER_SEARCH = append(req.AUTH_LDAP_4_USER_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_5_BIND_DN = o.AUTH_LDAP_5_BIND_DN.ValueString()
	req.AUTH_LDAP_5_BIND_PASSWORD = o.AUTH_LDAP_5_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_5_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_5_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_5_DENY_GROUP = o.AUTH_LDAP_5_DENY_GROUP.ValueString()
	req.AUTH_LDAP_5_GROUP_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_5_GROUP_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_5_GROUP_SEARCH = append(req.AUTH_LDAP_5_GROUP_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_5_GROUP_SEARCH = append(req.AUTH_LDAP_5_GROUP_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_5_GROUP_TYPE = o.AUTH_LDAP_5_GROUP_TYPE.ValueString()
	req.AUTH_LDAP_5_GROUP_TYPE_PARAMS = json.RawMessage(o.AUTH_LDAP_5_GROUP_TYPE_PARAMS.ValueString())
	req.AUTH_LDAP_5_ORGANIZATION_MAP = json.RawMessage(o.AUTH_LDAP_5_ORGANIZATION_MAP.ValueString())
	req.AUTH_LDAP_5_REQUIRE_GROUP = o.AUTH_LDAP_5_REQUIRE_GROUP.ValueString()
	req.AUTH_LDAP_5_SERVER_URI = o.AUTH_LDAP_5_SERVER_URI.ValueString()
	req.AUTH_LDAP_5_START_TLS = o.AUTH_LDAP_5_START_TLS.ValueBool()
	req.AUTH_LDAP_5_TEAM_MAP = json.RawMessage(o.AUTH_LDAP_5_TEAM_MAP.ValueString())
	req.AUTH_LDAP_5_USER_ATTR_MAP = json.RawMessage(o.AUTH_LDAP_5_USER_ATTR_MAP.ValueString())
	req.AUTH_LDAP_5_USER_DN_TEMPLATE = o.AUTH_LDAP_5_USER_DN_TEMPLATE.ValueString()
	req.AUTH_LDAP_5_USER_FLAGS_BY_GROUP = json.RawMessage(o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP.ValueString())
	req.AUTH_LDAP_5_USER_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_5_USER_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_5_USER_SEARCH = append(req.AUTH_LDAP_5_USER_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_5_USER_SEARCH = append(req.AUTH_LDAP_5_USER_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_BIND_DN = o.AUTH_LDAP_BIND_DN.ValueString()
	req.AUTH_LDAP_BIND_PASSWORD = o.AUTH_LDAP_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_DENY_GROUP = o.AUTH_LDAP_DENY_GROUP.ValueString()
	req.AUTH_LDAP_GROUP_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_GROUP_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_GROUP_SEARCH = append(req.AUTH_LDAP_GROUP_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_GROUP_SEARCH = append(req.AUTH_LDAP_GROUP_SEARCH, val.String())
		}
	}
	req.AUTH_LDAP_GROUP_TYPE = o.AUTH_LDAP_GROUP_TYPE.ValueString()
	req.AUTH_LDAP_GROUP_TYPE_PARAMS = json.RawMessage(o.AUTH_LDAP_GROUP_TYPE_PARAMS.ValueString())
	req.AUTH_LDAP_ORGANIZATION_MAP = json.RawMessage(o.AUTH_LDAP_ORGANIZATION_MAP.ValueString())
	req.AUTH_LDAP_REQUIRE_GROUP = o.AUTH_LDAP_REQUIRE_GROUP.ValueString()
	req.AUTH_LDAP_SERVER_URI = o.AUTH_LDAP_SERVER_URI.ValueString()
	req.AUTH_LDAP_START_TLS = o.AUTH_LDAP_START_TLS.ValueBool()
	req.AUTH_LDAP_TEAM_MAP = json.RawMessage(o.AUTH_LDAP_TEAM_MAP.ValueString())
	req.AUTH_LDAP_USER_ATTR_MAP = json.RawMessage(o.AUTH_LDAP_USER_ATTR_MAP.ValueString())
	req.AUTH_LDAP_USER_DN_TEMPLATE = o.AUTH_LDAP_USER_DN_TEMPLATE.ValueString()
	req.AUTH_LDAP_USER_FLAGS_BY_GROUP = json.RawMessage(o.AUTH_LDAP_USER_FLAGS_BY_GROUP.ValueString())
	req.AUTH_LDAP_USER_SEARCH = []string{}
	for _, val := range o.AUTH_LDAP_USER_SEARCH.Elements() {
		if _, ok := val.(types.String); ok {
			req.AUTH_LDAP_USER_SEARCH = append(req.AUTH_LDAP_USER_SEARCH, val.(types.String).ValueString())
		} else {
			req.AUTH_LDAP_USER_SEARCH = append(req.AUTH_LDAP_USER_SEARCH, val.String())
		}
	}
	return
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1BindDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_BIND_DN, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1BindPassword(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_BIND_PASSWORD, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_CONNECTION_OPTIONS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1DenyGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_DENY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1GroupSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_1_GROUP_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1GroupType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_GROUP_TYPE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_GROUP_TYPE_PARAMS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1RequireGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_REQUIRE_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1ServerUri(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_SERVER_URI, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1StartTls(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_LDAP_1_START_TLS, data)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_TEAM_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1UserAttrMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_USER_ATTR_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_1_USER_DN_TEMPLATE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap1UserSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_1_USER_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2BindDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_BIND_DN, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2BindPassword(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_BIND_PASSWORD, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_CONNECTION_OPTIONS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2DenyGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_DENY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2GroupSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_2_GROUP_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2GroupType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_GROUP_TYPE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_GROUP_TYPE_PARAMS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2RequireGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_REQUIRE_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2ServerUri(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_SERVER_URI, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2StartTls(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_LDAP_2_START_TLS, data)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_TEAM_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2UserAttrMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_USER_ATTR_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_2_USER_DN_TEMPLATE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap2UserSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_2_USER_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3BindDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_BIND_DN, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3BindPassword(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_BIND_PASSWORD, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_CONNECTION_OPTIONS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3DenyGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_DENY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3GroupSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_3_GROUP_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3GroupType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_GROUP_TYPE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_GROUP_TYPE_PARAMS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3RequireGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_REQUIRE_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3ServerUri(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_SERVER_URI, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3StartTls(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_LDAP_3_START_TLS, data)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_TEAM_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3UserAttrMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_USER_ATTR_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_3_USER_DN_TEMPLATE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap3UserSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_3_USER_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4BindDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_BIND_DN, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4BindPassword(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_BIND_PASSWORD, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_CONNECTION_OPTIONS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4DenyGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_DENY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4GroupSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_4_GROUP_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4GroupType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_GROUP_TYPE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_GROUP_TYPE_PARAMS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4RequireGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_REQUIRE_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4ServerUri(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_SERVER_URI, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4StartTls(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_LDAP_4_START_TLS, data)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_TEAM_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4UserAttrMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_USER_ATTR_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_4_USER_DN_TEMPLATE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap4UserSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_4_USER_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5BindDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_BIND_DN, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5BindPassword(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_BIND_PASSWORD, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_CONNECTION_OPTIONS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5DenyGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_DENY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5GroupSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_5_GROUP_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5GroupType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_GROUP_TYPE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_GROUP_TYPE_PARAMS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5RequireGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_REQUIRE_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5ServerUri(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_SERVER_URI, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5StartTls(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_LDAP_5_START_TLS, data)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_TEAM_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5UserAttrMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_USER_ATTR_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_5_USER_DN_TEMPLATE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdap5UserSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_5_USER_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapBindDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_BIND_DN, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapBindPassword(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_BIND_PASSWORD, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapConnectionOptions(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_CONNECTION_OPTIONS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapDenyGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_DENY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapGroupSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_GROUP_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapGroupType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_GROUP_TYPE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapGroupTypeParams(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_GROUP_TYPE_PARAMS, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapRequireGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_REQUIRE_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapServerUri(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_SERVER_URI, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapStartTls(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_LDAP_START_TLS, data)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_TEAM_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapUserAttrMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_USER_ATTR_MAP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapUserDnTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AUTH_LDAP_USER_DN_TEMPLATE, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapUserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.AUTH_LDAP_USER_FLAGS_BY_GROUP, data, false)
}

func (o *settingsAuthLdapTerraformModel) setAuthLdapUserSearch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTH_LDAP_USER_SEARCH, data, false)
}

func (o *settingsAuthLdapTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAuthLdap1BindDn(data["AUTH_LDAP_1_BIND_DN"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1BindPassword(data["AUTH_LDAP_1_BIND_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1ConnectionOptions(data["AUTH_LDAP_1_CONNECTION_OPTIONS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1DenyGroup(data["AUTH_LDAP_1_DENY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1GroupSearch(data["AUTH_LDAP_1_GROUP_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1GroupType(data["AUTH_LDAP_1_GROUP_TYPE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1GroupTypeParams(data["AUTH_LDAP_1_GROUP_TYPE_PARAMS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1OrganizationMap(data["AUTH_LDAP_1_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1RequireGroup(data["AUTH_LDAP_1_REQUIRE_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1ServerUri(data["AUTH_LDAP_1_SERVER_URI"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1StartTls(data["AUTH_LDAP_1_START_TLS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1TeamMap(data["AUTH_LDAP_1_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1UserAttrMap(data["AUTH_LDAP_1_USER_ATTR_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1UserDnTemplate(data["AUTH_LDAP_1_USER_DN_TEMPLATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1UserFlagsByGroup(data["AUTH_LDAP_1_USER_FLAGS_BY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap1UserSearch(data["AUTH_LDAP_1_USER_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2BindDn(data["AUTH_LDAP_2_BIND_DN"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2BindPassword(data["AUTH_LDAP_2_BIND_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2ConnectionOptions(data["AUTH_LDAP_2_CONNECTION_OPTIONS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2DenyGroup(data["AUTH_LDAP_2_DENY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2GroupSearch(data["AUTH_LDAP_2_GROUP_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2GroupType(data["AUTH_LDAP_2_GROUP_TYPE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2GroupTypeParams(data["AUTH_LDAP_2_GROUP_TYPE_PARAMS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2OrganizationMap(data["AUTH_LDAP_2_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2RequireGroup(data["AUTH_LDAP_2_REQUIRE_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2ServerUri(data["AUTH_LDAP_2_SERVER_URI"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2StartTls(data["AUTH_LDAP_2_START_TLS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2TeamMap(data["AUTH_LDAP_2_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2UserAttrMap(data["AUTH_LDAP_2_USER_ATTR_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2UserDnTemplate(data["AUTH_LDAP_2_USER_DN_TEMPLATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2UserFlagsByGroup(data["AUTH_LDAP_2_USER_FLAGS_BY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap2UserSearch(data["AUTH_LDAP_2_USER_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3BindDn(data["AUTH_LDAP_3_BIND_DN"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3BindPassword(data["AUTH_LDAP_3_BIND_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3ConnectionOptions(data["AUTH_LDAP_3_CONNECTION_OPTIONS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3DenyGroup(data["AUTH_LDAP_3_DENY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3GroupSearch(data["AUTH_LDAP_3_GROUP_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3GroupType(data["AUTH_LDAP_3_GROUP_TYPE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3GroupTypeParams(data["AUTH_LDAP_3_GROUP_TYPE_PARAMS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3OrganizationMap(data["AUTH_LDAP_3_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3RequireGroup(data["AUTH_LDAP_3_REQUIRE_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3ServerUri(data["AUTH_LDAP_3_SERVER_URI"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3StartTls(data["AUTH_LDAP_3_START_TLS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3TeamMap(data["AUTH_LDAP_3_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3UserAttrMap(data["AUTH_LDAP_3_USER_ATTR_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3UserDnTemplate(data["AUTH_LDAP_3_USER_DN_TEMPLATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3UserFlagsByGroup(data["AUTH_LDAP_3_USER_FLAGS_BY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap3UserSearch(data["AUTH_LDAP_3_USER_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4BindDn(data["AUTH_LDAP_4_BIND_DN"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4BindPassword(data["AUTH_LDAP_4_BIND_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4ConnectionOptions(data["AUTH_LDAP_4_CONNECTION_OPTIONS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4DenyGroup(data["AUTH_LDAP_4_DENY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4GroupSearch(data["AUTH_LDAP_4_GROUP_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4GroupType(data["AUTH_LDAP_4_GROUP_TYPE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4GroupTypeParams(data["AUTH_LDAP_4_GROUP_TYPE_PARAMS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4OrganizationMap(data["AUTH_LDAP_4_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4RequireGroup(data["AUTH_LDAP_4_REQUIRE_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4ServerUri(data["AUTH_LDAP_4_SERVER_URI"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4StartTls(data["AUTH_LDAP_4_START_TLS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4TeamMap(data["AUTH_LDAP_4_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4UserAttrMap(data["AUTH_LDAP_4_USER_ATTR_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4UserDnTemplate(data["AUTH_LDAP_4_USER_DN_TEMPLATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4UserFlagsByGroup(data["AUTH_LDAP_4_USER_FLAGS_BY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap4UserSearch(data["AUTH_LDAP_4_USER_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5BindDn(data["AUTH_LDAP_5_BIND_DN"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5BindPassword(data["AUTH_LDAP_5_BIND_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5ConnectionOptions(data["AUTH_LDAP_5_CONNECTION_OPTIONS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5DenyGroup(data["AUTH_LDAP_5_DENY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5GroupSearch(data["AUTH_LDAP_5_GROUP_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5GroupType(data["AUTH_LDAP_5_GROUP_TYPE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5GroupTypeParams(data["AUTH_LDAP_5_GROUP_TYPE_PARAMS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5OrganizationMap(data["AUTH_LDAP_5_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5RequireGroup(data["AUTH_LDAP_5_REQUIRE_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5ServerUri(data["AUTH_LDAP_5_SERVER_URI"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5StartTls(data["AUTH_LDAP_5_START_TLS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5TeamMap(data["AUTH_LDAP_5_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5UserAttrMap(data["AUTH_LDAP_5_USER_ATTR_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5UserDnTemplate(data["AUTH_LDAP_5_USER_DN_TEMPLATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5UserFlagsByGroup(data["AUTH_LDAP_5_USER_FLAGS_BY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdap5UserSearch(data["AUTH_LDAP_5_USER_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapBindDn(data["AUTH_LDAP_BIND_DN"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapBindPassword(data["AUTH_LDAP_BIND_PASSWORD"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapConnectionOptions(data["AUTH_LDAP_CONNECTION_OPTIONS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapDenyGroup(data["AUTH_LDAP_DENY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapGroupSearch(data["AUTH_LDAP_GROUP_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapGroupType(data["AUTH_LDAP_GROUP_TYPE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapGroupTypeParams(data["AUTH_LDAP_GROUP_TYPE_PARAMS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapOrganizationMap(data["AUTH_LDAP_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapRequireGroup(data["AUTH_LDAP_REQUIRE_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapServerUri(data["AUTH_LDAP_SERVER_URI"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapStartTls(data["AUTH_LDAP_START_TLS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapTeamMap(data["AUTH_LDAP_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapUserAttrMap(data["AUTH_LDAP_USER_ATTR_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapUserDnTemplate(data["AUTH_LDAP_USER_DN_TEMPLATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapUserFlagsByGroup(data["AUTH_LDAP_USER_FLAGS_BY_GROUP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthLdapUserSearch(data["AUTH_LDAP_USER_SEARCH"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthLdapBodyRequestModel maps the schema for SettingsAuthLDAP for creating and updating the data
type settingsAuthLdapBodyRequestModel struct {
	// AUTH_LDAP_1_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_1_BIND_DN string `json:"AUTH_LDAP_1_BIND_DN,omitempty"`
	// AUTH_LDAP_1_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_1_BIND_PASSWORD string `json:"AUTH_LDAP_1_BIND_PASSWORD,omitempty"`
	// AUTH_LDAP_1_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_1_CONNECTION_OPTIONS json.RawMessage `json:"AUTH_LDAP_1_CONNECTION_OPTIONS,omitempty"`
	// AUTH_LDAP_1_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_1_DENY_GROUP string `json:"AUTH_LDAP_1_DENY_GROUP,omitempty"`
	// AUTH_LDAP_1_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_1_GROUP_SEARCH []string `json:"AUTH_LDAP_1_GROUP_SEARCH,omitempty"`
	// AUTH_LDAP_1_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_1_GROUP_TYPE string `json:"AUTH_LDAP_1_GROUP_TYPE,omitempty"`
	// AUTH_LDAP_1_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_1_GROUP_TYPE_PARAMS json.RawMessage `json:"AUTH_LDAP_1_GROUP_TYPE_PARAMS,omitempty"`
	// AUTH_LDAP_1_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_1_ORGANIZATION_MAP json.RawMessage `json:"AUTH_LDAP_1_ORGANIZATION_MAP,omitempty"`
	// AUTH_LDAP_1_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_1_REQUIRE_GROUP string `json:"AUTH_LDAP_1_REQUIRE_GROUP,omitempty"`
	// AUTH_LDAP_1_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_1_SERVER_URI string `json:"AUTH_LDAP_1_SERVER_URI,omitempty"`
	// AUTH_LDAP_1_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_1_START_TLS bool `json:"AUTH_LDAP_1_START_TLS"`
	// AUTH_LDAP_1_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_1_TEAM_MAP json.RawMessage `json:"AUTH_LDAP_1_TEAM_MAP,omitempty"`
	// AUTH_LDAP_1_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_1_USER_ATTR_MAP json.RawMessage `json:"AUTH_LDAP_1_USER_ATTR_MAP,omitempty"`
	// AUTH_LDAP_1_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_1_USER_DN_TEMPLATE string `json:"AUTH_LDAP_1_USER_DN_TEMPLATE,omitempty"`
	// AUTH_LDAP_1_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_1_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_1_USER_FLAGS_BY_GROUP,omitempty"`
	// AUTH_LDAP_1_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_1_USER_SEARCH []string `json:"AUTH_LDAP_1_USER_SEARCH,omitempty"`
	// AUTH_LDAP_2_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_2_BIND_DN string `json:"AUTH_LDAP_2_BIND_DN,omitempty"`
	// AUTH_LDAP_2_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_2_BIND_PASSWORD string `json:"AUTH_LDAP_2_BIND_PASSWORD,omitempty"`
	// AUTH_LDAP_2_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_2_CONNECTION_OPTIONS json.RawMessage `json:"AUTH_LDAP_2_CONNECTION_OPTIONS,omitempty"`
	// AUTH_LDAP_2_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_2_DENY_GROUP string `json:"AUTH_LDAP_2_DENY_GROUP,omitempty"`
	// AUTH_LDAP_2_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_2_GROUP_SEARCH []string `json:"AUTH_LDAP_2_GROUP_SEARCH,omitempty"`
	// AUTH_LDAP_2_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_2_GROUP_TYPE string `json:"AUTH_LDAP_2_GROUP_TYPE,omitempty"`
	// AUTH_LDAP_2_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_2_GROUP_TYPE_PARAMS json.RawMessage `json:"AUTH_LDAP_2_GROUP_TYPE_PARAMS,omitempty"`
	// AUTH_LDAP_2_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_2_ORGANIZATION_MAP json.RawMessage `json:"AUTH_LDAP_2_ORGANIZATION_MAP,omitempty"`
	// AUTH_LDAP_2_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_2_REQUIRE_GROUP string `json:"AUTH_LDAP_2_REQUIRE_GROUP,omitempty"`
	// AUTH_LDAP_2_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_2_SERVER_URI string `json:"AUTH_LDAP_2_SERVER_URI,omitempty"`
	// AUTH_LDAP_2_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_2_START_TLS bool `json:"AUTH_LDAP_2_START_TLS"`
	// AUTH_LDAP_2_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_2_TEAM_MAP json.RawMessage `json:"AUTH_LDAP_2_TEAM_MAP,omitempty"`
	// AUTH_LDAP_2_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_2_USER_ATTR_MAP json.RawMessage `json:"AUTH_LDAP_2_USER_ATTR_MAP,omitempty"`
	// AUTH_LDAP_2_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_2_USER_DN_TEMPLATE string `json:"AUTH_LDAP_2_USER_DN_TEMPLATE,omitempty"`
	// AUTH_LDAP_2_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_2_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_2_USER_FLAGS_BY_GROUP,omitempty"`
	// AUTH_LDAP_2_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_2_USER_SEARCH []string `json:"AUTH_LDAP_2_USER_SEARCH,omitempty"`
	// AUTH_LDAP_3_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_3_BIND_DN string `json:"AUTH_LDAP_3_BIND_DN,omitempty"`
	// AUTH_LDAP_3_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_3_BIND_PASSWORD string `json:"AUTH_LDAP_3_BIND_PASSWORD,omitempty"`
	// AUTH_LDAP_3_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_3_CONNECTION_OPTIONS json.RawMessage `json:"AUTH_LDAP_3_CONNECTION_OPTIONS,omitempty"`
	// AUTH_LDAP_3_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_3_DENY_GROUP string `json:"AUTH_LDAP_3_DENY_GROUP,omitempty"`
	// AUTH_LDAP_3_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_3_GROUP_SEARCH []string `json:"AUTH_LDAP_3_GROUP_SEARCH,omitempty"`
	// AUTH_LDAP_3_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_3_GROUP_TYPE string `json:"AUTH_LDAP_3_GROUP_TYPE,omitempty"`
	// AUTH_LDAP_3_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_3_GROUP_TYPE_PARAMS json.RawMessage `json:"AUTH_LDAP_3_GROUP_TYPE_PARAMS,omitempty"`
	// AUTH_LDAP_3_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_3_ORGANIZATION_MAP json.RawMessage `json:"AUTH_LDAP_3_ORGANIZATION_MAP,omitempty"`
	// AUTH_LDAP_3_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_3_REQUIRE_GROUP string `json:"AUTH_LDAP_3_REQUIRE_GROUP,omitempty"`
	// AUTH_LDAP_3_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_3_SERVER_URI string `json:"AUTH_LDAP_3_SERVER_URI,omitempty"`
	// AUTH_LDAP_3_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_3_START_TLS bool `json:"AUTH_LDAP_3_START_TLS"`
	// AUTH_LDAP_3_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_3_TEAM_MAP json.RawMessage `json:"AUTH_LDAP_3_TEAM_MAP,omitempty"`
	// AUTH_LDAP_3_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_3_USER_ATTR_MAP json.RawMessage `json:"AUTH_LDAP_3_USER_ATTR_MAP,omitempty"`
	// AUTH_LDAP_3_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_3_USER_DN_TEMPLATE string `json:"AUTH_LDAP_3_USER_DN_TEMPLATE,omitempty"`
	// AUTH_LDAP_3_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_3_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_3_USER_FLAGS_BY_GROUP,omitempty"`
	// AUTH_LDAP_3_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_3_USER_SEARCH []string `json:"AUTH_LDAP_3_USER_SEARCH,omitempty"`
	// AUTH_LDAP_4_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_4_BIND_DN string `json:"AUTH_LDAP_4_BIND_DN,omitempty"`
	// AUTH_LDAP_4_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_4_BIND_PASSWORD string `json:"AUTH_LDAP_4_BIND_PASSWORD,omitempty"`
	// AUTH_LDAP_4_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_4_CONNECTION_OPTIONS json.RawMessage `json:"AUTH_LDAP_4_CONNECTION_OPTIONS,omitempty"`
	// AUTH_LDAP_4_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_4_DENY_GROUP string `json:"AUTH_LDAP_4_DENY_GROUP,omitempty"`
	// AUTH_LDAP_4_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_4_GROUP_SEARCH []string `json:"AUTH_LDAP_4_GROUP_SEARCH,omitempty"`
	// AUTH_LDAP_4_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_4_GROUP_TYPE string `json:"AUTH_LDAP_4_GROUP_TYPE,omitempty"`
	// AUTH_LDAP_4_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_4_GROUP_TYPE_PARAMS json.RawMessage `json:"AUTH_LDAP_4_GROUP_TYPE_PARAMS,omitempty"`
	// AUTH_LDAP_4_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_4_ORGANIZATION_MAP json.RawMessage `json:"AUTH_LDAP_4_ORGANIZATION_MAP,omitempty"`
	// AUTH_LDAP_4_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_4_REQUIRE_GROUP string `json:"AUTH_LDAP_4_REQUIRE_GROUP,omitempty"`
	// AUTH_LDAP_4_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_4_SERVER_URI string `json:"AUTH_LDAP_4_SERVER_URI,omitempty"`
	// AUTH_LDAP_4_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_4_START_TLS bool `json:"AUTH_LDAP_4_START_TLS"`
	// AUTH_LDAP_4_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_4_TEAM_MAP json.RawMessage `json:"AUTH_LDAP_4_TEAM_MAP,omitempty"`
	// AUTH_LDAP_4_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_4_USER_ATTR_MAP json.RawMessage `json:"AUTH_LDAP_4_USER_ATTR_MAP,omitempty"`
	// AUTH_LDAP_4_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_4_USER_DN_TEMPLATE string `json:"AUTH_LDAP_4_USER_DN_TEMPLATE,omitempty"`
	// AUTH_LDAP_4_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_4_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_4_USER_FLAGS_BY_GROUP,omitempty"`
	// AUTH_LDAP_4_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_4_USER_SEARCH []string `json:"AUTH_LDAP_4_USER_SEARCH,omitempty"`
	// AUTH_LDAP_5_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_5_BIND_DN string `json:"AUTH_LDAP_5_BIND_DN,omitempty"`
	// AUTH_LDAP_5_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_5_BIND_PASSWORD string `json:"AUTH_LDAP_5_BIND_PASSWORD,omitempty"`
	// AUTH_LDAP_5_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_5_CONNECTION_OPTIONS json.RawMessage `json:"AUTH_LDAP_5_CONNECTION_OPTIONS,omitempty"`
	// AUTH_LDAP_5_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_5_DENY_GROUP string `json:"AUTH_LDAP_5_DENY_GROUP,omitempty"`
	// AUTH_LDAP_5_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_5_GROUP_SEARCH []string `json:"AUTH_LDAP_5_GROUP_SEARCH,omitempty"`
	// AUTH_LDAP_5_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_5_GROUP_TYPE string `json:"AUTH_LDAP_5_GROUP_TYPE,omitempty"`
	// AUTH_LDAP_5_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_5_GROUP_TYPE_PARAMS json.RawMessage `json:"AUTH_LDAP_5_GROUP_TYPE_PARAMS,omitempty"`
	// AUTH_LDAP_5_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_5_ORGANIZATION_MAP json.RawMessage `json:"AUTH_LDAP_5_ORGANIZATION_MAP,omitempty"`
	// AUTH_LDAP_5_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_5_REQUIRE_GROUP string `json:"AUTH_LDAP_5_REQUIRE_GROUP,omitempty"`
	// AUTH_LDAP_5_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_5_SERVER_URI string `json:"AUTH_LDAP_5_SERVER_URI,omitempty"`
	// AUTH_LDAP_5_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_5_START_TLS bool `json:"AUTH_LDAP_5_START_TLS"`
	// AUTH_LDAP_5_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_5_TEAM_MAP json.RawMessage `json:"AUTH_LDAP_5_TEAM_MAP,omitempty"`
	// AUTH_LDAP_5_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_5_USER_ATTR_MAP json.RawMessage `json:"AUTH_LDAP_5_USER_ATTR_MAP,omitempty"`
	// AUTH_LDAP_5_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_5_USER_DN_TEMPLATE string `json:"AUTH_LDAP_5_USER_DN_TEMPLATE,omitempty"`
	// AUTH_LDAP_5_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_5_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_5_USER_FLAGS_BY_GROUP,omitempty"`
	// AUTH_LDAP_5_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_5_USER_SEARCH []string `json:"AUTH_LDAP_5_USER_SEARCH,omitempty"`
	// AUTH_LDAP_BIND_DN "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax."
	AUTH_LDAP_BIND_DN string `json:"AUTH_LDAP_BIND_DN,omitempty"`
	// AUTH_LDAP_BIND_PASSWORD "Password used to bind LDAP user account."
	AUTH_LDAP_BIND_PASSWORD string `json:"AUTH_LDAP_BIND_PASSWORD,omitempty"`
	// AUTH_LDAP_CONNECTION_OPTIONS "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set."
	AUTH_LDAP_CONNECTION_OPTIONS json.RawMessage `json:"AUTH_LDAP_CONNECTION_OPTIONS,omitempty"`
	// AUTH_LDAP_DENY_GROUP "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported."
	AUTH_LDAP_DENY_GROUP string `json:"AUTH_LDAP_DENY_GROUP,omitempty"`
	// AUTH_LDAP_GROUP_SEARCH "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion."
	AUTH_LDAP_GROUP_SEARCH []string `json:"AUTH_LDAP_GROUP_SEARCH,omitempty"`
	// AUTH_LDAP_GROUP_TYPE "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups"
	AUTH_LDAP_GROUP_TYPE string `json:"AUTH_LDAP_GROUP_TYPE,omitempty"`
	// AUTH_LDAP_GROUP_TYPE_PARAMS "Key value parameters to send the chosen group type init method."
	AUTH_LDAP_GROUP_TYPE_PARAMS json.RawMessage `json:"AUTH_LDAP_GROUP_TYPE_PARAMS,omitempty"`
	// AUTH_LDAP_ORGANIZATION_MAP "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation."
	AUTH_LDAP_ORGANIZATION_MAP json.RawMessage `json:"AUTH_LDAP_ORGANIZATION_MAP,omitempty"`
	// AUTH_LDAP_REQUIRE_GROUP "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported."
	AUTH_LDAP_REQUIRE_GROUP string `json:"AUTH_LDAP_REQUIRE_GROUP,omitempty"`
	// AUTH_LDAP_SERVER_URI "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty."
	AUTH_LDAP_SERVER_URI string `json:"AUTH_LDAP_SERVER_URI,omitempty"`
	// AUTH_LDAP_START_TLS "Whether to enable TLS when the LDAP connection is not using SSL."
	AUTH_LDAP_START_TLS bool `json:"AUTH_LDAP_START_TLS"`
	// AUTH_LDAP_TEAM_MAP "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation."
	AUTH_LDAP_TEAM_MAP json.RawMessage `json:"AUTH_LDAP_TEAM_MAP,omitempty"`
	// AUTH_LDAP_USER_ATTR_MAP "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details."
	AUTH_LDAP_USER_ATTR_MAP json.RawMessage `json:"AUTH_LDAP_USER_ATTR_MAP,omitempty"`
	// AUTH_LDAP_USER_DN_TEMPLATE "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH."
	AUTH_LDAP_USER_DN_TEMPLATE string `json:"AUTH_LDAP_USER_DN_TEMPLATE,omitempty"`
	// AUTH_LDAP_USER_FLAGS_BY_GROUP "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail."
	AUTH_LDAP_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_USER_FLAGS_BY_GROUP,omitempty"`
	// AUTH_LDAP_USER_SEARCH "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details."
	AUTH_LDAP_USER_SEARCH []string `json:"AUTH_LDAP_USER_SEARCH,omitempty"`
}
