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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialThycoticDsvTerraformModel exposes the typed AWX Thycotic DevOps Secrets Vault
// credential (credential_thycotic_dsv) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialThycoticDsvTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	ClientId       types.String `tfsdk:"client_id" json:"-"`
	ClientSecret   types.String `tfsdk:"client_secret" json:"-"`
	Tenant         types.String `tfsdk:"tenant" json:"-"`
	Tld            types.String `tfsdk:"tld" json:"-"`
}

func (o *credentialThycoticDsvTerraformModel) Clone() credentialThycoticDsvTerraformModel {
	return *o
}

type credentialThycoticDsvBodyRequestModel struct {
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
func (o *credentialThycoticDsvTerraformModel) BodyRequest() *credentialThycoticDsvBodyRequestModel {
	req := &credentialThycoticDsvBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.ClientId.IsNull() && !o.ClientId.IsUnknown() {
		inputs["client_id"] = o.ClientId.ValueString()
	}
	if !o.ClientSecret.IsNull() && !o.ClientSecret.IsUnknown() {
		inputs["client_secret"] = o.ClientSecret.ValueString()
	}
	if !o.Tenant.IsNull() && !o.Tenant.IsUnknown() {
		inputs["tenant"] = o.Tenant.ValueString()
	}
	if !o.Tld.IsNull() && !o.Tld.IsUnknown() {
		inputs["tld"] = o.Tld.ValueString()
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
func (o *credentialThycoticDsvTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.ClientId, inputs["client_id"], false))
		collect(helpers.AttrValueSetString(&o.ClientSecret, inputs["client_secret"], false))
		collect(helpers.AttrValueSetString(&o.Tenant, inputs["tenant"], false))
		collect(helpers.AttrValueSetString(&o.Tld, inputs["tld"], false))
	}
	return diags, nil
}

// hookCredentialThycoticDsv reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialThycoticDsv(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialThycoticDsvTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.ClientSecret.IsNull() || orig.ClientSecret.IsUnknown() {
			state.ClientSecret = types.StringNull()
		} else {
			state.ClientSecret = orig.ClientSecret
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.ClientSecret, state.ClientSecret); subbed {
			state.ClientSecret = v
		}
	}
	return nil
}

// credentialThycoticDsvDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialThycoticDsvDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	ClientId       types.String `tfsdk:"client_id" json:"-"`
	Tenant         types.String `tfsdk:"tenant" json:"-"`
	Tld            types.String `tfsdk:"tld" json:"-"`
}

func (o *credentialThycoticDsvDataSourceTerraformModel) Clone() credentialThycoticDsvDataSourceTerraformModel {
	return *o
}

func (o *credentialThycoticDsvDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.ClientId, inputs["client_id"], false))
		collect(helpers.AttrValueSetString(&o.Tenant, inputs["tenant"], false))
		collect(helpers.AttrValueSetString(&o.Tld, inputs["tld"], false))
	}
	return diags, nil
}

// credentialThycoticDsvTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialThycoticDsvTypeLookup = framework.NewCredentialTypeLookup()

type credentialThycoticDsvResource = framework.GenericResource[credentialThycoticDsvTerraformModel, credentialThycoticDsvBodyRequestModel, *credentialThycoticDsvTerraformModel]

// NewCredentialThycoticDsvResource constructs the typed Thycotic DevOps Secrets Vault credential resource.
// The credential_type ID is resolved by namespace (thycotic_dsv) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialThycoticDsvResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["client_id"] = schema.StringAttribute{
		Description: "Client ID.",
		Required:    true,
	}
	attrs["client_secret"] = schema.StringAttribute{
		Description: "Client Secret.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["tenant"] = schema.StringAttribute{
		Description: "The tenant e.g. \"ex\" when the URL is https://ex.secretsvaultcloud.com.",
		Required:    true,
	}
	attrs["tld"] = schema.StringAttribute{
		Description: "The TLD of the tenant e.g. \"com\" when the URL is https://ex.secretsvaultcloud.com. Allowed values: \"ca\", \"com\", \"com.au\", \"eu\". AWX defaults this to \"com\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Validators: []validator.String{
			stringvalidator.OneOf("ca", "com", "com.au", "eu"),
		},
	}
	return &credentialThycoticDsvResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_thycotic_dsv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialThycoticDsvTerraformModel, credentialThycoticDsvBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Thycotic DevOps Secrets Vault` (thycotic_dsv) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.thycotic_dsv.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialThycoticDsvTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialThycoticDsv,
			OnConfigure: credentialThycoticDsvTypeLookup.OnConfigure("thycotic_dsv"),
			MutateBody: func(plan *credentialThycoticDsvTerraformModel, body *credentialThycoticDsvBodyRequestModel) {
				body.CredentialType = credentialThycoticDsvTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialThycoticDsvTerraformModel, body *credentialThycoticDsvBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialThycoticDsvTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialThycoticDsvTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialThycoticDsv",
		},
	}
}

type credentialThycoticDsvDataSource = framework.GenericDataSource[credentialThycoticDsvDataSourceTerraformModel, *credentialThycoticDsvDataSourceTerraformModel]

// NewCredentialThycoticDsvDataSource constructs the typed Thycotic DevOps Secrets Vault credential data source.
func NewCredentialThycoticDsvDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["client_id"] = dschema.StringAttribute{
		Description: "Client ID.",
		Computed:    true,
	}
	attrs["tenant"] = dschema.StringAttribute{
		Description: "The tenant e.g. \"ex\" when the URL is https://ex.secretsvaultcloud.com.",
		Computed:    true,
	}
	attrs["tld"] = dschema.StringAttribute{
		Description: "The TLD of the tenant e.g. \"com\" when the URL is https://ex.secretsvaultcloud.com. Allowed values: \"ca\", \"com\", \"com.au\", \"eu\". AWX defaults this to \"com\" when unset.",
		Computed:    true,
	}
	return &credentialThycoticDsvDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_thycotic_dsv", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialThycoticDsvDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Thycotic DevOps Secrets Vault` (thycotic_dsv) credential by ID or name.",
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
			OnConfigure:  credentialThycoticDsvTypeLookup.OnConfigure("thycotic_dsv"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialThycoticDsv",
		},
	}
}
