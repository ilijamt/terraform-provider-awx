package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsMiscAuthenticationTerraformModel struct {
	ALLOW_METRICS_FOR_ANONYMOUS_USERS  types.Bool   `tfsdk:"allow_metrics_for_anonymous_users" json:"ALLOW_METRICS_FOR_ANONYMOUS_USERS"`
	ALLOW_OAUTH2_FOR_EXTERNAL_USERS    types.Bool   `tfsdk:"allow_oauth2_for_external_users" json:"ALLOW_OAUTH2_FOR_EXTERNAL_USERS"`
	AUTHENTICATION_BACKENDS            types.List   `tfsdk:"authentication_backends" json:"AUTHENTICATION_BACKENDS"`
	AUTH_BASIC_ENABLED                 types.Bool   `tfsdk:"auth_basic_enabled" json:"AUTH_BASIC_ENABLED"`
	DISABLE_LOCAL_AUTH                 types.Bool   `tfsdk:"disable_local_auth" json:"DISABLE_LOCAL_AUTH"`
	LOCAL_PASSWORD_MIN_DIGITS          types.Int64  `tfsdk:"local_password_min_digits" json:"LOCAL_PASSWORD_MIN_DIGITS"`
	LOCAL_PASSWORD_MIN_LENGTH          types.Int64  `tfsdk:"local_password_min_length" json:"LOCAL_PASSWORD_MIN_LENGTH"`
	LOCAL_PASSWORD_MIN_SPECIAL         types.Int64  `tfsdk:"local_password_min_special" json:"LOCAL_PASSWORD_MIN_SPECIAL"`
	LOCAL_PASSWORD_MIN_UPPER           types.Int64  `tfsdk:"local_password_min_upper" json:"LOCAL_PASSWORD_MIN_UPPER"`
	LOGIN_REDIRECT_OVERRIDE            types.String `tfsdk:"login_redirect_override" json:"LOGIN_REDIRECT_OVERRIDE"`
	OAUTH2_PROVIDER                    types.String `tfsdk:"oauth2_provider" json:"OAUTH2_PROVIDER"`
	SESSIONS_PER_USER                  types.Int64  `tfsdk:"sessions_per_user" json:"SESSIONS_PER_USER"`
	SESSION_COOKIE_AGE                 types.Int64  `tfsdk:"session_cookie_age" json:"SESSION_COOKIE_AGE"`
	SOCIAL_AUTH_ORGANIZATION_MAP       types.String `tfsdk:"social_auth_organization_map" json:"SOCIAL_AUTH_ORGANIZATION_MAP"`
	SOCIAL_AUTH_TEAM_MAP               types.String `tfsdk:"social_auth_team_map" json:"SOCIAL_AUTH_TEAM_MAP"`
	SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL types.Bool   `tfsdk:"social_auth_username_is_full_email" json:"SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL"`
	SOCIAL_AUTH_USER_FIELDS            types.List   `tfsdk:"social_auth_user_fields" json:"SOCIAL_AUTH_USER_FIELDS"`
}

func (o *settingsMiscAuthenticationTerraformModel) Clone() settingsMiscAuthenticationTerraformModel {
	return *o
}

func (o *settingsMiscAuthenticationTerraformModel) BodyRequest() *settingsMiscAuthenticationBodyRequestModel {
	var req settingsMiscAuthenticationBodyRequestModel
	req.ALLOW_METRICS_FOR_ANONYMOUS_USERS = o.ALLOW_METRICS_FOR_ANONYMOUS_USERS.ValueBool()
	req.ALLOW_OAUTH2_FOR_EXTERNAL_USERS = o.ALLOW_OAUTH2_FOR_EXTERNAL_USERS.ValueBool()
	req.AUTH_BASIC_ENABLED = o.AUTH_BASIC_ENABLED.ValueBool()
	req.DISABLE_LOCAL_AUTH = o.DISABLE_LOCAL_AUTH.ValueBool()
	req.LOCAL_PASSWORD_MIN_DIGITS = o.LOCAL_PASSWORD_MIN_DIGITS.ValueInt64()
	req.LOCAL_PASSWORD_MIN_LENGTH = o.LOCAL_PASSWORD_MIN_LENGTH.ValueInt64()
	req.LOCAL_PASSWORD_MIN_SPECIAL = o.LOCAL_PASSWORD_MIN_SPECIAL.ValueInt64()
	req.LOCAL_PASSWORD_MIN_UPPER = o.LOCAL_PASSWORD_MIN_UPPER.ValueInt64()
	req.LOGIN_REDIRECT_OVERRIDE = o.LOGIN_REDIRECT_OVERRIDE.ValueString()
	req.OAUTH2_PROVIDER = json.RawMessage(o.OAUTH2_PROVIDER.ValueString())
	req.SESSIONS_PER_USER = o.SESSIONS_PER_USER.ValueInt64()
	req.SESSION_COOKIE_AGE = o.SESSION_COOKIE_AGE.ValueInt64()
	req.SOCIAL_AUTH_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL = o.SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL.ValueBool()
	req.SOCIAL_AUTH_USER_FIELDS = helpers.ListAsStringSlice(o.SOCIAL_AUTH_USER_FIELDS, false)
	return &req
}

func (o *settingsMiscAuthenticationTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.ALLOW_METRICS_FOR_ANONYMOUS_USERS, data["ALLOW_METRICS_FOR_ANONYMOUS_USERS"]))
	collect(helpers.AttrValueSetBool(&o.ALLOW_OAUTH2_FOR_EXTERNAL_USERS, data["ALLOW_OAUTH2_FOR_EXTERNAL_USERS"]))
	collect(helpers.AttrValueSetListString(&o.AUTHENTICATION_BACKENDS, data["AUTHENTICATION_BACKENDS"], false))
	collect(helpers.AttrValueSetBool(&o.AUTH_BASIC_ENABLED, data["AUTH_BASIC_ENABLED"]))
	collect(helpers.AttrValueSetBool(&o.DISABLE_LOCAL_AUTH, data["DISABLE_LOCAL_AUTH"]))
	collect(helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_DIGITS, data["LOCAL_PASSWORD_MIN_DIGITS"]))
	collect(helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_LENGTH, data["LOCAL_PASSWORD_MIN_LENGTH"]))
	collect(helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_SPECIAL, data["LOCAL_PASSWORD_MIN_SPECIAL"]))
	collect(helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_UPPER, data["LOCAL_PASSWORD_MIN_UPPER"]))
	collect(helpers.AttrValueSetString(&o.LOGIN_REDIRECT_OVERRIDE, data["LOGIN_REDIRECT_OVERRIDE"], false))
	collect(helpers.AttrValueSetJsonString(&o.OAUTH2_PROVIDER, data["OAUTH2_PROVIDER"], false))
	collect(helpers.AttrValueSetInt64(&o.SESSIONS_PER_USER, data["SESSIONS_PER_USER"]))
	collect(helpers.AttrValueSetInt64(&o.SESSION_COOKIE_AGE, data["SESSION_COOKIE_AGE"]))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_ORGANIZATION_MAP, data["SOCIAL_AUTH_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_TEAM_MAP, data["SOCIAL_AUTH_TEAM_MAP"], false))
	collect(helpers.AttrValueSetBool(&o.SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL, data["SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL"]))
	collect(helpers.AttrValueSetListString(&o.SOCIAL_AUTH_USER_FIELDS, data["SOCIAL_AUTH_USER_FIELDS"], false))
	return diags, nil
}

type settingsMiscAuthenticationBodyRequestModel struct {
	ALLOW_METRICS_FOR_ANONYMOUS_USERS  bool            `json:"ALLOW_METRICS_FOR_ANONYMOUS_USERS"`
	ALLOW_OAUTH2_FOR_EXTERNAL_USERS    bool            `json:"ALLOW_OAUTH2_FOR_EXTERNAL_USERS"`
	AUTH_BASIC_ENABLED                 bool            `json:"AUTH_BASIC_ENABLED"`
	DISABLE_LOCAL_AUTH                 bool            `json:"DISABLE_LOCAL_AUTH"`
	LOCAL_PASSWORD_MIN_DIGITS          int64           `json:"LOCAL_PASSWORD_MIN_DIGITS,omitempty"`
	LOCAL_PASSWORD_MIN_LENGTH          int64           `json:"LOCAL_PASSWORD_MIN_LENGTH,omitempty"`
	LOCAL_PASSWORD_MIN_SPECIAL         int64           `json:"LOCAL_PASSWORD_MIN_SPECIAL,omitempty"`
	LOCAL_PASSWORD_MIN_UPPER           int64           `json:"LOCAL_PASSWORD_MIN_UPPER,omitempty"`
	LOGIN_REDIRECT_OVERRIDE            string          `json:"LOGIN_REDIRECT_OVERRIDE,omitempty"`
	OAUTH2_PROVIDER                    json.RawMessage `json:"OAUTH2_PROVIDER,omitempty"`
	SESSIONS_PER_USER                  int64           `json:"SESSIONS_PER_USER,omitempty"`
	SESSION_COOKIE_AGE                 int64           `json:"SESSION_COOKIE_AGE,omitempty"`
	SOCIAL_AUTH_ORGANIZATION_MAP       json.RawMessage `json:"SOCIAL_AUTH_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_TEAM_MAP               json.RawMessage `json:"SOCIAL_AUTH_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL bool            `json:"SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL"`
	SOCIAL_AUTH_USER_FIELDS            []string        `json:"SOCIAL_AUTH_USER_FIELDS,omitempty"`
}

type settingsMiscAuthenticationResource = framework.GenericResource[settingsMiscAuthenticationTerraformModel, settingsMiscAuthenticationBodyRequestModel, *settingsMiscAuthenticationTerraformModel]

