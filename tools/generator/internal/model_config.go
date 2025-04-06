package internal

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// ModelConfig holds the entire structure's configuration
type ModelConfig struct {
	PackageName                 string                       `json:"package_name" yaml:"package_name"`
	ApiVersion                  string                       `json:"api_version" yaml:"api_version"`
	Endpoint                    string                       `json:"endpoint" yaml:"endpoint"`
	TypeName                    string                       `json:"type_name" yaml:"type_name"`
	Description                 string                       `json:"description" yaml:"description"`
	HasObjectRoles              bool                         `json:"has_object_roles" yaml:"has_object_roles"`
	HasSurveySpec               bool                         `json:"has_survey_spec" yaml:"has_survey_spec"`
	RenderApiDocs               bool                         `json:"render_api_docs" yaml:"render_api_docs"`
	NoTerraformDataSource       bool                         `json:"no_terraform_data_source" yaml:"no_terraform_data_source"`
	NoTerraformResource         bool                         `json:"no_terraform_resource" yaml:"no_terraform_resource"`
	HasSearchFields             bool                         `json:"has_search_fields" yaml:"has_search_fields"`
	SearchFields                []SearchGroup                `json:"search_fields,omitempty" yaml:"search_fields"`
	Enabled                     bool                         `json:"enabled" yaml:"enabled"`
	Name                        string                       `json:"name" yaml:"name"`
	NoId                        bool                         `json:"no_id" yaml:"no_id"`
	ReadProperties              map[string]*Property         `json:"read_properties" yaml:"read_properties"`
	WriteProperties             map[string]*Property         `json:"write_properties" yaml:"write_properties"`
	IdProperty                  *Property                    `json:"id_property" yaml:"id_property"`
	IdKey                       string                       `json:"id_key" yaml:"id_key"`
	UnDeletable                 bool                         `json:"un_deletable" yaml:"un_deletable"`
	PreStateSetHookFunction     string                       `json:"pre_state_set_hook_function" yaml:"pre_state_set_hook_function"`
	FieldConstraints            []FieldConstraint            `json:"field_constraints" yaml:"field_constraints" mapstructure:"field_constraints"`
	AssociateDisassociateGroups []AssociateDisassociateGroup `json:"associate_disassociate_groups" yaml:"associate_disassociate_groups"`
	WriteOnlyKeys               []string                     `json:"write_only_keys" yaml:"write_only_keys"`
	Deprecated                  bool                         `json:"deprecated" yaml:"deprecated"`
	DeprecatedParts             map[string]bool              `json:"deprecated_parts" yaml:"deprecated_parts"`
	DeprecatedReadProperties    []string                     `json:"deprecated_read_properties" yaml:"deprecated_read_properties"`
	DeprecatedWriteProperties   []string                     `json:"deprecated_write_properties" yaml:"deprecated_write_properties"`
}

// Property represents a single property in the model
type Property struct {
	IdKey             string            `json:"id_key" yaml:"id_key"`
	Name              string            `json:"name" yaml:"name"`               // The name of the property (e.g., "id", "name")
	Label             string            `json:"label" yaml:"label"`             // The label of the property
	Description       string            `json:"description" yaml:"description"` // The description of the property
	Type              string            `json:"type" yaml:"type"`               // The type for the property
	HasDefaultValue   bool              `json:"has_default_value" yaml:"has_default_value"`
	DefaultValue      string            `json:"default_value" yaml:"default_value"`  // The default value for the property
	ElementType       string            `json:"element_type" yaml:"element_type"`    // The element type of the property
	IsSensitive       bool              `json:"is_sensitive" yaml:"is_sensitive"`    // Indicates if the property is sensitive
	IsRequired        bool              `json:"is_required" yaml:"is_required" `     // Indicates if the property is required in the schema
	IsWriteOnly       bool              `json:"is_write_only" yaml:"is_write_only" ` // Indicates if the property is write-only (used only in requests)
	IsReadOnly        bool              `json:"is_read_only" yaml:"is_read_only"`    // Indicates if the property is read-only (used only in responses)
	IsComputed        bool              `json:"is_computed" yaml:"is_computed" `     // Indicates if the property is computed
	IsTypeRead        bool              `json:"is_type_read" yaml:"is_type_read" `   // Indicates if the property is a read type
	IsTypeWrite       bool              `json:"is_type_write" yaml:"is_type_write" ` // Indicates if the property is a write type
	IsInReadProperty  bool              `json:"is_in_read_property" yaml:"is_in_read_property" `
	IsInWriteProperty bool              `json:"is_in_write_property" yaml:"is_in_write_property" `
	Validators        []string          `json:"validators" yaml:"validators"`
	IsHidden          bool              `json:"is_hidden" yaml:"is_hidden"`
	PostWrap          bool              `json:"post_wrap" yaml:"post_wrap"`
	Trim              bool              `json:"trim" yaml:"trim"`
	IsSearchable      bool              `json:"is_searchable" yaml:"is_searchable"`
	Generated         PropertyGenerated `json:"generated" yaml:"generated"`
	ValidatorData     map[string]any    `json:"validator_data" yaml:"validator_data"`
	Constraints       []FieldConstraint `json:"constraints" yaml:"constraints"`
	Deprecated        bool              `json:"deprecated" yaml:"deprecated"`
}

