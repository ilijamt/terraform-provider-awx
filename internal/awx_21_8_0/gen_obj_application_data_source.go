package awx_21_8_0

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

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
	_ datasource.DataSource              = &applicationDataSource{}
	_ datasource.DataSourceWithConfigure = &applicationDataSource{}
)

// NewApplicationDataSource is a helper function to instantiate the Application data source.
func NewApplicationDataSource() datasource.DataSource {
	return &applicationDataSource{}
}

// applicationDataSource is the data source implementation.
type applicationDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *applicationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/applications/"
}

// Metadata returns the data source type name.
func (o *applicationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

// Schema defines the schema for the data source.
func (o *applicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"authorization_grant_type": schema.StringAttribute{
				Description: "The Grant type the user must use for acquire tokens for this application.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"authorization-code", "password"}...),
				},
			},
			"client_id": schema.StringAttribute{
				Description: "Client id",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"client_secret": schema.StringAttribute{
				Description: "Used for more stringent verification of access to an application when creating a token.",
				Sensitive:   true,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"client_type": schema.StringAttribute{
				Description: "Set to Public or Confidential depending on how secure the client device is.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"confidential", "public"}...),
				},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this application.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this application.",
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
				Description: "Name of this application.",
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
				Description: "Organization containing this application.",
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
			"redirect_uris": schema.StringAttribute{
				Description: "Allowed URIs list, space separated",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"skip_authorization": schema.BoolAttribute{
				Description: "Set True to skip authorization step for completely trusted applications.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			// Write only elements
		},
	}
}

func (o *applicationDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *applicationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state applicationTerraformModel
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

	// Creates a new request for Application
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Application on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Application
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Application on %s", endpoint),
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
	if err = hookApplication(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on Application",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
