{{ define "tf_data_source" }}
var (
    _ datasource.DataSource                     = &{{ .Name | lowerCamelCase }}DataSource{}
    _ datasource.DataSourceWithConfigure        = &{{ .Name | lowerCamelCase }}DataSource{}
)

// New{{ .Name }}DataSource is a helper function to instantiate the {{ .Name }} data source.
func New{{ .Name }}DataSource() datasource.DataSource {
    return &{{ .Name | lowerCamelCase }}DataSource{}
}

// {{ .Name | lowerCamelCase }}DataSource is the data source implementation.
type {{ .Name | lowerCamelCase }}DataSource struct{
    client   c.Client
    endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *{{ .Name | lowerCamelCase }}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }

    o.client = req.ProviderData.(c.Client)
    o.endpoint = "{{ $.Endpoint }}"
}

// Metadata returns the data source type name.
func (o *{{ .Name | lowerCamelCase }}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_{{ $.Config.TypeName }}"
}

// GetSchema defines the schema for the data source.
func (o *{{ .Name | lowerCamelCase }}DataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
    return processSchema(
        SourceData,
        "{{ .Name }}",
        tfsdk.Schema{
		Version: helpers.SchemaVersion,
        Attributes: map[string]tfsdk.Attribute{
			// Data only elements
{{- range $key := .PropertyGetKeys }}
{{- with (index $.PropertyGetData $key) }}
            "{{ $key | lowerCase }}": {
                Description: {{ escape_quotes (default .help_text .label) }},
                Type:        {{ awx2tf_type . }},
{{- if (hasKey . "sensitive") }}
                Sensitive:   {{ .sensitive }},
{{- end }}
{{- if awx_is_property_searchable $.Config.SearchFields $key }}
                Optional:    true,
                Computed:    true,
{{- else }}
                Computed:    true,
{{- end }}
				Validators: []tfsdk.AttributeValidator{
{{- if eq .type "choice" }}
					stringvalidator.OneOf({{ awx_type_choice_data .choices }}...),
{{- end }}
{{- if awx_is_property_searchable $.Config.SearchFields $key }}
{{- range $key, $attrs := awx_generate_attribute_validator $.Config.SearchFields $key }}
					schemavalidator.{{ $key }}(
{{- range $attr := $attrs }}
						path.MatchRoot("{{ $attr }}"),
{{- end }}
					),
{{- end }}
{{- end }}
				},
            },
{{- end }}
{{- end }}
            // Write only elements
{{- range $key := .PropertyWriteOnlyKeys }}
{{- with (index $.PropertyWriteOnlyData $key) }}
            "{{ $key | lowerCase }}": {
                Description: {{ escape_quotes (default .help_text .label) }},
                Type:        {{ awx2tf_type . }},
                Computed:    true,
{{- if (hasKey . "sensitive") }}
                Sensitive:   {{ .sensitive }},
{{- end }}
{{- end }}
            },
{{- end }}
		},
	}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *{{ .Name | lowerCamelCase }}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state {{ .Name | lowerCamelCase }}TerraformModel
	var err error
	var endpoint string
{{- if gt (len $.Config.SearchFields) 0 }}
    var searchDefined bool

    // Only one group should evaluate to True, terraform should prevent from being able to set
    // the conflicting groups
{{- else }}
    endpoint = o.endpoint
{{ end }}

{{ range $field := $.Config.SearchFields }}
    // Evaluate group '{{ $field.Name }}' based on the schema definition
    var group{{ $field.Name | camelCase }}Exists = func() bool {
         var group{{ $field.Name | camelCase }}Exists = true
         var params{{ $field.Name | camelCase }} = []any{o.endpoint}
{{- range $attr := $field.Fields }}
         var attr{{ $attr.Name | camelCase }} {{ awx2go_type (index $.PropertyGetData $attr.Name) }}
         req.Config.GetAttribute(ctx, path.Root("{{ $attr.Name }}"), &attr{{ $attr.Name | camelCase }})
         group{{ $field.Name | camelCase }}Exists = group{{ $field.Name | camelCase }}Exists && (!attr{{ $attr.Name | camelCase }}.IsNull() && !attr{{ $attr.Name | camelCase }}.IsUnknown())
{{- if $attr.UrlEscapeValue }}
         params{{ $field.Name | camelCase }} = append(params{{ $field.Name | camelCase }}, url.PathEscape(attr{{ $attr.Name | camelCase }}.{{ tf2go_primitive_value (index $.PropertyGetData $attr.Name) }}()))
{{- else }}
         params{{ $field.Name | camelCase }} = append(params{{ $field.Name | camelCase }}, attr{{ $attr.Name | camelCase }}.{{ tf2go_primitive_value (index $.PropertyGetData $attr.Name) }}())
{{- end }}
{{- end }}
        if group{{ $field.Name | camelCase }}Exists {
            endpoint = p.Clean(fmt.Sprintf("%s/{{ $field.UrlSuffix }}", params{{ $field.Name | camelCase }}...))
        }
         return group{{ $field.Name | camelCase }}Exists
    }()
    searchDefined = searchDefined || group{{ $field.Name | camelCase }}Exists
{{ end }}

{{ if gt (len $.Config.SearchFields) 0 }}
    if !searchDefined {
        var detailMessage string
        resp.Diagnostics.AddError(
            fmt.Sprintf("missing configuration for one of the predefined search groups"),
            detailMessage,
        )
        return
    }
{{ end }}

	// Creates a new request for {{ .Name }}
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for {{ .Name }}
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
            fmt.Sprintf("Unable to read resource for {{ .Name }} on %s", o.endpoint),
			err.Error(),
		)
		return
	}

    var d diag.Diagnostics

{{ if gt (len $.Config.SearchFields) 0 }}
	if data, d, err = extractDataIfSearchResult(data); err != nil {
        resp.Diagnostics.Append(d...)
        return
	}
{{ end }}

    if d, err = state.updateFromApiData(data); err != nil {
        resp.Diagnostics.Append(d...)
        return
    }

    // Set state
{{- if $.Config.PreStateSetHookFunction }}
    if err = {{ $.Config.PreStateSetHookFunction }}(SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on {{ .Name }}"),
			err.Error(),
		)
		return
    }
{{ end }}
    resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
    if resp.Diagnostics.HasError() {
        return
    }
}
{{ end }}
{{ block "tf_data_source" . }}{{ end }}