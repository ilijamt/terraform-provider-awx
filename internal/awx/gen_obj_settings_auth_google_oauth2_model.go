package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGoogleOauth2TerraformModel maps the schema for SettingsAuthGoogleOauth2 when using Data Source
type settingsAuthGoogleOauth2TerraformModel struct {
	// SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail."
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS types.String `tfsdk:"social_auth_google_oauth2_auth_extra_arguments" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL types.String `tfsdk:"social_auth_google_oauth2_callback_url" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_KEY "The OAuth2 key from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY types.String `tfsdk:"social_auth_google_oauth2_key" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP types.String `tfsdk:"social_auth_google_oauth2_organization_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET "The OAuth2 secret from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET types.String `tfsdk:"social_auth_google_oauth2_secret" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP types.String `tfsdk:"social_auth_google_oauth2_team_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS "Update this setting to restrict the domains who are allowed to login using Google OAuth2."
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS types.List `tfsdk:"social_auth_google_oauth2_whitelisted_domains" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"`
}

// Clone the object
func (o *settingsAuthGoogleOauth2TerraformModel) Clone() settingsAuthGoogleOauth2TerraformModel {
	return settingsAuthGoogleOauth2TerraformModel{
		SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS: o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS,
		SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL:         o.SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL,
		SOCIAL_AUTH_GOOGLE_OAUTH2_KEY:                  o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY,
		SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP:     o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP,
		SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET:               o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET,
		SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP:             o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP,
		SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS:  o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGoogleOauth2
func (o *settingsAuthGoogleOauth2TerraformModel) BodyRequest() (req settingsAuthGoogleOauth2BodyRequestModel) {
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY = o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET = o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = []string{}
	for _, val := range o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS.Elements() {
		if _, ok := val.(types.String); ok {
			req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = append(req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, val.(types.String).ValueString())
		} else {
			req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = append(req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, val.String())
		}
	}
	return
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2AuthExtraArguments(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2CallbackUrl(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2Key(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2OrganizationMap(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2Secret(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2TeamMap(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2WhitelistedDomains(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetListString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGoogleOauth2AuthExtraArguments(data["SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2CallbackUrl(data["SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2Key(data["SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2OrganizationMap(data["SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2Secret(data["SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2TeamMap(data["SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2WhitelistedDomains(data["SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGoogleOauth2BodyRequestModel maps the schema for SettingsAuthGoogleOauth2 for creating and updating the data
type settingsAuthGoogleOauth2BodyRequestModel struct {
	// SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail."
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_KEY "The OAuth2 key from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY string `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET "The OAuth2 secret from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET string `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS "Update this setting to restrict the domains who are allowed to login using Google OAuth2."
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS []string `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS,omitempty"`
}
