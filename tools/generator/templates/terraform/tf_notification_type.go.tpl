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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

type {{ .Name | lowerCamelCase }}TerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
{{- range .Fields }}
	{{ .PropertyName }} types.{{ template "nt_tf_type" . }} `tfsdk:"{{ .ID }}" json:"-"`
{{- end }}
}

func (o *{{ .Name | lowerCamelCase }}TerraformModel) Clone() {{ .Name | lowerCamelCase }}TerraformModel {
	return *o
}

type {{ .Name | lowerCamelCase }}BodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *{{ .Name | lowerCamelCase }}TerraformModel) BodyRequest() *{{ .Name | lowerCamelCase }}BodyRequestModel {
	req := &{{ .Name | lowerCamelCase }}BodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "{{ .Namespace }}",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
{{- range .Fields }}
{{- if .IsObject }}
	// AWX requires the key even when the map is empty.
	config["{{ .ID }}"] = helpers.MapAsStringMap(o.{{ .PropertyName }}, false)
{{- else }}
	if !o.{{ .PropertyName }}.IsNull() && !o.{{ .PropertyName }}.IsUnknown() {
{{- if .IsList }}
		config["{{ .ID }}"] = helpers.ListAsStringSlice(o.{{ .PropertyName }}, false)
{{- else }}
		config["{{ .ID }}"] = o.{{ .PropertyName }}.Value{{ template "nt_tf_type" . }}()
{{- end }}
	}
{{- end }}
{{- end }}
	req.NotificationConfiguration = config
	return req
}

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
	collect(helpers.AttrValueSetString(&o.NotificationType, data["notification_type"], false))
	collect(helpers.AttrValueSetJsonString(&o.Messages, data["messages"], false))

	if config, ok := data["notification_configuration"].(map[string]any); ok {
{{- range .Fields }}
{{- if .IsBool }}
		collect(helpers.AttrValueSetBool(&o.{{ .PropertyName }}, config["{{ .ID }}"]))
{{- else if .IsInt }}
		collect(helpers.AttrValueSetInt64(&o.{{ .PropertyName }}, config["{{ .ID }}"]))
{{- else if .IsList }}
		collect(helpers.AttrValueSetListString(&o.{{ .PropertyName }}, config["{{ .ID }}"], false))
{{- else if .IsObject }}
		collect(helpers.AttrValueSetMapString(&o.{{ .PropertyName }}, config["{{ .ID }}"], false))
{{- else }}
		collect(helpers.AttrValueSetString(&o.{{ .PropertyName }}, config["{{ .ID }}"], false))
{{- end }}
{{- end }}
	}
	return diags, nil
}

{{- if .HasSecrets }}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hook{{ .Name }}(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *{{ .Name | lowerCamelCase }}TerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// AWX never returns this value, so state has to carry a known null.
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
{{- end }}

type {{ .Name | lowerCamelCase }}DataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
{{- range .DataSourceFields }}
	{{ .PropertyName }} types.{{ template "nt_tf_type" . }} `tfsdk:"{{ .ID }}" json:"-"`
{{- end }}
}

func (o *{{ .Name | lowerCamelCase }}DataSourceTerraformModel) Clone() {{ .Name | lowerCamelCase }}DataSourceTerraformModel {
	return *o
}

func (o *{{ .Name | lowerCamelCase }}DataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.NotificationType, data["notification_type"], false))
	collect(helpers.AttrValueSetJsonString(&o.Messages, data["messages"], false))
{{- if .DataSourceFields }}

	if config, ok := data["notification_configuration"].(map[string]any); ok {
{{- range .DataSourceFields }}
{{- if .IsBool }}
		collect(helpers.AttrValueSetBool(&o.{{ .PropertyName }}, config["{{ .ID }}"]))
{{- else if .IsInt }}
		collect(helpers.AttrValueSetInt64(&o.{{ .PropertyName }}, config["{{ .ID }}"]))
{{- else if .IsList }}
		collect(helpers.AttrValueSetListString(&o.{{ .PropertyName }}, config["{{ .ID }}"], false))
{{- else if .IsObject }}
		collect(helpers.AttrValueSetMapString(&o.{{ .PropertyName }}, config["{{ .ID }}"], false))
{{- else }}
		collect(helpers.AttrValueSetString(&o.{{ .PropertyName }}, config["{{ .ID }}"], false))
{{- end }}
{{- end }}
	}
{{- end }}
	return diags, nil
}

