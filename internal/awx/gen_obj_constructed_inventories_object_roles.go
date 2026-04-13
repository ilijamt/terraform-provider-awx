package awx

import (
	"context"
	"fmt"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mitchellh/mapstructure"
)

var (
	_ datasource.DataSource              = &constructedInventoriesObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &constructedInventoriesObjectRolesDataSource{}
)

// NewConstructedInventoriesObjectRolesDataSource is a helper function to instantiate the ConstructedInventories data source.
func NewConstructedInventoriesObjectRolesDataSource() datasource.DataSource {
	return &constructedInventoriesObjectRolesDataSource{DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "constructed_inventories_object_roles", Endpoint: "/api/v2/constructed_inventories/%d/object_roles/"}}}
}

// constructedInventoriesObjectRolesDataSource is the data source implementation.
type constructedInventoriesObjectRolesDataSource struct {
	framework.DataSourceBase
}

// Schema defines the schema for the data source.
func (o *constructedInventoriesObjectRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		DeprecationMessage: "This data source has been deprecated and will be removed in a future release.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "ConstructedInventories ID",
				Required:    true,
			},
			"roles": schema.MapAttribute{
				Description: "Roles for constructed_inventories",
				ElementType: types.Int64Type,
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *constructedInventoriesObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state constructedInventoriesObjectRolesModel
	var err error
	var id types.Int64

	if framework.DiagnosticsHasError(&resp.Diagnostics, req.Config.GetAttribute(ctx, path.Root("id"), &id)...) {
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	var endpoint = fmt.Sprintf(o.Endpoint, id.ValueInt64())
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "ConstructedInventories/ObjectRoles")
	if framework.DiagnosticsHasError(&resp.Diagnostics, d...) {
		return
	}

	var sr models.SearchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for constructedinventories",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var dg diag.Diagnostics
	if state.Roles, dg = types.MapValue(types.Int64Type, in); framework.DiagnosticsHasError(&resp.Diagnostics, dg...) {
		return
	}

	if framework.DiagnosticsHasError(&resp.Diagnostics, resp.State.Set(ctx, &state)...) {
		return
	}
}
