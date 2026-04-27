package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
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

type notificationTemplateDataSource = framework.GenericDataSource[notificationTemplateTerraformModel, *notificationTemplateTerraformModel]

// NewNotificationTemplateDataSource is a helper function to instantiate the NotificationTemplate data source.
func NewNotificationTemplateDataSource() datasource.DataSource {
	return &notificationTemplateDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this notification template.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this notification template.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"messages": dschema.StringAttribute{
						Description: "Optional custom messages for notification template.",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this notification template.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"notification_configuration": dschema.StringAttribute{
						Description: "Notification configuration",
						Computed:    true,
					},
					"notification_type": dschema.StringAttribute{
						Description: "Notification type",
						Computed:    true,
					},
					"organization": dschema.Int64Attribute{
						Description: "Organization",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook:         hookNotificationTemplate,
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplate",
		},
	}
}
