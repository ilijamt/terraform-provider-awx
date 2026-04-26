package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsAuthGithubEnterpriseResource = framework.GenericResource[settingsAuthGithubEnterpriseTerraformModel, settingsAuthGithubEnterpriseBodyRequestModel, *settingsAuthGithubEnterpriseTerraformModel]

// NewSettingsAuthGithubEnterpriseResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubEnterpriseResource() resource.Resource {
	return &settingsAuthGithubEnterpriseResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github_enterprise", Endpoint: "/api/v2/settings/github-enterprise/"}},
		Cfg: framework.ResourceCfg[settingsAuthGithubEnterpriseTerraformModel, settingsAuthGithubEnterpriseBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_github_enterprise_api_url": schema.StringAttribute{
						Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub Enterprise developer application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_url": schema.StringAttribute{
						Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_github_enterprise_callback_url": schema.StringAttribute{
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
			Hook:         hookSettingsAuthGithubEnterprise,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGithubEnterprise",
		},
	}
}
