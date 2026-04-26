package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
					"description": schema.StringAttribute{
						Description: "Optional description of this notification template.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"messages": schema.StringAttribute{
						Description: "Optional custom messages for notification template.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"error":null,"started":null,"success":null,"workflow_approval":null}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this notification template.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"notification_configuration": schema.StringAttribute{
						Description: "Notification configuration",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"notification_type": schema.StringAttribute{
						Description: "Notification type",
						Required:    true,
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
						Description: "Organization",
						Required:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this notification template.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *notificationTemplateTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			Hook:         hookNotificationTemplate,
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplate",
		},
	}
}
