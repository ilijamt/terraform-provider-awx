package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &inventorySourceResource{}
	_ resource.ResourceWithConfigure   = &inventorySourceResource{}
	_ resource.ResourceWithImportState = &inventorySourceResource{}
)

// NewInventorySourceResource is a helper function to simplify the provider implementation.
func NewInventorySourceResource() resource.Resource {
	return &inventorySourceResource{}
}

type inventorySourceResource struct {
	client   c.Client
	endpoint string
}

func (o *inventorySourceResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventory_sources/"
}

func (o *inventorySourceResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_inventory_source"
}

func (o *inventorySourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Request elements
			"credential": schema.Int64Attribute{
				Description: "Cloud credential to use for inventory updates.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this inventory source.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"enabled_value": schema.StringAttribute{
				Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"enabled_var": schema.StringAttribute{
				Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"execution_environment": schema.Int64Attribute{
				Description: "The container image to be used for execution.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"host_filter": schema.StringAttribute{
				Description: "Regex where only matching hosts will be imported.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"inventory": schema.Int64Attribute{
				Description:   "Inventory",
				Sensitive:     false,
				Required:      true,
				Optional:      false,
				Computed:      false,
				PlanModifiers: []planmodifier.Int64{},
				Validators:    []validator.Int64{},
			},
			"name": schema.StringAttribute{
				Description:   "Name of this inventory source.",
				Sensitive:     false,
				Required:      true,
				Optional:      false,
				Computed:      false,
				PlanModifiers: []planmodifier.String{},
				Validators: []validator.String{
					stringvalidator.LengthAtMost(512),
				},
			},
			"overwrite": schema.BoolAttribute{
				Description: "Overwrite local groups and hosts from remote inventory source.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"overwrite_vars": schema.BoolAttribute{
				Description: "Overwrite local variables from remote inventory source.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"source": schema.StringAttribute{
				Description: "Source",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"file", "scm", "ec2", "gce", "azure_rm", "vmware", "satellite6", "openstack", "rhv", "controller", "insights"}...),
				},
			},
			"source_path": schema.StringAttribute{
				Description: "Source path",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1024),
				},
			},
			"source_project": schema.Int64Attribute{
				Description: "Project containing inventory file used as source.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{},
			},
			"source_vars": schema.StringAttribute{
				Description: "Inventory source variables in YAML or JSON format.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"timeout": schema.Int64Attribute{
				Description: "The amount of time (in seconds) to run before the task is canceled.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			"update_cache_timeout": schema.Int64Attribute{
				Description: "Update cache timeout",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"update_on_launch": schema.BoolAttribute{
				Description: "Update on launch",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Bool{},
			},
			"verbosity": schema.StringAttribute{
				Description: "Verbosity",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`%!s(float64=1)`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"0", "1", "2"}...),
				},
			},
			// Write only elements
			// Data only elements
			"id": schema.Int64Attribute{
				Description: "Database ID for this inventory source.",
				Required:    false,
				Optional:    false,
				Computed:    true,
				Sensitive:   false,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o *inventorySourceResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the InventorySource.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *inventorySourceResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state inventorySourceTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[InventorySource/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for create", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new InventorySource resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for InventorySource on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hooks.RequireResourceStateOrOrig(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on InventorySource",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *inventorySourceResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state inventorySourceTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for InventorySource
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for read", endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for InventorySource from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InventorySource on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hooks.RequireResourceStateOrOrig(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on InventorySource",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *inventorySourceResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state inventorySourceTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[InventorySource/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for update", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new InventorySource resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for InventorySource on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hooks.RequireResourceStateOrOrig(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on InventorySource",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *inventorySourceResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state inventorySourceTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing InventorySource
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for InventorySource on %s", endpoint),
			err.Error(),
		)
		return
	}
}
