package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGithubEnterpriseTeamTerraformModel struct {
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL          types.String `tfsdk:"social_auth_github_enterprise_team_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL     types.String `tfsdk:"social_auth_github_enterprise_team_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID               types.String `tfsdk:"social_auth_github_enterprise_team_id" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY              types.String `tfsdk:"social_auth_github_enterprise_team_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_team_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET           types.String `tfsdk:"social_auth_github_enterprise_team_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP         types.String `tfsdk:"social_auth_github_enterprise_team_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL              types.String `tfsdk:"social_auth_github_enterprise_team_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"`
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) Clone() settingsAuthGithubEnterpriseTeamTerraformModel {
	return *o
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) BodyRequest() *settingsAuthGithubEnterpriseTeamBodyRequestModel {
	var req settingsAuthGithubEnterpriseTeamBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL.ValueString()
	return &req
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"], false))
	return diags, nil
}

type settingsAuthGithubEnterpriseTeamBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL          string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID               string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY              string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL              string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL,omitempty"`
}
