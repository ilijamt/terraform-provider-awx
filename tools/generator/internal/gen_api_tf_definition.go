package internal

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"text/template"
)

func getItemElementListType(value map[string]any) (any, error) {
	if v, ok := value["child"]; ok {
		if t, ok := v.(map[string]any)["type"]; ok {
			switch t.(type) {
			case string:
				if t == "field" {
					t = "string"
				}
				return t, nil
			}
			return "", fmt.Errorf("unknown type for list type")
		}
	}
	return "", fmt.Errorf("no list element type found")
}

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

	/*
		{{- range $key := .PropertyPostKeys }}
		{{- with (index $.PropertyPostData $key) }}
		            "{{ $key | lowerCase }}": schema.{{ tf_attribute_type . }}Attribute{
		{{- if eq (tf_attribute_type .) "List" }}
						ElementType: types.{{ camelCase .element_type }}Type,
		{{- end }}
		                Description: {{ escape_quotes (default .help_text .label) }},
		{{- if (hasKey . "sensitive") }}
		                Sensitive:   {{ .sensitive }},
		{{- end }}
		                Required:    {{ .required }},
		                Optional:    {{ not .required }},
		{{- if not .required }}
		                Computed:    true,
		{{- end }}
		{{- if (hasKey . "default") }}
		{{- if and (hasKey $.PropertyPostData $key) (eq (awx2go_value (index $.PropertyPostData $key)) "types.StringValue") (ne .default nil) }}
						Default:     {{ tf_attribute_type . | lowerCase }}default.Static{{ tf_attribute_type . }}(`{{ convertDefaultValue .default }}`),
		{{- else if and (hasKey $.PropertyPostData $key) (eq (awx2go_value (index $.PropertyPostData $key)) "types.Int64Value") (ne .default nil) }}
						Default:     {{ tf_attribute_type . | lowerCase }}default.Static{{ tf_attribute_type . }}({{ convertDefaultValue .default }}),
		{{- end }}
		{{- end }}
				        PlanModifiers: []planmodifier.{{ tf_attribute_type . }} {
		{{- if not .required }}
		                    {{ tf_attribute_type . | lowerCase }}planmodifier.UseStateForUnknown(),
		{{- end }}
		                },
	*/

	var processOverride = func(
		value map[string]any,
		key string,
		override PropertyOverride,
	) {
		value["name"] = key
		if v, err := getItemElementListType(value); err == nil {
			value["element_type"] = v
		}

		if override, ok := val.PropertyOverrides[key]; ok {
			if "" != override.Type {
				value["type"] = override.Type
			}
			if "" != override.DefaultValue {
				value["default"] = override.DefaultValue
			}
			if "" != override.Description {
				value["help_text"] = override.Description
			}
			if override.Sensitive {
				value["sensitive"] = override.Sensitive
			}
			if override.Required {
				value["required"] = override.Required
			}
			if "" != override.ElementType {
				value["element_type"] = override.ElementType
			}
			value["trim"] = override.Trim
			value["post_wrap"] = override.PostWrap
		}
	}

	// ---------------------
	var propertyGetData = make(map[string]any)
	var propertyGetKeys []string
	if props, ok := objmap["actions"].(map[string]any)[val.ApiPropertyDataKey].(map[string]any); ok {

		for _, field := range append(config.DefaultRemoveApiDataSource, val.RemoveFieldsDataSource...) {
			delete(props, field)
		}

		for key, value := range props {
			var override PropertyOverride
			if val, ok := val.PropertyOverrides[key]; ok {
				override = val
			}
			processOverride(value.(map[string]any), key, override)
			//value.(map[string]any)["name"] = key
			//if v, err := getItemElementListType(value.(map[string]any)); err == nil {
			//	value.(map[string]any)["element_type"] = v
			//}
			//
			//if override, ok := val.PropertyOverrides[key]; ok {
			//	if "" != override.Type {
			//		value.(map[string]any)["type"] = override.Type
			//	}
			//	if "" != override.DefaultValue {
			//		value.(map[string]any)["default"] = override.DefaultValue
			//	}
			//	if "" != override.Description {
			//		value.(map[string]any)["help_text"] = override.Description
			//	}
			//	if override.Sensitive {
			//		value.(map[string]any)["sensitive"] = override.Sensitive
			//	}
			//	if override.Required {
			//		value.(map[string]any)["required"] = override.Required
			//	}
			//	if "" != override.ElementType {
			//		value.(map[string]any)["element_type"] = override.ElementType
			//	}
			//	value.(map[string]any)["trim"] = override.Trim
			//	value.(map[string]any)["post_wrap"] = override.PostWrap
			//}
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
			if v, err := getItemElementListType(value.(map[string]any)); err == nil {
				value.(map[string]any)["element_type"] = v
			}

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
				if "" != override.ElementType {
					value.(map[string]any)["element_type"] = override.ElementType
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

	var tpls = []struct {
		Filename string
		Template string
		Render   bool
		Data     map[string]any
	}{
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_model.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_model.go.tpl",
			Render:   true,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_data_source.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_data_source.go.tpl",
			Render:   !val.NoTerraformDataSource,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_resource.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_resource.go.tpl",
			Render:   !val.NoTerraformResource,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_object_roles.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_resource_object_role.go.tpl",
			Render:   val.HasObjectRoles,
		},
		{
			Filename: fmt.Sprintf("%s/gen_obj_%s_survey_spec.go", resourcePath, strings.ToLower(val.TypeName)),
			Template: "tf_survey_spec.go.tpl",
			Render:   val.HasSurveySpec,
		},
		{
			Filename: fmt.Sprintf("resources/api/docs/%s.md", strings.ToLower(val.TypeName)),
			Template: "tf_api_description.go.tpl",
			Render:   config.RenderApiDocs,
		},
	}

	for _, adg := range val.AssociateDisassociateGroups {
		tpls = append(tpls, struct {
			Filename string
			Template string
			Render   bool
			Data     map[string]any
		}{
			Filename: fmt.Sprintf("%s/gen_obj_%s_adg_%s.go", resourcePath,
				strings.ToLower(val.TypeName), strings.ToLower(adg.Type)),
			Template: "tf_associate_disassociate.go.tpl",
			Render:   true,
			Data:     adg.Map(),
		})
	}

	// ---------------------

	if val.Enabled {
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
			if err = renderTemplate(
				tpl,
				t.Filename,
				t.Template,
				d,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
