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

// credentialGceTerraformModel exposes the typed AWX Google Compute Engine
// credential (credential_gce) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialGceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Project        types.String `tfsdk:"project" json:"-"`
	SshKeyData     types.String `tfsdk:"ssh_key_data" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialGceTerraformModel) Clone() credentialGceTerraformModel {
	return *o
}

type credentialGceBodyRequestModel struct {
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
func (o *credentialGceTerraformModel) BodyRequest() *credentialGceBodyRequestModel {
	req := &credentialGceBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Project.IsNull() && !o.Project.IsUnknown() {
		inputs["project"] = o.Project.ValueString()
	}
	if !o.SshKeyData.IsNull() && !o.SshKeyData.IsUnknown() {
		inputs["ssh_key_data"] = o.SshKeyData.ValueString()
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
func (o *credentialGceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Project, inputs["project"], false))
		collect(helpers.AttrValueSetString(&o.SshKeyData, inputs["ssh_key_data"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// hookCredentialGce reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialGce(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialGceTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.SshKeyData.IsNull() || orig.SshKeyData.IsUnknown() {
			state.SshKeyData = types.StringNull()
		} else {
			state.SshKeyData = orig.SshKeyData
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.SshKeyData, state.SshKeyData); subbed {
			state.SshKeyData = v
		}
	}
	return nil
}

// credentialGceDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialGceDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Project        types.String `tfsdk:"project" json:"-"`
	Username       types.String `tfsdk:"username" json:"-"`
}

func (o *credentialGceDataSourceTerraformModel) Clone() credentialGceDataSourceTerraformModel {
	return *o
}

func (o *credentialGceDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Project, inputs["project"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
	}
	return diags, nil
}

// credentialGceTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialGceTypeLookup = framework.NewCredentialTypeLookup()

type credentialGceResource = framework.GenericResource[credentialGceTerraformModel, credentialGceBodyRequestModel, *credentialGceTerraformModel]

// NewCredentialGceResource constructs the typed Google Compute Engine credential resource.
// The credential_type ID is resolved by namespace (gce) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialGceResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["project"] = schema.StringAttribute{
		Description: "The Project ID is the GCE assigned identification. It is often constructed as three words or two words followed by a three-digit number. Examples: project-id-000 and another-project-id.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["ssh_key_data"] = schema.StringAttribute{
		Description: "Paste the contents of the PEM file associated with the service account email.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "The email address assigned to the Google Compute Engine service account.",
		Required:    true,
	}
	return &credentialGceResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_gce", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialGceTerraformModel, credentialGceBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `Google Compute Engine` (gce) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.gce.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialGceTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialGce,
			OnConfigure: credentialGceTypeLookup.OnConfigure("gce"),
			MutateBody: func(plan *credentialGceTerraformModel, body *credentialGceBodyRequestModel) {
				body.CredentialType = credentialGceTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialGceTerraformModel, body *credentialGceBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialGceTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialGceTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialGce",
		},
	}
}

type credentialGceDataSource = framework.GenericDataSource[credentialGceDataSourceTerraformModel, *credentialGceDataSourceTerraformModel]

// NewCredentialGceDataSource constructs the typed Google Compute Engine credential data source.
func NewCredentialGceDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["project"] = dschema.StringAttribute{
		Description: "The Project ID is the GCE assigned identification. It is often constructed as three words or two words followed by a three-digit number. Examples: project-id-000 and another-project-id.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "The email address assigned to the Google Compute Engine service account.",
		Computed:    true,
	}
	return &credentialGceDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_gce", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialGceDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Google Compute Engine` (gce) credential by ID or name.",
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
			OnConfigure:  credentialGceTypeLookup.OnConfigure("gce"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialGce",
		},
	}
}
