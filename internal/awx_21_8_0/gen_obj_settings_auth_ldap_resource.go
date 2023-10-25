package awx_21_8_0

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsAuthLdapResource{}
	_ resource.ResourceWithConfigure = &settingsAuthLdapResource{}
)

// NewSettingsAuthLDAPResource is a helper function to simplify the provider implementation.
func NewSettingsAuthLDAPResource() resource.Resource {
	return &settingsAuthLdapResource{}
}

type settingsAuthLdapResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthLdapResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ldap/"
}

func (o *settingsAuthLdapResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_ldap"
}

func (o *settingsAuthLdapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Request elements
			"auth_ldap_1_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_1_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`MemberDNGroupType`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
				},
			},
			"auth_ldap_1_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"member_attr":"member","name_attr":"cn"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"auth_ldap_1_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_1_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_2_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_2_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`MemberDNGroupType`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
				},
			},
			"auth_ldap_2_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"member_attr":"member","name_attr":"cn"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"auth_ldap_2_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_2_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_3_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_3_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`MemberDNGroupType`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
				},
			},
			"auth_ldap_3_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"member_attr":"member","name_attr":"cn"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"auth_ldap_3_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_3_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_4_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_4_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`MemberDNGroupType`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
				},
			},
			"auth_ldap_4_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"member_attr":"member","name_attr":"cn"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"auth_ldap_4_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_4_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_5_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_5_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`MemberDNGroupType`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
				},
			},
			"auth_ldap_5_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"member_attr":"member","name_attr":"cn"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"auth_ldap_5_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_5_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_bind_dn": schema.StringAttribute{
				Description: "DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_bind_password": schema.StringAttribute{
				Description: "Password used to bind LDAP user account.",
				Sensitive:   true,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_connection_options": schema.StringAttribute{
				Description: "Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \"OPT_REFERRALS\"). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"OPT_NETWORK_TIMEOUT":30,"OPT_REFERRALS":0}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_deny_group": schema.StringAttribute{
				Description: "Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_group_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			"auth_ldap_group_type": schema.StringAttribute{
				Description: "The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`MemberDNGroupType`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PosixGroupType", "GroupOfNamesType", "GroupOfUniqueNamesType", "ActiveDirectoryGroupType", "OrganizationalRoleGroupType", "MemberDNGroupType", "NestedGroupOfNamesType", "NestedGroupOfUniqueNamesType", "NestedActiveDirectoryGroupType", "NestedOrganizationalRoleGroupType", "NestedMemberDNGroupType", "PosixUIDGroupType"}...),
				},
			},
			"auth_ldap_group_type_params": schema.StringAttribute{
				Description: "Key value parameters to send the chosen group type init method.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{"member_attr":"member","name_attr":"cn"}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_organization_map": schema.StringAttribute{
				Description: "Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_require_group": schema.StringAttribute{
				Description: "Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_server_uri": schema.StringAttribute{
				Description: "URI to connect to LDAP server, such as \"ldap://ldap.example.com:389\" (non-SSL) or \"ldaps://ldap.example.com:636\" (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_start_tls": schema.BoolAttribute{
				Description: "Whether to enable TLS when the LDAP connection is not using SSL.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"auth_ldap_team_map": schema.StringAttribute{
				Description: "Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_user_attr_map": schema.StringAttribute{
				Description: "Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_user_dn_template": schema.StringAttribute{
				Description: "Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_user_flags_by_group": schema.StringAttribute{
				Description: "Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"auth_ldap_user_search": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \"LDAPUnion\" is possible. See the documentation for details.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{},
			},
			// Write only elements
			// Data only elements
		},
	}
}

func (o *settingsAuthLdapResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthLdapTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsAuthLDAP/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for create", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthLDAP resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthLDAP on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthLdap(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthLDAP",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthLdapResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthLdapTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for read", endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthLDAP from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthLDAP on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthLdap(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthLDAP",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthLdapResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthLdapTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthLDAP
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsAuthLDAP/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthLDAP on %s for update", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthLDAP resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthLDAP on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthLdap(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthLDAP",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthLdapResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
