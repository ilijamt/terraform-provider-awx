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

// applicationTerraformModel maps the schema for Application when using Data Source
type applicationTerraformModel struct {
	// AuthorizationGrantType "The Grant type the user must use for acquire tokens for this application."
	AuthorizationGrantType types.String `tfsdk:"authorization_grant_type" json:"authorization_grant_type"`
	// ClientId ""
	ClientId types.String `tfsdk:"client_id" json:"client_id"`
	// ClientSecret "Used for more stringent verification of access to an application when creating a token."
	ClientSecret types.String `tfsdk:"client_secret" json:"client_secret"`
	// ClientType "Set to Public or Confidential depending on how secure the client device is."
	ClientType types.String `tfsdk:"client_type" json:"client_type"`
	// Description "Optional description of this application."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this application."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this application."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Organization containing this application."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// RedirectUris "Allowed URIs list, space separated"
	RedirectUris types.String `tfsdk:"redirect_uris" json:"redirect_uris"`
	// SkipAuthorization "Set True to skip authorization step for completely trusted applications."
	SkipAuthorization types.Bool `tfsdk:"skip_authorization" json:"skip_authorization"`
}

// Clone the object
func (o applicationTerraformModel) Clone() applicationTerraformModel {
	return applicationTerraformModel{
		AuthorizationGrantType: o.AuthorizationGrantType,
		ClientId:               o.ClientId,
		ClientSecret:           o.ClientSecret,
		ClientType:             o.ClientType,
		Description:            o.Description,
		ID:                     o.ID,
		Name:                   o.Name,
		Organization:           o.Organization,
		RedirectUris:           o.RedirectUris,
		SkipAuthorization:      o.SkipAuthorization,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Application
func (o applicationTerraformModel) BodyRequest() (req applicationBodyRequestModel) {
	req.AuthorizationGrantType = o.AuthorizationGrantType.ValueString()
	req.ClientType = o.ClientType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.RedirectUris = o.RedirectUris.ValueString()
	req.SkipAuthorization = o.SkipAuthorization.ValueBool()
	return
}

func (o *applicationTerraformModel) setAuthorizationGrantType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AuthorizationGrantType, data, false)
}

func (o *applicationTerraformModel) setClientId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ClientId, data, false)
}

func (o *applicationTerraformModel) setClientSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ClientSecret, data, false)
}

func (o *applicationTerraformModel) setClientType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ClientType, data, false)
}

func (o *applicationTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *applicationTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *applicationTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *applicationTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *applicationTerraformModel) setRedirectUris(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.RedirectUris, data, false)
}

func (o *applicationTerraformModel) setSkipAuthorization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SkipAuthorization, data)
}

