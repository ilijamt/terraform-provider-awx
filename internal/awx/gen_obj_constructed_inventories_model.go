package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// constructedInventoriesTerraformModel maps the schema for ConstructedInventories when using Data Source
type constructedInventoriesTerraformModel struct {
	// Description "Optional description of this inventory."
	Description types.String `tfsdk:"description" json:"description"`
	// HasActiveFailures "Flag indicating whether any hosts in this inventory have failed."
	HasActiveFailures types.Bool `tfsdk:"has_active_failures" json:"has_active_failures"`
	// HasInventorySources "Flag indicating whether this inventory has any external inventory sources."
	HasInventorySources types.Bool `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	// HostsWithActiveFailures "Number of hosts in this inventory with active failures."
	HostsWithActiveFailures types.Int64 `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	// ID "Database ID for this inventory."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InventorySourcesWithFailures "Number of external inventory sources in this inventory with failures."
	InventorySourcesWithFailures types.Int64 `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	// Kind "Kind of inventory being represented."
	Kind types.String `tfsdk:"kind" json:"kind"`
	// Limit "The limit to restrict the returned hosts for the related auto-created inventory source, special to constructed inventory."
	Limit types.String `tfsdk:"limit" json:"limit"`
	// Name "Name of this inventory."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Organization containing this inventory."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// PendingDeletion "Flag indicating the inventory is being deleted."
	PendingDeletion types.Bool `tfsdk:"pending_deletion" json:"pending_deletion"`
	// PreventInstanceGroupFallback "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback types.Bool `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	// SourceVars "The source_vars for the related auto-created inventory source, special to constructed inventory."
	SourceVars types.String `tfsdk:"source_vars" json:"source_vars"`
	// TotalGroups "Total number of groups in this inventory."
	TotalGroups types.Int64 `tfsdk:"total_groups" json:"total_groups"`
	// TotalHosts "Total number of hosts in this inventory."
	TotalHosts types.Int64 `tfsdk:"total_hosts" json:"total_hosts"`
	// TotalInventorySources "Total number of external inventory sources configured within this inventory."
	TotalInventorySources types.Int64 `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	// UpdateCacheTimeout "The cache timeout for the related auto-created inventory source, special to constructed inventory"
	UpdateCacheTimeout types.Int64 `tfsdk:"update_cache_timeout" json:"update_cache_timeout"`
	// Variables "Inventory variables in JSON or YAML format."
	Variables types.String `tfsdk:"variables" json:"variables"`
	// Verbosity "The verbosity level for the related auto-created inventory source, special to constructed inventory"
	Verbosity types.Int64 `tfsdk:"verbosity" json:"verbosity"`
}

// Clone the object
func (o *constructedInventoriesTerraformModel) Clone() constructedInventoriesTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for ConstructedInventories
func (o *constructedInventoriesTerraformModel) BodyRequest() *constructedInventoriesBodyRequestModel {
	var req constructedInventoriesBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.SourceVars = o.SourceVars.ValueString()
	req.UpdateCacheTimeout = o.UpdateCacheTimeout.ValueInt64()
	req.Variables = json.RawMessage(o.Variables.ValueString())
	req.Verbosity = o.Verbosity.ValueInt64()
	return &req
}

func (o *constructedInventoriesTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
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
		dg, _ := helpers.AttrValueSetString(&o.Limit, data["limit"], false)
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
		dg, _ := helpers.AttrValueSetString(&o.SourceVars, data["source_vars"], false)
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
		dg, _ := helpers.AttrValueSetInt64(&o.UpdateCacheTimeout, data["update_cache_timeout"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Verbosity, data["verbosity"])
		diags.Append(dg...)
	}
	return diags, nil
}

// constructedInventoriesBodyRequestModel maps the schema for ConstructedInventories for creating and updating the data
type constructedInventoriesBodyRequestModel struct {
	// Description "Optional description of this inventory."
	Description string `json:"description,omitempty"`
	// Limit "The limit to restrict the returned hosts for the related auto-created inventory source, special to constructed inventory."
	Limit string `json:"limit,omitempty"`
	// Name "Name of this inventory."
	Name string `json:"name"`
	// Organization "Organization containing this inventory."
	Organization int64 `json:"organization"`
	// PreventInstanceGroupFallback "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback bool `json:"prevent_instance_group_fallback"`
	// SourceVars "The source_vars for the related auto-created inventory source, special to constructed inventory."
	SourceVars string `json:"source_vars,omitempty"`
	// UpdateCacheTimeout "The cache timeout for the related auto-created inventory source, special to constructed inventory"
	UpdateCacheTimeout int64 `json:"update_cache_timeout,omitempty"`
	// Variables "Inventory variables in JSON or YAML format."
	Variables json.RawMessage `json:"variables,omitempty"`
	// Verbosity "The verbosity level for the related auto-created inventory source, special to constructed inventory"
	Verbosity int64 `json:"verbosity,omitempty"`
}
