package awx_21_8_0

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
	// JobsRunning "Count of jobs in the running or waiting state that are targeted for this instance group"
	JobsRunning types.Int64 `tfsdk:"jobs_running" json:"jobs_running"`
	// JobsTotal "Count of all jobs that target this instance group"
	JobsTotal types.Int64 `tfsdk:"jobs_total" json:"jobs_total"`
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
	return instanceGroupTerraformModel{
		Capacity:                 o.Capacity,
		ConsumedCapacity:         o.ConsumedCapacity,
		Credential:               o.Credential,
		ID:                       o.ID,
		Instances:                o.Instances,
		IsContainerGroup:         o.IsContainerGroup,
		JobsRunning:              o.JobsRunning,
		JobsTotal:                o.JobsTotal,
		Name:                     o.Name,
		PercentCapacityRemaining: o.PercentCapacityRemaining,
		PodSpecOverride:          o.PodSpecOverride,
		PolicyInstanceList:       o.PolicyInstanceList,
		PolicyInstanceMinimum:    o.PolicyInstanceMinimum,
		PolicyInstancePercentage: o.PolicyInstancePercentage,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for InstanceGroup
func (o *instanceGroupTerraformModel) BodyRequest() (req instanceGroupBodyRequestModel) {
	req.Credential = o.Credential.ValueInt64()
	req.IsContainerGroup = o.IsContainerGroup.ValueBool()
	req.Name = o.Name.ValueString()
	req.PodSpecOverride = o.PodSpecOverride.ValueString()
	req.PolicyInstanceList = json.RawMessage(o.PolicyInstanceList.ValueString())
	req.PolicyInstanceMinimum = o.PolicyInstanceMinimum.ValueInt64()
	req.PolicyInstancePercentage = o.PolicyInstancePercentage.ValueInt64()
	return
}

func (o *instanceGroupTerraformModel) setCapacity(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Capacity, data)
}

func (o *instanceGroupTerraformModel) setConsumedCapacity(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetFloat64(&o.ConsumedCapacity, data)
}

func (o *instanceGroupTerraformModel) setCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Credential, data)
}

func (o *instanceGroupTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *instanceGroupTerraformModel) setInstances(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Instances, data)
}

func (o *instanceGroupTerraformModel) setIsContainerGroup(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsContainerGroup, data)
}

func (o *instanceGroupTerraformModel) setJobsRunning(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.JobsRunning, data)
}

func (o *instanceGroupTerraformModel) setJobsTotal(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.JobsTotal, data)
}

func (o *instanceGroupTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *instanceGroupTerraformModel) setPercentCapacityRemaining(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetFloat64(&o.PercentCapacityRemaining, data)
}

func (o *instanceGroupTerraformModel) setPodSpecOverride(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.PodSpecOverride, data, false)
}

func (o *instanceGroupTerraformModel) setPolicyInstanceList(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.PolicyInstanceList, data, false)
}

func (o *instanceGroupTerraformModel) setPolicyInstanceMinimum(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.PolicyInstanceMinimum, data)
}

func (o *instanceGroupTerraformModel) setPolicyInstancePercentage(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.PolicyInstancePercentage, data)
}

func (o *instanceGroupTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setCapacity(data["capacity"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setConsumedCapacity(data["consumed_capacity"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCredential(data["credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInstances(data["instances"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsContainerGroup(data["is_container_group"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobsRunning(data["jobs_running"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobsTotal(data["jobs_total"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPercentCapacityRemaining(data["percent_capacity_remaining"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPodSpecOverride(data["pod_spec_override"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPolicyInstanceList(data["policy_instance_list"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPolicyInstanceMinimum(data["policy_instance_minimum"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPolicyInstancePercentage(data["policy_instance_percentage"]); dg.HasError() {
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
