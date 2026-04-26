package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type tokensDataSource = framework.GenericDataSource[tokensTerraformModel, *tokensTerraformModel]

// NewTokensDataSource is a helper function to instantiate the Tokens data source.
func NewTokensDataSource() datasource.DataSource {
	return &tokensDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "token", Endpoint: "/api/v2/tokens/"}},
		Cfg: framework.DataSourceCfg[tokensTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"application": schema.Int64Attribute{
						Description: "Application",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this access token.",
						Computed:    true,
					},
					"expires": schema.StringAttribute{
						Description: "Expires",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this access token.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"refresh_token": schema.StringAttribute{
						Description: "Refresh token",
						Computed:    true,
					},
					"scope": schema.StringAttribute{
						Description: "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].",
						Computed:    true,
					},
					"token": schema.StringAttribute{
						Description: "Token",
						Computed:    true,
					},
					"user": schema.Int64Attribute{
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
