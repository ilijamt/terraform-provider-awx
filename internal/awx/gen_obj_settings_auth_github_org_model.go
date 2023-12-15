package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGithubOrgTerraformModel maps the schema for SettingsAuthGithubOrg when using Data Source
type settingsAuthGithubOrgTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL types.String `tfsdk:"social_auth_github_org_callback_url" json:"SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ORG_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_KEY types.String `tfsdk:"social_auth_github_org_key" json:"SOCIAL_AUTH_GITHUB_ORG_KEY"`
	// SOCIAL_AUTH_GITHUB_ORG_NAME "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ORG_NAME types.String `tfsdk:"social_auth_github_org_name" json:"SOCIAL_AUTH_GITHUB_ORG_NAME"`
	// SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_org_organization_map" json:"SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_SECRET types.String `tfsdk:"social_auth_github_org_secret" json:"SOCIAL_AUTH_GITHUB_ORG_SECRET"`
	// SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP types.String `tfsdk:"social_auth_github_org_team_map" json:"SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP"`
}

// Clone the object
func (o *settingsAuthGithubOrgTerraformModel) Clone() settingsAuthGithubOrgTerraformModel {
	return settingsAuthGithubOrgTerraformModel{
		SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ORG_KEY:              o.SOCIAL_AUTH_GITHUB_ORG_KEY,
		SOCIAL_AUTH_GITHUB_ORG_NAME:             o.SOCIAL_AUTH_GITHUB_ORG_NAME,
		SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ORG_SECRET:           o.SOCIAL_AUTH_GITHUB_ORG_SECRET,
		SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubOrg
func (o *settingsAuthGithubOrgTerraformModel) BodyRequest() (req settingsAuthGithubOrgBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ORG_KEY = o.SOCIAL_AUTH_GITHUB_ORG_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_NAME = o.SOCIAL_AUTH_GITHUB_ORG_NAME.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ORG_SECRET = o.SOCIAL_AUTH_GITHUB_ORG_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP.ValueString())
	return
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_KEY, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_NAME, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_SECRET, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubOrgCallbackUrl(data["SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgKey(data["SOCIAL_AUTH_GITHUB_ORG_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgName(data["SOCIAL_AUTH_GITHUB_ORG_NAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgOrganizationMap(data["SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgSecret(data["SOCIAL_AUTH_GITHUB_ORG_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgTeamMap(data["SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubOrgBodyRequestModel maps the schema for SettingsAuthGithubOrg for creating and updating the data
type settingsAuthGithubOrgBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ORG_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_KEY string `json:"SOCIAL_AUTH_GITHUB_ORG_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_NAME "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ORG_NAME string `json:"SOCIAL_AUTH_GITHUB_ORG_NAME,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_SECRET string `json:"SOCIAL_AUTH_GITHUB_ORG_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP,omitempty"`
}
