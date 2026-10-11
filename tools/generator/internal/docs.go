package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/iancoleman/strcase"
)

// ImportExample is the import ID of one resource, rendered into the
// examples/resources/<type>/ files tfplugindocs turns into its Import section.
type ImportExample struct {
	TypeName string
	Parts    []string
}

func (e ImportExample) Format() string {
	return "<" + strings.Join(e.Parts, ">/<") + ">"
}

func (e ImportExample) ExampleID() string {
	ids := make([]string, len(e.Parts))
	for i := range e.Parts {
		ids[i] = strconv.Itoa(i + 1)
	}
	return strings.Join(ids, "/")
}

// ImportExamples mirrors the ImportState each resource is generated with: see
// tf_object_resource.partial.tpl, tf_associate_disassociate.go.tpl and
// tf_survey_spec.go.tpl. Notification associations carry an option and are
// not importable.
func ImportExamples(item Item) []ImportExample {
	if !item.Enabled || item.NoTerraformResource {
		return nil
	}
	if IsCredentialTypeItem(item) || IsNotificationTypeItem(item) {
		return []ImportExample{{TypeName: "awx_" + item.TypeName, Parts: []string{"id"}}}
	}

	var examples []ImportExample
	if !item.NoId && !item.NoImport {
		parts := []string{item.IdKey}
		if item.CreateEndpoint != nil {
			parts = []string{item.CreateEndpoint.IdAttribute, item.IdKey}
		}
		examples = append(examples, ImportExample{TypeName: "awx_" + item.TypeName, Parts: parts})
	}
	for _, adg := range item.AssociateDisassociateGroups {
		if strings.HasPrefix(adg.AssociateType, "notification_") {
			continue
		}
		parent, child := strcase.ToSnake(adg.Name), strcase.ToSnake(adg.Type)
		examples = append(examples, ImportExample{
			TypeName: fmt.Sprintf("awx_%s_associate_%s", parent, child),
			Parts:    []string{parent + "_id", child + "_id"},
		})
	}
	if item.HasSurveySpec {
		name := strcase.ToSnake(item.Name)
		examples = append(examples, ImportExample{TypeName: fmt.Sprintf("awx_%s_survey_spec", name), Parts: []string{name + "_id"}})
	}
	return examples
}

func GenerateImportExamples(tpl *template.Template, dir string, examples []ImportExample) error {
	for _, e := range examples {
		resDir := filepath.Join(dir, "resources", e.TypeName)
		if err := os.MkdirAll(resDir, os.ModePerm); err != nil {
			return err
		}
		if err := renderTemplate(tpl, filepath.Join(resDir, "import.sh"), "import.sh.tpl", e); err != nil {
			return err
		}
		if err := renderTemplate(tpl, filepath.Join(resDir, "import-by-string-id.tf"), "import-by-string-id.tf.tpl", e); err != nil {
			return err
		}
	}
	return nil
}

// Docs subcategories group the registry sidebar and mirror the AWX UI
// navigation. A prefix matches the whole name or the name followed by "_", so
// "group" does not swallow "instance_group".
var docsSubcategories = []struct {
	name     string
	prefixes []string
}{
	{"Templates", []string{"job_template", "workflow_job_template", "schedule", "label", "ad_hoc_command"}},
	{"Projects", []string{"project"}},
	{"Inventories", []string{"inventory", "constructed_inventories", "host", "group"}},
	{"Credentials", []string{"credential"}},
	{"Notifications", []string{"notification_template"}},
	{"Access", []string{"organization", "team", "user", "role", "token", "application", "me"}},
	{"Infrastructure", []string{"instance", "execution_environment"}},
	{"Settings", []string{"settings"}},
	{"System", []string{"config", "ping"}},
}

func DocsSubcategory(name string) (string, bool) {
	for _, s := range docsSubcategories {
		for _, p := range s.prefixes {
			if name == p || strings.HasPrefix(name, p+"_") {
				return s.name, true
			}
		}
	}
	return "", false
}

const emptySubcategory = "subcategory: \"\"\n"

