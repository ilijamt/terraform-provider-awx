package internal

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// Every typed notification resource is generated from this one item's OPTIONS
// payload, so the name is fixed here rather than repeated in each config file.
const NotificationTemplateResourceName = "NotificationTemplate"

// DRF field-description keys that share actions.POST.notification_configuration
// with the per-type schemas. They are not notification types.
var notificationConfigurationMetaKeys = []string{"default", "filterable", "hidden", "label", "required", "type"}

type NotificationTypeField struct {
	ID           string
	PropertyName string
	Label        string
	Description  string
	Type         string // string / password / int / bool / list / object
	Secret       bool
	Required     bool
	// Checked instead of Default's truthiness: a `false` bool and an
	// empty-string default are real defaults that still have to be emitted.
	HasDefault bool
	Default    any
}

func (f NotificationTypeField) IsBool() bool   { return f.Type == notificationFieldTypeBool }
func (f NotificationTypeField) IsInt() bool    { return f.Type == notificationFieldTypeInt }
func (f NotificationTypeField) IsList() bool   { return f.Type == notificationFieldTypeList }
func (f NotificationTypeField) IsObject() bool { return f.Type == notificationFieldTypeObject }

const (
	notificationFieldTypeBool     = "bool"
	notificationFieldTypeInt      = "int"
	notificationFieldTypeList     = "list"
	notificationFieldTypeObject   = "object"
	notificationFieldTypePassword = "password"
)

type NotificationTypeTplData struct {
	ApiVersion  string
	PackageName string
	Name        string // NotificationTemplateWebhook
	TypeName    string // notification_template_webhook
	Endpoint    string
	Namespace   string // webhook
	DisplayName string // Webhook
	Fields      []NotificationTypeField
	HasSecrets  bool
	Enabled     bool
}

// DataSourceFields drops the secret fields. AWX answers those with the literal
// "$encrypted$", so a read-only attribute for one returns nothing usable.
func (d NotificationTypeTplData) DataSourceFields() []NotificationTypeField {
	out := make([]NotificationTypeField, 0, len(d.Fields))
	for _, f := range d.Fields {
		if !f.Secret {
			out = append(out, f)
		}
	}
	return out
}

func GenerateNotificationTypeTfDefinition(tpl *template.Template, config Config, item Item, resourcePath string, payload map[string]any) error {
	if item.NotificationType == "" {
		return fmt.Errorf("notification_type is empty for item %q", item.Name)
	}
	if payload == nil {
		return fmt.Errorf("no %s payload found for notification type %q", NotificationTemplateResourceName, item.NotificationType)
	}

	log.Printf("Generating typed notification template resource for %s (notification_type=%s)", item.Name, item.NotificationType)

	data, err := buildNotificationTypeTplData(config, item, payload)
	if err != nil {
		return err
	}

	if !item.Enabled {
		log.Printf("Skipping %s, disabled ...", item.Name)
		return nil
	}

	filename := fmt.Sprintf("%s/gen_obj_%s.go", resourcePath, strings.ToLower(item.TypeName))
	return renderTemplate(tpl, filename, "tf_notification_type.go.tpl", data)
}

