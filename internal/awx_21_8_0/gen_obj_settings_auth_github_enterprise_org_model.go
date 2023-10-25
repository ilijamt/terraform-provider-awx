package awx_21_8_0

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGithubEnterpriseOrgTerraformModel maps the schema for SettingsAuthGithubEnterpriseOrg when using Data Source
type settingsAuthGithubEnterpriseOrgTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL types.String `tfsdk:"social_auth_github_enterprise_org_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL types.String `tfsdk:"social_auth_github_enterprise_org_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY types.String `tfsdk:"social_auth_github_enterprise_org_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME types.String `tfsdk:"social_auth_github_enterprise_org_name" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_org_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET types.String `tfsdk:"social_auth_github_enterprise_org_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP types.String `tfsdk:"social_auth_github_enterprise_org_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL types.String `tfsdk:"social_auth_github_enterprise_org_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL"`
}

// Clone the object
func (o *settingsAuthGithubEnterpriseOrgTerraformModel) Clone() settingsAuthGithubEnterpriseOrgTerraformModel {
	return settingsAuthGithubEnterpriseOrgTerraformModel{
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL:          o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME:             o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET:           o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterpriseOrg
func (o *settingsAuthGithubEnterpriseOrgTerraformModel) BodyRequest() (req settingsAuthGithubEnterpriseOrgBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL.ValueString()
	return
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgApiUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgApiUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgCallbackUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgKey(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgName(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgOrganizationMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgSecret(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgTeamMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubEnterpriseOrgBodyRequestModel maps the schema for SettingsAuthGithubEnterpriseOrg for creating and updating the data
type settingsAuthGithubEnterpriseOrgBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL,omitempty"`
}
