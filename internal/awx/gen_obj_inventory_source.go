package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
func (o inventorySourceTerraformModel) Clone() inventorySourceTerraformModel {
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
func (o inventorySourceTerraformModel) BodyRequest() (req inventorySourceBodyRequestModel) {
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
	req.SourceVars = json.RawMessage(o.SourceVars.ValueString())
	req.Timeout = o.Timeout.ValueInt64()
	req.UpdateCacheTimeout = o.UpdateCacheTimeout.ValueInt64()
	req.UpdateOnLaunch = o.UpdateOnLaunch.ValueBool()
	req.Verbosity = o.Verbosity.ValueString()
	return
}

func (o *inventorySourceTerraformModel) setCredential(data any) (d diag.Diagnostics, err error) {
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

func (o *inventorySourceTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	// Decode "description"
	if val, ok := data.(string); ok {
		o.Description = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Description = types.StringValue(val.String())
	} else {
		o.Description = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setEnabledValue(data any) (d diag.Diagnostics, err error) {
	// Decode "enabled_value"
	if val, ok := data.(string); ok {
		o.EnabledValue = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.EnabledValue = types.StringValue(val.String())
	} else {
		o.EnabledValue = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setEnabledVar(data any) (d diag.Diagnostics, err error) {
	// Decode "enabled_var"
	if val, ok := data.(string); ok {
		o.EnabledVar = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.EnabledVar = types.StringValue(val.String())
	} else {
		o.EnabledVar = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setExecutionEnvironment(data any) (d diag.Diagnostics, err error) {
	// Decode "execution_environment"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.ExecutionEnvironment = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.ExecutionEnvironment = types.Int64Value(val)
	} else {
		o.ExecutionEnvironment = types.Int64Null()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setHostFilter(data any) (d diag.Diagnostics, err error) {
	// Decode "host_filter"
	if val, ok := data.(string); ok {
		o.HostFilter = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.HostFilter = types.StringValue(val.String())
	} else {
		o.HostFilter = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
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

func (o *inventorySourceTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	// Decode "inventory"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Inventory = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Inventory = types.Int64Value(val)
	} else {
		o.Inventory = types.Int64Null()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
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

func (o *inventorySourceTerraformModel) setOverwrite(data any) (d diag.Diagnostics, err error) {
	// Decode "overwrite"
	if val, ok := data.(bool); ok {
		o.Overwrite = types.BoolValue(val)
	} else {
		o.Overwrite = types.BoolNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setOverwriteVars(data any) (d diag.Diagnostics, err error) {
	// Decode "overwrite_vars"
	if val, ok := data.(bool); ok {
		o.OverwriteVars = types.BoolValue(val)
	} else {
		o.OverwriteVars = types.BoolNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setSource(data any) (d diag.Diagnostics, err error) {
	// Decode "source"
	if val, ok := data.(string); ok {
		o.Source = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Source = types.StringValue(val.String())
	} else {
		o.Source = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setSourcePath(data any) (d diag.Diagnostics, err error) {
	// Decode "source_path"
	if val, ok := data.(string); ok {
		o.SourcePath = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.SourcePath = types.StringValue(val.String())
	} else {
		o.SourcePath = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setSourceProject(data any) (d diag.Diagnostics, err error) {
	// Decode "source_project"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.SourceProject = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.SourceProject = types.Int64Value(val)
	} else {
		o.SourceProject = types.Int64Null()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setSourceVars(data any) (d diag.Diagnostics, err error) {
	// Decode "source_vars"
	if val, ok := data.(string); ok {
		o.SourceVars = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.SourceVars = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.SourceVars = types.StringNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setTimeout(data any) (d diag.Diagnostics, err error) {
	// Decode "timeout"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Timeout = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Timeout = types.Int64Value(val)
	} else {
		o.Timeout = types.Int64Null()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setUpdateCacheTimeout(data any) (d diag.Diagnostics, err error) {
	// Decode "update_cache_timeout"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.UpdateCacheTimeout = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.UpdateCacheTimeout = types.Int64Value(val)
	} else {
		o.UpdateCacheTimeout = types.Int64Null()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setUpdateOnLaunch(data any) (d diag.Diagnostics, err error) {
	// Decode "update_on_launch"
	if val, ok := data.(bool); ok {
		o.UpdateOnLaunch = types.BoolValue(val)
	} else {
		o.UpdateOnLaunch = types.BoolNull()
	}
	return d, nil
}

func (o *inventorySourceTerraformModel) setVerbosity(data any) (d diag.Diagnostics, err error) {
	// Decode "verbosity"
	if val, ok := data.(string); ok {
		o.Verbosity = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Verbosity = types.StringValue(val.String())
	} else {
		o.Verbosity = types.StringNull()
	}
	return d, nil
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

var (
	_ datasource.DataSource              = &inventorySourceDataSource{}
	_ datasource.DataSourceWithConfigure = &inventorySourceDataSource{}
)

// NewInventorySourceDataSource is a helper function to instantiate the InventorySource data source.
func NewInventorySourceDataSource() datasource.DataSource {
	return &inventorySourceDataSource{}
}

// inventorySourceDataSource is the data source implementation.
type inventorySourceDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *inventorySourceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventory_sources/"
}

// Metadata returns the data source type name.
func (o *inventorySourceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_source"
}

// GetSchema defines the schema for the data source.
func (o *inventorySourceDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"InventorySource",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"credential": {
					Description: "Cloud credential to use for inventory updates.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this inventory source.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"enabled_value": {
					Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"enabled_var": {
					Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"execution_environment": {
					Description: "The container image to be used for execution.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"host_filter": {
					Description: "Regex where only matching hosts will be imported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this inventory source.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
				"inventory": {
					Description: "Inventory",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this inventory source.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
				"overwrite": {
					Description: "Overwrite local groups and hosts from remote inventory source.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"overwrite_vars": {
					Description: "Overwrite local variables from remote inventory source.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"source": {
					Description: "Source",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"file", "scm", "ec2", "gce", "azure_rm", "vmware", "satellite6", "openstack", "rhv", "controller", "insights"}...),
					},
				},
				"source_path": {
					Description: "Source path",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"source_project": {
					Description: "Project containing inventory file used as source.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"source_vars": {
					Description: "Inventory source variables in YAML or JSON format.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"timeout": {
					Description: "The amount of time (in seconds) to run before the task is canceled.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"update_cache_timeout": {
					Description: "Update cache timeout",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"update_on_launch": {
					Description: "Update on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"verbosity": {
					Description: "Verbosity",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"0", "1", "2"}...),
					},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *inventorySourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state inventorySourceTerraformModel
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

	// Evaluate group 'by_name' based on the schema definition
	var groupByNameExists = func() bool {
		var groupByNameExists = true
		var paramsByName = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameExists = groupByNameExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByName = append(paramsByName, url.PathEscape(attrName.ValueString()))
		if groupByNameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s", paramsByName...))
		}
		return groupByNameExists
	}()
	searchDefined = searchDefined || groupByNameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			fmt.Sprintf("missing configuration for one of the predefined search groups"),
			detailMessage,
		)
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for InventorySource
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InventorySource on %s", o.endpoint),
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
	_ resource.Resource                = &inventorySourceResource{}
	_ resource.ResourceWithConfigure   = &inventorySourceResource{}
	_ resource.ResourceWithImportState = &inventorySourceResource{}
)

// NewInventorySourceResource is a helper function to simplify the provider implementation.
func NewInventorySourceResource() resource.Resource {
	return &inventorySourceResource{}
}

type inventorySourceResource struct {
	client   c.Client
	endpoint string
}

func (o *inventorySourceResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventory_sources/"
}

func (o inventorySourceResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_inventory_source"
}

func (o inventorySourceResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"InventorySource",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"credential": {
					Description: "Cloud credential to use for inventory updates.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this inventory source.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"enabled_value": {
					Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"enabled_var": {
					Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"execution_environment": {
					Description: "The container image to be used for execution.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"host_filter": {
					Description: "Regex where only matching hosts will be imported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"inventory": {
					Description:   "Inventory",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this inventory source.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"overwrite": {
					Description: "Overwrite local groups and hosts from remote inventory source.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"overwrite_vars": {
					Description: "Overwrite local variables from remote inventory source.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"source": {
					Description: "Source",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"file", "scm", "ec2", "gce", "azure_rm", "vmware", "satellite6", "openstack", "rhv", "controller", "insights"}...),
					},
				},
				"source_path": {
					Description: "Source path",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"source_project": {
					Description: "Project containing inventory file used as source.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"source_vars": {
					Description: "Inventory source variables in YAML or JSON format.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"timeout": {
					Description: "The amount of time (in seconds) to run before the task is canceled.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						int64validator.Between(-2.147483648e+09, 2.147483647e+09),
					},
				},
				"update_cache_timeout": {
					Description: "Update cache timeout",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						int64validator.Between(0, 2.147483647e+09),
					},
				},
				"update_on_launch": {
					Description: "Update on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"verbosity": {
					Description: "Verbosity",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`1`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"0", "1", "2"}...),
					},
				},
				// Write only elements
				// Data only elements
				"id": {
					Description: "Database ID for this inventory source.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *inventorySourceResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the InventorySource.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *inventorySourceResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state inventorySourceTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new InventorySource resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for InventorySource on %s", o.endpoint),
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

func (o *inventorySourceResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state inventorySourceTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for InventorySource from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InventorySource on %s", o.endpoint),
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

func (o *inventorySourceResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state inventorySourceTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new InventorySource resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for InventorySource on %s", o.endpoint),
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

func (o *inventorySourceResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state inventorySourceTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing InventorySource
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for InventorySource on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
