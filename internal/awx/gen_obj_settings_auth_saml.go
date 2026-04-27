package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsAuthSamlTerraformModel struct {
	SAML_AUTO_CREATE_OBJECTS            types.Bool   `tfsdk:"saml_auto_create_objects" json:"SAML_AUTO_CREATE_OBJECTS"`
	SOCIAL_AUTH_SAML_CALLBACK_URL       types.String `tfsdk:"social_auth_saml_callback_url" json:"SOCIAL_AUTH_SAML_CALLBACK_URL"`
	SOCIAL_AUTH_SAML_ENABLED_IDPS       types.String `tfsdk:"social_auth_saml_enabled_idps" json:"SOCIAL_AUTH_SAML_ENABLED_IDPS"`
	SOCIAL_AUTH_SAML_EXTRA_DATA         types.List   `tfsdk:"social_auth_saml_extra_data" json:"SOCIAL_AUTH_SAML_EXTRA_DATA"`
	SOCIAL_AUTH_SAML_METADATA_URL       types.String `tfsdk:"social_auth_saml_metadata_url" json:"SOCIAL_AUTH_SAML_METADATA_URL"`
	SOCIAL_AUTH_SAML_ORGANIZATION_ATTR  types.String `tfsdk:"social_auth_saml_organization_attr" json:"SOCIAL_AUTH_SAML_ORGANIZATION_ATTR"`
	SOCIAL_AUTH_SAML_ORGANIZATION_MAP   types.String `tfsdk:"social_auth_saml_organization_map" json:"SOCIAL_AUTH_SAML_ORGANIZATION_MAP"`
	SOCIAL_AUTH_SAML_ORG_INFO           types.String `tfsdk:"social_auth_saml_org_info" json:"SOCIAL_AUTH_SAML_ORG_INFO"`
	SOCIAL_AUTH_SAML_SECURITY_CONFIG    types.String `tfsdk:"social_auth_saml_security_config" json:"SOCIAL_AUTH_SAML_SECURITY_CONFIG"`
	SOCIAL_AUTH_SAML_SP_ENTITY_ID       types.String `tfsdk:"social_auth_saml_sp_entity_id" json:"SOCIAL_AUTH_SAML_SP_ENTITY_ID"`
	SOCIAL_AUTH_SAML_SP_EXTRA           types.String `tfsdk:"social_auth_saml_sp_extra" json:"SOCIAL_AUTH_SAML_SP_EXTRA"`
	SOCIAL_AUTH_SAML_SP_PRIVATE_KEY     types.String `tfsdk:"social_auth_saml_sp_private_key" json:"SOCIAL_AUTH_SAML_SP_PRIVATE_KEY"`
	SOCIAL_AUTH_SAML_SP_PUBLIC_CERT     types.String `tfsdk:"social_auth_saml_sp_public_cert" json:"SOCIAL_AUTH_SAML_SP_PUBLIC_CERT"`
	SOCIAL_AUTH_SAML_SUPPORT_CONTACT    types.String `tfsdk:"social_auth_saml_support_contact" json:"SOCIAL_AUTH_SAML_SUPPORT_CONTACT"`
	SOCIAL_AUTH_SAML_TEAM_ATTR          types.String `tfsdk:"social_auth_saml_team_attr" json:"SOCIAL_AUTH_SAML_TEAM_ATTR"`
	SOCIAL_AUTH_SAML_TEAM_MAP           types.String `tfsdk:"social_auth_saml_team_map" json:"SOCIAL_AUTH_SAML_TEAM_MAP"`
	SOCIAL_AUTH_SAML_TECHNICAL_CONTACT  types.String `tfsdk:"social_auth_saml_technical_contact" json:"SOCIAL_AUTH_SAML_TECHNICAL_CONTACT"`
	SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR types.String `tfsdk:"social_auth_saml_user_flags_by_attr" json:"SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR"`
}

func (o *settingsAuthSamlTerraformModel) Clone() settingsAuthSamlTerraformModel {
	return *o
}

