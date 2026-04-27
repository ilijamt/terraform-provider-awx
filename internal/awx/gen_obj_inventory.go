package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type inventoryTerraformModel struct {
	Description                  types.String `tfsdk:"description" json:"description"`
	HasActiveFailures            types.Bool   `tfsdk:"has_active_failures" json:"has_active_failures"`
	HasInventorySources          types.Bool   `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	HostFilter                   types.String `tfsdk:"host_filter" json:"host_filter"`
	HostsWithActiveFailures      types.Int64  `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	ID                           types.Int64  `tfsdk:"id" json:"id"`
	InventorySourcesWithFailures types.Int64  `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	Kind                         types.String `tfsdk:"kind" json:"kind"`
	Name                         types.String `tfsdk:"name" json:"name"`
	Organization                 types.Int64  `tfsdk:"organization" json:"organization"`
	PendingDeletion              types.Bool   `tfsdk:"pending_deletion" json:"pending_deletion"`
	PreventInstanceGroupFallback types.Bool   `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	TotalGroups                  types.Int64  `tfsdk:"total_groups" json:"total_groups"`
	TotalHosts                   types.Int64  `tfsdk:"total_hosts" json:"total_hosts"`
	TotalInventorySources        types.Int64  `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	Variables                    types.String `tfsdk:"variables" json:"variables"`
}

func (o *inventoryTerraformModel) Clone() inventoryTerraformModel {
	return *o
}

func (o *inventoryTerraformModel) BodyRequest() *inventoryBodyRequestModel {
	var req inventoryBodyRequestModel
	req.Description = o.Description.ValueString()
	req.HostFilter = o.HostFilter.ValueString()
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.Variables = json.RawMessage(o.Variables.String())
	return &req
}

func (o *inventoryTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.HasActiveFailures, data["has_active_failures"]))
	collect(helpers.AttrValueSetBool(&o.HasInventorySources, data["has_inventory_sources"]))
	collect(helpers.AttrValueSetString(&o.HostFilter, data["host_filter"], false))
	collect(helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data["hosts_with_active_failures"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data["inventory_sources_with_failures"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetBool(&o.PendingDeletion, data["pending_deletion"]))
	collect(helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data["prevent_instance_group_fallback"]))
	collect(helpers.AttrValueSetInt64(&o.TotalGroups, data["total_groups"]))
	collect(helpers.AttrValueSetInt64(&o.TotalHosts, data["total_hosts"]))
	collect(helpers.AttrValueSetInt64(&o.TotalInventorySources, data["total_inventory_sources"]))
	collect(helpers.AttrValueSetJsonYamlString(&o.Variables, data["variables"], false))
	return diags, nil
}

type inventoryBodyRequestModel struct {
	Description                  string          `json:"description,omitempty"`
	HostFilter                   string          `json:"host_filter,omitempty"`
	Kind                         string          `json:"kind,omitempty"`
	Name                         string          `json:"name"`
	Organization                 int64           `json:"organization"`
	PreventInstanceGroupFallback bool            `json:"prevent_instance_group_fallback"`
	Variables                    json.RawMessage `json:"variables,omitempty"`
}

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

type inventoryDataSource = framework.GenericDataSource[inventoryTerraformModel, *inventoryTerraformModel]

// NewInventoryDataSource is a helper function to instantiate the Inventory data source.
func NewInventoryDataSource() datasource.DataSource {
	return &inventoryDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "inventory", Endpoint: "/api/v2/inventories/"}},
		Cfg: framework.DataSourceCfg[inventoryTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this inventory.",
						Computed:    true,
					},
					"has_active_failures": dschema.BoolAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Flag indicating whether any hosts in this inventory have failed.",
						Computed:           true,
					},
					"has_inventory_sources": dschema.BoolAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Flag indicating whether this inventory has any external inventory sources.",
						Computed:           true,
					},
					"host_filter": dschema.StringAttribute{
						Description: "Filter that will be applied to the hosts of this inventory.",
						Computed:    true,
					},
					"hosts_with_active_failures": dschema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Number of hosts in this inventory with active failures.",
						Computed:           true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this inventory.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"inventory_sources_with_failures": dschema.Int64Attribute{
						Description: "Number of external inventory sources in this inventory with failures.",
						Computed:    true,
					},
					"kind": dschema.StringAttribute{
						Description: "Kind of inventory being represented.",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this inventory.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"organization": dschema.Int64Attribute{
						Description: "Organization containing this inventory.",
						Computed:    true,
					},
					"pending_deletion": dschema.BoolAttribute{
						Description: "Flag indicating the inventory is being deleted.",
						Computed:    true,
					},
					"prevent_instance_group_fallback": dschema.BoolAttribute{
						Description: "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
						Computed:    true,
					},
					"total_groups": dschema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Total number of groups in this inventory.",
						Computed:           true,
					},
					"total_hosts": dschema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Total number of hosts in this inventory.",
						Computed:           true,
					},
					"total_inventory_sources": dschema.Int64Attribute{
						Description: "Total number of external inventory sources configured within this inventory.",
						Computed:    true,
					},
					"variables": dschema.StringAttribute{
						Description: "Inventory variables in JSON format",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Inventory",
		},
	}
}
