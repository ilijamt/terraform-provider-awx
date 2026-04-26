package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGithubOrgTerraformModel struct {
	SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL     types.String `tfsdk:"social_auth_github_org_callback_url" json:"SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_ORG_KEY              types.String `tfsdk:"social_auth_github_org_key" json:"SOCIAL_AUTH_GITHUB_ORG_KEY"`
	SOCIAL_AUTH_GITHUB_ORG_NAME             types.String `tfsdk:"social_auth_github_org_name" json:"SOCIAL_AUTH_GITHUB_ORG_NAME"`
	SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_org_organization_map" json:"SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_ORG_SECRET           types.String `tfsdk:"social_auth_github_org_secret" json:"SOCIAL_AUTH_GITHUB_ORG_SECRET"`
	SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP         types.String `tfsdk:"social_auth_github_org_team_map" json:"SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP"`
}

func (o *settingsAuthGithubOrgTerraformModel) Clone() settingsAuthGithubOrgTerraformModel {
	return *o
}

func (o *settingsAuthGithubOrgTerraformModel) BodyRequest() *settingsAuthGithubOrgBodyRequestModel {
	var req settingsAuthGithubOrgBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_ORG_KEY = o.SOCIAL_AUTH_GITHUB_ORG_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_NAME = o.SOCIAL_AUTH_GITHUB_ORG_NAME.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ORG_SECRET = o.SOCIAL_AUTH_GITHUB_ORG_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP.ValueString())
	return &req
}

func (o *settingsAuthGithubOrgTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_KEY, data["SOCIAL_AUTH_GITHUB_ORG_KEY"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_NAME, data["SOCIAL_AUTH_GITHUB_ORG_NAME"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_SECRET, data["SOCIAL_AUTH_GITHUB_ORG_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP"], false))
	return diags, nil
}

type settingsAuthGithubOrgBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_ORG_KEY              string          `json:"SOCIAL_AUTH_GITHUB_ORG_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_ORG_NAME             string          `json:"SOCIAL_AUTH_GITHUB_ORG_NAME,omitempty"`
	SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ORG_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_ORG_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP,omitempty"`
}
