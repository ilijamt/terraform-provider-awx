package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Reads the real payload so the test sees what the generator sees.
func loadNotificationTemplatePayload(t *testing.T) map[string]any {
	t.Helper()
	wd, err := os.Getwd()
	require.NoError(t, err)
	repoRoot := filepath.Clean(filepath.Join(wd, "..", "..", ".."))
	path := filepath.Join(repoRoot, "resources", "api", "24.6.1", "payload", "resource_notificationtemplate.json")
	raw, err := os.ReadFile(path)
	require.NoError(t, err, "load %s", path)
	var payload map[string]any
	require.NoError(t, json.Unmarshal(raw, &payload))
	return payload
}

func notificationItem(notificationType, name, typeName string) Item {
	return Item{
		Name:             name,
		TypeName:         typeName,
		Endpoint:         "/api/v2/notification_templates/",
		NotificationType: notificationType,
		Enabled:          true,
	}
}

func fieldsByID(fields []NotificationTypeField) map[string]NotificationTypeField {
	out := map[string]NotificationTypeField{}
	for _, f := range fields {
		out[f.ID] = f
	}
	return out
}

func TestBuildNotificationTypeTplData_Webhook(t *testing.T) {
	cfg := Config{ApiVersion: "24.6.1"}
	item := notificationItem("webhook", "NotificationTemplateWebhook", "notification_template_webhook")

	data, err := buildNotificationTypeTplData(cfg, item, loadNotificationTemplatePayload(t))
	require.NoError(t, err)

	assert.Equal(t, "Webhook", data.DisplayName)
	assert.Equal(t, "webhook", data.Namespace)
	assert.True(t, data.HasSecrets)

	assert.Equal(t,
		[]string{"disable_ssl_verification", "headers", "http_method", "password", "url", "username"},
		[]string{data.Fields[0].ID, data.Fields[1].ID, data.Fields[2].ID, data.Fields[3].ID, data.Fields[4].ID, data.Fields[5].ID})

	byID := fieldsByID(data.Fields)

	// AWX answers an empty webhook config with "Missing required fields for
	// Notification Configuration: ['url', 'headers']".
	assert.True(t, byID["url"].Required)
	assert.True(t, byID["headers"].Required)
	for _, id := range []string{"http_method", "disable_ssl_verification", "username", "password"} {
		assert.False(t, byID[id].Required, "%s carries a default, so AWX fills it in", id)
	}

	assert.True(t, byID["password"].Secret)
	assert.False(t, byID["username"].Secret)

	assert.True(t, byID["headers"].IsObject())
	assert.True(t, byID["disable_ssl_verification"].IsBool())
	assert.Equal(t, "DisableSslVerification", byID["disable_ssl_verification"].PropertyName)

	// A false default is still a default.
	assert.True(t, byID["disable_ssl_verification"].HasDefault)
	assert.Equal(t, false, byID["disable_ssl_verification"].Default)
	assert.Equal(t, "POST", byID["http_method"].Default)
	assert.Equal(t, "HTTP Method. AWX defaults this to \"POST\" when unset.", byID["http_method"].Description)
}

func TestBuildNotificationTypeTplData_Email(t *testing.T) {
	// The widest spec: int, list, bool, password and a defaulted int in one type.
	cfg := Config{ApiVersion: "24.6.1"}
	item := notificationItem("email", "NotificationTemplateEmail", "notification_template_email")

	data, err := buildNotificationTypeTplData(cfg, item, loadNotificationTemplatePayload(t))
	require.NoError(t, err)

	assert.Equal(t, "Email", data.DisplayName)
	require.Len(t, data.Fields, 9)

	byID := fieldsByID(data.Fields)
	assert.True(t, byID["port"].IsInt())
	assert.True(t, byID["recipients"].IsList())
	assert.True(t, byID["use_tls"].IsBool())
	assert.True(t, byID["password"].Secret)

	assert.False(t, byID["timeout"].Required)
	assert.True(t, byID["timeout"].HasDefault)
	assert.Equal(t, "Timeout. AWX defaults this to 30 when unset.", byID["timeout"].Description)
	for _, id := range []string{"host", "port", "username", "password", "use_tls", "use_ssl", "sender", "recipients"} {
		assert.True(t, byID[id].Required, "%s has no default in the spec", id)
	}

	assert.Len(t, data.DataSourceFields(), 8)
	for _, f := range data.DataSourceFields() {
		assert.NotEqual(t, "password", f.ID)
	}
}

