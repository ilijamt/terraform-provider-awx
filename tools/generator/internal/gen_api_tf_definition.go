package internal

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"text/template"
)

func GenerateApiTfDefinition(tpl *template.Template, config Config, val Item, resourcePath, name string, objmap map[string]any) error {
	var err error

	log.Printf("Generating resources for %s", name)

	if _, ok := objmap["actions"]; !ok {
		log.Printf("No actions for %s, skipping ....", name)
		return nil
	}

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
			if override, ok := val.PropertyOverrides[key]; ok {
				if "" != override.Type {
					value.(map[string]any)["type"] = override.Type
				}
				if "" != override.Description {
					value.(map[string]any)["help_text"] = override.Description
				}
				if override.Sensitive {
					value.(map[string]any)["sensitive"] = override.Sensitive
				}
				if override.Required {
					value.(map[string]any)["required"] = override.Required
				}
				value.(map[string]any)["trim"] = override.Trim
				value.(map[string]any)["post_wrap"] = override.PostWrap
			}
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
			if override, ok := val.PropertyOverrides[key]; ok {
				if "" != override.Type {
					value.(map[string]any)["type"] = override.Type
				}
				if "" != override.DefaultValue {
					value.(map[string]any)["default"] = override.DefaultValue
				}
				if "" != override.Description {
					value.(map[string]any)["help_text"] = override.Description
				}
				if override.Sensitive {
					value.(map[string]any)["sensitive"] = override.Sensitive
				}
				if override.Required {
					value.(map[string]any)["required"] = override.Required
				}
				value.(map[string]any)["trim"] = override.Trim
				value.(map[string]any)["post_wrap"] = override.PostWrap
			}

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

	// ---------------------

	var data = map[string]any{
		"ApiVersion":            config.ApiVersion,
		"PackageName":           "awx",
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

	// ---------------------

	if val.Enabled {
		if err = renderTemplate(tpl, fmt.Sprintf("%s/gen_obj_%s.go", resourcePath, strings.ToLower(val.TypeName)), "tf_full_element.go.tpl", data); err != nil {
			return err
		}
	}

	if config.RenderApiDocs {
		if err = renderTemplate(tpl, fmt.Sprintf("resources/api/docs/%s.md", strings.ToLower(val.TypeName)), "tf_api_description.md.tpl", data); err != nil {
			return err
		}
	}

	return nil
}
