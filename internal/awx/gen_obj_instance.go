package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type instanceTerraformModel struct {
	Capacity                 types.Int64   `tfsdk:"capacity" json:"capacity"`
	CapacityAdjustment       types.Float64 `tfsdk:"capacity_adjustment" json:"capacity_adjustment"`
	ConsumedCapacity         types.Int64   `tfsdk:"consumed_capacity" json:"consumed_capacity"`
	Cpu                      types.Float64 `tfsdk:"cpu" json:"cpu"`
	CpuCapacity              types.Int64   `tfsdk:"cpu_capacity" json:"cpu_capacity"`
	Enabled                  types.Bool    `tfsdk:"enabled" json:"enabled"`
	Errors                   types.String  `tfsdk:"errors" json:"errors"`
	HealthCheckPending       types.Bool    `tfsdk:"health_check_pending" json:"health_check_pending"`
	HealthCheckStarted       types.String  `tfsdk:"health_check_started" json:"health_check_started"`
	Hostname                 types.String  `tfsdk:"hostname" json:"hostname"`
	ID                       types.Int64   `tfsdk:"id" json:"id"`
	IpAddress                types.String  `tfsdk:"ip_address" json:"ip_address"`
	JobsRunning              types.Int64   `tfsdk:"jobs_running" json:"jobs_running"`
	JobsTotal                types.Int64   `tfsdk:"jobs_total" json:"jobs_total"`
	LastHealthCheck          types.String  `tfsdk:"last_health_check" json:"last_health_check"`
	LastSeen                 types.String  `tfsdk:"last_seen" json:"last_seen"`
	ListenerPort             types.Int64   `tfsdk:"listener_port" json:"listener_port"`
	Managed                  types.Bool    `tfsdk:"managed" json:"managed"`
	ManagedByPolicy          types.Bool    `tfsdk:"managed_by_policy" json:"managed_by_policy"`
	MemCapacity              types.Int64   `tfsdk:"mem_capacity" json:"mem_capacity"`
	Memory                   types.Int64   `tfsdk:"memory" json:"memory"`
	NodeState                types.String  `tfsdk:"node_state" json:"node_state"`
	NodeType                 types.String  `tfsdk:"node_type" json:"node_type"`
	Peers                    types.List    `tfsdk:"peers" json:"peers"`
	PeersFromControlNodes    types.Bool    `tfsdk:"peers_from_control_nodes" json:"peers_from_control_nodes"`
	PercentCapacityRemaining types.Float64 `tfsdk:"percent_capacity_remaining" json:"percent_capacity_remaining"`
	Protocol                 types.String  `tfsdk:"protocol" json:"protocol"`
	ReversePeers             types.List    `tfsdk:"reverse_peers" json:"reverse_peers"`
	Uuid                     types.String  `tfsdk:"uuid" json:"uuid"`
	Version                  types.String  `tfsdk:"version" json:"version"`
}

func (o *instanceTerraformModel) Clone() instanceTerraformModel {
	return *o
}

func (o *instanceTerraformModel) BodyRequest() *instanceBodyRequestModel {
	var req instanceBodyRequestModel
	req.CapacityAdjustment = o.CapacityAdjustment.ValueFloat64()
	req.Enabled = o.Enabled.ValueBool()
	req.Hostname = o.Hostname.ValueString()
	req.ListenerPort = o.ListenerPort.ValueInt64()
	req.ManagedByPolicy = o.ManagedByPolicy.ValueBool()
	req.NodeState = o.NodeState.ValueString()
	req.NodeType = o.NodeType.ValueString()
	req.Peers = helpers.ListAsInt64Slice(o.Peers)
	req.PeersFromControlNodes = o.PeersFromControlNodes.ValueBool()
	return &req
}

