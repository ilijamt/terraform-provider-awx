package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
					// Request elements
					"email": schema.StringAttribute{
						Description: "Email address",
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"is_system_auditor": schema.BoolAttribute{
						Description: "Is system auditor",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"last_name": schema.StringAttribute{
						Description: "Last name",
						Sensitive:   false,
						Required:    false,
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
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"username": schema.StringAttribute{
						Description:   "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					// Write only elements
					// Data only elements
					"external_account": schema.StringAttribute{
						Description: "Set if the account is managed by an external service",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this user.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_login": schema.StringAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"ldap_dn": schema.StringAttribute{
						Description: "",
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
			IDAccessor:      func(m *userTerraformModel) any { return m.ID.ValueInt64() },
			ImportStateFunc: userResourceImportState,
			Hook:            hookUser,
			ApiVersion:      ApiVersion,
			ResourceName:    "User",
		},
	}
}

func userResourceImportState(_ context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the User.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(context.Background(), path.Root("id"), id)...)
}