func TestBuildNotificationTypeTplData_AwssnsHasNoRequiredFields(t *testing.T) {
	// This is why AWX accepts a create with an empty notification_configuration.
	cfg := Config{ApiVersion: "24.6.1"}
	item := notificationItem("awssns", "NotificationTemplateAwssns", "notification_template_awssns")

	data, err := buildNotificationTypeTplData(cfg, item, loadNotificationTemplatePayload(t))
	require.NoError(t, err)

	assert.Equal(t, "AWS SNS", data.DisplayName)
	require.Len(t, data.Fields, 5)
	for _, f := range data.Fields {
		assert.False(t, f.Required, "%s carries a default", f.ID)
		assert.True(t, f.HasDefault)
	}
	assert.True(t, data.HasSecrets, "aws_secret_access_key and aws_session_token are passwords")
}

func TestBuildNotificationTypeTplData_MattermostHasNoSecrets(t *testing.T) {
	// A type with no secrets gets no hook in the template.
	cfg := Config{ApiVersion: "24.6.1"}
	item := notificationItem("mattermost", "NotificationTemplateMattermost", "notification_template_mattermost")

	data, err := buildNotificationTypeTplData(cfg, item, loadNotificationTemplatePayload(t))
	require.NoError(t, err)

	assert.False(t, data.HasSecrets)
	assert.Len(t, data.DataSourceFields(), len(data.Fields))
}

func TestBuildNotificationTypeTplData_Errors(t *testing.T) {
	cfg := Config{ApiVersion: "24.6.1"}
	payload := loadNotificationTemplatePayload(t)

	t.Run("unknown notification type", func(t *testing.T) {
		_, err := buildNotificationTypeTplData(cfg, notificationItem("nope", "X", "x"), payload)
		require.ErrorContains(t, err, "absent from actions.POST.notification_configuration")
	})

	t.Run("field description key is not a notification type", func(t *testing.T) {
		// These sit alongside the per-type schemas in the same object.
		_, err := buildNotificationTypeTplData(cfg, notificationItem("label", "X", "x"), payload)
		require.ErrorContains(t, err, "field-description key")
	})

	t.Run("payload without actions", func(t *testing.T) {
		_, err := buildNotificationTypeTplData(cfg, notificationItem("webhook", "X", "x"), map[string]any{})
		require.ErrorContains(t, err, "no actions object")
	})
}

func TestGenerateNotificationTypeTfDefinition_Guards(t *testing.T) {
	t.Run("empty notification type", func(t *testing.T) {
		err := GenerateNotificationTypeTfDefinition(nil, Config{}, Item{Name: "X"}, t.TempDir(), map[string]any{})
		require.ErrorContains(t, err, "notification_type is empty")
	})

	t.Run("nil payload", func(t *testing.T) {
		err := GenerateNotificationTypeTfDefinition(nil, Config{}, notificationItem("webhook", "X", "x"), t.TempDir(), nil)
		require.ErrorContains(t, err, "no NotificationTemplate payload found")
	})
}

func TestIsNotificationTypeItemAndNamespaces(t *testing.T) {
	items := []Item{
		notificationItem("webhook", "NotificationTemplateWebhook", "notification_template_webhook"),
		notificationItem("email", "NotificationTemplateEmail", "notification_template_email"),
		{Name: "NotificationTemplate", TypeName: "notification_template", Enabled: true},
		{Name: "CredentialAws", CredentialType: "aws", Enabled: true},
	}
	disabled := notificationItem("slack", "NotificationTemplateSlack", "notification_template_slack")
	disabled.Enabled = false
	items = append(items, disabled)

	assert.True(t, IsNotificationTypeItem(items[0]))
	assert.False(t, IsNotificationTypeItem(items[2]), "the untyped notification_template item stays on the standard path")
	assert.False(t, IsNotificationTypeItem(items[3]))
	assert.Equal(t, []string{"email", "webhook"}, NotificationTypeNamespaces(items))
}
