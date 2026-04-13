package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsAuthGithubResource = framework.GenericResource[settingsAuthGithubTerraformModel, settingsAuthGithubBodyRequestModel, *settingsAuthGithubTerraformModel]

// NewSettingsAuthGithubResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubResource() resource.Resource {
	return &settingsAuthGithubResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_github", Endpoint: "/api/v2/settings/github/"}},
		Cfg: framework.ResourceCfg[settingsAuthGithubTerraformModel, settingsAuthGithubBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"social_auth_github_key": schema.StringAttribute{
						Description: "The OAuth2 key (Client ID) from your GitHub developer application.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_github_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_github_secret": schema.StringAttribute{
						Description: "The OAuth2 secret (Client Secret) from your GitHub developer application.",
						Sensitive:   true,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_github_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					// Write only elements
					// Data only elements
					"social_auth_github_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
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
