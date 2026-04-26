package awx

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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterprise
func (o *settingsAuthGithubEnterpriseTerraformModel) BodyRequest() *settingsAuthGithubEnterpriseBodyRequestModel {
	var req settingsAuthGithubEnterpriseBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL.ValueString()
	return &req
}

func (o *settingsAuthGithubEnterpriseTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_URL"], false)
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
