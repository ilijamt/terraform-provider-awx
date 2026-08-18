package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type configTerraformModel struct {
	AnalyticsCollectors types.String `tfsdk:"analytics_collectors" json:"analytics_collectors"`
	AnalyticsStatus     types.String `tfsdk:"analytics_status" json:"analytics_status"`
	BecomeMethods       types.String `tfsdk:"become_methods" json:"become_methods"`
	CustomVirtualenvs   types.List   `tfsdk:"custom_virtualenvs" json:"custom_virtualenvs"`
	Eula                types.String `tfsdk:"eula" json:"eula"`
	LicenseInfo         types.String `tfsdk:"license_info" json:"license_info"`
	ProjectBaseDir      types.String `tfsdk:"project_base_dir" json:"project_base_dir"`
	ProjectLocalPaths   types.List   `tfsdk:"project_local_paths" json:"project_local_paths"`
	TimeZone            types.String `tfsdk:"time_zone" json:"time_zone"`
	UiNext              types.Bool   `tfsdk:"ui_next" json:"ui_next"`
	Version             types.String `tfsdk:"version" json:"version"`
}

func (o *configTerraformModel) Clone() configTerraformModel {
	return *o
}

func (o *configTerraformModel) BodyRequest() *configBodyRequestModel {
	var req configBodyRequestModel
	return &req
}

func (o *configTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetJsonString(&o.AnalyticsCollectors, data["analytics_collectors"], false))
	collect(helpers.AttrValueSetString(&o.AnalyticsStatus, data["analytics_status"], false))
	collect(helpers.AttrValueSetJsonString(&o.BecomeMethods, data["become_methods"], false))
	collect(helpers.AttrValueSetListString(&o.CustomVirtualenvs, data["custom_virtualenvs"], false))
	collect(helpers.AttrValueSetString(&o.Eula, data["eula"], false))
	collect(helpers.AttrValueSetJsonString(&o.LicenseInfo, data["license_info"], false))
	collect(helpers.AttrValueSetString(&o.ProjectBaseDir, data["project_base_dir"], false))
	collect(helpers.AttrValueSetListString(&o.ProjectLocalPaths, data["project_local_paths"], false))
	collect(helpers.AttrValueSetString(&o.TimeZone, data["time_zone"], false))
	collect(helpers.AttrValueSetBool(&o.UiNext, data["ui_next"]))
	collect(helpers.AttrValueSetString(&o.Version, data["version"], false))
	return diags, nil
}

type configBodyRequestModel struct {
}

type configDataSource = framework.GenericDataSource[configTerraformModel, *configTerraformModel]

// NewConfigDataSource is a helper function to instantiate the Config data source.
func NewConfigDataSource() datasource.DataSource {
	return &configDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "config", Endpoint: "/api/v2/config/"}},
		Cfg: framework.DataSourceCfg[configTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"analytics_collectors": dschema.StringAttribute{
						Description: "The analytics collectors registered on this installation, keyed by collector name.",
						Computed:    true,
					},
					"analytics_status": dschema.StringAttribute{
						Description: "Whether analytics gathering is turned on.",
						Computed:    true,
					},
					"become_methods": dschema.StringAttribute{
						Description: "The privilege escalation methods AWX accepts, as value and label pairs.",
						Computed:    true,
					},
					"custom_virtualenvs": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "Paths of the virtual environments left over from before execution environments. AWX treats these as deprecated.",
						Computed:    true,
					},
					"eula": dschema.StringAttribute{
						Description: "The end user license agreement text. Empty on an open installation.",
						Computed:    true,
					},
					"license_info": dschema.StringAttribute{
						Description: "Subscription details, including license type, expiry and the host counts it permits.",
						Computed:    true,
					},
					"project_base_dir": dschema.StringAttribute{
						Description: "Filesystem path the installation stores project checkouts under.",
						Computed:    true,
					},
					"project_local_paths": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "Directories under the project base dir that a manual project can be pointed at.",
						Computed:    true,
					},
					"time_zone": dschema.StringAttribute{
						Description: "The time zone the installation runs in, which schedules are evaluated against.",
						Computed:    true,
					},
					"ui_next": dschema.BoolAttribute{
						Description: "Whether this installation serves the next generation user interface.",
						Computed:    true,
					},
					"version": dschema.StringAttribute{
						Description: "The AWX version this installation runs, which is not necessarily the version the provider was generated against.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Config",
		},
	}
}
