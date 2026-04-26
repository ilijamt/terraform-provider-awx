package {{ .PackageName }}

import (
	"context"
	"fmt"
	"net/http"
	p "path"

    "github.com/ilijamt/terraform-provider-awx/internal/hooks"
	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type {{ .Name | lowerCamelCase }}DataSource = framework.GenericDataSource[{{ .Name | lowerCamelCase }}TerraformModel, *{{ .Name | lowerCamelCase }}TerraformModel]

// New{{ .Name }}DataSource is a helper function to instantiate the {{ .Name }} data source.
func New{{ .Name }}DataSource() datasource.DataSource {
    return &{{ .Name | lowerCamelCase }}DataSource{
        DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ .TypeName }}", Endpoint: "{{ .Endpoint }}"}},
        Cfg: framework.DataSourceCfg[{{ .Name | lowerCamelCase }}TerraformModel]{
            Schema: schema.Schema{
{{- if .Deprecated }}
                DeprecationMessage: "This data source has been deprecated and will be removed in a future release.",
{{- end }}
                Attributes: map[string]schema.Attribute{
                    // Data only elements
{{- range $key, $value := $.ReadProperties }}
                    "{{ $key | lowerCase }}": schema.{{ $value.Generated.AttributeType }}Attribute{
{{- if $value.Deprecated }}
                        DeprecationMessage: "This field is deprecated and will be removed in a future release.",
{{- end }}
{{- if and (eq $value.Generated.AttributeType "List") (eq $value.ElementType "choice") }}
                        ElementType: types.ListType{ElemType: types.StringType},
{{- else if eq $value.Generated.AttributeType "List" }}
                        ElementType: types.StringType,
{{- end }}
                        Description: {{ escape_quotes (or .Description .Label) }},
                        Sensitive: {{ .IsSensitive }},
{{- if $value.IsSearchable }}
                        Optional:    true,
                        Computed:    true,
{{- else }}
                        Computed:    true,
{{- end }}
{{- if $value.IsSearchable }}
                        Validators: []validator.{{ $value.Generated.AttributeType }}{
{{- range $key, $attrs := $value.Generated.AttributeValidationData }}
                            {{ $value.Generated.AttributeType | lowerCase }}validator.{{ $key }}(
{{- range $attr := $attrs }}
                                path.MatchRoot("{{ $attr }}"),
{{- end }}
                            ),
{{- end }}
                        },
{{- end }}
                    },
{{- end }}
{{- range $key, $value := $.WriteProperties }}
{{- if $value.IsWriteOnly }}
                    "{{ $key | lowerCase }}": schema.{{ $value.Generated.AttributeType }}Attribute{
{{- if eq $value.Generated.AttributeType "List" }}
                        ElementType: types.{{ camelCase $value.ElementType }}Type,
{{- end }}
{{- if $value.Deprecated }}
                        DeprecationMessage: "This field is deprecated and will be removed in a future release.",
{{- end }}
                        Description: {{ escape_quotes (or .Description .Label) }},
                        Sensitive:   {{ $value.IsSensitive }},
                        Optional:    true,
                        Computed:    true,
                    },
{{- end }}
{{- end }}
                },
            },
{{- if .HasSearchFields }}
            SearchGroups: []framework.SearchGroup{
{{- range $field := .SearchFields }}
                {Name: "{{ $field.Name }}", URLSuffix: "{{ $field.UrlSuffix }}", Fields: []framework.SearchField{
{{- range $attr := $field.Fields }}
                    {Name: "{{ $attr.Name }}", Type: "{{ if eq (index $.ReadProperties $attr.Name).Generated.AwxGoType "types.Int64" }}int64{{ else }}string{{ end }}", URLEscape: {{ $attr.UrlEscapeValue }}},
{{- end }}
                }},
{{- end }}
            },
{{- end }}
{{- if .PreStateSetHookFunction }}
{{- if eq .PreStateSetHookFunction "hooks.RequireResourceStateOrOrig" }}
            Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *{{ .Name | lowerCamelCase }}TerraformModel) error {
                return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
            },
{{- else }}
            Hook: {{ .PreStateSetHookFunction }},
{{- end }}
{{- end }}
            ApiVersion: ApiVersion,
            ResourceName: "{{ .Name }}",
        },
    }
}

