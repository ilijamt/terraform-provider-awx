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

type inventoryResource = framework.GenericResource[inventoryTerraformModel, inventoryBodyRequestModel, *inventoryTerraformModel]

// NewInventoryResource is a helper function to simplify the provider implementation.
func NewInventoryResource() resource.Resource {
	return &inventoryResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "inventory", Endpoint: "/api/v2/inventories/"}},
		Cfg: framework.ResourceCfg[inventoryTerraformModel, inventoryBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this inventory.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"host_filter": schema.StringAttribute{
						Description: "Filter that will be applied to the hosts of this inventory.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"kind": schema.StringAttribute{
						Description: "Kind of inventory being represented.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"smart",
								"constructed",
							),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this inventory.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization containing this inventory.",
						Required:    true,
					},
					"prevent_instance_group_fallback": schema.BoolAttribute{
						Description: "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"variables": schema.StringAttribute{
						Description: "Inventory variables in JSON format",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"has_active_failures": schema.BoolAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Flag indicating whether any hosts in this inventory have failed.",
						Computed:           true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"has_inventory_sources": schema.BoolAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Flag indicating whether this inventory has any external inventory sources.",
						Computed:           true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"hosts_with_active_failures": schema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Number of hosts in this inventory with active failures.",
						Computed:           true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this inventory.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"inventory_sources_with_failures": schema.Int64Attribute{
						Description: "Number of external inventory sources in this inventory with failures.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"pending_deletion": schema.BoolAttribute{
						Description: "Flag indicating the inventory is being deleted.",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"total_groups": schema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Total number of groups in this inventory.",
						Computed:           true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"total_hosts": schema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Total number of hosts in this inventory.",
						Computed:           true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"total_inventory_sources": schema.Int64Attribute{
						Description: "Total number of external inventory sources configured within this inventory.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *inventoryTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Inventory",
		},
	}
}
