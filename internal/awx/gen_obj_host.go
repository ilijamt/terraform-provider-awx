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
	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// hostTerraformModel maps the schema for Host when using Data Source
type hostTerraformModel struct {
	// Description "Optional description of this host."
	Description types.String `tfsdk:"description" json:"description"`
	// Enabled "Is this host online and available for running jobs?"
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
	// ID "Database ID for this host."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InstanceId "The value used by the remote inventory source to uniquely identify the host"
	InstanceId types.String `tfsdk:"instance_id" json:"instance_id"`
	// Inventory ""
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// LastJob ""
	LastJob types.Int64 `tfsdk:"last_job" json:"last_job"`
	// LastJobHostSummary ""
	LastJobHostSummary types.Int64 `tfsdk:"last_job_host_summary" json:"last_job_host_summary"`
	// Name "Name of this host."
	Name types.String `tfsdk:"name" json:"name"`
	// Variables "Host variables in JSON or YAML format."
	Variables types.String `tfsdk:"variables" json:"variables"`
}

// Clone the object
func (o *hostTerraformModel) Clone() hostTerraformModel {
	return hostTerraformModel{
		Description:        o.Description,
		Enabled:            o.Enabled,
		ID:                 o.ID,
		InstanceId:         o.InstanceId,
		Inventory:          o.Inventory,
		LastJob:            o.LastJob,
		LastJobHostSummary: o.LastJobHostSummary,
		Name:               o.Name,
		Variables:          o.Variables,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Host
func (o *hostTerraformModel) BodyRequest() (req hostBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Enabled = o.Enabled.ValueBool()
	req.InstanceId = o.InstanceId.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.String())
	return
}

func (o *hostTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *hostTerraformModel) setEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Enabled, data)
}

func (o *hostTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *hostTerraformModel) setInstanceId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.InstanceId, data, false)
}

func (o *hostTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *hostTerraformModel) setLastJob(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.LastJob, data)
}

func (o *hostTerraformModel) setLastJobHostSummary(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.LastJobHostSummary, data)
}

func (o *hostTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *hostTerraformModel) setVariables(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.Variables, data, false)
}

