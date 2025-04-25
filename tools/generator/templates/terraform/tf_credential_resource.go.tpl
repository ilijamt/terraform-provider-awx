package {{ .PackageName }}

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	r "github.com/ilijamt/terraform-provider-awx/internal/resource"
    "github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &{{ .Name | lowerCamelCase }}CredentialResource{}
	_ resource.ResourceWithConfigure   = &{{ .Name | lowerCamelCase }}CredentialResource{}
	_ resource.ResourceWithImportState = &{{ .Name | lowerCamelCase }}CredentialResource{}
)

// New{{ .Name }}CredentialResource is a helper function to simplify the provider implementation.
func New{{ .Name }}CredentialResource() resource.Resource {
	return &{{ .Name | lowerCamelCase }}CredentialResource{}
}

type {{ .Name | lowerCamelCase }}CredentialResource struct {
    client           c.Client
    rci              r.CallInfo
    endpoint         string
    name             string
    credentialTypeId int
}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
    if request.ProviderData == nil {
        return
    }
    o.rci = r.CallInfo{Name: "{{ .Name }}", Endpoint: "/api/v2/credentials/", TypeName: "{{ .TypeName }}" }
    o.name = "{{ .Name }}"
    o.client = request.ProviderData.(c.Client)
    o.endpoint = "/api/v2/credentials/"
    o.credentialTypeId = {{ .Id }}
}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential_{{ .TypeName }}"
}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Database ID of this credential",
				Computed:    true,
			},
{{- range $key, $value := .Fields }}
            "{{ $value.Id | lowerCase }}": schema.{{ $value.Generated.TerraformAttributeType }}Attribute{
                Description:         {{ escape_quotes (or $value.HelpText $value.Label) }},
				Required:            {{ $value.Generated.Required }},
				Optional:            {{ $value.Generated.Optional }},
				Computed:            {{ $value.Generated.Computed }},
				Sensitive:           {{ $value.Secret }},
            },
{{- end }}
        },
    }
}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the {{ .Name }}.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("{{ .IdKey }}"), id)...)
}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {}
func (o *{{ .Name | lowerCamelCase }}CredentialResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state {{ .Name | lowerCamelCase }}CredentialTerraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeRead)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
    var d, _ = r.Read(ctx, o.client, rci, state.ID, &state)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {}

func (o *{{ .Name | lowerCamelCase }}CredentialResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state {{ .Name | lowerCamelCase }}CredentialTerraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeDelete)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var d, _ = r.Delete(ctx, o.client, rci, state.ID)
	response.Diagnostics.Append(d...)
}
