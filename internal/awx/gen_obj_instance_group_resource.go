package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type instanceGroupResource = framework.GenericResource[instanceGroupTerraformModel, instanceGroupBodyRequestModel, *instanceGroupTerraformModel]

// NewInstanceGroupResource is a helper function to simplify the provider implementation.
func NewInstanceGroupResource() resource.Resource {
	return &instanceGroupResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "instance_group", Endpoint: "/api/v2/instance_groups/"}},
		Cfg: framework.ResourceCfg[instanceGroupTerraformModel, instanceGroupBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"is_container_group": schema.BoolAttribute{
						Description: "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"max_concurrent_jobs": schema.Int64Attribute{
						Description: "Maximum number of concurrent jobs to run on a group. When set to zero, no maximum is enforced.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"max_forks": schema.Int64Attribute{
						Description: "Maximum number of forks to execute concurrently on a group. When set to zero, no maximum is enforced.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this instance group.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(250),
						},
					},
					"pod_spec_override": schema.StringAttribute{
						Description: "Pod spec override",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"apiVersion":"v1","kind":"Pod","metadata":{"namespace":"default"},"spec":{"automountServiceAccountToken":false,"containers":[{"args":["ansible-runner","worker","--private-data-dir=/runner"],"image":"quay.io/ansible/awx-ee:latest","name":"worker","resources":{"requests":{"cpu":"250m","memory":"100Mi"}}}],"serviceAccountName":"default"}}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"policy_instance_list": schema.StringAttribute{
						Description: "List of exact-match Instances that will be assigned to this group",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"policy_instance_minimum": schema.Int64Attribute{
						Description: "Static minimum number of Instances that will be automatically assign to this group when new instances come online.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"policy_instance_percentage": schema.Int64Attribute{
						Description: "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"capacity": schema.Int64Attribute{
						Description: "Capacity",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"consumed_capacity": schema.Float64Attribute{
						Description: "Consumed capacity",
						Computed:    true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this instance group.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"instances": schema.Int64Attribute{
						Description: "Instances",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"jobs_running": schema.Int64Attribute{
						Description: "Jobs running",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"jobs_total": schema.Int64Attribute{
						Description: "Count of all jobs that target this instance group",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"percent_capacity_remaining": schema.Float64Attribute{
						Description: "Percent capacity remaining",
						Computed:    true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *instanceGroupTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "InstanceGroup",
		},
	}
}
