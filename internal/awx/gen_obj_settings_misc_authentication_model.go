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
	// ALLOW_METRICS_FOR_ANONYMOUS_USERS "If true, anonymous users are allowed to poll metrics."
	ALLOW_METRICS_FOR_ANONYMOUS_USERS types.Bool `tfsdk:"allow_metrics_for_anonymous_users" json:"ALLOW_METRICS_FOR_ANONYMOUS_USERS"`
	// ALLOW_OAUTH2_FOR_EXTERNAL_USERS "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off."
	ALLOW_OAUTH2_FOR_EXTERNAL_USERS types.Bool `tfsdk:"allow_oauth2_for_external_users" json:"ALLOW_OAUTH2_FOR_EXTERNAL_USERS"`
	// AUTHENTICATION_BACKENDS "List of authentication backends that are enabled based on license features and other authentication settings."
	AUTHENTICATION_BACKENDS types.List `tfsdk:"authentication_backends" json:"AUTHENTICATION_BACKENDS"`
	// AUTH_BASIC_ENABLED "Enable HTTP Basic Auth for the API Browser."
	AUTH_BASIC_ENABLED types.Bool `tfsdk:"auth_basic_enabled" json:"AUTH_BASIC_ENABLED"`
	// DISABLE_LOCAL_AUTH "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration."
	DISABLE_LOCAL_AUTH types.Bool `tfsdk:"disable_local_auth" json:"DISABLE_LOCAL_AUTH"`
	// LOCAL_PASSWORD_MIN_DIGITS "Minimum number of digit characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_DIGITS types.Int64 `tfsdk:"local_password_min_digits" json:"LOCAL_PASSWORD_MIN_DIGITS"`
	// LOCAL_PASSWORD_MIN_LENGTH "Minimum number of characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_LENGTH types.Int64 `tfsdk:"local_password_min_length" json:"LOCAL_PASSWORD_MIN_LENGTH"`
	// LOCAL_PASSWORD_MIN_SPECIAL "Minimum number of special characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_SPECIAL types.Int64 `tfsdk:"local_password_min_special" json:"LOCAL_PASSWORD_MIN_SPECIAL"`
	// LOCAL_PASSWORD_MIN_UPPER "Minimum number of uppercase characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_UPPER types.Int64 `tfsdk:"local_password_min_upper" json:"LOCAL_PASSWORD_MIN_UPPER"`
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
	// SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL "Enabling this setting will tell social auth to use the full Email as username instead of the full name"
	SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL types.Bool `tfsdk:"social_auth_username_is_full_email" json:"SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL"`
	// SOCIAL_AUTH_USER_FIELDS "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login."
	SOCIAL_AUTH_USER_FIELDS types.List `tfsdk:"social_auth_user_fields" json:"SOCIAL_AUTH_USER_FIELDS"`
}

