package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type groupDataSource = framework.GenericDataSource[groupTerraformModel, *groupTerraformModel]

// NewGroupDataSource is a helper function to instantiate the Group data source.
func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "group", Endpoint: "/api/v2/groups/"}},
		Cfg: framework.DataSourceCfg[groupTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"description": schema.StringAttribute{
						Description: "Optional description of this group.",
						Sensitive:   false,
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this group.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Sensitive:   false,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this group.",
						Sensitive:   false,
						Computed:    true,
					},
					"variables": schema.StringAttribute{
						Description: "Group variables in JSON or YAML format.",
						Sensitive:   false,
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
			ResourceName: "Group",
		},
	}
}
