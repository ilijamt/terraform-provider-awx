package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsMiscAuthenticationTerraformModel maps the schema for SettingsMiscAuthentication when using Data Source
type settingsMiscAuthenticationTerraformModel struct {
	// ALLOW_OAUTH2_FOR_EXTERNAL_USERS "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off."
	ALLOW_OAUTH2_FOR_EXTERNAL_USERS types.Bool `tfsdk:"allow_oauth2_for_external_users" json:"ALLOW_OAUTH2_FOR_EXTERNAL_USERS"`
	// AUTHENTICATION_BACKENDS "List of authentication backends that are enabled based on license features and other authentication settings."
	AUTHENTICATION_BACKENDS types.List `tfsdk:"authentication_backends" json:"AUTHENTICATION_BACKENDS"`
	// AUTH_BASIC_ENABLED "Enable HTTP Basic Auth for the API Browser."
	AUTH_BASIC_ENABLED types.Bool `tfsdk:"auth_basic_enabled" json:"AUTH_BASIC_ENABLED"`
	// DISABLE_LOCAL_AUTH "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration."
	DISABLE_LOCAL_AUTH types.Bool `tfsdk:"disable_local_auth" json:"DISABLE_LOCAL_AUTH"`
	// LOGIN_REDIRECT_OVERRIDE "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page."
	LOGIN_REDIRECT_OVERRIDE types.String `tfsdk:"login_redirect_override" json:"LOGIN_REDIRECT_OVERRIDE"`
	// OAUTH2_PROVIDER "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds."
	OAUTH2_PROVIDER types.String `tfsdk:"oauth2_provider" json:"OAUTH2_PROVIDER"`
	// SESSIONS_PER_USER "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1."
	SESSIONS_PER_USER types.Int64 `tfsdk:"sessions_per_user" json:"SESSIONS_PER_USER"`
	// SESSION_COOKIE_AGE "Number of seconds that a user is inactive before they will need to login again."
	SESSION_COOKIE_AGE types.Int64 `tfsdk:"session_cookie_age" json:"SESSION_COOKIE_AGE"`
	// SOCIAL_AUTH_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_ORGANIZATION_MAP types.String `tfsdk:"social_auth_organization_map" json:"SOCIAL_AUTH_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_TEAM_MAP types.String `tfsdk:"social_auth_team_map" json:"SOCIAL_AUTH_TEAM_MAP"`
	// SOCIAL_AUTH_USER_FIELDS "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login."
	SOCIAL_AUTH_USER_FIELDS types.List `tfsdk:"social_auth_user_fields" json:"SOCIAL_AUTH_USER_FIELDS"`
}

// Clone the object
func (o *settingsMiscAuthenticationTerraformModel) Clone() settingsMiscAuthenticationTerraformModel {
	return settingsMiscAuthenticationTerraformModel{
		ALLOW_OAUTH2_FOR_EXTERNAL_USERS: o.ALLOW_OAUTH2_FOR_EXTERNAL_USERS,
		AUTHENTICATION_BACKENDS:         o.AUTHENTICATION_BACKENDS,
		AUTH_BASIC_ENABLED:              o.AUTH_BASIC_ENABLED,
		DISABLE_LOCAL_AUTH:              o.DISABLE_LOCAL_AUTH,
		LOGIN_REDIRECT_OVERRIDE:         o.LOGIN_REDIRECT_OVERRIDE,
		OAUTH2_PROVIDER:                 o.OAUTH2_PROVIDER,
		SESSIONS_PER_USER:               o.SESSIONS_PER_USER,
		SESSION_COOKIE_AGE:              o.SESSION_COOKIE_AGE,
		SOCIAL_AUTH_ORGANIZATION_MAP:    o.SOCIAL_AUTH_ORGANIZATION_MAP,
		SOCIAL_AUTH_TEAM_MAP:            o.SOCIAL_AUTH_TEAM_MAP,
		SOCIAL_AUTH_USER_FIELDS:         o.SOCIAL_AUTH_USER_FIELDS,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsMiscAuthentication
func (o *settingsMiscAuthenticationTerraformModel) BodyRequest() (req settingsMiscAuthenticationBodyRequestModel) {
	req.ALLOW_OAUTH2_FOR_EXTERNAL_USERS = o.ALLOW_OAUTH2_FOR_EXTERNAL_USERS.ValueBool()
	req.AUTH_BASIC_ENABLED = o.AUTH_BASIC_ENABLED.ValueBool()
	req.DISABLE_LOCAL_AUTH = o.DISABLE_LOCAL_AUTH.ValueBool()
	req.LOGIN_REDIRECT_OVERRIDE = o.LOGIN_REDIRECT_OVERRIDE.ValueString()
	req.OAUTH2_PROVIDER = json.RawMessage(o.OAUTH2_PROVIDER.ValueString())
	req.SESSIONS_PER_USER = o.SESSIONS_PER_USER.ValueInt64()
	req.SESSION_COOKIE_AGE = o.SESSION_COOKIE_AGE.ValueInt64()
	req.SOCIAL_AUTH_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_USER_FIELDS = []string{}
	for _, val := range o.SOCIAL_AUTH_USER_FIELDS.Elements() {
		if _, ok := val.(types.String); ok {
			req.SOCIAL_AUTH_USER_FIELDS = append(req.SOCIAL_AUTH_USER_FIELDS, val.(types.String).ValueString())
		} else {
			req.SOCIAL_AUTH_USER_FIELDS = append(req.SOCIAL_AUTH_USER_FIELDS, val.String())
		}
	}
	return
}

func (o *settingsMiscAuthenticationTerraformModel) setAllowOauth2ForExternalUsers(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.ALLOW_OAUTH2_FOR_EXTERNAL_USERS, data)
}

func (o *settingsMiscAuthenticationTerraformModel) setAuthenticationBackends(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.AUTHENTICATION_BACKENDS, data, false)
}