// Clone the object
func (o *settingsMiscAuthenticationTerraformModel) Clone() settingsMiscAuthenticationTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsMiscAuthentication
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
	{
		dg, _ := helpers.AttrValueSetBool(&o.ALLOW_METRICS_FOR_ANONYMOUS_USERS, data["ALLOW_METRICS_FOR_ANONYMOUS_USERS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ALLOW_OAUTH2_FOR_EXTERNAL_USERS, data["ALLOW_OAUTH2_FOR_EXTERNAL_USERS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.AUTHENTICATION_BACKENDS, data["AUTHENTICATION_BACKENDS"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AUTH_BASIC_ENABLED, data["AUTH_BASIC_ENABLED"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.DISABLE_LOCAL_AUTH, data["DISABLE_LOCAL_AUTH"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_DIGITS, data["LOCAL_PASSWORD_MIN_DIGITS"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_LENGTH, data["LOCAL_PASSWORD_MIN_LENGTH"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_SPECIAL, data["LOCAL_PASSWORD_MIN_SPECIAL"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LOCAL_PASSWORD_MIN_UPPER, data["LOCAL_PASSWORD_MIN_UPPER"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LOGIN_REDIRECT_OVERRIDE, data["LOGIN_REDIRECT_OVERRIDE"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.OAUTH2_PROVIDER, data["OAUTH2_PROVIDER"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.SESSIONS_PER_USER, data["SESSIONS_PER_USER"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.SESSION_COOKIE_AGE, data["SESSION_COOKIE_AGE"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_ORGANIZATION_MAP, data["SOCIAL_AUTH_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_TEAM_MAP, data["SOCIAL_AUTH_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL, data["SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetListString(&o.SOCIAL_AUTH_USER_FIELDS, data["SOCIAL_AUTH_USER_FIELDS"], false)
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsMiscAuthenticationBodyRequestModel maps the schema for SettingsMiscAuthentication for creating and updating the data
type settingsMiscAuthenticationBodyRequestModel struct {
	// ALLOW_METRICS_FOR_ANONYMOUS_USERS "If true, anonymous users are allowed to poll metrics."
	ALLOW_METRICS_FOR_ANONYMOUS_USERS bool `json:"ALLOW_METRICS_FOR_ANONYMOUS_USERS"`
	// ALLOW_OAUTH2_FOR_EXTERNAL_USERS "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off."
	ALLOW_OAUTH2_FOR_EXTERNAL_USERS bool `json:"ALLOW_OAUTH2_FOR_EXTERNAL_USERS"`
	// AUTH_BASIC_ENABLED "Enable HTTP Basic Auth for the API Browser."
	AUTH_BASIC_ENABLED bool `json:"AUTH_BASIC_ENABLED"`
	// DISABLE_LOCAL_AUTH "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration."
	DISABLE_LOCAL_AUTH bool `json:"DISABLE_LOCAL_AUTH"`
	// LOCAL_PASSWORD_MIN_DIGITS "Minimum number of digit characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_DIGITS int64 `json:"LOCAL_PASSWORD_MIN_DIGITS,omitempty"`
	// LOCAL_PASSWORD_MIN_LENGTH "Minimum number of characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_LENGTH int64 `json:"LOCAL_PASSWORD_MIN_LENGTH,omitempty"`
	// LOCAL_PASSWORD_MIN_SPECIAL "Minimum number of special characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_SPECIAL int64 `json:"LOCAL_PASSWORD_MIN_SPECIAL,omitempty"`
	// LOCAL_PASSWORD_MIN_UPPER "Minimum number of uppercase characters required in a local password. 0 means no minimum"
	LOCAL_PASSWORD_MIN_UPPER int64 `json:"LOCAL_PASSWORD_MIN_UPPER,omitempty"`
	// LOGIN_REDIRECT_OVERRIDE "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page."
	LOGIN_REDIRECT_OVERRIDE string `json:"LOGIN_REDIRECT_OVERRIDE,omitempty"`
	// OAUTH2_PROVIDER "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds."
	OAUTH2_PROVIDER json.RawMessage `json:"OAUTH2_PROVIDER,omitempty"`
	// SESSIONS_PER_USER "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1."
	SESSIONS_PER_USER int64 `json:"SESSIONS_PER_USER,omitempty"`
	// SESSION_COOKIE_AGE "Number of seconds that a user is inactive before they will need to login again."
	SESSION_COOKIE_AGE int64 `json:"SESSION_COOKIE_AGE,omitempty"`
	// SOCIAL_AUTH_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL "Enabling this setting will tell social auth to use the full Email as username instead of the full name"
	SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL bool `json:"SOCIAL_AUTH_USERNAME_IS_FULL_EMAIL"`
	// SOCIAL_AUTH_USER_FIELDS "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login."
	SOCIAL_AUTH_USER_FIELDS []string `json:"SOCIAL_AUTH_USER_FIELDS,omitempty"`
}
