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

// credentialVaultTerraformModel exposes the typed AWX Vault
// credential (credential_vault) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialVaultTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	VaultId        types.String `tfsdk:"vault_id" json:"-"`
	VaultPassword  types.String `tfsdk:"vault_password" json:"-"`
}

func (o *credentialVaultTerraformModel) Clone() credentialVaultTerraformModel {
	return *o
}

type credentialVaultBodyRequestModel struct {
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
func (o *credentialVaultTerraformModel) BodyRequest() *credentialVaultBodyRequestModel {
	req := &credentialVaultBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.VaultId.IsNull() && !o.VaultId.IsUnknown() {
		inputs["vault_id"] = o.VaultId.ValueString()
	}
	if !o.VaultPassword.IsNull() && !o.VaultPassword.IsUnknown() {
		inputs["vault_password"] = o.VaultPassword.ValueString()
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
func (o *credentialVaultTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.VaultId, inputs["vault_id"], false))
		collect(helpers.AttrValueSetString(&o.VaultPassword, inputs["vault_password"], false))
	}
	return diags, nil
}

// hookCredentialVault reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialVault(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialVaultTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.VaultPassword.IsNull() || orig.VaultPassword.IsUnknown() {
			state.VaultPassword = types.StringNull()
		} else {
			state.VaultPassword = orig.VaultPassword
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.VaultPassword, state.VaultPassword); subbed {
			state.VaultPassword = v
		}
	}
	return nil
}

// credentialVaultDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialVaultDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	VaultId        types.String `tfsdk:"vault_id" json:"-"`
}

func (o *credentialVaultDataSourceTerraformModel) Clone() credentialVaultDataSourceTerraformModel {
	return *o
}

func (o *credentialVaultDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.VaultId, inputs["vault_id"], false))
	}
	return diags, nil
}

// credentialVaultTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialVaultTypeLookup = framework.NewCredentialTypeLookup()

type credentialVaultResource = framework.GenericResource[credentialVaultTerraformModel, credentialVaultBodyRequestModel, *credentialVaultTerraformModel]

// NewCredentialVaultResource constructs the typed Vault credential resource.
// The credential_type ID is resolved by namespace (vault) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialVaultResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["vault_id"] = schema.StringAttribute{
		Description: "Specify an (optional) Vault ID. This is equivalent to specifying the --vault-id Ansible parameter for providing multiple Vault passwords.  Note: this feature only works in Ansible 2.4+. Changing this forces a new credential to be created.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
			stringplanmodifier.RequiresReplace(),
		},
	}
	attrs["vault_password"] = schema.StringAttribute{
		Description: "Vault Password.",
		Required:    true,
		Sensitive:   true,
	}
	return &credentialVaultResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_vault", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialVaultTerraformModel, credentialVaultBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Vault` (vault) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.vault.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialVaultTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialVault,
			OnConfigure: credentialVaultTypeLookup.OnConfigure("vault"),
			MutateBody: func(plan *credentialVaultTerraformModel, body *credentialVaultBodyRequestModel) {
				body.CredentialType = credentialVaultTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialVaultTerraformModel, body *credentialVaultBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialVaultTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialVaultTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialVault",
		},
	}
}

type credentialVaultDataSource = framework.GenericDataSource[credentialVaultDataSourceTerraformModel, *credentialVaultDataSourceTerraformModel]

// NewCredentialVaultDataSource constructs the typed Vault credential data source.
func NewCredentialVaultDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["vault_id"] = dschema.StringAttribute{
		Description: "Specify an (optional) Vault ID. This is equivalent to specifying the --vault-id Ansible parameter for providing multiple Vault passwords.  Note: this feature only works in Ansible 2.4+. Changing this forces a new credential to be created.",
		Computed:    true,
	}
	return &credentialVaultDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_vault", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialVaultDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Vault` (vault) credential by ID or name.",
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
			OnConfigure:  credentialVaultTypeLookup.OnConfigure("vault"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialVault",
		},
	}
}
