package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsOpenIdconnectTerraformModel struct {
	SOCIAL_AUTH_OIDC_KEY           types.String `tfsdk:"social_auth_oidc_key" json:"SOCIAL_AUTH_OIDC_KEY"`
	SOCIAL_AUTH_OIDC_OIDC_ENDPOINT types.String `tfsdk:"social_auth_oidc_oidc_endpoint" json:"SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"`
	SOCIAL_AUTH_OIDC_SECRET        types.String `tfsdk:"social_auth_oidc_secret" json:"SOCIAL_AUTH_OIDC_SECRET"`
	SOCIAL_AUTH_OIDC_VERIFY_SSL    types.Bool   `tfsdk:"social_auth_oidc_verify_ssl" json:"SOCIAL_AUTH_OIDC_VERIFY_SSL"`
}

func (o *settingsOpenIdconnectTerraformModel) Clone() settingsOpenIdconnectTerraformModel {
	return *o
}

func (o *settingsOpenIdconnectTerraformModel) BodyRequest() *settingsOpenIdconnectBodyRequestModel {
	var req settingsOpenIdconnectBodyRequestModel
	req.SOCIAL_AUTH_OIDC_KEY = o.SOCIAL_AUTH_OIDC_KEY.ValueString()
	req.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT = o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT.ValueString()
	req.SOCIAL_AUTH_OIDC_SECRET = o.SOCIAL_AUTH_OIDC_SECRET.ValueString()
	req.SOCIAL_AUTH_OIDC_VERIFY_SSL = o.SOCIAL_AUTH_OIDC_VERIFY_SSL.ValueBool()
	return &req
}

func (o *settingsOpenIdconnectTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_KEY, data["SOCIAL_AUTH_OIDC_KEY"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT, data["SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_SECRET, data["SOCIAL_AUTH_OIDC_SECRET"], false))
	collect(helpers.AttrValueSetBool(&o.SOCIAL_AUTH_OIDC_VERIFY_SSL, data["SOCIAL_AUTH_OIDC_VERIFY_SSL"]))
	return diags, nil
}

type settingsOpenIdconnectBodyRequestModel struct {
	SOCIAL_AUTH_OIDC_KEY           string `json:"SOCIAL_AUTH_OIDC_KEY,omitempty"`
	SOCIAL_AUTH_OIDC_OIDC_ENDPOINT string `json:"SOCIAL_AUTH_OIDC_OIDC_ENDPOINT,omitempty"`
	SOCIAL_AUTH_OIDC_SECRET        string `json:"SOCIAL_AUTH_OIDC_SECRET,omitempty"`
	SOCIAL_AUTH_OIDC_VERIFY_SSL    bool   `json:"SOCIAL_AUTH_OIDC_VERIFY_SSL"`
}

type settingsOpenIdconnectResource = framework.GenericResource[settingsOpenIdconnectTerraformModel, settingsOpenIdconnectBodyRequestModel, *settingsOpenIdconnectTerraformModel]

// NewSettingsOpenIDConnectResource is a helper function to simplify the provider implementation.
func NewSettingsOpenIDConnectResource() resource.Resource {
	return &settingsOpenIdconnectResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_oidc", Endpoint: "/api/v2/settings/oidc/"}},
		Cfg: framework.ResourceCfg[settingsOpenIdconnectTerraformModel, settingsOpenIdconnectBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_oidc_key": schema.StringAttribute{
						Description: "The OIDC key (Client ID) from your IDP.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_oidc_oidc_endpoint": schema.StringAttribute{
						Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_oidc_secret": schema.StringAttribute{
						Description: "The OIDC secret (Client Secret) from your IDP.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_oidc_verify_ssl": schema.BoolAttribute{
						Description: "Verify the OIDC provider ssl certificate.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsOidc,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsOpenIDConnect",
		},
	}
}

type settingsOpenIdconnectDataSource = framework.GenericDataSource[settingsOpenIdconnectTerraformModel, *settingsOpenIdconnectTerraformModel]

// NewSettingsOpenIDConnectDataSource is a helper function to instantiate the SettingsOpenIDConnect data source.
func NewSettingsOpenIDConnectDataSource() datasource.DataSource {
	return &settingsOpenIdconnectDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_oidc", Endpoint: "/api/v2/settings/oidc/"}},
		Cfg: framework.DataSourceCfg[settingsOpenIdconnectTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_oidc_key": dschema.StringAttribute{
						Description: "The OIDC key (Client ID) from your IDP.",
						Computed:    true,
					},
					"social_auth_oidc_oidc_endpoint": dschema.StringAttribute{
						Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
						Computed:    true,
					},
					"social_auth_oidc_secret": dschema.StringAttribute{
						Description: "The OIDC secret (Client Secret) from your IDP.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_oidc_verify_ssl": dschema.BoolAttribute{
						Description: "Verify the OIDC provider ssl certificate.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsOidc,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsOpenIDConnect",
		},
	}
}
