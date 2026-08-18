package internal

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/tools/generator"
)

// Docs must land under apiResourcePath rather than the working directory, or
// generating against a scratch API tree writes into the repo.
func TestGenerateApiTfDefinitionWritesDocsUnderApiResourcePath(t *testing.T) {
	tpl, err := template.New("").Funcs(FuncMap).ParseFS(generator.Fs(), "templates/*.tpl", "templates/terraform/*.tpl")
	require.NoError(t, err)

	apiResourcePath := t.TempDir()
	resourcePath := t.TempDir()
	require.NoError(t, os.Mkdir(filepath.Join(apiResourcePath, "docs"), os.ModePerm))

	cfg := Config{ApiVersion: "24.6.1", RenderApiDocs: true}
	item := Item{
		Name:                   "Widget",
		TypeName:               "widget",
		Endpoint:               "/api/v2/widgets/",
		IdKey:                  "id",
		Enabled:                true,
		ApiPropertyResourceKey: "POST",
		ApiPropertyDataKey:     "GET",
	}
	objmap := map[string]any{
		"description": "Widgets",
		"actions": map[string]any{
			"GET":  map[string]any{"id": map[string]any{"type": "integer", "label": "ID"}},
			"POST": map[string]any{"name": map[string]any{"type": "string", "label": "Name", "required": true}},
		},
	}

	// Run from somewhere else entirely so a CWD-relative path cannot pass.
	wd, err := os.Getwd()
	require.NoError(t, err)
	require.NoError(t, os.Chdir(t.TempDir()))
	t.Cleanup(func() { _ = os.Chdir(wd) })

	_, _, _, err = GenerateApiTfDefinition(tpl, cfg, item, apiResourcePath, resourcePath, item.Name, objmap)
	require.NoError(t, err)

	assert.FileExists(t, filepath.Join(apiResourcePath, "docs", "widget.md"))
	assert.FileExists(t, filepath.Join(resourcePath, "gen_obj_widget.go"))
}

// A create endpoint carries its parent id in the URL, so the attribute holding
// it exists only on the write side. Leaking it into the read action would have
// UpdateFromApiData null it on every read, since no response ever carries it.
func TestGenerateApiTfDefinitionSplitCreateEndpoint(t *testing.T) {
	tpl, err := template.New("").Funcs(FuncMap).ParseFS(generator.Fs(), "templates/*.tpl", "templates/terraform/*.tpl")
	require.NoError(t, err)

	apiResourcePath := t.TempDir()
	resourcePath := t.TempDir()
	require.NoError(t, os.Mkdir(filepath.Join(apiResourcePath, "docs"), os.ModePerm))

	item := Item{
		Name:                   "Widget",
		TypeName:               "widget",
		Endpoint:               "/api/v2/widgets/",
		IdKey:                  "id",
		Enabled:                true,
		ApiPropertyResourceKey: "POST",
		ApiPropertyDataKey:     "GET",
		CreateEndpoint: &CreateEndpointConfig{
			Endpoint:    "/api/v2/gadgets/%d/create_widget/",
			IdAttribute: "gadget_id",
		},
		ApiDataOverrideResource: map[string]map[string]any{
			"gadget_id": {"type": "integer", "required": true, "write_only": true, "label": "Gadget"},
		},
	}
	objmap := map[string]any{
		"description": "Widgets",
		"actions": map[string]any{
			"GET":  map[string]any{"id": map[string]any{"type": "integer", "label": "ID"}},
			"POST": map[string]any{"name": map[string]any{"type": "string", "label": "Name", "required": true}},
		},
	}

	_, _, _, err = GenerateApiTfDefinition(tpl, Config{ApiVersion: "24.6.1"}, item, apiResourcePath, resourcePath, item.Name, objmap)
	require.NoError(t, err)

	generated, err := os.ReadFile(filepath.Join(resourcePath, "gen_obj_widget.go"))
	require.NoError(t, err)

	assert.Contains(t, string(generated), `"gadget_id": schema.Int64Attribute{`)
	assert.Contains(t, string(generated), `fmt.Sprintf("/api/v2/gadgets/%d/create_widget/", m.GadgetId.ValueInt64())`)
	assert.Contains(t, string(generated), `ImportIDParts: []string{"gadget_id", "id"}`)

	_, rest, found := strings.Cut(string(generated), "func (o *widgetTerraformModel) UpdateFromApiData")
	require.True(t, found)
	reader, _, found := strings.Cut(rest, "\n}")
	require.True(t, found)
	assert.NotContains(t, reader, "GadgetId")
}
