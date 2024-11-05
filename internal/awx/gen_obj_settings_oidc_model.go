package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsOpenIdconnectTerraformModel maps the schema for SettingsOpenIDConnect when using Data Source
type settingsOpenIdconnectTerraformModel struct {
	// SOCIAL_AUTH_OIDC_KEY "The OIDC key (Client ID) from your IDP."
	SOCIAL_AUTH_OIDC_KEY types.String `tfsdk:"social_auth_oidc_key" json:"SOCIAL_AUTH_OIDC_KEY"`
	// SOCIAL_AUTH_OIDC_OIDC_ENDPOINT "The URL for your OIDC provider including the path up to /.well-known/openid-configuration"
	SOCIAL_AUTH_OIDC_OIDC_ENDPOINT types.String `tfsdk:"social_auth_oidc_oidc_endpoint" json:"SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"`
	// SOCIAL_AUTH_OIDC_SECRET "The OIDC secret (Client Secret) from your IDP."
	SOCIAL_AUTH_OIDC_SECRET types.String `tfsdk:"social_auth_oidc_secret" json:"SOCIAL_AUTH_OIDC_SECRET"`
	// SOCIAL_AUTH_OIDC_VERIFY_SSL "Verify the OIDC provider ssl certificate."
	SOCIAL_AUTH_OIDC_VERIFY_SSL types.Bool `tfsdk:"social_auth_oidc_verify_ssl" json:"SOCIAL_AUTH_OIDC_VERIFY_SSL"`
}

// Clone the object
func (o *settingsOpenIdconnectTerraformModel) Clone() settingsOpenIdconnectTerraformModel {
	return settingsOpenIdconnectTerraformModel{
		SOCIAL_AUTH_OIDC_KEY:           o.SOCIAL_AUTH_OIDC_KEY,
		SOCIAL_AUTH_OIDC_OIDC_ENDPOINT: o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT,
		SOCIAL_AUTH_OIDC_SECRET:        o.SOCIAL_AUTH_OIDC_SECRET,
		SOCIAL_AUTH_OIDC_VERIFY_SSL:    o.SOCIAL_AUTH_OIDC_VERIFY_SSL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsOpenIDConnect
func (o *settingsOpenIdconnectTerraformModel) BodyRequest() (req settingsOpenIdconnectBodyRequestModel) {
	req.SOCIAL_AUTH_OIDC_KEY = o.SOCIAL_AUTH_OIDC_KEY.ValueString()
	req.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT = o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT.ValueString()
	req.SOCIAL_AUTH_OIDC_SECRET = o.SOCIAL_AUTH_OIDC_SECRET.ValueString()
	req.SOCIAL_AUTH_OIDC_VERIFY_SSL = o.SOCIAL_AUTH_OIDC_VERIFY_SSL.ValueBool()
	return
}

func (o *settingsOpenIdconnectTerraformModel) setSocialAuthOidcKey(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_KEY, data, false)
}

func (o *settingsOpenIdconnectTerraformModel) setSocialAuthOidcOidcEndpoint(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT, data, false)
}

func (o *settingsOpenIdconnectTerraformModel) setSocialAuthOidcSecret(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_SECRET, data, false)
}

func (o *settingsOpenIdconnectTerraformModel) setSocialAuthOidcVerifySsl(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.SOCIAL_AUTH_OIDC_VERIFY_SSL, data)
}

func (o *settingsOpenIdconnectTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthOidcKey(data["SOCIAL_AUTH_OIDC_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOidcOidcEndpoint(data["SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOidcSecret(data["SOCIAL_AUTH_OIDC_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOidcVerifySsl(data["SOCIAL_AUTH_OIDC_VERIFY_SSL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsOpenIdconnectBodyRequestModel maps the schema for SettingsOpenIDConnect for creating and updating the data
type settingsOpenIdconnectBodyRequestModel struct {
	// SOCIAL_AUTH_OIDC_KEY "The OIDC key (Client ID) from your IDP."
	SOCIAL_AUTH_OIDC_KEY string `json:"SOCIAL_AUTH_OIDC_KEY,omitempty"`
	// SOCIAL_AUTH_OIDC_OIDC_ENDPOINT "The URL for your OIDC provider including the path up to /.well-known/openid-configuration"
	SOCIAL_AUTH_OIDC_OIDC_ENDPOINT string `json:"SOCIAL_AUTH_OIDC_OIDC_ENDPOINT,omitempty"`
	// SOCIAL_AUTH_OIDC_SECRET "The OIDC secret (Client Secret) from your IDP."
	SOCIAL_AUTH_OIDC_SECRET string `json:"SOCIAL_AUTH_OIDC_SECRET,omitempty"`
	// SOCIAL_AUTH_OIDC_VERIFY_SSL "Verify the OIDC provider ssl certificate."
	SOCIAL_AUTH_OIDC_VERIFY_SSL bool `json:"SOCIAL_AUTH_OIDC_VERIFY_SSL"`
}
