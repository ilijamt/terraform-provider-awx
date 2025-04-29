package internal

import (
	"embed"
	"path/filepath"
	"strings"
	"text/template"
)

func LoadTemplates(fs embed.FS) (*template.Template, error) {
	tpl := template.New("").Funcs(FuncMap)
	err := parseTemplatesFromFS(tpl, fs, "templates")
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

// Helper function to recursively parse templates from the filesystem
func parseTemplatesFromFS(tpl *template.Template, fs embed.FS, dir string) error {
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			// Recursively process subdirectories
			if err := parseTemplatesFromFS(tpl, fs, path); err != nil {
				return err
			}
			continue
		}

		// Only process template files
		if !strings.HasSuffix(entry.Name(), ".tpl") {
			continue
		}

		// Read the template content
		content, err := fs.ReadFile(path)
		if err != nil {
			return err
		}

		// Use the relative path as the template name (without the "templates/" prefix)
		templateName := path
		if strings.HasPrefix(templateName, "templates/") {
			templateName = templateName[len("templates/"):]
		}

		// Parse the template
		_, err = tpl.New(templateName).Parse(string(content))
		if err != nil {
			return err
		}
	}

	return nil
}