func (o *instanceTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Capacity, data["capacity"]))
	collect(helpers.AttrValueSetFloat64(&o.CapacityAdjustment, data["capacity_adjustment"]))
	collect(helpers.AttrValueSetInt64(&o.ConsumedCapacity, data["consumed_capacity"]))
	collect(helpers.AttrValueSetFloat64(&o.Cpu, data["cpu"]))
	collect(helpers.AttrValueSetInt64(&o.CpuCapacity, data["cpu_capacity"]))
	collect(helpers.AttrValueSetBool(&o.Enabled, data["enabled"]))
	collect(helpers.AttrValueSetString(&o.Errors, data["errors"], false))
	collect(helpers.AttrValueSetBool(&o.HealthCheckPending, data["health_check_pending"]))
	collect(helpers.AttrValueSetString(&o.HealthCheckStarted, data["health_check_started"], false))
	collect(helpers.AttrValueSetString(&o.Hostname, data["hostname"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.IpAddress, data["ip_address"], false))
	collect(helpers.AttrValueSetInt64(&o.JobsRunning, data["jobs_running"]))
	collect(helpers.AttrValueSetInt64(&o.JobsTotal, data["jobs_total"]))
	collect(helpers.AttrValueSetString(&o.LastHealthCheck, data["last_health_check"], false))
	collect(helpers.AttrValueSetString(&o.LastSeen, data["last_seen"], false))
	collect(helpers.AttrValueSetInt64(&o.ListenerPort, data["listener_port"]))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetBool(&o.ManagedByPolicy, data["managed_by_policy"]))
	collect(helpers.AttrValueSetInt64(&o.MemCapacity, data["mem_capacity"]))
	collect(helpers.AttrValueSetInt64(&o.Memory, data["memory"]))
	collect(helpers.AttrValueSetString(&o.NodeState, data["node_state"], false))
	collect(helpers.AttrValueSetString(&o.NodeType, data["node_type"], false))
	collect(helpers.AttrValueSetListInt64(&o.Peers, data["peers"]))
	collect(helpers.AttrValueSetBool(&o.PeersFromControlNodes, data["peers_from_control_nodes"]))
	collect(helpers.AttrValueSetFloat64(&o.PercentCapacityRemaining, data["percent_capacity_remaining"]))
	collect(helpers.AttrValueSetString(&o.Protocol, data["protocol"], false))
	collect(helpers.AttrValueSetListInt64(&o.ReversePeers, data["reverse_peers"]))
	collect(helpers.AttrValueSetString(&o.Uuid, data["uuid"], false))
	collect(helpers.AttrValueSetString(&o.Version, data["version"], false))
	return diags, nil
}

type instanceBodyRequestModel struct {
	CapacityAdjustment    float64 `json:"capacity_adjustment"`
	Enabled               bool    `json:"enabled"`
	Hostname              string  `json:"hostname"`
	ListenerPort          int64   `json:"listener_port,omitempty"`
	ManagedByPolicy       bool    `json:"managed_by_policy"`
	NodeState             string  `json:"node_state,omitempty"`
	NodeType              string  `json:"node_type,omitempty"`
	Peers                 []int64 `json:"peers,omitempty"`
	PeersFromControlNodes bool    `json:"peers_from_control_nodes"`
}

type instanceResource = framework.GenericResource[instanceTerraformModel, instanceBodyRequestModel, *instanceTerraformModel]