func (o *settingsAuthSamlTerraformModel) BodyRequest() *settingsAuthSamlBodyRequestModel {
	var req settingsAuthSamlBodyRequestModel
	req.SAML_AUTO_CREATE_OBJECTS = o.SAML_AUTO_CREATE_OBJECTS.ValueBool()
	req.SOCIAL_AUTH_SAML_ENABLED_IDPS = json.RawMessage(o.SOCIAL_AUTH_SAML_ENABLED_IDPS.ValueString())
	req.SOCIAL_AUTH_SAML_EXTRA_DATA = helpers.ListAsStringSlice(o.SOCIAL_AUTH_SAML_EXTRA_DATA, false)
	req.SOCIAL_AUTH_SAML_ORGANIZATION_ATTR = json.RawMessage(o.SOCIAL_AUTH_SAML_ORGANIZATION_ATTR.ValueString())
	req.SOCIAL_AUTH_SAML_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_SAML_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_SAML_ORG_INFO = json.RawMessage(o.SOCIAL_AUTH_SAML_ORG_INFO.ValueString())
	req.SOCIAL_AUTH_SAML_SECURITY_CONFIG = json.RawMessage(o.SOCIAL_AUTH_SAML_SECURITY_CONFIG.ValueString())
	req.SOCIAL_AUTH_SAML_SP_ENTITY_ID = o.SOCIAL_AUTH_SAML_SP_ENTITY_ID.ValueString()
	req.SOCIAL_AUTH_SAML_SP_EXTRA = json.RawMessage(o.SOCIAL_AUTH_SAML_SP_EXTRA.ValueString())
	req.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY = o.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY.ValueString()
	req.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT = o.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT.ValueString()
	req.SOCIAL_AUTH_SAML_SUPPORT_CONTACT = json.RawMessage(o.SOCIAL_AUTH_SAML_SUPPORT_CONTACT.ValueString())
	req.SOCIAL_AUTH_SAML_TEAM_ATTR = json.RawMessage(o.SOCIAL_AUTH_SAML_TEAM_ATTR.ValueString())
	req.SOCIAL_AUTH_SAML_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_SAML_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_SAML_TECHNICAL_CONTACT = json.RawMessage(o.SOCIAL_AUTH_SAML_TECHNICAL_CONTACT.ValueString())
	req.SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR = json.RawMessage(o.SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR.ValueString())
	return &req
}

