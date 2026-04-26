package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGithubTeamTerraformModel struct {
	SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL     types.String `tfsdk:"social_auth_github_team_callback_url" json:"SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_TEAM_ID               types.String `tfsdk:"social_auth_github_team_id" json:"SOCIAL_AUTH_GITHUB_TEAM_ID"`
	SOCIAL_AUTH_GITHUB_TEAM_KEY              types.String `tfsdk:"social_auth_github_team_key" json:"SOCIAL_AUTH_GITHUB_TEAM_KEY"`
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_team_organization_map" json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_TEAM_SECRET           types.String `tfsdk:"social_auth_github_team_secret" json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET"`
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP         types.String `tfsdk:"social_auth_github_team_team_map" json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"`
}

func (o *settingsAuthGithubTeamTerraformModel) Clone() settingsAuthGithubTeamTerraformModel {
	return *o
}

func (o *settingsAuthGithubTeamTerraformModel) BodyRequest() *settingsAuthGithubTeamBodyRequestModel {
	var req settingsAuthGithubTeamBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_TEAM_ID = o.SOCIAL_AUTH_GITHUB_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP.ValueString())
	return &req
}

func (o *settingsAuthGithubTeamTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_ID, data["SOCIAL_AUTH_GITHUB_TEAM_ID"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_KEY, data["SOCIAL_AUTH_GITHUB_TEAM_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_SECRET, data["SOCIAL_AUTH_GITHUB_TEAM_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"], false))
	return diags, nil
}

type settingsAuthGithubTeamBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_TEAM_ID               string          `json:"SOCIAL_AUTH_GITHUB_TEAM_ID,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_KEY              string          `json:"SOCIAL_AUTH_GITHUB_TEAM_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP,omitempty"`
}
