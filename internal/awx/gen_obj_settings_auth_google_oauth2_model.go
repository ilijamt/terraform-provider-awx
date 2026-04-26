package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGoogleOauth2TerraformModel struct {
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS types.String `tfsdk:"social_auth_google_oauth2_auth_extra_arguments" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL         types.String `tfsdk:"social_auth_google_oauth2_callback_url" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY                  types.String `tfsdk:"social_auth_google_oauth2_key" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP     types.String `tfsdk:"social_auth_google_oauth2_organization_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET               types.String `tfsdk:"social_auth_google_oauth2_secret" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP             types.String `tfsdk:"social_auth_google_oauth2_team_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS  types.List   `tfsdk:"social_auth_google_oauth2_whitelisted_domains" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"`
}

func (o *settingsAuthGoogleOauth2TerraformModel) Clone() settingsAuthGoogleOauth2TerraformModel {
	return *o
}

func (o *settingsAuthGoogleOauth2TerraformModel) BodyRequest() *settingsAuthGoogleOauth2BodyRequestModel {
	var req settingsAuthGoogleOauth2BodyRequestModel
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY = o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET = o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = helpers.ListAsStringSlice(o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, false)
	return &req
}

func (o *settingsAuthGoogleOauth2TerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS, data["SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL, data["SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY, data["SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP, data["SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET, data["SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP, data["SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"], false))
	collect(helpers.AttrValueSetListString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, data["SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"], false))
	return diags, nil
}

type settingsAuthGoogleOauth2BodyRequestModel struct {
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY                  string          `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP     json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET               string          `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP             json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS  []string        `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS,omitempty"`
}
