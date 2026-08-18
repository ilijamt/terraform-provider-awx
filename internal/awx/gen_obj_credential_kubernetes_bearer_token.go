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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialKubernetesBearerTokenTerraformModel exposes the typed AWX OpenShift or Kubernetes API Bearer Token
// credential (credential_kubernetes_bearer_token) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialKubernetesBearerTokenTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	BearerToken    types.String `tfsdk:"bearer_token" json:"-"`
	Host           types.String `tfsdk:"host" json:"-"`
	SslCaCert      types.String `tfsdk:"ssl_ca_cert" json:"-"`
	VerifySsl      types.Bool   `tfsdk:"verify_ssl" json:"-"`
}

func (o *credentialKubernetesBearerTokenTerraformModel) Clone() credentialKubernetesBearerTokenTerraformModel {
	return *o
}

type credentialKubernetesBearerTokenBodyRequestModel struct {
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
func (o *credentialKubernetesBearerTokenTerraformModel) BodyRequest() *credentialKubernetesBearerTokenBodyRequestModel {
	req := &credentialKubernetesBearerTokenBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.BearerToken.IsNull() && !o.BearerToken.IsUnknown() {
		inputs["bearer_token"] = o.BearerToken.ValueString()
	}
	if !o.Host.IsNull() && !o.Host.IsUnknown() {
		inputs["host"] = o.Host.ValueString()
	}
	if !o.SslCaCert.IsNull() && !o.SslCaCert.IsUnknown() {
		inputs["ssl_ca_cert"] = o.SslCaCert.ValueString()
	}
	if !o.VerifySsl.IsNull() && !o.VerifySsl.IsUnknown() {
		inputs["verify_ssl"] = o.VerifySsl.ValueBool()
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
func (o *credentialKubernetesBearerTokenTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.BearerToken, inputs["bearer_token"], false))
		collect(helpers.AttrValueSetString(&o.Host, inputs["host"], false))
		collect(helpers.AttrValueSetString(&o.SslCaCert, inputs["ssl_ca_cert"], false))
		collect(helpers.AttrValueSetBool(&o.VerifySsl, inputs["verify_ssl"]))
	}
	return diags, nil
}

// hookCredentialKubernetesBearerToken reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialKubernetesBearerToken(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialKubernetesBearerTokenTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.BearerToken.IsNull() || orig.BearerToken.IsUnknown() {
			state.BearerToken = types.StringNull()
		} else {
			state.BearerToken = orig.BearerToken
		}
		if orig.SslCaCert.IsNull() || orig.SslCaCert.IsUnknown() {
			state.SslCaCert = types.StringNull()
		} else {
			state.SslCaCert = orig.SslCaCert
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.BearerToken, state.BearerToken); subbed {
			state.BearerToken = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.SslCaCert, state.SslCaCert); subbed {
			state.SslCaCert = v
		}
	}
	return nil
}

// credentialKubernetesBearerTokenDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialKubernetesBearerTokenDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Host           types.String `tfsdk:"host" json:"-"`
	VerifySsl      types.Bool   `tfsdk:"verify_ssl" json:"-"`
}

func (o *credentialKubernetesBearerTokenDataSourceTerraformModel) Clone() credentialKubernetesBearerTokenDataSourceTerraformModel {
	return *o
}

func (o *credentialKubernetesBearerTokenDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Host, inputs["host"], false))
		collect(helpers.AttrValueSetBool(&o.VerifySsl, inputs["verify_ssl"]))
	}
	return diags, nil
}

// credentialKubernetesBearerTokenTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialKubernetesBearerTokenTypeLookup = framework.NewCredentialTypeLookup()

type credentialKubernetesBearerTokenResource = framework.GenericResource[credentialKubernetesBearerTokenTerraformModel, credentialKubernetesBearerTokenBodyRequestModel, *credentialKubernetesBearerTokenTerraformModel]

// NewCredentialKubernetesBearerTokenResource constructs the typed OpenShift or Kubernetes API Bearer Token credential resource.
// The credential_type ID is resolved by namespace (kubernetes_bearer_token) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialKubernetesBearerTokenResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["bearer_token"] = schema.StringAttribute{
		Description: "API authentication bearer token.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["host"] = schema.StringAttribute{
		Description: "The OpenShift or Kubernetes API Endpoint to authenticate with.",
		Required:    true,
	}
	attrs["ssl_ca_cert"] = schema.StringAttribute{
		Description: "Certificate Authority data.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["verify_ssl"] = schema.BoolAttribute{
		Description: "Verify SSL. AWX defaults this to \"true\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	}
	return &credentialKubernetesBearerTokenResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_kubernetes_bearer_token", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialKubernetesBearerTokenTerraformModel, credentialKubernetesBearerTokenBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `OpenShift or Kubernetes API Bearer Token` (kubernetes_bearer_token) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.kubernetes_bearer_token.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialKubernetesBearerTokenTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialKubernetesBearerToken,
			OnConfigure: credentialKubernetesBearerTokenTypeLookup.OnConfigure("kubernetes_bearer_token"),
			MutateBody: func(plan *credentialKubernetesBearerTokenTerraformModel, body *credentialKubernetesBearerTokenBodyRequestModel) {
				body.CredentialType = credentialKubernetesBearerTokenTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialKubernetesBearerTokenTerraformModel, body *credentialKubernetesBearerTokenBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialKubernetesBearerTokenTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialKubernetesBearerTokenTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialKubernetesBearerToken",
		},
	}
}

type credentialKubernetesBearerTokenDataSource = framework.GenericDataSource[credentialKubernetesBearerTokenDataSourceTerraformModel, *credentialKubernetesBearerTokenDataSourceTerraformModel]

// NewCredentialKubernetesBearerTokenDataSource constructs the typed OpenShift or Kubernetes API Bearer Token credential data source.
func NewCredentialKubernetesBearerTokenDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["host"] = dschema.StringAttribute{
		Description: "The OpenShift or Kubernetes API Endpoint to authenticate with.",
		Computed:    true,
	}
	attrs["verify_ssl"] = dschema.BoolAttribute{
		Description: "Verify SSL. AWX defaults this to \"true\" when unset.",
		Computed:    true,
	}
	return &credentialKubernetesBearerTokenDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_kubernetes_bearer_token", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialKubernetesBearerTokenDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `OpenShift or Kubernetes API Bearer Token` (kubernetes_bearer_token) credential by ID or name.",
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
			OnConfigure:  credentialKubernetesBearerTokenTypeLookup.OnConfigure("kubernetes_bearer_token"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialKubernetesBearerToken",
		},
	}
}
