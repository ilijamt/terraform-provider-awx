package framework_test

import (
	"context"
	"testing"

	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestCredentialTypeLookup(t *testing.T) {
	ctx := context.Background()

	t.Run("OnConfigure populates Load()", func(t *testing.T) {
		lookup := framework.NewCredentialTypeLookup()
		assert.Equal(t, int64(0), lookup.Load())

		req := namespaceLookupRequester(t, "vault", []map[string]any{{"id": float64(13), "namespace": "vault"}})
		diags := lookup.OnConfigure("vault")(ctx, req)

		require.False(t, diags.HasError())
		assert.Equal(t, int64(13), lookup.Load())
	})

	t.Run("error path leaves Load() at zero", func(t *testing.T) {
		lookup := framework.NewCredentialTypeLookup()
		req := namespaceLookupRequester(t, "ssh", []map[string]any{})
		diags := lookup.OnConfigure("ssh")(ctx, req)

		require.True(t, diags.HasError())
		assert.Equal(t, int64(0), lookup.Load())
	})
}

func TestCredentialBaseResourceAttrs(t *testing.T) {
	attrs := framework.CredentialBaseResourceAttrs()

	wantKeys := []string{"id", "name", "description", "organization", "team", "user", "credential_type", "kind", "managed"}
	for _, k := range wantKeys {
		_, ok := attrs[k]
		assert.Truef(t, ok, "missing attr %q", k)
	}
	assert.Len(t, attrs, len(wantKeys))

	// Spot-check the typing — the contract callers depend on.
	assert.IsType(t, rschema.Int64Attribute{}, attrs["id"])
	assert.IsType(t, rschema.StringAttribute{}, attrs["name"])
	assert.IsType(t, rschema.StringAttribute{}, attrs["description"])
	assert.IsType(t, rschema.Int64Attribute{}, attrs["organization"])
	assert.IsType(t, rschema.Int64Attribute{}, attrs["credential_type"])
	assert.IsType(t, rschema.BoolAttribute{}, attrs["managed"])

	// Each call returns a fresh map: mutating one must not leak into the next.
	attrs["mine"] = rschema.StringAttribute{}
	fresh := framework.CredentialBaseResourceAttrs()
	_, leaked := fresh["mine"]
	assert.False(t, leaked, "helper must return a fresh map per call")

	// name is the only required attr (caller adds the per-credential-type fields on top).
	name := attrs["name"].(rschema.StringAttribute)
	assert.True(t, name.Required)
}

func TestCredentialBaseDataSourceAttrs(t *testing.T) {
	attrs := framework.CredentialBaseDataSourceAttrs()

	wantKeys := []string{"id", "name", "description", "organization", "team", "user", "credential_type", "kind", "managed"}
	for _, k := range wantKeys {
		_, ok := attrs[k]
		assert.Truef(t, ok, "missing attr %q", k)
	}
	assert.Len(t, attrs, len(wantKeys))

	// id and name are Optional+Computed (lookup keys); the rest are Computed-only.
	id := attrs["id"].(dschema.Int64Attribute)
	assert.True(t, id.Optional)
	assert.True(t, id.Computed)

	desc := attrs["description"].(dschema.StringAttribute)
	assert.False(t, desc.Optional)
	assert.True(t, desc.Computed)
}
