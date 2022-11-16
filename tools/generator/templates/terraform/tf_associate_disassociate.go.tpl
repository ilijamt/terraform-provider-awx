{{ define "tf_associate_disassociate" }}

var (
	_ resource.Resource                  = &{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}{}
	_ resource.ResourceWithConfigure     = &{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}{}
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
{{- else }}
	_ resource.ResourceWithImportState   = &{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}{}
{{- end }}
)

type {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}TerraformModel struct {
	{{ .Name }}ID  types.Int64  `tfsdk:"{{ .Name | snakeCase }}_id"`
	{{ .Type }}ID  types.Int64  `tfsdk:"{{ .Type | snakeCase }}_id"`
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
	Option         types.String  `tfsdk:"option"`
{{- end }}
}

// New{{ .Name }}AssociateDisassociate{{ .Type }}Resource is a helper function to simplify the provider implementation.
func New{{ .Name }}AssociateDisassociate{{ .Type }}Resource() resource.Resource {
	return &{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}{}
}

type {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }} struct {
    client   c.Client
    endpoint string
}

func (o *{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
    if request.ProviderData == nil {
        return
    }

    o.client = request.ProviderData.(c.Client)
    o.endpoint = "{{ .Endpoint }}"
}

func (o {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
    response.TypeName = request.ProviderTypeName + "_{{ .Name | snakeCase }}_associate_{{ .Type | snakeCase }}"
}

func (o {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
    return processSchema(
        SourceResource,
        "{{ .Name }}/Associate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
			    "{{ .Name | snakeCase }}_id": {
					Description: "Database ID for this {{ .Name }}.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
							path.MatchRoot("option"),
							path.MatchRoot("{{ .Type | snakeCase }}_id"),
{{- else }}
							path.MatchRoot("{{ .Type | snakeCase }}_id"),
{{- end }}
						),
					},
			    },
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
				"option": {
					Description: "Notification Option",
					Required:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("{{ .Name | snakeCase }}_id"),
						),
{{- if eq .AssociateType "notification_job_template" }}
						stringvalidator.OneOf([]string{"started", "success", "error"}...),
{{- else if  (eq .AssociateType "notification_job_workflow_template") }}
						stringvalidator.OneOf([]string{"approval", "started", "success", "error"}...),
{{- else }}
{{- end }}
					},
				},
{{- end }}
				"{{ .Type | snakeCase }}_id": {
					Description: "Database ID of the {{ .Type | lowerCase }} to assign.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
							path.MatchRoot("option"),
{{- end }}
							path.MatchRoot("{{ .Name | snakeCase }}_id"),
						),
					},
				},
			},
		},
	), nil
}

{{ if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
{{- else }}
func (o *{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}TerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <{{ .Name | snakeCase }}_id>/<{{ .Type | snakeCase }}_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for {{ .Name }} association, invalid format.",
			err.Error(),
		)
		return
	}

	var {{ .Name | lowerCamelCase }}Id, {{ .Type | lowerCamelCase }}Id int64

	{{ .Name | lowerCamelCase }}Id, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the {{ .Name | snakeCase }}Id for the {{ .Name }} association.", request.ID),
			err.Error(),
		)
		return
	}
	state.{{ .Name }}ID = types.Int64Value({{ .Name | lowerCamelCase }}Id)

	{{ .Type | lowerCamelCase }}Id, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the {{ .Type | lowerCamelCase }}_id for the {{ .Name }} association.", request.ID),
			err.Error(),
		)
		return
	}
	state.{{ .Type }}ID = types.Int64Value({{ .Type | lowerCamelCase }}Id)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}
{{ end }}

func (o *{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of {{ .Name }}
	var r *http.Request
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.{{ .Name }}ID.ValueInt64(), plan.Option.ValueString())) + "/"
{{- else }}
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.{{ .Name }}ID.ValueInt64())) + "/"
{{- end }}
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.{{ .Type }}ID.ValueInt64(), Disassociate: false}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for create of type {{ default .AssociateType "default" }}", o.endpoint),
			err.Error(),
		)
		return
	}

    if _, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to associate for {{ .Name }} on %s", o.endpoint),
            err.Error(),
        )
        return
    }

	state.{{ .Name }}ID = plan.{{ .Name }}ID
	state.{{ .Type }}ID = plan.{{ .Type }}ID
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
	state.Option = plan.Option
{{ end }}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state {{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}TerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of {{ .Name }}
	var r *http.Request
{{- if or (eq .AssociateType "notification_job_template") (eq .AssociateType "notification_job_workflow_template") }}
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.{{ .Name }}ID.ValueInt64(), state.Option.ValueString())) + "/"
{{- else }}
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.{{ .Name }}ID.ValueInt64())) + "/"
{{- end }}
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.{{ .Type | camelCase }}ID.ValueInt64(), Disassociate: true}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
            fmt.Sprintf("Unable to create a new request for {{ .Name }} on %s for delete of type {{ default .AssociateType "default" }}" , o.endpoint),
			err.Error(),
		)
		return
	}

    if _, err = o.client.Do(ctx, r); err != nil {
        response.Diagnostics.AddError(
            fmt.Sprintf("Unable to disassociate for {{ .Name }} on %s", o.endpoint),
            err.Error(),
        )
        return
    }
}

func (o *{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *{{ .Name | lowerCamelCase }}AssociateDisassociate{{ .Type }}) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
{{ end }}