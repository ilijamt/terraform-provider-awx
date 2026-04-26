package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
