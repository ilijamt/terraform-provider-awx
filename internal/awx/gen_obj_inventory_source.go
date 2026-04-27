package awx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
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

type inventorySourceResource = framework.GenericResource[inventorySourceTerraformModel, inventorySourceBodyRequestModel, *inventorySourceTerraformModel]

// NewInventorySourceResource is a helper function to simplify the provider implementation.
func NewInventorySourceResource() resource.Resource {
	return &inventorySourceResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "inventory_source", Endpoint: "/api/v2/inventory_sources/"}},
		Cfg: framework.ResourceCfg[inventorySourceTerraformModel, inventorySourceBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential": schema.Int64Attribute{
						Description: "Cloud credential to use for inventory updates.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this inventory source.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"enabled_value": schema.StringAttribute{
						Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"enabled_var": schema.StringAttribute{
						Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"host_filter": schema.StringAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Regex where only matching hosts will be imported.",
						Optional:           true,
						Computed:           true,
						Default:            stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Required:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Enter host, group or pattern match",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this inventory source.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"overwrite": schema.BoolAttribute{
						Description: "Overwrite local groups and hosts from remote inventory source.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"overwrite_vars": schema.BoolAttribute{
						Description: "Overwrite local variables from remote inventory source.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"scm_branch": schema.StringAttribute{
						Description: "Inventory source SCM branch. Project default used if blank. Only allowed if project allow_override field is set to true.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"source": schema.StringAttribute{
						Description: "Source",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"file",
								"constructed",
								"scm",
								"ec2",
								"gce",
								"azure_rm",
								"vmware",
								"satellite6",
								"openstack",
								"rhv",
								"controller",
								"insights",
								"terraform",
								"openshift_virtualization",
							),
						},
					},
					"source_path": schema.StringAttribute{
						Description: "Source path",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"source_project": schema.Int64Attribute{
						Description: "Project containing inventory file used as source.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"source_vars": schema.StringAttribute{
						Description: "Inventory source variables in YAML or JSON format.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(-2147483648, 2147483647),
						},
					},
					"update_cache_timeout": schema.Int64Attribute{
						Description: "Update cache timeout",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 2147483647),
						},
					},
					"update_on_launch": schema.BoolAttribute{
						Description: "Update on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`1`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"0",
								"1",
								"2",
							),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this inventory source.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *inventorySourceTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *inventorySourceTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "InventorySource",
		},
	}
}

type inventorySourceDataSource = framework.GenericDataSource[inventorySourceTerraformModel, *inventorySourceTerraformModel]

// NewInventorySourceDataSource is a helper function to instantiate the InventorySource data source.
func NewInventorySourceDataSource() datasource.DataSource {
	return &inventorySourceDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "inventory_source", Endpoint: "/api/v2/inventory_sources/"}},
		Cfg: framework.DataSourceCfg[inventorySourceTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"credential": dschema.Int64Attribute{
						Description: "Cloud credential to use for inventory updates.",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this inventory source.",
						Computed:    true,
					},
					"enabled_value": dschema.StringAttribute{
						Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
						Computed:    true,
					},
					"enabled_var": dschema.StringAttribute{
						Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
						Computed:    true,
					},
					"execution_environment": dschema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"host_filter": dschema.StringAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Regex where only matching hosts will be imported.",
						Computed:           true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this inventory source.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"inventory": dschema.Int64Attribute{
						Description: "Inventory",
						Computed:    true,
					},
					"limit": dschema.StringAttribute{
						Description: "Enter host, group or pattern match",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this inventory source.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"overwrite": dschema.BoolAttribute{
						Description: "Overwrite local groups and hosts from remote inventory source.",
						Computed:    true,
					},
					"overwrite_vars": dschema.BoolAttribute{
						Description: "Overwrite local variables from remote inventory source.",
						Computed:    true,
					},
					"scm_branch": dschema.StringAttribute{
						Description: "Inventory source SCM branch. Project default used if blank. Only allowed if project allow_override field is set to true.",
						Computed:    true,
					},
					"source": dschema.StringAttribute{
						Description: "Source",
						Computed:    true,
					},
					"source_path": dschema.StringAttribute{
						Description: "Source path",
						Computed:    true,
					},
					"source_project": dschema.Int64Attribute{
						Description: "Project containing inventory file used as source.",
						Computed:    true,
					},
					"source_vars": dschema.StringAttribute{
						Description: "Inventory source variables in YAML or JSON format.",
						Computed:    true,
					},
					"timeout": dschema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Computed:    true,
					},
					"update_cache_timeout": dschema.Int64Attribute{
						Description: "Update cache timeout",
						Computed:    true,
					},
					"update_on_launch": dschema.BoolAttribute{
						Description: "Update on launch",
						Computed:    true,
					},
					"verbosity": dschema.StringAttribute{
						Description: "Verbosity",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *inventorySourceTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "InventorySource",
		},
	}
}
