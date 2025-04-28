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
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
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
	_ resource.Resource                = &{{ $.TypeName | lowerCamelCase }}CredentialResource{}
	_ resource.ResourceWithConfigValidators = &{{ $.TypeName | lowerCamelCase }}CredentialResource{}
	_ resource.ResourceWithConfigure   = &{{ $.TypeName | lowerCamelCase }}CredentialResource{}
	_ resource.ResourceWithImportState = &{{ $.TypeName | lowerCamelCase }}CredentialResource{}
)

// New{{ $.TypeName | pascalCase }}CredentialResource is a helper function to simplify the provider implementation.
func New{{ $.TypeName | pascalCase }}CredentialResource() resource.Resource {
	return &{{ $.TypeName | lowerCamelCase }}CredentialResource{}
}

type {{ $.TypeName | lowerCamelCase }}CredentialResource struct {
    client           c.Client
    rci              r.CallInfo
    endpoint         string
    name             string
    typeName         string
    credentialTypeId int
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
    if request.ProviderData == nil {
        return
    }
    o.name = "{{ $.Name }}"
    o.typeName = "{{ $.TypeName }}"
    o.rci = r.CallInfo{Name: o.name, TypeName: o.typeName, Endpoint: "/api/v2/credentials/" }
    o.client = request.ProviderData.(c.Client)
    o.endpoint = "/api/v2/credentials/"
    o.credentialTypeId = {{ .Id }}
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential_{{ .TypeName }}"
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
{{- range $key, $value := .Fields }}
{{- if $value.IsUTO }}
			path.MatchRoot("{{ $value.Id }}"),
{{- end }}
{{- end }}
		),
	}
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
		Description: "{{ .Description }}",
        Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Database ID of this credential",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
{{- range $key, $value := .Fields }}
            "{{ $value.Id | lowerCase }}": schema.{{ $value.Generated.TerraformAttributeType }}Attribute{
                Description:         {{ escape_quotes (or $value.HelpText $value.Label) }},
				Required:            {{ $value.Generated.Required }},
				Optional:            {{ $value.Generated.Optional }},
				Computed:            {{ $value.Generated.Computed }},
				Sensitive:           {{ $value.Secret }},
				WriteOnly:           {{ $value.Generated.WriteOnly }},
            },
{{- end }}
        },
    }
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the {{ $.TypeName }}.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("{{ .IdKey }}"), id)...)
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var rci = o.rci.With(r.SourceResource, r.CalleeCreate)
	var plan {{ $.TypeName | lowerCamelCase }}CredentialTerraformModel
    response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
    if response.Diagnostics.HasError() {
        return
    }
    var d, _ = r.Create(ctx, o.client, rci, &plan)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state {{ $.TypeName | lowerCamelCase }}CredentialTerraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeRead)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
    var d, _ = r.Read(ctx, o.client, rci, &state)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var rci = o.rci.With(r.SourceResource, r.CalleeUpdate)
	var plan {{ $.TypeName | lowerCamelCase }}CredentialTerraformModel
    response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
    if response.Diagnostics.HasError() {
        return
    }
    var d, _ = r.Update(ctx, o.client, rci, &plan)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (o *{{ $.TypeName | lowerCamelCase }}CredentialResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state {{ $.TypeName | lowerCamelCase }}CredentialTerraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeDelete)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var d, _ = r.Delete(ctx, o.client, rci, &state)
	response.Diagnostics.Append(d...)
}
