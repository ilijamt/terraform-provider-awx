package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGithubTerraformModel struct {
	SOCIAL_AUTH_GITHUB_CALLBACK_URL     types.String `tfsdk:"social_auth_github_callback_url" json:"SOCIAL_AUTH_GITHUB_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_KEY              types.String `tfsdk:"social_auth_github_key" json:"SOCIAL_AUTH_GITHUB_KEY"`
	SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_organization_map" json:"SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_SECRET           types.String `tfsdk:"social_auth_github_secret" json:"SOCIAL_AUTH_GITHUB_SECRET"`
	SOCIAL_AUTH_GITHUB_TEAM_MAP         types.String `tfsdk:"social_auth_github_team_map" json:"SOCIAL_AUTH_GITHUB_TEAM_MAP"`
}

func (o *settingsAuthGithubTerraformModel) Clone() settingsAuthGithubTerraformModel {
	return *o
}

func (o *settingsAuthGithubTerraformModel) BodyRequest() *settingsAuthGithubBodyRequestModel {
	var req settingsAuthGithubBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_KEY = o.SOCIAL_AUTH_GITHUB_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_SECRET = o.SOCIAL_AUTH_GITHUB_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_MAP.ValueString())
	return &req
}

func (o *settingsAuthGithubTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_KEY, data["SOCIAL_AUTH_GITHUB_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_SECRET, data["SOCIAL_AUTH_GITHUB_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_TEAM_MAP"], false))
	return diags, nil
}

type settingsAuthGithubBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_KEY              string          `json:"SOCIAL_AUTH_GITHUB_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_MAP,omitempty"`
}