// NewInstanceResource is a helper function to simplify the provider implementation.
func NewInstanceResource() resource.Resource {
	return &instanceResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "instance", Endpoint: "/api/v2/instances/"}},
		Cfg: framework.ResourceCfg[instanceTerraformModel, instanceBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"capacity_adjustment": schema.Float64Attribute{
						Description: "Capacity adjustment",
						Optional:    true,
						Computed:    true,
						Default:     float64default.StaticFloat64(1),
					},
					"enabled": schema.BoolAttribute{
						Description: "Enabled",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(true),
					},
					"hostname": schema.StringAttribute{
						Description: "Hostname",
						Required:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(250),
						},
					},
					"listener_port": schema.Int64Attribute{
						Description: "Listener port",
						Optional:    true,
						Computed:    true,
					},
					"managed_by_policy": schema.BoolAttribute{
						Description: "Managed by policy",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(true),
					},
					"node_state": schema.StringAttribute{
						Description: "Indicates the current life cycle stage of this instance.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"provisioning",
								"provision-fail",
								"installed",
								"ready",
								"unavailable",
								"deprovisioning",
								"deprovision-fail",
							),
						},
					},
					"node_type": schema.StringAttribute{
						Description: "Role that this node plays in the mesh.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`execution`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"control",
								"execution",
								"hybrid",
								"hop",
							),
						},
					},
					"peers": schema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Primary keys of receptor addresses to peer to.",
						Optional:    true,
						Computed:    true,
					},
					"peers_from_control_nodes": schema.BoolAttribute{
						Description: "Peers from control nodes",
						Optional:    true,
						Computed:    true,
					},
					"capacity": schema.Int64Attribute{
						Description: "Capacity",
						Computed:    true,
					},
					"consumed_capacity": schema.Int64Attribute{
						Description: "Consumed capacity",
						Computed:    true,
					},
					"cpu": schema.Float64Attribute{
						Description: "Cpu",
						Computed:    true,
					},
					"cpu_capacity": schema.Int64Attribute{
						Description: "Cpu capacity",
						Computed:    true,
					},
					"errors": schema.StringAttribute{
						Description: "Any error details from the last health check.",
						Computed:    true,
					},
					"health_check_pending": schema.BoolAttribute{
						Description: "Health check pending",
						Computed:    true,
					},
					"health_check_started": schema.StringAttribute{
						Description: "The last time a health check was initiated on this instance.",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this instance.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"ip_address": schema.StringAttribute{
						Description: "Ip address",
						Computed:    true,
					},
					"jobs_running": schema.Int64Attribute{
						Description: "Count of jobs in the running or waiting state that are targeted for this instance",
						Computed:    true,
					},
					"jobs_total": schema.Int64Attribute{
						Description: "Count of all jobs that target this instance",
						Computed:    true,
					},
					"last_health_check": schema.StringAttribute{
						Description: "Last time a health check was ran on this instance to refresh cpu, memory, and capacity.",
						Computed:    true,
					},
					"last_seen": schema.StringAttribute{
						Description: "Last time instance ran its heartbeat task for main cluster nodes. Last known connection to receptor mesh for execution nodes.",
						Computed:    true,
					},
					"managed": schema.BoolAttribute{
						Description: "If True, this instance is managed by the control plane.",
						Computed:    true,
					},
					"mem_capacity": schema.Int64Attribute{
						Description: "Mem capacity",
						Computed:    true,
					},
					"memory": schema.Int64Attribute{
						Description: "Total system memory of this instance in bytes.",
						Computed:    true,
					},
					"percent_capacity_remaining": schema.Float64Attribute{
						Description: "Percent capacity remaining",
						Computed:    true,
					},
					"protocol": schema.StringAttribute{
						Description: "Protocol",
						Computed:    true,
					},
					"reverse_peers": schema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Reverse peers",
						Computed:    true,
					},
					"uuid": schema.StringAttribute{
						Description: "Uuid",
						Computed:    true,
					},
					"version": schema.StringAttribute{
						Description: "Version",
						Computed:    true,
					},
				},
			},
			IDAccessor: func(m *instanceTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			SoftDelete: map[string]any{
				"node_state": "deprovisioning",
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Instance",
		},
	}
}

type instanceDataSource = framework.GenericDataSource[instanceTerraformModel, *instanceTerraformModel]

