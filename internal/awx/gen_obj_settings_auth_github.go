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

type settingsAuthGithubResource = framework.GenericResource[settingsAuthGithubTerraformModel, settingsAuthGithubBodyRequestModel, *settingsAuthGithubTerraformModel]

// NewSettingsAuthGithubResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubResource() resource.Resource {
	return &settingsAuthGithubResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github", Endpoint: "/api/v2/settings/github/"}},
		Cfg: framework.ResourceCfg[settingsAuthGithubTerraformModel, settingsAuthGithubBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub developer application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub developer application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_callback_url": schema.StringAttribute{
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
			Hook:         hookSettingsAuthGithub,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithub",
		},
	}
}

type settingsAuthGithubDataSource = framework.GenericDataSource[settingsAuthGithubTerraformModel, *settingsAuthGithubTerraformModel]

// NewSettingsAuthGithubDataSource is a helper function to instantiate the SettingsAuthGithub data source.
func NewSettingsAuthGithubDataSource() datasource.DataSource {
	return &settingsAuthGithubDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github", Endpoint: "/api/v2/settings/github/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGithubTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_github_callback_url": dschema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
					},
					"social_auth_github_key": dschema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub developer application.",
						Computed:    true,
					},
					"social_auth_github_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_github_secret": dschema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub developer application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_github_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGithub,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithub",
		},
	}
}