func (o *hostTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setEnabled(data["enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInstanceId(data["instance_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventory(data["inventory"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastJob(data["last_job"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastJobHostSummary(data["last_job_host_summary"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVariables(data["variables"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// hostBodyRequestModel maps the schema for Host for creating and updating the data
type hostBodyRequestModel struct {
	// Description "Optional description of this host."
	Description string `json:"description,omitempty"`
	// Enabled "Is this host online and available for running jobs?"
	Enabled bool `json:"enabled"`
	// InstanceId "The value used by the remote inventory source to uniquely identify the host"
	InstanceId string `json:"instance_id,omitempty"`
	// Inventory ""
	Inventory int64 `json:"inventory"`
	// Name "Name of this host."
	Name string `json:"name"`
	// Variables "Host variables in JSON or YAML format."
	Variables json.RawMessage `json:"variables,omitempty"`
}

type hostObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}

var (
	_ datasource.DataSource              = &hostDataSource{}
	_ datasource.DataSourceWithConfigure = &hostDataSource{}
)

// NewHostDataSource is a helper function to instantiate the Host data source.
func NewHostDataSource() datasource.DataSource {
	return &hostDataSource{}
}

// hostDataSource is the data source implementation.
type hostDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *hostDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/hosts/"
}

// Metadata returns the data source type name.
func (o *hostDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host"
}

// GetSchema defines the schema for the data source.
func (o *hostDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Host",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"description": {
					Description: "Optional description of this host.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"enabled": {
					Description: "Is this host online and available for running jobs?",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this host.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
				"instance_id": {
					Description: "The value used by the remote inventory source to uniquely identify the host",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"inventory": {
					Description: "Inventory",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"last_job": {
					Description: "Last job",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"last_job_host_summary": {
					Description: "Last job host summary",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this host.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
				"variables": {
					Description: "Host variables in JSON or YAML format.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *hostDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state hostTerraformModel
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

	// Evaluate group 'by_name' based on the schema definition
	var groupByNameExists = func() bool {
		var groupByNameExists = true
		var paramsByName = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameExists = groupByNameExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByName = append(paramsByName, url.PathEscape(attrName.ValueString()))
		if groupByNameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s", paramsByName...))
		}
		return groupByNameExists
	}()
	searchDefined = searchDefined || groupByNameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for Host
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Host
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Host on %s", o.endpoint),
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
	_ resource.Resource                = &hostResource{}
	_ resource.ResourceWithConfigure   = &hostResource{}
	_ resource.ResourceWithImportState = &hostResource{}
)

// NewHostResource is a helper function to simplify the provider implementation.
func NewHostResource() resource.Resource {
	return &hostResource{}
}

type hostResource struct {
	client   c.Client
	endpoint string
}

func (o *hostResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/hosts/"
}

func (o *hostResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_host"
}

func (o *hostResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Host",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"description": {
					Description: "Optional description of this host.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"enabled": {
					Description: "Is this host online and available for running jobs?",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"instance_id": {
					Description: "The value used by the remote inventory source to uniquely identify the host",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"inventory": {
					Description:   "Inventory",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this host.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"variables": {
					Description: "Host variables in JSON or YAML format.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"id": {
					Description: "Database ID for this host.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"last_job": {
					Description: "",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"last_job_host_summary": {
					Description: "",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *hostResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Host.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *hostResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state hostTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Host
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[Host/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Host resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Host on %s", o.endpoint),
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

func (o *hostResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state hostTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Host
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Host from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Host on %s", o.endpoint),
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

func (o *hostResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state hostTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Host
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[Host/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Host resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Host on %s", o.endpoint),
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

func (o *hostResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state hostTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Host
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Host
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Host on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

var (
	_ datasource.DataSource              = &hostObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &hostObjectRolesDataSource{}
)

// NewHostObjectRolesDataSource is a helper function to instantiate the Host data source.
func NewHostObjectRolesDataSource() datasource.DataSource {
	return &hostObjectRolesDataSource{}
}

// hostObjectRolesDataSource is the data source implementation.
type hostObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *hostObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/hosts/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *hostObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host_object_roles"
}

// GetSchema defines the schema for the data source.
func (o *hostObjectRolesDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: helpers.SchemaVersion,
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "Host ID",
				Type:        types.Int64Type,
				Required:    true,
			},
			"roles": {
				Description: "Roles for host",
				Type:        types.MapType{ElemType: types.Int64Type},
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (o *hostObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state hostObjectRolesModel
	var err error
	var id types.Int64

	if d := req.Config.GetAttribute(ctx, path.Root("id"), &id); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	// Creates a new request for Credential
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(o.endpoint, id.ValueInt64()), nil); err != nil {
		resp.Diagnostics.AddError(
			"Unable to create a new request for host",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch the request for host object roles",
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for host",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var d diag.Diagnostics
	if state.Roles, d = types.MapValue(types.Int64Type, in); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

var (
	_ resource.Resource                = &hostAssociateDisassociateGroup{}
	_ resource.ResourceWithConfigure   = &hostAssociateDisassociateGroup{}
	_ resource.ResourceWithImportState = &hostAssociateDisassociateGroup{}
)

type hostAssociateDisassociateGroupTerraformModel struct {
	HostID  types.Int64 `tfsdk:"host_id"`
	GroupID types.Int64 `tfsdk:"group_id"`
}

// NewHostAssociateDisassociateGroupResource is a helper function to simplify the provider implementation.
func NewHostAssociateDisassociateGroupResource() resource.Resource {
	return &hostAssociateDisassociateGroup{}
}

type hostAssociateDisassociateGroup struct {
	client   c.Client
	endpoint string
}

func (o *hostAssociateDisassociateGroup) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/hosts/%d/groups/"
}

func (o hostAssociateDisassociateGroup) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_host_associate_group"
}

func (o hostAssociateDisassociateGroup) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Host/Associate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				"host_id": {
					Description: "Database ID for this Host.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("group_id"),
						),
					},
				},
				"group_id": {
					Description: "Database ID of the group to assign.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("host_id"),
						),
					},
				},
			},
		},
	), nil
}

func (o *hostAssociateDisassociateGroup) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state hostAssociateDisassociateGroupTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <host_id>/<group_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for Host association, invalid format.",
			err.Error(),
		)
		return
	}

	var hostId, groupId int64

	hostId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the hostId for the Host association.", request.ID),
			err.Error(),
		)
		return
	}
	state.HostID = types.Int64Value(hostId)

	groupId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the group_id for the Host association.", request.ID),
			err.Error(),
		)
		return
	}
	state.GroupID = types.Int64Value(groupId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *hostAssociateDisassociateGroup) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state hostAssociateDisassociateGroupTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of Host
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.HostID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.GroupID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[Host/Create/Associate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for create of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for Host on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.HostID = plan.HostID
	state.GroupID = plan.GroupID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *hostAssociateDisassociateGroup) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state hostAssociateDisassociateGroupTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of Host
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.HostID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.GroupID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[Host/Delete/Disassociate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for delete of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for Host on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *hostAssociateDisassociateGroup) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *hostAssociateDisassociateGroup) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
