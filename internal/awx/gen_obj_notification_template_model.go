package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for NotificationTemplate
func (o *notificationTemplateTerraformModel) BodyRequest() *notificationTemplateBodyRequestModel {
	var req notificationTemplateBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Messages = json.RawMessage(o.Messages.ValueString())
	req.Name = o.Name.ValueString()
	req.NotificationConfiguration = json.RawMessage(o.NotificationConfiguration.ValueString())
	req.NotificationType = o.NotificationType.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *notificationTemplateTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.Messages, data["messages"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.NotificationConfiguration, data["notification_configuration"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.NotificationType, data["notification_type"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
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
