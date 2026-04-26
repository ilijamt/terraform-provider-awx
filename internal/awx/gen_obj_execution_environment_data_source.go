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

type executionEnvironmentDataSource = framework.GenericDataSource[executionEnvironmentTerraformModel, *executionEnvironmentTerraformModel]

// NewExecutionEnvironmentDataSource is a helper function to instantiate the ExecutionEnvironment data source.
func NewExecutionEnvironmentDataSource() datasource.DataSource {
	return &executionEnvironmentDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "execution_environment", Endpoint: "/api/v2/execution_environments/"}},
		Cfg: framework.DataSourceCfg[executionEnvironmentTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this execution environment.",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this execution environment.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"image": schema.StringAttribute{
						Description: "The full image location, including the container registry, image name, and version tag.",
						Computed:    true,
					},
					"managed": schema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this execution environment.",
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
						Description: "The organization used to determine access to this execution environment.",
						Computed:    true,
					},
					"pull": schema.StringAttribute{
						Description: "Pull image before running?",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "ExecutionEnvironment",
		},
	}
}
