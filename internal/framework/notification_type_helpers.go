package framework

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// What AWX stores when a template is created without a `messages` body. Pinning
// it as the schema default keeps an unset attribute from planning as drift.
const DefaultNotificationMessages = `{"error":null,"started":null,"success":null,"workflow_approval":null}`

// Returns a fresh map per call so callers can layer the per-type attrs on top.
func NotificationBaseResourceAttrs() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Description: "Database ID of this notification template.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Description: "Name of this notification template.",
			Required:    true,
			Validators: []validator.String{
				stringvalidator.LengthAtMost(512),
			},
		},
		"description": schema.StringAttribute{
			Description: "Optional description of this notification template.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"organization": schema.Int64Attribute{
			Description: "Organization this notification template belongs to.",
			Required:    true,
		},
		// Computed, not Required: the resource type pins this, so a
		// user-supplied value could only disagree with it.
		"notification_type": schema.StringAttribute{
			Description: "AWX notification type, fixed by this resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"messages": schema.StringAttribute{
			Description: "Optional custom messages for this notification template, as a JSON object.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(DefaultNotificationMessages),
		},
	}
}

func NotificationBaseDataSourceAttrs() map[string]dschema.Attribute {
	idOrName := []path.Expression{path.MatchRoot("id"), path.MatchRoot("name")}
	return map[string]dschema.Attribute{
		"id": dschema.Int64Attribute{
			Description: "Database ID of this notification template.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{
				int64validator.ExactlyOneOf(idOrName...),
			},
		},
		"name": dschema.StringAttribute{
			Description: "Name of this notification template.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(idOrName...),
			},
		},
		"description":       dschema.StringAttribute{Description: "Optional description of this notification template.", Computed: true},
		"organization":      dschema.Int64Attribute{Description: "Organization this notification template belongs to.", Computed: true},
		"notification_type": dschema.StringAttribute{Description: "AWX notification type, fixed by this resource.", Computed: true},
		"messages":          dschema.StringAttribute{Description: "Optional custom messages for this notification template, as a JSON object.", Computed: true},
	}
}
