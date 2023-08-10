package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// settingsAuthSamlTerraformModel maps the schema for SettingsAuthSAML when using Data Source
type settingsAuthSamlTerraformModel struct {
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
func (o *settingsAuthSamlTerraformModel) Clone() settingsAuthSamlTerraformModel {
	return settingsAuthSamlTerraformModel{
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
func (o *settingsAuthSamlTerraformModel) BodyRequest() (req settingsAuthSamlBodyRequestModel) {
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

func (o *settingsAuthSamlTerraformModel) setSamlAutoCreateObjects(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SAML_AUTO_CREATE_OBJECTS, data)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_CALLBACK_URL, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlEnabledIdps(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ENABLED_IDPS, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlExtraData(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.SOCIAL_AUTH_SAML_EXTRA_DATA, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlMetadataUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_METADATA_URL, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlOrganizationAttr(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORGANIZATION_ATTR, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlOrgInfo(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_ORG_INFO, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlSecurityConfig(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SECURITY_CONFIG, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlSpEntityId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_ENTITY_ID, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlSpExtra(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SP_EXTRA, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlSpPrivateKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY, data, true)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlSpPublicCert(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT, data, true)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlSupportContact(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_SUPPORT_CONTACT, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlTeamAttr(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TEAM_ATTR, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TEAM_MAP, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlTechnicalContact(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_TECHNICAL_CONTACT, data, false)
}

func (o *settingsAuthSamlTerraformModel) setSocialAuthSamlUserFlagsByAttr(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR, data, false)
}

func (o *settingsAuthSamlTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
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

// settingsAuthSamlBodyRequestModel maps the schema for SettingsAuthSAML for creating and updating the data
type settingsAuthSamlBodyRequestModel struct {
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
	SOCIAL_AUTH_SAML_ORG_INFO json.RawMessage `json:"SOCIAL_AUTH_SAML_ORG_INFO,omitempty"`
	// SOCIAL_AUTH_SAML_SECURITY_CONFIG "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings"
	SOCIAL_AUTH_SAML_SECURITY_CONFIG json.RawMessage `json:"SOCIAL_AUTH_SAML_SECURITY_CONFIG,omitempty"`
	// SOCIAL_AUTH_SAML_SP_ENTITY_ID "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service."
	SOCIAL_AUTH_SAML_SP_ENTITY_ID string `json:"SOCIAL_AUTH_SAML_SP_ENTITY_ID,omitempty"`
	// SOCIAL_AUTH_SAML_SP_EXTRA "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting."
	SOCIAL_AUTH_SAML_SP_EXTRA json.RawMessage `json:"SOCIAL_AUTH_SAML_SP_EXTRA,omitempty"`
	// SOCIAL_AUTH_SAML_SP_PRIVATE_KEY "Create a keypair to use as a service provider (SP) and include the private key content here."
	SOCIAL_AUTH_SAML_SP_PRIVATE_KEY string `json:"SOCIAL_AUTH_SAML_SP_PRIVATE_KEY,omitempty"`
	// SOCIAL_AUTH_SAML_SP_PUBLIC_CERT "Create a keypair to use as a service provider (SP) and include the certificate content here."
	SOCIAL_AUTH_SAML_SP_PUBLIC_CERT string `json:"SOCIAL_AUTH_SAML_SP_PUBLIC_CERT,omitempty"`
	// SOCIAL_AUTH_SAML_SUPPORT_CONTACT "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_SUPPORT_CONTACT json.RawMessage `json:"SOCIAL_AUTH_SAML_SUPPORT_CONTACT,omitempty"`
	// SOCIAL_AUTH_SAML_TEAM_ATTR "Used to translate user team membership."
	SOCIAL_AUTH_SAML_TEAM_ATTR json.RawMessage `json:"SOCIAL_AUTH_SAML_TEAM_ATTR,omitempty"`
	// SOCIAL_AUTH_SAML_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_SAML_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_SAML_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_SAML_TECHNICAL_CONTACT "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax."
	SOCIAL_AUTH_SAML_TECHNICAL_CONTACT json.RawMessage `json:"SOCIAL_AUTH_SAML_TECHNICAL_CONTACT,omitempty"`
	// SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR "Used to map super users and system auditors from SAML."
	SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR json.RawMessage `json:"SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR,omitempty"`
}
