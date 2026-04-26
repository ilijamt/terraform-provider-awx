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

type hostDataSource = framework.GenericDataSource[hostTerraformModel, *hostTerraformModel]

// NewHostDataSource is a helper function to instantiate the Host data source.
func NewHostDataSource() datasource.DataSource {
	return &hostDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "host", Endpoint: "/api/v2/hosts/"}},
		Cfg: framework.DataSourceCfg[hostTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this host.",
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Is this host online and available for running jobs?",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this host.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"instance_id": schema.StringAttribute{
						Description: "The value used by the remote inventory source to uniquely identify the host",
						Computed:    true,
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Computed:    true,
					},
					"last_job": schema.Int64Attribute{
						Description: "Last job",
						Computed:    true,
					},
					"last_job_host_summary": schema.Int64Attribute{
						Description: "Last job host summary",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this host.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"variables": schema.StringAttribute{
						Description: "Host variables in JSON or YAML format.",
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
			ResourceName: "Host",
		},
	}
}
