package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsOpenIdconnectResource = framework.GenericResource[settingsOpenIdconnectTerraformModel, settingsOpenIdconnectBodyRequestModel, *settingsOpenIdconnectTerraformModel]

// NewSettingsOpenIDConnectResource is a helper function to simplify the provider implementation.
func NewSettingsOpenIDConnectResource() resource.Resource {
	return &settingsOpenIdconnectResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_oidc", Endpoint: "/api/v2/settings/oidc/"}},
		Cfg: framework.ResourceCfg[settingsOpenIdconnectTerraformModel, settingsOpenIdconnectBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"social_auth_oidc_key": schema.StringAttribute{
						Description: "The OIDC key (Client ID) from your IDP.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_oidc_oidc_endpoint": schema.StringAttribute{
						Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
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
					"social_auth_oidc_secret": schema.StringAttribute{
						Description: "The OIDC secret (Client Secret) from your IDP.",
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
					"social_auth_oidc_verify_ssl": schema.BoolAttribute{
						Description: "Verify the OIDC provider ssl certificate.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					// Write only elements
					// Data only elements
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsOidc,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsOpenIDConnect",
		},
	}
}
