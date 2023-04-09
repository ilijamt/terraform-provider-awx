package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
