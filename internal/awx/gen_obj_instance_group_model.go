package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// instanceGroupTerraformModel maps the schema for InstanceGroup when using Data Source
type instanceGroupTerraformModel struct {
	// Capacity ""
	Capacity types.Int64 `tfsdk:"capacity" json:"capacity"`
	// ConsumedCapacity ""
	ConsumedCapacity types.Float64 `tfsdk:"consumed_capacity" json:"consumed_capacity"`
	// Credential ""
	Credential types.Int64 `tfsdk:"credential" json:"credential"`
	// ID "Database ID for this instance group."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Instances ""
	Instances types.Int64 `tfsdk:"instances" json:"instances"`
	// IsContainerGroup "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster."
	IsContainerGroup types.Bool `tfsdk:"is_container_group" json:"is_container_group"`
	// JobsRunning ""
	JobsRunning types.Int64 `tfsdk:"jobs_running" json:"jobs_running"`
	// JobsTotal "Count of all jobs that target this instance group"
	JobsTotal types.Int64 `tfsdk:"jobs_total" json:"jobs_total"`
	// MaxConcurrentJobs "Maximum number of concurrent jobs to run on a group. When set to zero, no maximum is enforced."
	MaxConcurrentJobs types.Int64 `tfsdk:"max_concurrent_jobs" json:"max_concurrent_jobs"`
	// MaxForks "Maximum number of forks to execute concurrently on a group. When set to zero, no maximum is enforced."
	MaxForks types.Int64 `tfsdk:"max_forks" json:"max_forks"`
	// Name "Name of this instance group."
	Name types.String `tfsdk:"name" json:"name"`
	// PercentCapacityRemaining ""
	PercentCapacityRemaining types.Float64 `tfsdk:"percent_capacity_remaining" json:"percent_capacity_remaining"`
	// PodSpecOverride ""
	PodSpecOverride types.String `tfsdk:"pod_spec_override" json:"pod_spec_override"`
	// PolicyInstanceList "List of exact-match Instances that will be assigned to this group"
	PolicyInstanceList types.String `tfsdk:"policy_instance_list" json:"policy_instance_list"`
	// PolicyInstanceMinimum "Static minimum number of Instances that will be automatically assign to this group when new instances come online."
	PolicyInstanceMinimum types.Int64 `tfsdk:"policy_instance_minimum" json:"policy_instance_minimum"`
	// PolicyInstancePercentage "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online."
	PolicyInstancePercentage types.Int64 `tfsdk:"policy_instance_percentage" json:"policy_instance_percentage"`
}

// Clone the object
func (o *instanceGroupTerraformModel) Clone() instanceGroupTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for InstanceGroup
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
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Capacity, data["capacity"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetFloat64(&o.ConsumedCapacity, data["consumed_capacity"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Credential, data["credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Instances, data["instances"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.IsContainerGroup, data["is_container_group"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.JobsRunning, data["jobs_running"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.JobsTotal, data["jobs_total"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.MaxConcurrentJobs, data["max_concurrent_jobs"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.MaxForks, data["max_forks"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetFloat64(&o.PercentCapacityRemaining, data["percent_capacity_remaining"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.PodSpecOverride, data["pod_spec_override"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.PolicyInstanceList, data["policy_instance_list"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.PolicyInstanceMinimum, data["policy_instance_minimum"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.PolicyInstancePercentage, data["policy_instance_percentage"])
		diags.Append(dg...)
	}
	return diags, nil
}

// instanceGroupBodyRequestModel maps the schema for InstanceGroup for creating and updating the data
type instanceGroupBodyRequestModel struct {
	// Credential ""
	Credential int64 `json:"credential,omitempty"`
	// IsContainerGroup "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster."
	IsContainerGroup bool `json:"is_container_group"`
	// MaxConcurrentJobs "Maximum number of concurrent jobs to run on a group. When set to zero, no maximum is enforced."
	MaxConcurrentJobs int64 `json:"max_concurrent_jobs,omitempty"`
	// MaxForks "Maximum number of forks to execute concurrently on a group. When set to zero, no maximum is enforced."
	MaxForks int64 `json:"max_forks,omitempty"`
	// Name "Name of this instance group."
	Name string `json:"name"`
	// PodSpecOverride ""
	PodSpecOverride string `json:"pod_spec_override,omitempty"`
	// PolicyInstanceList "List of exact-match Instances that will be assigned to this group"
	PolicyInstanceList json.RawMessage `json:"policy_instance_list,omitempty"`
	// PolicyInstanceMinimum "Static minimum number of Instances that will be automatically assign to this group when new instances come online."
	PolicyInstanceMinimum int64 `json:"policy_instance_minimum,omitempty"`
	// PolicyInstancePercentage "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online."
	PolicyInstancePercentage int64 `json:"policy_instance_percentage,omitempty"`
}
