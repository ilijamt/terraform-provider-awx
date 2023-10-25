package awx_21_8_0

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// inventorySourceTerraformModel maps the schema for InventorySource when using Data Source
type inventorySourceTerraformModel struct {
	// Credential "Cloud credential to use for inventory updates."
	Credential types.Int64 `tfsdk:"credential" json:"credential"`
	// Description "Optional description of this inventory source."
	Description types.String `tfsdk:"description" json:"description"`
	// EnabledValue "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled"
	EnabledValue types.String `tfsdk:"enabled_value" json:"enabled_value"`
	// EnabledVar "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)"
	EnabledVar types.String `tfsdk:"enabled_var" json:"enabled_var"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment types.Int64 `tfsdk:"execution_environment" json:"execution_environment"`
	// HostFilter "Regex where only matching hosts will be imported."
	HostFilter types.String `tfsdk:"host_filter" json:"host_filter"`
	// ID "Database ID for this inventory source."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inventory ""
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// Name "Name of this inventory source."
	Name types.String `tfsdk:"name" json:"name"`
	// Overwrite "Overwrite local groups and hosts from remote inventory source."
	Overwrite types.Bool `tfsdk:"overwrite" json:"overwrite"`
	// OverwriteVars "Overwrite local variables from remote inventory source."
	OverwriteVars types.Bool `tfsdk:"overwrite_vars" json:"overwrite_vars"`
	// Source ""
	Source types.String `tfsdk:"source" json:"source"`
	// SourcePath ""
	SourcePath types.String `tfsdk:"source_path" json:"source_path"`
	// SourceProject "Project containing inventory file used as source."
	SourceProject types.Int64 `tfsdk:"source_project" json:"source_project"`
	// SourceVars "Inventory source variables in YAML or JSON format."
	SourceVars types.String `tfsdk:"source_vars" json:"source_vars"`
	// Timeout "The amount of time (in seconds) to run before the task is canceled."
	Timeout types.Int64 `tfsdk:"timeout" json:"timeout"`
	// UpdateCacheTimeout ""
	UpdateCacheTimeout types.Int64 `tfsdk:"update_cache_timeout" json:"update_cache_timeout"`
	// UpdateOnLaunch ""
	UpdateOnLaunch types.Bool `tfsdk:"update_on_launch" json:"update_on_launch"`
	// Verbosity ""
	Verbosity types.String `tfsdk:"verbosity" json:"verbosity"`
}

// Clone the object
func (o *inventorySourceTerraformModel) Clone() inventorySourceTerraformModel {
	return inventorySourceTerraformModel{
		Credential:           o.Credential,
		Description:          o.Description,
		EnabledValue:         o.EnabledValue,
		EnabledVar:           o.EnabledVar,
		ExecutionEnvironment: o.ExecutionEnvironment,
		HostFilter:           o.HostFilter,
		ID:                   o.ID,
		Inventory:            o.Inventory,
		Name:                 o.Name,
		Overwrite:            o.Overwrite,
		OverwriteVars:        o.OverwriteVars,
		Source:               o.Source,
		SourcePath:           o.SourcePath,
		SourceProject:        o.SourceProject,
		SourceVars:           o.SourceVars,
		Timeout:              o.Timeout,
		UpdateCacheTimeout:   o.UpdateCacheTimeout,
		UpdateOnLaunch:       o.UpdateOnLaunch,
		Verbosity:            o.Verbosity,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for InventorySource
func (o *inventorySourceTerraformModel) BodyRequest() (req inventorySourceBodyRequestModel) {
	req.Credential = o.Credential.ValueInt64()
	req.Description = o.Description.ValueString()
	req.EnabledValue = o.EnabledValue.ValueString()
	req.EnabledVar = o.EnabledVar.ValueString()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.HostFilter = o.HostFilter.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Overwrite = o.Overwrite.ValueBool()
	req.OverwriteVars = o.OverwriteVars.ValueBool()
	req.Source = o.Source.ValueString()
	req.SourcePath = o.SourcePath.ValueString()
	req.SourceProject = o.SourceProject.ValueInt64()
	req.SourceVars = json.RawMessage(o.SourceVars.String())
	req.Timeout = o.Timeout.ValueInt64()
	req.UpdateCacheTimeout = o.UpdateCacheTimeout.ValueInt64()
	req.UpdateOnLaunch = o.UpdateOnLaunch.ValueBool()
	req.Verbosity = o.Verbosity.ValueString()
	return
}

func (o *inventorySourceTerraformModel) setCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Credential, data)
}

func (o *inventorySourceTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *inventorySourceTerraformModel) setEnabledValue(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.EnabledValue, data, false)
}

func (o *inventorySourceTerraformModel) setEnabledVar(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.EnabledVar, data, false)
}

func (o *inventorySourceTerraformModel) setExecutionEnvironment(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data)
}

func (o *inventorySourceTerraformModel) setHostFilter(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.HostFilter, data, false)
}

func (o *inventorySourceTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *inventorySourceTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *inventorySourceTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *inventorySourceTerraformModel) setOverwrite(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Overwrite, data)
}

func (o *inventorySourceTerraformModel) setOverwriteVars(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.OverwriteVars, data)
}

func (o *inventorySourceTerraformModel) setSource(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Source, data, false)
}

func (o *inventorySourceTerraformModel) setSourcePath(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SourcePath, data, false)
}

func (o *inventorySourceTerraformModel) setSourceProject(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SourceProject, data)
}

func (o *inventorySourceTerraformModel) setSourceVars(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SourceVars, data, false)
}

func (o *inventorySourceTerraformModel) setTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Timeout, data)
}

func (o *inventorySourceTerraformModel) setUpdateCacheTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.UpdateCacheTimeout, data)
}

func (o *inventorySourceTerraformModel) setUpdateOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.UpdateOnLaunch, data)
}

func (o *inventorySourceTerraformModel) setVerbosity(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Verbosity, data, false)
}

func (o *inventorySourceTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setCredential(data["credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setEnabledValue(data["enabled_value"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setEnabledVar(data["enabled_var"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExecutionEnvironment(data["execution_environment"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHostFilter(data["host_filter"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventory(data["inventory"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOverwrite(data["overwrite"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOverwriteVars(data["overwrite_vars"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSource(data["source"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSourcePath(data["source_path"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSourceProject(data["source_project"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSourceVars(data["source_vars"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTimeout(data["timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUpdateCacheTimeout(data["update_cache_timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUpdateOnLaunch(data["update_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVerbosity(data["verbosity"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// inventorySourceBodyRequestModel maps the schema for InventorySource for creating and updating the data
type inventorySourceBodyRequestModel struct {
	// Credential "Cloud credential to use for inventory updates."
	Credential int64 `json:"credential,omitempty"`
	// Description "Optional description of this inventory source."
	Description string `json:"description,omitempty"`
	// EnabledValue "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled"
	EnabledValue string `json:"enabled_value,omitempty"`
	// EnabledVar "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)"
	EnabledVar string `json:"enabled_var,omitempty"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment int64 `json:"execution_environment,omitempty"`
	// HostFilter "Regex where only matching hosts will be imported."
	HostFilter string `json:"host_filter,omitempty"`
	// Inventory ""
	Inventory int64 `json:"inventory"`
	// Name "Name of this inventory source."
	Name string `json:"name"`
	// Overwrite "Overwrite local groups and hosts from remote inventory source."
	Overwrite bool `json:"overwrite"`
	// OverwriteVars "Overwrite local variables from remote inventory source."
	OverwriteVars bool `json:"overwrite_vars"`
	// Source ""
	Source string `json:"source,omitempty"`
	// SourcePath ""
	SourcePath string `json:"source_path,omitempty"`
	// SourceProject "Project containing inventory file used as source."
	SourceProject int64 `json:"source_project,omitempty"`
	// SourceVars "Inventory source variables in YAML or JSON format."
	SourceVars json.RawMessage `json:"source_vars,omitempty"`
	// Timeout "The amount of time (in seconds) to run before the task is canceled."
	Timeout int64 `json:"timeout,omitempty"`
	// UpdateCacheTimeout ""
	UpdateCacheTimeout int64 `json:"update_cache_timeout,omitempty"`
	// UpdateOnLaunch ""
	UpdateOnLaunch bool `json:"update_on_launch"`
	// Verbosity ""
	Verbosity string `json:"verbosity,omitempty"`
}
