package awx

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type projectResource = framework.GenericResource[projectTerraformModel, projectBodyRequestModel, *projectTerraformModel]

// NewProjectResource is a helper function to simplify the provider implementation.
func NewProjectResource() resource.Resource {
	return &projectResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "project", Endpoint: "/api/v2/projects/"}},
		Cfg: framework.ResourceCfg[projectTerraformModel, projectBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"allow_override": schema.BoolAttribute{
						Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"default_environment": schema.Int64Attribute{
						Description: "The default execution environment for jobs run using this project.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this project.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"local_path": schema.StringAttribute{
						Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
						Sensitive:   false,
						Required:    false,
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
						Description:   "Name of this project.",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"scm_branch": schema.StringAttribute{
						Description: "Specific branch, tag or commit to checkout.",
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"scm_delete_on_update": schema.BoolAttribute{
						Description: "Delete the project before syncing.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"scm_refspec": schema.StringAttribute{
						Description: "For git projects, an additional refspec to fetch.",
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"scm_type": schema.StringAttribute{
						Description: "Specifies the source control system used to store the project.",
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"scm_url": schema.StringAttribute{
						Description: "The location where the project is stored.",
						Sensitive:   false,
						Required:    false,
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
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Sensitive:   false,
						Required:    false,
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
					// Write only elements
					// Data only elements
					"id": schema.Int64Attribute{
						Description: "Database ID for this project.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					// Terraform-only lifecycle controls
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
