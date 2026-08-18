package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthTacacsplusTerraformModel struct {
	TACACSPLUS_AUTH_PROTOCOL   types.String `tfsdk:"tacacsplus_auth_protocol" json:"TACACSPLUS_AUTH_PROTOCOL"`
	TACACSPLUS_HOST            types.String `tfsdk:"tacacsplus_host" json:"TACACSPLUS_HOST"`
	TACACSPLUS_PORT            types.Int64  `tfsdk:"tacacsplus_port" json:"TACACSPLUS_PORT"`
	TACACSPLUS_REM_ADDR        types.Bool   `tfsdk:"tacacsplus_rem_addr" json:"TACACSPLUS_REM_ADDR"`
	TACACSPLUS_SECRET          types.String `tfsdk:"tacacsplus_secret" json:"TACACSPLUS_SECRET"`
	TACACSPLUS_SESSION_TIMEOUT types.Int64  `tfsdk:"tacacsplus_session_timeout" json:"TACACSPLUS_SESSION_TIMEOUT"`
}

func (o *settingsAuthTacacsplusTerraformModel) Clone() settingsAuthTacacsplusTerraformModel {
	return *o
}

func (o *settingsAuthTacacsplusTerraformModel) BodyRequest() *settingsAuthTacacsplusBodyRequestModel {
	var req settingsAuthTacacsplusBodyRequestModel
	req.TACACSPLUS_AUTH_PROTOCOL = o.TACACSPLUS_AUTH_PROTOCOL.ValueString()
	req.TACACSPLUS_HOST = o.TACACSPLUS_HOST.ValueString()
	req.TACACSPLUS_PORT = o.TACACSPLUS_PORT.ValueInt64()
	req.TACACSPLUS_REM_ADDR = o.TACACSPLUS_REM_ADDR.ValueBool()
	req.TACACSPLUS_SECRET = o.TACACSPLUS_SECRET.ValueString()
	req.TACACSPLUS_SESSION_TIMEOUT = o.TACACSPLUS_SESSION_TIMEOUT.ValueInt64()
	return &req
}

func (o *settingsAuthTacacsplusTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.TACACSPLUS_AUTH_PROTOCOL, data["TACACSPLUS_AUTH_PROTOCOL"], false))
	collect(helpers.AttrValueSetString(&o.TACACSPLUS_HOST, data["TACACSPLUS_HOST"], false))
	collect(helpers.AttrValueSetInt64(&o.TACACSPLUS_PORT, data["TACACSPLUS_PORT"]))
	collect(helpers.AttrValueSetBool(&o.TACACSPLUS_REM_ADDR, data["TACACSPLUS_REM_ADDR"]))
	collect(helpers.AttrValueSetString(&o.TACACSPLUS_SECRET, data["TACACSPLUS_SECRET"], false))
	collect(helpers.AttrValueSetInt64(&o.TACACSPLUS_SESSION_TIMEOUT, data["TACACSPLUS_SESSION_TIMEOUT"]))
	return diags, nil
}

type settingsAuthTacacsplusBodyRequestModel struct {
	TACACSPLUS_AUTH_PROTOCOL   string `json:"TACACSPLUS_AUTH_PROTOCOL,omitempty"`
	TACACSPLUS_HOST            string `json:"TACACSPLUS_HOST,omitempty"`
	TACACSPLUS_PORT            int64  `json:"TACACSPLUS_PORT,omitempty"`
	TACACSPLUS_REM_ADDR        bool   `json:"TACACSPLUS_REM_ADDR"`
	TACACSPLUS_SECRET          string `json:"TACACSPLUS_SECRET,omitempty"`
	TACACSPLUS_SESSION_TIMEOUT int64  `json:"TACACSPLUS_SESSION_TIMEOUT"`
}

type settingsAuthTacacsplusResource = framework.GenericResource[settingsAuthTacacsplusTerraformModel, settingsAuthTacacsplusBodyRequestModel, *settingsAuthTacacsplusTerraformModel]

// NewSettingsAuthTACACSPlusResource is a helper function to simplify the provider implementation.
func NewSettingsAuthTACACSPlusResource() resource.Resource {
	return &settingsAuthTacacsplusResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_tacacsplus", Endpoint: "/api/v2/settings/tacacsplus/"}},
		Cfg: framework.ResourceCfg[settingsAuthTacacsplusTerraformModel, settingsAuthTacacsplusBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"tacacsplus_auth_protocol": schema.StringAttribute{
						Description: "Choose the authentication protocol used by TACACS+ client.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`ascii`),
						Validators: []validator.String{
							stringvalidator.OneOf(
								"ascii",
								"pap",
							),
						},
					},
					"tacacsplus_host": schema.StringAttribute{
						Description: "Hostname of TACACS+ server.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
					},
					"tacacsplus_port": schema.Int64Attribute{
						Description: "Port number of TACACS+ server.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(49),
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
					},
					"tacacsplus_rem_addr": schema.BoolAttribute{
						Description: "Enable the client address sending by TACACS+ client.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"tacacsplus_secret": schema.StringAttribute{
						Description: "Shared secret for authenticating to TACACS+ server.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
					},
					"tacacsplus_session_timeout": schema.Int64Attribute{
						Description: "TACACS+ session timeout value in seconds, 0 disables timeout.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(5),
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsAuthTacacsPlus,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthTACACSPlus",
		},
	}
}

type settingsAuthTacacsplusDataSource = framework.GenericDataSource[settingsAuthTacacsplusTerraformModel, *settingsAuthTacacsplusTerraformModel]

// NewSettingsAuthTACACSPlusDataSource is a helper function to instantiate the SettingsAuthTACACSPlus data source.
func NewSettingsAuthTACACSPlusDataSource() datasource.DataSource {
	return &settingsAuthTacacsplusDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_tacacsplus", Endpoint: "/api/v2/settings/tacacsplus/"}},
		Cfg: framework.DataSourceCfg[settingsAuthTacacsplusTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"tacacsplus_auth_protocol": dschema.StringAttribute{
						Description: "Choose the authentication protocol used by TACACS+ client.",
						Computed:    true,
					},
					"tacacsplus_host": dschema.StringAttribute{
						Description: "Hostname of TACACS+ server.",
						Computed:    true,
					},
					"tacacsplus_port": dschema.Int64Attribute{
						Description: "Port number of TACACS+ server.",
						Computed:    true,
					},
					"tacacsplus_rem_addr": dschema.BoolAttribute{
						Description: "Enable the client address sending by TACACS+ client.",
						Computed:    true,
					},
					"tacacsplus_secret": dschema.StringAttribute{
						Description: "Shared secret for authenticating to TACACS+ server.",
						Sensitive:   true,
						Computed:    true,
					},
					"tacacsplus_session_timeout": dschema.Int64Attribute{
						Description: "TACACS+ session timeout value in seconds, 0 disables timeout.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthTacacsPlus,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthTACACSPlus",
		},
	}
}