// NewInstanceDataSource is a helper function to instantiate the Instance data source.
func NewInstanceDataSource() datasource.DataSource {
	return &instanceDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "instance", Endpoint: "/api/v2/instances/"}},
		Cfg: framework.DataSourceCfg[instanceTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"capacity": dschema.Int64Attribute{
						Description: "Capacity",
						Computed:    true,
					},
					"capacity_adjustment": dschema.Float64Attribute{
						Description: "Capacity adjustment",
						Computed:    true,
					},
					"consumed_capacity": dschema.Int64Attribute{
						Description: "Consumed capacity",
						Computed:    true,
					},
					"cpu": dschema.Float64Attribute{
						Description: "Cpu",
						Computed:    true,
					},
					"cpu_capacity": dschema.Int64Attribute{
						Description: "Cpu capacity",
						Computed:    true,
					},
					"enabled": dschema.BoolAttribute{
						Description: "Enabled",
						Computed:    true,
					},
					"errors": dschema.StringAttribute{
						Description: "Any error details from the last health check.",
						Computed:    true,
					},
					"health_check_pending": dschema.BoolAttribute{
						Description: "Health check pending",
						Computed:    true,
					},
					"health_check_started": dschema.StringAttribute{
						Description: "The last time a health check was initiated on this instance.",
						Computed:    true,
					},
					"hostname": dschema.StringAttribute{
						Description: "Hostname",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("hostname"),
							),
						},
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this instance.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("hostname"),
							),
						},
					},
					"ip_address": dschema.StringAttribute{
						Description: "Ip address",
						Computed:    true,
					},
					"jobs_running": dschema.Int64Attribute{
						Description: "Count of jobs in the running or waiting state that are targeted for this instance",
						Computed:    true,
					},
					"jobs_total": dschema.Int64Attribute{
						Description: "Count of all jobs that target this instance",
						Computed:    true,
					},
					"last_health_check": dschema.StringAttribute{
						Description: "Last time a health check was ran on this instance to refresh cpu, memory, and capacity.",
						Computed:    true,
					},
					"last_seen": dschema.StringAttribute{
						Description: "Last time instance ran its heartbeat task for main cluster nodes. Last known connection to receptor mesh for execution nodes.",
						Computed:    true,
					},
					"listener_port": dschema.Int64Attribute{
						Description: "Listener port",
						Computed:    true,
					},
					"managed": dschema.BoolAttribute{
						Description: "If True, this instance is managed by the control plane.",
						Computed:    true,
					},
					"managed_by_policy": dschema.BoolAttribute{
						Description: "Managed by policy",
						Computed:    true,
					},
					"mem_capacity": dschema.Int64Attribute{
						Description: "Mem capacity",
						Computed:    true,
					},
					"memory": dschema.Int64Attribute{
						Description: "Total system memory of this instance in bytes.",
						Computed:    true,
					},
					"node_state": dschema.StringAttribute{
						Description: "Indicates the current life cycle stage of this instance.",
						Computed:    true,
					},
					"node_type": dschema.StringAttribute{
						Description: "Role that this node plays in the mesh.",
						Computed:    true,
					},
					"peers": dschema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Primary keys of receptor addresses to peer to.",
						Computed:    true,
					},
					"peers_from_control_nodes": dschema.BoolAttribute{
						Description: "Peers from control nodes",
						Computed:    true,
					},
					"percent_capacity_remaining": dschema.Float64Attribute{
						Description: "Percent capacity remaining",
						Computed:    true,
					},
					"protocol": dschema.StringAttribute{
						Description: "Protocol",
						Computed:    true,
					},
					"reverse_peers": dschema.ListAttribute{
						ElementType: types.Int64Type,
						Description: "Reverse peers",
						Computed:    true,
					},
					"uuid": dschema.StringAttribute{
						Description: "Uuid",
						Computed:    true,
					},
					"version": dschema.StringAttribute{
						Description: "Version",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_hostname", URLSuffix: "?hostname=%s", Fields: []framework.SearchField{
					{Name: "hostname", Type: "string", URLEscape: true},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Instance",
		},
	}
}
