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

type constructedInventoriesTerraformModel struct {
	Description                  types.String `tfsdk:"description" json:"description"`
	HasActiveFailures            types.Bool   `tfsdk:"has_active_failures" json:"has_active_failures"`
	HasInventorySources          types.Bool   `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	HostsWithActiveFailures      types.Int64  `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	ID                           types.Int64  `tfsdk:"id" json:"id"`
	InventorySourcesWithFailures types.Int64  `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	Kind                         types.String `tfsdk:"kind" json:"kind"`
	Limit                        types.String `tfsdk:"limit" json:"limit"`
	Name                         types.String `tfsdk:"name" json:"name"`
	Organization                 types.Int64  `tfsdk:"organization" json:"organization"`
	PendingDeletion              types.Bool   `tfsdk:"pending_deletion" json:"pending_deletion"`
	PreventInstanceGroupFallback types.Bool   `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	SourceVars                   types.String `tfsdk:"source_vars" json:"source_vars"`
	TotalGroups                  types.Int64  `tfsdk:"total_groups" json:"total_groups"`
	TotalHosts                   types.Int64  `tfsdk:"total_hosts" json:"total_hosts"`
	TotalInventorySources        types.Int64  `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	UpdateCacheTimeout           types.Int64  `tfsdk:"update_cache_timeout" json:"update_cache_timeout"`
	Variables                    types.String `tfsdk:"variables" json:"variables"`
	Verbosity                    types.Int64  `tfsdk:"verbosity" json:"verbosity"`
}

func (o *constructedInventoriesTerraformModel) Clone() constructedInventoriesTerraformModel {
	return *o
}

func (o *constructedInventoriesTerraformModel) BodyRequest() *constructedInventoriesBodyRequestModel {
	var req constructedInventoriesBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.SourceVars = o.SourceVars.ValueString()
	req.UpdateCacheTimeout = o.UpdateCacheTimeout.ValueInt64()
	req.Variables = json.RawMessage(o.Variables.ValueString())
	req.Verbosity = o.Verbosity.ValueInt64()
	return &req
}

func (o *constructedInventoriesTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.HasActiveFailures, data["has_active_failures"]))
	collect(helpers.AttrValueSetBool(&o.HasInventorySources, data["has_inventory_sources"]))
	collect(helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data["hosts_with_active_failures"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data["inventory_sources_with_failures"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetBool(&o.PendingDeletion, data["pending_deletion"]))
	collect(helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data["prevent_instance_group_fallback"]))
	collect(helpers.AttrValueSetString(&o.SourceVars, data["source_vars"], false))
	collect(helpers.AttrValueSetInt64(&o.TotalGroups, data["total_groups"]))
	collect(helpers.AttrValueSetInt64(&o.TotalHosts, data["total_hosts"]))
	collect(helpers.AttrValueSetInt64(&o.TotalInventorySources, data["total_inventory_sources"]))
	collect(helpers.AttrValueSetInt64(&o.UpdateCacheTimeout, data["update_cache_timeout"]))
	collect(helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false))
	collect(helpers.AttrValueSetInt64(&o.Verbosity, data["verbosity"]))
	return diags, nil
}

type constructedInventoriesBodyRequestModel struct {
	Description                  string          `json:"description,omitempty"`
	Limit                        string          `json:"limit,omitempty"`
	Name                         string          `json:"name"`
	Organization                 int64           `json:"organization"`
	PreventInstanceGroupFallback bool            `json:"prevent_instance_group_fallback"`
	SourceVars                   string          `json:"source_vars,omitempty"`
	UpdateCacheTimeout           int64           `json:"update_cache_timeout,omitempty"`
	Variables                    json.RawMessage `json:"variables,omitempty"`
	Verbosity                    int64           `json:"verbosity,omitempty"`
}

type constructedInventoriesResource = framework.GenericResource[constructedInventoriesTerraformModel, constructedInventoriesBodyRequestModel, *constructedInventoriesTerraformModel]

// NewConstructedInventoriesResource is a helper function to simplify the provider implementation.
func NewConstructedInventoriesResource() resource.Resource {
	return &constructedInventoriesResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "constructed_inventories", Endpoint: "/api/v2/constructed_inventories/"}},
		Cfg: framework.ResourceCfg[constructedInventoriesTerraformModel, constructedInventoriesBodyRequestModel]{
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
					"limit": schema.StringAttribute{
						Description: "The limit to restrict the returned hosts for the related auto-created inventory source, special to constructed inventory.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"source_vars": schema.StringAttribute{
						Description: "The source_vars for the related auto-created inventory source, special to constructed inventory.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"update_cache_timeout": schema.Int64Attribute{
						Description: "The cache timeout for the related auto-created inventory source, special to constructed inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"variables": schema.StringAttribute{
						Description: "Inventory variables in JSON or YAML format.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"verbosity": schema.Int64Attribute{
						Description: "The verbosity level for the related auto-created inventory source, special to constructed inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 2),
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
					"kind": schema.StringAttribute{
						Description: "Kind of inventory being represented.",
						Computed:    true,
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
			IDAccessor:   func(m *constructedInventoriesTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "ConstructedInventories",
		},
	}
}

type constructedInventoriesDataSource = framework.GenericDataSource[constructedInventoriesTerraformModel, *constructedInventoriesTerraformModel]

// NewConstructedInventoriesDataSource is a helper function to instantiate the ConstructedInventories data source.
func NewConstructedInventoriesDataSource() datasource.DataSource {
	return &constructedInventoriesDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "constructed_inventories", Endpoint: "/api/v2/constructed_inventories/"}},
		Cfg: framework.DataSourceCfg[constructedInventoriesTerraformModel]{
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
					"limit": dschema.StringAttribute{
						Description: "The limit to restrict the returned hosts for the related auto-created inventory source, special to constructed inventory.",
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
					"source_vars": dschema.StringAttribute{
						Description: "The source_vars for the related auto-created inventory source, special to constructed inventory.",
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
					"update_cache_timeout": dschema.Int64Attribute{
						Description: "The cache timeout for the related auto-created inventory source, special to constructed inventory",
						Computed:    true,
					},
					"variables": dschema.StringAttribute{
						Description: "Inventory variables in JSON or YAML format.",
						Computed:    true,
					},
					"verbosity": dschema.Int64Attribute{
						Description: "The verbosity level for the related auto-created inventory source, special to constructed inventory",
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
			ResourceName: "ConstructedInventories",
		},
	}
}
