package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	p "path"
	"strconv"
	"strings"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// userTerraformModel maps the schema for User when using Data Source
type userTerraformModel struct {
	// Email ""
	Email types.String `tfsdk:"email" json:"email"`
	// ExternalAccount "Set if the account is managed by an external service"
	ExternalAccount types.String `tfsdk:"external_account" json:"external_account"`
	// FirstName ""
	FirstName types.String `tfsdk:"first_name" json:"first_name"`
	// ID "Database ID for this user."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// IsSuperuser "Designates that this user has all permissions without explicitly assigning them."
	IsSuperuser types.Bool `tfsdk:"is_superuser" json:"is_superuser"`
	// IsSystemAuditor ""
	IsSystemAuditor types.Bool `tfsdk:"is_system_auditor" json:"is_system_auditor"`
	// LastLogin ""
	LastLogin types.String `tfsdk:"last_login" json:"last_login"`
	// LastName ""
	LastName types.String `tfsdk:"last_name" json:"last_name"`
	// LdapDn ""
	LdapDn types.String `tfsdk:"ldap_dn" json:"ldap_dn"`
	// Username "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
	Username types.String `tfsdk:"username" json:"username"`
	// Password "Write-only field used to change the password."
	Password types.String `tfsdk:"password" json:"password"`
}

// Clone the object
func (o userTerraformModel) Clone() userTerraformModel {
	return userTerraformModel{
		Email:           o.Email,
		ExternalAccount: o.ExternalAccount,
		FirstName:       o.FirstName,
		ID:              o.ID,
		IsSuperuser:     o.IsSuperuser,
		IsSystemAuditor: o.IsSystemAuditor,
		LastLogin:       o.LastLogin,
		LastName:        o.LastName,
		LdapDn:          o.LdapDn,
		Username:        o.Username,
		Password:        o.Password,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for User
func (o userTerraformModel) BodyRequest() (req userBodyRequestModel) {
	req.Email = o.Email.ValueString()
	req.FirstName = o.FirstName.ValueString()
	req.IsSuperuser = o.IsSuperuser.ValueBool()
	req.IsSystemAuditor = o.IsSystemAuditor.ValueBool()
	req.LastName = o.LastName.ValueString()
	req.Username = o.Username.ValueString()
	return
}

func (o *userTerraformModel) setEmail(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Email, data, false)
}

func (o *userTerraformModel) setExternalAccount(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ExternalAccount, data, false)
}

func (o *userTerraformModel) setFirstName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.FirstName, data, false)
}

func (o *userTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *userTerraformModel) setIsSuperuser(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsSuperuser, data)
}

func (o *userTerraformModel) setIsSystemAuditor(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsSystemAuditor, data)
}

func (o *userTerraformModel) setLastLogin(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LastLogin, data, false)
}

func (o *userTerraformModel) setLastName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LastName, data, false)
}

func (o *userTerraformModel) setLdapDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LdapDn, data, false)
}

