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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthLDAP
func (o *settingsAuthLdapTerraformModel) BodyRequest() *settingsAuthLdapBodyRequestModel {
	var req settingsAuthLdapBodyRequestModel
	req.AUTH_LDAP_1_BIND_DN = o.AUTH_LDAP_1_BIND_DN.ValueString()
	req.AUTH_LDAP_1_BIND_PASSWORD = o.AUTH_LDAP_1_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_1_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_1_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_1_DENY_GROUP = o.AUTH_LDAP_1_DENY_GROUP.ValueString()
	req.AUTH_LDAP_1_GROUP_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_1_GROUP_SEARCH, false)
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
	req.AUTH_LDAP_1_USER_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_1_USER_SEARCH, false)
	req.AUTH_LDAP_2_BIND_DN = o.AUTH_LDAP_2_BIND_DN.ValueString()
	req.AUTH_LDAP_2_BIND_PASSWORD = o.AUTH_LDAP_2_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_2_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_2_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_2_DENY_GROUP = o.AUTH_LDAP_2_DENY_GROUP.ValueString()
	req.AUTH_LDAP_2_GROUP_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_2_GROUP_SEARCH, false)
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
	req.AUTH_LDAP_2_USER_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_2_USER_SEARCH, false)
	req.AUTH_LDAP_3_BIND_DN = o.AUTH_LDAP_3_BIND_DN.ValueString()
	req.AUTH_LDAP_3_BIND_PASSWORD = o.AUTH_LDAP_3_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_3_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_3_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_3_DENY_GROUP = o.AUTH_LDAP_3_DENY_GROUP.ValueString()
	req.AUTH_LDAP_3_GROUP_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_3_GROUP_SEARCH, false)
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
	req.AUTH_LDAP_3_USER_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_3_USER_SEARCH, false)
	req.AUTH_LDAP_4_BIND_DN = o.AUTH_LDAP_4_BIND_DN.ValueString()
	req.AUTH_LDAP_4_BIND_PASSWORD = o.AUTH_LDAP_4_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_4_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_4_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_4_DENY_GROUP = o.AUTH_LDAP_4_DENY_GROUP.ValueString()
	req.AUTH_LDAP_4_GROUP_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_4_GROUP_SEARCH, false)
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
	req.AUTH_LDAP_4_USER_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_4_USER_SEARCH, false)
	req.AUTH_LDAP_5_BIND_DN = o.AUTH_LDAP_5_BIND_DN.ValueString()
	req.AUTH_LDAP_5_BIND_PASSWORD = o.AUTH_LDAP_5_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_5_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_5_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_5_DENY_GROUP = o.AUTH_LDAP_5_DENY_GROUP.ValueString()
	req.AUTH_LDAP_5_GROUP_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_5_GROUP_SEARCH, false)
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
	req.AUTH_LDAP_5_USER_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_5_USER_SEARCH, false)
	req.AUTH_LDAP_BIND_DN = o.AUTH_LDAP_BIND_DN.ValueString()
	req.AUTH_LDAP_BIND_PASSWORD = o.AUTH_LDAP_BIND_PASSWORD.ValueString()
	req.AUTH_LDAP_CONNECTION_OPTIONS = json.RawMessage(o.AUTH_LDAP_CONNECTION_OPTIONS.ValueString())
	req.AUTH_LDAP_DENY_GROUP = o.AUTH_LDAP_DENY_GROUP.ValueString()
	req.AUTH_LDAP_GROUP_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_GROUP_SEARCH, false)
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
	req.AUTH_LDAP_USER_SEARCH = helpers.ListAsStringSlice(o.AUTH_LDAP_USER_SEARCH, false)
	return &req
}

