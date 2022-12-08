package internal

import (
	"encoding/json"
	"os"
)

type PropertyOverride struct {
	Type         string `json:"type"`
	Description  string `json:"description"`
	Sensitive    bool   `json:"sensitive"`
	Required     bool   `json:"required"`
	Trim         bool   `json:"trim"`
	PostWrap     bool   `json:"post_wrap"`
	DefaultValue string `json:"default_value"`
}

type SearchField struct {
	Name           string `json:"name"`
	UrlEscapeValue bool   `json:"url_escape_value,omitempty"`
}

type SearchGroup struct {
	UrlSuffix       string        `json:"url_suffix"`
	Name            string        `json:"name"`
	Fields          []SearchField `json:"fields"`
	MultipleResults bool          `json:"multiple_results,omitempty"`
}

type AssociateDisassociateGroup struct {
	Name          string `json:"name"`
	Endpoint      string `json:"endpoint"`
	Type          string `json:"type"`
	AssociateType string `json:"associate_type"`
}

type Item struct {
	Endpoint                    string                       `json:"endpoint"`
	Name                        string                       `json:"name"`
	TypeName                    string                       `json:"type_name"`
	IdKey                       string                       `json:"id_key"`
	PropertyOverrides           map[string]PropertyOverride  `json:"property_overrides,omitempty"`
	SearchFields                []SearchGroup                `json:"search_fields,omitempty"`
	Enabled                     bool                         `json:"enabled"`
	HasObjectRoles              bool                         `json:"has_object_roles"`
	HasSurveySpec               bool                         `json:"has_survey_spec"`
	AssociateDisassociateGroups []AssociateDisassociateGroup `json:"associate_disassociate_groups"`
	SkipWriteOnly               bool                         `json:"skip_write_only"`
	Undeletable                 bool                         `json:"undeletable"`
	PreStateSetHookFunction     string                       `json:"pre_state_set_hook_function"`
	NoId                        bool                         `json:"no_id"`
	NoTerraformDataSource       bool                         `json:"no_terraform_data_source"`
	NoTerraformResource         bool                         `json:"no_terraform_resource"`
	ApiPropertyResourceKey      string                       `json:"api_property_resource_key"`
	ApiPropertyDataKey          string                       `json:"api_property_data_key"`
	PropertyNameLeaveAsIs       bool                         `json:"property_name_leave_as_is"`
	ApiDataOverride             map[string]map[string]any    `json:"api_data_override"`
	RemoveFieldsDataSource      []string                     `json:"remove_fields_data_source"`
	RemoveFieldsResource        []string                     `json:"remove_fields_resource"`
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
