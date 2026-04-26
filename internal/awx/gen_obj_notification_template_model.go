package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type notificationTemplateTerraformModel struct {
	Description               types.String `tfsdk:"description" json:"description"`
	ID                        types.Int64  `tfsdk:"id" json:"id"`
	Messages                  types.String `tfsdk:"messages" json:"messages"`
	Name                      types.String `tfsdk:"name" json:"name"`
	NotificationConfiguration types.String `tfsdk:"notification_configuration" json:"notification_configuration"`
	NotificationType          types.String `tfsdk:"notification_type" json:"notification_type"`
	Organization              types.Int64  `tfsdk:"organization" json:"organization"`
}

func (o *notificationTemplateTerraformModel) Clone() notificationTemplateTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetJsonString(&o.Messages, data["messages"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetJsonString(&o.NotificationConfiguration, data["notification_configuration"], false))
	collect(helpers.AttrValueSetString(&o.NotificationType, data["notification_type"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type notificationTemplateBodyRequestModel struct {
	Description               string          `json:"description,omitempty"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
	Name                      string          `json:"name"`
	NotificationConfiguration json.RawMessage `json:"notification_configuration,omitempty"`
	NotificationType          string          `json:"notification_type"`
	Organization              int64           `json:"organization"`
}
