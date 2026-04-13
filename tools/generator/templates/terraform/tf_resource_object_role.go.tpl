package {{ .PackageName }}

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mitchellh/mapstructure"
)

var (
    _ datasource.DataSource                     = &{{ .Name | lowerCamelCase }}ObjectRolesDataSource{}
    _ datasource.DataSourceWithConfigure        = &{{ .Name | lowerCamelCase }}ObjectRolesDataSource{}
)

// New{{ .Name }}ObjectRolesDataSource is a helper function to instantiate the {{ .Name }} data source.
func New{{ .Name }}ObjectRolesDataSource() datasource.DataSource {
    return &{{ .Name | lowerCamelCase }}ObjectRolesDataSource{DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "{{ $.TypeName }}_object_roles", Endpoint: "{{ .Endpoint }}%d/object_roles/"}}}
}

// {{ .Name | lowerCamelCase }}ObjectRolesDataSource is the data source implementation.
type {{ .Name | lowerCamelCase }}ObjectRolesDataSource struct{
    framework.DataSourceBase
}

// Schema defines the schema for the data source.
func (o *{{ .Name | lowerCamelCase }}ObjectRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
{{- if (index .DeprecatedParts "ObjectRoles") }}
        DeprecationMessage: "This data source has been deprecated and will be removed in a future release.",
{{- end }}
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
				Description: "{{ .Name }} ID",
				Required:    true,
			},
			"roles": schema.MapAttribute{
				Description: "Roles for {{ $.TypeName }}",
				ElementType: types.Int64Type,
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *{{ .Name | lowerCamelCase }}ObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    var state {{ .Name | lowerCamelCase }}ObjectRolesModel
	var err error
	var id types.Int64

	if framework.DiagnosticsHasError(&resp.Diagnostics, req.Config.GetAttribute(ctx, path.Root("id"), &id)...) { return }
	state.ID = types.Int64Value(id.ValueInt64())

	var endpoint = fmt.Sprintf(o.Endpoint, id.ValueInt64())
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "{{ .Name }}/ObjectRoles")
	if framework.DiagnosticsHasError(&resp.Diagnostics, d...) { return }

	var sr models.SearchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for {{ .Name | lowerCase }}",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var dg diag.Diagnostics
	if state.Roles, dg = types.MapValue(types.Int64Type, in); framework.DiagnosticsHasError(&resp.Diagnostics, dg...) { return }

	if framework.DiagnosticsHasError(&resp.Diagnostics, resp.State.Set(ctx, &state)...) { return }
}
