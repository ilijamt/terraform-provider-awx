package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGithubEnterpriseOrgTerraformModel struct {
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL          types.String `tfsdk:"social_auth_github_enterprise_org_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL     types.String `tfsdk:"social_auth_github_enterprise_org_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY              types.String `tfsdk:"social_auth_github_enterprise_org_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME             types.String `tfsdk:"social_auth_github_enterprise_org_name" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_org_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET           types.String `tfsdk:"social_auth_github_enterprise_org_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP         types.String `tfsdk:"social_auth_github_enterprise_org_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL              types.String `tfsdk:"social_auth_github_enterprise_org_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL"`
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) Clone() settingsAuthGithubEnterpriseOrgTerraformModel {
	return *o
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) BodyRequest() *settingsAuthGithubEnterpriseOrgBodyRequestModel {
	var req settingsAuthGithubEnterpriseOrgBodyRequestModel
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL.ValueString()
	return &req
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL, data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL"], false))
	return diags, nil
}

type settingsAuthGithubEnterpriseOrgBodyRequestModel struct {
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL          string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY              string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME             string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET           string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP         json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL              string          `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL,omitempty"`
}

type settingsAuthGithubEnterpriseOrgResource = framework.GenericResource[settingsAuthGithubEnterpriseOrgTerraformModel, settingsAuthGithubEnterpriseOrgBodyRequestModel, *settingsAuthGithubEnterpriseOrgTerraformModel]

// NewSettingsAuthGithubEnterpriseOrgResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubEnterpriseOrgResource() resource.Resource {
	return &settingsAuthGithubEnterpriseOrgResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise_org", Endpoint: "/api/v2/settings/github-enterprise-org/"}},
		Cfg: framework.ResourceCfg[settingsAuthGithubEnterpriseOrgTerraformModel, settingsAuthGithubEnterpriseOrgBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_enterprise_org_api_url": schema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_name": schema.StringAttribute{
						Description: "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_url": schema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_org_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsAuthGithubEnterpriseOrg,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterpriseOrg",
		},
	}
}

type settingsAuthGithubEnterpriseOrgDataSource = framework.GenericDataSource[settingsAuthGithubEnterpriseOrgTerraformModel, *settingsAuthGithubEnterpriseOrgTerraformModel]

// NewSettingsAuthGithubEnterpriseOrgDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseOrg data source.
func NewSettingsAuthGithubEnterpriseOrgDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseOrgDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise_org", Endpoint: "/api/v2/settings/github-enterprise-org/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubEnterpriseOrgTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_github_enterprise_org_api_url": dschema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
					"social_auth_github_enterprise_org_callback_url": dschema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
					},
					"social_auth_github_enterprise_org_key": dschema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
						Computed:    true,
					},
					"social_auth_github_enterprise_org_name": dschema.StringAttribute{
						Description: "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/.",
						Computed:    true,
					},
					"social_auth_github_enterprise_org_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_org_secret": dschema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_enterprise_org_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_github_enterprise_org_url": dschema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithubEnterpriseOrg,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterpriseOrg",
		},
	}
}
