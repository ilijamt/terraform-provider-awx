package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type tokensResource = framework.GenericResource[tokensTerraformModel, tokensBodyRequestModel, *tokensTerraformModel]

// NewTokensResource is a helper function to simplify the provider implementation.
func NewTokensResource() resource.Resource {
	return &tokensResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "token", Endpoint: "/api/v2/tokens/"}},
		Cfg: framework.ResourceCfg[tokensTerraformModel, tokensBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"application": schema.Int64Attribute{
						Description: "Application",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this access token.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"scope": schema.StringAttribute{
						Description: "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`write`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"expires": schema.StringAttribute{
						Description: "Expires",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this access token.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"refresh_token": schema.StringAttribute{
						Description: "Refresh token",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"token": schema.StringAttribute{
						Description: "Token",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"user": schema.Int64Attribute{
						Description: "The user representing the token owner",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *tokensTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Tokens",
		},
	}
}
