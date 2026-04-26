package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// loadCredentialTypePayload reads a credential_type spec straight from the
// 24.6.1 payload directory so the test reflects what the generator sees in
// production. Path is computed relative to the test file location.
func loadCredentialTypePayload(t *testing.T, namespace string) map[string]any {
	t.Helper()
	wd, err := os.Getwd()
	require.NoError(t, err)
	repoRoot := filepath.Clean(filepath.Join(wd, "..", "..", ".."))
	path := filepath.Join(repoRoot, "resources", "api", "24.6.1", "payload", "credential_type_"+namespace+".json")
	raw, err := os.ReadFile(path)
	require.NoError(t, err, "load %s", path)
	var payload map[string]any
	require.NoError(t, json.Unmarshal(raw, &payload))
	return payload
}

func TestBuildCredentialTypeTplData_AWS(t *testing.T) {
	cfg := Config{ApiVersion: "24.6.1"}
	item := Item{
		Name:           "CredentialAws",
		TypeName:       "credential_aws",
		Endpoint:       "/api/v2/credentials/",
		CredentialType: "aws",
		Enabled:        true,
	}

	data, err := buildCredentialTypeTplData(cfg, item, loadCredentialTypePayload(t, "aws"))
	require.NoError(t, err)

	assert.Equal(t, "Amazon Web Services", data.DisplayName)
	assert.Equal(t, "cloud", data.Kind)
	assert.Equal(t, "aws", data.Namespace)
	require.Len(t, data.Fields, 3)

	// Sorted by ID — password, security_token, username.
	assert.Equal(t, []string{"password", "security_token", "username"}, []string{data.Fields[0].ID, data.Fields[1].ID, data.Fields[2].ID})

	password := data.Fields[0]
	assert.Equal(t, "Password", password.PropertyName)
	assert.Equal(t, "Secret Key", password.Label)
	assert.True(t, password.Secret)
	assert.True(t, password.Required, "password is in inputs.required")

	username := data.Fields[2]
	assert.Equal(t, "Username", username.PropertyName)
	assert.False(t, username.Secret)
	assert.True(t, username.Required, "username is in inputs.required")

	securityToken := data.Fields[1]
	assert.True(t, securityToken.Secret, "security_token is marked secret in spec")
	assert.False(t, securityToken.Required, "security_token is optional")

	assert.True(t, data.HasSecrets)
}

func TestBuildCredentialTypeTplData_SSH(t *testing.T) {
	// SSH is the most exotic spec — it has multiline + format + ask_at_runtime
	// fields. Confirms the mapper doesn't choke on the variety.
	cfg := Config{ApiVersion: "24.6.1"}
	item := Item{
		Name:           "CredentialSsh",
		TypeName:       "credential_ssh",
		Endpoint:       "/api/v2/credentials/",
		CredentialType: "ssh",
		Enabled:        true,
	}

	data, err := buildCredentialTypeTplData(cfg, item, loadCredentialTypePayload(t, "ssh"))
	require.NoError(t, err)

	assert.Equal(t, "ssh", data.Kind)
	assert.Equal(t, "Machine", data.DisplayName)
	assert.Len(t, data.Fields, 8, "ssh spec has 8 input fields")

	byID := map[string]CredentialTypeField{}
	for _, f := range data.Fields {
		byID[f.ID] = f
	}

	sshKey := byID["ssh_key_data"]
	assert.True(t, sshKey.Secret)
	assert.True(t, sshKey.Multiline)
	assert.Equal(t, "ssh_private_key", sshKey.Format)

	password := byID["password"]
	assert.True(t, password.Secret, "password is sensitive even when ask_at_runtime")

	becomeMethod := byID["become_method"]
	assert.False(t, becomeMethod.Secret)
	assert.False(t, becomeMethod.Multiline)
}

func TestIsCredentialTypeItem(t *testing.T) {
	assert.True(t, IsCredentialTypeItem(Item{CredentialType: "aws"}))
	assert.False(t, IsCredentialTypeItem(Item{CredentialType: ""}))
	assert.False(t, IsCredentialTypeItem(Item{CredentialType: "  "}))
}

func TestCredentialTypeNamespaces(t *testing.T) {
	items := []Item{
		{CredentialType: "aws", Enabled: true},
		{CredentialType: "ssh", Enabled: false}, // disabled — excluded
		{CredentialType: "vault", Enabled: true},
		{CredentialType: "", Enabled: true}, // not a credential type
	}
	assert.Equal(t, []string{"aws", "vault"}, CredentialTypeNamespaces(items))
}

func TestBuildCredentialTypeTplData_RequiresInputs(t *testing.T) {
	cfg := Config{ApiVersion: "24.6.1"}
	item := Item{Name: "Bad", TypeName: "bad", CredentialType: "broken", Enabled: true}
	_, err := buildCredentialTypeTplData(cfg, item, map[string]any{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no inputs object")
}
