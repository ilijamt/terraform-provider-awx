package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsAuthGoogleOauth2Resource = framework.GenericResource[settingsAuthGoogleOauth2TerraformModel, settingsAuthGoogleOauth2BodyRequestModel, *settingsAuthGoogleOauth2TerraformModel]

// NewSettingsAuthGoogleOauth2Resource is a helper function to simplify the provider implementation.
func NewSettingsAuthGoogleOauth2Resource() resource.Resource {
	return &settingsAuthGoogleOauth2Resource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_google_oauth2", Endpoint: "/api/v2/settings/google-oauth2/"}},
		Cfg: framework.ResourceCfg[settingsAuthGoogleOauth2TerraformModel, settingsAuthGoogleOauth2BodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_google_oauth2_auth_extra_arguments": schema.StringAttribute{
						Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_key": schema.StringAttribute{
						Description: "The OAuth2 key from your web application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_secret": schema.StringAttribute{
						Description: "The OAuth2 secret from your web application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_whitelisted_domains": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_callback_url": schema.StringAttribute{
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
			Hook:         hookSettingsAuthGoogleOauth2,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGoogleOauth2",
		},
	}
}
