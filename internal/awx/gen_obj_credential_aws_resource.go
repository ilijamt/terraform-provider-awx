package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	r "github.com/ilijamt/terraform-provider-awx/internal/resource"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &awsCredentialResource{}
	_ resource.ResourceWithConfigure   = &awsCredentialResource{}
	_ resource.ResourceWithImportState = &awsCredentialResource{}
)

// NewAWSCredentialResource is a helper function to simplify the provider implementation.
func NewAWSCredentialResource() resource.Resource {
	return &awsCredentialResource{}
}

type awsCredentialResource struct {
	client   c.Client
	rci      r.CallInfo
	endpoint string
	name     string
}

func (o *awsCredentialResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	o.rci = r.CallInfo{Name: "AWS", Endpoint: "/api/v2/credentials/", TypeName: "aws"}
	o.name = "AWS"
	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credentials/"
}

func (o *awsCredentialResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential_aws"
}

func (o *awsCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Database ID of this credential",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of this credential",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   false,
			},
			"description": schema.StringAttribute{
				Description: "Description of this credential",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   false,
			},
			"organization": schema.Int64Attribute{
				Description: "Organization of this credential",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   false,
			},
			"username": schema.StringAttribute{
				Description: "Access Key",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   false,
			},
			"password": schema.StringAttribute{
				Description: "Secret Key",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   true,
			},
			"security_token": schema.StringAttribute{
				Description: "STS Token",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   true,
			},
		},
	}
}

func (o *awsCredentialResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the AWS.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *awsCredentialResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
}
func (o *awsCredentialResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state awsCredentialTerraformModel
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

func (o *awsCredentialResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}

func (o *awsCredentialResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state awsCredentialTerraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeDelete)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var d, _ = r.Delete(ctx, o.client, rci, state.ID)
	response.Diagnostics.Append(d...)
}
