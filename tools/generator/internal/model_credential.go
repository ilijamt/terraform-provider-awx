package internal

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/go-viper/mapstructure/v2"
)

type ModelCredential struct {
	Enabled     bool              `json:"enabled" yaml:"enabled"`
	PackageName string            `json:"package_name" yaml:"package_name"`
	ApiVersion  string            `json:"api_version" yaml:"api_version"`
	TypeName    string            `json:"type_name" yaml:"type_name"`
	Id          int64             `json:"id" yaml:"id"`
	Description string            `json:"description" yaml:"description"`
	Name        string            `json:"name" yaml:"name"`
	Namespace   string            `json:"namespace" yaml:"namespace"`
	Required    []string          `json:"required" yaml:"required"`
	IdKey       string            `json:"id_key" yaml:"id_key"`
	Fields      []CredentialField `json:"fields" yaml:"fields"`
	Kind        string            `json:"kind" yaml:"kind"`
	Managed     bool              `json:"managed" yaml:"managed"`
}

type CredentialField struct {
	HelpText     string                   `json:"help_text" yaml:"help_text"`
	Id           string                   `json:"id" yaml:"id"`
	Label        string                   `json:"label" yaml:"label"`
	Type         string                   `json:"type" yaml:"type"`
	Secret       bool                     `json:"secret" yaml:"secret"`
	Multiline    bool                     `json:"multiline" yaml:"multiline"`
	Format       string                   `json:"format" yaml:"format"`
	IsInput      bool                     `json:"is_input" yaml:"is_input"`
	IsUTO        bool                     `json:"is_uto" yaml:"is_uto"`
	AskAtRuntime bool                     `json:"ask_at_runtime" yaml:"ask_at_runtime"`
	Generated    CredentialFieldGenerated `json:"generated" yaml:"generated"`
}

type CredentialFieldGenerated struct {
	Name                   string   `json:"name" yaml:"name"`
	Type                   string   `json:"type" yaml:"type"`
	GoType                 string   `json:"go_type" yaml:"go_type"`
	Required               bool     `json:"required" yaml:"required"`
	Optional               bool     `json:"optional" yaml:"optional"`
	Computed               bool     `json:"computed" yaml:"computed"`
	TerraformValue         string   `json:"terraform_value" yaml:"terraform_value"`
	TerraformAttributeType string   `json:"terraform_attribute_type" yaml:"terraform_attribute_type"`
	ValidatorsOneOf        []string `json:"validators_one_of" yaml:"validators_one_of"`
	WriteOnly              bool     `json:"write_only" yaml:"write_only"`
	Pointer                bool     `json:"pointer" yaml:"pointer"`
}

