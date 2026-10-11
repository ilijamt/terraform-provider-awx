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

func TestImportExamples(t *testing.T) {
	for _, tt := range []struct {
		name string
		item Item
		want []ImportExample
	}{
		{
			name: "plain id",
			item: Item{Enabled: true, Name: "Organization", TypeName: "organization", IdKey: "id"},
			want: []ImportExample{{TypeName: "awx_organization", Parts: []string{"id"}}},
		},
		{
			name: "create endpoint makes the id composite",
			item: Item{Enabled: true, TypeName: "node_approval", IdKey: "id", CreateEndpoint: &CreateEndpointConfig{IdAttribute: "node_id"}},
			want: []ImportExample{{TypeName: "awx_node_approval", Parts: []string{"node_id", "id"}}},
		},
		{
			name: "no import keeps associations and survey",
			item: Item{
				Enabled: true, Name: "JobTemplate", TypeName: "job_template", IdKey: "id", NoImport: true, HasSurveySpec: true,
				AssociateDisassociateGroups: []AssociateDisassociateGroup{
					{Name: "JobTemplate", Type: "Credential"},
					{Name: "JobTemplate", Type: "NotificationTemplate", AssociateType: "notification_job_template"},
				},
			},
			want: []ImportExample{
				{TypeName: "awx_job_template_associate_credential", Parts: []string{"job_template_id", "credential_id"}},
				{TypeName: "awx_job_template_survey_spec", Parts: []string{"job_template_id"}},
			},
		},
		{name: "no id", item: Item{Enabled: true, TypeName: "settings_ui", NoId: true}},
		{name: "disabled", item: Item{TypeName: "organization", IdKey: "id"}},
		{name: "no terraform resource", item: Item{Enabled: true, TypeName: "me", IdKey: "id", NoTerraformResource: true}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ImportExamples(tt.item))
		})
	}
}

func TestImportExample_Format(t *testing.T) {
	e := ImportExample{Parts: []string{"parent_id", "id"}}
	assert.Equal(t, "<parent_id>/<id>", e.Format())
	assert.Equal(t, "1/2", e.ExampleID())
}

func TestDocsSubcategory(t *testing.T) {
	for _, tt := range []struct {
		name   string
		want   string
		wantOk bool
	}{
		{name: "credential_aws", want: "Credentials", wantOk: true},
		{name: "instance_group", want: "Infrastructure", wantOk: true},
		{name: "group", want: "Inventories", wantOk: true},
		{name: "workflow_job_template_node_associate_always_node", want: "Templates", wantOk: true},
		{name: "me", want: "Access", wantOk: true},
		{name: "metrics"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := DocsSubcategory(tt.name)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetDocsSubcategories(t *testing.T) {
	dir := t.TempDir()
	page := filepath.Join(dir, "resources", "host.md")
	require.NoError(t, os.MkdirAll(filepath.Dir(page), 0o755))
	require.NoError(t, os.WriteFile(page, []byte("---\nsubcategory: \"\"\n---\n"), 0o644))

	require.NoError(t, SetDocsSubcategories(dir))

	content, err := os.ReadFile(page)
	require.NoError(t, err)
	assert.Equal(t, "---\nsubcategory: \"Inventories\"\n---\n", string(content))
}

func TestGenerateProviderDocs(t *testing.T) {
	tpl, err := template.New("").Funcs(FuncMap).ParseFS(generator.Fs(), "templates/*.tpl")
	require.NoError(t, err)
	dir := t.TempDir()

	require.NoError(t, GenerateProviderDocs(tpl, Config{ApiVersion: "24.6.1"}, filepath.Join(dir, "examples"), filepath.Join(dir, "templates")))

	index, err := os.ReadFile(filepath.Join(dir, "templates", "index.md.tmpl"))
	require.NoError(t, err)
	assert.Contains(t, string(index), "AWX 24.6.1,")
	assert.Contains(t, string(index), "{{ tffile .ExampleFile }}", "tfplugindocs actions must pass through unrendered")
	assert.Contains(t, string(index), "{{ .SchemaMarkdown | trimspace }}")
	assert.FileExists(t, filepath.Join(dir, "examples", "provider", "provider.tf"))
}
