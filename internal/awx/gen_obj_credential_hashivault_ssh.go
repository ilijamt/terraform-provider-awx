package awx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialHashivaultSshTerraformModel exposes the typed AWX HashiCorp Vault Signed SSH
// credential (credential_hashivault_ssh) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialHashivaultSshTerraformModel struct {
	ID                types.Int64  `tfsdk:"id" json:"id"`
	Name              types.String `tfsdk:"name" json:"name"`
	Description       types.String `tfsdk:"description" json:"description"`
	Organization      types.Int64  `tfsdk:"organization" json:"organization"`
	Team              types.Int64  `tfsdk:"team" json:"team"`
	User              types.Int64  `tfsdk:"user" json:"user"`
	Kind              types.String `tfsdk:"kind" json:"kind"`
	Managed           types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType    types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Cacert            types.String `tfsdk:"cacert" json:"-"`
	ClientCertPrivate types.String `tfsdk:"client_cert_private" json:"-"`
	ClientCertPublic  types.String `tfsdk:"client_cert_public" json:"-"`
	ClientCertRole    types.String `tfsdk:"client_cert_role" json:"-"`
	DefaultAuthPath   types.String `tfsdk:"default_auth_path" json:"-"`
	KubernetesRole    types.String `tfsdk:"kubernetes_role" json:"-"`
	Namespace         types.String `tfsdk:"namespace" json:"-"`
	Password          types.String `tfsdk:"password" json:"-"`
	RoleId            types.String `tfsdk:"role_id" json:"-"`
	SecretId          types.String `tfsdk:"secret_id" json:"-"`
	Token             types.String `tfsdk:"token" json:"-"`
	Url               types.String `tfsdk:"url" json:"-"`
	Username          types.String `tfsdk:"username" json:"-"`
}

func (o *credentialHashivaultSshTerraformModel) Clone() credentialHashivaultSshTerraformModel {
	return *o
}

type credentialHashivaultSshBodyRequestModel struct {
	CredentialType int64           `json:"credential_type"`
	Description    string          `json:"description,omitempty"`
	Inputs         json.RawMessage `json:"inputs,omitempty"`
	Name           string          `json:"name"`
	Organization   int64           `json:"organization,omitempty"`
	Team           int64           `json:"team,omitempty"`
	User           int64           `json:"user,omitempty"`
}

