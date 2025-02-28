package internal

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GenerateApiTfDefinition(tpl *template.Template, config Config, val Item, resourcePath, name string, objmap map[string]any) (data map[string]any, p *ModelConfig, dr Deprecated, err error) {
	log.Printf("Generating resources for %s", name)

	if _, ok := objmap["actions"]; !ok {
		log.Printf("No actions for %s, skipping ....", name)
		return nil, nil, dr, nil
	}

	var description string
	if v, ok := objmap["description"].(string); ok {
		description = v
	}

	var item = &ModelConfig{
		Name:        name,
		Description: description,
	}
	_ = item.Update(config, val)

	// ---------------------
	var propertyWriteOnlyData = make(map[string]any)
	var propertyWriteOnlyKeys []string

	// ---------------------
	var propertyGetData = make(map[string]any)
	var propertyGetKeys []string
	if props, ok := objmap["actions"].(map[string]any)[val.ApiPropertyDataKey].(map[string]any); ok {
		for _, field := range append(config.DefaultRemoveApiDataSource, val.RemoveFieldsDataSource...) {
			delete(props, field)
		}

		for key, value := range props {
			value.(map[string]any)["name"] = key
			_, _ = item.UpdateProperty(TypeRead, key, val.PropertyOverrides[key], value.(map[string]any), val)
			propertyGetKeys = append(propertyGetKeys, key)
			propertyGetData[key] = value
		}
	}

	// ---------------------
	var propertyPostData = make(map[string]any)
	var propertyPostKeys []string
	if props, ok := objmap["actions"].(map[string]any)[val.ApiPropertyResourceKey].(map[string]any); ok {
		for _, field := range append(config.DefaultRemoveApiResource, val.RemoveFieldsResource...) {
			delete(props, field)
		}

		for key, value := range props {
			value.(map[string]any)["name"] = key
			_, _ = item.UpdateProperty(TypeWrite, key, val.PropertyOverrides[key], value.(map[string]any), val)
			if writeOnly, ok := value.(map[string]any)["write_only"].(bool); ok && writeOnly {
				if val.SkipWriteOnly {
					continue
				}
				propertyWriteOnlyKeys = append(propertyWriteOnlyKeys, key)
				propertyWriteOnlyData[key] = value
				continue
			}

			propertyPostKeys = append(propertyPostKeys, key)
			propertyPostData[key] = value
		}
	}

	sort.Strings(propertyGetKeys)
	sort.Strings(propertyPostKeys)
	sort.Strings(propertyWriteOnlyKeys)

	// ---------------------

	data = map[string]any{
		"ApiVersion":            config.ApiVersion,
		"PackageName":           config.PackageName("awx"),
		"Name":                  name,
		"Endpoint":              val.Endpoint,
		"Description":           objmap["description"],
		"PropertyGetKeys":       propertyGetKeys,
		"PropertyGetData":       propertyGetData,
		"PropertyPostKeys":      propertyPostKeys,
		"PropertyPostData":      propertyPostData,
		"PropertyWriteOnlyKeys": propertyWriteOnlyKeys,
		"PropertyWriteOnlyData": propertyWriteOnlyData,
		"Config":                val,
	}

	var tpls = []struct {
		Filename string
		Template string
		Render   bool
		IsNew    bool
		Data     map[string]any
	}{
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_model.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_model.go.tpl",
			Render:   true,
			IsNew:    true,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_data_source.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_data_source.go.tpl",
			Render:   !item.NoTerraformDataSource,
			IsNew:    true,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_resource.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_resource.go.tpl",
			Render:   !item.NoTerraformResource,
			IsNew:    true,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_object_roles.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_resource_object_role.go.tpl",
			Render:   item.HasObjectRoles,
			IsNew:    true,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_survey_spec.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_survey_spec.go.tpl",
			Render:   item.HasSurveySpec,
			IsNew:    true,
		},
		{
			Filename: fmt.Sprintf("resources/api/%s/docs/%s.md", config.ApiVersion, strings.ToLower(val.TypeName)),
			Template: "tf_api_description.md.tpl",
			Render:   item.RenderApiDocs,
			IsNew:    true,
		},
	}

	for _, adg := range val.AssociateDisassociateGroups {
		_, deprecated := item.DeprecatedParts["AssociateDisassociateGroups"]
		tpls = append(tpls, struct {
			Filename string
			Template string
			Render   bool
			IsNew    bool
			Data     map[string]any
		}{
			Filename: fmt.Sprintf("%s/gen_obj_%s_adg_%s.go", resourcePath,
				strings.ToLower(val.TypeName), strings.ToLower(adg.Type)),
			Template: "tf_associate_disassociate.go.tpl",
			Render:   true,
			Data:     adg.Map(deprecated),
		})

		if deprecated && val.Enabled {
			dr.Resources = append(
				dr.Resources,
				strcase.ToDelimited(
					fmt.Sprintf(
						"%sAssociateDisassociate%s",
						strcase.ToLowerCamel(adg.Name), adg.Type,
					), '_',
				),
			)
			dr.DataSources = append(
				dr.DataSources,
				strcase.ToDelimited(
					fmt.Sprintf(
						"%sObjectRoles",
						strcase.ToLowerCamel(item.Name),
					), '_',
				),
			)
		}
	}

	_ = item.Process(config, val)

	// ---------------------

	if val.Enabled {
		dr.Properties = []DeprecatedProperties{
			{
				Resource:        item.Name,
				WriteProperties: item.DeprecatedWriteProperties,
				ReadProperties:  item.DeprecatedReadProperties,
			},
		}
		for _, t := range tpls {
			if !t.Render {
				log.Printf("Rendering of %s into %s skipped.", t.Template, t.Filename)
				continue
			}
			d := data
			if len(t.Data) > 0 {
				d = t.Data
				d["PackageName"] = data["PackageName"]
			}

			if t.IsNew {
				d = item.ToMap()
				d["Config"] = val
			}

			if err = renderTemplate(
				tpl,
				t.Filename,
				t.Template,
				d,
			); err != nil {
				return data, item, dr, err
			}
		}
	}

	return data, item, dr, nil
}
