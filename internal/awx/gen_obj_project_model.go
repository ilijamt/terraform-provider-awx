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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Project
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
	{
		dg, _ := helpers.AttrValueSetBool(&o.AllowOverride, data["allow_override"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Credential, data["credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.DefaultEnvironment, data["default_environment"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LocalPath, data["local_path"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ScmClean, data["scm_clean"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ScmDeleteOnUpdate, data["scm_delete_on_update"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ScmRefspec, data["scm_refspec"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ScmTrackSubmodules, data["scm_track_submodules"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ScmType, data["scm_type"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ScmUpdateCacheTimeout, data["scm_update_cache_timeout"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.ScmUpdateOnLaunch, data["scm_update_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ScmUrl, data["scm_url"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.SignatureValidationCredential, data["signature_validation_credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Timeout, data["timeout"])
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
