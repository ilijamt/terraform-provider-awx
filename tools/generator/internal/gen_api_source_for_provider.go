package internal

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func GenerateApiSourcesForProvider(tpl *template.Template, config Config, resourcePath string, resources []string, dataSources []string) error {
	var f *os.File
	var err error
	var filename = fmt.Sprintf("%s/gen_sources.go", resourcePath)
	log.Printf("Generating datasources into %s", filename)
	if f, err = os.Create(filename); err != nil {
		return err
	}

	return tpl.ExecuteTemplate(f, "sources.go.tpl", map[string]any{
		"ApiVersion":  config.ApiVersion,
		"PackageName": config.PackageName("awx"),
		"Resources":   resources,
		"DataSources": dataSources,
	})
}
