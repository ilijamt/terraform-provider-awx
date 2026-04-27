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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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

type adHocCommandTerraformModel struct {
	BecomeEnabled        types.Bool    `tfsdk:"become_enabled" json:"become_enabled"`
	CanceledOn           types.String  `tfsdk:"canceled_on" json:"canceled_on"`
	ControllerNode       types.String  `tfsdk:"controller_node" json:"controller_node"`
	Credential           types.Int64   `tfsdk:"credential" json:"credential"`
	DiffMode             types.Bool    `tfsdk:"diff_mode" json:"diff_mode"`
	Elapsed              types.Float64 `tfsdk:"elapsed" json:"elapsed"`
	ExecutionEnvironment types.Int64   `tfsdk:"execution_environment" json:"execution_environment"`
	ExecutionNode        types.String  `tfsdk:"execution_node" json:"execution_node"`
	ExtraVars            types.String  `tfsdk:"extra_vars" json:"extra_vars"`
	Failed               types.Bool    `tfsdk:"failed" json:"failed"`
	Finished             types.String  `tfsdk:"finished" json:"finished"`
	Forks                types.Int64   `tfsdk:"forks" json:"forks"`
	ID                   types.Int64   `tfsdk:"id" json:"id"`
	Inventory            types.Int64   `tfsdk:"inventory" json:"inventory"`
	JobExplanation       types.String  `tfsdk:"job_explanation" json:"job_explanation"`
	JobType              types.String  `tfsdk:"job_type" json:"job_type"`
	LaunchType           types.String  `tfsdk:"launch_type" json:"launch_type"`
	LaunchedBy           types.Int64   `tfsdk:"launched_by" json:"launched_by"`
	Limit                types.String  `tfsdk:"limit" json:"limit"`
	ModuleArgs           types.String  `tfsdk:"module_args" json:"module_args"`
	ModuleName           types.String  `tfsdk:"module_name" json:"module_name"`
	Name                 types.String  `tfsdk:"name" json:"name"`
	Started              types.String  `tfsdk:"started" json:"started"`
	Status               types.String  `tfsdk:"status" json:"status"`
	Verbosity            types.String  `tfsdk:"verbosity" json:"verbosity"`
	WorkUnitId           types.String  `tfsdk:"work_unit_id" json:"work_unit_id"`
}

func (o *adHocCommandTerraformModel) Clone() adHocCommandTerraformModel {
	return *o
}

func (o *adHocCommandTerraformModel) BodyRequest() *adHocCommandBodyRequestModel {
	var req adHocCommandBodyRequestModel
	req.BecomeEnabled = o.BecomeEnabled.ValueBool()
	req.Credential = o.Credential.ValueInt64()
	req.DiffMode = o.DiffMode.ValueBool()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.ExtraVars = json.RawMessage(o.ExtraVars.String())
	req.Forks = o.Forks.ValueInt64()
	req.Inventory = o.Inventory.ValueInt64()
	req.JobType = o.JobType.ValueString()
	req.Limit = o.Limit.ValueString()
	req.ModuleArgs = o.ModuleArgs.ValueString()
	req.ModuleName = o.ModuleName.ValueString()
	req.Verbosity = o.Verbosity.ValueString()
	return &req
}

func (o *adHocCommandTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.BecomeEnabled, data["become_enabled"]))
	collect(helpers.AttrValueSetString(&o.CanceledOn, data["canceled_on"], false))
	collect(helpers.AttrValueSetString(&o.ControllerNode, data["controller_node"], false))
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"]))
	collect(helpers.AttrValueSetFloat64(&o.Elapsed, data["elapsed"]))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetString(&o.ExecutionNode, data["execution_node"], false))
	collect(helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false))
	collect(helpers.AttrValueSetBool(&o.Failed, data["failed"]))
	collect(helpers.AttrValueSetString(&o.Finished, data["finished"], false))
	collect(helpers.AttrValueSetInt64(&o.Forks, data["forks"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.JobExplanation, data["job_explanation"], false))
	collect(helpers.AttrValueSetString(&o.JobType, data["job_type"], false))
	collect(helpers.AttrValueSetString(&o.LaunchType, data["launch_type"], false))
	collect(helpers.AttrValueSetInt64(&o.LaunchedBy, data["launched_by"]))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.ModuleArgs, data["module_args"], false))
	collect(helpers.AttrValueSetString(&o.ModuleName, data["module_name"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Started, data["started"], false))
	collect(helpers.AttrValueSetString(&o.Status, data["status"], false))
	collect(helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false))
	collect(helpers.AttrValueSetString(&o.WorkUnitId, data["work_unit_id"], false))
	return diags, nil
}

type adHocCommandBodyRequestModel struct {
	BecomeEnabled        bool            `json:"become_enabled"`
	Credential           int64           `json:"credential,omitempty"`
	DiffMode             bool            `json:"diff_mode"`
	ExecutionEnvironment int64           `json:"execution_environment,omitempty"`
	ExtraVars            json.RawMessage `json:"extra_vars,omitempty"`
	Forks                int64           `json:"forks,omitempty"`
	Inventory            int64           `json:"inventory,omitempty"`
	JobType              string          `json:"job_type,omitempty"`
	Limit                string          `json:"limit,omitempty"`
	ModuleArgs           string          `json:"module_args,omitempty"`
	ModuleName           string          `json:"module_name,omitempty"`
	Verbosity            string          `json:"verbosity,omitempty"`
}

