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
	_ datasource.DataSource              = &credentialTypeDataSource{}
	_ datasource.DataSourceWithConfigure = &credentialTypeDataSource{}
)

// NewCredentialTypeDataSource is a helper function to instantiate the CredentialType data source.
func NewCredentialTypeDataSource() datasource.DataSource {
	return &credentialTypeDataSource{}
}

// credentialTypeDataSource is the data source implementation.
type credentialTypeDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *credentialTypeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credential_types/"
}

// Metadata returns the data source type name.
func (o *credentialTypeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credential_type"
}

// Schema defines the schema for the data source.
func (o *credentialTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"description": schema.StringAttribute{
				Description: "Optional description of this credential type.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this credential type.",
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
			"injectors": schema.StringAttribute{
				Description: "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"inputs": schema.StringAttribute{
				Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"kind": schema.StringAttribute{
				Description: "The credential type",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ssh",
						"vault",
						"net",
						"scm",
						"cloud",
						"registry",
						"token",
						"insights",
						"external",
						"kubernetes",
						"galaxy",
						"cryptography",
					),
				},
			},
			"managed": schema.BoolAttribute{
				Description: "Is the resource managed",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this credential type.",
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
			"namespace": schema.StringAttribute{
				Description: "The namespace to which the resource belongs to",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
		},
	}
}

func (o *credentialTypeDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *credentialTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state credentialTypeTerraformModel
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

	// Creates a new request for CredentialType
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for CredentialType on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for CredentialType
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for CredentialType on %s", endpoint),
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