func (o *applicationTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAuthorizationGrantType(data["authorization_grant_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setClientId(data["client_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setClientSecret(data["client_secret"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setClientType(data["client_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRedirectUris(data["redirect_uris"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSkipAuthorization(data["skip_authorization"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// applicationBodyRequestModel maps the schema for Application for creating and updating the data
type applicationBodyRequestModel struct {
	// AuthorizationGrantType "The Grant type the user must use for acquire tokens for this application."
	AuthorizationGrantType string `json:"authorization_grant_type"`
	// ClientType "Set to Public or Confidential depending on how secure the client device is."
	ClientType string `json:"client_type"`
	// Description "Optional description of this application."
	Description string `json:"description,omitempty"`
	// Name "Name of this application."
	Name string `json:"name"`
	// Organization "Organization containing this application."
	Organization int64 `json:"organization"`
	// RedirectUris "Allowed URIs list, space separated"
	RedirectUris string `json:"redirect_uris,omitempty"`
	// SkipAuthorization "Set True to skip authorization step for completely trusted applications."
	SkipAuthorization bool `json:"skip_authorization"`
}

var (
	_ datasource.DataSource              = &applicationDataSource{}
	_ datasource.DataSourceWithConfigure = &applicationDataSource{}
)

// NewApplicationDataSource is a helper function to instantiate the Application data source.
func NewApplicationDataSource() datasource.DataSource {
	return &applicationDataSource{}
}

// applicationDataSource is the data source implementation.
type applicationDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *applicationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/applications/"
}

// Metadata returns the data source type name.
func (o *applicationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

// GetSchema defines the schema for the data source.
func (o *applicationDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Application",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"authorization_grant_type": {
					Description: "The Grant type the user must use for acquire tokens for this application.",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"authorization-code", "password"}...),
					},
				},
				"client_id": {
					Description: "Client id",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"client_secret": {
					Description: "Used for more stringent verification of access to an application when creating a token.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"client_type": {
					Description: "Set to Public or Confidential depending on how secure the client device is.",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"confidential", "public"}...),
					},
				},
				"description": {
					Description: "Optional description of this application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this application.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ConflictsWith(
							path.MatchRoot("name"),
							path.MatchRoot("organization"),
						),
					},
				},
				"name": {
					Description: "Name of this application.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("organization"),
						),
						schemavalidator.ConflictsWith(
							path.MatchRoot("id"),
						),
					},
				},
				"organization": {
					Description: "Organization containing this application.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("name"),
						),
						schemavalidator.ConflictsWith(
							path.MatchRoot("id"),
						),
					},
				},
				"redirect_uris": {
					Description: "Allowed URIs list, space separated",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"skip_authorization": {
					Description: "Set True to skip authorization step for completely trusted applications.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *applicationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state applicationTerraformModel
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

	// Evaluate group 'by_name_organization' based on the schema definition
	var groupByNameOrganizationExists = func() bool {
		var groupByNameOrganizationExists = true
		var paramsByNameOrganization = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameOrganizationExists = groupByNameOrganizationExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByNameOrganization = append(paramsByNameOrganization, url.PathEscape(attrName.ValueString()))
		var attrOrganization types.Int64
		req.Config.GetAttribute(ctx, path.Root("organization"), &attrOrganization)
		groupByNameOrganizationExists = groupByNameOrganizationExists && (!attrOrganization.IsNull() && !attrOrganization.IsUnknown())
		paramsByNameOrganization = append(paramsByNameOrganization, attrOrganization.ValueInt64())
		if groupByNameOrganizationExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s&organization=%d", paramsByNameOrganization...))
		}
		return groupByNameOrganizationExists
	}()
	searchDefined = searchDefined || groupByNameOrganizationExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for Application
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Application on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Application
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Application on %s", o.endpoint),
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
	if err = hookApplication(ctx, SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on Application",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &applicationResource{}
	_ resource.ResourceWithConfigure   = &applicationResource{}
	_ resource.ResourceWithImportState = &applicationResource{}
)

// NewApplicationResource is a helper function to simplify the provider implementation.
func NewApplicationResource() resource.Resource {
	return &applicationResource{}
}

type applicationResource struct {
	client   c.Client
	endpoint string
}

func (o *applicationResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/applications/"
}

func (o applicationResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_application"
}

func (o applicationResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Application",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"authorization_grant_type": {
					Description:   "The Grant type the user must use for acquire tokens for this application.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"authorization-code", "password"}...),
					},
				},
				"client_type": {
					Description:   "Set to Public or Confidential depending on how secure the client device is.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"confidential", "public"}...),
					},
				},
				"description": {
					Description: "Optional description of this application.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this application.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(255),
					},
				},
				"organization": {
					Description:   "Organization containing this application.",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"redirect_uris": {
					Description: "Allowed URIs list, space separated",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"skip_authorization": {
					Description: "Set True to skip authorization step for completely trusted applications.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"client_id": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"client_secret": {
					Description: "Used for more stringent verification of access to an application when creating a token.",
					Computed:    true,
					Type:        types.StringType,
					Sensitive:   true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this application.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *applicationResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Application.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *applicationResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state applicationTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Application
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Application on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Application resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Application on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookApplication(ctx, SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on Application",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *applicationResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state applicationTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for Application
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Application on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Application from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Application on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookApplication(ctx, SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on Application",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *applicationResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state applicationTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Application
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Application on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Application resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Application on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookApplication(ctx, SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on Application",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *applicationResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state applicationTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Application
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Application on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Application
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Application on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