func (o *settingsAuthLdapTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_BIND_DN, data["AUTH_LDAP_1_BIND_DN"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_BIND_PASSWORD, data["AUTH_LDAP_1_BIND_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_CONNECTION_OPTIONS, data["AUTH_LDAP_1_CONNECTION_OPTIONS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_DENY_GROUP, data["AUTH_LDAP_1_DENY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_1_GROUP_SEARCH, data["AUTH_LDAP_1_GROUP_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_GROUP_TYPE, data["AUTH_LDAP_1_GROUP_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_GROUP_TYPE_PARAMS, data["AUTH_LDAP_1_GROUP_TYPE_PARAMS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_ORGANIZATION_MAP, data["AUTH_LDAP_1_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_REQUIRE_GROUP, data["AUTH_LDAP_1_REQUIRE_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_SERVER_URI, data["AUTH_LDAP_1_SERVER_URI"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_LDAP_1_START_TLS, data["AUTH_LDAP_1_START_TLS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_TEAM_MAP, data["AUTH_LDAP_1_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_USER_ATTR_MAP, data["AUTH_LDAP_1_USER_ATTR_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_1_USER_DN_TEMPLATE, data["AUTH_LDAP_1_USER_DN_TEMPLATE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_1_USER_FLAGS_BY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_1_USER_SEARCH, data["AUTH_LDAP_1_USER_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_BIND_DN, data["AUTH_LDAP_2_BIND_DN"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_BIND_PASSWORD, data["AUTH_LDAP_2_BIND_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_CONNECTION_OPTIONS, data["AUTH_LDAP_2_CONNECTION_OPTIONS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_DENY_GROUP, data["AUTH_LDAP_2_DENY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_2_GROUP_SEARCH, data["AUTH_LDAP_2_GROUP_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_GROUP_TYPE, data["AUTH_LDAP_2_GROUP_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_GROUP_TYPE_PARAMS, data["AUTH_LDAP_2_GROUP_TYPE_PARAMS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_ORGANIZATION_MAP, data["AUTH_LDAP_2_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_REQUIRE_GROUP, data["AUTH_LDAP_2_REQUIRE_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_SERVER_URI, data["AUTH_LDAP_2_SERVER_URI"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_LDAP_2_START_TLS, data["AUTH_LDAP_2_START_TLS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_TEAM_MAP, data["AUTH_LDAP_2_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_USER_ATTR_MAP, data["AUTH_LDAP_2_USER_ATTR_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_2_USER_DN_TEMPLATE, data["AUTH_LDAP_2_USER_DN_TEMPLATE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_2_USER_FLAGS_BY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_2_USER_SEARCH, data["AUTH_LDAP_2_USER_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_BIND_DN, data["AUTH_LDAP_3_BIND_DN"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_BIND_PASSWORD, data["AUTH_LDAP_3_BIND_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_CONNECTION_OPTIONS, data["AUTH_LDAP_3_CONNECTION_OPTIONS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_DENY_GROUP, data["AUTH_LDAP_3_DENY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_3_GROUP_SEARCH, data["AUTH_LDAP_3_GROUP_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_GROUP_TYPE, data["AUTH_LDAP_3_GROUP_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_GROUP_TYPE_PARAMS, data["AUTH_LDAP_3_GROUP_TYPE_PARAMS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_ORGANIZATION_MAP, data["AUTH_LDAP_3_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_REQUIRE_GROUP, data["AUTH_LDAP_3_REQUIRE_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_SERVER_URI, data["AUTH_LDAP_3_SERVER_URI"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_LDAP_3_START_TLS, data["AUTH_LDAP_3_START_TLS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_TEAM_MAP, data["AUTH_LDAP_3_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_USER_ATTR_MAP, data["AUTH_LDAP_3_USER_ATTR_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_3_USER_DN_TEMPLATE, data["AUTH_LDAP_3_USER_DN_TEMPLATE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_3_USER_FLAGS_BY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_3_USER_SEARCH, data["AUTH_LDAP_3_USER_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_BIND_DN, data["AUTH_LDAP_4_BIND_DN"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_BIND_PASSWORD, data["AUTH_LDAP_4_BIND_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_CONNECTION_OPTIONS, data["AUTH_LDAP_4_CONNECTION_OPTIONS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_DENY_GROUP, data["AUTH_LDAP_4_DENY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_4_GROUP_SEARCH, data["AUTH_LDAP_4_GROUP_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_GROUP_TYPE, data["AUTH_LDAP_4_GROUP_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_GROUP_TYPE_PARAMS, data["AUTH_LDAP_4_GROUP_TYPE_PARAMS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_ORGANIZATION_MAP, data["AUTH_LDAP_4_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_REQUIRE_GROUP, data["AUTH_LDAP_4_REQUIRE_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_SERVER_URI, data["AUTH_LDAP_4_SERVER_URI"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_LDAP_4_START_TLS, data["AUTH_LDAP_4_START_TLS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_TEAM_MAP, data["AUTH_LDAP_4_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_USER_ATTR_MAP, data["AUTH_LDAP_4_USER_ATTR_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_4_USER_DN_TEMPLATE, data["AUTH_LDAP_4_USER_DN_TEMPLATE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_4_USER_FLAGS_BY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_4_USER_SEARCH, data["AUTH_LDAP_4_USER_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_BIND_DN, data["AUTH_LDAP_5_BIND_DN"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_BIND_PASSWORD, data["AUTH_LDAP_5_BIND_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_CONNECTION_OPTIONS, data["AUTH_LDAP_5_CONNECTION_OPTIONS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_DENY_GROUP, data["AUTH_LDAP_5_DENY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_5_GROUP_SEARCH, data["AUTH_LDAP_5_GROUP_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_GROUP_TYPE, data["AUTH_LDAP_5_GROUP_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_GROUP_TYPE_PARAMS, data["AUTH_LDAP_5_GROUP_TYPE_PARAMS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_ORGANIZATION_MAP, data["AUTH_LDAP_5_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_REQUIRE_GROUP, data["AUTH_LDAP_5_REQUIRE_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_SERVER_URI, data["AUTH_LDAP_5_SERVER_URI"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_LDAP_5_START_TLS, data["AUTH_LDAP_5_START_TLS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_TEAM_MAP, data["AUTH_LDAP_5_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_USER_ATTR_MAP, data["AUTH_LDAP_5_USER_ATTR_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_5_USER_DN_TEMPLATE, data["AUTH_LDAP_5_USER_DN_TEMPLATE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_5_USER_FLAGS_BY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_5_USER_SEARCH, data["AUTH_LDAP_5_USER_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_BIND_DN, data["AUTH_LDAP_BIND_DN"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_BIND_PASSWORD, data["AUTH_LDAP_BIND_PASSWORD"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_CONNECTION_OPTIONS, data["AUTH_LDAP_CONNECTION_OPTIONS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_DENY_GROUP, data["AUTH_LDAP_DENY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_GROUP_SEARCH, data["AUTH_LDAP_GROUP_SEARCH"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_GROUP_TYPE, data["AUTH_LDAP_GROUP_TYPE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_GROUP_TYPE_PARAMS, data["AUTH_LDAP_GROUP_TYPE_PARAMS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_ORGANIZATION_MAP, data["AUTH_LDAP_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_REQUIRE_GROUP, data["AUTH_LDAP_REQUIRE_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_SERVER_URI, data["AUTH_LDAP_SERVER_URI"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_LDAP_START_TLS, data["AUTH_LDAP_START_TLS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_TEAM_MAP, data["AUTH_LDAP_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_USER_ATTR_MAP, data["AUTH_LDAP_USER_ATTR_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AUTH_LDAP_USER_DN_TEMPLATE, data["AUTH_LDAP_USER_DN_TEMPLATE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.AUTH_LDAP_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_USER_FLAGS_BY_GROUP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTH_LDAP_USER_SEARCH, data["AUTH_LDAP_USER_SEARCH"], false)
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