func (c *ModelCredential) Update(config Config, item Credential, objmap map[string]any) (err error) {
	c.Enabled = item.Enabled
	c.Name = item.Name
	c.TypeName = item.TypeName
	c.Id, _ = objmap["id"].(json.Number).Int64()
	c.IdKey = cmp.Or(item.IdKey, "id")
	c.ApiVersion = config.ApiVersion
	c.PackageName = config.PackageName("awx")
	c.Required = make([]string, 0)
	c.Kind, _ = objmap["kind"].(string)
	c.Managed, _ = objmap["managed"].(bool)

	if val, ok := objmap["namespace"]; ok {
		c.Namespace = val.(string)
	}

	if val, ok := objmap["description"]; ok {
		c.Description = val.(string)
	}

	// Add the default properties defined for all the credentials
	c.Fields = []CredentialField{
		{
			HelpText: "Name of this credential",
			Id:       "name",
			Label:    "Name",
			Type:     "string",
			IsInput:  false,
			Generated: CredentialFieldGenerated{
				Name:                   setPropertyCase("name"),
				Type:                   awxGoType("string"),
				GoType:                 awxPrimitiveType("string"),
				Required:               true,
				TerraformValue:         tfGoPrimitiveValue("string", false),
				TerraformAttributeType: tfAttributeType("string"),
			},
		},
		{
			HelpText: "Description of this credential",
			Id:       "description",
			Label:    "Description",
			Type:     "string",
			IsInput:  false,
			Generated: CredentialFieldGenerated{
				Name:                   setPropertyCase("description"),
				Type:                   awxGoType("string"),
				GoType:                 awxPrimitiveType("string"),
				Optional:               true,
				TerraformValue:         tfGoPrimitiveValue("string", false),
				TerraformAttributeType: tfAttributeType("string"),
			},
		},
		{
			HelpText: "Organization of this credential",
			Id:       "organization",
			Label:    "Organization",
			Type:     "integer",
			IsUTO:    true,
			IsInput:  false,
			Generated: CredentialFieldGenerated{
				Name:                   setPropertyCase("organization"),
				Type:                   awxGoType("integer"),
				GoType:                 awxPrimitiveType("integer"),
				Optional:               true,
				TerraformValue:         tfGoPrimitiveValue("integer", false),
				TerraformAttributeType: tfAttributeType("integer"),
				WriteOnly:              false,
				Pointer:                true,
			},
		},
		// {
		// 	HelpText: "User of this credential, only provided during initial creation of the credential",
		// 	Id:       "user",
		// 	Label:    "User",
		// 	Type:     "integer",
		// 	IsUTO:    true,
		// 	IsInput:  false,
		// 	Generated: CredentialFieldGenerated{
		// 		Name:                   setPropertyCase("user"),
		// 		Type:                   awxGoType("integer"),
		// 		GoType:                 awxPrimitiveType("integer"),
		// 		Optional:               true,
		// 		TerraformValue:         tfGoPrimitiveValue("integer", false),
		// 		TerraformAttributeType: tfAttributeType("integer"),
		// 		WriteOnly:              true,
		// 	},
		// },
		// {
		// 	HelpText: "Team of this credential, only provided during initial creation of the credential",
		// 	Id:       "team",
		// 	Label:    "Team",
		// 	Type:     "integer",
		// 	IsUTO:    true,
		// 	IsInput:  false,
		// 	Generated: CredentialFieldGenerated{
		// 		Name:                   setPropertyCase("team"),
		// 		Type:                   awxGoType("integer"),
		// 		GoType:                 awxPrimitiveType("integer"),
		// 		Optional:               true,
		// 		TerraformValue:         tfGoPrimitiveValue("integer", false),
		// 		TerraformAttributeType: tfAttributeType("integer"),
		// 		WriteOnly:              true,
		// 	},
		// },
	}

	type inputsData struct {
		Fields   []CredentialField `json:"fields" yaml:"fields"`
		Required []string          `json:"required" yaml:"required"`
	}

	var inputs inputsData
	err = mapstructure.Decode(objmap["inputs"], &inputs)
	c.Required = inputs.Required

	for _, field := range inputs.Fields {
		field.Generated.Type = awxGoType(field.Type)
		field.Generated.TerraformValue = tfGoPrimitiveValue(field.Type, false)
		field.Generated.TerraformAttributeType = tfAttributeType(field.Type)
		field.Generated.GoType = awxPrimitiveType(field.Type)
		field.Generated.Name = setPropertyCase(field.Id)
		field.Generated.Required = slices.Contains(c.Required, field.Id)
		field.Generated.Optional = !field.Generated.Required
		field.IsInput = true
		field.IsUTO = false
		c.Fields = append(c.Fields, field)
	}

	return err
}

func (c *ModelCredential) Save(path string) error {
	genDataFile := fmt.Sprintf("%s/%s.json", path, c.TypeName)
	log.Printf("Storing generated data for '%s' in '%s'\n", c.Name, genDataFile)
	payload, _ := json.MarshalIndent(c, "", "  ")
	return os.WriteFile(genDataFile, payload, os.ModePerm)
}
