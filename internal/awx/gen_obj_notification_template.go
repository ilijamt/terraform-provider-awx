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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// notificationTemplateTerraformModel maps the schema for NotificationTemplate when using Data Source
type notificationTemplateTerraformModel struct {
	// Description "Optional description of this notification template."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this notification template."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Messages "Optional custom messages for notification template."
	Messages types.String `tfsdk:"messages" json:"messages"`
	// Name "Name of this notification template."
	Name types.String `tfsdk:"name" json:"name"`
	// NotificationConfiguration ""
	NotificationConfiguration types.String `tfsdk:"notification_configuration" json:"notification_configuration"`
	// NotificationType ""
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	// Organization ""
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
}

// Clone the object
func (o *notificationTemplateTerraformModel) Clone() notificationTemplateTerraformModel {
	return notificationTemplateTerraformModel{
		Description:               o.Description,
		ID:                        o.ID,
		Messages:                  o.Messages,
		Name:                      o.Name,
		NotificationConfiguration: o.NotificationConfiguration,
		NotificationType:          o.NotificationType,
		Organization:              o.Organization,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for NotificationTemplate
func (o *notificationTemplateTerraformModel) BodyRequest() (req notificationTemplateBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Messages = json.RawMessage(o.Messages.ValueString())
	req.Name = o.Name.ValueString()
	req.NotificationConfiguration = json.RawMessage(o.NotificationConfiguration.ValueString())
	req.NotificationType = o.NotificationType.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return
}

func (o *notificationTemplateTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *notificationTemplateTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *notificationTemplateTerraformModel) setMessages(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.Messages, data, false)
}

func (o *notificationTemplateTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *notificationTemplateTerraformModel) setNotificationConfiguration(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.NotificationConfiguration, data, false)
}

func (o *notificationTemplateTerraformModel) setNotificationType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.NotificationType, data, false)
}

func (o *notificationTemplateTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *notificationTemplateTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMessages(data["messages"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setNotificationConfiguration(data["notification_configuration"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setNotificationType(data["notification_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// notificationTemplateBodyRequestModel maps the schema for NotificationTemplate for creating and updating the data
type notificationTemplateBodyRequestModel struct {
	// Description "Optional description of this notification template."
	Description string `json:"description,omitempty"`
	// Messages "Optional custom messages for notification template."
	Messages json.RawMessage `json:"messages,omitempty"`
	// Name "Name of this notification template."
	Name string `json:"name"`
	// NotificationConfiguration ""
	NotificationConfiguration json.RawMessage `json:"notification_configuration,omitempty"`
	// NotificationType ""
	NotificationType string `json:"notification_type"`
	// Organization ""
	Organization int64 `json:"organization"`
}

var (
	_ datasource.DataSource              = &notificationTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &notificationTemplateDataSource{}
)

// NewNotificationTemplateDataSource is a helper function to instantiate the NotificationTemplate data source.
func NewNotificationTemplateDataSource() datasource.DataSource {
	return &notificationTemplateDataSource{}
}

// notificationTemplateDataSource is the data source implementation.
type notificationTemplateDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *notificationTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/notification_templates/"
}

// Metadata returns the data source type name.
func (o *notificationTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_notification_template"
}

// GetSchema defines the schema for the data source.
func (o *notificationTemplateDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"NotificationTemplate",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"description": {
					Description: "Optional description of this notification template.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this notification template.",
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
				"messages": {
					Description: "Optional custom messages for notification template.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this notification template.",
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
				"notification_configuration": {
					Description: "Notification configuration",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"notification_type": {
					Description: "Notification type",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"email", "grafana", "irc", "mattermost", "pagerduty", "rocketchat", "slack", "twilio", "webhook"}...),
					},
				},
				"organization": {
					Description: "Organization",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *notificationTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state notificationTemplateTerraformModel
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

	// Creates a new request for NotificationTemplate
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for NotificationTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for NotificationTemplate
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for NotificationTemplate on %s", o.endpoint),
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
	_ resource.Resource                = &notificationTemplateResource{}
	_ resource.ResourceWithConfigure   = &notificationTemplateResource{}
	_ resource.ResourceWithImportState = &notificationTemplateResource{}
)

// NewNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewNotificationTemplateResource() resource.Resource {
	return &notificationTemplateResource{}
}

type notificationTemplateResource struct {
	client   c.Client
	endpoint string
}

func (o *notificationTemplateResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/notification_templates/"
}

func (o *notificationTemplateResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_notification_template"
}

func (o *notificationTemplateResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"NotificationTemplate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"description": {
					Description: "Optional description of this notification template.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"messages": {
					Description: "Optional custom messages for notification template.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{"error":null,"started":null,"success":null,"workflow_approval":null}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this notification template.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"notification_configuration": {
					Description: "Notification configuration",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"notification_type": {
					Description:   "Notification type",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"email", "grafana", "irc", "mattermost", "pagerduty", "rocketchat", "slack", "twilio", "webhook"}...),
					},
				},
				"organization": {
					Description:   "Organization",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"id": {
					Description: "Database ID for this notification template.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *notificationTemplateResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the NotificationTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *notificationTemplateResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state notificationTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for NotificationTemplate
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[NotificationTemplate/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for NotificationTemplate on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new NotificationTemplate resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for NotificationTemplate on %s", o.endpoint),
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

func (o *notificationTemplateResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state notificationTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for NotificationTemplate
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for NotificationTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for NotificationTemplate from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for NotificationTemplate on %s", o.endpoint),
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

func (o *notificationTemplateResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state notificationTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for NotificationTemplate
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[NotificationTemplate/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for NotificationTemplate on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new NotificationTemplate resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for NotificationTemplate on %s", o.endpoint),
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

func (o *notificationTemplateResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state notificationTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for NotificationTemplate
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for NotificationTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing NotificationTemplate
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for NotificationTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
