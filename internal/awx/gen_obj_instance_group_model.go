package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
