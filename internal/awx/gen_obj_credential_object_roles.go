package awx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/mitchellh/mapstructure"
)

var (
	_ datasource.DataSource              = &credentialObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &credentialObjectRolesDataSource{}
)

// NewCredentialObjectRolesDataSource is a helper function to instantiate the Credential data source.
func NewCredentialObjectRolesDataSource() datasource.DataSource {
	return &credentialObjectRolesDataSource{}
}

// credentialObjectRolesDataSource is the data source implementation.
type credentialObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *credentialObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credentials/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *credentialObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credential_object_roles"
}

// Schema defines the schema for the data source.
func (o *credentialObjectRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Credential ID",
				Required:    true,
			},
			"roles": schema.MapAttribute{
				Description: "Roles for credential",
				ElementType: types.Int64Type,
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *credentialObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state credentialObjectRolesModel
	var err error
	var id types.Int64

	if d := req.Config.GetAttribute(ctx, path.Root("id"), &id); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	// Creates a new request for Credential
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(o.endpoint, id.ValueInt64()), nil); err != nil {
		resp.Diagnostics.AddError(
			"Unable to create a new request for credential",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch the request for credential object roles",
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for credential",
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
