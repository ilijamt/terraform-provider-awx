package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// tokensTerraformModel maps the schema for Tokens when using Data Source
type tokensTerraformModel struct {
	// Application ""
	Application types.Int64 `tfsdk:"application" json:"application"`
	// Description "Optional description of this access token."
	Description types.String `tfsdk:"description" json:"description"`
	// Expires ""
	Expires types.String `tfsdk:"expires" json:"expires"`
	// ID "Database ID for this access token."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// RefreshToken ""
	RefreshToken types.String `tfsdk:"refresh_token" json:"refresh_token"`
	// Scope "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write']."
	Scope types.String `tfsdk:"scope" json:"scope"`
	// Token ""
	Token types.String `tfsdk:"token" json:"token"`
	// User "The user representing the token owner"
	User types.Int64 `tfsdk:"user" json:"user"`
}

// Clone the object
func (o tokensTerraformModel) Clone() tokensTerraformModel {
	return tokensTerraformModel{
		Application:  o.Application,
		Description:  o.Description,
		Expires:      o.Expires,
		ID:           o.ID,
		RefreshToken: o.RefreshToken,
		Scope:        o.Scope,
		Token:        o.Token,
		User:         o.User,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Tokens
func (o tokensTerraformModel) BodyRequest() (req tokensBodyRequestModel) {
	req.Application = o.Application.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Scope = o.Scope.ValueString()
	return
}

func (o *tokensTerraformModel) setApplication(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Application, data)
}

func (o *tokensTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *tokensTerraformModel) setExpires(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Expires, data, false)
}

func (o *tokensTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *tokensTerraformModel) setRefreshToken(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.RefreshToken, data, false)
}

func (o *tokensTerraformModel) setScope(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Scope, data, false)
}

func (o *tokensTerraformModel) setToken(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Token, data, false)
}

func (o *tokensTerraformModel) setUser(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.User, data)
}

func (o *tokensTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setApplication(data["application"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExpires(data["expires"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRefreshToken(data["refresh_token"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScope(data["scope"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setToken(data["token"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUser(data["user"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// tokensBodyRequestModel maps the schema for Tokens for creating and updating the data
type tokensBodyRequestModel struct {
	// Application ""
	Application int64 `json:"application,omitempty"`
	// Description "Optional description of this access token."
	Description string `json:"description,omitempty"`
	// Scope "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write']."
	Scope string `json:"scope,omitempty"`
}

var (
	_ datasource.DataSource              = &tokensDataSource{}
	_ datasource.DataSourceWithConfigure = &tokensDataSource{}
)

// NewTokensDataSource is a helper function to instantiate the Tokens data source.
func NewTokensDataSource() datasource.DataSource {
	return &tokensDataSource{}
}

// tokensDataSource is the data source implementation.
type tokensDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *tokensDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/tokens/"
}

// Metadata returns the data source type name.
func (o *tokensDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_token"
}

// GetSchema defines the schema for the data source.
func (o *tokensDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Tokens",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"application": {
					Description: "Application",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this access token.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"expires": {
					Description: "Expires",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this access token.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
						),
					},
				},
				"refresh_token": {
					Description: "Refresh token",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scope": {
					Description: "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"token": {
					Description: "Token",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"user": {
					Description: "The user representing the token owner",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *tokensDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state tokensTerraformModel
	var err error
	var endpoint string
	var searchDefined bool

	// Only one group should evaluate to True, terraform should prevent from being able to set
	// the conflicting groups

	// Evaluate group 'by_id' based on the schema definition
	var groupByIdExists = func() bool {
		var groupByIdExists = true
		var paramsById = []any{o.endpoint}
		var attrID types.Int64
		req.Config.GetAttribute(ctx, path.Root("id"), &attrID)
		groupByIdExists = groupByIdExists && (!attrID.IsNull() && !attrID.IsUnknown())
		paramsById = append(paramsById, attrID.ValueInt64())
		if groupByIdExists {
			endpoint = p.Clean(fmt.Sprintf("%s/%d/", paramsById...))
		}
		return groupByIdExists
	}()
	searchDefined = searchDefined || groupByIdExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			fmt.Sprintf("missing configuration for one of the predefined search groups"),
			detailMessage,
		)
		return
	}

	// Creates a new request for Tokens
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Tokens on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Tokens
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Tokens on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = extractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &tokensResource{}
	_ resource.ResourceWithConfigure   = &tokensResource{}
	_ resource.ResourceWithImportState = &tokensResource{}
)

// NewTokensResource is a helper function to simplify the provider implementation.
func NewTokensResource() resource.Resource {
	return &tokensResource{}
}

type tokensResource struct {
	client   c.Client
	endpoint string
}

func (o *tokensResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/tokens/"
}

func (o tokensResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_token"
}

func (o tokensResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Tokens",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"application": {
					Description: "Application",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this access token.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scope": {
					Description: "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`write`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"expires": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this access token.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"refresh_token": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"token": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"user": {
					Description: "The user representing the token owner",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *tokensResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Tokens.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *tokensResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state tokensTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Tokens
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Tokens on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Tokens resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Tokens on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *tokensResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state tokensTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Tokens
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Tokens on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Tokens from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Tokens on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *tokensResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state tokensTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Tokens
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Tokens on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Tokens resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Tokens on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *tokensResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state tokensTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Tokens
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Tokens on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Tokens
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Tokens on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