type adHocCommandResource = framework.GenericResource[adHocCommandTerraformModel, adHocCommandBodyRequestModel, *adHocCommandTerraformModel]

// NewAdHocCommandResource is a helper function to simplify the provider implementation.
func NewAdHocCommandResource() resource.Resource {
	return &adHocCommandResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "ad_hoc_command", Endpoint: "/api/v2/ad_hoc_commands/"}},
		Cfg: framework.ResourceCfg[adHocCommandTerraformModel, adHocCommandBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"become_enabled": schema.BoolAttribute{
						Description: "Become enabled",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
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
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
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
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`run`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"run",
								"check",
							),
						},
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"module_args": schema.StringAttribute{
						Description: "Module args",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"module_name": schema.StringAttribute{
						Description: "Module name",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`command`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"command",
								"shell",
								"yum",
								"apt",
								"apt_key",
								"apt_repository",
								"apt_rpm",
								"service",
								"group",
								"user",
								"mount",
								"ping",
								"selinux",
								"setup",
								"win_ping",
								"win_service",
								"win_updates",
								"win_group",
								"win_user",
							),
						},
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`0`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"0",
								"1",
								"2",
								"3",
								"4",
								"5",
							),
						},
					},
					"canceled_on": schema.StringAttribute{
						Description: "The date and time when the cancel request was sent.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"controller_node": schema.StringAttribute{
						Description: "The instance that managed the execution environment.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"elapsed": schema.Float64Attribute{
						Description: "Elapsed time in seconds that the job ran.",
						Computed:    true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"execution_node": schema.StringAttribute{
						Description: "The node the job executed on.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"failed": schema.BoolAttribute{
						Description: "Failed",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"finished": schema.StringAttribute{
						Description: "The date and time the job finished execution.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this ad hoc command.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_explanation": schema.StringAttribute{
						Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"launch_type": schema.StringAttribute{
						Description: "Launch type",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"manual",
								"relaunch",
								"callback",
								"scheduled",
								"dependency",
								"workflow",
								"webhook",
								"sync",
								"scm",
							),
						},
					},
					"launched_by": schema.Int64Attribute{
						Description: "Launched by",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this ad hoc command.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"started": schema.StringAttribute{
						Description: "The date and time the job was queued for starting.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"status": schema.StringAttribute{
						Description: "Status",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"new",
								"pending",
								"waiting",
								"running",
								"successful",
								"failed",
								"error",
								"canceled",
							),
						},
					},
					"work_unit_id": schema.StringAttribute{
						Description: "The Receptor work unit ID associated with this job.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *adHocCommandTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *adHocCommandTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "AdHocCommand",
		},
	}
}

type adHocCommandDataSource = framework.GenericDataSource[adHocCommandTerraformModel, *adHocCommandTerraformModel]

// NewAdHocCommandDataSource is a helper function to instantiate the AdHocCommand data source.
func NewAdHocCommandDataSource() datasource.DataSource {
	return &adHocCommandDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "ad_hoc_command", Endpoint: "/api/v2/ad_hoc_commands/"}},
		Cfg: framework.DataSourceCfg[adHocCommandTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"become_enabled": dschema.BoolAttribute{
						Description: "Become enabled",
						Computed:    true,
					},
					"canceled_on": dschema.StringAttribute{
						Description: "The date and time when the cancel request was sent.",
						Computed:    true,
					},
					"controller_node": dschema.StringAttribute{
						Description: "The instance that managed the execution environment.",
						Computed:    true,
					},
					"credential": dschema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"diff_mode": dschema.BoolAttribute{
						Description: "Diff mode",
						Computed:    true,
					},
					"elapsed": dschema.Float64Attribute{
						Description: "Elapsed time in seconds that the job ran.",
						Computed:    true,
					},
					"execution_environment": dschema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"execution_node": dschema.StringAttribute{
						Description: "The node the job executed on.",
						Computed:    true,
					},
					"extra_vars": dschema.StringAttribute{
						Description: "Extra vars",
						Computed:    true,
					},
					"failed": dschema.BoolAttribute{
						Description: "Failed",
						Computed:    true,
					},
					"finished": dschema.StringAttribute{
						Description: "The date and time the job finished execution.",
						Computed:    true,
					},
					"forks": dschema.Int64Attribute{
						Description: "Forks",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this ad hoc command.",
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
					"job_explanation": dschema.StringAttribute{
						Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
						Computed:    true,
					},
					"job_type": dschema.StringAttribute{
						Description: "Job type",
						Computed:    true,
					},
					"launch_type": dschema.StringAttribute{
						Description: "Launch type",
						Computed:    true,
					},
					"launched_by": dschema.Int64Attribute{
						Description: "Launched by",
						Computed:    true,
					},
					"limit": dschema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"module_args": dschema.StringAttribute{
						Description: "Module args",
						Computed:    true,
					},
					"module_name": dschema.StringAttribute{
						Description: "Module name",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this ad hoc command.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"started": dschema.StringAttribute{
						Description: "The date and time the job was queued for starting.",
						Computed:    true,
					},
					"status": dschema.StringAttribute{
						Description: "Status",
						Computed:    true,
					},
					"verbosity": dschema.StringAttribute{
						Description: "Verbosity",
						Computed:    true,
					},
					"work_unit_id": dschema.StringAttribute{
						Description: "The Receptor work unit ID associated with this job.",
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
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *adHocCommandTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "AdHocCommand",
		},
	}
}
