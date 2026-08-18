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

// credentialThycoticTssTerraformModel exposes the typed AWX Thycotic Secret Server
// credential (credential_thycotic_tss) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialThycoticTssTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Domain         types.String `tfsdk:"domain" json:"-"`
	Password       types.String `tfsdk:"password" json:"-"`
	ServerUrl      types.String `tfsdk:"server_url" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialThycoticTssTerraformModel) Clone() credentialThycoticTssTerraformModel {
	return *o
}

type credentialThycoticTssBodyRequestModel struct {
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
func (o *credentialThycoticTssTerraformModel) BodyRequest() *credentialThycoticTssBodyRequestModel {
	req := &credentialThycoticTssBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Domain.IsNull() && !o.Domain.IsUnknown() {
		inputs["domain"] = o.Domain.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.ServerUrl.IsNull() && !o.ServerUrl.IsUnknown() {
		inputs["server_url"] = o.ServerUrl.ValueString()
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
func (o *credentialThycoticTssTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Domain, inputs["domain"], false))
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.ServerUrl, inputs["server_url"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialThycoticTss reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialThycoticTss(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialThycoticTssTerraformModel) error {
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

// credentialThycoticTssDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialThycoticTssDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Domain         types.String `tfsdk:"domain" json:"-"`
	ServerUrl      types.String `tfsdk:"server_url" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialThycoticTssDataSourceTerraformModel) Clone() credentialThycoticTssDataSourceTerraformModel {
	return *o
}

func (o *credentialThycoticTssDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Domain, inputs["domain"], false))
		collect(helpers.AttrValueSetString(&o.ServerUrl, inputs["server_url"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialThycoticTssTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialThycoticTssTypeLookup = framework.NewCredentialTypeLookup()

type credentialThycoticTssResource = framework.GenericResource[credentialThycoticTssTerraformModel, credentialThycoticTssBodyRequestModel, *credentialThycoticTssTerraformModel]

// NewCredentialThycoticTssResource constructs the typed Thycotic Secret Server credential resource.
// The credential_type ID is resolved by namespace (thycotic_tss) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialThycoticTssResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["domain"] = schema.StringAttribute{
		Description: "The (Application) user domain.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["password"] = schema.StringAttribute{
		Description: "The corresponding password.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["server_url"] = schema.StringAttribute{
		Description: "The Base URL of Secret Server e.g. https://myserver/SecretServer or https://mytenant.secretservercloud.com.",
		Required:    true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "The (Application) user username.",
		Required:    true,
	}
	return &credentialThycoticTssResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_thycotic_tss", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialThycoticTssTerraformModel, credentialThycoticTssBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Thycotic Secret Server` (thycotic_tss) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.thycotic_tss.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialThycoticTssTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialThycoticTss,
			OnConfigure: credentialThycoticTssTypeLookup.OnConfigure("thycotic_tss"),
			MutateBody: func(plan *credentialThycoticTssTerraformModel, body *credentialThycoticTssBodyRequestModel) {
				body.CredentialType = credentialThycoticTssTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialThycoticTssTerraformModel, body *credentialThycoticTssBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialThycoticTssTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialThycoticTssTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialThycoticTss",
		},
	}
}

type credentialThycoticTssDataSource = framework.GenericDataSource[credentialThycoticTssDataSourceTerraformModel, *credentialThycoticTssDataSourceTerraformModel]

// NewCredentialThycoticTssDataSource constructs the typed Thycotic Secret Server credential data source.
func NewCredentialThycoticTssDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["domain"] = dschema.StringAttribute{
		Description: "The (Application) user domain.",
		Computed:    true,
	}
	attrs["server_url"] = dschema.StringAttribute{
		Description: "The Base URL of Secret Server e.g. https://myserver/SecretServer or https://mytenant.secretservercloud.com.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "The (Application) user username.",
		Computed:    true,
	}
	return &credentialThycoticTssDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_thycotic_tss", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialThycoticTssDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Thycotic Secret Server` (thycotic_tss) credential by ID or name.",
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
			OnConfigure:  credentialThycoticTssTypeLookup.OnConfigure("thycotic_tss"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialThycoticTss",
		},
	}
}
