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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// credentialInputSourceTerraformModel maps the schema for CredentialInputSource when using Data Source
type credentialInputSourceTerraformModel struct {
	// Description "Optional description of this credential input source."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this credential input source."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InputFieldName ""
	InputFieldName types.String `tfsdk:"input_field_name" json:"input_field_name"`
	// Metadata ""
	Metadata types.String `tfsdk:"metadata" json:"metadata"`
	// SourceCredential ""
	SourceCredential types.Int64 `tfsdk:"source_credential" json:"source_credential"`
	// TargetCredential ""
	TargetCredential types.Int64 `tfsdk:"target_credential" json:"target_credential"`
}

// Clone the object
func (o credentialInputSourceTerraformModel) Clone() credentialInputSourceTerraformModel {
	return credentialInputSourceTerraformModel{
		Description:      o.Description,
		ID:               o.ID,
		InputFieldName:   o.InputFieldName,
		Metadata:         o.Metadata,
		SourceCredential: o.SourceCredential,
		TargetCredential: o.TargetCredential,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for CredentialInputSource
func (o credentialInputSourceTerraformModel) BodyRequest() (req credentialInputSourceBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.InputFieldName = o.InputFieldName.ValueString()
	req.Metadata = json.RawMessage(o.Metadata.ValueString())
	req.SourceCredential = o.SourceCredential.ValueInt64()
	req.TargetCredential = o.TargetCredential.ValueInt64()
	return
}

func (o *credentialInputSourceTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *credentialInputSourceTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *credentialInputSourceTerraformModel) setInputFieldName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.InputFieldName, data, false)
}

func (o *credentialInputSourceTerraformModel) setMetadata(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.Metadata, data, false)
}

func (o *credentialInputSourceTerraformModel) setSourceCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SourceCredential, data)
}

func (o *credentialInputSourceTerraformModel) setTargetCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.TargetCredential, data)
}

func (o *credentialInputSourceTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInputFieldName(data["input_field_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMetadata(data["metadata"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSourceCredential(data["source_credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTargetCredential(data["target_credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// credentialInputSourceBodyRequestModel maps the schema for CredentialInputSource for creating and updating the data
type credentialInputSourceBodyRequestModel struct {
	// Description "Optional description of this credential input source."
	Description string `json:"description,omitempty"`
	// InputFieldName ""
	InputFieldName string `json:"input_field_name"`
	// Metadata ""
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// SourceCredential ""
	SourceCredential int64 `json:"source_credential"`
	// TargetCredential ""
	TargetCredential int64 `json:"target_credential"`
}

var (
	_ datasource.DataSource              = &credentialInputSourceDataSource{}
	_ datasource.DataSourceWithConfigure = &credentialInputSourceDataSource{}
)

// NewCredentialInputSourceDataSource is a helper function to instantiate the CredentialInputSource data source.
func NewCredentialInputSourceDataSource() datasource.DataSource {
	return &credentialInputSourceDataSource{}
}

// credentialInputSourceDataSource is the data source implementation.
type credentialInputSourceDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *credentialInputSourceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credential_input_sources/"
}

// Metadata returns the data source type name.
func (o *credentialInputSourceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credential_input_source"
}

// GetSchema defines the schema for the data source.
func (o *credentialInputSourceDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"CredentialInputSource",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"description": {
					Description: "Optional description of this credential input source.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this credential input source.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
						),
					},
				},
				"input_field_name": {
					Description: "Input field name",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"metadata": {
					Description: "Metadata",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"source_credential": {
					Description: "Source credential",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"target_credential": {
					Description: "Target credential",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *credentialInputSourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state credentialInputSourceTerraformModel
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
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for CredentialInputSource
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for CredentialInputSource on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for CredentialInputSource
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for CredentialInputSource on %s", o.endpoint),
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
	_ resource.Resource                = &credentialInputSourceResource{}
	_ resource.ResourceWithConfigure   = &credentialInputSourceResource{}
	_ resource.ResourceWithImportState = &credentialInputSourceResource{}
)

// NewCredentialInputSourceResource is a helper function to simplify the provider implementation.
func NewCredentialInputSourceResource() resource.Resource {
	return &credentialInputSourceResource{}
}

type credentialInputSourceResource struct {
	client   c.Client
	endpoint string
}

func (o *credentialInputSourceResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credential_input_sources/"
}

func (o credentialInputSourceResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential_input_source"
}

func (o credentialInputSourceResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"CredentialInputSource",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"description": {
					Description: "Optional description of this credential input source.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"input_field_name": {
					Description:   "Input field name",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"metadata": {
					Description: "Metadata",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"source_credential": {
					Description:   "Source credential",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"target_credential": {
					Description:   "Target credential",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"id": {
					Description: "Database ID for this credential input source.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *credentialInputSourceResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the CredentialInputSource.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *credentialInputSourceResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state credentialInputSourceTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for CredentialInputSource
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for CredentialInputSource on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new CredentialInputSource resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for CredentialInputSource on %s", o.endpoint),
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

func (o *credentialInputSourceResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state credentialInputSourceTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for CredentialInputSource
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for CredentialInputSource on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for CredentialInputSource from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for CredentialInputSource on %s", o.endpoint),
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

func (o *credentialInputSourceResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state credentialInputSourceTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for CredentialInputSource
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for CredentialInputSource on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new CredentialInputSource resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for CredentialInputSource on %s", o.endpoint),
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

func (o *credentialInputSourceResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state credentialInputSourceTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for CredentialInputSource
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for CredentialInputSource on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing CredentialInputSource
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for CredentialInputSource on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
