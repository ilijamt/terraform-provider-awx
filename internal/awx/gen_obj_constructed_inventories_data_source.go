package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type constructedInventoriesDataSource = framework.GenericDataSource[constructedInventoriesTerraformModel, *constructedInventoriesTerraformModel]

// NewConstructedInventoriesDataSource is a helper function to instantiate the ConstructedInventories data source.
func NewConstructedInventoriesDataSource() datasource.DataSource {
	return &constructedInventoriesDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "constructed_inventories", Endpoint: "/api/v2/constructed_inventories/"}},
		Cfg: framework.DataSourceCfg[constructedInventoriesTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"description": schema.StringAttribute{
						Description: "Optional description of this inventory.",
						Sensitive:   false,
						Computed:    true,
					},
					"has_active_failures": schema.BoolAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Flag indicating whether any hosts in this inventory have failed.",
						Sensitive:          false,
						Computed:           true,
					},
					"has_inventory_sources": schema.BoolAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Flag indicating whether this inventory has any external inventory sources.",
						Sensitive:          false,
						Computed:           true,
					},
					"hosts_with_active_failures": schema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Number of hosts in this inventory with active failures.",
						Sensitive:          false,
						Computed:           true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this inventory.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"inventory_sources_with_failures": schema.Int64Attribute{
						Description: "Number of external inventory sources in this inventory with failures.",
						Sensitive:   false,
						Computed:    true,
					},
					"kind": schema.StringAttribute{
						Description: "Kind of inventory being represented.",
						Sensitive:   false,
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "The limit to restrict the returned hosts for the related auto-created inventory source, special to constructed inventory.",
						Sensitive:   false,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this inventory.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization containing this inventory.",
						Sensitive:   false,
						Computed:    true,
					},
					"pending_deletion": schema.BoolAttribute{
						Description: "Flag indicating the inventory is being deleted.",
						Sensitive:   false,
						Computed:    true,
					},
					"prevent_instance_group_fallback": schema.BoolAttribute{
						Description: "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
						Sensitive:   false,
						Computed:    true,
					},
					"source_vars": schema.StringAttribute{
						Description: "The source_vars for the related auto-created inventory source, special to constructed inventory.",
						Sensitive:   false,
						Computed:    true,
					},
					"total_groups": schema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Total number of groups in this inventory.",
						Sensitive:          false,
						Computed:           true,
					},
					"total_hosts": schema.Int64Attribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Total number of hosts in this inventory.",
						Sensitive:          false,
						Computed:           true,
					},
					"total_inventory_sources": schema.Int64Attribute{
						Description: "Total number of external inventory sources configured within this inventory.",
						Sensitive:   false,
						Computed:    true,
					},
					"update_cache_timeout": schema.Int64Attribute{
						Description: "The cache timeout for the related auto-created inventory source, special to constructed inventory",
						Sensitive:   false,
						Computed:    true,
					},
					"variables": schema.StringAttribute{
						Description: "Inventory variables in JSON or YAML format.",
						Sensitive:   false,
						Computed:    true,
					},
					"verbosity": schema.Int64Attribute{
						Description: "The verbosity level for the related auto-created inventory source, special to constructed inventory",
						Sensitive:   false,
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
