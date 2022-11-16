package internal

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	p "path"
	"strconv"
	"strings"
	"text/template"
)

func init() {
	strcase.ConfigureAcronym("id", "ID")
}

func renderTemplate(tpl *template.Template, filename, template string, data any) (err error) {
	log.Printf("Rendering %s to %s", template, filename)
	var f *os.File
	if f, err = os.Create(filename); err != nil {
		return err
	}

	if err = tpl.ExecuteTemplate(f, template, data); err != nil {
		return err
	}
	return nil
}

var FuncMap = template.FuncMap{
	"url_path_clean": func(in string) string {
		return p.Clean(in)
	},
	"escape_quotes": func(in string) string {
		return fmt.Sprintf("%q", in)
	},
	"property_case": func(in string, config Item) string {
		if config.PropertyNameLeaveAsIs {
			return in
		}
		return strcase.ToCamel(in)
	},
	"convertDefaultValue": func(in any) any {
		switch in.(type) {
		case map[string]any:
			payload, err := json.Marshal(in)
			if err != nil {
				return "{}"
			}
			return string(payload)
		}
		return in
	},
	"default": func(in any, def any) any {
		if in == nil {
			return def
		}
		return in
	},
	"snakeCase": strcase.ToSnake,
	"camelCase": strcase.ToCamel,
	"setPropertyCase": func(in string) string {
		return strcase.ToCamel(strcase.ToSnake(in))
	},
	"lowerCase": func(in string) string {
		return strings.ToLower(in)
	},
	"lowerCamelCase": strcase.ToLowerCamel,
	"hasKey": func(d map[string]any, key string) bool {
		_, ok := d[key]
		return ok
	},
	"awx_is_property_searchable": func(fields []SearchGroup, needle string) bool {
		for _, field := range fields {
			for _, v := range field.Fields {
				if v.Name == needle {
					return true
				}
			}
		}
		return false
	},
	"awx_generate_attribute_validator": func(fields []SearchGroup, needle string) (attrs map[string][]string) {
		attrs = make(map[string][]string)

		var hasMultiFieldSearch = false
		for _, field := range fields {
			hasMultiFieldSearch = len(field.Fields) > 1
			if hasMultiFieldSearch {
				break
			}
		}

		if hasMultiFieldSearch {
			var sourceFieldName string
			// 1. Find the one this needle belongs to and create a AlsoRequires for all of them
			for _, field := range fields {
				fieldHasKey := false
				for _, v := range field.Fields {
					fieldHasKey = v.Name == needle
					if fieldHasKey {
						sourceFieldName = field.Name
						break
					}
				}

				if fieldHasKey {
					for _, v := range field.Fields {
						if v.Name == needle {
							continue
						}
						attrs["AlsoRequires"] = append(attrs["AlsoRequires"], v.Name)
					}
				}
			}

			// 2. Add all the other fields as Conflicts With
			for _, field := range fields {
				if sourceFieldName == field.Name {
					continue
				}
				for _, v := range field.Fields {
					attrs["ConflictsWith"] = append(attrs["ConflictsWith"], v.Name)
				}
			}
		} else {
			for _, field := range fields {
				for _, v := range field.Fields {
					attrs["ExactlyOneOf"] = append(attrs["ExactlyOneOf"], v.Name)
				}
			}
		}
		return
	},
	"awx2go_primitive_type": func(item map[string]any) string {
		var t string
		if val, ok := item["type"]; ok {
			t = val.(string)
		}
		switch t {
		case "integer", "id":
			return "int64"
		case "float", "decimal":
			return "float64"
		case "string", "choice", "datetime", "json":
			return "string"
		case "boolean", "bool":
			return "bool"
		case "list":
			return "[]string"
		case "nested object":
			return "map[string]any"
		}
		return t
	},
	"tf2go_primitive_value": func(item map[string]any) string {
		var t string
		if val, ok := item["type"]; ok {
			t = val.(string)
		}
		var post_wrap bool
		if val, ok := item["post_wrap"].(bool); ok {
			post_wrap = val
		}
		switch t {
		case "integer", "id":
			return "ValueInt64"
		case "float", "decimal":
			return "ValueFloat64"
		case "string", "choice", "datetime", "json":
			if post_wrap {
				return "String"
			}
			return "ValueString"
		case "boolean", "bool":
			return "ValueBool"
		case "list":
			return "Elements"
		}
		return t
	},
	"awx2tf_type": func(item map[string]any) string {
		var t string
		if val, ok := item["type"]; ok {
			t = val.(string)
		}
		switch t {
		case "integer", "id":
			return types.Int64Type.String()
		case "float", "decimal":
			return types.Float64Type.String()
		case "string", "choice", "datetime", "json":
			return types.StringType.String()
		case "boolean", "bool":
			return types.BoolType.String()
		case "list":
			return "types.ListType{ElemType: types.StringType}"
		}
		return t
	},
	"awx2go_type": func(item map[string]any) string {
		var t string
		if val, ok := item["type"]; ok {
			t = val.(string)
		}
		switch t {
		case "integer", "id":
			return "types.Int64"
		case "float", "decimal":
			return "types.Float64"
		case "string", "choice", "datetime", "json":
			return "types.String"
		case "boolean", "bool":
			return "types.Bool"
		case "list":
			return "types.List"
		}
		return t
	},
	"awx2go_value": func(item map[string]any) string {
		var t string
		if val, ok := item["type"]; ok {
			t = val.(string)
		}
		switch t {
		case "integer", "id":
			return "types.Int64Value"
		case "float", "decimal":
			return "types.Float64Value"
		case "string", "choice", "datetime", "json":
			return "types.StringValue"
		case "boolean", "bool":
			return "types.BoolValue"
		case "list":
			return "types.ListValueMust(types.StringType, val.Elements())"
		}
		return t
	},
	"awx2go_value_cast": func(item map[string]any) string {
		var t string
		if val, ok := item["type"]; ok {
			t = val.(string)
		}
		switch t {
		case "integer", "id", "float", "decimal":
			return "json.Number"
		case "string", "choice", "datetime", "json":
			return "string"
		case "boolean", "bool":
			return "bool"
		case "list":
			return "types.List"
		}
		return t
	},
	"awx_type_choice_data": func(choices []interface{}) (ret string) {
		var arr []string
		var val interface{}
		for _, choice := range choices {
			val = (choice.([]interface{}))[0]
			switch val := val.(type) {
			case string:
				arr = append(arr, fmt.Sprintf("\"%s\"", val))
			case float64:
				arr = append(arr, fmt.Sprintf("\"%s\"", strconv.FormatInt(int64(val), 10)))
			}
		}

		return fmt.Sprintf("[]string{%s}", strings.Join(arr, ", "))
	},
}
