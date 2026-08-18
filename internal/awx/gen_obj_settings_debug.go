package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsDebugTerraformModel struct {
	AWX_CLEANUP_PATHS     types.Bool `tfsdk:"awx_cleanup_paths" json:"AWX_CLEANUP_PATHS"`
	AWX_REQUEST_PROFILE   types.Bool `tfsdk:"awx_request_profile" json:"AWX_REQUEST_PROFILE"`
	RECEPTOR_RELEASE_WORK types.Bool `tfsdk:"receptor_release_work" json:"RECEPTOR_RELEASE_WORK"`
}

func (o *settingsDebugTerraformModel) Clone() settingsDebugTerraformModel {
	return *o
}

func (o *settingsDebugTerraformModel) BodyRequest() *settingsDebugBodyRequestModel {
	var req settingsDebugBodyRequestModel
	req.AWX_CLEANUP_PATHS = o.AWX_CLEANUP_PATHS.ValueBool()
	req.AWX_REQUEST_PROFILE = o.AWX_REQUEST_PROFILE.ValueBool()
	req.RECEPTOR_RELEASE_WORK = o.RECEPTOR_RELEASE_WORK.ValueBool()
	return &req
}

func (o *settingsDebugTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.AWX_CLEANUP_PATHS, data["AWX_CLEANUP_PATHS"]))
	collect(helpers.AttrValueSetBool(&o.AWX_REQUEST_PROFILE, data["AWX_REQUEST_PROFILE"]))
	collect(helpers.AttrValueSetBool(&o.RECEPTOR_RELEASE_WORK, data["RECEPTOR_RELEASE_WORK"]))
	return diags, nil
}

type settingsDebugBodyRequestModel struct {
	AWX_CLEANUP_PATHS     bool `json:"AWX_CLEANUP_PATHS"`
	AWX_REQUEST_PROFILE   bool `json:"AWX_REQUEST_PROFILE"`
	RECEPTOR_RELEASE_WORK bool `json:"RECEPTOR_RELEASE_WORK"`
}

type settingsDebugResource = framework.GenericResource[settingsDebugTerraformModel, settingsDebugBodyRequestModel, *settingsDebugTerraformModel]

// NewSettingsDebugResource is a helper function to simplify the provider implementation.
func NewSettingsDebugResource() resource.Resource {
	return &settingsDebugResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_debug", Endpoint: "/api/v2/settings/debug/"}},
		Cfg: framework.ResourceCfg[settingsDebugTerraformModel, settingsDebugBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"awx_cleanup_paths": schema.BoolAttribute{
						Description: "Enable or Disable TMP Dir cleanup",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(true),
					},
					"awx_request_profile": schema.BoolAttribute{
						Description: "Debug web request python timing",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"receptor_release_work": schema.BoolAttribute{
						Description: "Release receptor work",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(true),
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsDebug",
		},
	}
}

type settingsDebugDataSource = framework.GenericDataSource[settingsDebugTerraformModel, *settingsDebugTerraformModel]

// NewSettingsDebugDataSource is a helper function to instantiate the SettingsDebug data source.
func NewSettingsDebugDataSource() datasource.DataSource {
	return &settingsDebugDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_debug", Endpoint: "/api/v2/settings/debug/"}},
		Cfg: framework.DataSourceCfg[settingsDebugTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"awx_cleanup_paths": dschema.BoolAttribute{
						Description: "Enable or Disable TMP Dir cleanup",
						Computed:    true,
					},
					"awx_request_profile": dschema.BoolAttribute{
						Description: "Debug web request python timing",
						Computed:    true,
					},
					"receptor_release_work": dschema.BoolAttribute{
						Description: "Release receptor work",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsDebug",
		},
	}
}
