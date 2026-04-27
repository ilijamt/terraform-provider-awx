package awx

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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

type projectTerraformModel struct {
	AllowOverride                 types.Bool   `tfsdk:"allow_override" json:"allow_override"`
	Credential                    types.Int64  `tfsdk:"credential" json:"credential"`
	DefaultEnvironment            types.Int64  `tfsdk:"default_environment" json:"default_environment"`
	Description                   types.String `tfsdk:"description" json:"description"`
	ID                            types.Int64  `tfsdk:"id" json:"id"`
	LocalPath                     types.String `tfsdk:"local_path" json:"local_path"`
	Name                          types.String `tfsdk:"name" json:"name"`
	Organization                  types.Int64  `tfsdk:"organization" json:"organization"`
	ScmBranch                     types.String `tfsdk:"scm_branch" json:"scm_branch"`
	ScmClean                      types.Bool   `tfsdk:"scm_clean" json:"scm_clean"`
	ScmDeleteOnUpdate             types.Bool   `tfsdk:"scm_delete_on_update" json:"scm_delete_on_update"`
	ScmRefspec                    types.String `tfsdk:"scm_refspec" json:"scm_refspec"`
	ScmTrackSubmodules            types.Bool   `tfsdk:"scm_track_submodules" json:"scm_track_submodules"`
	ScmType                       types.String `tfsdk:"scm_type" json:"scm_type"`
	ScmUpdateCacheTimeout         types.Int64  `tfsdk:"scm_update_cache_timeout" json:"scm_update_cache_timeout"`
	ScmUpdateOnLaunch             types.Bool   `tfsdk:"scm_update_on_launch" json:"scm_update_on_launch"`
	ScmUrl                        types.String `tfsdk:"scm_url" json:"scm_url"`
	SignatureValidationCredential types.Int64  `tfsdk:"signature_validation_credential" json:"signature_validation_credential"`
	Timeout                       types.Int64  `tfsdk:"timeout" json:"timeout"`
	// WaitForSync is a Terraform-only toggle, not synced to the AWX API.
	WaitForSync types.Bool     `tfsdk:"wait_for_sync" json:"-"`
	Timeouts    timeouts.Value `tfsdk:"timeouts" json:"-"`
}

func (o *projectTerraformModel) Clone() projectTerraformModel {
	return *o
}

func (o *projectTerraformModel) BodyRequest() *projectBodyRequestModel {
	var req projectBodyRequestModel
	req.AllowOverride = o.AllowOverride.ValueBool()
	req.Credential = o.Credential.ValueInt64()
	req.DefaultEnvironment = o.DefaultEnvironment.ValueInt64()
	req.Description = o.Description.ValueString()
	req.LocalPath = o.LocalPath.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.ScmClean = o.ScmClean.ValueBool()
	req.ScmDeleteOnUpdate = o.ScmDeleteOnUpdate.ValueBool()
	req.ScmRefspec = o.ScmRefspec.ValueString()
	req.ScmTrackSubmodules = o.ScmTrackSubmodules.ValueBool()
	req.ScmType = o.ScmType.ValueString()
	req.ScmUpdateCacheTimeout = o.ScmUpdateCacheTimeout.ValueInt64()
	req.ScmUpdateOnLaunch = o.ScmUpdateOnLaunch.ValueBool()
	req.ScmUrl = o.ScmUrl.ValueString()
	req.SignatureValidationCredential = o.SignatureValidationCredential.ValueInt64()
	req.Timeout = o.Timeout.ValueInt64()
	return &req
}

