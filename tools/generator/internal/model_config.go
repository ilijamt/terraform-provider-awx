package internal

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// ModelConfig holds the entire structure's configuration
type ModelConfig struct {
	PackageName     string               `json:"package_name" yaml:"package_name"`         // The name of the package (e.g., "example")
	ApiVersion      string               `json:"api_version" yaml:"api_version"`           // The api version of the API
	Endpoint        string               `json:"endpoint" yaml:"endpoint"`                 // The endpoint for the API request
	Enabled         bool                 `json:"enabled" yaml:"enabled"`                   // Is this model enabled
	Name            string               `json:"name" yaml:"name"`                         // The name of the model (e.g., "ExampleResource")
	ReadProperties  map[string]*Property `json:"read_properties" yaml:"read_properties"`   // All the read properties, regardless of type
	WriteProperties map[string]*Property `json:"write_properties" yaml:"write_properties"` // All the write properties, regardless of type
}

// Property represents a single property in the model
type Property struct {
	Name         string            `json:"name" yaml:"name" mapstructure:"name"`                            // The name of the property (e.g., "id", "name")
	Label        string            `json:"label" yaml:"label" mapstructure:"label"`                         // The label of the property
	Description  string            `json:"description" yaml:"description" mapstructure:"description"`       // The description of the property
	Type         string            `json:"type" yaml:"type" mapstructure:"type"`                            // The type for the property
	DefaultValue string            `json:"default_value" yaml:"default_value" mapstructure:"default_value"` // The default value for the property
	ElementType  string            `json:"element_type" yaml:"element_type" mapstructure:"element_type"`    // The element type of the property
	IsSensitive  bool              `json:"is_sensitive" yaml:"is_sensitive" mapstructure:"is_sensitive"`    // Indicates if the property is sensitive
	IsRequired   bool              `json:"is_required" yaml:"is_required" mapstructure:"is_required"`       // Indicates if the property is required in the schema
	IsWriteOnly  bool              `json:"is_write_only" yaml:"is_write_only" mapstructure:"is_write_only"` // Indicates if the property is write-only (used only in requests)
	IsReadOnly   bool              `json:"is_read_only" yaml:"is_read_only" mapstructure:"is_read_only"`    // Indicates if the property is read-only (used only in responses)
	IsComputed   bool              `json:"is_computed" yaml:"is_computed" mapstructure:"is_computed"`       // Indicates if the property is computed
	IsTypeRead   bool              `json:"is_type_read" yaml:"is_type_read" mapstructure:"is_type_read"`    // Indicates if the property is a read type
	IsTypeWrite  bool              `json:"is_type_write" yaml:"is_type_write" mapstructure:"is_type_write"` // Indicates if the property is a write type
	PostWrap     bool              `json:"post_wrap" yaml:"post_wrap" mapstructure:"post_wrap"`
	Trim         bool              `json:"trim" yaml:"trim" mapstructure:"trim"`
	Generated    PropertyGenerated `json:"generated" yaml:"generated" mapstructure:"generated"`
}

type PropertyGenerated struct {
	AwxGoType             string `json:"awx_go_type" yaml:"awx_go_type"`
	AwxGoValue            string `json:"awx_go_value" yaml:"awx_go_value"`
	PropertyName          string `json:"property_name" yaml:"property_name"`
	BodyRequestModelType  string `json:"body_request_model_type" yaml:"body_request_model_type"`
	TfGoPrimitiveValue    string `json:"tf_go_primitive_value" yaml:"tf_go_primitive_value"`
	ModelBodyRequestValue string `json:"model_body_request_value" yaml:"model_body_request_value"`
	AttributeType         string `json:"attribute_type" yaml:"attribute_type"`
}

type AwxKeyValueType string

const (
	TypeRead  AwxKeyValueType = "read"
	TypeWrite AwxKeyValueType = "write"
)

func (p *Property) Update(vt AwxKeyValueType, override PropertyOverride, values map[string]any, item Item) error {
	p.IsTypeRead = vt == TypeRead
	p.IsTypeWrite = vt == TypeWrite
	p.Trim = override.Trim
	p.PostWrap = override.PostWrap

	p.setWriteOnly(values, override)
	p.setDescription(values, override)
	p.setLabel(values, override)
	p.setType(values, override)
	p.setSensitive(values, override)
	p.setRequired(values, override)
	p.setDefaultValue(values, override)
	p.setElementType(values, override)
	p.setGenerated(values, override, item)

	return nil
}

