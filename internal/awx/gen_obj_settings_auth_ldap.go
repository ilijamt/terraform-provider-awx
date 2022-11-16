package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// settingsAuthLDAPTerraformModel maps the schema for SettingsAuthLDAP when using Data Source
type settingsAuthLDAPTerraformModel struct {
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
func (o settingsAuthLDAPTerraformModel) Clone() settingsAuthLDAPTerraformModel {
	return settingsAuthLDAPTerraformModel{
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
func (o settingsAuthLDAPTerraformModel) BodyRequest() (req settingsAuthLDAPBodyRequestModel) {
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

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1BindDn(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_BIND_DN"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_BIND_DN = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_BIND_DN = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_BIND_DN = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1BindPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_BIND_PASSWORD"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_BIND_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_BIND_PASSWORD = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_BIND_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_CONNECTION_OPTIONS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_1_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_1_CONNECTION_OPTIONS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1DenyGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_DENY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_DENY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_DENY_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_DENY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1GroupSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_GROUP_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_1_GROUP_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_1_GROUP_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_1_GROUP_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1GroupType(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_GROUP_TYPE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_GROUP_TYPE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_GROUP_TYPE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_GROUP_TYPE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_GROUP_TYPE_PARAMS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_1_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_1_GROUP_TYPE_PARAMS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1OrganizationMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_ORGANIZATION_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_1_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_1_ORGANIZATION_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1RequireGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_REQUIRE_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_REQUIRE_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_REQUIRE_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_REQUIRE_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1ServerUri(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_SERVER_URI"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_SERVER_URI = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_SERVER_URI = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_SERVER_URI = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1StartTls(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_START_TLS"
	if val, ok := data.(bool); ok {
		o.AUTH_LDAP_1_START_TLS = types.BoolValue(val)
	} else {
		o.AUTH_LDAP_1_START_TLS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1TeamMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_TEAM_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_1_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_1_TEAM_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1UserAttrMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_USER_ATTR_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_1_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_1_USER_ATTR_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_USER_DN_TEMPLATE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_USER_DN_TEMPLATE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_1_USER_DN_TEMPLATE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_1_USER_DN_TEMPLATE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_USER_FLAGS_BY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap1UserSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_1_USER_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_1_USER_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_1_USER_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_1_USER_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2BindDn(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_BIND_DN"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_BIND_DN = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_BIND_DN = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_BIND_DN = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2BindPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_BIND_PASSWORD"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_BIND_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_BIND_PASSWORD = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_BIND_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_CONNECTION_OPTIONS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_2_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_2_CONNECTION_OPTIONS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2DenyGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_DENY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_DENY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_DENY_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_DENY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2GroupSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_GROUP_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_2_GROUP_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_2_GROUP_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_2_GROUP_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2GroupType(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_GROUP_TYPE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_GROUP_TYPE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_GROUP_TYPE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_GROUP_TYPE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_GROUP_TYPE_PARAMS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_2_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_2_GROUP_TYPE_PARAMS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2OrganizationMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_ORGANIZATION_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_2_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_2_ORGANIZATION_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2RequireGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_REQUIRE_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_REQUIRE_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_REQUIRE_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_REQUIRE_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2ServerUri(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_SERVER_URI"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_SERVER_URI = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_SERVER_URI = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_SERVER_URI = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2StartTls(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_START_TLS"
	if val, ok := data.(bool); ok {
		o.AUTH_LDAP_2_START_TLS = types.BoolValue(val)
	} else {
		o.AUTH_LDAP_2_START_TLS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2TeamMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_TEAM_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_2_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_2_TEAM_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2UserAttrMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_USER_ATTR_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_2_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_2_USER_ATTR_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_USER_DN_TEMPLATE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_USER_DN_TEMPLATE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_2_USER_DN_TEMPLATE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_2_USER_DN_TEMPLATE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_USER_FLAGS_BY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap2UserSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_2_USER_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_2_USER_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_2_USER_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_2_USER_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3BindDn(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_BIND_DN"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_BIND_DN = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_BIND_DN = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_BIND_DN = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3BindPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_BIND_PASSWORD"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_BIND_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_BIND_PASSWORD = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_BIND_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_CONNECTION_OPTIONS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_3_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_3_CONNECTION_OPTIONS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3DenyGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_DENY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_DENY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_DENY_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_DENY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3GroupSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_GROUP_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_3_GROUP_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_3_GROUP_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_3_GROUP_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3GroupType(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_GROUP_TYPE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_GROUP_TYPE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_GROUP_TYPE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_GROUP_TYPE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_GROUP_TYPE_PARAMS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_3_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_3_GROUP_TYPE_PARAMS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3OrganizationMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_ORGANIZATION_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_3_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_3_ORGANIZATION_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3RequireGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_REQUIRE_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_REQUIRE_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_REQUIRE_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_REQUIRE_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3ServerUri(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_SERVER_URI"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_SERVER_URI = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_SERVER_URI = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_SERVER_URI = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3StartTls(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_START_TLS"
	if val, ok := data.(bool); ok {
		o.AUTH_LDAP_3_START_TLS = types.BoolValue(val)
	} else {
		o.AUTH_LDAP_3_START_TLS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3TeamMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_TEAM_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_3_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_3_TEAM_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3UserAttrMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_USER_ATTR_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_3_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_3_USER_ATTR_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_USER_DN_TEMPLATE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_USER_DN_TEMPLATE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_3_USER_DN_TEMPLATE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_3_USER_DN_TEMPLATE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_USER_FLAGS_BY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap3UserSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_3_USER_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_3_USER_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_3_USER_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_3_USER_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4BindDn(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_BIND_DN"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_BIND_DN = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_BIND_DN = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_BIND_DN = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4BindPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_BIND_PASSWORD"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_BIND_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_BIND_PASSWORD = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_BIND_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_CONNECTION_OPTIONS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_4_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_4_CONNECTION_OPTIONS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4DenyGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_DENY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_DENY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_DENY_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_DENY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4GroupSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_GROUP_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_4_GROUP_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_4_GROUP_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_4_GROUP_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4GroupType(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_GROUP_TYPE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_GROUP_TYPE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_GROUP_TYPE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_GROUP_TYPE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_GROUP_TYPE_PARAMS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_4_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_4_GROUP_TYPE_PARAMS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4OrganizationMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_ORGANIZATION_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_4_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_4_ORGANIZATION_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4RequireGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_REQUIRE_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_REQUIRE_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_REQUIRE_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_REQUIRE_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4ServerUri(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_SERVER_URI"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_SERVER_URI = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_SERVER_URI = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_SERVER_URI = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4StartTls(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_START_TLS"
	if val, ok := data.(bool); ok {
		o.AUTH_LDAP_4_START_TLS = types.BoolValue(val)
	} else {
		o.AUTH_LDAP_4_START_TLS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4TeamMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_TEAM_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_4_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_4_TEAM_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4UserAttrMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_USER_ATTR_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_4_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_4_USER_ATTR_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_USER_DN_TEMPLATE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_USER_DN_TEMPLATE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_4_USER_DN_TEMPLATE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_4_USER_DN_TEMPLATE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_USER_FLAGS_BY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap4UserSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_4_USER_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_4_USER_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_4_USER_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_4_USER_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5BindDn(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_BIND_DN"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_BIND_DN = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_BIND_DN = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_BIND_DN = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5BindPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_BIND_PASSWORD"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_BIND_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_BIND_PASSWORD = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_BIND_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5ConnectionOptions(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_CONNECTION_OPTIONS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_5_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_5_CONNECTION_OPTIONS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5DenyGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_DENY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_DENY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_DENY_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_DENY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5GroupSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_GROUP_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_5_GROUP_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_5_GROUP_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_5_GROUP_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5GroupType(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_GROUP_TYPE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_GROUP_TYPE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_GROUP_TYPE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_GROUP_TYPE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5GroupTypeParams(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_GROUP_TYPE_PARAMS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_5_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_5_GROUP_TYPE_PARAMS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5OrganizationMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_ORGANIZATION_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_5_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_5_ORGANIZATION_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5RequireGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_REQUIRE_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_REQUIRE_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_REQUIRE_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_REQUIRE_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5ServerUri(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_SERVER_URI"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_SERVER_URI = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_SERVER_URI = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_SERVER_URI = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5StartTls(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_START_TLS"
	if val, ok := data.(bool); ok {
		o.AUTH_LDAP_5_START_TLS = types.BoolValue(val)
	} else {
		o.AUTH_LDAP_5_START_TLS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5TeamMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_TEAM_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_5_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_5_TEAM_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5UserAttrMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_USER_ATTR_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_5_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_5_USER_ATTR_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5UserDnTemplate(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_USER_DN_TEMPLATE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_USER_DN_TEMPLATE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_5_USER_DN_TEMPLATE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_5_USER_DN_TEMPLATE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5UserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_USER_FLAGS_BY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdap5UserSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_5_USER_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_5_USER_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_5_USER_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_5_USER_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapBindDn(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_BIND_DN"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_BIND_DN = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_BIND_DN = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_BIND_DN = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapBindPassword(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_BIND_PASSWORD"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_BIND_PASSWORD = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_BIND_PASSWORD = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_BIND_PASSWORD = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapConnectionOptions(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_CONNECTION_OPTIONS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_CONNECTION_OPTIONS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_CONNECTION_OPTIONS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapDenyGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_DENY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_DENY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_DENY_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_DENY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapGroupSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_GROUP_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_GROUP_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_GROUP_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_GROUP_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapGroupType(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_GROUP_TYPE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_GROUP_TYPE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_GROUP_TYPE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_GROUP_TYPE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapGroupTypeParams(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_GROUP_TYPE_PARAMS"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_GROUP_TYPE_PARAMS = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_GROUP_TYPE_PARAMS = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapOrganizationMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_ORGANIZATION_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_ORGANIZATION_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_ORGANIZATION_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapRequireGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_REQUIRE_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_REQUIRE_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_REQUIRE_GROUP = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_REQUIRE_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapServerUri(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_SERVER_URI"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_SERVER_URI = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_SERVER_URI = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_SERVER_URI = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapStartTls(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_START_TLS"
	if val, ok := data.(bool); ok {
		o.AUTH_LDAP_START_TLS = types.BoolValue(val)
	} else {
		o.AUTH_LDAP_START_TLS = types.BoolNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapTeamMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_TEAM_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_TEAM_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_TEAM_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapUserAttrMap(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_USER_ATTR_MAP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_USER_ATTR_MAP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_USER_ATTR_MAP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapUserDnTemplate(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_USER_DN_TEMPLATE"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_USER_DN_TEMPLATE = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.AUTH_LDAP_USER_DN_TEMPLATE = types.StringValue(val.String())
	} else {
		o.AUTH_LDAP_USER_DN_TEMPLATE = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapUserFlagsByGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_USER_FLAGS_BY_GROUP"
	if val, ok := data.(string); ok {
		o.AUTH_LDAP_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.AUTH_LDAP_USER_FLAGS_BY_GROUP = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.AUTH_LDAP_USER_FLAGS_BY_GROUP = types.StringNull()
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) setAuthLdapUserSearch(data any) (d diag.Diagnostics, err error) {
	// Decode "AUTH_LDAP_USER_SEARCH"
	if val, ok := data.(types.List); ok {
		o.AUTH_LDAP_USER_SEARCH = types.ListValueMust(types.StringType, val.Elements())
	} else if val, ok := data.([]any); ok {
		var list []attr.Value
		for _, v := range val {
			list = append(list, types.StringValue(helpers.TrimString(false, false, v.(string))))
		}
		o.AUTH_LDAP_USER_SEARCH = types.ListValueMust(types.StringType, list)
	} else if data == nil {
		o.AUTH_LDAP_USER_SEARCH = types.ListValueMust(types.StringType, []attr.Value{})
	} else {
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
		return d, err
	}
	return d, nil
}

func (o *settingsAuthLDAPTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
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

// settingsAuthLDAPBodyRequestModel maps the schema for SettingsAuthLDAP for creating and updating the data
type settingsAuthLDAPBodyRequestModel struct {
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

var (
	_ datasource.DataSource              = &settingsAuthLDAPDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthLDAPDataSource{}
)

// NewSettingsAuthLDAPDataSource is a helper function to instantiate the SettingsAuthLDAP data source.
func NewSettingsAuthLDAPDataSource() datasource.DataSource {
	return &settingsAuthLDAPDataSource{}
}

// settingsAuthLDAPDataSource is the data source implementation.
type settingsAuthLDAPDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthLDAPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ldap/"
}

// Metadata returns the data source type name.
func (o *settingsAuthLDAPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_ldap"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthLDAPDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthLDAP",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"auth_ldap_1_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_1_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_2_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_3_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_4_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_5_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthLDAPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthLDAPTerraformModel
	var err error
	var endpoint string
	endpoint = o.endpoint

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthLDAP
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthLDAP on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	if err = hookSettingsAuthLdap(SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthLDAP"),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsAuthLDAPResource{}
	_ resource.ResourceWithConfigure = &settingsAuthLDAPResource{}
)

// NewSettingsAuthLDAPResource is a helper function to simplify the provider implementation.
func NewSettingsAuthLDAPResource() resource.Resource {
	return &settingsAuthLDAPResource{}
}

type settingsAuthLDAPResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthLDAPResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ldap/"
}

func (o settingsAuthLDAPResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_ldap"
}

func (o settingsAuthLDAPResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthLDAP",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"auth_ldap_1_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`MemberDNGroupType`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_1_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"member_attr":"member","name_attr":"cn"}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_1_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`MemberDNGroupType`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_2_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"member_attr":"member","name_attr":"cn"}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_2_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`MemberDNGroupType`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_3_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"member_attr":"member","name_attr":"cn"}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_3_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`MemberDNGroupType`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_4_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"member_attr":"member","name_attr":"cn"}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_4_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`MemberDNGroupType`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_5_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"member_attr":"member","name_attr":"cn"}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_5_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_bind_dn": {
					Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_bind_password": {
					Description: "Password used to bind LDAP user account.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_connection_options": {
					Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_deny_group": {
					Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_group_search": {
					Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_group_type": {
					Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`MemberDNGroupType`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
					},
				},
				"auth_ldap_group_type_params": {
					Description: "Key value parameters to send the chosen group type init method.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"member_attr":"member","name_attr":"cn"}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_organization_map": {
					Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_require_group": {
					Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_server_uri": {
					Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_start_tls": {
					Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_team_map": {
					Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_attr_map": {
					Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_dn_template": {
					Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_flags_by_group": {
					Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"auth_ldap_user_search": {
					Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
			},
		}), nil
}

func (o *settingsAuthLDAPResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthLDAPTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthLDAP resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthLDAP on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthLdap(SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthLDAP"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthLDAPResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthLDAPTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthLDAP from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthLDAP on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthLdap(SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthLDAP"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthLDAPResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthLDAPTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthLDAP resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthLDAP on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthLdap(SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthLDAP"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthLDAPResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	return
}
