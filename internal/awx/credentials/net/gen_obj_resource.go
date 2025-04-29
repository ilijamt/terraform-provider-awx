package net

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
	_ resource.Resource                     = &TerraformResource{}
	_ resource.ResourceWithConfigValidators = &TerraformResource{}
	_ resource.ResourceWithConfigure        = &TerraformResource{}
	_ resource.ResourceWithImportState      = &TerraformResource{}
)

// NewResource is a helper function to simplify the provider implementation.
func NewResource() resource.Resource {
	return &TerraformResource{}
}

type TerraformResource struct {
	client           c.Client
	rci              r.CallInfo
	endpoint         string
	name             string
	typeName         string
	credentialTypeId int64
}

func (o *TerraformResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	o.name = "Network"
	o.typeName = "net"
	o.rci = r.CallInfo{Name: o.name, TypeName: o.typeName, Endpoint: "/api/v2/credentials/"}
	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credentials/"
	o.credentialTypeId = 4
}

func (o *TerraformResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential_net"
}

func (o *TerraformResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("organization"),
		),
	}
}

func (o *TerraformResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description: "Username",
				Required:    true,
				Optional:    false,
				Computed:    false,
				Sensitive:   false,
				WriteOnly:   false,
			},
			"password": schema.StringAttribute{
				Description: "Password",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   true,
				WriteOnly:   false,
			},
			"ssh_key_data": schema.StringAttribute{
				Description: "SSH Private Key",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   true,
				WriteOnly:   false,
			},
			"ssh_key_unlock": schema.StringAttribute{
				Description: "Private Key Passphrase",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   true,
				WriteOnly:   false,
			},
			"authorize": schema.BoolAttribute{
				Description: "Authorize",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   false,
				WriteOnly:   false,
			},
			"authorize_password": schema.StringAttribute{
				Description: "Authorize Password",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   true,
				WriteOnly:   false,
			},
		},
	}
}

func (o *TerraformResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the net.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *TerraformResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var rci = o.rci.With(r.SourceResource, r.CalleeCreate)
	var plan TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}
	if plan.Organization.IsNull() || plan.Organization.IsUnknown() {
		user, err := o.client.User(ctx)
		if err != nil {
			response.Diagnostics.AddError("failed to retrieve current user", err.Error())
			return
		}
		plan.userId = user.ID
	}
	plan.credentialTypeId = o.credentialTypeId
	var d, _ = r.Create(ctx, o.client, rci, &plan)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (o *TerraformResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state TerraformModel
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

func (o *TerraformResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var rci = o.rci.With(r.SourceResource, r.CalleeUpdate)
	var plan TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}
	plan.credentialTypeId = o.credentialTypeId
	var d, _ = r.Update(ctx, o.client, rci, &plan)
	if d.HasError() {
		response.Diagnostics.Append(d...)
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (o *TerraformResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state TerraformModel
	var rci = o.rci.With(r.SourceResource, r.CalleeDelete)
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var d, _ = r.Delete(ctx, o.client, rci, &state)
	response.Diagnostics.Append(d...)
}
