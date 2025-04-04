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
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

// NewUserDataSource is a helper function to instantiate the User data source.
func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

// userDataSource is the data source implementation.
type userDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/users/"
}

// Metadata returns the data source type name.
func (o *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the data source.
func (o *userDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"email": schema.StringAttribute{
				Description: "Email address",
				Sensitive:   false,
				Computed:    true,
			},
			"external_account": schema.StringAttribute{
				Description: "Set if the account is managed by an external service",
				Sensitive:   false,
				Computed:    true,
			},
			"first_name": schema.StringAttribute{
				Description: "First name",
				Sensitive:   false,
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this user.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("username"),
					),
				},
			},
			"is_superuser": schema.BoolAttribute{
				Description: "Designates that this user has all permissions without explicitly assigning them.",
				Sensitive:   false,
				Computed:    true,
			},
			"is_system_auditor": schema.BoolAttribute{
				Description: "Is system auditor",
				Sensitive:   false,
				Computed:    true,
			},
			"last_login": schema.StringAttribute{
				Description: "Last login",
				Sensitive:   false,
				Computed:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "Last name",
				Sensitive:   false,
				Computed:    true,
			},
			"ldap_dn": schema.StringAttribute{
				Description: "Ldap dn",
				Sensitive:   false,
				Computed:    true,
			},
			"password": schema.StringAttribute{
				Description: "Field used to change the password.",
				Sensitive:   true,
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("username"),
					),
				},
			},
		},
	}
}

func (o *userDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state userTerraformModel
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

	// Evaluate group 'by_username' based on the schema definition
	var groupByUsernameExists = func() bool {
		var groupByUsernameExists = true
		var paramsByUsername = []any{o.endpoint}
		var attrUsername types.String
		req.Config.GetAttribute(ctx, path.Root("username"), &attrUsername)
		groupByUsernameExists = groupByUsernameExists && (!attrUsername.IsNull() && !attrUsername.IsUnknown())
		paramsByUsername = append(paramsByUsername, url.PathEscape(attrUsername.ValueString()))
		if groupByUsernameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?username__exact=%s", paramsByUsername...))
		}
		return groupByUsernameExists
	}()
	searchDefined = searchDefined || groupByUsernameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for User
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for User
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for User on %s", endpoint),
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
	if err = hookUser(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on User",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
