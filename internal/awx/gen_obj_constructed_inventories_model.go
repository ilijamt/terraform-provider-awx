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
	return constructedInventoriesTerraformModel{
		Description:                  o.Description,
		HasActiveFailures:            o.HasActiveFailures,
		HasInventorySources:          o.HasInventorySources,
		HostsWithActiveFailures:      o.HostsWithActiveFailures,
		ID:                           o.ID,
		InventorySourcesWithFailures: o.InventorySourcesWithFailures,
		Kind:                         o.Kind,
		Limit:                        o.Limit,
		Name:                         o.Name,
		Organization:                 o.Organization,
		PendingDeletion:              o.PendingDeletion,
		PreventInstanceGroupFallback: o.PreventInstanceGroupFallback,
		SourceVars:                   o.SourceVars,
		TotalGroups:                  o.TotalGroups,
		TotalHosts:                   o.TotalHosts,
		TotalInventorySources:        o.TotalInventorySources,
		UpdateCacheTimeout:           o.UpdateCacheTimeout,
		Variables:                    o.Variables,
		Verbosity:                    o.Verbosity,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for ConstructedInventories
func (o *constructedInventoriesTerraformModel) BodyRequest() (req constructedInventoriesBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.SourceVars = o.SourceVars.ValueString()
	req.UpdateCacheTimeout = o.UpdateCacheTimeout.ValueInt64()
	req.Variables = json.RawMessage(o.Variables.ValueString())
	req.Verbosity = o.Verbosity.ValueInt64()
	return
}

func (o *constructedInventoriesTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *constructedInventoriesTerraformModel) setHasActiveFailures(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.HasActiveFailures, data)
}

func (o *constructedInventoriesTerraformModel) setHasInventorySources(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.HasInventorySources, data)
}

func (o *constructedInventoriesTerraformModel) setHostsWithActiveFailures(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data)
}

func (o *constructedInventoriesTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *constructedInventoriesTerraformModel) setInventorySourcesWithFailures(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data)
}

func (o *constructedInventoriesTerraformModel) setKind(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Kind, data, false)
}

func (o *constructedInventoriesTerraformModel) setLimit(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *constructedInventoriesTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *constructedInventoriesTerraformModel) setOrganization(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *constructedInventoriesTerraformModel) setPendingDeletion(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.PendingDeletion, data)
}

func (o *constructedInventoriesTerraformModel) setPreventInstanceGroupFallback(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data)
}

func (o *constructedInventoriesTerraformModel) setSourceVars(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SourceVars, data, false)
}

func (o *constructedInventoriesTerraformModel) setTotalGroups(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.TotalGroups, data)
}

func (o *constructedInventoriesTerraformModel) setTotalHosts(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.TotalHosts, data)
}

func (o *constructedInventoriesTerraformModel) setTotalInventorySources(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.TotalInventorySources, data)
}

func (o *constructedInventoriesTerraformModel) setUpdateCacheTimeout(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.UpdateCacheTimeout, data)
}

func (o *constructedInventoriesTerraformModel) setVariables(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.Variables, data, false)
}

func (o *constructedInventoriesTerraformModel) setVerbosity(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Verbosity, data)
}

func (o *constructedInventoriesTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
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
	if dg, _ := o.setLimit(data["limit"]); dg.HasError() {
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
	if dg, _ := o.setSourceVars(data["source_vars"]); dg.HasError() {
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
	if dg, _ := o.setUpdateCacheTimeout(data["update_cache_timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVariables(data["variables"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVerbosity(data["verbosity"]); dg.HasError() {
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

type constructedInventoriesObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
