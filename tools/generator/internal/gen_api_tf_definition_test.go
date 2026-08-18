package internal

import (
	"os"
	"path/filepath"
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
