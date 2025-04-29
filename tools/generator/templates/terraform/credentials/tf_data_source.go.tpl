package {{ .PackageName }}

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
    "github.com/ilijamt/terraform-provider-awx/internal/hooks"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
    _ datasource.DataSource                     = &{{ .Name | lowerCamelCase }}CredentialDataSource{}
    _ datasource.DataSourceWithConfigure        = &{{ .Name | lowerCamelCase }}CredentialDataSource{}
)

// New{{ .Name }}CredentialDataSource is a helper function to instantiate the {{ .Name }} data source.
func New{{ .Name }}CredentialDataSource() datasource.DataSource {
    return &{{ .Name | lowerCamelCase }}CredentialDataSource{}
}

// {{ .Name | lowerCamelCase }}CredentialDataSource is the data source implementation.
type {{ .Name | lowerCamelCase }}CredentialDataSource struct{
    client   c.Client
    endpoint string
    name     string
}

func (o *{{ .Name | lowerCamelCase }}CredentialDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Database ID of this credential",
				Optional:    true,
			},
{{- range $key, $value := .Fields }}
{{- if not $value.Secret }}
            "{{ $value.Id | lowerCase }}": schema.{{ $value.Generated.TerraformAttributeType }}Attribute{
                Description:         {{ escape_quotes (or $value.HelpText $value.Label) }},
{{- if eq $value.Id "name" }}
				Optional:           true,
{{ else }}
				Computed:           true,
{{- end }}
            },
{{- end }}
{{- end }}
        },
    }
}


// Configure adds the provider configured client to the data source.
func (o *{{ .Name | lowerCamelCase }}CredentialDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }

    o.name = "{{ .Name }}"
    o.client = req.ProviderData.(c.Client)
    o.endpoint = "/api/v2/credentials/"
}

// Metadata returns the data source type name.
func (o *{{ .Name | lowerCamelCase }}CredentialDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_credential_{{ .TypeName }}"
}

func (o *{{ .Name | lowerCamelCase }}CredentialDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
    return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *{{ .Name | lowerCamelCase }}CredentialDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    // @TODO
}