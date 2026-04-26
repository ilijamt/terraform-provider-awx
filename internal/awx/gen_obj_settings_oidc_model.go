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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsOpenIDConnect
func (o *settingsOpenIdconnectTerraformModel) BodyRequest() *settingsOpenIdconnectBodyRequestModel {
	var req settingsOpenIdconnectBodyRequestModel
	req.SOCIAL_AUTH_OIDC_KEY = o.SOCIAL_AUTH_OIDC_KEY.ValueString()
	req.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT = o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT.ValueString()
	req.SOCIAL_AUTH_OIDC_SECRET = o.SOCIAL_AUTH_OIDC_SECRET.ValueString()
	req.SOCIAL_AUTH_OIDC_VERIFY_SSL = o.SOCIAL_AUTH_OIDC_VERIFY_SSL.ValueBool()
	return &req
}

func (o *settingsOpenIdconnectTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_KEY, data["SOCIAL_AUTH_OIDC_KEY"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT, data["SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_SECRET, data["SOCIAL_AUTH_OIDC_SECRET"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.SOCIAL_AUTH_OIDC_VERIFY_SSL, data["SOCIAL_AUTH_OIDC_VERIFY_SSL"])
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
