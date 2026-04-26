package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type inventoryTerraformModel struct {
	Description                  types.String `tfsdk:"description" json:"description"`
	HasActiveFailures            types.Bool   `tfsdk:"has_active_failures" json:"has_active_failures"`
	HasInventorySources          types.Bool   `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	HostFilter                   types.String `tfsdk:"host_filter" json:"host_filter"`
	HostsWithActiveFailures      types.Int64  `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	ID                           types.Int64  `tfsdk:"id" json:"id"`
	InventorySourcesWithFailures types.Int64  `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	Kind                         types.String `tfsdk:"kind" json:"kind"`
	Name                         types.String `tfsdk:"name" json:"name"`
	Organization                 types.Int64  `tfsdk:"organization" json:"organization"`
	PendingDeletion              types.Bool   `tfsdk:"pending_deletion" json:"pending_deletion"`
	PreventInstanceGroupFallback types.Bool   `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	TotalGroups                  types.Int64  `tfsdk:"total_groups" json:"total_groups"`
	TotalHosts                   types.Int64  `tfsdk:"total_hosts" json:"total_hosts"`
	TotalInventorySources        types.Int64  `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	Variables                    types.String `tfsdk:"variables" json:"variables"`
}

func (o *inventoryTerraformModel) Clone() inventoryTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.HasActiveFailures, data["has_active_failures"]))
	collect(helpers.AttrValueSetBool(&o.HasInventorySources, data["has_inventory_sources"]))
	collect(helpers.AttrValueSetString(&o.HostFilter, data["host_filter"], false))
	collect(helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data["hosts_with_active_failures"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data["inventory_sources_with_failures"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetBool(&o.PendingDeletion, data["pending_deletion"]))
	collect(helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data["prevent_instance_group_fallback"]))
	collect(helpers.AttrValueSetInt64(&o.TotalGroups, data["total_groups"]))
	collect(helpers.AttrValueSetInt64(&o.TotalHosts, data["total_hosts"]))
	collect(helpers.AttrValueSetInt64(&o.TotalInventorySources, data["total_inventory_sources"]))
	collect(helpers.AttrValueSetJsonYamlString(&o.Variables, data["variables"], false))
	return diags, nil
}

type inventoryBodyRequestModel struct {
	Description                  string          `json:"description,omitempty"`
	HostFilter                   string          `json:"host_filter,omitempty"`
	Kind                         string          `json:"kind,omitempty"`
	Name                         string          `json:"name"`
	Organization                 int64           `json:"organization"`
	PreventInstanceGroupFallback bool            `json:"prevent_instance_group_fallback"`
	Variables                    json.RawMessage `json:"variables,omitempty"`
}