type PropertyGenerated struct {
	AwxGoType                     string              `json:"awx_go_type" yaml:"awx_go_type"`
	AwxGoValue                    string              `json:"awx_go_value" yaml:"awx_go_value"`
	PropertyName                  string              `json:"property_name" yaml:"property_name"`
	PropertyCase                  string              `json:"property_case" yaml:"property_case"`
	BodyRequestModelType          string              `json:"body_request_model_type" yaml:"body_request_model_type"`
	TfGoPrimitiveValue            string              `json:"tf_go_primitive_value" yaml:"tf_go_primitive_value"`
	ModelBodyRequestValue         string              `json:"model_body_request_value" yaml:"model_body_request_value"`
	AttributeType                 string              `json:"attribute_type" yaml:"attribute_type"`
	ValidationAvailableChoiceData []string            `json:"validation_available_choice_data" yaml:"validation_available_choice_data"`
	AttributeValidationData       map[string][]string `json:"attribute_validation_data" yaml:"attribute_validation_data"`
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
	p.Validators = make([]string, 0)
	p.Generated.ValidationAvailableChoiceData = make([]string, 0)
	p.Generated.AttributeValidationData = make(map[string][]string)

	p.setIdKey(values, override)
	p.setType(values, override)
	p.setWriteOnly(values, override)
	p.setDescription(values, override)
	p.setLabel(values, override)
	p.setSensitive(values, override)
	p.setRequired(values, override)
	p.setElementType(values, override)
	p.setDefaultValue(values, override)
	p.setConstraints(item.FieldConstraints)
	p.setHidden(values)
	p.setValidatorData(values)
	p.setPropertyCustom(values, override)
	p.IsSearchable = fieldIsSearchable(item.SearchFields, p.Name)

	p.Deprecated = strings.Contains(strings.ToLower(p.Description), "this field is deprecated")
	if p.Deprecated {
		p.Description = strings.TrimSpace(strings.ReplaceAll(p.Description, "This field is deprecated and will be removed in a future release.", ""))
	}

	p.setGenerated(values, override, item)

	return nil
}

func (p *Property) setPropertyCustom(values map[string]any, override PropertyOverride) {
	if len(override.Validators) > 0 {
		p.Validators = override.Validators
	}
}

func (p *Property) setValidatorData(values map[string]any) {
	p.ValidatorData = make(map[string]any)

	for _, key := range []string{"max_length", "min_value", "max_value", "choices"} {
		if v, ok := values[key]; ok {
			p.ValidatorData[key] = v
		}
	}

	if v, ok := values["child"].(map[string]any); ok {
		if choices, ok := v["choices"].([]any); ok {
			p.ValidatorData["choices"] = choices
		}
	}
}

func (p *Property) setHidden(values map[string]any) {
	if v, ok := values["hidden"].(bool); ok {
		p.IsHidden = v
	}
}

func (p *Property) setConstraints(constraints []FieldConstraint) {
	p.Constraints = make([]FieldConstraint, 0)
	for _, constraint := range constraints {
		if slices.Contains(constraint.Fields, p.Name) {
			p.Constraints = append(p.Constraints, constraint)
		}
	}
}

func (p *Property) setGenerated(values map[string]any, override PropertyOverride, item Item) {
	p.Generated.AwxGoType = awxGoType(p.Type)
	p.Generated.AwxGoValue = awxGoValue(p.Type)
	p.Generated.PropertyName = awxPropertyCase(p.Name, item)
	p.Generated.PropertyCase = setPropertyCase(p.Name)
	p.Generated.TfGoPrimitiveValue = tfGoPrimitiveValue(p.Type, p.PostWrap)
	p.Generated.AttributeType = tfAttributeType(p.Type)

	if slices.Contains([]string{"json", "json-yaml"}, p.Type) {
		p.Generated.BodyRequestModelType = "json.RawMessage"
		p.Generated.ModelBodyRequestValue = fmt.Sprintf("json.RawMessage(o.%s.%s())", p.Generated.PropertyName, p.Generated.TfGoPrimitiveValue)
	} else {
		p.Generated.BodyRequestModelType = awxPrimitiveType(p.Type)
		p.Generated.ModelBodyRequestValue = fmt.Sprintf("o.%s.%s()", p.Generated.PropertyName, p.Generated.TfGoPrimitiveValue)
	}

	switch p.Type {
	case "choice", "list":
		if v, ok := p.ValidatorData["choices"].([]any); ok {
			p.Generated.ValidationAvailableChoiceData = availableChoicesData(v)
		}
	}

	if p.IsSearchable {
		p.Generated.AttributeValidationData = generateAttributeValidationData(item.SearchFields, p.Name)
	}

}

