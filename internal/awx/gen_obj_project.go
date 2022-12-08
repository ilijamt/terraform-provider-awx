package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// projectTerraformModel maps the schema for Project when using Data Source
type projectTerraformModel struct {
	// AllowOverride "Allow changing the SCM branch or revision in a job template that uses this project."
	AllowOverride types.Bool `tfsdk:"allow_override" json:"allow_override"`
	// Credential ""
	Credential types.Int64 `tfsdk:"credential" json:"credential"`
	// DefaultEnvironment "The default execution environment for jobs run using this project."
	DefaultEnvironment types.Int64 `tfsdk:"default_environment" json:"default_environment"`
	// Description "Optional description of this project."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this project."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// LocalPath "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project."
	LocalPath types.String `tfsdk:"local_path" json:"local_path"`
	// Name "Name of this project."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "The organization used to determine access to this template."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// ScmBranch "Specific branch, tag or commit to checkout."
	ScmBranch types.String `tfsdk:"scm_branch" json:"scm_branch"`
	// ScmClean "Discard any local changes before syncing the project."
	ScmClean types.Bool `tfsdk:"scm_clean" json:"scm_clean"`
	// ScmDeleteOnUpdate "Delete the project before syncing."
	ScmDeleteOnUpdate types.Bool `tfsdk:"scm_delete_on_update" json:"scm_delete_on_update"`
	// ScmRefspec "For git projects, an additional refspec to fetch."
	ScmRefspec types.String `tfsdk:"scm_refspec" json:"scm_refspec"`
	// ScmRevision "The last revision fetched by a project update"
	ScmRevision types.String `tfsdk:"scm_revision" json:"scm_revision"`
	// ScmTrackSubmodules "Track submodules latest commits on defined branch."
	ScmTrackSubmodules types.Bool `tfsdk:"scm_track_submodules" json:"scm_track_submodules"`
	// ScmType "Specifies the source control system used to store the project."
	ScmType types.String `tfsdk:"scm_type" json:"scm_type"`
	// ScmUpdateCacheTimeout "The number of seconds after the last project update ran that a new project update will be launched as a job dependency."
	ScmUpdateCacheTimeout types.Int64 `tfsdk:"scm_update_cache_timeout" json:"scm_update_cache_timeout"`
	// ScmUpdateOnLaunch "Update the project when a job is launched that uses the project."
	ScmUpdateOnLaunch types.Bool `tfsdk:"scm_update_on_launch" json:"scm_update_on_launch"`
	// ScmUrl "The location where the project is stored."
	ScmUrl types.String `tfsdk:"scm_url" json:"scm_url"`
	// SignatureValidationCredential "An optional credential used for validating files in the project against unexpected changes."
	SignatureValidationCredential types.Int64 `tfsdk:"signature_validation_credential" json:"signature_validation_credential"`
	// Timeout "The amount of time (in seconds) to run before the task is canceled."
	Timeout types.Int64 `tfsdk:"timeout" json:"timeout"`
}

// Clone the object
func (o *projectTerraformModel) Clone() projectTerraformModel {
	return projectTerraformModel{
		AllowOverride:                 o.AllowOverride,
		Credential:                    o.Credential,
		DefaultEnvironment:            o.DefaultEnvironment,
		Description:                   o.Description,
		ID:                            o.ID,
		LocalPath:                     o.LocalPath,
		Name:                          o.Name,
		Organization:                  o.Organization,
		ScmBranch:                     o.ScmBranch,
		ScmClean:                      o.ScmClean,
		ScmDeleteOnUpdate:             o.ScmDeleteOnUpdate,
		ScmRefspec:                    o.ScmRefspec,
		ScmRevision:                   o.ScmRevision,
		ScmTrackSubmodules:            o.ScmTrackSubmodules,
		ScmType:                       o.ScmType,
		ScmUpdateCacheTimeout:         o.ScmUpdateCacheTimeout,
		ScmUpdateOnLaunch:             o.ScmUpdateOnLaunch,
		ScmUrl:                        o.ScmUrl,
		SignatureValidationCredential: o.SignatureValidationCredential,
		Timeout:                       o.Timeout,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Project
func (o *projectTerraformModel) BodyRequest() (req projectBodyRequestModel) {
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
	return
}

func (o *projectTerraformModel) setAllowOverride(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AllowOverride, data)
}

func (o *projectTerraformModel) setCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Credential, data)
}

func (o *projectTerraformModel) setDefaultEnvironment(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.DefaultEnvironment, data)
}

func (o *projectTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *projectTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *projectTerraformModel) setLocalPath(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LocalPath, data, false)
}

func (o *projectTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *projectTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *projectTerraformModel) setScmBranch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmBranch, data, false)
}

