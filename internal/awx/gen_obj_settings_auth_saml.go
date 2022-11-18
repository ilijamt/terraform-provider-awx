package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// settingsAuthSAMLTerraformModel maps the schema for SettingsAuthSAML when using Data Source
type settingsAuthSAMLTerraformModel struct {
	// SAML_AUTO_CREATE_OBJECTS "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login."
	SAML_AUTO_CREATE_OBJECTS types.Bool `tfsdk:"saml_auto_create_objects" json:"SAML_AUTO_CREATE_OBJECTS"`
	// SOCIAL_AUTH_SAML_CALLBACK_URL "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application."
	SOCIAL_AUTH_SAML_CALLBACK_URL types.String `tfsdk:"social_auth_saml_callback_url" json:"SOCIAL_AUTH_SAML_CALLBACK_URL"`
	// SOCIAL_AUTH_SAML_ENABLED_IDPS "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax."
	SOCIAL_AUTH_SAML_ENABLED_IDPS types.String `tfsdk:"social_auth_saml_enabled_idps" json:"SOCIAL_AUTH_SAML_ENABLED_IDPS"`
	// SOCIAL_AUTH_SAML_EXTRA_DATA "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value."
	SOCIAL_AUTH_SAML_EXTRA_DATA types.List `tfsdk:"social_auth_saml_extra_data" json:"SOCIAL_AUTH_SAML_EXTRA_DATA"`
	// SOCIAL_AUTH_SAML_METADATA_URL "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL."
	SOCIAL_AUTH_SAML_METADATA_URL types.String `tfsdk:"social_auth_saml_metadata_url" json:"SOCIAL_AUTH_SAML_METADATA_URL"`
	// SOCIAL_AUTH_SAML_ORGANIZATION_ATTR "Used to translate user organization membership."
	SOCIAL_AUTH_SAML_ORGANIZATION_ATTR types.String `tfsdk:"social_auth_saml_organization_attr" json:"SOCIAL_AUTH_SAML_ORGANIZATION_ATTR"`
	// SOCIAL_AUTH_SAML_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_SAML_ORGANIZATION_MAP types.String `tfsdk:"social_auth_saml_organization_map" json:"SOCIAL_AUTH_SAML_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_SAML_ORG_INFO "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_ORG_INFO types.String `tfsdk:"social_auth_saml_org_info" json:"SOCIAL_AUTH_SAML_ORG_INFO"`
	// SOCIAL_AUTH_SAML_SECURITY_CONFIG "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings"
	SOCIAL_AUTH_SAML_SECURITY_CONFIG types.String `tfsdk:"social_auth_saml_security_config" json:"SOCIAL_AUTH_SAML_SECURITY_CONFIG"`
	// SOCIAL_AUTH_SAML_SP_ENTITY_ID "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service."
	SOCIAL_AUTH_SAML_SP_ENTITY_ID types.String `tfsdk:"social_auth_saml_sp_entity_id" json:"SOCIAL_AUTH_SAML_SP_ENTITY_ID"`
	// SOCIAL_AUTH_SAML_SP_EXTRA "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting."
	SOCIAL_AUTH_SAML_SP_EXTRA types.String `tfsdk:"social_auth_saml_sp_extra" json:"SOCIAL_AUTH_SAML_SP_EXTRA"`
	// SOCIAL_AUTH_SAML_SP_PRIVATE_KEY "Create a keypair to use as a service provider (SP) and include the private key content here."
	SOCIAL_AUTH_SAML_SP_PRIVATE_KEY types.String `tfsdk:"social_auth_saml_sp_private_key" json:"SOCIAL_AUTH_SAML_SP_PRIVATE_KEY"`
	// SOCIAL_AUTH_SAML_SP_PUBLIC_CERT "Create a keypair to use as a service provider (SP) and include the certificate content here."
	SOCIAL_AUTH_SAML_SP_PUBLIC_CERT types.String `tfsdk:"social_auth_saml_sp_public_cert" json:"SOCIAL_AUTH_SAML_SP_PUBLIC_CERT"`
	// SOCIAL_AUTH_SAML_SUPPORT_CONTACT "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_SUPPORT_CONTACT types.String `tfsdk:"social_auth_saml_support_contact" json:"SOCIAL_AUTH_SAML_SUPPORT_CONTACT"`
	// SOCIAL_AUTH_SAML_TEAM_ATTR "Used to translate user team membership."
	SOCIAL_AUTH_SAML_TEAM_ATTR types.String `tfsdk:"social_auth_saml_team_attr" json:"SOCIAL_AUTH_SAML_TEAM_ATTR"`
	// SOCIAL_AUTH_SAML_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_SAML_TEAM_MAP types.String `tfsdk:"social_auth_saml_team_map" json:"SOCIAL_AUTH_SAML_TEAM_MAP"`
	// SOCIAL_AUTH_SAML_TECHNICAL_CONTACT "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_TECHNICAL_CONTACT types.String `tfsdk:"social_auth_saml_technical_contact" json:"SOCIAL_AUTH_SAML_TECHNICAL_CONTACT"`
	// SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR "Used to map super users and system auditors from SAML."
	SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR types.String `tfsdk:"social_auth_saml_user_flags_by_attr" json:"SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR"`
}

