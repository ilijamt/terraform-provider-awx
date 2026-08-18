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

// credentialConjurTerraformModel exposes the typed AWX CyberArk Conjur Secrets Manager Lookup
// credential (credential_conjur) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialConjurTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Account        types.String `tfsdk:"account" json:"-"`
	ApiKey         types.String `tfsdk:"api_key" json:"-"`
	Cacert         types.String `tfsdk:"cacert" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialConjurTerraformModel) Clone() credentialConjurTerraformModel {
	return *o
}

type credentialConjurBodyRequestModel struct {
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
func (o *credentialConjurTerraformModel) BodyRequest() *credentialConjurBodyRequestModel {
	req := &credentialConjurBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Account.IsNull() && !o.Account.IsUnknown() {
		inputs["account"] = o.Account.ValueString()
	}
	if !o.ApiKey.IsNull() && !o.ApiKey.IsUnknown() {
		inputs["api_key"] = o.ApiKey.ValueString()
	}
	if !o.Cacert.IsNull() && !o.Cacert.IsUnknown() {
		inputs["cacert"] = o.Cacert.ValueString()
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
func (o *credentialConjurTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Account, inputs["account"], false))
		collect(helpers.AttrValueSetString(&o.ApiKey, inputs["api_key"], false))
		collect(helpers.AttrValueSetString(&o.Cacert, inputs["cacert"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialConjur reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialConjur(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialConjurTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.ApiKey.IsNull() || orig.ApiKey.IsUnknown() {
			state.ApiKey = types.StringNull()
		} else {
			state.ApiKey = orig.ApiKey
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.ApiKey, state.ApiKey); subbed {
			state.ApiKey = v
		}
	}
	return nil
}

// credentialConjurDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialConjurDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Account        types.String `tfsdk:"account" json:"-"`
	Cacert         types.String `tfsdk:"cacert" json:"-"`
	Url            types.String `tfsdk:"url" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialConjurDataSourceTerraformModel) Clone() credentialConjurDataSourceTerraformModel {
	return *o
}

func (o *credentialConjurDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Account, inputs["account"], false))
		collect(helpers.AttrValueSetString(&o.Cacert, inputs["cacert"], false))
		collect(helpers.AttrValueSetString(&o.Url, inputs["url"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialConjurTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialConjurTypeLookup = framework.NewCredentialTypeLookup()

type credentialConjurResource = framework.GenericResource[credentialConjurTerraformModel, credentialConjurBodyRequestModel, *credentialConjurTerraformModel]

// NewCredentialConjurResource constructs the typed CyberArk Conjur Secrets Manager Lookup credential resource.
// The credential_type ID is resolved by namespace (conjur) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialConjurResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["account"] = schema.StringAttribute{
		Description: "Account.",
		Required:    true,
	}
	attrs["api_key"] = schema.StringAttribute{
		Description: "API Key.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["cacert"] = schema.StringAttribute{
		Description: "Public Key Certificate.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["url"] = schema.StringAttribute{
		Description: "Conjur URL.",
		Required:    true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Required:    true,
	}
	return &credentialConjurResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_conjur", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialConjurTerraformModel, credentialConjurBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `CyberArk Conjur Secrets Manager Lookup` (conjur) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.conjur.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialConjurTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialConjur,
			OnConfigure: credentialConjurTypeLookup.OnConfigure("conjur"),
			MutateBody: func(plan *credentialConjurTerraformModel, body *credentialConjurBodyRequestModel) {
				body.CredentialType = credentialConjurTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialConjurTerraformModel, body *credentialConjurBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialConjurTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialConjurTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialConjur",
		},
	}
}

type credentialConjurDataSource = framework.GenericDataSource[credentialConjurDataSourceTerraformModel, *credentialConjurDataSourceTerraformModel]

// NewCredentialConjurDataSource constructs the typed CyberArk Conjur Secrets Manager Lookup credential data source.
func NewCredentialConjurDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["account"] = dschema.StringAttribute{
		Description: "Account.",
		Computed:    true,
	}
	attrs["cacert"] = dschema.StringAttribute{
		Description: "Public Key Certificate.",
		Computed:    true,
	}
	attrs["url"] = dschema.StringAttribute{
		Description: "Conjur URL.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	return &credentialConjurDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_conjur", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialConjurDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `CyberArk Conjur Secrets Manager Lookup` (conjur) credential by ID or name.",
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
			OnConfigure:  credentialConjurTypeLookup.OnConfigure("conjur"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialConjur",
		},
	}
}
