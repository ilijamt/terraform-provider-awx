package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type applicationDataSource = framework.GenericDataSource[applicationTerraformModel, *applicationTerraformModel]

// NewApplicationDataSource is a helper function to instantiate the Application data source.
func NewApplicationDataSource() datasource.DataSource {
	return &applicationDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "application", Endpoint: "/api/v2/applications/"}},
		Cfg: framework.DataSourceCfg[applicationTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"authorization_grant_type": schema.StringAttribute{
						Description: "The Grant type the user must use for acquire tokens for this application.",
						Sensitive:   false,
						Computed:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "Client id",
						Sensitive:   false,
						Computed:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "Used for more stringent verification of access to an application when creating a token.",
						Sensitive:   true,
						Computed:    true,
					},
					"client_type": schema.StringAttribute{
						Description: "Set to Public or Confidential depending on how secure the client device is.",
						Sensitive:   false,
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this application.",
						Sensitive:   false,
						Computed:    true,
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
					},
					"skip_authorization": schema.BoolAttribute{
						Description: "Set True to skip authorization step for completely trusted applications.",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name_organization", URLSuffix: "?name__exact=%s&organization=%d", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
					{Name: "organization", Type: "int64", URLEscape: false},
				}},
			},
			Hook:         hookApplication,
			ApiVersion:   ApiVersion,
			ResourceName: "Application",
		},
	}
}
