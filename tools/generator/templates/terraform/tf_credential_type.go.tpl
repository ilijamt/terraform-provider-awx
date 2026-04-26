package {{ .PackageName }}

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

// {{ .Name | lowerCamelCase }}TerraformModel exposes the typed AWX {{ .DisplayName }}
// credential ({{ .TypeName }}) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type {{ .Name | lowerCamelCase }}TerraformModel struct {
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Name           types.String `tfsdk:"name" json:"name"`
	Description    types.String `tfsdk:"description" json:"description"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
{{- range .Fields }}
	{{ .PropertyName }} types.String `tfsdk:"{{ .ID }}" json:"-"`
{{- end }}
}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
	return *o
}

type {{ .Name | lowerCamelCase }}BodyRequestModel struct {
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
func (o *{{ .Name | lowerCamelCase }}TerraformModel) BodyRequest() *{{ .Name | lowerCamelCase }}BodyRequestModel {
	req := &{{ .Name | lowerCamelCase }}BodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
{{- range .Fields }}
	if !o.{{ .PropertyName }}.IsNull() && !o.{{ .PropertyName }}.IsUnknown() {
		inputs["{{ .ID }}"] = o.{{ .PropertyName }}.ValueString()
	}
{{- end }}
	if len(inputs) > 0 {
		payload, _ := json.Marshal(inputs)
		req.Inputs = payload
	}
	return req
}

// UpdateFromApiData unfolds the AWX response back into the typed model. Secret
// fields come back as `$encrypted$` placeholders; the per-credential-type
// pre-state-set hook reconciles them against prior plan state.
func (o *{{ .Name | lowerCamelCase }}TerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
{{- range .Fields }}
		collect(helpers.AttrValueSetString(&o.{{ .PropertyName }}, inputs["{{ .ID }}"], false))
{{- end }}
	}
	return diags, nil
}

// hook{{ .Name }} reconciles `$encrypted$` placeholders that AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan. Data-source reads have orig==nil and skip reconciliation.
func hook{{ .Name }}(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *{{ .Name | lowerCamelCase }}TerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
{{- range .Fields }}{{ if .Secret }}
		if orig.{{ .PropertyName }}.IsNull() || orig.{{ .PropertyName }}.IsUnknown() {
			state.{{ .PropertyName }} = types.StringNull()
		} else {
			state.{{ .PropertyName }} = orig.{{ .PropertyName }}
		}
{{- end }}{{ end }}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
{{- range .Fields }}{{ if .Secret }}
		if v, subbed := helpers.MergeEncryptedField(orig.{{ .PropertyName }}, state.{{ .PropertyName }}); subbed {
			state.{{ .PropertyName }} = v
		}
{{- end }}{{ end }}
	}
	return nil
}

// {{ .Name | lowerCamelCase }}TypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var {{ .Name | lowerCamelCase }}TypeLookup = framework.NewCredentialTypeLookup()

type {{ .Name | lowerCamelCase }}Resource = framework.GenericResource[{{ .Name | lowerCamelCase }}TerraformModel, {{ .Name | lowerCamelCase }}BodyRequestModel, *{{ .Name | lowerCamelCase }}TerraformModel]

// New{{ .Name }}Resource constructs the typed {{ .DisplayName }} credential resource.
// The credential_type ID is resolved by namespace ({{ .Namespace }}) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func New{{ .Name }}Resource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
{{- range .Fields }}
	attrs["{{ .ID }}"] = schema.StringAttribute{
		Description: {{ if .HelpText }}{{ .HelpText | escape_quotes }}{{ else }}{{ .Label | escape_quotes }}{{ end }},
{{- if .Required }}
		Required:    true,
{{- else }}
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
{{- end }}
{{- if .Secret }}
		Sensitive:   true,
{{- end }}
	}
{{- end }}
	return &{{ .Name | lowerCamelCase }}Resource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .TypeName }}", Endpoint: "{{ .Endpoint }}"}},
		Cfg: framework.ResourceCfg[{{ .Name | lowerCamelCase }}TerraformModel, {{ .Name | lowerCamelCase }}BodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `{{ .DisplayName }}` ({{ .Namespace }}) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.{{ .Namespace }}.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *{{ .Name | lowerCamelCase }}TerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hook{{ .Name }},
			OnConfigure: {{ .Name | lowerCamelCase }}TypeLookup.OnConfigure("{{ .Namespace }}"),
			MutateBody: func(plan *{{ .Name | lowerCamelCase }}TerraformModel, body *{{ .Name | lowerCamelCase }}BodyRequestModel) {
				body.CredentialType = {{ .Name | lowerCamelCase }}TypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *{{ .Name | lowerCamelCase }}TerraformModel, body *{{ .Name | lowerCamelCase }}BodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *{{ .Name | lowerCamelCase }}TerraformModel) {
				state.Team = types.Int64Value(plan.Team.ValueInt64())
				state.User = types.Int64Value(plan.User.ValueInt64())
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value({{ .Name | lowerCamelCase }}TypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "{{ .Name }}",
		},
	}
}

type {{ .Name | lowerCamelCase }}DataSource = framework.GenericDataSource[{{ .Name | lowerCamelCase }}TerraformModel, *{{ .Name | lowerCamelCase }}TerraformModel]

// New{{ .Name }}DataSource constructs the typed {{ .DisplayName }} credential data source.
func New{{ .Name }}DataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
{{- range .Fields }}
	attrs["{{ .ID }}"] = dschema.StringAttribute{
		Description: {{ if .HelpText }}{{ .HelpText | escape_quotes }}{{ else }}{{ .Label | escape_quotes }}{{ end }},
		Computed:    true,
{{- if .Secret }}
		Sensitive:   true,
{{- end }}
	}
{{- end }}
	return &{{ .Name | lowerCamelCase }}DataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .TypeName }}", Endpoint: "{{ .Endpoint }}"}},
		Cfg: framework.DataSourceCfg[{{ .Name | lowerCamelCase }}TerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `{{ .DisplayName }}` ({{ .Namespace }}) credential by ID or name.",
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
			OnConfigure:  {{ .Name | lowerCamelCase }}TypeLookup.OnConfigure("{{ .Namespace }}"),
			Hook:         hook{{ .Name }},
			ApiVersion:   ApiVersion,
			ResourceName: "{{ .Name }}",
		},
	}
}
