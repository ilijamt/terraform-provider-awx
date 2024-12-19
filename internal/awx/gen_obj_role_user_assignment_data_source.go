package awx

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &roleUserAssignmentDataSource{}
	_ datasource.DataSourceWithConfigure = &roleUserAssignmentDataSource{}
)

// NewRoleUserAssignmentDataSource is a helper function to instantiate the RoleUserAssignment data source.
func NewRoleUserAssignmentDataSource() datasource.DataSource {
	return &roleUserAssignmentDataSource{}
}

// roleUserAssignmentDataSource is the data source implementation.
type roleUserAssignmentDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *roleUserAssignmentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/role_user_assignments/"
}

// Metadata returns the data source type name.
func (o *roleUserAssignmentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_role_user_assignment"
}

// Schema defines the schema for the data source.
func (o *roleUserAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"content_type": schema.Int64Attribute{
				Description: "The type of resource this applies to",
				Sensitive:   false,
				Computed:    true,
			},
			"created_by": schema.Int64Attribute{
				Description: "The user who created this resource",
				Sensitive:   false,
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this role user assignment.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(
						path.MatchRoot("id"),
					),
				},
			},
			"object_ansible_id": schema.StringAttribute{
				Description: "Resource id of the object this role applies to. Alternative to the object_id field.",
				Sensitive:   false,
				Computed:    true,
			},
			"object_id": schema.StringAttribute{
				Description: "Primary key of the object this assignment applies to, null value indicates system-wide assignment",
				Sensitive:   false,
				Computed:    true,
			},
			"role_definition": schema.Int64Attribute{
				Description: "The role definition which defines permissions conveyed by this assignment",
				Sensitive:   false,
				Computed:    true,
			},
			"user": schema.Int64Attribute{
				Description: "User",
				Sensitive:   false,
				Computed:    true,
			},
			"user_ansible_id": schema.StringAttribute{
				Description: "Resource id of the user who will receive permissions from this assignment. Alternative to user field.",
				Sensitive:   false,
				Computed:    true,
			},
		},
	}
}

func (o *roleUserAssignmentDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *roleUserAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state roleUserAssignmentTerraformModel
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

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for RoleUserAssignment
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for RoleUserAssignment on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for RoleUserAssignment
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for RoleUserAssignment on %s", endpoint),
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