func (o *settingsAuthSamlTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.SAML_AUTO_CREATE_OBJECTS, data["SAML_AUTO_CREATE_OBJECTS"]))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_CALLBACK_URL, data["SOCIAL_AUTH_SAML_CALLBACK_URL"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ENABLED_IDPS, data["SOCIAL_AUTH_SAML_ENABLED_IDPS"], false))
	collect(helpers.AttrValueSetListString(&o.SOCIAL_AUTH_SAML_EXTRA_DATA, data["SOCIAL_AUTH_SAML_EXTRA_DATA"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_METADATA_URL, data["SOCIAL_AUTH_SAML_METADATA_URL"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORGANIZATION_ATTR, data["SOCIAL_AUTH_SAML_ORGANIZATION_ATTR"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORGANIZATION_MAP, data["SOCIAL_AUTH_SAML_ORGANIZATION_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORG_INFO, data["SOCIAL_AUTH_SAML_ORG_INFO"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SECURITY_CONFIG, data["SOCIAL_AUTH_SAML_SECURITY_CONFIG"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_ENTITY_ID, data["SOCIAL_AUTH_SAML_SP_ENTITY_ID"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SP_EXTRA, data["SOCIAL_AUTH_SAML_SP_EXTRA"], false))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY, data["SOCIAL_AUTH_SAML_SP_PRIVATE_KEY"], true))
	collect(helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT, data["SOCIAL_AUTH_SAML_SP_PUBLIC_CERT"], true))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SUPPORT_CONTACT, data["SOCIAL_AUTH_SAML_SUPPORT_CONTACT"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TEAM_ATTR, data["SOCIAL_AUTH_SAML_TEAM_ATTR"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TEAM_MAP, data["SOCIAL_AUTH_SAML_TEAM_MAP"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TECHNICAL_CONTACT, data["SOCIAL_AUTH_SAML_TECHNICAL_CONTACT"], false))
	collect(helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR, data["SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR"], false))
	return diags, nil
}

type settingsAuthSamlBodyRequestModel struct {
	SAML_AUTO_CREATE_OBJECTS            bool            `json:"SAML_AUTO_CREATE_OBJECTS"`
	SOCIAL_AUTH_SAML_ENABLED_IDPS       json.RawMessage `json:"SOCIAL_AUTH_SAML_ENABLED_IDPS,omitempty"`
	SOCIAL_AUTH_SAML_EXTRA_DATA         []string        `json:"SOCIAL_AUTH_SAML_EXTRA_DATA,omitempty"`
	SOCIAL_AUTH_SAML_ORGANIZATION_ATTR  json.RawMessage `json:"SOCIAL_AUTH_SAML_ORGANIZATION_ATTR,omitempty"`
	SOCIAL_AUTH_SAML_ORGANIZATION_MAP   json.RawMessage `json:"SOCIAL_AUTH_SAML_ORGANIZATION_MAP,omitempty"`
	SOCIAL_AUTH_SAML_ORG_INFO           json.RawMessage `json:"SOCIAL_AUTH_SAML_ORG_INFO,omitempty"`
	SOCIAL_AUTH_SAML_SECURITY_CONFIG    json.RawMessage `json:"SOCIAL_AUTH_SAML_SECURITY_CONFIG,omitempty"`
	SOCIAL_AUTH_SAML_SP_ENTITY_ID       string          `json:"SOCIAL_AUTH_SAML_SP_ENTITY_ID,omitempty"`
	SOCIAL_AUTH_SAML_SP_EXTRA           json.RawMessage `json:"SOCIAL_AUTH_SAML_SP_EXTRA,omitempty"`
	SOCIAL_AUTH_SAML_SP_PRIVATE_KEY     string          `json:"SOCIAL_AUTH_SAML_SP_PRIVATE_KEY,omitempty"`
	SOCIAL_AUTH_SAML_SP_PUBLIC_CERT     string          `json:"SOCIAL_AUTH_SAML_SP_PUBLIC_CERT,omitempty"`
	SOCIAL_AUTH_SAML_SUPPORT_CONTACT    json.RawMessage `json:"SOCIAL_AUTH_SAML_SUPPORT_CONTACT,omitempty"`
	SOCIAL_AUTH_SAML_TEAM_ATTR          json.RawMessage `json:"SOCIAL_AUTH_SAML_TEAM_ATTR,omitempty"`
	SOCIAL_AUTH_SAML_TEAM_MAP           json.RawMessage `json:"SOCIAL_AUTH_SAML_TEAM_MAP,omitempty"`
	SOCIAL_AUTH_SAML_TECHNICAL_CONTACT  json.RawMessage `json:"SOCIAL_AUTH_SAML_TECHNICAL_CONTACT,omitempty"`
	SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR json.RawMessage `json:"SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR,omitempty"`
}

type settingsAuthSamlResource = framework.GenericResource[settingsAuthSamlTerraformModel, settingsAuthSamlBodyRequestModel, *settingsAuthSamlTerraformModel]

// NewSettingsAuthSAMLResource is a helper function to simplify the provider implementation.
func NewSettingsAuthSAMLResource() resource.Resource {
	return &settingsAuthSamlResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_saml", Endpoint: "/api/v2/settings/saml/"}},
		Cfg: framework.ResourceCfg[settingsAuthSamlTerraformModel, settingsAuthSamlBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"saml_auto_create_objects": schema.BoolAttribute{
						Description: "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_enabled_idps": schema.StringAttribute{
						Description: "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_extra_data": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_organization_attr": schema.StringAttribute{
						Description: "Used to translate user organization membership.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_org_info": schema.StringAttribute{
						Description: "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_security_config": schema.StringAttribute{
						Description: "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"requestedAuthnContext":false}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_sp_entity_id": schema.StringAttribute{
						Description: "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_sp_extra": schema.StringAttribute{
						Description: "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_sp_private_key": schema.StringAttribute{
						Description: "Create a keypair to use as a service provider (SP) and include the private key content here.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_sp_public_cert": schema.StringAttribute{
						Description: "Create a keypair to use as a service provider (SP) and include the certificate content here.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_support_contact": schema.StringAttribute{
						Description: "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_team_attr": schema.StringAttribute{
						Description: "Used to translate user team membership.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_technical_contact": schema.StringAttribute{
						Description: "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_user_flags_by_attr": schema.StringAttribute{
						Description: "Used to map super users and system auditors from SAML.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_callback_url": schema.StringAttribute{
						Description: "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_metadata_url": schema.StringAttribute{
						Description: "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			Hook:         hookSettingsSaml,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthSAML",
		},
	}
}

type settingsAuthSamlDataSource = framework.GenericDataSource[settingsAuthSamlTerraformModel, *settingsAuthSamlTerraformModel]

// NewSettingsAuthSAMLDataSource is a helper function to instantiate the SettingsAuthSAML data source.
func NewSettingsAuthSAMLDataSource() datasource.DataSource {
	return &settingsAuthSamlDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_saml", Endpoint: "/api/v2/settings/saml/"}},
		Cfg: framework.DataSourceCfg[settingsAuthSamlTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"saml_auto_create_objects": dschema.BoolAttribute{
						Description: "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login.",
						Computed:    true,
					},
					"social_auth_saml_callback_url": dschema.StringAttribute{
						Description: "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application.",
						Computed:    true,
					},
					"social_auth_saml_enabled_idps": dschema.StringAttribute{
						Description: "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax.",
						Computed:    true,
					},
					"social_auth_saml_extra_data": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value.",
						Computed:    true,
					},
					"social_auth_saml_metadata_url": dschema.StringAttribute{
						Description: "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL.",
						Computed:    true,
					},
					"social_auth_saml_organization_attr": dschema.StringAttribute{
						Description: "Used to translate user organization membership.",
						Computed:    true,
					},
					"social_auth_saml_organization_map": dschema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Computed:    true,
					},
					"social_auth_saml_org_info": dschema.StringAttribute{
						Description: "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"social_auth_saml_security_config": dschema.StringAttribute{
						Description: "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings",
						Computed:    true,
					},
					"social_auth_saml_sp_entity_id": dschema.StringAttribute{
						Description: "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service.",
						Computed:    true,
					},
					"social_auth_saml_sp_extra": dschema.StringAttribute{
						Description: "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting.",
						Computed:    true,
					},
					"social_auth_saml_sp_private_key": dschema.StringAttribute{
						Description: "Create a keypair to use as a service provider (SP) and include the private key content here.",
						Sensitive:   true,
						Computed:    true,
					},
					"social_auth_saml_sp_public_cert": dschema.StringAttribute{
						Description: "Create a keypair to use as a service provider (SP) and include the certificate content here.",
						Computed:    true,
					},
					"social_auth_saml_support_contact": dschema.StringAttribute{
						Description: "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"social_auth_saml_team_attr": dschema.StringAttribute{
						Description: "Used to translate user team membership.",
						Computed:    true,
					},
					"social_auth_saml_team_map": dschema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Computed:    true,
					},
					"social_auth_saml_technical_contact": dschema.StringAttribute{
						Description: "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"social_auth_saml_user_flags_by_attr": dschema.StringAttribute{
						Description: "Used to map super users and system auditors from SAML.",
						Computed:    true,
					},
				},
			},
			Hook:         hookSettingsSaml,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsAuthSAML",
		},
	}
}