func buildNotificationTypeTplData(config Config, item Item, payload map[string]any) (*NotificationTypeTplData, error) {
	out := &NotificationTypeTplData{
		ApiVersion:  config.ApiVersion,
		PackageName: config.PackageName("awx"),
		Name:        item.Name,
		TypeName:    item.TypeName,
		Endpoint:    item.Endpoint,
		Namespace:   item.NotificationType,
		DisplayName: notificationTypeDisplayName(payload, item.NotificationType),
		Enabled:     item.Enabled,
	}

	schema, err := notificationConfigurationSchema(payload, item.NotificationType)
	if err != nil {
		return nil, err
	}

	fields := make([]NotificationTypeField, 0, len(schema))
	for id, raw := range schema {
		fm, ok := raw.(map[string]any)
		if !ok {
			continue
		}

		ft := NotificationTypeField{
			ID:           id,
			PropertyName: strcase.ToCamel(id),
			Type:         "string",
		}
		if v, ok := fm["type"].(string); ok && v != "" {
			ft.Type = v
		}
		if v, ok := fm["label"].(string); ok {
			ft.Label = v
		}
		ft.Secret = ft.Type == notificationFieldTypePassword
		// AWX fills a missing field from its schema default and rejects the
		// request when there is none, so presence of the key is exactly the
		// required/optional split (NotificationTemplateSerializer.validate).
		ft.Default, ft.HasDefault = fm["default"]
		ft.Required = !ft.HasDefault
		ft.Description = notificationFieldDescription(ft)
		if ft.Secret {
			out.HasSecrets = true
		}
		fields = append(fields, ft)
	}

	// Keeps generated diffs stable if AWX reshuffles its schema object.
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].ID < fields[j].ID })
	out.Fields = fields

	return out, nil
}

func notificationConfigurationSchema(payload map[string]any, notificationType string) (map[string]any, error) {
	actions, ok := payload["actions"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s payload has no actions object", NotificationTemplateResourceName)
	}
	post, ok := actions["POST"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s payload has no actions.POST object", NotificationTemplateResourceName)
	}
	config, ok := post["notification_configuration"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s payload has no actions.POST.notification_configuration object", NotificationTemplateResourceName)
	}
	if slices.Contains(notificationConfigurationMetaKeys, notificationType) {
		return nil, fmt.Errorf("notification_type %q is a field-description key, not a notification type", notificationType)
	}
	schema, ok := config[notificationType].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("notification_type %q is absent from actions.POST.notification_configuration", notificationType)
	}
	return schema, nil
}

func notificationTypeDisplayName(payload map[string]any, notificationType string) string {
	actions, _ := payload["actions"].(map[string]any)
	post, _ := actions["POST"].(map[string]any)
	field, _ := post["notification_type"].(map[string]any)
	choices, _ := field["choices"].([]any)
	for _, choice := range choices {
		pair, ok := choice.([]any)
		if !ok || len(pair) != 2 {
			continue
		}
		if value, ok := pair[0].(string); ok && value == notificationType {
			if label, ok := pair[1].(string); ok && label != "" {
				return label
			}
		}
	}
	return notificationType
}

// Records the AWX default rather than applying it: a defaulted field is
// Optional+Computed, so AWX fills it in.
func notificationFieldDescription(f NotificationTypeField) string {
	lead := f.Label
	if lead != "" && !strings.HasSuffix(lead, ".") {
		lead += "."
	}

	parts := []string{lead}
	// Optional in the schema even when AWX requires the key, because
	// BodyRequest always sends it. Without this the docs contradict the API.
	if f.IsObject() && f.Required {
		parts = append(parts, "AWX requires this key; the provider sends an empty map when unset.")
	}
	switch v := f.Default.(type) {
	case string:
		if v != "" {
			parts = append(parts, fmt.Sprintf("AWX defaults this to %q when unset.", v))
		}
	case bool:
		parts = append(parts, fmt.Sprintf("AWX defaults this to %q when unset.", strconv.FormatBool(v)))
	case float64:
		parts = append(parts, fmt.Sprintf("AWX defaults this to %s when unset.", strconv.FormatFloat(v, 'f', -1, 64)))
	}
	return strings.TrimSpace(strings.Join(parts, " "))
}

func IsNotificationTypeItem(item Item) bool {
	return strings.TrimSpace(item.NotificationType) != ""
}

func NotificationTypeNamespaces(items []Item) []string {
	out := make([]string, 0)
	for _, it := range items {
		if IsNotificationTypeItem(it) && it.Enabled {
			out = append(out, it.NotificationType)
		}
	}
	slices.Sort(out)
	return out
}
