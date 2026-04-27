package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type instanceGroupTerraformModel struct {
	Capacity                 types.Int64   `tfsdk:"capacity" json:"capacity"`
	ConsumedCapacity         types.Float64 `tfsdk:"consumed_capacity" json:"consumed_capacity"`
	Credential               types.Int64   `tfsdk:"credential" json:"credential"`
	ID                       types.Int64   `tfsdk:"id" json:"id"`
	Instances                types.Int64   `tfsdk:"instances" json:"instances"`
	IsContainerGroup         types.Bool    `tfsdk:"is_container_group" json:"is_container_group"`
	JobsRunning              types.Int64   `tfsdk:"jobs_running" json:"jobs_running"`
	JobsTotal                types.Int64   `tfsdk:"jobs_total" json:"jobs_total"`
	MaxConcurrentJobs        types.Int64   `tfsdk:"max_concurrent_jobs" json:"max_concurrent_jobs"`
	MaxForks                 types.Int64   `tfsdk:"max_forks" json:"max_forks"`
	Name                     types.String  `tfsdk:"name" json:"name"`
	PercentCapacityRemaining types.Float64 `tfsdk:"percent_capacity_remaining" json:"percent_capacity_remaining"`
	PodSpecOverride          types.String  `tfsdk:"pod_spec_override" json:"pod_spec_override"`
	PolicyInstanceList       types.String  `tfsdk:"policy_instance_list" json:"policy_instance_list"`
	PolicyInstanceMinimum    types.Int64   `tfsdk:"policy_instance_minimum" json:"policy_instance_minimum"`
	PolicyInstancePercentage types.Int64   `tfsdk:"policy_instance_percentage" json:"policy_instance_percentage"`
}

func (o *instanceGroupTerraformModel) Clone() instanceGroupTerraformModel {
	return *o
}

func (o *instanceGroupTerraformModel) BodyRequest() *instanceGroupBodyRequestModel {
	var req instanceGroupBodyRequestModel
	req.Credential = o.Credential.ValueInt64()
	req.IsContainerGroup = o.IsContainerGroup.ValueBool()
	req.MaxConcurrentJobs = o.MaxConcurrentJobs.ValueInt64()
	req.MaxForks = o.MaxForks.ValueInt64()
	req.Name = o.Name.ValueString()
	req.PodSpecOverride = o.PodSpecOverride.ValueString()
	req.PolicyInstanceList = json.RawMessage(o.PolicyInstanceList.ValueString())
	req.PolicyInstanceMinimum = o.PolicyInstanceMinimum.ValueInt64()
	req.PolicyInstancePercentage = o.PolicyInstancePercentage.ValueInt64()
	return &req
}

func (o *instanceGroupTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Capacity, data["capacity"]))
	collect(helpers.AttrValueSetFloat64(&o.ConsumedCapacity, data["consumed_capacity"]))
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Instances, data["instances"]))
	collect(helpers.AttrValueSetBool(&o.IsContainerGroup, data["is_container_group"]))
	collect(helpers.AttrValueSetInt64(&o.JobsRunning, data["jobs_running"]))
	collect(helpers.AttrValueSetInt64(&o.JobsTotal, data["jobs_total"]))
	collect(helpers.AttrValueSetInt64(&o.MaxConcurrentJobs, data["max_concurrent_jobs"]))
	collect(helpers.AttrValueSetInt64(&o.MaxForks, data["max_forks"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetFloat64(&o.PercentCapacityRemaining, data["percent_capacity_remaining"]))
	collect(helpers.AttrValueSetString(&o.PodSpecOverride, data["pod_spec_override"], false))
	collect(helpers.AttrValueSetJsonString(&o.PolicyInstanceList, data["policy_instance_list"], false))
	collect(helpers.AttrValueSetInt64(&o.PolicyInstanceMinimum, data["policy_instance_minimum"]))
	collect(helpers.AttrValueSetInt64(&o.PolicyInstancePercentage, data["policy_instance_percentage"]))
	return diags, nil
}

type instanceGroupBodyRequestModel struct {
	Credential               int64           `json:"credential,omitempty"`
	IsContainerGroup         bool            `json:"is_container_group"`
	MaxConcurrentJobs        int64           `json:"max_concurrent_jobs,omitempty"`
	MaxForks                 int64           `json:"max_forks,omitempty"`
	Name                     string          `json:"name"`
	PodSpecOverride          string          `json:"pod_spec_override,omitempty"`
	PolicyInstanceList       json.RawMessage `json:"policy_instance_list,omitempty"`
	PolicyInstanceMinimum    int64           `json:"policy_instance_minimum,omitempty"`
	PolicyInstancePercentage int64           `json:"policy_instance_percentage,omitempty"`
}

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

type instanceGroupDataSource = framework.GenericDataSource[instanceGroupTerraformModel, *instanceGroupTerraformModel]

// NewInstanceGroupDataSource is a helper function to instantiate the InstanceGroup data source.
func NewInstanceGroupDataSource() datasource.DataSource {
	return &instanceGroupDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "instance_group", Endpoint: "/api/v2/instance_groups/"}},
		Cfg: framework.DataSourceCfg[instanceGroupTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"capacity": dschema.Int64Attribute{
						Description: "Capacity",
						Computed:    true,
					},
					"consumed_capacity": dschema.Float64Attribute{
						Description: "Consumed capacity",
						Computed:    true,
					},
					"credential": dschema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
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
					"instances": dschema.Int64Attribute{
						Description: "Instances",
						Computed:    true,
					},
					"is_container_group": dschema.BoolAttribute{
						Description: "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.",
						Computed:    true,
					},
					"jobs_running": dschema.Int64Attribute{
						Description: "Jobs running",
						Computed:    true,
					},
					"jobs_total": dschema.Int64Attribute{
						Description: "Count of all jobs that target this instance group",
						Computed:    true,
					},
					"max_concurrent_jobs": dschema.Int64Attribute{
						Description: "Maximum number of concurrent jobs to run on a group. When set to zero, no maximum is enforced.",
						Computed:    true,
					},
					"max_forks": dschema.Int64Attribute{
						Description: "Maximum number of forks to execute concurrently on a group. When set to zero, no maximum is enforced.",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
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
					"percent_capacity_remaining": dschema.Float64Attribute{
						Description: "Percent capacity remaining",
						Computed:    true,
					},
					"pod_spec_override": dschema.StringAttribute{
						Description: "Pod spec override",
						Computed:    true,
					},
					"policy_instance_list": dschema.StringAttribute{
						Description: "List of exact-match Instances that will be assigned to this group",
						Computed:    true,
					},
					"policy_instance_minimum": dschema.Int64Attribute{
						Description: "Static minimum number of Instances that will be automatically assign to this group when new instances come online.",
						Computed:    true,
					},
					"policy_instance_percentage": dschema.Int64Attribute{
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
