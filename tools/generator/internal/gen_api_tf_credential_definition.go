package internal

import (
	"fmt"
	"log"
	"strings"
	"text/template"
)

func GenerateApiTfCredentialDefinition(tpl *template.Template, config Config, item Credential, name string, resourcePath string, objmap map[string]any) (p *ModelCredential, inclDatasource bool, err error) {
	log.Printf("Generating resources for %s (Credential)", name)
	inclDatasource = false

	p = &ModelCredential{}
	if err = p.Update(config, item, objmap); err != nil {
		return nil, inclDatasource, err
	}

	resourcePath = fmt.Sprintf("%s/credentials/%s", resourcePath, strings.ToLower(p.TypeName))

	var tpls = []struct {
		Filename string
		Template string
		Skip     bool
		Data     map[string]any
	}{
		{
			Filename: fmt.Sprintf("%s/model.go", resourcePath),
			Template: "tf_credential_model.go.tpl",
		},
		{
			Filename: fmt.Sprintf("%s/resource.go", resourcePath),
			Template: "tf_credential_resource.go.tpl",
		},
		{
			Filename: fmt.Sprintf("%s/resource_test.go", resourcePath),
			Template: "tf_credential_resource_test.go.tpl",
		},
		{
			Filename: fmt.Sprintf("%s/data_source.go", resourcePath),
			Template: "tf_credential_data_source.go.tpl",
			Skip:     !inclDatasource,
		},
	}

	for _, t := range tpls {
		if !t.Skip {
			_ = renderTemplate(tpl, t.Filename, t.Template, p)
		}
	}

	return p, inclDatasource, err
}
