package awx

import "github.com/hashicorp/terraform-plugin-framework/tfsdk"

type fnSchema func(Source, tfsdk.Schema) tfsdk.Schema

var definedSchemaFn = make(map[string]fnSchema)

func processSchema(source Source, target string, schema tfsdk.Schema) tfsdk.Schema {
	if fn, ok := definedSchemaFn[target]; ok {
		return fn(source, schema)
	}
	return schema
}
