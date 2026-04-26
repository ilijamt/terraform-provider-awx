package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type constructedInventoriesTerraformModel struct {
	Description                  types.String `tfsdk:"description" json:"description"`
	HasActiveFailures            types.Bool   `tfsdk:"has_active_failures" json:"has_active_failures"`
	HasInventorySources          types.Bool   `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	HostsWithActiveFailures      types.Int64  `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	ID                           types.Int64  `tfsdk:"id" json:"id"`
	InventorySourcesWithFailures types.Int64  `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	Kind                         types.String `tfsdk:"kind" json:"kind"`
	Limit                        types.String `tfsdk:"limit" json:"limit"`
	Name                         types.String `tfsdk:"name" json:"name"`
	Organization                 types.Int64  `tfsdk:"organization" json:"organization"`
	PendingDeletion              types.Bool   `tfsdk:"pending_deletion" json:"pending_deletion"`
	PreventInstanceGroupFallback types.Bool   `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	SourceVars                   types.String `tfsdk:"source_vars" json:"source_vars"`
	TotalGroups                  types.Int64  `tfsdk:"total_groups" json:"total_groups"`
	TotalHosts                   types.Int64  `tfsdk:"total_hosts" json:"total_hosts"`
	TotalInventorySources        types.Int64  `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	UpdateCacheTimeout           types.Int64  `tfsdk:"update_cache_timeout" json:"update_cache_timeout"`
	Variables                    types.String `tfsdk:"variables" json:"variables"`
	Verbosity                    types.Int64  `tfsdk:"verbosity" json:"verbosity"`
}

func (o *constructedInventoriesTerraformModel) Clone() constructedInventoriesTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.HasActiveFailures, data["has_active_failures"]))
	collect(helpers.AttrValueSetBool(&o.HasInventorySources, data["has_inventory_sources"]))
	collect(helpers.AttrValueSetInt64(&o.HostsWithActiveFailures, data["hosts_with_active_failures"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.InventorySourcesWithFailures, data["inventory_sources_with_failures"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetBool(&o.PendingDeletion, data["pending_deletion"]))
	collect(helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data["prevent_instance_group_fallback"]))
	collect(helpers.AttrValueSetString(&o.SourceVars, data["source_vars"], false))
	collect(helpers.AttrValueSetInt64(&o.TotalGroups, data["total_groups"]))
	collect(helpers.AttrValueSetInt64(&o.TotalHosts, data["total_hosts"]))
	collect(helpers.AttrValueSetInt64(&o.TotalInventorySources, data["total_inventory_sources"]))
	collect(helpers.AttrValueSetInt64(&o.UpdateCacheTimeout, data["update_cache_timeout"]))
	collect(helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false))
	collect(helpers.AttrValueSetInt64(&o.Verbosity, data["verbosity"]))
	return diags, nil
}

type constructedInventoriesBodyRequestModel struct {
	Description                  string          `json:"description,omitempty"`
	Limit                        string          `json:"limit,omitempty"`
	Name                         string          `json:"name"`
	Organization                 int64           `json:"organization"`
	PreventInstanceGroupFallback bool            `json:"prevent_instance_group_fallback"`
	SourceVars                   string          `json:"source_vars,omitempty"`
	UpdateCacheTimeout           int64           `json:"update_cache_timeout,omitempty"`
	Variables                    json.RawMessage `json:"variables,omitempty"`
	Verbosity                    int64           `json:"verbosity,omitempty"`
}
