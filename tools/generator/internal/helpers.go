package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	p "path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

func generateAttributeValidationData(fields []SearchGroup, needle string) (attrs map[string][]string) {
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
}

func availableChoicesData(choices []any) (ret []string) {
	var arr []string
	var val any
	for _, choice := range choices {
		val = (choice.([]any))[0]
		switch val := val.(type) {
		case string:
			arr = append(arr, val)
		case float64:
			arr = append(arr, strconv.FormatInt(int64(val), 10))
		}
	}

	return arr
}

func fieldIsSearchable(fields []SearchGroup, needle string) bool {
	for _, field := range fields {
		for _, v := range field.Fields {
			if v.Name == needle {
				return true
			}
		}
	}
	return false
}

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

func awxGoValue(t string) string {
	switch t {
	case "integer", "id":
		return "types.Int64Value"
	case "float", "decimal":
		return "types.Float64Value"
	case "string", "choice", "datetime", "json", "json-yaml":
		return "types.StringValue"
	case "boolean", "bool":
		return "types.BoolValue"
	case "list":
		return "types.ListValueMust(types.StringType, val.Elements())"
	}
	return t
}

func awxGoType(t string) string {
	switch t {
	case "integer", "id":
		return "types.Int64"
	case "float", "decimal":
		return "types.Float64"
	case "string", "choice", "datetime", "json", "json-yaml":
		return "types.String"
	case "boolean", "bool":
		return "types.Bool"
	case "list":
		return "types.List"
	}
	return t
}

func awxPropertyCase(in string, config Item) string {
	if config.PropertyNameLeaveAsIs {
		return in
	}
	return strcase.ToCamel(in)
}

func awxPrimitiveType(t string) string {
	switch t {
	case "integer", "id":
		return "int64"
	case "float", "decimal":
		return "float64"
	case "string", "choice", "datetime", "json", "json-yaml":
		return "string"
	case "boolean", "bool":
		return "bool"
	case "list":
		return "[]string"
	case "nested object":
		return "map[string]any"
	}
	return t
}

func tfAttributeType(t string) string {
	switch t {
	case "integer", "id":
		return "Int64"
	case "float", "decimal":
		return "Float64"
	case "string", "choice", "datetime", "json", "json-yaml":
		return "String"
	case "boolean", "bool":
		return "Bool"
	case "list":
		return "List"
	}
	return t
}

func tfGoPrimitiveValue(t string, postWrap bool) string {
	switch t {
	case "integer", "id":
		return "ValueInt64"
	case "float", "decimal":
		return "ValueFloat64"
	case "string", "choice", "datetime", "json", "json-yaml":
		if postWrap {
			return "String"
		}
		return "ValueString"
	case "boolean", "bool":
		return "ValueBool"
	case "list":
		return "Elements"
	}
	return t
}

func setPropertyCase(in string) string {
	return strcase.ToCamel(strcase.ToSnake(in))
}

func init() {
	strcase.ConfigureAcronym("id", "ID")
}

func renderTemplate(tpl *template.Template, filename, template string, data any) (err error) {
	log.Printf("Rendering of %s into %s.", template, filename)
	var f *os.File
	resourcePathDir := filepath.Dir(filename)
	if err = os.MkdirAll(resourcePathDir, os.ModePerm); err != nil {
		return err
	}

	if f, err = os.Create(filename); err != nil {
		return err
	}

	if err = tpl.ExecuteTemplate(f, template, data); err != nil {
		return err
	}
	return nil
}

var lowerCase = func(in string) string {
	return strings.ToLower(in)
}

var PascalCase = func(s string) string {
	// Split the string by non-alphanumeric characters
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Capitalize the first letter of each word
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			// Make the rest of the letters lowercase
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}

	// Join the words back together
	return strings.Join(words, "")
}

var convertDefaultValue = func(in any) any {
	switch in.(type) {
	case map[string]any:
		payload, err := json.Marshal(in)
		if err != nil {
			return "{}"
		}
		return string(payload)
	}
	return in
}

var FuncMap = template.FuncMap{
	"format_number": func(in any) any {
		switch v := in.(type) {
		case float32:
			return int64(v)
		case float64:
			return int64(v)
		}
		return in
	},
	"testTfValue": func(t any, name string) any {
		switch t {
		case "integer", "id":
			return "types.Int64Value(1)"
		case "boolean", "bool":
			return "types.BoolValue(true)"
		}
		return fmt.Sprintf("types.StringValue(%q)", name)
	},
	"url_path_clean": func(in string) string {
		return p.Clean(in)
	},
	"escape_quotes": func(in string) string {
		return fmt.Sprintf("%q", in)
	},
	"pascalCase":     PascalCase,
	"snakeCase":      strcase.ToSnake,
	"camelCase":      strcase.ToCamel,
	"lowerCase":      lowerCase,
	"lowerCamelCase": strcase.ToLowerCamel,
	"hasKey": func(d map[string]any, key string) bool {
		_, ok := d[key]
		return ok
	},
	"quote": func(in any) string {
		return fmt.Sprintf("%q", in)
	},
	"toJson": func(in any) string {
		payload, _ := json.MarshalIndent(in, "", "  ")
		return string(payload)
	},
	"toYaml": func(in any) string {
		payload, _ := yaml.Marshal(in)
		return string(payload)
	},
}
