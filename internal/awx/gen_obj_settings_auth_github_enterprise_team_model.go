package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthGithubEnterpriseTeamTerraformModel maps the schema for SettingsAuthGithubEnterpriseTeam when using Data Source
type settingsAuthGithubEnterpriseTeamTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL types.String `tfsdk:"social_auth_github_enterprise_team_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL types.String `tfsdk:"social_auth_github_enterprise_team_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID types.String `tfsdk:"social_auth_github_enterprise_team_id" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY types.String `tfsdk:"social_auth_github_enterprise_team_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_team_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET types.String `tfsdk:"social_auth_github_enterprise_team_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP types.String `tfsdk:"social_auth_github_enterprise_team_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL types.String `tfsdk:"social_auth_github_enterprise_team_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"`
}

// Clone the object
func (o *settingsAuthGithubEnterpriseTeamTerraformModel) Clone() settingsAuthGithubEnterpriseTeamTerraformModel {
	return settingsAuthGithubEnterpriseTeamTerraformModel{
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL:          o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID:               o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET:           o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterpriseTeam
func (o *settingsAuthGithubEnterpriseTeamTerraformModel) BodyRequest() (req settingsAuthGithubEnterpriseTeamBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL.ValueString()
	return
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamApiUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamApiUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamCallbackUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamId(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamKey(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamOrganizationMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamSecret(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamTeamMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubEnterpriseTeamBodyRequestModel maps the schema for SettingsAuthGithubEnterpriseTeam for creating and updating the data
type settingsAuthGithubEnterpriseTeamBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL,omitempty"`
}
