{{ define "tf_credential" }}
{{- range $ct := $.Config.CredentialTypes }}
var (
    _ datasource.DataSource                     = &{{ $.Config.TypeName | lowerCamelCase }}{{ $ct.Name | camelCase }}DataSource{}
    _ datasource.DataSourceWithConfigure        = &{{ $.Config.TypeName | lowerCamelCase }}{{ $ct.Name | camelCase }}DataSource{}
)

// {{ $.Config.TypeName | lowerCamelCase }}{{ $ct.Name | camelCase }}DataSource is the data source implementation.
type {{ credentialNameStruct $.Config.TypeName $ct.Name }}DataSource struct{
    client   c.Client
    endpoint string
}

// New{{ $.Config.TypeName | camelCase }}{{ $ct.Name | camelCase }}DataSource is a helper function to instantiate the {{ .Name }} data source.
func New{{ $.Config.TypeName | camelCase }}{{ $ct.Name | camelCase }}DataSource() datasource.DataSource {
    return &{{  credentialNameStruct $.Config.TypeName $ct.Name  }}DataSource{}
}

// Configure adds the provider configured client to the data source.
func (o *{{  credentialNameStruct $.Config.TypeName $ct.Name  }}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }

    o.client = req.ProviderData.(c.Client)
    o.endpoint = "{{ $.Endpoint }}"
}

// Metadata returns the data source type name.
func (o *{{  credentialNameStruct $.Config.TypeName $ct.Name  }}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_{{ $.Config.TypeName }}_{{ replace ($ct.Name | snakeCase) "/" "_" }}"
}

// GetSchema defines the schema for the data source.
func (o *{{  credentialNameStruct $.Config.TypeName $ct.Name  }}DataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
    return processSchema(
        SourceData,
        "{{ $.Config.TypeName | camelCase }}{{ .Name | camelCase }}",
        tfsdk.Schema{
		Version: helpers.SchemaVersion,
        Attributes: map[string]tfsdk.Attribute{
{{- range $attr := credAttrs $.Config.TypeName $ct.Inputs }}
        "{{ $attr.id | lowerCase }}": {
            Sensitive: {{ default .secret false }},
            Description: {{ escape_quotes (default .help_text .label) }},
            Validators: []tfsdk.AttributeValidator{
{{- if $attr.choices -}}
                stringvalidator.OneOf([]string{
{{- range $val := $attr.choices -}}
                    "{{ $val }}",
{{- end -}}
                }...),
{{- end }}
            },
{{- if $attr.choices }}
            Type:        types.StringType,
{{- else }}
            Type:        {{ awx2tf_type . }},
{{- end }}
         },
{{- end }}
		},
	}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *{{ $.Config.TypeName | lowerCamelCase }}{{ $ct.Name | camelCase }}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    panic("not implemented")
}
{{ end }}
{{ end }}