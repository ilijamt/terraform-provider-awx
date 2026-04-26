package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
