package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type userResource = framework.GenericResource[userTerraformModel, userBodyRequestModel, *userTerraformModel]

// NewUserResource is a helper function to simplify the provider implementation.
func NewUserResource() resource.Resource {
	return &userResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "user", Endpoint: "/api/v2/users/"}},
		Cfg: framework.ResourceCfg[userTerraformModel, userBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description: "Email address",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(254),
						},
					},
					"first_name": schema.StringAttribute{
						Description: "First name",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					"is_superuser": schema.BoolAttribute{
						Description: "Designates that this user has all permissions without explicitly assigning them.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"is_system_auditor": schema.BoolAttribute{
						Description: "Is system auditor",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"last_name": schema.StringAttribute{
						Description: "Last name",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					"password": schema.StringAttribute{
						Description: "Field used to change the password.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"username": schema.StringAttribute{
						Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					"external_account": schema.StringAttribute{
						Description: "Set if the account is managed by an external service",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this user.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_login": schema.StringAttribute{
						Description: "Last login",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"ldap_dn": schema.StringAttribute{
						Description: "Ldap dn",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *userTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			Hook:         hookUser,
			ApiVersion:   ApiVersion,
			ResourceName: "User",
		},
	}
}
