package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthAzureAdoauth2TerraformModel maps the schema for SettingsAuthAzureADOauth2 when using Data Source
type settingsAuthAzureAdoauth2TerraformModel struct {
	// SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail. "
	SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL types.String `tfsdk:"social_auth_azuread_oauth2_callback_url" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_KEY "The OAuth2 key (Client ID) from your Azure AD application."
	SOCIAL_AUTH_AZUREAD_OAUTH2_KEY types.String `tfsdk:"social_auth_azuread_oauth2_key" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP types.String `tfsdk:"social_auth_azuread_oauth2_organization_map" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET "The OAuth2 secret (Client Secret) from your Azure AD application."
	SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET types.String `tfsdk:"social_auth_azuread_oauth2_secret" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP types.String `tfsdk:"social_auth_azuread_oauth2_team_map" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"`
}

// Clone the object
func (o *settingsAuthAzureAdoauth2TerraformModel) Clone() settingsAuthAzureAdoauth2TerraformModel {
	return settingsAuthAzureAdoauth2TerraformModel{
		SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL:     o.SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL,
		SOCIAL_AUTH_AZUREAD_OAUTH2_KEY:              o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY,
		SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP: o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP,
		SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET:           o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET,
		SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP:         o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthAzureADOauth2
func (o *settingsAuthAzureAdoauth2TerraformModel) BodyRequest() (req settingsAuthAzureAdoauth2BodyRequestModel) {
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY = o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET = o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP.ValueString())
	return
}

func (o *settingsAuthAzureAdoauth2TerraformModel) setSocialAuthAzureadOauth2CallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL, data, false)
}

func (o *settingsAuthAzureAdoauth2TerraformModel) setSocialAuthAzureadOauth2Key(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY, data, false)
}

func (o *settingsAuthAzureAdoauth2TerraformModel) setSocialAuthAzureadOauth2OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthAzureAdoauth2TerraformModel) setSocialAuthAzureadOauth2Secret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET, data, false)
}

func (o *settingsAuthAzureAdoauth2TerraformModel) setSocialAuthAzureadOauth2TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP, data, false)
}

func (o *settingsAuthAzureAdoauth2TerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthAzureadOauth2CallbackUrl(data["SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthAzureadOauth2Key(data["SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthAzureadOauth2OrganizationMap(data["SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthAzureadOauth2Secret(data["SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthAzureadOauth2TeamMap(data["SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthAzureAdoauth2BodyRequestModel maps the schema for SettingsAuthAzureADOauth2 for creating and updating the data
type settingsAuthAzureAdoauth2BodyRequestModel struct {
	// SOCIAL_AUTH_AZUREAD_OAUTH2_KEY "The OAuth2 key (Client ID) from your Azure AD application."
	SOCIAL_AUTH_AZUREAD_OAUTH2_KEY string `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY,omitempty"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET "The OAuth2 secret (Client Secret) from your Azure AD application."
	SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET string `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET,omitempty"`
	// SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP,omitempty"`
}
