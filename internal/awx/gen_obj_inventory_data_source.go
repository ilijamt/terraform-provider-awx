package awx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &inventoryDataSource{}
	_ datasource.DataSourceWithConfigure = &inventoryDataSource{}
)

// NewInventoryDataSource is a helper function to instantiate the Inventory data source.
func NewInventoryDataSource() datasource.DataSource {
	return &inventoryDataSource{}
}

// inventoryDataSource is the data source implementation.
type inventoryDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *inventoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventories/"
}

// Metadata returns the data source type name.
func (o *inventoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory"
}

// Schema defines the schema for the data source.
func (o *inventoryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"description": schema.StringAttribute{
				Description: "Optional description of this inventory.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"has_active_failures": schema.BoolAttribute{
				Description: "This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"has_inventory_sources": schema.BoolAttribute{
				Description: "This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"host_filter": schema.StringAttribute{
				Description: "Filter that will be applied to the hosts of this inventory.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"hosts_with_active_failures": schema.Int64Attribute{
				Description: "This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
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
				Validators:  []validator.Int64{},
			},
			"kind": schema.StringAttribute{
				Description: "Kind of inventory being represented.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"", "smart", "constructed"}...),
				},
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
				Validators:  []validator.Int64{},
			},
			"pending_deletion": schema.BoolAttribute{
				Description: "Flag indicating the inventory is being deleted.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"prevent_instance_group_fallback": schema.BoolAttribute{
				Description: "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"total_groups": schema.Int64Attribute{
				Description: "This field is deprecated and will be removed in a future release. Total number of groups in this inventory.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"total_hosts": schema.Int64Attribute{
				Description: "This field is deprecated and will be removed in a future release. Total number of hosts in this inventory.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"total_inventory_sources": schema.Int64Attribute{
				Description: "Total number of external inventory sources configured within this inventory.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"variables": schema.StringAttribute{
				Description: "Inventory variables in JSON format",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			// Write only elements
		},
	}
}

func (o *inventoryDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *inventoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state inventoryTerraformModel
	var err error
	var endpoint string
	var searchDefined bool

	// Only one group should evaluate to True, terraform should prevent from being able to set
	// the conflicting groups

	// Evaluate group 'by_id' based on the schema definition
	var groupByIdExists = func() bool {
		var groupByIdExists = true
		var paramsById = []any{o.endpoint}
		var attrID types.Int64
		req.Config.GetAttribute(ctx, path.Root("id"), &attrID)
		groupByIdExists = groupByIdExists && (!attrID.IsNull() && !attrID.IsUnknown())
		paramsById = append(paramsById, attrID.ValueInt64())
		if groupByIdExists {
			endpoint = p.Clean(fmt.Sprintf("%s/%d/", paramsById...))
		}
		return groupByIdExists
	}()
	searchDefined = searchDefined || groupByIdExists

	// Evaluate group 'by_name' based on the schema definition
	var groupByNameExists = func() bool {
		var groupByNameExists = true
		var paramsByName = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameExists = groupByNameExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByName = append(paramsByName, url.PathEscape(attrName.ValueString()))
		if groupByNameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s", paramsByName...))
		}
		return groupByNameExists
	}()
	searchDefined = searchDefined || groupByNameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for Inventory
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Inventory on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Inventory
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Inventory on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = helpers.ExtractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
