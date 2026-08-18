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

// credentialRhvTerraformModel exposes the typed AWX Red Hat Virtualization
// credential (credential_rhv) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialRhvTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	CaFile         types.String `tfsdk:"ca_file" json:"-"`
	Host           types.String `tfsdk:"host" json:"-"`
	Password       types.String `tfsdk:"password" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialRhvTerraformModel) Clone() credentialRhvTerraformModel {
	return *o
}

type credentialRhvBodyRequestModel struct {
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
func (o *credentialRhvTerraformModel) BodyRequest() *credentialRhvBodyRequestModel {
	req := &credentialRhvBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.CaFile.IsNull() && !o.CaFile.IsUnknown() {
		inputs["ca_file"] = o.CaFile.ValueString()
	}
	if !o.Host.IsNull() && !o.Host.IsUnknown() {
		inputs["host"] = o.Host.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
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
func (o *credentialRhvTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.CaFile, inputs["ca_file"], false))
		collect(helpers.AttrValueSetString(&o.Host, inputs["host"], false))
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialRhv reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialRhv(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialRhvTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Password.IsNull() || orig.Password.IsUnknown() {
			state.Password = types.StringNull()
		} else {
			state.Password = orig.Password
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
	}
	return nil
}

// credentialRhvDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialRhvDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	CaFile         types.String `tfsdk:"ca_file" json:"-"`
	Host           types.String `tfsdk:"host" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialRhvDataSourceTerraformModel) Clone() credentialRhvDataSourceTerraformModel {
	return *o
}

func (o *credentialRhvDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.CaFile, inputs["ca_file"], false))
		collect(helpers.AttrValueSetString(&o.Host, inputs["host"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialRhvTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialRhvTypeLookup = framework.NewCredentialTypeLookup()

type credentialRhvResource = framework.GenericResource[credentialRhvTerraformModel, credentialRhvBodyRequestModel, *credentialRhvTerraformModel]

// NewCredentialRhvResource constructs the typed Red Hat Virtualization credential resource.
// The credential_type ID is resolved by namespace (rhv) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialRhvResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["ca_file"] = schema.StringAttribute{
		Description: "Absolute file path to the CA file to use (optional).",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["host"] = schema.StringAttribute{
		Description: "The host to authenticate with.",
		Required:    true,
	}
	attrs["password"] = schema.StringAttribute{
		Description: "Password.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Required:    true,
	}
	return &credentialRhvResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_rhv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialRhvTerraformModel, credentialRhvBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Red Hat Virtualization` (rhv) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.rhv.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialRhvTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialRhv,
			OnConfigure: credentialRhvTypeLookup.OnConfigure("rhv"),
			MutateBody: func(plan *credentialRhvTerraformModel, body *credentialRhvBodyRequestModel) {
				body.CredentialType = credentialRhvTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialRhvTerraformModel, body *credentialRhvBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialRhvTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialRhvTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialRhv",
		},
	}
}

type credentialRhvDataSource = framework.GenericDataSource[credentialRhvDataSourceTerraformModel, *credentialRhvDataSourceTerraformModel]

// NewCredentialRhvDataSource constructs the typed Red Hat Virtualization credential data source.
func NewCredentialRhvDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["ca_file"] = dschema.StringAttribute{
		Description: "Absolute file path to the CA file to use (optional).",
		Computed:    true,
	}
	attrs["host"] = dschema.StringAttribute{
		Description: "The host to authenticate with.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	return &credentialRhvDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_rhv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialRhvDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Red Hat Virtualization` (rhv) credential by ID or name.",
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
			OnConfigure:  credentialRhvTypeLookup.OnConfigure("rhv"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialRhv",
		},
	}
}
