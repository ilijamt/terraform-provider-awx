package awx

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
	_ datasource.DataSource              = &credentialDataSource{}
	_ datasource.DataSourceWithConfigure = &credentialDataSource{}
)

// NewCredentialDataSource is a helper function to instantiate the Credential data source.
func NewCredentialDataSource() datasource.DataSource {
	return &credentialDataSource{}
}

// credentialDataSource is the data source implementation.
type credentialDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *credentialDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credentials/"
}

// Metadata returns the data source type name.
func (o *credentialDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credential"
}

// Schema defines the schema for the data source.
func (o *credentialDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"cloud": schema.BoolAttribute{
				Description: "Cloud",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"credential_type": schema.Int64Attribute{
				Description: "Specify the type of credential you want to create. Refer to the documentation for details on each type.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this credential.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this credential.",
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
			"inputs": schema.StringAttribute{
				Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"kind": schema.StringAttribute{
				Description: "Kind",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"kubernetes": schema.BoolAttribute{
				Description: "Kubernetes",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"managed": schema.BoolAttribute{
				Description: "Managed",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this credential.",
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
			"organization": schema.Int64Attribute{
				Description: "Inherit permissions from organization roles. If provided on creation, do not give either user or team.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			// Write only elements
			"user": schema.Int64Attribute{
				Description: "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation.",
				Computed:    true,
				Sensitive:   false,
			},
			"team": schema.Int64Attribute{
				Description: "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation.",
				Computed:    true,
				Sensitive:   false,
			},
		},
	}
}

func (o *credentialDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *credentialDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state credentialTerraformModel
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
			endpoint = p.Clean(fmt.Sprintf("%s//?name__exact=%s", paramsByName...))
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

	// Creates a new request for Credential
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Credential on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Credential on %s", endpoint),
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
	if err = hookCredential(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on Credential",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
