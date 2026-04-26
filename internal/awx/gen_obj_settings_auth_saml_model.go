package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
