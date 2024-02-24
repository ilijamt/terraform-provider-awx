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
	// HasActiveFailures "This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed."
	HasActiveFailures types.Bool `tfsdk:"has_active_failures" json:"has_active_failures"`
	// HasInventorySources "This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources."
	HasInventorySources types.Bool `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	// HostFilter "Filter that will be applied to the hosts of this inventory."
	HostFilter types.String `tfsdk:"host_filter" json:"host_filter"`
	// HostsWithActiveFailures "This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures."
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
	// TotalGroups "This field is deprecated and will be removed in a future release. Total number of groups in this inventory."
	TotalGroups types.Int64 `tfsdk:"total_groups" json:"total_groups"`
	// TotalHosts "This field is deprecated and will be removed in a future release. Total number of hosts in this inventory."
	TotalHosts types.Int64 `tfsdk:"total_hosts" json:"total_hosts"`
	// TotalInventorySources "Total number of external inventory sources configured within this inventory."
	TotalInventorySources types.Int64 `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	// Variables "Inventory variables in JSON format"
	Variables types.String `tfsdk:"variables" json:"variables"`
}

// Clone the object
func (o *inventoryTerraformModel) Clone() inventoryTerraformModel {
	return inventoryTerraformModel{
		Description:                  o.Description,
		HasActiveFailures:            o.HasActiveFailures,
		HasInventorySources:          o.HasInventorySources,
		HostFilter:                   o.HostFilter,
		HostsWithActiveFailures:      o.HostsWithActiveFailures,
		ID:                           o.ID,
		InventorySourcesWithFailures: o.InventorySourcesWithFailures,
		Kind:                         o.Kind,
		Name:                         o.Name,
		Organization:                 o.Organization,
		PendingDeletion:              o.PendingDeletion,
		PreventInstanceGroupFallback: o.PreventInstanceGroupFallback,
		TotalGroups:                  o.TotalGroups,
		TotalHosts:                   o.TotalHosts,
		TotalInventorySources:        o.TotalInventorySources,
		Variables:                    o.Variables,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Inventory
func (o *inventoryTerraformModel) BodyRequest() (req inventoryBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.HostFilter = o.HostFilter.ValueString()
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.Variables = json.RawMessage(o.Variables.String())
	return
}

func (o *inventoryTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *inventoryTerraformModel) setHasActiveFailures(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.HasActiveFailures, data)
}

func (o *inventoryTerraformModel) setHasInventorySources(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.HasInventorySources, data)
}

func (o *inventoryTerraformModel) setHostFilter(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.HostFilter, data, false)
}

func (o *inventoryTerraformModel) setHostsWithActiveFailures(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data)
}

func (o *inventoryTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *inventoryTerraformModel) setInventorySourcesWithFailures(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data)
}

func (o *inventoryTerraformModel) setKind(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Kind, data, false)
}

func (o *inventoryTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *inventoryTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *inventoryTerraformModel) setPendingDeletion(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.PendingDeletion, data)
}

func (o *inventoryTerraformModel) setPreventInstanceGroupFallback(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data)
}

func (o *inventoryTerraformModel) setTotalGroups(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.TotalGroups, data)
}

func (o *inventoryTerraformModel) setTotalHosts(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.TotalHosts, data)
}

func (o *inventoryTerraformModel) setTotalInventorySources(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.TotalInventorySources, data)
}

func (o *inventoryTerraformModel) setVariables(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonYamlString(&o.Variables, data, false)
}

func (o *inventoryTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHasActiveFailures(data["has_active_failures"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHasInventorySources(data["has_inventory_sources"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHostFilter(data["host_filter"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHostsWithActiveFailures(data["hosts_with_active_failures"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventorySourcesWithFailures(data["inventory_sources_with_failures"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setKind(data["kind"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPendingDeletion(data["pending_deletion"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPreventInstanceGroupFallback(data["prevent_instance_group_fallback"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTotalGroups(data["total_groups"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTotalHosts(data["total_hosts"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTotalInventorySources(data["total_inventory_sources"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVariables(data["variables"]); dg.HasError() {
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

type inventoryObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
