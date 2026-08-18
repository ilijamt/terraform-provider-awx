package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeJSON(t *testing.T) {
	for _, tt := range []struct {
		name string
		dst  string
		src  string
		want string
	}{
		{
			name: "appends to a list rather than replacing it",
			dst:  `{"validators": ["a"]}`,
			src:  `{"validators": ["b"]}`,
			want: `{"validators": ["a", "b"]}`,
		},
		{
			name: "appends a scalar onto a list",
			dst:  `{"validators": ["a"]}`,
			src:  `{"validators": "b"}`,
			want: `{"validators": ["a", "b"]}`,
		},
		{
			name: "merges nested objects key by key",
			dst:  `{"property_overrides": {"name": {"trim": true}}}`,
			src:  `{"property_overrides": {"name": {"required": true}, "id": {"type": "int"}}}`,
			want: `{"property_overrides": {"name": {"trim": true, "required": true}, "id": {"type": "int"}}}`,
		},
		{
			name: "overwrites a scalar",
			dst:  `{"enabled": false}`,
			src:  `{"enabled": true}`,
			want: `{"enabled": true}`,
		},
		{
			name: "keeps a key the source does not mention",
			dst:  `{"endpoint": "/api/v2/hosts/"}`,
			src:  `{"enabled": true}`,
			want: `{"endpoint": "/api/v2/hosts/", "enabled": true}`,
		},
		{
			name: "replaces a scalar with an object",
			dst:  `{"search_fields": "name"}`,
			src:  `{"search_fields": {"name": "id"}}`,
			want: `{"search_fields": {"name": "id"}}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var merged = mergeJSON(mustDecodeJSON(t, tt.dst), mustDecodeJSON(t, tt.src))
			payload, err := encodeJSONDocument(merged)
			require.NoError(t, err)
			require.JSONEq(t, tt.want, string(payload))
		})
	}
}

func TestMergeJSONKeepsTheKeyPositionItOverwrites(t *testing.T) {
	var merged = mergeJSON(
		mustDecodeJSON(t, `{"endpoint": "/api/v2/hosts/", "enabled": false}`),
		mustDecodeJSON(t, `{"enabled": true, "no_id": true}`),
	)
	require.Equal(t, []string{"endpoint", "enabled", "no_id"}, merged.(*jsonObject).Keys())
}

func TestMergeConfig(t *testing.T) {
	var dir = t.TempDir()
	var configDir = filepath.Join(dir, "config")
	var apiDir = filepath.Join(dir, "api", "24.6.1")

	writeFixture(t, filepath.Join(configDir, "default.json"), `{
  "api_version": "0.0.0",
  "render_api_docs": true,
  "default_remove_api_resource": ["url"]
}`)
	// The filenames order the items, not the type names.
	writeFixture(t, filepath.Join(configDir, "types", "a_thing.json"), `{
  "type_name": "Thing",
  "endpoint": "/api/v2/things/?name__exact=%s&organization=%d",
  "search_fields": [{"name": "by_name"}]
}`)
	writeFixture(t, filepath.Join(configDir, "types", "b_other.json"), `{"type_name": "Other", "enabled": true}`)
	writeFixture(t, filepath.Join(configDir, "types", "notes.txt"), `not a config`)
	writeFixture(t, filepath.Join(apiDir, "config", "default.json"), `{"api_version": "24.6.1"}`)
	writeFixture(t, filepath.Join(apiDir, "config", "types", "thing.json"), `{
  "type_name": "Thing",
  "enabled": true,
  "search_fields": [{"name": "by_id"}]
}`)

	require.NoError(t, MergeConfig(configDir, apiDir))

	payload, err := os.ReadFile(filepath.Join(apiDir, "config.json"))
	require.NoError(t, err)
	require.Equal(t, `{
  "api_version": "24.6.1",
  "render_api_docs": true,
  "default_remove_api_resource": [
    "url"
  ],
  "items": [
    {
      "type_name": "Thing",
      "endpoint": "/api/v2/things/?name__exact=%s&organization=%d",
      "search_fields": [
        {
          "name": "by_name"
        },
        {
          "name": "by_id"
        }
      ],
      "enabled": true
    },
    {
      "type_name": "Other",
      "enabled": true
    }
  ]
}`, string(payload))
}

func TestMergeConfigWithoutAnOverlay(t *testing.T) {
	var dir = t.TempDir()
	var configDir = filepath.Join(dir, "config")
	var apiDir = filepath.Join(dir, "api", "24.6.1")

	writeFixture(t, filepath.Join(configDir, "default.json"), `{"api_version": "0.0.0"}`)
	writeFixture(t, filepath.Join(configDir, "types", "thing.json"), `{"type_name": "Thing"}`)
	writeFixture(t, filepath.Join(apiDir, "config", "default.json"), `{"api_version": "24.6.1"}`)

	require.NoError(t, MergeConfig(configDir, apiDir))

	var cfg Config
	require.NoError(t, cfg.Load(filepath.Join(apiDir, "config.json")))
	require.Equal(t, "24.6.1", cfg.ApiVersion)
	require.Len(t, cfg.Items, 1)
	require.Equal(t, "Thing", cfg.Items[0].TypeName)
}

func TestMergeConfigRejectsATypeConfigWithoutATypeName(t *testing.T) {
	var dir = t.TempDir()
	var configDir = filepath.Join(dir, "config")
	var apiDir = filepath.Join(dir, "api", "24.6.1")

	writeFixture(t, filepath.Join(configDir, "default.json"), `{"api_version": "0.0.0"}`)
	writeFixture(t, filepath.Join(configDir, "types", "thing.json"), `{"name": "Thing"}`)
	writeFixture(t, filepath.Join(apiDir, "config", "default.json"), `{"api_version": "24.6.1"}`)

	require.ErrorContains(t, MergeConfig(configDir, apiDir), "no type_name")
}

func mustDecodeJSON(t *testing.T, payload string) any {
	t.Helper()
	var document, err = decodeJSONDocument([]byte(payload))
	require.NoError(t, err)
	return document
}

func writeFixture(t *testing.T, path, payload string) {
	t.Helper()
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(payload), 0o644))
}
