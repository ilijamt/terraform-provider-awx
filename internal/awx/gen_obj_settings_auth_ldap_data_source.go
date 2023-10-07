package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &settingsAuthLdapDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthLdapDataSource{}
)

// NewSettingsAuthLDAPDataSource is a helper function to instantiate the SettingsAuthLDAP data source.
func NewSettingsAuthLDAPDataSource() datasource.DataSource {
	return &settingsAuthLdapDataSource{}
}

// settingsAuthLdapDataSource is the data source implementation.
type settingsAuthLdapDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthLdapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ldap/"
}

// Metadata returns the data source type name.
func (o *settingsAuthLdapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_ldap"
}

// Schema defines the schema for the data source.
func (o *settingsAuthLdapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"auth_ldap_1_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Computed:    true,
			},
			"auth_ldap_1_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_1_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Computed:    true,
			},
			"auth_ldap_2_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_2_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Computed:    true,
			},
			"auth_ldap_3_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_3_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Computed:    true,
			},
			"auth_ldap_4_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_4_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Computed:    true,
			},
			"auth_ldap_5_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_5_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Computed:    true,
			},
			"auth_ldap_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_ldap_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *settingsAuthLdapDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthLdapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthLdapTerraformModel
	var err error
	var endpoint = o.endpoint

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
	if err = hookSettingsAuthLdap(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthLDAP",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