func (o *projectTerraformModel) setScmClean(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.ScmClean, data)
}

func (o *projectTerraformModel) setScmDeleteOnUpdate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.ScmDeleteOnUpdate, data)
}

func (o *projectTerraformModel) setScmRefspec(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmRefspec, data, false)
}

func (o *projectTerraformModel) setScmRevision(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmRevision, data, false)
}

func (o *projectTerraformModel) setScmTrackSubmodules(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.ScmTrackSubmodules, data)
}

func (o *projectTerraformModel) setScmType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmType, data, false)
}

func (o *projectTerraformModel) setScmUpdateCacheTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ScmUpdateCacheTimeout, data)
}

func (o *projectTerraformModel) setScmUpdateOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.ScmUpdateOnLaunch, data)
}

func (o *projectTerraformModel) setScmUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmUrl, data, false)
}

func (o *projectTerraformModel) setSignatureValidationCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SignatureValidationCredential, data)
}

func (o *projectTerraformModel) setTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Timeout, data)
}

func (o *projectTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAllowOverride(data["allow_override"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCredential(data["credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDefaultEnvironment(data["default_environment"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLocalPath(data["local_path"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmBranch(data["scm_branch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmClean(data["scm_clean"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmDeleteOnUpdate(data["scm_delete_on_update"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmRefspec(data["scm_refspec"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmRevision(data["scm_revision"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmTrackSubmodules(data["scm_track_submodules"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmType(data["scm_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmUpdateCacheTimeout(data["scm_update_cache_timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmUpdateOnLaunch(data["scm_update_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmUrl(data["scm_url"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSignatureValidationCredential(data["signature_validation_credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTimeout(data["timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// projectBodyRequestModel maps the schema for Project for creating and updating the data
type projectBodyRequestModel struct {
	// AllowOverride "Allow changing the SCM branch or revision in a job template that uses this project."
	AllowOverride bool `json:"allow_override"`
	// Credential ""
	Credential int64 `json:"credential,omitempty"`
	// DefaultEnvironment "The default execution environment for jobs run using this project."
	DefaultEnvironment int64 `json:"default_environment,omitempty"`
	// Description "Optional description of this project."
	Description string `json:"description,omitempty"`
	// LocalPath "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project."
	LocalPath string `json:"local_path,omitempty"`
	// Name "Name of this project."
	Name string `json:"name"`
	// Organization "The organization used to determine access to this template."
	Organization int64 `json:"organization,omitempty"`
	// ScmBranch "Specific branch, tag or commit to checkout."
	ScmBranch string `json:"scm_branch,omitempty"`
	// ScmClean "Discard any local changes before syncing the project."
	ScmClean bool `json:"scm_clean"`
	// ScmDeleteOnUpdate "Delete the project before syncing."
	ScmDeleteOnUpdate bool `json:"scm_delete_on_update"`
	// ScmRefspec "For git projects, an additional refspec to fetch."
	ScmRefspec string `json:"scm_refspec,omitempty"`
	// ScmTrackSubmodules "Track submodules latest commits on defined branch."
	ScmTrackSubmodules bool `json:"scm_track_submodules"`
	// ScmType "Specifies the source control system used to store the project."
	ScmType string `json:"scm_type,omitempty"`
	// ScmUpdateCacheTimeout "The number of seconds after the last project update ran that a new project update will be launched as a job dependency."
	ScmUpdateCacheTimeout int64 `json:"scm_update_cache_timeout,omitempty"`
	// ScmUpdateOnLaunch "Update the project when a job is launched that uses the project."
	ScmUpdateOnLaunch bool `json:"scm_update_on_launch"`
	// ScmUrl "The location where the project is stored."
	ScmUrl string `json:"scm_url,omitempty"`
	// SignatureValidationCredential "An optional credential used for validating files in the project against unexpected changes."
	SignatureValidationCredential int64 `json:"signature_validation_credential,omitempty"`
	// Timeout "The amount of time (in seconds) to run before the task is canceled."
	Timeout int64 `json:"timeout,omitempty"`
}

type projectObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}

var (
	_ datasource.DataSource              = &projectDataSource{}
	_ datasource.DataSourceWithConfigure = &projectDataSource{}
)

// NewProjectDataSource is a helper function to instantiate the Project data source.
func NewProjectDataSource() datasource.DataSource {
	return &projectDataSource{}
}

// projectDataSource is the data source implementation.
type projectDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *projectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/projects/"
}

// Metadata returns the data source type name.
func (o *projectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// GetSchema defines the schema for the data source.
func (o *projectDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Project",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"allow_override": {
					Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"credential": {
					Description: "Credential",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"default_environment": {
					Description: "The default execution environment for jobs run using this project.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this project.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this project.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
						),
					},
				},
				"local_path": {
					Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this project.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"organization": {
					Description: "The organization used to determine access to this template.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Specific branch, tag or commit to checkout.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_clean": {
					Description: "Discard any local changes before syncing the project.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_delete_on_update": {
					Description: "Delete the project before syncing.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_refspec": {
					Description: "For git projects, an additional refspec to fetch.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_revision": {
					Description: "The last revision fetched by a project update",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_track_submodules": {
					Description: "Track submodules latest commits on defined branch.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_type": {
					Description: "Specifies the source control system used to store the project.",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "git", "svn", "insights", "archive"}...),
					},
				},
				"scm_update_cache_timeout": {
					Description: "The number of seconds after the last project update ran that a new project update will be launched as a job dependency.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_update_on_launch": {
					Description: "Update the project when a job is launched that uses the project.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_url": {
					Description: "The location where the project is stored.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"signature_validation_credential": {
					Description: "An optional credential used for validating files in the project against unexpected changes.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"timeout": {
					Description: "The amount of time (in seconds) to run before the task is canceled.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *projectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state projectTerraformModel
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

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for Project
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Project on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Project
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Project on %s", o.endpoint),
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
	_ resource.Resource                = &projectResource{}
	_ resource.ResourceWithConfigure   = &projectResource{}
	_ resource.ResourceWithImportState = &projectResource{}
)

// NewProjectResource is a helper function to simplify the provider implementation.
func NewProjectResource() resource.Resource {
	return &projectResource{}
}

type projectResource struct {
	client   c.Client
	endpoint string
}

func (o *projectResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/projects/"
}

func (o *projectResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_project"
}

func (o *projectResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Project",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"allow_override": {
					Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"credential": {
					Description: "Credential",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"default_environment": {
					Description: "The default execution environment for jobs run using this project.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this project.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"local_path": {
					Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"name": {
					Description:   "Name of this project.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"organization": {
					Description: "The organization used to determine access to this template.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Specific branch, tag or commit to checkout.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(256),
					},
				},
				"scm_clean": {
					Description: "Discard any local changes before syncing the project.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_delete_on_update": {
					Description: "Delete the project before syncing.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_refspec": {
					Description: "For git projects, an additional refspec to fetch.",
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
				"scm_track_submodules": {
					Description: "Track submodules latest commits on defined branch.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_type": {
					Description: "Specifies the source control system used to store the project.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "git", "svn", "insights", "archive"}...),
					},
				},
				"scm_update_cache_timeout": {
					Description: "The number of seconds after the last project update ran that a new project update will be launched as a job dependency.",
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
				"scm_update_on_launch": {
					Description: "Update the project when a job is launched that uses the project.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_url": {
					Description: "The location where the project is stored.",
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
				"signature_validation_credential": {
					Description: "An optional credential used for validating files in the project against unexpected changes.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
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
				// Write only elements
				// Data only elements
				"id": {
					Description: "Database ID for this project.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"scm_revision": {
					Description: "The last revision fetched by a project update",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *projectResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Project.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *projectResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state projectTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Project
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[Project/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Project on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Project resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Project on %s", o.endpoint),
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

func (o *projectResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state projectTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Project
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Project on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Project from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Project on %s", o.endpoint),
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

func (o *projectResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state projectTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Project
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[Project/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Project on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Project resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Project on %s", o.endpoint),
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

func (o *projectResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state projectTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Project
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Project on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Project
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Project on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

var (
	_ datasource.DataSource              = &projectObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &projectObjectRolesDataSource{}
)

// NewProjectObjectRolesDataSource is a helper function to instantiate the Project data source.
func NewProjectObjectRolesDataSource() datasource.DataSource {
	return &projectObjectRolesDataSource{}
}

// projectObjectRolesDataSource is the data source implementation.
type projectObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *projectObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/projects/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *projectObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_object_roles"
}

// GetSchema defines the schema for the data source.
func (o *projectObjectRolesDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: helpers.SchemaVersion,
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "Project ID",
				Type:        types.Int64Type,
				Required:    true,
			},
			"roles": {
				Description: "Roles for project",
				Type:        types.MapType{ElemType: types.Int64Type},
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (o *projectObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state projectObjectRolesModel
	var err error
	var id types.Int64

	if d := req.Config.GetAttribute(ctx, path.Root("id"), &id); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	// Creates a new request for Credential
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(o.endpoint, id.ValueInt64()), nil); err != nil {
		resp.Diagnostics.AddError(
			"Unable to create a new request for project",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch the request for project object roles",
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for project",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var d diag.Diagnostics
	if state.Roles, d = types.MapValue(types.Int64Type, in); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
