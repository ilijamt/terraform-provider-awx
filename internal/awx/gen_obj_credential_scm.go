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

// credentialScmTerraformModel exposes the typed AWX Source Control
// credential (credential_scm) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialScmTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Password       types.String `tfsdk:"password" json:"-"`
	SshKeyData     types.String `tfsdk:"ssh_key_data" json:"-"`
	SshKeyUnlock   types.String `tfsdk:"ssh_key_unlock" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialScmTerraformModel) Clone() credentialScmTerraformModel {
	return *o
}

type credentialScmBodyRequestModel struct {
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
func (o *credentialScmTerraformModel) BodyRequest() *credentialScmBodyRequestModel {
	req := &credentialScmBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.SshKeyData.IsNull() && !o.SshKeyData.IsUnknown() {
		inputs["ssh_key_data"] = o.SshKeyData.ValueString()
	}
	if !o.SshKeyUnlock.IsNull() && !o.SshKeyUnlock.IsUnknown() {
		inputs["ssh_key_unlock"] = o.SshKeyUnlock.ValueString()
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
func (o *credentialScmTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.SshKeyData, inputs["ssh_key_data"], false))
		collect(helpers.AttrValueSetString(&o.SshKeyUnlock, inputs["ssh_key_unlock"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialScm reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialScm(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialScmTerraformModel) error {
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
		if orig.SshKeyData.IsNull() || orig.SshKeyData.IsUnknown() {
			state.SshKeyData = types.StringNull()
		} else {
			state.SshKeyData = orig.SshKeyData
		}
		if orig.SshKeyUnlock.IsNull() || orig.SshKeyUnlock.IsUnknown() {
			state.SshKeyUnlock = types.StringNull()
		} else {
			state.SshKeyUnlock = orig.SshKeyUnlock
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.SshKeyData, state.SshKeyData); subbed {
			state.SshKeyData = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.SshKeyUnlock, state.SshKeyUnlock); subbed {
			state.SshKeyUnlock = v
		}
	}
	return nil
}

// credentialScmDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialScmDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialScmDataSourceTerraformModel) Clone() credentialScmDataSourceTerraformModel {
	return *o
}

func (o *credentialScmDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialScmTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialScmTypeLookup = framework.NewCredentialTypeLookup()

type credentialScmResource = framework.GenericResource[credentialScmTerraformModel, credentialScmBodyRequestModel, *credentialScmTerraformModel]

// NewCredentialScmResource constructs the typed Source Control credential resource.
// The credential_type ID is resolved by namespace (scm) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialScmResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["password"] = schema.StringAttribute{
		Description: "Password.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["ssh_key_data"] = schema.StringAttribute{
		Description: "SCM Private Key.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["ssh_key_unlock"] = schema.StringAttribute{
		Description: "Private Key Passphrase.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	return &credentialScmResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_scm", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialScmTerraformModel, credentialScmBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Source Control` (scm) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.scm.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialScmTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialScm,
			OnConfigure: credentialScmTypeLookup.OnConfigure("scm"),
			MutateBody: func(plan *credentialScmTerraformModel, body *credentialScmBodyRequestModel) {
				body.CredentialType = credentialScmTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialScmTerraformModel, body *credentialScmBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialScmTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialScmTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialScm",
		},
	}
}

type credentialScmDataSource = framework.GenericDataSource[credentialScmDataSourceTerraformModel, *credentialScmDataSourceTerraformModel]

// NewCredentialScmDataSource constructs the typed Source Control credential data source.
func NewCredentialScmDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	return &credentialScmDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_scm", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialScmDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Source Control` (scm) credential by ID or name.",
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
			OnConfigure:  credentialScmTypeLookup.OnConfigure("scm"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialScm",
		},
	}
}
