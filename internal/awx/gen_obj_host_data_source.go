package awx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &hostDataSource{}
	_ datasource.DataSourceWithConfigure = &hostDataSource{}
)

// NewHostDataSource is a helper function to instantiate the Host data source.
func NewHostDataSource() datasource.DataSource {
	return &hostDataSource{}
}

// hostDataSource is the data source implementation.
type hostDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *hostDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/hosts/"
}

// Metadata returns the data source type name.
func (o *hostDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host"
}

// Schema defines the schema for the data source.
func (o *hostDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"description": schema.StringAttribute{
				Description: "Optional description of this host.",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Is this host online and available for running jobs?",
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this host.",
				Optional:    true,
				Computed:    true,
			},
			"instance_id": schema.StringAttribute{
				Description: "The value used by the remote inventory source to uniquely identify the host",
				Computed:    true,
			},
			"inventory": schema.Int64Attribute{
				Description: "Inventory",
				Computed:    true,
			},
			"last_job": schema.Int64Attribute{
				Description: "Last job",
				Computed:    true,
			},
			"last_job_host_summary": schema.Int64Attribute{
				Description: "Last job host summary",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of this host.",
				Optional:    true,
				Computed:    true,
			},
			"variables": schema.StringAttribute{
				Description: "Host variables in JSON or YAML format.",
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *hostDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *hostDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state hostTerraformModel
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

	// Creates a new request for Host
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Host
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Host on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = extractDataIfSearchResult(data); err != nil {
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
