package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type notificationTemplateResource = framework.GenericResource[notificationTemplateTerraformModel, notificationTemplateBodyRequestModel, *notificationTemplateTerraformModel]

// NewNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewNotificationTemplateResource() resource.Resource {
	return &notificationTemplateResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateTerraformModel, notificationTemplateBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"description": schema.StringAttribute{
						Description: "Optional description of this notification template.",
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
					"messages": schema.StringAttribute{
						Description: "Optional custom messages for notification template.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"error":null,"started":null,"success":null,"workflow_approval":null}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"name": schema.StringAttribute{
						Description:   "Name of this notification template.",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"notification_configuration": schema.StringAttribute{
						Description: "Notification configuration",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"notification_type": schema.StringAttribute{
						Description:   "Notification type",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"awssns",
								"email",
								"grafana",
								"irc",
								"mattermost",
								"pagerduty",
								"rocketchat",
								"slack",
								"twilio",
								"webhook",
							),
						},
					},
					"organization": schema.Int64Attribute{
						Description:   "Organization",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.Int64{},
						Validators:    []validator.Int64{},
					},
					// Write only elements
					// Data only elements
					"id": schema.Int64Attribute{
						Description: "Database ID for this notification template.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:      func(m *notificationTemplateTerraformModel) any { return m.ID.ValueInt64() },
			ImportStateFunc: notificationTemplateResourceImportState,
			Hook:            hookNotificationTemplate,
			ApiVersion:      ApiVersion,
			ResourceName:    "NotificationTemplate",
		},
	}
}

func notificationTemplateResourceImportState(_ context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the NotificationTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(context.Background(), path.Root("id"), id)...)
}