func (o *projectTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.AllowOverride, data["allow_override"]))
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetInt64(&o.DefaultEnvironment, data["default_environment"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.LocalPath, data["local_path"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetBool(&o.ScmClean, data["scm_clean"]))
	collect(helpers.AttrValueSetBool(&o.ScmDeleteOnUpdate, data["scm_delete_on_update"]))
	collect(helpers.AttrValueSetString(&o.ScmRefspec, data["scm_refspec"], false))
	collect(helpers.AttrValueSetBool(&o.ScmTrackSubmodules, data["scm_track_submodules"]))
	collect(helpers.AttrValueSetString(&o.ScmType, data["scm_type"], false))
	collect(helpers.AttrValueSetInt64(&o.ScmUpdateCacheTimeout, data["scm_update_cache_timeout"]))
	collect(helpers.AttrValueSetBool(&o.ScmUpdateOnLaunch, data["scm_update_on_launch"]))
	collect(helpers.AttrValueSetString(&o.ScmUrl, data["scm_url"], false))
	collect(helpers.AttrValueSetInt64(&o.SignatureValidationCredential, data["signature_validation_credential"]))
	collect(helpers.AttrValueSetInt64(&o.Timeout, data["timeout"]))
	return diags, nil
}

type projectBodyRequestModel struct {
	AllowOverride                 bool   `json:"allow_override"`
	Credential                    int64  `json:"credential,omitempty"`
	DefaultEnvironment            int64  `json:"default_environment,omitempty"`
	Description                   string `json:"description,omitempty"`
	LocalPath                     string `json:"local_path,omitempty"`
	Name                          string `json:"name"`
	Organization                  int64  `json:"organization,omitempty"`
	ScmBranch                     string `json:"scm_branch,omitempty"`
	ScmClean                      bool   `json:"scm_clean"`
	ScmDeleteOnUpdate             bool   `json:"scm_delete_on_update"`
	ScmRefspec                    string `json:"scm_refspec,omitempty"`
	ScmTrackSubmodules            bool   `json:"scm_track_submodules"`
	ScmType                       string `json:"scm_type,omitempty"`
	ScmUpdateCacheTimeout         int64  `json:"scm_update_cache_timeout,omitempty"`
	ScmUpdateOnLaunch             bool   `json:"scm_update_on_launch"`
	ScmUrl                        string `json:"scm_url,omitempty"`
	SignatureValidationCredential int64  `json:"signature_validation_credential,omitempty"`
	Timeout                       int64  `json:"timeout,omitempty"`
}

type projectResource = framework.GenericResource[projectTerraformModel, projectBodyRequestModel, *projectTerraformModel]

// NewProjectResource is a helper function to simplify the provider implementation.
func NewProjectResource() resource.Resource {
	return &projectResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "project", Endpoint: "/api/v2/projects/"}},
		Cfg: framework.ResourceCfg[projectTerraformModel, projectBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"allow_override": schema.BoolAttribute{
						Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
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
					"default_environment": schema.Int64Attribute{
						Description: "The default execution environment for jobs run using this project.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this project.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"local_path": schema.StringAttribute{
						Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this project.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"scm_branch": schema.StringAttribute{
						Description: "Specific branch, tag or commit to checkout.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(256),
						},
					},
					"scm_clean": schema.BoolAttribute{
						Description: "Discard any local changes before syncing the project.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"scm_delete_on_update": schema.BoolAttribute{
						Description: "Delete the project before syncing.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"scm_refspec": schema.StringAttribute{
						Description: "For git projects, an additional refspec to fetch.",
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
					"scm_track_submodules": schema.BoolAttribute{
						Description: "Track submodules latest commits on defined branch.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"scm_type": schema.StringAttribute{
						Description: "Specifies the source control system used to store the project.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"git",
								"svn",
								"insights",
								"archive",
							),
						},
					},
					"scm_update_cache_timeout": schema.Int64Attribute{
						Description: "The number of seconds after the last project update ran that a new project update will be launched as a job dependency.",
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
					"scm_update_on_launch": schema.BoolAttribute{
						Description: "Update the project when a job is launched that uses the project.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"scm_url": schema.StringAttribute{
						Description: "The location where the project is stored.",
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
					"signature_validation_credential": schema.Int64Attribute{
						Description: "An optional credential used for validating files in the project against unexpected changes.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
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
					"id": schema.Int64Attribute{
						Description: "Database ID for this project.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"wait_for_sync": schema.BoolAttribute{
						Description: "If true, wait for AWX to finish the SCM update kicked off on create or update before returning. Configure the maximum wait via the timeouts block.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *projectTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			EmitTimeouts: true,
			CopyExtraAttributes: func(plan, state *projectTerraformModel) {
				state.WaitForSync = plan.WaitForSync
				state.Timeouts = plan.Timeouts
			},
			WaitLifecycle: &framework.WaitLifecycleCfg[projectTerraformModel]{
				ShouldWait: func(plan *projectTerraformModel) bool {
					return !plan.WaitForSync.IsNull() && plan.WaitForSync.ValueBool()
				},
				EndpointForModel: func(m *projectTerraformModel) string {
					if m.ID.IsNull() || m.ID.IsUnknown() {
						return ""
					}
					if m.ID.ValueInt64() == 0 {
						return ""
					}
					return framework.EndpointWithID("/api/v2/projects/", m.ID.ValueInt64())
				},
				Field:          "status",
				SuccessValues:  []string{"successful", "ok", "never updated"},
				FailureValues:  []string{"failed", "error", "canceled"},
				PollInterval:   5 * time.Second,
				DefaultTimeout: 5 * time.Minute,
				ResolveTimeout: func(ctx context.Context, plan *projectTerraformModel, callee hooks.Callee) (time.Duration, diag.Diagnostics) {
					if callee == hooks.CalleeUpdate {
						return plan.Timeouts.Update(ctx, 5*time.Minute)
					}
					return plan.Timeouts.Create(ctx, 5*time.Minute)
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Project",
		},
	}
}

type projectDataSource = framework.GenericDataSource[projectTerraformModel, *projectTerraformModel]

// NewProjectDataSource is a helper function to instantiate the Project data source.
func NewProjectDataSource() datasource.DataSource {
	return &projectDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "project", Endpoint: "/api/v2/projects/"}},
		Cfg: framework.DataSourceCfg[projectTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"allow_override": dschema.BoolAttribute{
						Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
						Computed:    true,
					},
					"credential": dschema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"default_environment": dschema.Int64Attribute{
						Description: "The default execution environment for jobs run using this project.",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this project.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this project.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"local_path": dschema.StringAttribute{
						Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this project.",
						Computed:    true,
					},
					"organization": dschema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Computed:    true,
					},
					"scm_branch": dschema.StringAttribute{
						Description: "Specific branch, tag or commit to checkout.",
						Computed:    true,
					},
					"scm_clean": dschema.BoolAttribute{
						Description: "Discard any local changes before syncing the project.",
						Computed:    true,
					},
					"scm_delete_on_update": dschema.BoolAttribute{
						Description: "Delete the project before syncing.",
						Computed:    true,
					},
					"scm_refspec": dschema.StringAttribute{
						Description: "For git projects, an additional refspec to fetch.",
						Computed:    true,
					},
					"scm_track_submodules": dschema.BoolAttribute{
						Description: "Track submodules latest commits on defined branch.",
						Computed:    true,
					},
					"scm_type": dschema.StringAttribute{
						Description: "Specifies the source control system used to store the project.",
						Computed:    true,
					},
					"scm_update_cache_timeout": dschema.Int64Attribute{
						Description: "The number of seconds after the last project update ran that a new project update will be launched as a job dependency.",
						Computed:    true,
					},
					"scm_update_on_launch": dschema.BoolAttribute{
						Description: "Update the project when a job is launched that uses the project.",
						Computed:    true,
					},
					"scm_url": dschema.StringAttribute{
						Description: "The location where the project is stored.",
						Computed:    true,
					},
					"signature_validation_credential": dschema.Int64Attribute{
						Description: "An optional credential used for validating files in the project against unexpected changes.",
						Computed:    true,
					},
					"timeout": dschema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Project",
		},
	}
}
