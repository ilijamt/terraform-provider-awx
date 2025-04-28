package aws

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	r "github.com/ilijamt/terraform-provider-awx/internal/resource"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                     = &awsCredentialResource{}
	_ resource.ResourceWithConfigValidators = &awsCredentialResource{}
	_ resource.ResourceWithConfigure        = &awsCredentialResource{}
	_ resource.ResourceWithImportState      = &awsCredentialResource{}
)

// NewResource is a helper function to simplify the provider implementation.
func NewResource() resource.Resource {
	return &awsCredentialResource{}
}

type awsCredentialResource struct {
	client           c.Client
	rci              r.CallInfo
	endpoint         string
	name             string
	typeName         string
	credentialTypeId int
}

func (o *awsCredentialResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	o.name = "Amazon Web Services"
	o.typeName = "aws"
	o.rci = r.CallInfo{Name: o.name, TypeName: o.typeName, Endpoint: "/api/v2/credentials/"}
	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credentials/"
	o.credentialTypeId = 5
}

func (o *awsCredentialResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential_aws"
}

func (o *awsCredentialResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("organization"),
		),
	}
}

func (o *awsCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Database ID of this credential",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of this credential",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   false,
				WriteOnly:   false,
			},
			"description": schema.StringAttribute{
				Description: "Description of this credential",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   false,
				WriteOnly:   false,
			},
			"organization": schema.Int64Attribute{
				Description: "Organization of this credential",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   false,
				WriteOnly:   false,
			},
			"username": schema.StringAttribute{
				Description: "Access Key",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   false,
				WriteOnly:   false,
			},
			"password": schema.StringAttribute{
				Description: "Secret Key",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   true,
				WriteOnly:   false,
			},
			"security_token": schema.StringAttribute{
				Description: "STS Token",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   true,
				WriteOnly:   false,
			},
		},
	}
}

func (o *awsCredentialResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the aws.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *awsCredentialResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var rci = o.rci.With(r.SourceResource, r.CalleeCreate)
	var plan terraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// fetch the UserID if organization is not set
	if plan.Organization.IsNull() || plan.Organization.IsUnknown() {
		user, err := o.client.User(ctx)
		if err != nil {
			response.Diagnostics.AddError("failed to retrieve current user", err.Error())
			return
		}
		plan.userId = user.ID
	}

	var d, _ = r.Create(ctx, o.client, rci, &plan)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (o *awsCredentialResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state terraformModel
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

func (o *awsCredentialResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var rci = o.rci.With(r.SourceResource, r.CalleeUpdate)
	var plan terraformModel
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

func (o *awsCredentialResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state terraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeDelete)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var d, _ = r.Delete(ctx, o.client, rci, &state)
	response.Diagnostics.Append(d...)
}
