package awx_21_8_0

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGithubTerraformModel maps the schema for SettingsAuthGithub when using Data Source
type settingsAuthGithubTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_CALLBACK_URL types.String `tfsdk:"social_auth_github_callback_url" json:"SOCIAL_AUTH_GITHUB_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_KEY "The OAuth2 key (Client ID) from your GitHub developer application."
	SOCIAL_AUTH_GITHUB_KEY types.String `tfsdk:"social_auth_github_key" json:"SOCIAL_AUTH_GITHUB_KEY"`
	// SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_organization_map" json:"SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_SECRET "The OAuth2 secret (Client Secret) from your GitHub developer application."
	SOCIAL_AUTH_GITHUB_SECRET types.String `tfsdk:"social_auth_github_secret" json:"SOCIAL_AUTH_GITHUB_SECRET"`
	// SOCIAL_AUTH_GITHUB_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_TEAM_MAP types.String `tfsdk:"social_auth_github_team_map" json:"SOCIAL_AUTH_GITHUB_TEAM_MAP"`
}

// Clone the object
func (o *settingsAuthGithubTerraformModel) Clone() settingsAuthGithubTerraformModel {
	return settingsAuthGithubTerraformModel{
		SOCIAL_AUTH_GITHUB_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_KEY:              o.SOCIAL_AUTH_GITHUB_KEY,
		SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_SECRET:           o.SOCIAL_AUTH_GITHUB_SECRET,
		SOCIAL_AUTH_GITHUB_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_TEAM_MAP,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithub
func (o *settingsAuthGithubTerraformModel) BodyRequest() (req settingsAuthGithubBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_KEY = o.SOCIAL_AUTH_GITHUB_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_SECRET = o.SOCIAL_AUTH_GITHUB_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_MAP.ValueString())
	return
}

func (o *settingsAuthGithubTerraformModel) setSocialAuthGithubCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubTerraformModel) setSocialAuthGithubKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_KEY, data, false)
}

func (o *settingsAuthGithubTerraformModel) setSocialAuthGithubOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubTerraformModel) setSocialAuthGithubSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_SECRET, data, false)
}

func (o *settingsAuthGithubTerraformModel) setSocialAuthGithubTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubCallbackUrl(data["SOCIAL_AUTH_GITHUB_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubKey(data["SOCIAL_AUTH_GITHUB_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrganizationMap(data["SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubSecret(data["SOCIAL_AUTH_GITHUB_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamMap(data["SOCIAL_AUTH_GITHUB_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubBodyRequestModel maps the schema for SettingsAuthGithub for creating and updating the data
type settingsAuthGithubBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_KEY "The OAuth2 key (Client ID) from your GitHub developer application."
	SOCIAL_AUTH_GITHUB_KEY string `json:"SOCIAL_AUTH_GITHUB_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_SECRET "The OAuth2 secret (Client Secret) from your GitHub developer application."
	SOCIAL_AUTH_GITHUB_SECRET string `json:"SOCIAL_AUTH_GITHUB_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_MAP,omitempty"`
}
