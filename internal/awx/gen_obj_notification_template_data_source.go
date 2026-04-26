package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type notificationTemplateDataSource = framework.GenericDataSource[notificationTemplateTerraformModel, *notificationTemplateTerraformModel]

// NewNotificationTemplateDataSource is a helper function to instantiate the NotificationTemplate data source.
func NewNotificationTemplateDataSource() datasource.DataSource {
	return &notificationTemplateDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"description": schema.StringAttribute{
						Description: "Optional description of this notification template.",
						Sensitive:   false,
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this notification template.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"messages": schema.StringAttribute{
						Description: "Optional custom messages for notification template.",
						Sensitive:   false,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this notification template.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"notification_configuration": schema.StringAttribute{
						Description: "Notification configuration",
						Sensitive:   false,
						Computed:    true,
					},
					"notification_type": schema.StringAttribute{
						Description: "Notification type",
						Sensitive:   false,
						Computed:    true,
					},
					"organization": schema.Int64Attribute{
						Description: "Organization",
						Sensitive:   false,
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
