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

type applicationResource = framework.GenericResource[applicationTerraformModel, applicationBodyRequestModel, *applicationTerraformModel]

// NewApplicationResource is a helper function to simplify the provider implementation.
func NewApplicationResource() resource.Resource {
	return &applicationResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "application", Endpoint: "/api/v2/applications/"}},
		Cfg: framework.ResourceCfg[applicationTerraformModel, applicationBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"authorization_grant_type": schema.StringAttribute{
						Description: "The Grant type the user must use for acquire tokens for this application.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"authorization-code",
								"password",
							),
						},
					},
					"client_type": schema.StringAttribute{
						Description: "Set to Public or Confidential depending on how secure the client device is.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"confidential",
								"public",
							),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this application.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization containing this application.",
						Required:    true,
					},
					"redirect_uris": schema.StringAttribute{
						Description: "Allowed URIs list, space separated",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"skip_authorization": schema.BoolAttribute{
						Description: "Set True to skip authorization step for completely trusted applications.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"client_id": schema.StringAttribute{
						Description: "Client id",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"client_secret": schema.StringAttribute{
						Description: "Used for more stringent verification of access to an application when creating a token.",
						Sensitive:   true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this application.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *applicationTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			Hook:         hookApplication,
			ApiVersion:   ApiVersion,
			ResourceName: "Application",
		},
	}
}
