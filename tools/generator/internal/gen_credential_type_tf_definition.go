package internal

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// CredentialTypeField describes a single input field on a credential type.
type CredentialTypeField struct {
	ID              string // raw AWX id, e.g. "security_token"
	PropertyName    string // Go field name, e.g. "SecurityToken"
	Label           string
	HelpText        string
	Description     string
	Type            string // "string" or "boolean"
	Secret          bool
	Required        bool
	Multiline       bool
	Format          string // e.g. "ssh_private_key"
	Choices         []string
	RequiresReplace bool
}

func (f CredentialTypeField) IsBool() bool { return f.Type == credentialFieldTypeBoolean }

const credentialFieldTypeBoolean = "boolean"

// CredentialTypeTplData is the template payload for generated typed credential
// resources/data sources/models.
type CredentialTypeTplData struct {
	ApiVersion  string
	PackageName string
	Name        string                // e.g. CredentialAws
	TypeName    string                // e.g. credential_aws
	Endpoint    string                // /api/v2/credentials/
	Namespace   string                // aws
	DisplayName string                // Amazon Web Services
	Kind        string                // cloud / ssh / vault
	Fields      []CredentialTypeField // typed inputs from credential_type spec
	HasSecrets  bool
	Enabled     bool
}

// DataSourceFields drops the secret inputs. AWX answers those with the literal
// "$encrypted$", so a read-only attribute for one returns nothing usable.
func (d CredentialTypeTplData) DataSourceFields() []CredentialTypeField {
	out := make([]CredentialTypeField, 0, len(d.Fields))
	for _, f := range d.Fields {
		if !f.Secret {
			out = append(out, f)
		}
	}
	return out
}

// GenerateCredentialTypeTfDefinition emits the model/resource/data source
// trio for one typed credential type (e.g. awx_credential_aws). The payload is
// the parsed credential_type_<namespace>.json document loaded by
// ApiResources.CredentialTypes.
func GenerateCredentialTypeTfDefinition(tpl *template.Template, config Config, item Item, resourcePath string, payload map[string]any) error {
	if item.CredentialType == "" {
		return fmt.Errorf("credential_type is empty for item %q", item.Name)
	}
	if payload == nil {
		return fmt.Errorf("no credential_type payload found for namespace %q", item.CredentialType)
	}

	log.Printf("Generating typed credential resource for %s (namespace=%s)", item.Name, item.CredentialType)

	data, err := buildCredentialTypeTplData(config, item, payload)
	if err != nil {
		return err
	}

	if !item.Enabled {
		log.Printf("Skipping %s, disabled ...", item.Name)
		return nil
	}

	filename := fmt.Sprintf("%s/gen_obj_%s.go", resourcePath, strings.ToLower(item.TypeName))
	return renderTemplate(tpl, filename, "tf_credential_type.go.tpl", data)
}

func buildCredentialTypeTplData(config Config, item Item, payload map[string]any) (*CredentialTypeTplData, error) {
	out := &CredentialTypeTplData{
		ApiVersion:  config.ApiVersion,
		PackageName: config.PackageName("awx"),
		Name:        item.Name,
		TypeName:    item.TypeName,
		Endpoint:    item.Endpoint,
		Namespace:   item.CredentialType,
		Enabled:     item.Enabled,
	}
	if v, ok := payload["name"].(string); ok {
		out.DisplayName = v
	}
	if v, ok := payload["kind"].(string); ok {
		out.Kind = v
	}

	inputs, ok := payload["inputs"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("credential_type %q has no inputs object", item.CredentialType)
	}

	requiredSet := map[string]bool{}
	if reqList, ok := inputs["required"].([]any); ok {
		for _, r := range reqList {
			if s, ok := r.(string); ok {
				requiredSet[s] = true
			}
		}
	}

	fieldsAny, ok := inputs["fields"].([]any)
	if !ok {
		return nil, fmt.Errorf("credential_type %q has no inputs.fields array", item.CredentialType)
	}

	fields := make([]CredentialTypeField, 0, len(fieldsAny))
	for _, f := range fieldsAny {
		fm, ok := f.(map[string]any)
		if !ok {
			continue
		}
		id, _ := fm["id"].(string)
		if id == "" {
			continue
		}

		ft := CredentialTypeField{
			ID:              id,
			PropertyName:    strcase.ToCamel(id),
			Required:        requiredSet[id],
			Type:            "string",
			RequiresReplace: slices.Contains(item.CredentialRequiresReplace, id),
		}
		if v, ok := fm["type"].(string); ok && v != "" {
			ft.Type = v
		}
		if v, ok := fm["label"].(string); ok {
			ft.Label = v
		}
		if v, ok := fm["help_text"].(string); ok {
			ft.HelpText = v
		}
		if v, ok := fm["secret"].(bool); ok {
			ft.Secret = v
		}
		if v, ok := fm["multiline"].(bool); ok {
			ft.Multiline = v
		}
		if v, ok := fm["format"].(string); ok {
			ft.Format = v
		}
		if choices, ok := fm["choices"].([]any); ok {
			for _, c := range choices {
				if s, ok := c.(string); ok {
					ft.Choices = append(ft.Choices, s)
				}
			}
		}
		ft.Description = credentialFieldDescription(ft, fm["default"])
		if ft.Secret {
			out.HasSecrets = true
		}
		fields = append(fields, ft)
	}

	// Stable order — sorting by ID keeps generated diffs deterministic across
	// runs even if AWX reshuffles its inputs array.
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].ID < fields[j].ID })
	out.Fields = fields

	return out, nil
}

// credentialFieldDescription records the AWX default rather than applying it,
// because AWX never echoes a field it wasn't sent. It also records choices and
// replace-on-change, which tfplugindocs cannot read off the validator and plan
// modifier that enforce them.
func credentialFieldDescription(f CredentialTypeField, def any) string {
	lead := f.HelpText
	if lead == "" {
		lead = f.Label
	}
	if lead != "" && !strings.HasSuffix(lead, ".") {
		lead += "."
	}

	parts := []string{lead}
	if len(f.Choices) > 0 {
		quoted := make([]string, 0, len(f.Choices))
		for _, c := range f.Choices {
			quoted = append(quoted, strconv.Quote(c))
		}
		parts = append(parts, fmt.Sprintf("Allowed values: %s.", strings.Join(quoted, ", ")))
	}
	switch v := def.(type) {
	case string:
		if v != "" {
			parts = append(parts, fmt.Sprintf("AWX defaults this to %q when unset.", v))
		}
	case bool:
		parts = append(parts, fmt.Sprintf("AWX defaults this to %q when unset.", strconv.FormatBool(v)))
	}
	if f.RequiresReplace {
		parts = append(parts, "Changing this forces a new credential to be created.")
	}
	return strings.TrimSpace(strings.Join(parts, " "))
}

// IsCredentialTypeItem reports whether an item should be routed through
// GenerateCredentialTypeTfDefinition rather than the standard pipeline.
func IsCredentialTypeItem(item Item) bool {
	return strings.TrimSpace(item.CredentialType) != ""
}

// CredentialTypeNamespaces returns sorted enabled credential-type namespaces
// from the config — handy for sanity-checking config loads in tests.
func CredentialTypeNamespaces(items []Item) []string {
	out := make([]string, 0)
	for _, it := range items {
		if IsCredentialTypeItem(it) && it.Enabled {
			out = append(out, it.CredentialType)
		}
	}
	slices.Sort(out)
	return out
}
