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
)

// inventoryTerraformModel maps the schema for Inventory when using Data Source
type inventoryTerraformModel struct {
	// Description "Optional description of this inventory."
	Description types.String `tfsdk:"description" json:"description"`
	// HasActiveFailures "This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed."
	HasActiveFailures types.Bool `tfsdk:"has_active_failures" json:"has_active_failures"`
	// HasInventorySources "This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources."
	HasInventorySources types.Bool `tfsdk:"has_inventory_sources" json:"has_inventory_sources"`
	// HostFilter "Filter that will be applied to the hosts of this inventory."
	HostFilter types.String `tfsdk:"host_filter" json:"host_filter"`
	// HostsWithActiveFailures "This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures."
	HostsWithActiveFailures types.Int64 `tfsdk:"hosts_with_active_failures" json:"hosts_with_active_failures"`
	// ID "Database ID for this inventory."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InventorySourcesWithFailures "Number of external inventory sources in this inventory with failures."
	InventorySourcesWithFailures types.Int64 `tfsdk:"inventory_sources_with_failures" json:"inventory_sources_with_failures"`
	// Kind "Kind of inventory being represented."
	Kind types.String `tfsdk:"kind" json:"kind"`
	// Name "Name of this inventory."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Organization containing this inventory."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// PendingDeletion "Flag indicating the inventory is being deleted."
	PendingDeletion types.Bool `tfsdk:"pending_deletion" json:"pending_deletion"`
	// PreventInstanceGroupFallback "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback types.Bool `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	// TotalGroups "This field is deprecated and will be removed in a future release. Total number of groups in this inventory."
	TotalGroups types.Int64 `tfsdk:"total_groups" json:"total_groups"`
	// TotalHosts "This field is deprecated and will be removed in a future release. Total number of hosts in this inventory."
	TotalHosts types.Int64 `tfsdk:"total_hosts" json:"total_hosts"`
	// TotalInventorySources "Total number of external inventory sources configured within this inventory."
	TotalInventorySources types.Int64 `tfsdk:"total_inventory_sources" json:"total_inventory_sources"`
	// Variables "Inventory variables in JSON or YAML format."
	Variables types.String `tfsdk:"variables" json:"variables"`
}

