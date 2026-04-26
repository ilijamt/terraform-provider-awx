package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthLdapTerraformModel struct {
	AUTH_LDAP_1_BIND_DN             types.String `tfsdk:"auth_ldap_1_bind_dn" json:"AUTH_LDAP_1_BIND_DN"`
	AUTH_LDAP_1_BIND_PASSWORD       types.String `tfsdk:"auth_ldap_1_bind_password" json:"AUTH_LDAP_1_BIND_PASSWORD"`
	AUTH_LDAP_1_CONNECTION_OPTIONS  types.String `tfsdk:"auth_ldap_1_connection_options" json:"AUTH_LDAP_1_CONNECTION_OPTIONS"`
	AUTH_LDAP_1_DENY_GROUP          types.String `tfsdk:"auth_ldap_1_deny_group" json:"AUTH_LDAP_1_DENY_GROUP"`
	AUTH_LDAP_1_GROUP_SEARCH        types.List   `tfsdk:"auth_ldap_1_group_search" json:"AUTH_LDAP_1_GROUP_SEARCH"`
	AUTH_LDAP_1_GROUP_TYPE          types.String `tfsdk:"auth_ldap_1_group_type" json:"AUTH_LDAP_1_GROUP_TYPE"`
	AUTH_LDAP_1_GROUP_TYPE_PARAMS   types.String `tfsdk:"auth_ldap_1_group_type_params" json:"AUTH_LDAP_1_GROUP_TYPE_PARAMS"`
	AUTH_LDAP_1_ORGANIZATION_MAP    types.String `tfsdk:"auth_ldap_1_organization_map" json:"AUTH_LDAP_1_ORGANIZATION_MAP"`
	AUTH_LDAP_1_REQUIRE_GROUP       types.String `tfsdk:"auth_ldap_1_require_group" json:"AUTH_LDAP_1_REQUIRE_GROUP"`
	AUTH_LDAP_1_SERVER_URI          types.String `tfsdk:"auth_ldap_1_server_uri" json:"AUTH_LDAP_1_SERVER_URI"`
	AUTH_LDAP_1_START_TLS           types.Bool   `tfsdk:"auth_ldap_1_start_tls" json:"AUTH_LDAP_1_START_TLS"`
	AUTH_LDAP_1_TEAM_MAP            types.String `tfsdk:"auth_ldap_1_team_map" json:"AUTH_LDAP_1_TEAM_MAP"`
	AUTH_LDAP_1_USER_ATTR_MAP       types.String `tfsdk:"auth_ldap_1_user_attr_map" json:"AUTH_LDAP_1_USER_ATTR_MAP"`
	AUTH_LDAP_1_USER_DN_TEMPLATE    types.String `tfsdk:"auth_ldap_1_user_dn_template" json:"AUTH_LDAP_1_USER_DN_TEMPLATE"`
	AUTH_LDAP_1_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_1_user_flags_by_group" json:"AUTH_LDAP_1_USER_FLAGS_BY_GROUP"`
	AUTH_LDAP_1_USER_SEARCH         types.List   `tfsdk:"auth_ldap_1_user_search" json:"AUTH_LDAP_1_USER_SEARCH"`
	AUTH_LDAP_2_BIND_DN             types.String `tfsdk:"auth_ldap_2_bind_dn" json:"AUTH_LDAP_2_BIND_DN"`
	AUTH_LDAP_2_BIND_PASSWORD       types.String `tfsdk:"auth_ldap_2_bind_password" json:"AUTH_LDAP_2_BIND_PASSWORD"`
	AUTH_LDAP_2_CONNECTION_OPTIONS  types.String `tfsdk:"auth_ldap_2_connection_options" json:"AUTH_LDAP_2_CONNECTION_OPTIONS"`
	AUTH_LDAP_2_DENY_GROUP          types.String `tfsdk:"auth_ldap_2_deny_group" json:"AUTH_LDAP_2_DENY_GROUP"`
	AUTH_LDAP_2_GROUP_SEARCH        types.List   `tfsdk:"auth_ldap_2_group_search" json:"AUTH_LDAP_2_GROUP_SEARCH"`
	AUTH_LDAP_2_GROUP_TYPE          types.String `tfsdk:"auth_ldap_2_group_type" json:"AUTH_LDAP_2_GROUP_TYPE"`
	AUTH_LDAP_2_GROUP_TYPE_PARAMS   types.String `tfsdk:"auth_ldap_2_group_type_params" json:"AUTH_LDAP_2_GROUP_TYPE_PARAMS"`
	AUTH_LDAP_2_ORGANIZATION_MAP    types.String `tfsdk:"auth_ldap_2_organization_map" json:"AUTH_LDAP_2_ORGANIZATION_MAP"`
	AUTH_LDAP_2_REQUIRE_GROUP       types.String `tfsdk:"auth_ldap_2_require_group" json:"AUTH_LDAP_2_REQUIRE_GROUP"`
	AUTH_LDAP_2_SERVER_URI          types.String `tfsdk:"auth_ldap_2_server_uri" json:"AUTH_LDAP_2_SERVER_URI"`
	AUTH_LDAP_2_START_TLS           types.Bool   `tfsdk:"auth_ldap_2_start_tls" json:"AUTH_LDAP_2_START_TLS"`
	AUTH_LDAP_2_TEAM_MAP            types.String `tfsdk:"auth_ldap_2_team_map" json:"AUTH_LDAP_2_TEAM_MAP"`
	AUTH_LDAP_2_USER_ATTR_MAP       types.String `tfsdk:"auth_ldap_2_user_attr_map" json:"AUTH_LDAP_2_USER_ATTR_MAP"`
	AUTH_LDAP_2_USER_DN_TEMPLATE    types.String `tfsdk:"auth_ldap_2_user_dn_template" json:"AUTH_LDAP_2_USER_DN_TEMPLATE"`
	AUTH_LDAP_2_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_2_user_flags_by_group" json:"AUTH_LDAP_2_USER_FLAGS_BY_GROUP"`
	AUTH_LDAP_2_USER_SEARCH         types.List   `tfsdk:"auth_ldap_2_user_search" json:"AUTH_LDAP_2_USER_SEARCH"`
	AUTH_LDAP_3_BIND_DN             types.String `tfsdk:"auth_ldap_3_bind_dn" json:"AUTH_LDAP_3_BIND_DN"`
	AUTH_LDAP_3_BIND_PASSWORD       types.String `tfsdk:"auth_ldap_3_bind_password" json:"AUTH_LDAP_3_BIND_PASSWORD"`
	AUTH_LDAP_3_CONNECTION_OPTIONS  types.String `tfsdk:"auth_ldap_3_connection_options" json:"AUTH_LDAP_3_CONNECTION_OPTIONS"`
	AUTH_LDAP_3_DENY_GROUP          types.String `tfsdk:"auth_ldap_3_deny_group" json:"AUTH_LDAP_3_DENY_GROUP"`
	AUTH_LDAP_3_GROUP_SEARCH        types.List   `tfsdk:"auth_ldap_3_group_search" json:"AUTH_LDAP_3_GROUP_SEARCH"`
	AUTH_LDAP_3_GROUP_TYPE          types.String `tfsdk:"auth_ldap_3_group_type" json:"AUTH_LDAP_3_GROUP_TYPE"`
	AUTH_LDAP_3_GROUP_TYPE_PARAMS   types.String `tfsdk:"auth_ldap_3_group_type_params" json:"AUTH_LDAP_3_GROUP_TYPE_PARAMS"`
	AUTH_LDAP_3_ORGANIZATION_MAP    types.String `tfsdk:"auth_ldap_3_organization_map" json:"AUTH_LDAP_3_ORGANIZATION_MAP"`
	AUTH_LDAP_3_REQUIRE_GROUP       types.String `tfsdk:"auth_ldap_3_require_group" json:"AUTH_LDAP_3_REQUIRE_GROUP"`
	AUTH_LDAP_3_SERVER_URI          types.String `tfsdk:"auth_ldap_3_server_uri" json:"AUTH_LDAP_3_SERVER_URI"`
	AUTH_LDAP_3_START_TLS           types.Bool   `tfsdk:"auth_ldap_3_start_tls" json:"AUTH_LDAP_3_START_TLS"`
	AUTH_LDAP_3_TEAM_MAP            types.String `tfsdk:"auth_ldap_3_team_map" json:"AUTH_LDAP_3_TEAM_MAP"`
	AUTH_LDAP_3_USER_ATTR_MAP       types.String `tfsdk:"auth_ldap_3_user_attr_map" json:"AUTH_LDAP_3_USER_ATTR_MAP"`
	AUTH_LDAP_3_USER_DN_TEMPLATE    types.String `tfsdk:"auth_ldap_3_user_dn_template" json:"AUTH_LDAP_3_USER_DN_TEMPLATE"`
	AUTH_LDAP_3_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_3_user_flags_by_group" json:"AUTH_LDAP_3_USER_FLAGS_BY_GROUP"`
	AUTH_LDAP_3_USER_SEARCH         types.List   `tfsdk:"auth_ldap_3_user_search" json:"AUTH_LDAP_3_USER_SEARCH"`
	AUTH_LDAP_4_BIND_DN             types.String `tfsdk:"auth_ldap_4_bind_dn" json:"AUTH_LDAP_4_BIND_DN"`
	AUTH_LDAP_4_BIND_PASSWORD       types.String `tfsdk:"auth_ldap_4_bind_password" json:"AUTH_LDAP_4_BIND_PASSWORD"`
	AUTH_LDAP_4_CONNECTION_OPTIONS  types.String `tfsdk:"auth_ldap_4_connection_options" json:"AUTH_LDAP_4_CONNECTION_OPTIONS"`
	AUTH_LDAP_4_DENY_GROUP          types.String `tfsdk:"auth_ldap_4_deny_group" json:"AUTH_LDAP_4_DENY_GROUP"`
	AUTH_LDAP_4_GROUP_SEARCH        types.List   `tfsdk:"auth_ldap_4_group_search" json:"AUTH_LDAP_4_GROUP_SEARCH"`
	AUTH_LDAP_4_GROUP_TYPE          types.String `tfsdk:"auth_ldap_4_group_type" json:"AUTH_LDAP_4_GROUP_TYPE"`
	AUTH_LDAP_4_GROUP_TYPE_PARAMS   types.String `tfsdk:"auth_ldap_4_group_type_params" json:"AUTH_LDAP_4_GROUP_TYPE_PARAMS"`
	AUTH_LDAP_4_ORGANIZATION_MAP    types.String `tfsdk:"auth_ldap_4_organization_map" json:"AUTH_LDAP_4_ORGANIZATION_MAP"`
	AUTH_LDAP_4_REQUIRE_GROUP       types.String `tfsdk:"auth_ldap_4_require_group" json:"AUTH_LDAP_4_REQUIRE_GROUP"`
	AUTH_LDAP_4_SERVER_URI          types.String `tfsdk:"auth_ldap_4_server_uri" json:"AUTH_LDAP_4_SERVER_URI"`
	AUTH_LDAP_4_START_TLS           types.Bool   `tfsdk:"auth_ldap_4_start_tls" json:"AUTH_LDAP_4_START_TLS"`
	AUTH_LDAP_4_TEAM_MAP            types.String `tfsdk:"auth_ldap_4_team_map" json:"AUTH_LDAP_4_TEAM_MAP"`
	AUTH_LDAP_4_USER_ATTR_MAP       types.String `tfsdk:"auth_ldap_4_user_attr_map" json:"AUTH_LDAP_4_USER_ATTR_MAP"`
	AUTH_LDAP_4_USER_DN_TEMPLATE    types.String `tfsdk:"auth_ldap_4_user_dn_template" json:"AUTH_LDAP_4_USER_DN_TEMPLATE"`
	AUTH_LDAP_4_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_4_user_flags_by_group" json:"AUTH_LDAP_4_USER_FLAGS_BY_GROUP"`
	AUTH_LDAP_4_USER_SEARCH         types.List   `tfsdk:"auth_ldap_4_user_search" json:"AUTH_LDAP_4_USER_SEARCH"`
	AUTH_LDAP_5_BIND_DN             types.String `tfsdk:"auth_ldap_5_bind_dn" json:"AUTH_LDAP_5_BIND_DN"`
	AUTH_LDAP_5_BIND_PASSWORD       types.String `tfsdk:"auth_ldap_5_bind_password" json:"AUTH_LDAP_5_BIND_PASSWORD"`
	AUTH_LDAP_5_CONNECTION_OPTIONS  types.String `tfsdk:"auth_ldap_5_connection_options" json:"AUTH_LDAP_5_CONNECTION_OPTIONS"`
	AUTH_LDAP_5_DENY_GROUP          types.String `tfsdk:"auth_ldap_5_deny_group" json:"AUTH_LDAP_5_DENY_GROUP"`
	AUTH_LDAP_5_GROUP_SEARCH        types.List   `tfsdk:"auth_ldap_5_group_search" json:"AUTH_LDAP_5_GROUP_SEARCH"`
	AUTH_LDAP_5_GROUP_TYPE          types.String `tfsdk:"auth_ldap_5_group_type" json:"AUTH_LDAP_5_GROUP_TYPE"`
	AUTH_LDAP_5_GROUP_TYPE_PARAMS   types.String `tfsdk:"auth_ldap_5_group_type_params" json:"AUTH_LDAP_5_GROUP_TYPE_PARAMS"`
	AUTH_LDAP_5_ORGANIZATION_MAP    types.String `tfsdk:"auth_ldap_5_organization_map" json:"AUTH_LDAP_5_ORGANIZATION_MAP"`
	AUTH_LDAP_5_REQUIRE_GROUP       types.String `tfsdk:"auth_ldap_5_require_group" json:"AUTH_LDAP_5_REQUIRE_GROUP"`
	AUTH_LDAP_5_SERVER_URI          types.String `tfsdk:"auth_ldap_5_server_uri" json:"AUTH_LDAP_5_SERVER_URI"`
	AUTH_LDAP_5_START_TLS           types.Bool   `tfsdk:"auth_ldap_5_start_tls" json:"AUTH_LDAP_5_START_TLS"`
	AUTH_LDAP_5_TEAM_MAP            types.String `tfsdk:"auth_ldap_5_team_map" json:"AUTH_LDAP_5_TEAM_MAP"`
	AUTH_LDAP_5_USER_ATTR_MAP       types.String `tfsdk:"auth_ldap_5_user_attr_map" json:"AUTH_LDAP_5_USER_ATTR_MAP"`
	AUTH_LDAP_5_USER_DN_TEMPLATE    types.String `tfsdk:"auth_ldap_5_user_dn_template" json:"AUTH_LDAP_5_USER_DN_TEMPLATE"`
	AUTH_LDAP_5_USER_FLAGS_BY_GROUP types.String `tfsdk:"auth_ldap_5_user_flags_by_group" json:"AUTH_LDAP_5_USER_FLAGS_BY_GROUP"`
	AUTH_LDAP_5_USER_SEARCH         types.List   `tfsdk:"auth_ldap_5_user_search" json:"AUTH_LDAP_5_USER_SEARCH"`
	AUTH_LDAP_BIND_DN               types.String `tfsdk:"auth_ldap_bind_dn" json:"AUTH_LDAP_BIND_DN"`
	AUTH_LDAP_BIND_PASSWORD         types.String `tfsdk:"auth_ldap_bind_password" json:"AUTH_LDAP_BIND_PASSWORD"`
	AUTH_LDAP_CONNECTION_OPTIONS    types.String `tfsdk:"auth_ldap_connection_options" json:"AUTH_LDAP_CONNECTION_OPTIONS"`
	AUTH_LDAP_DENY_GROUP            types.String `tfsdk:"auth_ldap_deny_group" json:"AUTH_LDAP_DENY_GROUP"`
	AUTH_LDAP_GROUP_SEARCH          types.List   `tfsdk:"auth_ldap_group_search" json:"AUTH_LDAP_GROUP_SEARCH"`
	AUTH_LDAP_GROUP_TYPE            types.String `tfsdk:"auth_ldap_group_type" json:"AUTH_LDAP_GROUP_TYPE"`
	AUTH_LDAP_GROUP_TYPE_PARAMS     types.String `tfsdk:"auth_ldap_group_type_params" json:"AUTH_LDAP_GROUP_TYPE_PARAMS"`
	AUTH_LDAP_ORGANIZATION_MAP      types.String `tfsdk:"auth_ldap_organization_map" json:"AUTH_LDAP_ORGANIZATION_MAP"`
	AUTH_LDAP_REQUIRE_GROUP         types.String `tfsdk:"auth_ldap_require_group" json:"AUTH_LDAP_REQUIRE_GROUP"`
	AUTH_LDAP_SERVER_URI            types.String `tfsdk:"auth_ldap_server_uri" json:"AUTH_LDAP_SERVER_URI"`
	AUTH_LDAP_START_TLS             types.Bool   `tfsdk:"auth_ldap_start_tls" json:"AUTH_LDAP_START_TLS"`
	AUTH_LDAP_TEAM_MAP              types.String `tfsdk:"auth_ldap_team_map" json:"AUTH_LDAP_TEAM_MAP"`
	AUTH_LDAP_USER_ATTR_MAP         types.String `tfsdk:"auth_ldap_user_attr_map" json:"AUTH_LDAP_USER_ATTR_MAP"`
	AUTH_LDAP_USER_DN_TEMPLATE      types.String `tfsdk:"auth_ldap_user_dn_template" json:"AUTH_LDAP_USER_DN_TEMPLATE"`
	AUTH_LDAP_USER_FLAGS_BY_GROUP   types.String `tfsdk:"auth_ldap_user_flags_by_group" json:"AUTH_LDAP_USER_FLAGS_BY_GROUP"`
	AUTH_LDAP_USER_SEARCH           types.List   `tfsdk:"auth_ldap_user_search" json:"AUTH_LDAP_USER_SEARCH"`
}

