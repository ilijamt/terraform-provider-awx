package internal

import (
	"fmt"

	"github.com/iancoleman/strcase"
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

func tf2GoPrimitiveValue(t string, postWrap bool) string {
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
