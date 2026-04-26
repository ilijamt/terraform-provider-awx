package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type inventorySourceTerraformModel struct {
	Credential           types.Int64  `tfsdk:"credential" json:"credential"`
	Description          types.String `tfsdk:"description" json:"description"`
	EnabledValue         types.String `tfsdk:"enabled_value" json:"enabled_value"`
	EnabledVar           types.String `tfsdk:"enabled_var" json:"enabled_var"`
	ExecutionEnvironment types.Int64  `tfsdk:"execution_environment" json:"execution_environment"`
	HostFilter           types.String `tfsdk:"host_filter" json:"host_filter"`
	ID                   types.Int64  `tfsdk:"id" json:"id"`
	Inventory            types.Int64  `tfsdk:"inventory" json:"inventory"`
	Limit                types.String `tfsdk:"limit" json:"limit"`
	Name                 types.String `tfsdk:"name" json:"name"`
	Overwrite            types.Bool   `tfsdk:"overwrite" json:"overwrite"`
	OverwriteVars        types.Bool   `tfsdk:"overwrite_vars" json:"overwrite_vars"`
	ScmBranch            types.String `tfsdk:"scm_branch" json:"scm_branch"`
	Source               types.String `tfsdk:"source" json:"source"`
	SourcePath           types.String `tfsdk:"source_path" json:"source_path"`
	SourceProject        types.Int64  `tfsdk:"source_project" json:"source_project"`
	SourceVars           types.String `tfsdk:"source_vars" json:"source_vars"`
	Timeout              types.Int64  `tfsdk:"timeout" json:"timeout"`
	UpdateCacheTimeout   types.Int64  `tfsdk:"update_cache_timeout" json:"update_cache_timeout"`
	UpdateOnLaunch       types.Bool   `tfsdk:"update_on_launch" json:"update_on_launch"`
	Verbosity            types.String `tfsdk:"verbosity" json:"verbosity"`
}

func (o *inventorySourceTerraformModel) Clone() inventorySourceTerraformModel {
	return *o
}

func (o *inventorySourceTerraformModel) BodyRequest() *inventorySourceBodyRequestModel {
	var req inventorySourceBodyRequestModel
	req.Credential = o.Credential.ValueInt64()
	req.Description = o.Description.ValueString()
	req.EnabledValue = o.EnabledValue.ValueString()
	req.EnabledVar = o.EnabledVar.ValueString()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.HostFilter = o.HostFilter.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Overwrite = o.Overwrite.ValueBool()
	req.OverwriteVars = o.OverwriteVars.ValueBool()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.Source = o.Source.ValueString()
	req.SourcePath = o.SourcePath.ValueString()
	req.SourceProject = o.SourceProject.ValueInt64()
	req.SourceVars = json.RawMessage(o.SourceVars.String())
	req.Timeout = o.Timeout.ValueInt64()
	req.UpdateCacheTimeout = o.UpdateCacheTimeout.ValueInt64()
	req.UpdateOnLaunch = o.UpdateOnLaunch.ValueBool()
	req.Verbosity = o.Verbosity.ValueString()
	return &req
}

func (o *inventorySourceTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetString(&o.EnabledValue, data["enabled_value"], false))
	collect(helpers.AttrValueSetString(&o.EnabledVar, data["enabled_var"], false))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetString(&o.HostFilter, data["host_filter"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetBool(&o.Overwrite, data["overwrite"]))
	collect(helpers.AttrValueSetBool(&o.OverwriteVars, data["overwrite_vars"]))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetString(&o.Source, data["source"], false))
	collect(helpers.AttrValueSetString(&o.SourcePath, data["source_path"], false))
	collect(helpers.AttrValueSetInt64(&o.SourceProject, data["source_project"]))
	collect(helpers.AttrValueSetJsonString(&o.SourceVars, data["source_vars"], false))
	collect(helpers.AttrValueSetInt64(&o.Timeout, data["timeout"]))
	collect(helpers.AttrValueSetInt64(&o.UpdateCacheTimeout, data["update_cache_timeout"]))
	collect(helpers.AttrValueSetBool(&o.UpdateOnLaunch, data["update_on_launch"]))
	collect(helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false))
	return diags, nil
}

type inventorySourceBodyRequestModel struct {
	Credential           int64           `json:"credential,omitempty"`
	Description          string          `json:"description,omitempty"`
	EnabledValue         string          `json:"enabled_value,omitempty"`
	EnabledVar           string          `json:"enabled_var,omitempty"`
	ExecutionEnvironment int64           `json:"execution_environment,omitempty"`
	HostFilter           string          `json:"host_filter,omitempty"`
	Inventory            int64           `json:"inventory"`
	Limit                string          `json:"limit,omitempty"`
	Name                 string          `json:"name"`
	Overwrite            bool            `json:"overwrite"`
	OverwriteVars        bool            `json:"overwrite_vars"`
	ScmBranch            string          `json:"scm_branch,omitempty"`
	Source               string          `json:"source,omitempty"`
	SourcePath           string          `json:"source_path,omitempty"`
	SourceProject        int64           `json:"source_project,omitempty"`
	SourceVars           json.RawMessage `json:"source_vars,omitempty"`
	Timeout              int64           `json:"timeout,omitempty"`
	UpdateCacheTimeout   int64           `json:"update_cache_timeout,omitempty"`
	UpdateOnLaunch       bool            `json:"update_on_launch"`
	Verbosity            string          `json:"verbosity,omitempty"`
}
