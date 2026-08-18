package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthRadiusTerraformModel struct {
	RADIUS_PORT   types.Int64  `tfsdk:"radius_port" json:"RADIUS_PORT"`
	RADIUS_SECRET types.String `tfsdk:"radius_secret" json:"RADIUS_SECRET"`
	RADIUS_SERVER types.String `tfsdk:"radius_server" json:"RADIUS_SERVER"`
}

func (o *settingsAuthRadiusTerraformModel) Clone() settingsAuthRadiusTerraformModel {
	return *o
}

func (o *settingsAuthRadiusTerraformModel) BodyRequest() *settingsAuthRadiusBodyRequestModel {
	var req settingsAuthRadiusBodyRequestModel
	req.RADIUS_PORT = o.RADIUS_PORT.ValueInt64()
	req.RADIUS_SECRET = o.RADIUS_SECRET.ValueString()
	req.RADIUS_SERVER = o.RADIUS_SERVER.ValueString()
	return &req
}

func (o *settingsAuthRadiusTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.RADIUS_PORT, data["RADIUS_PORT"]))
	collect(helpers.AttrValueSetString(&o.RADIUS_SECRET, data["RADIUS_SECRET"], false))
	collect(helpers.AttrValueSetString(&o.RADIUS_SERVER, data["RADIUS_SERVER"], false))
	return diags, nil
}

type settingsAuthRadiusBodyRequestModel struct {
	RADIUS_PORT   int64  `json:"RADIUS_PORT,omitempty"`
	RADIUS_SECRET string `json:"RADIUS_SECRET,omitempty"`
	RADIUS_SERVER string `json:"RADIUS_SERVER,omitempty"`
}

type settingsAuthRadiusResource = framework.GenericResource[settingsAuthRadiusTerraformModel, settingsAuthRadiusBodyRequestModel, *settingsAuthRadiusTerraformModel]

// NewSettingsAuthRADIUSResource is a helper function to simplify the provider implementation.
func NewSettingsAuthRADIUSResource() resource.Resource {
	return &settingsAuthRadiusResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_radius", Endpoint: "/api/v2/settings/radius/"}},
		Cfg: framework.ResourceCfg[settingsAuthRadiusTerraformModel, settingsAuthRadiusBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"radius_port": schema.Int64Attribute{
						Description: "Port of RADIUS server.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1812),
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
					},
					"radius_secret": schema.StringAttribute{
						Description: "Shared secret for authenticating to RADIUS server.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
					},
					"radius_server": schema.StringAttribute{
						Description: "Hostname/IP of RADIUS server. RADIUS authentication is disabled if this setting is empty.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsAuthRadius,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthRADIUS",
		},
	}
}

type settingsAuthRadiusDataSource = framework.GenericDataSource[settingsAuthRadiusTerraformModel, *settingsAuthRadiusTerraformModel]

// NewSettingsAuthRADIUSDataSource is a helper function to instantiate the SettingsAuthRADIUS data source.
func NewSettingsAuthRADIUSDataSource() datasource.DataSource {
	return &settingsAuthRadiusDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_radius", Endpoint: "/api/v2/settings/radius/"}},
		Cfg: framework.DataSourceCfg[settingsAuthRadiusTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"radius_port": dschema.Int64Attribute{
						Description: "Port of RADIUS server.",
						Computed:    true,
					},
					"radius_secret": dschema.StringAttribute{
						Description: "Shared secret for authenticating to RADIUS server.",
						Sensitive:   true,
						Computed:    true,
					},
					"radius_server": dschema.StringAttribute{
						Description: "Hostname/IP of RADIUS server. RADIUS authentication is disabled if this setting is empty.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthRadius,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthRADIUS",
		},
	}
}