// SetDocsSubcategories fills in the subcategory of each resource and data
// source page. tfplugindocs always renders it empty and its templates cannot
// branch on the resource name.
func SetDocsSubcategories(dir string) error {
	for _, kind := range []string{"resources", "data-sources"} {
		files, err := filepath.Glob(filepath.Join(dir, kind, "*.md"))
		if err != nil {
			return err
		}
		for _, file := range files {
			name := strings.TrimSuffix(filepath.Base(file), ".md")
			category, ok := DocsSubcategory(name)
			if !ok {
				return fmt.Errorf("%s: no subcategory matches %q, add it to docsSubcategories", file, name)
			}
			content, err := os.ReadFile(file)
			if err != nil {
				return err
			}
			if !strings.Contains(string(content), emptySubcategory) {
				return fmt.Errorf("%s: no empty subcategory to replace", file)
			}
			updated := strings.Replace(string(content), emptySubcategory, fmt.Sprintf("subcategory: %q\n", category), 1)
			if err := os.WriteFile(file, []byte(updated), 0o644); err != nil {
				return err
			}
		}
	}
	return nil
}

// GenerateProviderDocs writes the provider usage example and the tfplugindocs
// template for the provider index page, which names the AWX version.
func GenerateProviderDocs(tpl *template.Template, config Config, examplesDir, templatesDir string) error {
	providerDir := filepath.Join(examplesDir, "provider")
	if err := os.MkdirAll(providerDir, os.ModePerm); err != nil {
		return err
	}
	if err := renderTemplate(tpl, filepath.Join(providerDir, "provider.tf"), "provider.tf.tpl", nil); err != nil {
		return err
	}
	if err := os.MkdirAll(templatesDir, os.ModePerm); err != nil {
		return err
	}
	return renderTemplate(tpl, filepath.Join(templatesDir, "index.md.tmpl"), "index.md.tmpl.tpl", config)
}

// usageExample is one resource or data source block lifted from the
// hand-written configurations under the examples directory.
type usageExample struct {
	kind     string // "resources" or "data-sources"
	typeName string
	source   []byte
	fromOwn  bool
}

// GenerateUsageExamples turns the hand-written example configurations in
// <examplesDir>/<name>/*.tf into the per-type example files tfplugindocs
// renders as Example Usage. The block from the directory named after the type
// wins, otherwise the first one in path order. Blocks holding a private key
// are skipped so test keys do not end up on the docs pages.
func GenerateUsageExamples(examplesDir string) error {
	files, err := filepath.Glob(filepath.Join(examplesDir, "*", "*.tf"))
	if err != nil {
		return err
	}
	chosen := map[string]usageExample{}
	for _, file := range files {
		dir := filepath.Base(filepath.Dir(file))
		if dir == "resources" || dir == "data-sources" || dir == "provider" {
			continue
		}
		examples, err := usageExamplesFrom(file)
		if err != nil {
			return err
		}
		for _, e := range examples {
			e.fromOwn = strings.TrimPrefix(e.typeName, "awx_") == dir
			key := e.kind + "/" + e.typeName
			if current, ok := chosen[key]; !ok || (e.fromOwn && !current.fromOwn) {
				chosen[key] = e
			}
		}
	}

	for _, e := range chosen {
		filename := "resource.tf"
		if e.kind == "data-sources" {
			filename = "data-source.tf"
		}
		dir := filepath.Join(examplesDir, e.kind, e.typeName)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, filename), hclwrite.Format(e.source), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func usageExamplesFrom(file string) ([]usageExample, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	f, diags := hclsyntax.ParseConfig(src, file, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, diags
	}
	var examples []usageExample
	for _, block := range f.Body.(*hclsyntax.Body).Blocks {
		if len(block.Labels) != 2 || !strings.HasPrefix(block.Labels[0], "awx_") {
			continue
		}
		kind := map[string]string{"resource": "resources", "data": "data-sources"}[block.Type]
		if kind == "" {
			continue
		}
		source := src[block.TypeRange.Start.Byte:block.CloseBraceRange.End.Byte]
		if strings.Contains(string(source), "PRIVATE KEY") {
			continue
		}
		examples = append(examples, usageExample{kind: kind, typeName: block.Labels[0], source: append(append([]byte{}, source...), '\n')})
	}
	return examples, nil
}