// BodyRequest folds typed input fields back into a single `inputs` JSON object;
// null/unknown values are dropped so the API doesn't receive empty strings for
// unset optionals.
func (o *credentialHashivaultSshTerraformModel) BodyRequest() *credentialHashivaultSshBodyRequestModel {
	req := &credentialHashivaultSshBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Cacert.IsNull() && !o.Cacert.IsUnknown() {
		inputs["cacert"] = o.Cacert.ValueString()
	}
	if !o.ClientCertPrivate.IsNull() && !o.ClientCertPrivate.IsUnknown() {
		inputs["client_cert_private"] = o.ClientCertPrivate.ValueString()
	}
	if !o.ClientCertPublic.IsNull() && !o.ClientCertPublic.IsUnknown() {
		inputs["client_cert_public"] = o.ClientCertPublic.ValueString()
	}
	if !o.ClientCertRole.IsNull() && !o.ClientCertRole.IsUnknown() {
		inputs["client_cert_role"] = o.ClientCertRole.ValueString()
	}
	if !o.DefaultAuthPath.IsNull() && !o.DefaultAuthPath.IsUnknown() {
		inputs["default_auth_path"] = o.DefaultAuthPath.ValueString()
	}
	if !o.KubernetesRole.IsNull() && !o.KubernetesRole.IsUnknown() {
		inputs["kubernetes_role"] = o.KubernetesRole.ValueString()
	}
	if !o.Namespace.IsNull() && !o.Namespace.IsUnknown() {
		inputs["namespace"] = o.Namespace.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.RoleId.IsNull() && !o.RoleId.IsUnknown() {
		inputs["role_id"] = o.RoleId.ValueString()
	}
	if !o.SecretId.IsNull() && !o.SecretId.IsUnknown() {
		inputs["secret_id"] = o.SecretId.ValueString()
	}
	if !o.Token.IsNull() && !o.Token.IsUnknown() {
		inputs["token"] = o.Token.ValueString()
	}
	if !o.Url.IsNull() && !o.Url.IsUnknown() {
		inputs["url"] = o.Url.ValueString()
	}
	if !o.Username.IsNull() && !o.Username.IsUnknown() {
		inputs["username"] = o.Username.ValueString()
	}
	if len(inputs) > 0 {
		payload, _ := json.Marshal(inputs)
		req.Inputs = payload
	}
	return req
}

// UpdateFromApiData unfolds the AWX response back into the typed model. Secret
// fields come back as `$encrypted$` placeholders; the per-credential-type
// pre-state-set hook reconciles them against prior plan state.
func (o *credentialHashivaultSshTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))

	if inputs, ok := data["inputs"].(map[string]any); ok {
		collect(helpers.AttrValueSetString(&o.Cacert, inputs["cacert"], false))
		collect(helpers.AttrValueSetString(&o.ClientCertPrivate, inputs["client_cert_private"], false))
		collect(helpers.AttrValueSetString(&o.ClientCertPublic, inputs["client_cert_public"], false))
		collect(helpers.AttrValueSetString(&o.ClientCertRole, inputs["client_cert_role"], false))
		collect(helpers.AttrValueSetString(&o.DefaultAuthPath, inputs["default_auth_path"], false))
		collect(helpers.AttrValueSetString(&o.KubernetesRole, inputs["kubernetes_role"], false))
		collect(helpers.AttrValueSetString(&o.Namespace, inputs["namespace"], false))
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.RoleId, inputs["role_id"], false))
		collect(helpers.AttrValueSetString(&o.SecretId, inputs["secret_id"], false))
		collect(helpers.AttrValueSetString(&o.Token, inputs["token"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialHashivaultSsh reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialHashivaultSsh(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialHashivaultSshTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.ClientCertPrivate.IsNull() || orig.ClientCertPrivate.IsUnknown() {
			state.ClientCertPrivate = types.StringNull()
		} else {
			state.ClientCertPrivate = orig.ClientCertPrivate
		}
		if orig.Password.IsNull() || orig.Password.IsUnknown() {
			state.Password = types.StringNull()
		} else {
			state.Password = orig.Password
		}
		if orig.SecretId.IsNull() || orig.SecretId.IsUnknown() {
			state.SecretId = types.StringNull()
		} else {
			state.SecretId = orig.SecretId
		}
		if orig.Token.IsNull() || orig.Token.IsUnknown() {
			state.Token = types.StringNull()
		} else {
			state.Token = orig.Token
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.ClientCertPrivate, state.ClientCertPrivate); subbed {
			state.ClientCertPrivate = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.SecretId, state.SecretId); subbed {
			state.SecretId = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.Token, state.Token); subbed {
			state.Token = v
		}
	}
	return nil
}

// credentialHashivaultSshDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialHashivaultSshDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	Team             types.Int64  `tfsdk:"team" json:"team"`
	User             types.Int64  `tfsdk:"user" json:"user"`
	Kind             types.String `tfsdk:"kind" json:"kind"`
	Managed          types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType   types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Cacert           types.String `tfsdk:"cacert" json:"-"`
	ClientCertPublic types.String `tfsdk:"client_cert_public" json:"-"`
	ClientCertRole   types.String `tfsdk:"client_cert_role" json:"-"`
	DefaultAuthPath  types.String `tfsdk:"default_auth_path" json:"-"`
	KubernetesRole   types.String `tfsdk:"kubernetes_role" json:"-"`
	Namespace        types.String `tfsdk:"namespace" json:"-"`
	RoleId           types.String `tfsdk:"role_id" json:"-"`
	Url              types.String `tfsdk:"url" json:"-"`
	Username         types.String `tfsdk:"username" json:"-"`
}

func (o *credentialHashivaultSshDataSourceTerraformModel) Clone() credentialHashivaultSshDataSourceTerraformModel {
	return *o
}

func (o *credentialHashivaultSshDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))

	if inputs, ok := data["inputs"].(map[string]any); ok {
		collect(helpers.AttrValueSetString(&o.Cacert, inputs["cacert"], false))
		collect(helpers.AttrValueSetString(&o.ClientCertPublic, inputs["client_cert_public"], false))
		collect(helpers.AttrValueSetString(&o.ClientCertRole, inputs["client_cert_role"], false))
		collect(helpers.AttrValueSetString(&o.DefaultAuthPath, inputs["default_auth_path"], false))
		collect(helpers.AttrValueSetString(&o.KubernetesRole, inputs["kubernetes_role"], false))
		collect(helpers.AttrValueSetString(&o.Namespace, inputs["namespace"], false))
		collect(helpers.AttrValueSetString(&o.RoleId, inputs["role_id"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialHashivaultSshTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialHashivaultSshTypeLookup = framework.NewCredentialTypeLookup()

type credentialHashivaultSshResource = framework.GenericResource[credentialHashivaultSshTerraformModel, credentialHashivaultSshBodyRequestModel, *credentialHashivaultSshTerraformModel]

// NewCredentialHashivaultSshResource constructs the typed HashiCorp Vault Signed SSH credential resource.
// The credential_type ID is resolved by namespace (hashivault_ssh) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialHashivaultSshResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["cacert"] = schema.StringAttribute{
		Description: "The CA certificate used to verify the SSL certificate of the Vault server.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["client_cert_private"] = schema.StringAttribute{
		Description: "The certificate private key used for TLS client authentication.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["client_cert_public"] = schema.StringAttribute{
		Description: "The PEM-encoded client certificate used for TLS client authentication. This should include the certificate and any intermediate certififcates.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["client_cert_role"] = schema.StringAttribute{
		Description: "The role configured in Hashicorp Vault for TLS client authentication. If not provided, Hashicorp Vault may assign roles based on the certificate used.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["default_auth_path"] = schema.StringAttribute{
		Description: "The Authentication path to use if one isn't provided in the metadata when linking to an input field. Defaults to 'approle'. AWX defaults this to \"approle\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["kubernetes_role"] = schema.StringAttribute{
		Description: "The Role for Kubernetes Authentication. This is the named role, configured in Vault server, for AWX pod auth policies. see https://www.vaultproject.io/docs/auth/kubernetes#configuration.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["namespace"] = schema.StringAttribute{
		Description: "Name of the namespace to use when authenticate and retrieve secrets.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["password"] = schema.StringAttribute{
		Description: "Password for user authentication.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["role_id"] = schema.StringAttribute{
		Description: "The Role ID for AppRole Authentication.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["secret_id"] = schema.StringAttribute{
		Description: "The Secret ID for AppRole Authentication.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["token"] = schema.StringAttribute{
		Description: "The access token used to authenticate to the Vault server.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["url"] = schema.StringAttribute{
		Description: "The URL to the HashiCorp Vault.",
		Required:    true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username for user authentication.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	return &credentialHashivaultSshResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_hashivault_ssh", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialHashivaultSshTerraformModel, credentialHashivaultSshBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `HashiCorp Vault Signed SSH` (hashivault_ssh) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.hashivault_ssh.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialHashivaultSshTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialHashivaultSsh,
			OnConfigure: credentialHashivaultSshTypeLookup.OnConfigure("hashivault_ssh"),
			MutateBody: func(plan *credentialHashivaultSshTerraformModel, body *credentialHashivaultSshBodyRequestModel) {
				body.CredentialType = credentialHashivaultSshTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialHashivaultSshTerraformModel, body *credentialHashivaultSshBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialHashivaultSshTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialHashivaultSshTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialHashivaultSsh",
		},
	}
}

type credentialHashivaultSshDataSource = framework.GenericDataSource[credentialHashivaultSshDataSourceTerraformModel, *credentialHashivaultSshDataSourceTerraformModel]

// NewCredentialHashivaultSshDataSource constructs the typed HashiCorp Vault Signed SSH credential data source.
func NewCredentialHashivaultSshDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["cacert"] = dschema.StringAttribute{
		Description: "The CA certificate used to verify the SSL certificate of the Vault server.",
		Computed:    true,
	}
	attrs["client_cert_public"] = dschema.StringAttribute{
		Description: "The PEM-encoded client certificate used for TLS client authentication. This should include the certificate and any intermediate certififcates.",
		Computed:    true,
	}
	attrs["client_cert_role"] = dschema.StringAttribute{
		Description: "The role configured in Hashicorp Vault for TLS client authentication. If not provided, Hashicorp Vault may assign roles based on the certificate used.",
		Computed:    true,
	}
	attrs["default_auth_path"] = dschema.StringAttribute{
		Description: "The Authentication path to use if one isn't provided in the metadata when linking to an input field. Defaults to 'approle'. AWX defaults this to \"approle\" when unset.",
		Computed:    true,
	}
	attrs["kubernetes_role"] = dschema.StringAttribute{
		Description: "The Role for Kubernetes Authentication. This is the named role, configured in Vault server, for AWX pod auth policies. see https://www.vaultproject.io/docs/auth/kubernetes#configuration.",
		Computed:    true,
	}
	attrs["namespace"] = dschema.StringAttribute{
		Description: "Name of the namespace to use when authenticate and retrieve secrets.",
		Computed:    true,
	}
	attrs["role_id"] = dschema.StringAttribute{
		Description: "The Role ID for AppRole Authentication.",
		Computed:    true,
	}
	attrs["url"] = dschema.StringAttribute{
		Description: "The URL to the HashiCorp Vault.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username for user authentication.",
		Computed:    true,
	}
	return &credentialHashivaultSshDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_hashivault_ssh", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialHashivaultSshDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `HashiCorp Vault Signed SSH` (hashivault_ssh) credential by ID or name.",
				Attributes:          attrs,
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "/?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			OnConfigure:  credentialHashivaultSshTypeLookup.OnConfigure("hashivault_ssh"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialHashivaultSsh",
		},
	}
}