func (o *settingsMiscAuthenticationTerraformModel) setAuthBasicEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AUTH_BASIC_ENABLED, data)
}

func (o *settingsMiscAuthenticationTerraformModel) setDisableLocalAuth(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.DISABLE_LOCAL_AUTH, data)
}

func (o *settingsMiscAuthenticationTerraformModel) setLoginRedirectOverride(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LOGIN_REDIRECT_OVERRIDE, data, false)
}

func (o *settingsMiscAuthenticationTerraformModel) setOauth2Provider(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.OAUTH2_PROVIDER, data, false)
}

func (o *settingsMiscAuthenticationTerraformModel) setSessionsPerUser(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SESSIONS_PER_USER, data)
}

func (o *settingsMiscAuthenticationTerraformModel) setSessionCookieAge(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SESSION_COOKIE_AGE, data)
}

func (o *settingsMiscAuthenticationTerraformModel) setSocialAuthOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_ORGANIZATION_MAP, data, false)
}

func (o *settingsMiscAuthenticationTerraformModel) setSocialAuthTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_TEAM_MAP, data, false)
}

func (o *settingsMiscAuthenticationTerraformModel) setSocialAuthUserFields(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.SOCIAL_AUTH_USER_FIELDS, data, false)
}

func (o *settingsMiscAuthenticationTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAllowOauth2ForExternalUsers(data["ALLOW_OAUTH2_FOR_EXTERNAL_USERS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthenticationBackends(data["AUTHENTICATION_BACKENDS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAuthBasicEnabled(data["AUTH_BASIC_ENABLED"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDisableLocalAuth(data["DISABLE_LOCAL_AUTH"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLoginRedirectOverride(data["LOGIN_REDIRECT_OVERRIDE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOauth2Provider(data["OAUTH2_PROVIDER"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSessionsPerUser(data["SESSIONS_PER_USER"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSessionCookieAge(data["SESSION_COOKIE_AGE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOrganizationMap(data["SOCIAL_AUTH_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthTeamMap(data["SOCIAL_AUTH_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthUserFields(data["SOCIAL_AUTH_USER_FIELDS"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsMiscAuthenticationBodyRequestModel maps the schema for SettingsMiscAuthentication for creating and updating the data
type settingsMiscAuthenticationBodyRequestModel struct {
	// ALLOW_OAUTH2_FOR_EXTERNAL_USERS "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off."
	ALLOW_OAUTH2_FOR_EXTERNAL_USERS bool `json:"ALLOW_OAUTH2_FOR_EXTERNAL_USERS"`
	// AUTH_BASIC_ENABLED "Enable HTTP Basic Auth for the API Browser."
	AUTH_BASIC_ENABLED bool `json:"AUTH_BASIC_ENABLED"`
	// DISABLE_LOCAL_AUTH "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration."
	DISABLE_LOCAL_AUTH bool `json:"DISABLE_LOCAL_AUTH"`
	// LOGIN_REDIRECT_OVERRIDE "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page."
	LOGIN_REDIRECT_OVERRIDE string `json:"LOGIN_REDIRECT_OVERRIDE,omitempty"`
	// OAUTH2_PROVIDER "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds."
	OAUTH2_PROVIDER json.RawMessage `json:"OAUTH2_PROVIDER,omitempty"`
	// SESSIONS_PER_USER "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1."
	SESSIONS_PER_USER int64 `json:"SESSIONS_PER_USER"`
	// SESSION_COOKIE_AGE "Number of seconds that a user is inactive before they will need to login again."
	SESSION_COOKIE_AGE int64 `json:"SESSION_COOKIE_AGE"`
	// SOCIAL_AUTH_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_USER_FIELDS "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login."
	SOCIAL_AUTH_USER_FIELDS []string `json:"SOCIAL_AUTH_USER_FIELDS,omitempty"`
}
