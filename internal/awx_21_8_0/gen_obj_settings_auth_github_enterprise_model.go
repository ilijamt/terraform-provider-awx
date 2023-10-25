package awx_21_8_0

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGithubEnterpriseTerraformModel maps the schema for SettingsAuthGithubEnterprise when using Data Source
type settingsAuthGithubEnterpriseTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL types.String `tfsdk:"social_auth_github_enterprise_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL types.String `tfsdk:"social_auth_github_enterprise_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY types.String `tfsdk:"social_auth_github_enterprise_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET types.String `tfsdk:"social_auth_github_enterprise_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP types.String `tfsdk:"social_auth_github_enterprise_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_URL types.String `tfsdk:"social_auth_github_enterprise_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_URL"`
}

// Clone the object
func (o *settingsAuthGithubEnterpriseTerraformModel) Clone() settingsAuthGithubEnterpriseTerraformModel {
	return settingsAuthGithubEnterpriseTerraformModel{
		SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL:          o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET:           o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_URL:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterprise
func (o *settingsAuthGithubEnterpriseTerraformModel) BodyRequest() (req settingsAuthGithubEnterpriseBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL.ValueString()
	return
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseApiUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseApiUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseCallbackUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseKey(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrganizationMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseSecret(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubEnterpriseBodyRequestModel maps the schema for SettingsAuthGithubEnterprise for creating and updating the data
type settingsAuthGithubEnterpriseBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_URL,omitempty"`
}
