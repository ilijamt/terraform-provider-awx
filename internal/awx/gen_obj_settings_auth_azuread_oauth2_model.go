package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthAzureAdoauth2TerraformModel struct {
	SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL     types.String `tfsdk:"social_auth_azuread_oauth2_callback_url" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_KEY              types.String `tfsdk:"social_auth_azuread_oauth2_key" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP types.String `tfsdk:"social_auth_azuread_oauth2_organization_map" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET           types.String `tfsdk:"social_auth_azuread_oauth2_secret" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP         types.String `tfsdk:"social_auth_azuread_oauth2_team_map" json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"`
}

func (o *settingsAuthAzureAdoauth2TerraformModel) Clone() settingsAuthAzureAdoauth2TerraformModel {
	return *o
}

func (o *settingsAuthAzureAdoauth2TerraformModel) BodyRequest() *settingsAuthAzureAdoauth2BodyRequestModel {
	var req settingsAuthAzureAdoauth2BodyRequestModel
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY = o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET = o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP.ValueString())
	return &req
}

func (o *settingsAuthAzureAdoauth2TerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL, data["SOCIAL_AUTH_AZUREAD_OAUTH2_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_KEY, data["SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP, data["SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET, data["SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP, data["SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"], false))
	return diags, nil
}

type settingsAuthAzureAdoauth2BodyRequestModel struct {
	SOCIAL_AUTH_AZUREAD_OAUTH2_KEY              string          `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY,omitempty"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET           string          `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET,omitempty"`
	SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP,omitempty"`
}
