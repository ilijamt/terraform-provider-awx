{{ define "tf_resource_object_role" }}
var (
    _ datasource.DataSource                     = &{{ .Name | lowerCamelCase }}ObjectRolesDataSource{}
    _ datasource.DataSourceWithConfigure        = &{{ .Name | lowerCamelCase }}ObjectRolesDataSource{}
)

// New{{ .Name }}ObjectRolesDataSource is a helper function to instantiate the {{ .Name }} data source.
func New{{ .Name }}ObjectRolesDataSource() datasource.DataSource {
    return &{{ .Name | lowerCamelCase }}ObjectRolesDataSource{}
}

// {{ .Name | lowerCamelCase }}ObjectRolesDataSource is the data source implementation.
type {{ .Name | lowerCamelCase }}ObjectRolesDataSource struct{
    client   c.Client
    endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *{{ .Name | lowerCamelCase }}ObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }

    o.client = req.ProviderData.(c.Client)
    o.endpoint = "{{ $.Endpoint }}%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *{{ .Name | lowerCamelCase }}ObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_{{ $.Config.TypeName }}_object_roles"
}

// GetSchema defines the schema for the data source.
func (o *{{ .Name | lowerCamelCase }}ObjectRolesDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
    return tfsdk.Schema{
		Version: helpers.SchemaVersion,
        Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "{{ .Name }} ID",
				Type:        types.Int64Type,
				Required:    true,
			},
			"roles": {
				Description: "Roles for {{ .Name | lowerCase }}",
				Type:        types.MapType{ElemType: types.Int64Type},
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (o *{{ .Name | lowerCamelCase }}ObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    var state {{ .Name | lowerCamelCase }}ObjectRolesModel
	var err error
	var id types.Int64

	if d := req.Config.GetAttribute(ctx, path.Root("id"), &id); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	// Creates a new request for Credential
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(o.endpoint, id.ValueInt64()), nil); err != nil {
		resp.Diagnostics.AddError(
			"Unable to create a new request for {{ .Name | lowerCamelCase }}",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch the request for {{ .Name | lowerCase }} object roles",
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
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

	var d diag.Diagnostics
	if state.Roles, d = types.MapValue(types.Int64Type, in); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}

    // Set state
    resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
    if resp.Diagnostics.HasError() {
        return
    }
}
{{ end }}
{{ block "tf_resource_object_role" . }}{{ end }}