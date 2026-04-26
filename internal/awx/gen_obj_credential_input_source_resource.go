package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type credentialInputSourceResource = framework.GenericResource[credentialInputSourceTerraformModel, credentialInputSourceBodyRequestModel, *credentialInputSourceTerraformModel]

// NewCredentialInputSourceResource is a helper function to simplify the provider implementation.
func NewCredentialInputSourceResource() resource.Resource {
	return &credentialInputSourceResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_input_source", Endpoint: "/api/v2/credential_input_sources/"}},
		Cfg: framework.ResourceCfg[credentialInputSourceTerraformModel, credentialInputSourceBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this credential input source.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"input_field_name": schema.StringAttribute{
						Description: "Input field name",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"metadata": schema.StringAttribute{
						Description: "Metadata",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"source_credential": schema.Int64Attribute{
						Description: "Source credential",
						Required:    true,
					},
					"target_credential": schema.Int64Attribute{
						Description: "Target credential",
						Required:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential input source.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *credentialInputSourceTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialInputSource",
		},
	}
}
