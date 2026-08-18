package internal

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// applyApiDataOverride creates the field when the source instance never
// reported it. A deployment that pins a setting in a config file drops it from
// the write action entirely (TOWER_URL_BASE arrives with defined_in_file true
// and no PUT entry), which would delete the attribute from the schema. Runs
// after the remove_fields_* pruning so the override wins.
//
// Item.ApiDataOverride reaches both actions; ApiDataOverrideResource only the
// write one, for fields that go out in a request but never come back, such as
// the parent id a create endpoint carries in its URL.
func applyApiDataOverride(props map[string]any, overrides map[string]map[string]any) {
	for key, override := range overrides {
		field, ok := props[key].(map[string]any)
		if !ok {
			field = make(map[string]any)
			props[key] = field
		}
		for k, v := range override {
			field[k] = v
		}
	}
}

func GenerateApiTfDefinition(tpl *template.Template, config Config, val Item, apiResourcePath, resourcePath, name string, objmap map[string]any) (data map[string]any, p *ModelConfig, dr Deprecated, err error) {
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
		applyApiDataOverride(props, val.ApiDataOverride)

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
		applyApiDataOverride(props, val.ApiDataOverride)
		applyApiDataOverride(props, val.ApiDataOverrideResource)

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
			// One file per object: model + resource + data source. Sections
			// inside tf_object.go.tpl are gated by NoTerraformResource /
			// NoTerraformDataSource so the consolidated file works for
			// data-source-only objects (e.g. Me) too.
			Filename: fmt.Sprintf("%s/gen_obj_%s.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_object.go.tpl",
			Render:   true,
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
			Filename: fmt.Sprintf("%s/docs/%s.md", apiResourcePath, strings.ToLower(val.TypeName)),
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
