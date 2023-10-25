package awx_21_8_0

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