// Clone the object
func (o inventoryTerraformModel) Clone() inventoryTerraformModel {
	return inventoryTerraformModel{
		Description:                  o.Description,
		HasActiveFailures:            o.HasActiveFailures,
		HasInventorySources:          o.HasInventorySources,
		HostFilter:                   o.HostFilter,
		HostsWithActiveFailures:      o.HostsWithActiveFailures,
		ID:                           o.ID,
		InventorySourcesWithFailures: o.InventorySourcesWithFailures,
		Kind:                         o.Kind,
		Name:                         o.Name,
		Organization:                 o.Organization,
		PendingDeletion:              o.PendingDeletion,
		PreventInstanceGroupFallback: o.PreventInstanceGroupFallback,
		TotalGroups:                  o.TotalGroups,
		TotalHosts:                   o.TotalHosts,
		TotalInventorySources:        o.TotalInventorySources,
		Variables:                    o.Variables,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Inventory
func (o inventoryTerraformModel) BodyRequest() (req inventoryBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.HostFilter = o.HostFilter.ValueString()
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.Variables = json.RawMessage(o.Variables.ValueString())
	return
}

func (o *inventoryTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	// Decode "description"
	if val, ok := data.(string); ok {
		o.Description = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Description = types.StringValue(val.String())
	} else {
		o.Description = types.StringNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setHasActiveFailures(data any) (d diag.Diagnostics, err error) {
	// Decode "has_active_failures"
	if val, ok := data.(bool); ok {
		o.HasActiveFailures = types.BoolValue(val)
	} else {
		o.HasActiveFailures = types.BoolNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setHasInventorySources(data any) (d diag.Diagnostics, err error) {
	// Decode "has_inventory_sources"
	if val, ok := data.(bool); ok {
		o.HasInventorySources = types.BoolValue(val)
	} else {
		o.HasInventorySources = types.BoolNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setHostFilter(data any) (d diag.Diagnostics, err error) {
	// Decode "host_filter"
	if val, ok := data.(string); ok {
		o.HostFilter = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.HostFilter = types.StringValue(val.String())
	} else {
		o.HostFilter = types.StringNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setHostsWithActiveFailures(data any) (d diag.Diagnostics, err error) {
	// Decode "hosts_with_active_failures"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.HostsWithActiveFailures = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.HostsWithActiveFailures = types.Int64Value(val)
	} else {
		o.HostsWithActiveFailures = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	// Decode "id"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.ID = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.ID = types.Int64Value(val)
	} else {
		o.ID = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setInventorySourcesWithFailures(data any) (d diag.Diagnostics, err error) {
	// Decode "inventory_sources_with_failures"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.InventorySourcesWithFailures = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.InventorySourcesWithFailures = types.Int64Value(val)
	} else {
		o.InventorySourcesWithFailures = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setKind(data any) (d diag.Diagnostics, err error) {
	// Decode "kind"
	if val, ok := data.(string); ok {
		o.Kind = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Kind = types.StringValue(val.String())
	} else {
		o.Kind = types.StringNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	// Decode "name"
	if val, ok := data.(string); ok {
		o.Name = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Name = types.StringValue(val.String())
	} else {
		o.Name = types.StringNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	// Decode "organization"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Organization = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Organization = types.Int64Value(val)
	} else {
		o.Organization = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setPendingDeletion(data any) (d diag.Diagnostics, err error) {
	// Decode "pending_deletion"
	if val, ok := data.(bool); ok {
		o.PendingDeletion = types.BoolValue(val)
	} else {
		o.PendingDeletion = types.BoolNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setPreventInstanceGroupFallback(data any) (d diag.Diagnostics, err error) {
	// Decode "prevent_instance_group_fallback"
	if val, ok := data.(bool); ok {
		o.PreventInstanceGroupFallback = types.BoolValue(val)
	} else {
		o.PreventInstanceGroupFallback = types.BoolNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setTotalGroups(data any) (d diag.Diagnostics, err error) {
	// Decode "total_groups"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.TotalGroups = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.TotalGroups = types.Int64Value(val)
	} else {
		o.TotalGroups = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setTotalHosts(data any) (d diag.Diagnostics, err error) {
	// Decode "total_hosts"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.TotalHosts = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.TotalHosts = types.Int64Value(val)
	} else {
		o.TotalHosts = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setTotalInventorySources(data any) (d diag.Diagnostics, err error) {
	// Decode "total_inventory_sources"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.TotalInventorySources = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.TotalInventorySources = types.Int64Value(val)
	} else {
		o.TotalInventorySources = types.Int64Null()
	}
	return d, nil
}

func (o *inventoryTerraformModel) setVariables(data any) (d diag.Diagnostics, err error) {
	// Decode "variables"
	if val, ok := data.(string); ok {
		o.Variables = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(map[string]any); ok {
		var v []byte
		if v, err = json.Marshal(val); err != nil {
			d.AddError(
				fmt.Sprintf("failed to decode map"),
				err.Error(),
			)
			return
		}
		o.Variables = types.StringValue(helpers.TrimString(false, false, string(v)))
	} else {
		o.Variables = types.StringNull()
	}
	return d, nil
}

func (o *inventoryTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHasActiveFailures(data["has_active_failures"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHasInventorySources(data["has_inventory_sources"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHostFilter(data["host_filter"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHostsWithActiveFailures(data["hosts_with_active_failures"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventorySourcesWithFailures(data["inventory_sources_with_failures"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setKind(data["kind"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPendingDeletion(data["pending_deletion"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPreventInstanceGroupFallback(data["prevent_instance_group_fallback"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTotalGroups(data["total_groups"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTotalHosts(data["total_hosts"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTotalInventorySources(data["total_inventory_sources"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVariables(data["variables"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// inventoryBodyRequestModel maps the schema for Inventory for creating and updating the data
type inventoryBodyRequestModel struct {
	// Description "Optional description of this inventory."
	Description string `json:"description,omitempty"`
	// HostFilter "Filter that will be applied to the hosts of this inventory."
	HostFilter string `json:"host_filter,omitempty"`
	// Kind "Kind of inventory being represented."
	Kind string `json:"kind,omitempty"`
	// Name "Name of this inventory."
	Name string `json:"name"`
	// Organization "Organization containing this inventory."
	Organization int64 `json:"organization"`
	// PreventInstanceGroupFallback "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback bool `json:"prevent_instance_group_fallback"`
	// Variables "Inventory variables in JSON or YAML format."
	Variables json.RawMessage `json:"variables,omitempty"`
}

type inventoryObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}

var (
	_ datasource.DataSource              = &inventoryDataSource{}
	_ datasource.DataSourceWithConfigure = &inventoryDataSource{}
)

// NewInventoryDataSource is a helper function to instantiate the Inventory data source.
func NewInventoryDataSource() datasource.DataSource {
	return &inventoryDataSource{}
}

// inventoryDataSource is the data source implementation.
type inventoryDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *inventoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventories/"
}

// Metadata returns the data source type name.
func (o *inventoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory"
}

// GetSchema defines the schema for the data source.
func (o *inventoryDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Inventory",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"description": {
					Description: "Optional description of this inventory.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"has_active_failures": {
					Description: "This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"has_inventory_sources": {
					Description: "This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"host_filter": {
					Description: "Filter that will be applied to the hosts of this inventory.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"hosts_with_active_failures": {
					Description: "This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this inventory.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
						),
					},
				},
				"inventory_sources_with_failures": {
					Description: "Number of external inventory sources in this inventory with failures.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"kind": {
					Description: "Kind of inventory being represented.",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "smart"}...),
					},
				},
				"name": {
					Description: "Name of this inventory.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"organization": {
					Description: "Organization containing this inventory.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"pending_deletion": {
					Description: "Flag indicating the inventory is being deleted.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"prevent_instance_group_fallback": {
					Description: "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"total_groups": {
					Description: "This field is deprecated and will be removed in a future release. Total number of groups in this inventory.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"total_hosts": {
					Description: "This field is deprecated and will be removed in a future release. Total number of hosts in this inventory.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"total_inventory_sources": {
					Description: "Total number of external inventory sources configured within this inventory.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"variables": {
					Description: "Inventory variables in JSON or YAML format.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *inventoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state inventoryTerraformModel
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

	// Creates a new request for Inventory
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Inventory on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Inventory
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Inventory on %s", o.endpoint),
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
	_ resource.Resource                = &inventoryResource{}
	_ resource.ResourceWithConfigure   = &inventoryResource{}
	_ resource.ResourceWithImportState = &inventoryResource{}
)

// NewInventoryResource is a helper function to simplify the provider implementation.
func NewInventoryResource() resource.Resource {
	return &inventoryResource{}
}

type inventoryResource struct {
	client   c.Client
	endpoint string
}

func (o *inventoryResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventories/"
}

func (o inventoryResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_inventory"
}

func (o inventoryResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Inventory",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"description": {
					Description: "Optional description of this inventory.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"host_filter": {
					Description: "Filter that will be applied to the hosts of this inventory.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"kind": {
					Description: "Kind of inventory being represented.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "smart"}...),
					},
				},
				"name": {
					Description:   "Name of this inventory.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"organization": {
					Description:   "Organization containing this inventory.",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"prevent_instance_group_fallback": {
					Description: "If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"variables": {
					Description: "Inventory variables in JSON or YAML format.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"has_active_failures": {
					Description: "This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed.",
					Computed:    true,
					Type:        types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"has_inventory_sources": {
					Description: "This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources.",
					Computed:    true,
					Type:        types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"hosts_with_active_failures": {
					Description: "This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this inventory.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"inventory_sources_with_failures": {
					Description: "Number of external inventory sources in this inventory with failures.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"pending_deletion": {
					Description: "Flag indicating the inventory is being deleted.",
					Computed:    true,
					Type:        types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"total_groups": {
					Description: "This field is deprecated and will be removed in a future release. Total number of groups in this inventory.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"total_hosts": {
					Description: "This field is deprecated and will be removed in a future release. Total number of hosts in this inventory.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"total_inventory_sources": {
					Description: "Total number of external inventory sources configured within this inventory.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *inventoryResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Inventory.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *inventoryResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state inventoryTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Inventory
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Inventory on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Inventory resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Inventory on %s", o.endpoint),
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

func (o *inventoryResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state inventoryTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Inventory
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Inventory on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Inventory from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Inventory on %s", o.endpoint),
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

func (o *inventoryResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state inventoryTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Inventory
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Inventory on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Inventory resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Inventory on %s", o.endpoint),
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

func (o *inventoryResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state inventoryTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Inventory
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Inventory on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Inventory
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Inventory on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

var (
	_ datasource.DataSource              = &inventoryObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &inventoryObjectRolesDataSource{}
)

// NewInventoryObjectRolesDataSource is a helper function to instantiate the Inventory data source.
func NewInventoryObjectRolesDataSource() datasource.DataSource {
	return &inventoryObjectRolesDataSource{}
}

// inventoryObjectRolesDataSource is the data source implementation.
type inventoryObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *inventoryObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventories/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *inventoryObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_object_roles"
}

// GetSchema defines the schema for the data source.
func (o *inventoryObjectRolesDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: helpers.SchemaVersion,
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "Inventory ID",
				Type:        types.Int64Type,
				Required:    true,
			},
			"roles": {
				Description: "Roles for inventory",
				Type:        types.MapType{ElemType: types.Int64Type},
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (o *inventoryObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state inventoryObjectRolesModel
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
			fmt.Sprintf("Unable to create a new request for inventory"),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to fetch the request for inventory object roles "),
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for inventory",
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