func (o *userTerraformModel) setUsername(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *userTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setEmail(data["email"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExternalAccount(data["external_account"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setFirstName(data["first_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsSuperuser(data["is_superuser"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsSystemAuditor(data["is_system_auditor"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastLogin(data["last_login"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastName(data["last_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLdapDn(data["ldap_dn"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUsername(data["username"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// userBodyRequestModel maps the schema for User for creating and updating the data
type userBodyRequestModel struct {
	// Email ""
	Email string `json:"email,omitempty"`
	// FirstName ""
	FirstName string `json:"first_name,omitempty"`
	// IsSuperuser "Designates that this user has all permissions without explicitly assigning them."
	IsSuperuser bool `json:"is_superuser"`
	// IsSystemAuditor ""
	IsSystemAuditor bool `json:"is_system_auditor"`
	// LastName ""
	LastName string `json:"last_name,omitempty"`
	// Username "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
	Username string `json:"username"`
	// Password "Write-only field used to change the password."
	Password string `json:"password,omitempty"`
}

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

// NewUserDataSource is a helper function to instantiate the User data source.
func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

// userDataSource is the data source implementation.
type userDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/users/"
}

// Metadata returns the data source type name.
func (o *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// GetSchema defines the schema for the data source.
func (o *userDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"User",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"email": {
					Description: "Email address",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"external_account": {
					Description: "Set if the account is managed by an external service",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"first_name": {
					Description: "First name",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this user.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("username"),
						),
					},
				},
				"is_superuser": {
					Description: "Designates that this user has all permissions without explicitly assigning them.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"is_system_auditor": {
					Description: "Is system auditor",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"last_login": {
					Description: "Last login",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"last_name": {
					Description: "Last name",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ldap_dn": {
					Description: "Ldap dn",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"username": {
					Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("username"),
						),
					},
				},
				// Write only elements
				"password": {
					Description: "Write-only field used to change the password.",
					Type:        types.StringType,
					Computed:    true,
					Sensitive:   true,
				},
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state userTerraformModel
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

	// Evaluate group 'by_username' based on the schema definition
	var groupByUsernameExists = func() bool {
		var groupByUsernameExists = true
		var paramsByUsername = []any{o.endpoint}
		var attrUsername types.String
		req.Config.GetAttribute(ctx, path.Root("username"), &attrUsername)
		groupByUsernameExists = groupByUsernameExists && (!attrUsername.IsNull() && !attrUsername.IsUnknown())
		paramsByUsername = append(paramsByUsername, url.PathEscape(attrUsername.ValueString()))
		if groupByUsernameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?username__exact=%s", paramsByUsername...))
		}
		return groupByUsernameExists
	}()
	searchDefined = searchDefined || groupByUsernameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			fmt.Sprintf("missing configuration for one of the predefined search groups"),
			detailMessage,
		)
		return
	}

	// Creates a new request for User
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for User
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for User on %s", o.endpoint),
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
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
)

// NewUserResource is a helper function to simplify the provider implementation.
func NewUserResource() resource.Resource {
	return &userResource{}
}

type userResource struct {
	client   c.Client
	endpoint string
}

func (o *userResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/users/"
}

func (o userResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user"
}

func (o userResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"User",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"email": {
					Description: "Email address",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(254),
					},
				},
				"first_name": {
					Description: "First name",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(150),
					},
				},
				"is_superuser": {
					Description: "Designates that this user has all permissions without explicitly assigning them.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"is_system_auditor": {
					Description: "Is system auditor",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"last_name": {
					Description: "Last name",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(150),
					},
				},
				"username": {
					Description:   "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(150),
					},
				},
				// Write only elements
				"password": {
					Description:   "Write-only field used to change the password.",
					Type:          types.StringType,
					Sensitive:     true,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				// Data only elements
				"external_account": {
					Description: "Set if the account is managed by an external service",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this user.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"last_login": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"ldap_dn": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *userResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the User.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *userResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state userTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for User
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	bodyRequest.Password = plan.Password.ValueString()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new User resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for User on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	state.Password = types.StringValue(plan.Password.ValueString())

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *userResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state userTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for User
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for User from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for User on %s", o.endpoint),
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

func (o *userResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state userTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for User
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	bodyRequest.Password = plan.Password.ValueString()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new User resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for User on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	state.Password = types.StringValue(plan.Password.ValueString())

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *userResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state userTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for User
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing User
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for User on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

var (
	_ resource.Resource                = &userAssociateDisassociateRole{}
	_ resource.ResourceWithConfigure   = &userAssociateDisassociateRole{}
	_ resource.ResourceWithImportState = &userAssociateDisassociateRole{}
)

type userAssociateDisassociateRoleTerraformModel struct {
	UserID types.Int64 `tfsdk:"user_id"`
	RoleID types.Int64 `tfsdk:"role_id"`
}

// NewUserAssociateDisassociateRoleResource is a helper function to simplify the provider implementation.
func NewUserAssociateDisassociateRoleResource() resource.Resource {
	return &userAssociateDisassociateRole{}
}

type userAssociateDisassociateRole struct {
	client   c.Client
	endpoint string
}

func (o *userAssociateDisassociateRole) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/users/%d/roles/"
}

func (o userAssociateDisassociateRole) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user_associate_role"
}

func (o userAssociateDisassociateRole) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"User/Associate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				"user_id": {
					Description: "Database ID for this User.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("role_id"),
						),
					},
				},
				"role_id": {
					Description: "Database ID of the role to assign.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("user_id"),
						),
					},
				},
			},
		},
	), nil
}

func (o *userAssociateDisassociateRole) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state userAssociateDisassociateRoleTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <user_id>/<role_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to import state for User association, invalid format."),
			err.Error(),
		)
		return
	}

	var userId, roleId int64

	userId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the userId for the User association.", request.ID),
			err.Error(),
		)
		return
	}
	state.UserID = types.Int64Value(userId)

	roleId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the role_id for the User association.", request.ID),
			err.Error(),
		)
		return
	}
	state.RoleID = types.Int64Value(roleId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *userAssociateDisassociateRole) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state userAssociateDisassociateRoleTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of User
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.UserID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.RoleID.ValueInt64(), Disassociate: false}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for create of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for User on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.UserID = plan.UserID
	state.RoleID = plan.RoleID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *userAssociateDisassociateRole) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state userAssociateDisassociateRoleTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of User
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.UserID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.RoleID.ValueInt64(), Disassociate: true}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for delete of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for User on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *userAssociateDisassociateRole) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *userAssociateDisassociateRole) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
