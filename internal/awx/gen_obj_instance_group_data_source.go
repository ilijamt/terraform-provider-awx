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

type instanceGroupDataSource = framework.GenericDataSource[instanceGroupTerraformModel, *instanceGroupTerraformModel]

// NewInstanceGroupDataSource is a helper function to instantiate the InstanceGroup data source.
func NewInstanceGroupDataSource() datasource.DataSource {
	return &instanceGroupDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "instance_group", Endpoint: "/api/v2/instance_groups/"}},
		Cfg: framework.DataSourceCfg[instanceGroupTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"capacity": schema.Int64Attribute{
						Description: "Capacity",
						Computed:    true,
					},
					"consumed_capacity": schema.Float64Attribute{
						Description: "Consumed capacity",
						Computed:    true,
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this instance group.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"instances": schema.Int64Attribute{
						Description: "Instances",
						Computed:    true,
					},
					"is_container_group": schema.BoolAttribute{
						Description: "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.",
						Computed:    true,
					},
					"jobs_running": schema.Int64Attribute{
						Description: "Jobs running",
						Computed:    true,
					},
					"jobs_total": schema.Int64Attribute{
						Description: "Count of all jobs that target this instance group",
						Computed:    true,
					},
					"max_concurrent_jobs": schema.Int64Attribute{
						Description: "Maximum number of concurrent jobs to run on a group. When set to zero, no maximum is enforced.",
						Computed:    true,
					},
					"max_forks": schema.Int64Attribute{
						Description: "Maximum number of forks to execute concurrently on a group. When set to zero, no maximum is enforced.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this instance group.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"percent_capacity_remaining": schema.Float64Attribute{
						Description: "Percent capacity remaining",
						Computed:    true,
					},
					"pod_spec_override": schema.StringAttribute{
						Description: "Pod spec override",
						Computed:    true,
					},
					"policy_instance_list": schema.StringAttribute{
						Description: "List of exact-match Instances that will be assigned to this group",
						Computed:    true,
					},
					"policy_instance_minimum": schema.Int64Attribute{
						Description: "Static minimum number of Instances that will be automatically assign to this group when new instances come online.",
						Computed:    true,
					},
					"policy_instance_percentage": schema.Int64Attribute{
						Description: "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.",
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
			ResourceName: "InstanceGroup",
		},
	}
}
