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
	_ datasource.DataSource              = &labelDataSource{}
	_ datasource.DataSourceWithConfigure = &labelDataSource{}
)

// NewLabelDataSource is a helper function to instantiate the Label data source.
func NewLabelDataSource() datasource.DataSource {
	return &labelDataSource{}
}

// labelDataSource is the data source implementation.
type labelDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *labelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/labels/"
}

// Metadata returns the data source type name.
func (o *labelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_label"
}

// Schema defines the schema for the data source.
func (o *labelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"id": schema.Int64Attribute{
				Description: "Database ID for this label.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ConflictsWith(
						path.MatchRoot("name"),
						path.MatchRoot("organization"),
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of this label.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRoot("organization"),
					),
					stringvalidator.ConflictsWith(
						path.MatchRoot("id"),
					),
				},
			},
			"organization": schema.Int64Attribute{
				Description: "Organization this label belongs to.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.AlsoRequires(
						path.MatchRoot("name"),
					),
					int64validator.ConflictsWith(
						path.MatchRoot("id"),
					),
				},
			},
		},
	}
}

func (o *labelDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *labelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state labelTerraformModel
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

	// Evaluate group 'by_name_organization' based on the schema definition
	var groupByNameOrganizationExists = func() bool {
		var groupByNameOrganizationExists = true
		var paramsByNameOrganization = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameOrganizationExists = groupByNameOrganizationExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByNameOrganization = append(paramsByNameOrganization, url.PathEscape(attrName.ValueString()))
		var attrOrganization types.Int64
		req.Config.GetAttribute(ctx, path.Root("organization"), &attrOrganization)
		groupByNameOrganizationExists = groupByNameOrganizationExists && (!attrOrganization.IsNull() && !attrOrganization.IsUnknown())
		paramsByNameOrganization = append(paramsByNameOrganization, attrOrganization.ValueInt64())
		if groupByNameOrganizationExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s&organization=%d", paramsByNameOrganization...))
		}
		return groupByNameOrganizationExists
	}()
	searchDefined = searchDefined || groupByNameOrganizationExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for Label
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Label on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Label
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Label on %s", endpoint),
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
