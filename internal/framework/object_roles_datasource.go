package framework

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mitchellh/mapstructure"

	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

var (
	_ datasource.DataSource              = (*ObjectRolesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ObjectRolesDataSource)(nil)
)

// ObjectRolesModel is the shared state model for every <resource>_object_roles
// data source. It is the same shape across all resources: an input ID and a
// computed map of role name → role ID.
type ObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}

// ObjectRolesDataSource is the generic implementation backing every
// <resource>_object_roles data source. The per-resource constructors in
// internal/awx are thin wrappers that supply TypeName, Endpoint, DisplayName,
// and the deprecation flag.
type ObjectRolesDataSource struct {
	DataSourceBase
	DisplayName string
	Deprecated  bool
}

// NewObjectRolesDataSource constructs an ObjectRolesDataSource with the given
// configuration. typeName is the full Terraform type-name suffix
// (e.g. "instance_group_object_roles"), endpoint is the API path with a single
// %d placeholder for the parent ID, and displayName is the human-readable
// resource name used in attribute descriptions and error messages
// (e.g. "InstanceGroup").
func NewObjectRolesDataSource(typeName, endpoint, displayName string, deprecated bool) datasource.DataSource {
	return &ObjectRolesDataSource{
		DataSourceBase: DataSourceBase{
			ProviderBase: ProviderBase{TypeName: typeName, Endpoint: endpoint},
		},
		DisplayName: displayName,
		Deprecated:  deprecated,
	}
}

// Schema defines the schema for the data source.
func (o *ObjectRolesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: fmt.Sprintf("%s ID", o.DisplayName),
				Required:    true,
			},
			"roles": schema.MapAttribute{
				Description: fmt.Sprintf("Roles for %s", strings.TrimSuffix(o.TypeName, "_object_roles")),
				ElementType: types.Int64Type,
				Computed:    true,
			},
		},
	}
	if o.Deprecated {
		s.DeprecationMessage = "This data source has been deprecated and will be removed in a future release."
	}
	resp.Schema = s
}

// Read refreshes the Terraform state with the latest data.
func (o *ObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ObjectRolesModel
	if DiagnosticsHasError(&resp.Diagnostics, req.Config.GetAttribute(ctx, path.Root("id"), &state.ID)...) {
		return
	}

	endpoint := fmt.Sprintf(o.Endpoint, state.ID.ValueInt64())
	data, d := ReadRequest(ctx, o.Client, endpoint, fmt.Sprintf("%s/ObjectRoles", o.DisplayName))
	if DiagnosticsHasError(&resp.Diagnostics, d...) {
		return
	}

	var sr models.SearchResultObjectRole
	if err := mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to decode the search result data for %s", strings.ToLower(o.DisplayName)),
			err.Error(),
		)
		return
	}

	in := make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var dg diag.Diagnostics
	if state.Roles, dg = types.MapValue(types.Int64Type, in); DiagnosticsHasError(&resp.Diagnostics, dg...) {
		return
	}

	if DiagnosticsHasError(&resp.Diagnostics, resp.State.Set(ctx, &state)...) {
		return
	}
}
