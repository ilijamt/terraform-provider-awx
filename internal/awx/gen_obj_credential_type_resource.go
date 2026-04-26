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

type credentialTypeResource = framework.GenericResource[credentialTypeTerraformModel, credentialTypeBodyRequestModel, *credentialTypeTerraformModel]

// NewCredentialTypeResource is a helper function to simplify the provider implementation.
func NewCredentialTypeResource() resource.Resource {
	return &credentialTypeResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_type", Endpoint: "/api/v2/credential_types/"}},
		Cfg: framework.ResourceCfg[credentialTypeTerraformModel, credentialTypeBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this credential type.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"injectors": schema.StringAttribute{
						Description: "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"inputs": schema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"kind": schema.StringAttribute{
						Description: "The credential type",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"net",
								"cloud",
							),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this credential type.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential type.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"managed": schema.BoolAttribute{
						Description: "Is the resource managed",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"namespace": schema.StringAttribute{
						Description: "The namespace to which the resource belongs to",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *credentialTypeTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialType",
		},
	}
}
