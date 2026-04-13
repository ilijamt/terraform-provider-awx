package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsAuthSamlResource = framework.GenericResource[settingsAuthSamlTerraformModel, settingsAuthSamlBodyRequestModel, *settingsAuthSamlTerraformModel]

// NewSettingsAuthSAMLResource is a helper function to simplify the provider implementation.
func NewSettingsAuthSAMLResource() resource.Resource {
	return &settingsAuthSamlResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_auth_saml", Endpoint: "/api/v2/settings/saml/"}},
		Cfg: framework.ResourceCfg[settingsAuthSamlTerraformModel, settingsAuthSamlBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"saml_auto_create_objects": schema.BoolAttribute{
						Description: "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"social_auth_saml_enabled_idps": schema.StringAttribute{
						Description: "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_extra_data": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.List{},
					},
					"social_auth_saml_organization_attr": schema.StringAttribute{
						Description: "Used to translate user organization membership.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_org_info": schema.StringAttribute{
						Description: "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_security_config": schema.StringAttribute{
						Description: "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"requestedAuthnContext":false}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_sp_entity_id": schema.StringAttribute{
						Description: "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_sp_extra": schema.StringAttribute{
						Description: "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_sp_private_key": schema.StringAttribute{
						Description: "Create a keypair to use as a service provider (SP) and include the private key content here.",
						Sensitive:   true,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_sp_public_cert": schema.StringAttribute{
						Description: "Create a keypair to use as a service provider (SP) and include the certificate content here.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_support_contact": schema.StringAttribute{
						Description: "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_team_attr": schema.StringAttribute{
						Description: "Used to translate user team membership.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_technical_contact": schema.StringAttribute{
						Description: "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_saml_user_flags_by_attr": schema.StringAttribute{
						Description: "Used to map super users and system auditors from SAML.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					// Write only elements
					// Data only elements
					"social_auth_saml_callback_url": schema.StringAttribute{
						Description: "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"social_auth_saml_metadata_url": schema.StringAttribute{
						Description: "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
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
