package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

var (
	_ datasource.DataSource              = &meDataSource{}
	_ datasource.DataSourceWithConfigure = &meDataSource{}
)

// NewMeDataSource is a helper function to instantiate the Me data source.
func NewMeDataSource() datasource.DataSource {
	return &meDataSource{}
}

// meDataSource is the data source implementation.
type meDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *meDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/me/"
}

// Metadata returns the data source type name.
func (o *meDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_me"
}

// Schema defines the schema for the data source.
func (o *meDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"email": schema.StringAttribute{
				Description: "Email address",
				Computed:    true,
			},
			"external_account": schema.StringAttribute{
				Description: "Set if the account is managed by an external service",
				Computed:    true,
			},
			"first_name": schema.StringAttribute{
				Description: "First name",
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this user.",
				Computed:    true,
			},
			"is_superuser": schema.BoolAttribute{
				Description: "Designates that this user has all permissions without explicitly assigning them.",
				Computed:    true,
			},
			"is_system_auditor": schema.BoolAttribute{
				Description: "Is system auditor",
				Computed:    true,
			},
			"last_login": schema.StringAttribute{
				Description: "Last login",
				Computed:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "Last name",
				Computed:    true,
			},
			"ldap_dn": schema.StringAttribute{
				Description: "Ldap dn",
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *meDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *meDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state meTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for Me
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Me on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Me
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Me on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

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