func (p *Property) setGenerated(values map[string]any, override PropertyOverride, item Item) {
	p.Generated.AwxGoType = awxGoType(p.Type)
	p.Generated.AwxGoValue = awxGoValue(p.Type)
	p.Generated.PropertyName = awxPropertyCase(p.Name, item)
	p.Generated.TfGoPrimitiveValue = tfGoPrimitiveValue(p.Type, p.PostWrap)
	p.Generated.AttributeType = tfAttributeType(p.Type)

	if slices.Contains([]string{"json", "json-yaml"}, p.Type) {
		p.Generated.BodyRequestModelType = "json.RawMessage"
		p.Generated.ModelBodyRequestValue = fmt.Sprintf("json.RawMessage(o.%s.%s())", p.Generated.PropertyName, p.Generated.TfGoPrimitiveValue)
	} else {
		p.Generated.BodyRequestModelType = awxPrimitiveType(p.Type)
		p.Generated.ModelBodyRequestValue = fmt.Sprintf("o.%s.%s()", p.Generated.PropertyName, p.Generated.TfGoPrimitiveValue)
	}
}

func (p *Property) setWriteOnly(values map[string]any, override PropertyOverride) {
	if val, ok := values["write_only"].(bool); ok {
		p.IsWriteOnly = val
	}
}

func (p *Property) setDefaultValue(values map[string]any, override PropertyOverride) {
	if "" != override.DefaultValue {
		values["default"] = override.DefaultValue
	}

	var hasDefault bool
	if _, ok := values["default"]; ok {
		hasDefault = fn_default(values["default"], nil) != nil
	}

	values["computed"] = !p.IsRequired || hasDefault
	p.IsComputed = !p.IsRequired || hasDefault

	if hasDefault {
		values["required"] = false
		p.IsRequired = false
		values["computed"] = true
		p.IsComputed = true

		attrType := tf_attribute_type(values)
		defValue := convertDefaultValue(values["default"])
		switch awx2go_value(values) {
		case "types.StringValue":
			values["default_value"] = fmt.Sprintf("%sdefault.Static%s(`%v`)", lowerCase(attrType), attrType, defValue)
			p.DefaultValue = values["default_value"].(string)
		case "types.Int64Value":
			values["default_value"] = fmt.Sprintf("%sdefault.Static%s(%v)", lowerCase(attrType), attrType, defValue)
			p.DefaultValue = values["default_value"].(string)
		}
	}

}

func (p *Property) setElementType(values map[string]any, override PropertyOverride) {
	if "" != strings.TrimSpace(override.ElementType) {
		values["element_type"] = override.ElementType
	} else if v, err := getItemElementListType(values); err == nil {
		values["element_type"] = v
	}
	if val, ok := values["element_type"].(string); ok {
		p.ElementType = val
	}
}

func (p *Property) setRequired(values map[string]any, override PropertyOverride) {
	if override.Required != nil {
		values["required"] = *override.Required
	}
	if val, ok := values["required"].(bool); ok {
		p.IsRequired = val
	}
	values["required"] = p.IsRequired
}

func (p *Property) setType(values map[string]any, override PropertyOverride) {
	if "" != override.Type {
		values["type"] = override.Type
	}
	if val, ok := values["type"].(string); ok {
		p.Type = val
	}
}

func (p *Property) setSensitive(values map[string]any, override PropertyOverride) {
	if override.Sensitive != nil {
		values["sensitive"] = *override.Sensitive
	}
	if val, ok := values["sensitive"].(bool); ok {
		p.IsSensitive = val
	}
	values["sensitive"] = p.IsSensitive
}

func (p *Property) setLabel(values map[string]any, override PropertyOverride) {
	if v, ok := values["label"].(string); ok {
		p.Label = v
	}

}
func (p *Property) setDescription(values map[string]any, override PropertyOverride) {
	if "" != strings.TrimSpace(override.Description) {
		values["help_text"] = override.Description
	}
	if val, ok := values["help_text"].(string); ok {
		p.Description = val
	}
}

func (c *ModelConfig) UpdateProperty(vt AwxKeyValueType, key string, overrides PropertyOverride, values map[string]any, item Item) (prop *Property, err error) {
	if !slices.Contains([]AwxKeyValueType{TypeRead, TypeWrite}, vt) {
		return prop, fmt.Errorf("unknown property type %q", vt)
	}

	if vt == TypeRead {
		prop = cmp.Or(c.ReadProperties[key], new(Property))
	} else {
		prop = cmp.Or(c.WriteProperties[key], new(Property))
	}

	prop.Name = key
	err = prop.Update(vt, overrides, values, item)

	if vt == TypeRead {
		c.ReadProperties[key] = prop
	} else if vt == TypeWrite {
		if prop.IsWriteOnly && !item.SkipWriteOnly || !prop.IsWriteOnly {
			c.WriteProperties[key] = prop
		}
	}

	return prop, err
}

func (c *ModelConfig) ToMap() (out map[string]any) {
	_ = mapstructure.Decode(c, &out)
	return out
}
