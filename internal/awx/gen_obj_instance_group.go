package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
func (o instanceGroupTerraformModel) Clone() instanceGroupTerraformModel {
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
func (o instanceGroupTerraformModel) BodyRequest() (req instanceGroupBodyRequestModel) {
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
	// Decode "capacity"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Capacity = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Capacity = types.Int64Value(val)
	} else {
		o.Capacity = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setConsumedCapacity(data any) (d diag.Diagnostics, err error) {
	// Decode "consumed_capacity"
	if val, ok := data.(json.Number); ok {
		v, err := val.Float64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to float64", val),
				err.Error(),
			)
			return d, err
		}
		o.ConsumedCapacity = types.Float64Value(v)
	} else if val, ok := data.(float64); ok {
		o.ConsumedCapacity = types.Float64Value(val)
	} else {
		o.ConsumedCapacity = types.Float64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setCredential(data any) (d diag.Diagnostics, err error) {
	// Decode "credential"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Credential = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Credential = types.Int64Value(val)
	} else {
		o.Credential = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	// Decode "id"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.ID = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.ID = types.Int64Value(val)
	} else {
		o.ID = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setInstances(data any) (d diag.Diagnostics, err error) {
	// Decode "instances"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Instances = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Instances = types.Int64Value(val)
	} else {
		o.Instances = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setIsContainerGroup(data any) (d diag.Diagnostics, err error) {
	// Decode "is_container_group"
	if val, ok := data.(bool); ok {
		o.IsContainerGroup = types.BoolValue(val)
	} else {
		o.IsContainerGroup = types.BoolNull()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setJobsRunning(data any) (d diag.Diagnostics, err error) {
	// Decode "jobs_running"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.JobsRunning = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.JobsRunning = types.Int64Value(val)
	} else {
		o.JobsRunning = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setJobsTotal(data any) (d diag.Diagnostics, err error) {
	// Decode "jobs_total"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.JobsTotal = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.JobsTotal = types.Int64Value(val)
	} else {
		o.JobsTotal = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	// Decode "name"
	if val, ok := data.(string); ok {
		o.Name = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Name = types.StringValue(val.String())
	} else {
		o.Name = types.StringNull()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setPercentCapacityRemaining(data any) (d diag.Diagnostics, err error) {
	// Decode "percent_capacity_remaining"
	if val, ok := data.(json.Number); ok {
		v, err := val.Float64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to float64", val),
				err.Error(),
			)
			return d, err
		}
		o.PercentCapacityRemaining = types.Float64Value(v)
	} else if val, ok := data.(float64); ok {
		o.PercentCapacityRemaining = types.Float64Value(val)
	} else {
		o.PercentCapacityRemaining = types.Float64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setPodSpecOverride(data any) (d diag.Diagnostics, err error) {
	// Decode "pod_spec_override"
	if val, ok := data.(string); ok {
		o.PodSpecOverride = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.PodSpecOverride = types.StringValue(val.String())
	} else {
		o.PodSpecOverride = types.StringNull()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setPolicyInstanceList(data any) (d diag.Diagnostics, err error) {
	// Decode "policy_instance_list"
	if val, ok := data.(string); ok {
		o.PolicyInstanceList = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.PolicyInstanceList = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.PolicyInstanceList = types.StringNull()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setPolicyInstanceMinimum(data any) (d diag.Diagnostics, err error) {
	// Decode "policy_instance_minimum"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.PolicyInstanceMinimum = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.PolicyInstanceMinimum = types.Int64Value(val)
	} else {
		o.PolicyInstanceMinimum = types.Int64Null()
	}
	return d, nil
}

func (o *instanceGroupTerraformModel) setPolicyInstancePercentage(data any) (d diag.Diagnostics, err error) {
	// Decode "policy_instance_percentage"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.PolicyInstancePercentage = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.PolicyInstancePercentage = types.Int64Value(val)
	} else {
		o.PolicyInstancePercentage = types.Int64Null()
	}
	return d, nil
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

var (
	_ datasource.DataSource              = &instanceGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &instanceGroupDataSource{}
)

// NewInstanceGroupDataSource is a helper function to instantiate the InstanceGroup data source.
func NewInstanceGroupDataSource() datasource.DataSource {
	return &instanceGroupDataSource{}
}

// instanceGroupDataSource is the data source implementation.
type instanceGroupDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *instanceGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/instance_groups/"
}

// Metadata returns the data source type name.
func (o *instanceGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance_group"
}

// GetSchema defines the schema for the data source.
func (o *instanceGroupDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"InstanceGroup",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"capacity": {
					Description: "Capacity",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"consumed_capacity": {
					Description: "Consumed capacity",
					Type:        types.Float64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"credential": {
					Description: "Credential",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this instance group.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
						),
					},
				},
				"instances": {
					Description: "Instances",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"is_container_group": {
					Description: "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"jobs_running": {
					Description: "Count of jobs in the running or waiting state that are targeted for this instance group",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"jobs_total": {
					Description: "Count of all jobs that target this instance group",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this instance group.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"percent_capacity_remaining": {
					Description: "Percent capacity remaining",
					Type:        types.Float64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"pod_spec_override": {
					Description: "Pod spec override",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"policy_instance_list": {
					Description: "List of exact-match Instances that will be assigned to this group",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"policy_instance_minimum": {
					Description: "Static minimum number of Instances that will be automatically assign to this group when new instances come online.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"policy_instance_percentage": {
					Description: "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *instanceGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state instanceGroupTerraformModel
	var err error
	var endpoint string
	var searchDefined bool

	// Only one group should evaluate to True, terraform should prevent from being able to set
	// the conflicting groups

	// Evaluate group 'by_id' based on the schema definition
	var groupByIdExists = func() bool {
		var groupByIdExists = true
		var paramsById = []any{o.endpoint}
		var attrID types.Int64
		req.Config.GetAttribute(ctx, path.Root("id"), &attrID)
		groupByIdExists = groupByIdExists && (!attrID.IsNull() && !attrID.IsUnknown())
		paramsById = append(paramsById, attrID.ValueInt64())
		if groupByIdExists {
			endpoint = p.Clean(fmt.Sprintf("%s/%d/", paramsById...))
		}
		return groupByIdExists
	}()
	searchDefined = searchDefined || groupByIdExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			fmt.Sprintf("missing configuration for one of the predefined search groups"),
			detailMessage,
		)
		return
	}

	// Creates a new request for InstanceGroup
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InstanceGroup on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for InstanceGroup
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InstanceGroup on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = extractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &instanceGroupResource{}
	_ resource.ResourceWithConfigure   = &instanceGroupResource{}
	_ resource.ResourceWithImportState = &instanceGroupResource{}
)

// NewInstanceGroupResource is a helper function to simplify the provider implementation.
func NewInstanceGroupResource() resource.Resource {
	return &instanceGroupResource{}
}

type instanceGroupResource struct {
	client   c.Client
	endpoint string
}

func (o *instanceGroupResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/instance_groups/"
}

func (o instanceGroupResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_instance_group"
}

func (o instanceGroupResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"InstanceGroup",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"credential": {
					Description: "Credential",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"is_container_group": {
					Description: "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this instance group.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(250),
					},
				},
				"pod_spec_override": {
					Description: "Pod spec override",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"apiVersion":"v1","kind":"Pod","metadata":{"namespace":"awx"},"spec":{"automountServiceAccountToken":false,"containers":[{"args":["ansible-runner","worker","--private-data-dir=/runner"],"image":"quay.io/ansible/awx-ee:latest","name":"worker","resources":{"requests":{"cpu":"250m","memory":"100Mi"}}}],"serviceAccountName":"default"}}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"policy_instance_list": {
					Description: "List of exact-match Instances that will be assigned to this group",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"policy_instance_minimum": {
					Description: "Static minimum number of Instances that will be automatically assign to this group when new instances come online.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"policy_instance_percentage": {
					Description: "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						int64validator.Between(0, 100),
					},
				},
				// Write only elements
				// Data only elements
				"capacity": {
					Description: "",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"consumed_capacity": {
					Description: "",
					Computed:    true,
					Type:        types.Float64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this instance group.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"instances": {
					Description: "",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"jobs_running": {
					Description: "Count of jobs in the running or waiting state that are targeted for this instance group",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"jobs_total": {
					Description: "Count of all jobs that target this instance group",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"percent_capacity_remaining": {
					Description: "",
					Computed:    true,
					Type:        types.Float64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *instanceGroupResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the InstanceGroup.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *instanceGroupResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state instanceGroupTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InstanceGroup
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InstanceGroup on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new InstanceGroup resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for InstanceGroup on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *instanceGroupResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state instanceGroupTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InstanceGroup
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InstanceGroup on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for InstanceGroup from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InstanceGroup on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *instanceGroupResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state instanceGroupTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InstanceGroup
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InstanceGroup on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new InstanceGroup resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for InstanceGroup on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *instanceGroupResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state instanceGroupTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InstanceGroup
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InstanceGroup on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing InstanceGroup
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for InstanceGroup on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