// Clone the object
func (o settingsAuthSAMLTerraformModel) Clone() settingsAuthSAMLTerraformModel {
	return settingsAuthSAMLTerraformModel{
		SAML_AUTO_CREATE_OBJECTS:            o.SAML_AUTO_CREATE_OBJECTS,
		SOCIAL_AUTH_SAML_CALLBACK_URL:       o.SOCIAL_AUTH_SAML_CALLBACK_URL,
		SOCIAL_AUTH_SAML_ENABLED_IDPS:       o.SOCIAL_AUTH_SAML_ENABLED_IDPS,
		SOCIAL_AUTH_SAML_EXTRA_DATA:         o.SOCIAL_AUTH_SAML_EXTRA_DATA,
		SOCIAL_AUTH_SAML_METADATA_URL:       o.SOCIAL_AUTH_SAML_METADATA_URL,
		SOCIAL_AUTH_SAML_ORGANIZATION_ATTR:  o.SOCIAL_AUTH_SAML_ORGANIZATION_ATTR,
		SOCIAL_AUTH_SAML_ORGANIZATION_MAP:   o.SOCIAL_AUTH_SAML_ORGANIZATION_MAP,
		SOCIAL_AUTH_SAML_ORG_INFO:           o.SOCIAL_AUTH_SAML_ORG_INFO,
		SOCIAL_AUTH_SAML_SECURITY_CONFIG:    o.SOCIAL_AUTH_SAML_SECURITY_CONFIG,
		SOCIAL_AUTH_SAML_SP_ENTITY_ID:       o.SOCIAL_AUTH_SAML_SP_ENTITY_ID,
		SOCIAL_AUTH_SAML_SP_EXTRA:           o.SOCIAL_AUTH_SAML_SP_EXTRA,
		SOCIAL_AUTH_SAML_SP_PRIVATE_KEY:     o.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY,
		SOCIAL_AUTH_SAML_SP_PUBLIC_CERT:     o.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT,
		SOCIAL_AUTH_SAML_SUPPORT_CONTACT:    o.SOCIAL_AUTH_SAML_SUPPORT_CONTACT,
		SOCIAL_AUTH_SAML_TEAM_ATTR:          o.SOCIAL_AUTH_SAML_TEAM_ATTR,
		SOCIAL_AUTH_SAML_TEAM_MAP:           o.SOCIAL_AUTH_SAML_TEAM_MAP,
		SOCIAL_AUTH_SAML_TECHNICAL_CONTACT:  o.SOCIAL_AUTH_SAML_TECHNICAL_CONTACT,
		SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR: o.SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthSAML
func (o settingsAuthSAMLTerraformModel) BodyRequest() (req settingsAuthSAMLBodyRequestModel) {
	req.SAML_AUTO_CREATE_OBJECTS = o.SAML_AUTO_CREATE_OBJECTS.ValueBool()
	req.SOCIAL_AUTH_SAML_ENABLED_IDPS = json.RawMessage(o.SOCIAL_AUTH_SAML_ENABLED_IDPS.ValueString())
	req.SOCIAL_AUTH_SAML_EXTRA_DATA = []string{}
	for _, val := range o.SOCIAL_AUTH_SAML_EXTRA_DATA.Elements() {
		if _, ok := val.(types.String); ok {
			req.SOCIAL_AUTH_SAML_EXTRA_DATA = append(req.SOCIAL_AUTH_SAML_EXTRA_DATA, val.(types.String).ValueString())
		} else {
			req.SOCIAL_AUTH_SAML_EXTRA_DATA = append(req.SOCIAL_AUTH_SAML_EXTRA_DATA, val.String())
		}
	}
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
	return
}

func (o *settingsAuthSAMLTerraformModel) setSamlAutoCreateObjects(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SAML_AUTO_CREATE_OBJECTS, data)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_CALLBACK_URL, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlEnabledIdps(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ENABLED_IDPS, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlExtraData(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.SOCIAL_AUTH_SAML_EXTRA_DATA, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlMetadataUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_METADATA_URL, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlOrganizationAttr(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORGANIZATION_ATTR, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlOrgInfo(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORG_INFO, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlSecurityConfig(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SECURITY_CONFIG, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlSpEntityId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_ENTITY_ID, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlSpExtra(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SP_EXTRA, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlSpPrivateKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY, data, true)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlSpPublicCert(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT, data, true)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlSupportContact(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SUPPORT_CONTACT, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlTeamAttr(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TEAM_ATTR, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TEAM_MAP, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlTechnicalContact(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TECHNICAL_CONTACT, data, false)
}

func (o *settingsAuthSAMLTerraformModel) setSocialAuthSamlUserFlagsByAttr(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR, data, false)
}

func (o *settingsAuthSAMLTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSamlAutoCreateObjects(data["SAML_AUTO_CREATE_OBJECTS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlCallbackUrl(data["SOCIAL_AUTH_SAML_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlEnabledIdps(data["SOCIAL_AUTH_SAML_ENABLED_IDPS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlExtraData(data["SOCIAL_AUTH_SAML_EXTRA_DATA"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlMetadataUrl(data["SOCIAL_AUTH_SAML_METADATA_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlOrganizationAttr(data["SOCIAL_AUTH_SAML_ORGANIZATION_ATTR"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlOrganizationMap(data["SOCIAL_AUTH_SAML_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlOrgInfo(data["SOCIAL_AUTH_SAML_ORG_INFO"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlSecurityConfig(data["SOCIAL_AUTH_SAML_SECURITY_CONFIG"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlSpEntityId(data["SOCIAL_AUTH_SAML_SP_ENTITY_ID"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlSpExtra(data["SOCIAL_AUTH_SAML_SP_EXTRA"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlSpPrivateKey(data["SOCIAL_AUTH_SAML_SP_PRIVATE_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlSpPublicCert(data["SOCIAL_AUTH_SAML_SP_PUBLIC_CERT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlSupportContact(data["SOCIAL_AUTH_SAML_SUPPORT_CONTACT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlTeamAttr(data["SOCIAL_AUTH_SAML_TEAM_ATTR"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlTeamMap(data["SOCIAL_AUTH_SAML_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlTechnicalContact(data["SOCIAL_AUTH_SAML_TECHNICAL_CONTACT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthSamlUserFlagsByAttr(data["SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthSAMLBodyRequestModel maps the schema for SettingsAuthSAML for creating and updating the data
type settingsAuthSAMLBodyRequestModel struct {
	// SAML_AUTO_CREATE_OBJECTS "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login."
	SAML_AUTO_CREATE_OBJECTS bool `json:"SAML_AUTO_CREATE_OBJECTS"`
	// SOCIAL_AUTH_SAML_ENABLED_IDPS "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax."
	SOCIAL_AUTH_SAML_ENABLED_IDPS json.RawMessage `json:"SOCIAL_AUTH_SAML_ENABLED_IDPS,omitempty"`
	// SOCIAL_AUTH_SAML_EXTRA_DATA "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value."
	SOCIAL_AUTH_SAML_EXTRA_DATA []string `json:"SOCIAL_AUTH_SAML_EXTRA_DATA,omitempty"`
	// SOCIAL_AUTH_SAML_ORGANIZATION_ATTR "Used to translate user organization membership."
	SOCIAL_AUTH_SAML_ORGANIZATION_ATTR json.RawMessage `json:"SOCIAL_AUTH_SAML_ORGANIZATION_ATTR,omitempty"`
	// SOCIAL_AUTH_SAML_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_SAML_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_SAML_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_SAML_ORG_INFO "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_ORG_INFO json.RawMessage `json:"SOCIAL_AUTH_SAML_ORG_INFO"`
	// SOCIAL_AUTH_SAML_SECURITY_CONFIG "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings"
	SOCIAL_AUTH_SAML_SECURITY_CONFIG json.RawMessage `json:"SOCIAL_AUTH_SAML_SECURITY_CONFIG,omitempty"`
	// SOCIAL_AUTH_SAML_SP_ENTITY_ID "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service."
	SOCIAL_AUTH_SAML_SP_ENTITY_ID string `json:"SOCIAL_AUTH_SAML_SP_ENTITY_ID,omitempty"`
	// SOCIAL_AUTH_SAML_SP_EXTRA "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting."
	SOCIAL_AUTH_SAML_SP_EXTRA json.RawMessage `json:"SOCIAL_AUTH_SAML_SP_EXTRA,omitempty"`
	// SOCIAL_AUTH_SAML_SP_PRIVATE_KEY "Create a keypair to use as a service provider (SP) and include the private key content here."
	SOCIAL_AUTH_SAML_SP_PRIVATE_KEY string `json:"SOCIAL_AUTH_SAML_SP_PRIVATE_KEY"`
	// SOCIAL_AUTH_SAML_SP_PUBLIC_CERT "Create a keypair to use as a service provider (SP) and include the certificate content here."
	SOCIAL_AUTH_SAML_SP_PUBLIC_CERT string `json:"SOCIAL_AUTH_SAML_SP_PUBLIC_CERT"`
	// SOCIAL_AUTH_SAML_SUPPORT_CONTACT "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_SUPPORT_CONTACT json.RawMessage `json:"SOCIAL_AUTH_SAML_SUPPORT_CONTACT"`
	// SOCIAL_AUTH_SAML_TEAM_ATTR "Used to translate user team membership."
	SOCIAL_AUTH_SAML_TEAM_ATTR json.RawMessage `json:"SOCIAL_AUTH_SAML_TEAM_ATTR,omitempty"`
	// SOCIAL_AUTH_SAML_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_SAML_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_SAML_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_SAML_TECHNICAL_CONTACT "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_TECHNICAL_CONTACT json.RawMessage `json:"SOCIAL_AUTH_SAML_TECHNICAL_CONTACT"`
	// SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR "Used to map super users and system auditors from SAML."
	SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR json.RawMessage `json:"SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthSAMLDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthSAMLDataSource{}
)

// NewSettingsAuthSAMLDataSource is a helper function to instantiate the SettingsAuthSAML data source.
func NewSettingsAuthSAMLDataSource() datasource.DataSource {
	return &settingsAuthSAMLDataSource{}
}

// settingsAuthSAMLDataSource is the data source implementation.
type settingsAuthSAMLDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthSAMLDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/saml/"
}

// Metadata returns the data source type name.
func (o *settingsAuthSAMLDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_saml"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthSAMLDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthSAML",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"saml_auto_create_objects": {
					Description: "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_callback_url": {
					Description: "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_enabled_idps": {
					Description: "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_extra_data": {
					Description: "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_metadata_url": {
					Description: "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_organization_attr": {
					Description: "Used to translate user organization membership.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_org_info": {
					Description: "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_security_config": {
					Description: "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_entity_id": {
					Description: "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_extra": {
					Description: "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_private_key": {
					Description: "Create a keypair to use as a service provider (SP) and include the private key content here.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_public_cert": {
					Description: "Create a keypair to use as a service provider (SP) and include the certificate content here.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_support_contact": {
					Description: "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_team_attr": {
					Description: "Used to translate user team membership.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_technical_contact": {
					Description: "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_user_flags_by_attr": {
					Description: "Used to map super users and system auditors from SAML.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthSAMLDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthSAMLTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthSAML
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthSAML on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthSAML
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthSAML on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	if err = hookSettingsSaml(ctx, SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthSAML",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsAuthSAMLResource{}
	_ resource.ResourceWithConfigure = &settingsAuthSAMLResource{}
)

// NewSettingsAuthSAMLResource is a helper function to simplify the provider implementation.
func NewSettingsAuthSAMLResource() resource.Resource {
	return &settingsAuthSAMLResource{}
}

type settingsAuthSAMLResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthSAMLResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/saml/"
}

func (o settingsAuthSAMLResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_saml"
}

func (o settingsAuthSAMLResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthSAML",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"saml_auto_create_objects": {
					Description: "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_enabled_idps": {
					Description: "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_extra_data": {
					Description: "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_organization_attr": {
					Description: "Used to translate user organization membership.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_org_info": {
					Description: "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_security_config": {
					Description: "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"requestedAuthnContext":false}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_entity_id": {
					Description: "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_extra": {
					Description: "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_private_key": {
					Description: "Create a keypair to use as a service provider (SP) and include the private key content here.",
					Type:        types.StringType,
					Sensitive:   true,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_sp_public_cert": {
					Description: "Create a keypair to use as a service provider (SP) and include the certificate content here.",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_support_contact": {
					Description: "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_team_attr": {
					Description: "Used to translate user team membership.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_technical_contact": {
					Description: "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax.",
					Type:        types.StringType,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_saml_user_flags_by_attr": {
					Description: "Used to map super users and system auditors from SAML.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"social_auth_saml_callback_url": {
					Description: "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"social_auth_saml_metadata_url": {
					Description: "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *settingsAuthSAMLResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthSAMLTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthSAML
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthSAML on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthSAML resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthSAML on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsSaml(ctx, SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthSAML",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthSAMLResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthSAMLTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthSAML
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthSAML on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthSAML from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthSAML on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsSaml(ctx, SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthSAML",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthSAMLResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthSAMLTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthSAML
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthSAML on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthSAML resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthSAML on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsSaml(ctx, SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthSAML",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthSAMLResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
