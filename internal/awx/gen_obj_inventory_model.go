package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// inventoryTerraformModel maps the schema for Inventory when using Data Source
type inventoryTerraformModel struct {
	// Description "Optional description of this inventory."
	Description types.String `tfsdk:"description" json:"description"`
	// HasActiveFailures "Flag indicating whether any hosts in this inventory have failed."
	HasActiveFailures types.Bool `tfsdk:"has_active_failures" json:"has_active_failures"`
	// HasInventorySources "Flag indicating whether this inventory has any external inventory sources."
	HasInventorySources types.Bool `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	// HostFilter "Filter that will be applied to the hosts of this inventory."
	HostFilter types.String `tfsdk:"host_filter" json:"host_filter"`
	// HostsWithActiveFailures "Number of hosts in this inventory with active failures."
	HostsWithActiveFailures types.Int64 `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	// ID "Database ID for this inventory."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InventorySourcesWithFailures "Number of external inventory sources in this inventory with failures."
	InventorySourcesWithFailures types.Int64 `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	// Kind "Kind of inventory being represented."
	Kind types.String `tfsdk:"kind" json:"kind"`
	// Name "Name of this inventory."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Organization containing this inventory."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// PendingDeletion "Flag indicating the inventory is being deleted."
	PendingDeletion types.Bool `tfsdk:"pending_deletion" json:"pending_deletion"`
	// PreventInstanceGroupFallback "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback types.Bool `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	// TotalGroups "Total number of groups in this inventory."
	TotalGroups types.Int64 `tfsdk:"total_groups" json:"total_groups"`
	// TotalHosts "Total number of hosts in this inventory."
	TotalHosts types.Int64 `tfsdk:"total_hosts" json:"total_hosts"`
	// TotalInventorySources "Total number of external inventory sources configured within this inventory."
	TotalInventorySources types.Int64 `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	// Variables "Inventory variables in JSON format"
	Variables types.String `tfsdk:"variables" json:"variables"`
}

// Clone the object
func (o *inventoryTerraformModel) Clone() inventoryTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Inventory
func (o *inventoryTerraformModel) BodyRequest() *inventoryBodyRequestModel {
	var req inventoryBodyRequestModel
	req.Description = o.Description.ValueString()
	req.HostFilter = o.HostFilter.ValueString()
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.Variables = json.RawMessage(o.Variables.String())
	return &req
}

func (o *inventoryTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.HasActiveFailures, data["has_active_failures"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.HasInventorySources, data["has_inventory_sources"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.HostFilter, data["host_filter"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data["hosts_with_active_failures"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data["inventory_sources_with_failures"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Kind, data["kind"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.PendingDeletion, data["pending_deletion"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data["prevent_instance_group_fallback"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.TotalGroups, data["total_groups"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.TotalHosts, data["total_hosts"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.TotalInventorySources, data["total_inventory_sources"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonYamlString(&o.Variables, data["variables"], false)
		diags.Append(dg...)
	}
	return diags, nil
}

// inventoryBodyRequestModel maps the schema for Inventory for creating and updating the data
type inventoryBodyRequestModel struct {
	// Description "Optional description of this inventory."
	Description string `json:"description,omitempty"`
	// HostFilter "Filter that will be applied to the hosts of this inventory."
	HostFilter string `json:"host_filter,omitempty"`
	// Kind "Kind of inventory being represented."
	Kind string `json:"kind,omitempty"`
	// Name "Name of this inventory."
	Name string `json:"name"`
	// Organization "Organization containing this inventory."
	Organization int64 `json:"organization"`
	// PreventInstanceGroupFallback "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback bool `json:"prevent_instance_group_fallback"`
	// Variables "Inventory variables in JSON format"
	Variables json.RawMessage `json:"variables,omitempty"`
}