func (o *settingsAuthLdapTerraformModel) Clone() settingsAuthLdapTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_BIND_DN, data["AUTH_LDAP_1_BIND_DN"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_BIND_PASSWORD, data["AUTH_LDAP_1_BIND_PASSWORD"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_CONNECTION_OPTIONS, data["AUTH_LDAP_1_CONNECTION_OPTIONS"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_DENY_GROUP, data["AUTH_LDAP_1_DENY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_1_GROUP_SEARCH, data["AUTH_LDAP_1_GROUP_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_GROUP_TYPE, data["AUTH_LDAP_1_GROUP_TYPE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_GROUP_TYPE_PARAMS, data["AUTH_LDAP_1_GROUP_TYPE_PARAMS"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_ORGANIZATION_MAP, data["AUTH_LDAP_1_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_REQUIRE_GROUP, data["AUTH_LDAP_1_REQUIRE_GROUP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_SERVER_URI, data["AUTH_LDAP_1_SERVER_URI"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_LDAP_1_START_TLS, data["AUTH_LDAP_1_START_TLS"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_TEAM_MAP, data["AUTH_LDAP_1_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_USER_ATTR_MAP, data["AUTH_LDAP_1_USER_ATTR_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_1_USER_DN_TEMPLATE, data["AUTH_LDAP_1_USER_DN_TEMPLATE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_1_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_1_USER_FLAGS_BY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_1_USER_SEARCH, data["AUTH_LDAP_1_USER_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_BIND_DN, data["AUTH_LDAP_2_BIND_DN"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_BIND_PASSWORD, data["AUTH_LDAP_2_BIND_PASSWORD"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_CONNECTION_OPTIONS, data["AUTH_LDAP_2_CONNECTION_OPTIONS"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_DENY_GROUP, data["AUTH_LDAP_2_DENY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_2_GROUP_SEARCH, data["AUTH_LDAP_2_GROUP_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_GROUP_TYPE, data["AUTH_LDAP_2_GROUP_TYPE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_GROUP_TYPE_PARAMS, data["AUTH_LDAP_2_GROUP_TYPE_PARAMS"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_ORGANIZATION_MAP, data["AUTH_LDAP_2_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_REQUIRE_GROUP, data["AUTH_LDAP_2_REQUIRE_GROUP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_SERVER_URI, data["AUTH_LDAP_2_SERVER_URI"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_LDAP_2_START_TLS, data["AUTH_LDAP_2_START_TLS"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_TEAM_MAP, data["AUTH_LDAP_2_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_USER_ATTR_MAP, data["AUTH_LDAP_2_USER_ATTR_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_2_USER_DN_TEMPLATE, data["AUTH_LDAP_2_USER_DN_TEMPLATE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_2_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_2_USER_FLAGS_BY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_2_USER_SEARCH, data["AUTH_LDAP_2_USER_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_BIND_DN, data["AUTH_LDAP_3_BIND_DN"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_BIND_PASSWORD, data["AUTH_LDAP_3_BIND_PASSWORD"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_CONNECTION_OPTIONS, data["AUTH_LDAP_3_CONNECTION_OPTIONS"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_DENY_GROUP, data["AUTH_LDAP_3_DENY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_3_GROUP_SEARCH, data["AUTH_LDAP_3_GROUP_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_GROUP_TYPE, data["AUTH_LDAP_3_GROUP_TYPE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_GROUP_TYPE_PARAMS, data["AUTH_LDAP_3_GROUP_TYPE_PARAMS"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_ORGANIZATION_MAP, data["AUTH_LDAP_3_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_REQUIRE_GROUP, data["AUTH_LDAP_3_REQUIRE_GROUP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_SERVER_URI, data["AUTH_LDAP_3_SERVER_URI"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_LDAP_3_START_TLS, data["AUTH_LDAP_3_START_TLS"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_TEAM_MAP, data["AUTH_LDAP_3_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_USER_ATTR_MAP, data["AUTH_LDAP_3_USER_ATTR_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_3_USER_DN_TEMPLATE, data["AUTH_LDAP_3_USER_DN_TEMPLATE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_3_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_3_USER_FLAGS_BY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_3_USER_SEARCH, data["AUTH_LDAP_3_USER_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_BIND_DN, data["AUTH_LDAP_4_BIND_DN"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_BIND_PASSWORD, data["AUTH_LDAP_4_BIND_PASSWORD"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_CONNECTION_OPTIONS, data["AUTH_LDAP_4_CONNECTION_OPTIONS"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_DENY_GROUP, data["AUTH_LDAP_4_DENY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_4_GROUP_SEARCH, data["AUTH_LDAP_4_GROUP_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_GROUP_TYPE, data["AUTH_LDAP_4_GROUP_TYPE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_GROUP_TYPE_PARAMS, data["AUTH_LDAP_4_GROUP_TYPE_PARAMS"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_ORGANIZATION_MAP, data["AUTH_LDAP_4_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_REQUIRE_GROUP, data["AUTH_LDAP_4_REQUIRE_GROUP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_SERVER_URI, data["AUTH_LDAP_4_SERVER_URI"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_LDAP_4_START_TLS, data["AUTH_LDAP_4_START_TLS"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_TEAM_MAP, data["AUTH_LDAP_4_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_USER_ATTR_MAP, data["AUTH_LDAP_4_USER_ATTR_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_4_USER_DN_TEMPLATE, data["AUTH_LDAP_4_USER_DN_TEMPLATE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_4_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_4_USER_FLAGS_BY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_4_USER_SEARCH, data["AUTH_LDAP_4_USER_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_BIND_DN, data["AUTH_LDAP_5_BIND_DN"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_BIND_PASSWORD, data["AUTH_LDAP_5_BIND_PASSWORD"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_CONNECTION_OPTIONS, data["AUTH_LDAP_5_CONNECTION_OPTIONS"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_DENY_GROUP, data["AUTH_LDAP_5_DENY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_5_GROUP_SEARCH, data["AUTH_LDAP_5_GROUP_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_GROUP_TYPE, data["AUTH_LDAP_5_GROUP_TYPE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_GROUP_TYPE_PARAMS, data["AUTH_LDAP_5_GROUP_TYPE_PARAMS"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_ORGANIZATION_MAP, data["AUTH_LDAP_5_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_REQUIRE_GROUP, data["AUTH_LDAP_5_REQUIRE_GROUP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_SERVER_URI, data["AUTH_LDAP_5_SERVER_URI"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_LDAP_5_START_TLS, data["AUTH_LDAP_5_START_TLS"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_TEAM_MAP, data["AUTH_LDAP_5_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_USER_ATTR_MAP, data["AUTH_LDAP_5_USER_ATTR_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_5_USER_DN_TEMPLATE, data["AUTH_LDAP_5_USER_DN_TEMPLATE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_5_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_5_USER_FLAGS_BY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_5_USER_SEARCH, data["AUTH_LDAP_5_USER_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_BIND_DN, data["AUTH_LDAP_BIND_DN"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_BIND_PASSWORD, data["AUTH_LDAP_BIND_PASSWORD"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_CONNECTION_OPTIONS, data["AUTH_LDAP_CONNECTION_OPTIONS"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_DENY_GROUP, data["AUTH_LDAP_DENY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_GROUP_SEARCH, data["AUTH_LDAP_GROUP_SEARCH"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_GROUP_TYPE, data["AUTH_LDAP_GROUP_TYPE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_GROUP_TYPE_PARAMS, data["AUTH_LDAP_GROUP_TYPE_PARAMS"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_ORGANIZATION_MAP, data["AUTH_LDAP_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_REQUIRE_GROUP, data["AUTH_LDAP_REQUIRE_GROUP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_SERVER_URI, data["AUTH_LDAP_SERVER_URI"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_LDAP_START_TLS, data["AUTH_LDAP_START_TLS"]))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_TEAM_MAP, data["AUTH_LDAP_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_USER_ATTR_MAP, data["AUTH_LDAP_USER_ATTR_MAP"], false))
	collect(helpers.AttrValueSetString(&o.AUTH_LDAP_USER_DN_TEMPLATE, data["AUTH_LDAP_USER_DN_TEMPLATE"], false))
	collect(helpers.AttrValueSetJsonString(&o.AUTH_LDAP_USER_FLAGS_BY_GROUP, data["AUTH_LDAP_USER_FLAGS_BY_GROUP"], false))
	collect(helpers.AttrValueSetListString(&o.AUTH_LDAP_USER_SEARCH, data["AUTH_LDAP_USER_SEARCH"], false))
	return diags, nil
}

type settingsAuthLdapBodyRequestModel struct {
	AUTH_LDAP_1_BIND_DN             string          `json:"AUTH_LDAP_1_BIND_DN,omitempty"`
	AUTH_LDAP_1_BIND_PASSWORD       string          `json:"AUTH_LDAP_1_BIND_PASSWORD,omitempty"`
	AUTH_LDAP_1_CONNECTION_OPTIONS  json.RawMessage `json:"AUTH_LDAP_1_CONNECTION_OPTIONS,omitempty"`
	AUTH_LDAP_1_DENY_GROUP          string          `json:"AUTH_LDAP_1_DENY_GROUP,omitempty"`
	AUTH_LDAP_1_GROUP_SEARCH        []string        `json:"AUTH_LDAP_1_GROUP_SEARCH,omitempty"`
	AUTH_LDAP_1_GROUP_TYPE          string          `json:"AUTH_LDAP_1_GROUP_TYPE,omitempty"`
	AUTH_LDAP_1_GROUP_TYPE_PARAMS   json.RawMessage `json:"AUTH_LDAP_1_GROUP_TYPE_PARAMS,omitempty"`
	AUTH_LDAP_1_ORGANIZATION_MAP    json.RawMessage `json:"AUTH_LDAP_1_ORGANIZATION_MAP,omitempty"`
	AUTH_LDAP_1_REQUIRE_GROUP       string          `json:"AUTH_LDAP_1_REQUIRE_GROUP,omitempty"`
	AUTH_LDAP_1_SERVER_URI          string          `json:"AUTH_LDAP_1_SERVER_URI,omitempty"`
	AUTH_LDAP_1_START_TLS           bool            `json:"AUTH_LDAP_1_START_TLS"`
	AUTH_LDAP_1_TEAM_MAP            json.RawMessage `json:"AUTH_LDAP_1_TEAM_MAP,omitempty"`
	AUTH_LDAP_1_USER_ATTR_MAP       json.RawMessage `json:"AUTH_LDAP_1_USER_ATTR_MAP,omitempty"`
	AUTH_LDAP_1_USER_DN_TEMPLATE    string          `json:"AUTH_LDAP_1_USER_DN_TEMPLATE,omitempty"`
	AUTH_LDAP_1_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_1_USER_FLAGS_BY_GROUP,omitempty"`
	AUTH_LDAP_1_USER_SEARCH         []string        `json:"AUTH_LDAP_1_USER_SEARCH,omitempty"`
	AUTH_LDAP_2_BIND_DN             string          `json:"AUTH_LDAP_2_BIND_DN,omitempty"`
	AUTH_LDAP_2_BIND_PASSWORD       string          `json:"AUTH_LDAP_2_BIND_PASSWORD,omitempty"`
	AUTH_LDAP_2_CONNECTION_OPTIONS  json.RawMessage `json:"AUTH_LDAP_2_CONNECTION_OPTIONS,omitempty"`
	AUTH_LDAP_2_DENY_GROUP          string          `json:"AUTH_LDAP_2_DENY_GROUP,omitempty"`
	AUTH_LDAP_2_GROUP_SEARCH        []string        `json:"AUTH_LDAP_2_GROUP_SEARCH,omitempty"`
	AUTH_LDAP_2_GROUP_TYPE          string          `json:"AUTH_LDAP_2_GROUP_TYPE,omitempty"`
	AUTH_LDAP_2_GROUP_TYPE_PARAMS   json.RawMessage `json:"AUTH_LDAP_2_GROUP_TYPE_PARAMS,omitempty"`
	AUTH_LDAP_2_ORGANIZATION_MAP    json.RawMessage `json:"AUTH_LDAP_2_ORGANIZATION_MAP,omitempty"`
	AUTH_LDAP_2_REQUIRE_GROUP       string          `json:"AUTH_LDAP_2_REQUIRE_GROUP,omitempty"`
	AUTH_LDAP_2_SERVER_URI          string          `json:"AUTH_LDAP_2_SERVER_URI,omitempty"`
	AUTH_LDAP_2_START_TLS           bool            `json:"AUTH_LDAP_2_START_TLS"`
	AUTH_LDAP_2_TEAM_MAP            json.RawMessage `json:"AUTH_LDAP_2_TEAM_MAP,omitempty"`
	AUTH_LDAP_2_USER_ATTR_MAP       json.RawMessage `json:"AUTH_LDAP_2_USER_ATTR_MAP,omitempty"`
	AUTH_LDAP_2_USER_DN_TEMPLATE    string          `json:"AUTH_LDAP_2_USER_DN_TEMPLATE,omitempty"`
	AUTH_LDAP_2_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_2_USER_FLAGS_BY_GROUP,omitempty"`
	AUTH_LDAP_2_USER_SEARCH         []string        `json:"AUTH_LDAP_2_USER_SEARCH,omitempty"`
	AUTH_LDAP_3_BIND_DN             string          `json:"AUTH_LDAP_3_BIND_DN,omitempty"`
	AUTH_LDAP_3_BIND_PASSWORD       string          `json:"AUTH_LDAP_3_BIND_PASSWORD,omitempty"`
	AUTH_LDAP_3_CONNECTION_OPTIONS  json.RawMessage `json:"AUTH_LDAP_3_CONNECTION_OPTIONS,omitempty"`
	AUTH_LDAP_3_DENY_GROUP          string          `json:"AUTH_LDAP_3_DENY_GROUP,omitempty"`
	AUTH_LDAP_3_GROUP_SEARCH        []string        `json:"AUTH_LDAP_3_GROUP_SEARCH,omitempty"`
	AUTH_LDAP_3_GROUP_TYPE          string          `json:"AUTH_LDAP_3_GROUP_TYPE,omitempty"`
	AUTH_LDAP_3_GROUP_TYPE_PARAMS   json.RawMessage `json:"AUTH_LDAP_3_GROUP_TYPE_PARAMS,omitempty"`
	AUTH_LDAP_3_ORGANIZATION_MAP    json.RawMessage `json:"AUTH_LDAP_3_ORGANIZATION_MAP,omitempty"`
	AUTH_LDAP_3_REQUIRE_GROUP       string          `json:"AUTH_LDAP_3_REQUIRE_GROUP,omitempty"`
	AUTH_LDAP_3_SERVER_URI          string          `json:"AUTH_LDAP_3_SERVER_URI,omitempty"`
	AUTH_LDAP_3_START_TLS           bool            `json:"AUTH_LDAP_3_START_TLS"`
	AUTH_LDAP_3_TEAM_MAP            json.RawMessage `json:"AUTH_LDAP_3_TEAM_MAP,omitempty"`
	AUTH_LDAP_3_USER_ATTR_MAP       json.RawMessage `json:"AUTH_LDAP_3_USER_ATTR_MAP,omitempty"`
	AUTH_LDAP_3_USER_DN_TEMPLATE    string          `json:"AUTH_LDAP_3_USER_DN_TEMPLATE,omitempty"`
	AUTH_LDAP_3_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_3_USER_FLAGS_BY_GROUP,omitempty"`
	AUTH_LDAP_3_USER_SEARCH         []string        `json:"AUTH_LDAP_3_USER_SEARCH,omitempty"`
	AUTH_LDAP_4_BIND_DN             string          `json:"AUTH_LDAP_4_BIND_DN,omitempty"`
	AUTH_LDAP_4_BIND_PASSWORD       string          `json:"AUTH_LDAP_4_BIND_PASSWORD,omitempty"`
	AUTH_LDAP_4_CONNECTION_OPTIONS  json.RawMessage `json:"AUTH_LDAP_4_CONNECTION_OPTIONS,omitempty"`
	AUTH_LDAP_4_DENY_GROUP          string          `json:"AUTH_LDAP_4_DENY_GROUP,omitempty"`
	AUTH_LDAP_4_GROUP_SEARCH        []string        `json:"AUTH_LDAP_4_GROUP_SEARCH,omitempty"`
	AUTH_LDAP_4_GROUP_TYPE          string          `json:"AUTH_LDAP_4_GROUP_TYPE,omitempty"`
	AUTH_LDAP_4_GROUP_TYPE_PARAMS   json.RawMessage `json:"AUTH_LDAP_4_GROUP_TYPE_PARAMS,omitempty"`
	AUTH_LDAP_4_ORGANIZATION_MAP    json.RawMessage `json:"AUTH_LDAP_4_ORGANIZATION_MAP,omitempty"`
	AUTH_LDAP_4_REQUIRE_GROUP       string          `json:"AUTH_LDAP_4_REQUIRE_GROUP,omitempty"`
	AUTH_LDAP_4_SERVER_URI          string          `json:"AUTH_LDAP_4_SERVER_URI,omitempty"`
	AUTH_LDAP_4_START_TLS           bool            `json:"AUTH_LDAP_4_START_TLS"`
	AUTH_LDAP_4_TEAM_MAP            json.RawMessage `json:"AUTH_LDAP_4_TEAM_MAP,omitempty"`
	AUTH_LDAP_4_USER_ATTR_MAP       json.RawMessage `json:"AUTH_LDAP_4_USER_ATTR_MAP,omitempty"`
	AUTH_LDAP_4_USER_DN_TEMPLATE    string          `json:"AUTH_LDAP_4_USER_DN_TEMPLATE,omitempty"`
	AUTH_LDAP_4_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_4_USER_FLAGS_BY_GROUP,omitempty"`
	AUTH_LDAP_4_USER_SEARCH         []string        `json:"AUTH_LDAP_4_USER_SEARCH,omitempty"`
	AUTH_LDAP_5_BIND_DN             string          `json:"AUTH_LDAP_5_BIND_DN,omitempty"`
	AUTH_LDAP_5_BIND_PASSWORD       string          `json:"AUTH_LDAP_5_BIND_PASSWORD,omitempty"`
	AUTH_LDAP_5_CONNECTION_OPTIONS  json.RawMessage `json:"AUTH_LDAP_5_CONNECTION_OPTIONS,omitempty"`
	AUTH_LDAP_5_DENY_GROUP          string          `json:"AUTH_LDAP_5_DENY_GROUP,omitempty"`
	AUTH_LDAP_5_GROUP_SEARCH        []string        `json:"AUTH_LDAP_5_GROUP_SEARCH,omitempty"`
	AUTH_LDAP_5_GROUP_TYPE          string          `json:"AUTH_LDAP_5_GROUP_TYPE,omitempty"`
	AUTH_LDAP_5_GROUP_TYPE_PARAMS   json.RawMessage `json:"AUTH_LDAP_5_GROUP_TYPE_PARAMS,omitempty"`
	AUTH_LDAP_5_ORGANIZATION_MAP    json.RawMessage `json:"AUTH_LDAP_5_ORGANIZATION_MAP,omitempty"`
	AUTH_LDAP_5_REQUIRE_GROUP       string          `json:"AUTH_LDAP_5_REQUIRE_GROUP,omitempty"`
	AUTH_LDAP_5_SERVER_URI          string          `json:"AUTH_LDAP_5_SERVER_URI,omitempty"`
	AUTH_LDAP_5_START_TLS           bool            `json:"AUTH_LDAP_5_START_TLS"`
	AUTH_LDAP_5_TEAM_MAP            json.RawMessage `json:"AUTH_LDAP_5_TEAM_MAP,omitempty"`
	AUTH_LDAP_5_USER_ATTR_MAP       json.RawMessage `json:"AUTH_LDAP_5_USER_ATTR_MAP,omitempty"`
	AUTH_LDAP_5_USER_DN_TEMPLATE    string          `json:"AUTH_LDAP_5_USER_DN_TEMPLATE,omitempty"`
	AUTH_LDAP_5_USER_FLAGS_BY_GROUP json.RawMessage `json:"AUTH_LDAP_5_USER_FLAGS_BY_GROUP,omitempty"`
	AUTH_LDAP_5_USER_SEARCH         []string        `json:"AUTH_LDAP_5_USER_SEARCH,omitempty"`
	AUTH_LDAP_BIND_DN               string          `json:"AUTH_LDAP_BIND_DN,omitempty"`
	AUTH_LDAP_BIND_PASSWORD         string          `json:"AUTH_LDAP_BIND_PASSWORD,omitempty"`
	AUTH_LDAP_CONNECTION_OPTIONS    json.RawMessage `json:"AUTH_LDAP_CONNECTION_OPTIONS,omitempty"`
	AUTH_LDAP_DENY_GROUP            string          `json:"AUTH_LDAP_DENY_GROUP,omitempty"`
	AUTH_LDAP_GROUP_SEARCH          []string        `json:"AUTH_LDAP_GROUP_SEARCH,omitempty"`
	AUTH_LDAP_GROUP_TYPE            string          `json:"AUTH_LDAP_GROUP_TYPE,omitempty"`
	AUTH_LDAP_GROUP_TYPE_PARAMS     json.RawMessage `json:"AUTH_LDAP_GROUP_TYPE_PARAMS,omitempty"`
	AUTH_LDAP_ORGANIZATION_MAP      json.RawMessage `json:"AUTH_LDAP_ORGANIZATION_MAP,omitempty"`
	AUTH_LDAP_REQUIRE_GROUP         string          `json:"AUTH_LDAP_REQUIRE_GROUP,omitempty"`
	AUTH_LDAP_SERVER_URI            string          `json:"AUTH_LDAP_SERVER_URI,omitempty"`
	AUTH_LDAP_START_TLS             bool            `json:"AUTH_LDAP_START_TLS"`
	AUTH_LDAP_TEAM_MAP              json.RawMessage `json:"AUTH_LDAP_TEAM_MAP,omitempty"`
	AUTH_LDAP_USER_ATTR_MAP         json.RawMessage `json:"AUTH_LDAP_USER_ATTR_MAP,omitempty"`
	AUTH_LDAP_USER_DN_TEMPLATE      string          `json:"AUTH_LDAP_USER_DN_TEMPLATE,omitempty"`
	AUTH_LDAP_USER_FLAGS_BY_GROUP   json.RawMessage `json:"AUTH_LDAP_USER_FLAGS_BY_GROUP,omitempty"`
	AUTH_LDAP_USER_SEARCH           []string        `json:"AUTH_LDAP_USER_SEARCH,omitempty"`
}
