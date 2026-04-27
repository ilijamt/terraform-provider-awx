package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthGoogleOauth2TerraformModel struct {
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS types.String `tfsdk:"social_auth_google_oauth2_auth_extra_arguments" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL         types.String `tfsdk:"social_auth_google_oauth2_callback_url" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY                  types.String `tfsdk:"social_auth_google_oauth2_key" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP     types.String `tfsdk:"social_auth_google_oauth2_organization_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET               types.String `tfsdk:"social_auth_google_oauth2_secret" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP             types.String `tfsdk:"social_auth_google_oauth2_team_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS  types.List   `tfsdk:"social_auth_google_oauth2_whitelisted_domains" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"`
}

func (o *settingsAuthGoogleOauth2TerraformModel) Clone() settingsAuthGoogleOauth2TerraformModel {
	return *o
}

func (o *settingsAuthGoogleOauth2TerraformModel) BodyRequest() *settingsAuthGoogleOauth2BodyRequestModel {
	var req settingsAuthGoogleOauth2BodyRequestModel
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY = o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET = o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = helpers.ListAsStringSlice(o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, false)
	return &req
}

func (o *settingsAuthGoogleOauth2TerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS, data["SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL, data["SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY, data["SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP, data["SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET, data["SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP, data["SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"], false))
	collect(helpers.AttrValueSetListString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, data["SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"], false))
	return diags, nil
}

type settingsAuthGoogleOauth2BodyRequestModel struct {
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY                  string          `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP     json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET               string          `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP             json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS  []string        `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS,omitempty"`
}

type settingsAuthGoogleOauth2Resource = framework.GenericResource[settingsAuthGoogleOauth2TerraformModel, settingsAuthGoogleOauth2BodyRequestModel, *settingsAuthGoogleOauth2TerraformModel]

// NewSettingsAuthGoogleOauth2Resource is a helper function to simplify the provider implementation.
func NewSettingsAuthGoogleOauth2Resource() resource.Resource {
	return &settingsAuthGoogleOauth2Resource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_google_oauth2", Endpoint: "/api/v2/settings/google-oauth2/"}},
		Cfg: framework.ResourceCfg[settingsAuthGoogleOauth2TerraformModel, settingsAuthGoogleOauth2BodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"social_auth_google_oauth2_auth_extra_arguments": schema.StringAttribute{
						Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_key": schema.StringAttribute{
						Description: "The OAuth2 key from your web application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_secret": schema.StringAttribute{
						Description: "The OAuth2 secret from your web application.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_whitelisted_domains": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_google_oauth2_callback_url": schema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsAuthGoogleOauth2,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGoogleOauth2",
		},
	}
}

type settingsAuthGoogleOauth2DataSource = framework.GenericDataSource[settingsAuthGoogleOauth2TerraformModel, *settingsAuthGoogleOauth2TerraformModel]

// NewSettingsAuthGoogleOauth2DataSource is a helper function to instantiate the SettingsAuthGoogleOauth2 data source.
func NewSettingsAuthGoogleOauth2DataSource() datasource.DataSource {
	return &settingsAuthGoogleOauth2DataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_google_oauth2", Endpoint: "/api/v2/settings/google-oauth2/"}},
		Cfg: framework.DataSourceCfg[settingsAuthGoogleOauth2TerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"social_auth_google_oauth2_auth_extra_arguments": dschema.StringAttribute{
						Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
						Computed:    true,
					},
					"social_auth_google_oauth2_callback_url": dschema.StringAttribute{
						Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
						Computed:    true,
					},
					"social_auth_google_oauth2_key": dschema.StringAttribute{
						Description: "The OAuth2 key from your web application.",
						Computed:    true,
					},
					"social_auth_google_oauth2_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_google_oauth2_secret": dschema.StringAttribute{
						Description: "The OAuth2 secret from your web application.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_google_oauth2_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_google_oauth2_whitelisted_domains": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsAuthGoogleOauth2,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthGoogleOauth2",
		},
	}
}