func (p *Property) setIdKey(values map[string]any, override PropertyOverride) {
	if val, ok := values["id_key"].(string); ok {
		p.IdKey = val
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
		hasDefault = cmp.Or(values["default"], nil) != nil
	}

	values["computed"] = !p.IsRequired || hasDefault
	p.IsComputed = !p.IsRequired || hasDefault

	if hasDefault {
		values["required"] = false
		p.IsRequired = false
		values["computed"] = true
		p.IsComputed = true

		attrType := tfAttributeType(p.Type)
		defValue := convertDefaultValue(values["default"])
		switch awxGoValue(p.Type) {
		case "types.StringValue":
			values["default_value"] = fmt.Sprintf("%sdefault.Static%s(`%v`)", lowerCase(attrType), attrType, defValue)
			p.DefaultValue = values["default_value"].(string)
			p.HasDefaultValue = true
		case "types.Int64Value":
			values["default_value"] = fmt.Sprintf("%sdefault.Static%s(%v)", lowerCase(attrType), attrType, defValue)
			p.DefaultValue = values["default_value"].(string)
			p.HasDefaultValue = true
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

func (c *ModelConfig) Update(config Config, item Item) error {
	c.Endpoint = item.Endpoint
	c.HasObjectRoles = item.HasObjectRoles
	c.HasSurveySpec = item.HasSurveySpec
	c.NoTerraformDataSource = item.NoTerraformDataSource
	c.NoTerraformResource = item.NoTerraformResource
	c.TypeName = item.TypeName
	c.Enabled = item.Enabled
	c.UnDeletable = item.Undeletable
	c.PreStateSetHookFunction = item.PreStateSetHookFunction
	c.PackageName = config.PackageName("awx")
	c.ApiVersion = config.ApiVersion
	c.RenderApiDocs = config.RenderApiDocs
	c.NoId = item.NoId
	c.IdKey = item.IdKey
	c.FieldConstraints = item.FieldConstraints
	c.AssociateDisassociateGroups = item.AssociateDisassociateGroups
	c.DeprecatedReadProperties = make([]string, 0)
	c.DeprecatedWriteProperties = make([]string, 0)
	c.WriteOnlyKeys = make([]string, 0)
	c.FieldConstraints = make([]FieldConstraint, 0)

	if c.AssociateDisassociateGroups == nil {
		c.AssociateDisassociateGroups = make([]AssociateDisassociateGroup, 0)
	}

	if c.ReadProperties == nil {
		c.ReadProperties = make(map[string]*Property)
	}

	if c.WriteProperties == nil {
		c.WriteProperties = make(map[string]*Property)
	}

	c.SearchFields = item.SearchFields
	c.HasSearchFields = len(item.SearchFields) > 0

	c.DeprecatedParts = make(map[string]bool)
	for _, d := range deprecations {
		_ = d.Check(c)
	}

	return nil
}

func (c *ModelConfig) Process(config Config, item Item) (_ error) {
	for key := range c.ReadProperties {
		_, ok := c.WriteProperties[key]
		c.ReadProperties[key].IsInWriteProperty = ok
	}
	for key := range c.WriteProperties {
		_, ok := c.ReadProperties[key]
		c.WriteProperties[key].IsInReadProperty = ok
		if c.WriteProperties[key].IsWriteOnly {
			c.WriteOnlyKeys = append(c.WriteOnlyKeys, key)
		}
	}
	slices.Sort(c.WriteOnlyKeys)
	slices.Sort(c.DeprecatedReadProperties)
	slices.Sort(c.DeprecatedWriteProperties)
	c.IdProperty = c.ReadProperties[c.IdKey]
	return nil
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

	switch vt {
	case TypeRead:
		if prop.Deprecated {
			c.DeprecatedReadProperties = append(c.DeprecatedReadProperties, key)
		}
		c.ReadProperties[key] = prop
	case TypeWrite:
		if prop.Deprecated {
			c.DeprecatedWriteProperties = append(c.DeprecatedWriteProperties, key)
		}
		if prop.IsWriteOnly && !item.SkipWriteOnly || !prop.IsWriteOnly {
			c.WriteProperties[key] = prop
		}
	default:
		panic(fmt.Sprintf("unknown property type %q", vt))
	}

	return prop, err
}

func (c *ModelConfig) ToMap() (out map[string]any) {
	_ = mapstructure.Decode(c, &out)
	return out
}

func (c *ModelConfig) Save(path string) error {
	genDataFile := fmt.Sprintf("%s/%s.json", path, c.TypeName)
	log.Printf("Storing generated data for '%s' in '%s'\n", c.Name, genDataFile)
	payload, _ := json.MarshalIndent(c, "", "  ")
	return os.WriteFile(genDataFile, payload, os.ModePerm)
}