// NewSettingsMiscAuthenticationResource is a helper function to simplify the provider implementation.
func NewSettingsMiscAuthenticationResource() resource.Resource {
	return &settingsMiscAuthenticationResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_authentication", Endpoint: "/api/v2/settings/authentication/"}},
		Cfg: framework.ResourceCfg[settingsMiscAuthenticationTerraformModel, settingsMiscAuthenticationBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"allow_metrics_for_anonymous_users": schema.BoolAttribute{
						Description: "If true, anonymous users are allowed to poll metrics.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"allow_oauth2_for_external_users": schema.BoolAttribute{
						Description: "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"auth_basic_enabled": schema.BoolAttribute{
						Description: "Enable HTTP Basic Auth for the API Browser.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"disable_local_auth": schema.BoolAttribute{
						Description: "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"local_password_min_digits": schema.Int64Attribute{
						Description: "Minimum number of digit characters required in a local password. 0 means no minimum",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"local_password_min_length": schema.Int64Attribute{
						Description: "Minimum number of characters required in a local password. 0 means no minimum",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"local_password_min_special": schema.Int64Attribute{
						Description: "Minimum number of special characters required in a local password. 0 means no minimum",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"local_password_min_upper": schema.Int64Attribute{
						Description: "Minimum number of uppercase characters required in a local password. 0 means no minimum",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"login_redirect_override": schema.StringAttribute{
						Description: "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"oauth2_provider": schema.StringAttribute{
						Description: "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"ACCESS_TOKEN_EXPIRE_SECONDS":31536000000,"AUTHORIZATION_CODE_EXPIRE_SECONDS":600,"REFRESH_TOKEN_EXPIRE_SECONDS":2628000}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"sessions_per_user": schema.Int64Attribute{
						Description: "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(-1),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"session_cookie_age": schema.Int64Attribute{
						Description: "Number of seconds that a user is inactive before they will need to login again.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1800),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(60, 30000000000),
						},
					},
					"social_auth_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_username_is_full_email": schema.BoolAttribute{
						Description: "Enabling this setting will tell social auth to use the full Email as username instead of the full name",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_user_fields": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"authentication_backends": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of authentication backends that are enabled based on license features and other authentication settings.",
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsMiscAuthentication",
		},
	}
}

type settingsMiscAuthenticationDataSource = framework.GenericDataSource[settingsMiscAuthenticationTerraformModel, *settingsMiscAuthenticationTerraformModel]

// NewSettingsMiscAuthenticationDataSource is a helper function to instantiate the SettingsMiscAuthentication data source.
func NewSettingsMiscAuthenticationDataSource() datasource.DataSource {
	return &settingsMiscAuthenticationDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_authentication", Endpoint: "/api/v2/settings/authentication/"}},
		Cfg: framework.DataSourceCfg[settingsMiscAuthenticationTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"allow_metrics_for_anonymous_users": dschema.BoolAttribute{
						Description: "If true, anonymous users are allowed to poll metrics.",
						Computed:    true,
					},
					"allow_oauth2_for_external_users": dschema.BoolAttribute{
						Description: "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off.",
						Computed:    true,
					},
					"authentication_backends": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of authentication backends that are enabled based on license features and other authentication settings.",
						Computed:    true,
					},
					"auth_basic_enabled": dschema.BoolAttribute{
						Description: "Enable HTTP Basic Auth for the API Browser.",
						Computed:    true,
					},
					"disable_local_auth": dschema.BoolAttribute{
						Description: "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration.",
						Computed:    true,
					},
					"local_password_min_digits": dschema.Int64Attribute{
						Description: "Minimum number of digit characters required in a local password. 0 means no minimum",
						Computed:    true,
					},
					"local_password_min_length": dschema.Int64Attribute{
						Description: "Minimum number of characters required in a local password. 0 means no minimum",
						Computed:    true,
					},
					"local_password_min_special": dschema.Int64Attribute{
						Description: "Minimum number of special characters required in a local password. 0 means no minimum",
						Computed:    true,
					},
					"local_password_min_upper": dschema.Int64Attribute{
						Description: "Minimum number of uppercase characters required in a local password. 0 means no minimum",
						Computed:    true,
					},
					"login_redirect_override": dschema.StringAttribute{
						Description: "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page.",
						Computed:    true,
					},
					"oauth2_provider": dschema.StringAttribute{
						Description: "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds.",
						Computed:    true,
					},
					"sessions_per_user": dschema.Int64Attribute{
						Description: "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1.",
						Computed:    true,
					},
					"session_cookie_age": dschema.Int64Attribute{
						Description: "Number of seconds that a user is inactive before they will need to login again.",
						Computed:    true,
					},
					"social_auth_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_username_is_full_email": dschema.BoolAttribute{
						Description: "Enabling this setting will tell social auth to use the full Email as username instead of the full name",
						Computed:    true,
					},
					"social_auth_user_fields": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsMiscAuthentication",
		},
	}
}
