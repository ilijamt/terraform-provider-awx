package awx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ilijamt/terraform-provider-awx/internal/models"

	"github.com/go-viper/mapstructure/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
)

var (
	_ datasource.DataSource              = &inventoryObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &inventoryObjectRolesDataSource{}
)

// NewInventoryObjectRolesDataSource is a helper function to instantiate the Inventory data source.
func NewInventoryObjectRolesDataSource() datasource.DataSource {
	return &inventoryObjectRolesDataSource{}
}

// inventoryObjectRolesDataSource is the data source implementation.
type inventoryObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *inventoryObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventories/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *inventoryObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_object_roles"
}

// Schema defines the schema for the data source.
func (o *inventoryObjectRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		DeprecationMessage: "This data source has been deprecated and will be removed in a future release.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Inventory ID",
				Required:    true,
			},
			"roles": schema.MapAttribute{
				Description: "Roles for inventory",
				ElementType: types.Int64Type,
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *inventoryObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state inventoryObjectRolesModel
	var err error
	var id types.Int64

	if d := req.Config.GetAttribute(ctx, path.Root("id"), &id); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	// Creates a new request for Credential
	var r *http.Request
	var endpoint = fmt.Sprintf(o.endpoint, id.ValueInt64())
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			"Unable to create a new request for inventory",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to fetch the request for inventory object roles on %s", endpoint),
			err.Error(),
		)
		return
	}

	var sr models.SearchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for inventory",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var d diag.Diagnostics
	if state.Roles, d = types.MapValue(types.Int64Type, in); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