type {{ .Name | lowerCamelCase }}Resource = framework.GenericResource[{{ .Name | lowerCamelCase }}TerraformModel, {{ .Name | lowerCamelCase }}BodyRequestModel, *{{ .Name | lowerCamelCase }}TerraformModel]

func New{{ .Name }}Resource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
{{- range .Fields }}
{{- if .IsBool }}
	attrs["{{ .ID }}"] = schema.BoolAttribute{
		Description: {{ .Description | escape_quotes }},
{{- if .Required }}
		Required:    true,
{{- else }}
		Optional:    true,
		Computed:    true,
{{- if .HasDefault }}
		Default:     booldefault.StaticBool({{ .Default }}),
{{- end }}
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
{{- end }}
	}
{{- else if .IsInt }}
	attrs["{{ .ID }}"] = schema.Int64Attribute{
		Description: {{ .Description | escape_quotes }},
{{- if .Required }}
		Required:    true,
{{- else }}
		Optional:    true,
		Computed:    true,
{{- if .HasDefault }}
		Default:     int64default.StaticInt64({{ .Default | format_number }}),
{{- end }}
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
{{- end }}
	}
{{- else if .IsList }}
	attrs["{{ .ID }}"] = schema.ListAttribute{
		Description: {{ .Description | escape_quotes }},
		ElementType: types.StringType,
{{- if .Required }}
		Required:    true,
{{- else }}
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.UseStateForUnknown(),
		},
{{- end }}
	}
{{- else if .IsObject }}
	attrs["{{ .ID }}"] = schema.MapAttribute{
		Description: {{ .Description | escape_quotes }},
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.UseStateForUnknown(),
		},
	}
{{- else }}
	attrs["{{ .ID }}"] = schema.StringAttribute{
		Description: {{ .Description | escape_quotes }},
{{- if .Required }}
		Required:    true,
{{- else }}
		Optional:    true,
		Computed:    true,
{{- if .HasDefault }}
		Default:     stringdefault.StaticString({{ .Default | escape_quotes }}),
{{- end }}
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
{{- end }}
{{- if .Secret }}
		Sensitive:   true,
{{- end }}
	}
{{- end }}
{{- end }}
	return &{{ .Name | lowerCamelCase }}Resource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .TypeName }}", Endpoint: "{{ .Endpoint }}"}},
		Cfg: framework.ResourceCfg[{{ .Name | lowerCamelCase }}TerraformModel, {{ .Name | lowerCamelCase }}BodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `{{ .DisplayName }}` (`{{ .Namespace }}`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"{{ .Namespace }}\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *{{ .Name | lowerCamelCase }}TerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
{{- if .HasSecrets }}
			Hook:       hook{{ .Name }},
{{- end }}
			WriteOnlyPlanToState: func(plan, state *{{ .Name | lowerCamelCase }}TerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("{{ .Namespace }}")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "{{ .Name }}",
		},
	}
}

type {{ .Name | lowerCamelCase }}DataSource = framework.GenericDataSource[{{ .Name | lowerCamelCase }}DataSourceTerraformModel, *{{ .Name | lowerCamelCase }}DataSourceTerraformModel]

func New{{ .Name }}DataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
{{- range .DataSourceFields }}
{{- if .IsList }}
	attrs["{{ .ID }}"] = dschema.ListAttribute{
		Description: {{ .Description | escape_quotes }},
		ElementType: types.StringType,
		Computed:    true,
	}
{{- else if .IsObject }}
	attrs["{{ .ID }}"] = dschema.MapAttribute{
		Description: {{ .Description | escape_quotes }},
		ElementType: types.StringType,
		Computed:    true,
	}
{{- else }}
	attrs["{{ .ID }}"] = dschema.{{ template "nt_tf_type" . }}Attribute{
		Description: {{ .Description | escape_quotes }},
		Computed:    true,
	}
{{- end }}
{{- end }}
	return &{{ .Name | lowerCamelCase }}DataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .TypeName }}", Endpoint: "{{ .Endpoint }}"}},
		Cfg: framework.DataSourceCfg[{{ .Name | lowerCamelCase }}DataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `{{ .DisplayName }}` (`{{ .Namespace }}`) notification template by ID or name.",
				Attributes:          attrs,
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "{{ .Name }}",
		},
	}
}

{{- define "nt_tf_type" -}}
{{ if .IsBool }}Bool{{ else if .IsInt }}Int64{{ else if .IsList }}List{{ else if .IsObject }}Map{{ else }}String{{ end }}
{{- end }}
