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

// credentialAimTerraformModel exposes the typed AWX CyberArk Central Credential Provider Lookup
// credential (credential_aim) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialAimTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	AppId          types.String `tfsdk:"app_id" json:"-"`
	ClientCert     types.String `tfsdk:"client_cert" json:"-"`
	ClientKey      types.String `tfsdk:"client_key" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
	Verify         types.Bool   `tfsdk:"verify" json:"-"`
	WebserviceId   types.String `tfsdk:"webservice_id" json:"-"`
}

func (o *credentialAimTerraformModel) Clone() credentialAimTerraformModel {
	return *o
}

type credentialAimBodyRequestModel struct {
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
func (o *credentialAimTerraformModel) BodyRequest() *credentialAimBodyRequestModel {
	req := &credentialAimBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.AppId.IsNull() && !o.AppId.IsUnknown() {
		inputs["app_id"] = o.AppId.ValueString()
	}
	if !o.ClientCert.IsNull() && !o.ClientCert.IsUnknown() {
		inputs["client_cert"] = o.ClientCert.ValueString()
	}
	if !o.ClientKey.IsNull() && !o.ClientKey.IsUnknown() {
		inputs["client_key"] = o.ClientKey.ValueString()
	}
	if !o.Url.IsNull() && !o.Url.IsUnknown() {
		inputs["url"] = o.Url.ValueString()
	}
	if !o.Verify.IsNull() && !o.Verify.IsUnknown() {
		inputs["verify"] = o.Verify.ValueBool()
	}
	if !o.WebserviceId.IsNull() && !o.WebserviceId.IsUnknown() {
		inputs["webservice_id"] = o.WebserviceId.ValueString()
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
func (o *credentialAimTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AppId, inputs["app_id"], false))
		collect(helpers.AttrValueSetString(&o.ClientCert, inputs["client_cert"], false))
		collect(helpers.AttrValueSetString(&o.ClientKey, inputs["client_key"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
		collect(helpers.AttrValueSetBool(&o.Verify, inputs["verify"]))
		collect(helpers.AttrValueSetString(&o.WebserviceId, inputs["webservice_id"], false))
	}
	return diags, nil
}

// hookCredentialAim reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialAim(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialAimTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.AppId.IsNull() || orig.AppId.IsUnknown() {
			state.AppId = types.StringNull()
		} else {
			state.AppId = orig.AppId
		}
		if orig.ClientCert.IsNull() || orig.ClientCert.IsUnknown() {
			state.ClientCert = types.StringNull()
		} else {
			state.ClientCert = orig.ClientCert
		}
		if orig.ClientKey.IsNull() || orig.ClientKey.IsUnknown() {
			state.ClientKey = types.StringNull()
		} else {
			state.ClientKey = orig.ClientKey
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.AppId, state.AppId); subbed {
			state.AppId = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.ClientCert, state.ClientCert); subbed {
			state.ClientCert = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.ClientKey, state.ClientKey); subbed {
			state.ClientKey = v
		}
	}
	return nil
}

// credentialAimDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialAimDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Url            types.String `tfsdk:"url" json:"-"`
	Verify         types.Bool   `tfsdk:"verify" json:"-"`
	WebserviceId   types.String `tfsdk:"webservice_id" json:"-"`
}

func (o *credentialAimDataSourceTerraformModel) Clone() credentialAimDataSourceTerraformModel {
	return *o
}

func (o *credentialAimDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
		collect(helpers.AttrValueSetBool(&o.Verify, inputs["verify"]))
		collect(helpers.AttrValueSetString(&o.WebserviceId, inputs["webservice_id"], false))
	}
	return diags, nil
}

// credentialAimTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialAimTypeLookup = framework.NewCredentialTypeLookup()

type credentialAimResource = framework.GenericResource[credentialAimTerraformModel, credentialAimBodyRequestModel, *credentialAimTerraformModel]

// NewCredentialAimResource constructs the typed CyberArk Central Credential Provider Lookup credential resource.
// The credential_type ID is resolved by namespace (aim) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialAimResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["app_id"] = schema.StringAttribute{
		Description: "Application ID.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["client_cert"] = schema.StringAttribute{
		Description: "Client Certificate.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["client_key"] = schema.StringAttribute{
		Description: "Client Key.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["url"] = schema.StringAttribute{
		Description: "CyberArk CCP URL.",
		Required:    true,
	}
	attrs["verify"] = schema.BoolAttribute{
		Description: "Verify SSL Certificates. AWX defaults this to \"true\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["webservice_id"] = schema.StringAttribute{
		Description: "The CCP Web Service ID. Leave blank to default to AIMWebService.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	return &credentialAimResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_aim", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialAimTerraformModel, credentialAimBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `CyberArk Central Credential Provider Lookup` (aim) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.aim.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialAimTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialAim,
			OnConfigure: credentialAimTypeLookup.OnConfigure("aim"),
			MutateBody: func(plan *credentialAimTerraformModel, body *credentialAimBodyRequestModel) {
				body.CredentialType = credentialAimTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialAimTerraformModel, body *credentialAimBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialAimTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialAimTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAim",
		},
	}
}

type credentialAimDataSource = framework.GenericDataSource[credentialAimDataSourceTerraformModel, *credentialAimDataSourceTerraformModel]

// NewCredentialAimDataSource constructs the typed CyberArk Central Credential Provider Lookup credential data source.
func NewCredentialAimDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["url"] = dschema.StringAttribute{
		Description: "CyberArk CCP URL.",
		Computed:    true,
	}
	attrs["verify"] = dschema.BoolAttribute{
		Description: "Verify SSL Certificates. AWX defaults this to \"true\" when unset.",
		Computed:    true,
	}
	attrs["webservice_id"] = dschema.StringAttribute{
		Description: "The CCP Web Service ID. Leave blank to default to AIMWebService.",
		Computed:    true,
	}
	return &credentialAimDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_aim", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialAimDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `CyberArk Central Credential Provider Lookup` (aim) credential by ID or name.",
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
			OnConfigure:  credentialAimTypeLookup.OnConfigure("aim"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialAim",
		},
	}
}
