package internal

import (
	"encoding/json"
	"os"
)

type PropertyOverride struct {
	Type         string   `json:"type" yaml:"type"`
	Description  string   `json:"description" yaml:"description"`
	Sensitive    *bool    `json:"sensitive,omitempty" yaml:"sensitive"`
	Required     *bool    `json:"required,omitempty" yaml:"required"`
	Trim         bool     `json:"trim" yaml:"trim"`
	PostWrap     bool     `json:"post_wrap" yaml:"post_wrap"`
	DefaultValue string   `json:"default_value" yaml:"default_value"`
	ElementType  string   `json:"element_type" yaml:"element_type"`
	Validators   []string `json:"validators" yaml:"validators"`
}

type SearchField struct {
	Name           string `json:"name" yaml:"name"`
	UrlEscapeValue bool   `json:"url_escape_value" yaml:"url_escape_value"`
}

type SearchGroup struct {
	UrlSuffix       string        `json:"url_suffix" yaml:"url_suffix"`
	Name            string        `json:"name" yaml:"name"`
	Fields          []SearchField `json:"fields" yaml:"fields"`
	MultipleResults bool          `json:"multiple_results,omitempty" yaml:"multiple_results"`
}

type AssociateDisassociateGroup struct {
	Name          string `json:"name" yaml:"name"`
	Endpoint      string `json:"endpoint" yaml:"endpoint"`
	Type          string `json:"type" yaml:"type"`
	AssociateType string `json:"associate_type" yaml:"associate_type"`
}

func (a AssociateDisassociateGroup) Map(deprecated bool) map[string]any {
	return map[string]any{
		"Name":          a.Name,
		"Endpoint":      a.Endpoint,
		"Type":          a.Type,
		"AssociateType": a.AssociateType,
		"Deprecated":    deprecated,
	}
}

type FieldConstraint struct {
	Id         string   `json:"id"`
	Constraint string   `json:"constraint"`
	Fields     []string `json:"fields"`
}

type Item struct {
	Endpoint                    string                       `json:"endpoint" yaml:"endpoint"`
	Name                        string                       `json:"name" yaml:"name"`
	TypeName                    string                       `json:"type_name" yaml:"type_name"`
	IdKey                       string                       `json:"id_key" yaml:"id_key"`
	PropertyOverrides           map[string]PropertyOverride  `json:"property_overrides,omitempty" yaml:"property_overrides"`
	SearchFields                []SearchGroup                `json:"search_fields,omitempty" yaml:"search_fields"`
	Enabled                     bool                         `json:"enabled" yaml:"enabled"`
	HasObjectRoles              bool                         `json:"has_object_roles" yaml:"has_object_roles"`
	HasSurveySpec               bool                         `json:"has_survey_spec" yaml:"has_survey_spec"`
	AssociateDisassociateGroups []AssociateDisassociateGroup `json:"associate_disassociate_groups" yaml:"associate_disassociate_groups"`
	FieldConstraints            []FieldConstraint            `json:"field_constraints" yaml:"field_constraints"`
	SkipWriteOnly               bool                         `json:"skip_write_only" yaml:"skip_write_only"`
	Undeletable                 bool                         `json:"undeletable" yaml:"undeletable"`
	PreStateSetHookFunction     string                       `json:"pre_state_set_hook_function" yaml:"pre_state_set_hook_function"`
	NoId                        bool                         `json:"no_id" yaml:"no_id"`
	NoTerraformDataSource       bool                         `json:"no_terraform_data_source" yaml:"no_terraform_data_source"`
	NoTerraformResource         bool                         `json:"no_terraform_resource" yaml:"no_terraform_resource"`
	ApiPropertyResourceKey      string                       `json:"api_property_resource_key" yaml:"api_property_resource_key"`
	ApiPropertyDataKey          string                       `json:"api_property_data_key" yaml:"api_property_data_key"`
	PropertyNameLeaveAsIs       bool                         `json:"property_name_leave_as_is" yaml:"property_name_leave_as_is"`
	ApiDataOverride             map[string]map[string]any    `json:"api_data_override" yaml:"api_data_override"`
	RemoveFieldsDataSource      []string                     `json:"remove_fields_data_source" yaml:"remove_fields_data_source"`
	RemoveFieldsResource        []string                     `json:"remove_fields_resource" yaml:"remove_fields_resource"`
	CredentialTypes             []CredentialTypes            `json:"credential_types" yaml:"credential_types"`
}

type CredentialTypes struct {
	Name         string         `json:"name" mapstructure:"name"`
	Description  string         `json:"description" mapstructure:"description"`
	Organization int            `json:"organization" mapstructure:"organization"`
	Inputs       map[string]any `json:"inputs" mapstructure:"inputs"`
}

type Config struct {
	DefaultRemoveApiDataSource   []string `json:"default_remove_api_data_source"`
	DefaultRemoveApiResource     []string `json:"default_remove_api_resource"`
	Items                        []Item   `json:"items"`
	ApiVersion                   string   `json:"api_version"`
	RenderApiDocs                bool     `json:"render_api_docs"`
	GeneratedApiResources        []string `json:"-"`
	GeneratedDataSourceResources []string `json:"-"`
}

func (c *Config) PackageName(name string) string {
	// return fmt.Sprintf("%s_%s", name, strings.ReplaceAll(c.ApiVersion, ".", "_"))
	return name
}

func (c *Config) Load(filename string) error {
	var payload, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(payload, &c)
	if err != nil {
		return err
	}
	for idx, item := range c.Items {
		if "" == item.ApiPropertyResourceKey {
			c.Items[idx].ApiPropertyResourceKey = "POST"
		}
		if "" == item.ApiPropertyDataKey {
			c.Items[idx].ApiPropertyDataKey = "GET"
		}
	}
	return nil
}

type Deprecated struct {
	Resources   []string
	DataSources []string
	Properties  []DeprecatedProperties
}

type DeprecatedProperties struct {
	Resource        string
	ReadProperties  []string
	WriteProperties []string
}
