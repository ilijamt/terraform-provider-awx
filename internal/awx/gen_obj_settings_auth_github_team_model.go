package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGithubTeamTerraformModel maps the schema for SettingsAuthGithubTeam when using Data Source
type settingsAuthGithubTeamTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application."
	SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL types.String `tfsdk:"social_auth_github_team_callback_url" json:"SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_TEAM_ID "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_TEAM_ID types.String `tfsdk:"social_auth_github_team_id" json:"SOCIAL_AUTH_GITHUB_TEAM_ID"`
	// SOCIAL_AUTH_GITHUB_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_KEY types.String `tfsdk:"social_auth_github_team_key" json:"SOCIAL_AUTH_GITHUB_TEAM_KEY"`
	// SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_team_organization_map" json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_SECRET types.String `tfsdk:"social_auth_github_team_secret" json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET"`
	// SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP types.String `tfsdk:"social_auth_github_team_team_map" json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"`
}

// Clone the object
func (o *settingsAuthGithubTeamTerraformModel) Clone() settingsAuthGithubTeamTerraformModel {
	return settingsAuthGithubTeamTerraformModel{
		SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_TEAM_ID:               o.SOCIAL_AUTH_GITHUB_TEAM_ID,
		SOCIAL_AUTH_GITHUB_TEAM_KEY:              o.SOCIAL_AUTH_GITHUB_TEAM_KEY,
		SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_TEAM_SECRET:           o.SOCIAL_AUTH_GITHUB_TEAM_SECRET,
		SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubTeam
func (o *settingsAuthGithubTeamTerraformModel) BodyRequest() (req settingsAuthGithubTeamBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_TEAM_ID = o.SOCIAL_AUTH_GITHUB_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP.ValueString())
	return
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamCallbackUrl(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_ID, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamKey(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_KEY, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamOrganizationMap(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamSecret(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_SECRET, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamTeamMap(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubTeamCallbackUrl(data["SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamId(data["SOCIAL_AUTH_GITHUB_TEAM_ID"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamKey(data["SOCIAL_AUTH_GITHUB_TEAM_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamOrganizationMap(data["SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamSecret(data["SOCIAL_AUTH_GITHUB_TEAM_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamTeamMap(data["SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubTeamBodyRequestModel maps the schema for SettingsAuthGithubTeam for creating and updating the data
type settingsAuthGithubTeamBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_TEAM_ID "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_TEAM_ID string `json:"SOCIAL_AUTH_GITHUB_TEAM_ID,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_KEY string `json:"SOCIAL_AUTH_GITHUB_TEAM_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_SECRET string `json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP,omitempty"`
}
