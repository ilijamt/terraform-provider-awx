package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type tokensTerraformModel struct {
	Application  types.Int64  `tfsdk:"application" json:"application"`
	Description  types.String `tfsdk:"description" json:"description"`
	Expires      types.String `tfsdk:"expires" json:"expires"`
	ID           types.Int64  `tfsdk:"id" json:"id"`
	RefreshToken types.String `tfsdk:"refresh_token" json:"refresh_token"`
	Scope        types.String `tfsdk:"scope" json:"scope"`
	Token        types.String `tfsdk:"token" json:"token"`
	User         types.Int64  `tfsdk:"user" json:"user"`
}

func (o *tokensTerraformModel) Clone() tokensTerraformModel {
	return *o
}

func (o *tokensTerraformModel) BodyRequest() *tokensBodyRequestModel {
	var req tokensBodyRequestModel
	req.Application = o.Application.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Scope = o.Scope.ValueString()
	return &req
}

func (o *tokensTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Application, data["application"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetString(&o.Expires, data["expires"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.RefreshToken, data["refresh_token"], false))
	collect(helpers.AttrValueSetString(&o.Scope, data["scope"], false))
	collect(helpers.AttrValueSetString(&o.Token, data["token"], false))
	collect(helpers.AttrValueSetInt64(&o.User, data["user"]))
	return diags, nil
}

type tokensBodyRequestModel struct {
	Application int64  `json:"application,omitempty"`
	Description string `json:"description,omitempty"`
	Scope       string `json:"scope,omitempty"`
}

type tokensResource = framework.GenericResource[tokensTerraformModel, tokensBodyRequestModel, *tokensTerraformModel]

// NewTokensResource is a helper function to simplify the provider implementation.
func NewTokensResource() resource.Resource {
	return &tokensResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "token", Endpoint: "/api/v2/tokens/"}},
		Cfg: framework.ResourceCfg[tokensTerraformModel, tokensBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"application": schema.Int64Attribute{
						Description: "Application",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this access token.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"scope": schema.StringAttribute{
						Description: "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`write`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"expires": schema.StringAttribute{
						Description: "Expires",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this access token.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"refresh_token": schema.StringAttribute{
						Description: "Refresh token",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"token": schema.StringAttribute{
						Description: "Token",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"user": schema.Int64Attribute{
						Description: "The user representing the token owner",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *tokensTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Tokens",
		},
	}
}

type tokensDataSource = framework.GenericDataSource[tokensTerraformModel, *tokensTerraformModel]

// NewTokensDataSource is a helper function to instantiate the Tokens data source.
func NewTokensDataSource() datasource.DataSource {
	return &tokensDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "token", Endpoint: "/api/v2/tokens/"}},
		Cfg: framework.DataSourceCfg[tokensTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"application": dschema.Int64Attribute{
						Description: "Application",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this access token.",
						Computed:    true,
					},
					"expires": dschema.StringAttribute{
						Description: "Expires",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this access token.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"refresh_token": dschema.StringAttribute{
						Description: "Refresh token",
						Computed:    true,
					},
					"scope": dschema.StringAttribute{
						Description: "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].",
						Computed:    true,
					},
					"token": dschema.StringAttribute{
						Description: "Token",
						Computed:    true,
					},
					"user": dschema.Int64Attribute{
						Description: "The user representing the token owner",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Tokens",
		},
	}
}
