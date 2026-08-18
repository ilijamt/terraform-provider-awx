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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialGitlabTokenTerraformModel exposes the typed AWX GitLab Personal Access Token
// credential (credential_gitlab_token) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialGitlabTokenTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Token          types.String `tfsdk:"token" json:"-"`
}

func (o *credentialGitlabTokenTerraformModel) Clone() credentialGitlabTokenTerraformModel {
	return *o
}

type credentialGitlabTokenBodyRequestModel struct {
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
func (o *credentialGitlabTokenTerraformModel) BodyRequest() *credentialGitlabTokenBodyRequestModel {
	req := &credentialGitlabTokenBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Token.IsNull() && !o.Token.IsUnknown() {
		inputs["token"] = o.Token.ValueString()
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
func (o *credentialGitlabTokenTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Token, inputs["token"], false))
	}
	return diags, nil
}

// hookCredentialGitlabToken reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialGitlabToken(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialGitlabTokenTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Token.IsNull() || orig.Token.IsUnknown() {
			state.Token = types.StringNull()
		} else {
			state.Token = orig.Token
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Token, state.Token); subbed {
			state.Token = v
		}
	}
	return nil
}

// credentialGitlabTokenDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialGitlabTokenDataSourceTerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
}

func (o *credentialGitlabTokenDataSourceTerraformModel) Clone() credentialGitlabTokenDataSourceTerraformModel {
	return *o
}

func (o *credentialGitlabTokenDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
	return diags, nil
}

// credentialGitlabTokenTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialGitlabTokenTypeLookup = framework.NewCredentialTypeLookup()

type credentialGitlabTokenResource = framework.GenericResource[credentialGitlabTokenTerraformModel, credentialGitlabTokenBodyRequestModel, *credentialGitlabTokenTerraformModel]

// NewCredentialGitlabTokenResource constructs the typed GitLab Personal Access Token credential resource.
// The credential_type ID is resolved by namespace (gitlab_token) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialGitlabTokenResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["token"] = schema.StringAttribute{
		Description: "This token needs to come from your profile settings in GitLab.",
		Required:    true,
		Sensitive:   true,
	}
	return &credentialGitlabTokenResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_gitlab_token", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialGitlabTokenTerraformModel, credentialGitlabTokenBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `GitLab Personal Access Token` (gitlab_token) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.gitlab_token.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialGitlabTokenTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialGitlabToken,
			OnConfigure: credentialGitlabTokenTypeLookup.OnConfigure("gitlab_token"),
			MutateBody: func(plan *credentialGitlabTokenTerraformModel, body *credentialGitlabTokenBodyRequestModel) {
				body.CredentialType = credentialGitlabTokenTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialGitlabTokenTerraformModel, body *credentialGitlabTokenBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialGitlabTokenTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialGitlabTokenTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialGitlabToken",
		},
	}
}

type credentialGitlabTokenDataSource = framework.GenericDataSource[credentialGitlabTokenDataSourceTerraformModel, *credentialGitlabTokenDataSourceTerraformModel]

// NewCredentialGitlabTokenDataSource constructs the typed GitLab Personal Access Token credential data source.
func NewCredentialGitlabTokenDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	return &credentialGitlabTokenDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_gitlab_token", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialGitlabTokenDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `GitLab Personal Access Token` (gitlab_token) credential by ID or name.",
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
			OnConfigure:  credentialGitlabTokenTypeLookup.OnConfigure("gitlab_token"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialGitlabToken",
		},
	}
}
