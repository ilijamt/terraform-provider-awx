package internal

import (
	"encoding/json"
	"os"
	"strings"
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
	// OmitEmpty controls whether the body-request struct field gets a
	// `,omitempty` JSON tag. Default is the legacy behavior (omit on zero
	// for non-required, non-bool fields). Set to false on int64/float64
	// fields where 0 is a meaningful value (e.g. AWX "0 = no limit"
	// settings), otherwise the zero gets stripped from the body and the
	// server keeps whatever it held, causing "Provider produced inconsistent
	// result after apply".
	//
	// Launch prompts are the exception: forks, job_slice_count and timeout on
	// schedule and workflow_job_template_node have to drop out of the body
	// when unset, for the same reason diff_mode is Nullable there.
	OmitEmpty *bool `json:"omit_empty,omitempty" yaml:"omit_empty,omitempty"`
	// Nullable makes the body-request field a pointer so an unset attribute is
	// dropped rather than sent as its zero value. Set on bool launch prompts
	// (schedule diff_mode): AWX rejects those with "Field is not allowed on
	// launch" when the unified job template does not prompt for them.
	Nullable *bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	// UseStateForUnknown defaults to true. Set false for values AWX recomputes
	// server-side, such as schedule next_run derived from rrule. Promising the
	// prior state value there yields "Provider produced inconsistent result
	// after apply".
	UseStateForUnknown *bool `json:"use_state_for_unknown,omitempty" yaml:"use_state_for_unknown,omitempty"`
	// NoDefault discards a default reported by AWX rather than pinning it into
	// the schema, leaving the attribute Optional+Computed so the server supplies
	// the value. AWX renders a callable model default as one fixed sample, so a
	// uuid4 default arrives as the literal "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	// (awx/api/metadata.py). Pinning it hands every resource the same value,
	// which collides with the per-workflow unique constraint on a node
	// identifier. Also set on values that describe the source deployment rather
	// than AWX itself, such as the k8s namespace in an instance group pod spec.
	NoDefault bool `json:"no_default,omitempty" yaml:"no_default,omitempty"`
	// RequiresReplace marks an attribute AWX accepts on create but ignores on
	// update, so Terraform has to recreate instead of PATCH. The node serializer
	// turns workflow_job_template read_only once the instance exists; without
	// this the PATCH returns the old value and Terraform reports "Provider
	// produced inconsistent result after apply".
	RequiresReplace bool `json:"requires_replace,omitempty" yaml:"requires_replace,omitempty"`
}

// MetadataDiscovery resolves the id that fills the %d in Item.MetadataEndpoint
// by reading Field off the first result of Endpoint. Field defaults to "id".
type MetadataDiscovery struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
	Field    string `json:"field,omitempty" yaml:"field,omitempty"`
}

// CreateEndpointConfig describes a POST target that differs from Item.Endpoint.
// Endpoint is a Sprintf pattern taking the id held by IdAttribute, a schema
// attribute the API never reports and so has to be injected through
// ApiDataOverrideResource.
type CreateEndpointConfig struct {
	Endpoint    string `json:"endpoint" yaml:"endpoint"`
	IdAttribute string `json:"id_attribute" yaml:"id_attribute"`
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
	NoImport                    bool                         `json:"no_import" yaml:"no_import"`
	NoTerraformDataSource       bool                         `json:"no_terraform_data_source" yaml:"no_terraform_data_source"`
	NoTerraformResource         bool                         `json:"no_terraform_resource" yaml:"no_terraform_resource"`
	ApiPropertyResourceKey      string                       `json:"api_property_resource_key" yaml:"api_property_resource_key"`
	ApiPropertyDataKey          string                       `json:"api_property_data_key" yaml:"api_property_data_key"`
	PropertyNameLeaveAsIs       bool                         `json:"property_name_leave_as_is" yaml:"property_name_leave_as_is"`
	ApiDataOverride             map[string]map[string]any    `json:"api_data_override" yaml:"api_data_override"`
	ApiDataOverrideResource     map[string]map[string]any    `json:"api_data_override_resource" yaml:"api_data_override_resource"`
	RemoveFieldsDataSource      []string                     `json:"remove_fields_data_source" yaml:"remove_fields_data_source"`
	RemoveFieldsResource        []string                     `json:"remove_fields_resource" yaml:"remove_fields_resource"`
	CredentialTypes             []CredentialTypes            `json:"credential_types" yaml:"credential_types"`
	WaitLifecycle               *WaitLifecycleConfig         `json:"wait_lifecycle,omitempty" yaml:"wait_lifecycle,omitempty"`

	// MetadataEndpoint overrides the URL fetch-api-resources runs OPTIONS
	// against, for a model AWX only describes from an instance URL.
	// WorkflowApprovalTemplate has no list view, so its collection 404s and only
	// /api/v2/workflow_approval_templates/{id}/ answers. A %d is filled from
	// MetadataDiscovery.
	MetadataEndpoint string `json:"metadata_endpoint,omitempty" yaml:"metadata_endpoint,omitempty"`

	MetadataDiscovery *MetadataDiscovery `json:"metadata_discovery,omitempty" yaml:"metadata_discovery,omitempty"`

	CreateEndpoint *CreateEndpointConfig `json:"create_endpoint,omitempty" yaml:"create_endpoint,omitempty"`

	// UseStateForUnknown sets the default for every property on this item, and
	// defaults to true. Turn it off for a resource whose reads are live
	// telemetry: an instance recomputes capacity on any change, so promising the
	// prior value fails the apply.
	UseStateForUnknown *bool `json:"use_state_for_unknown,omitempty" yaml:"use_state_for_unknown,omitempty"`

	// RequiresReplace marks every write property on this item, for an object AWX
	// creates and deletes but never updates. A role assignment detail answers
	// only GET and DELETE, so any change has to recreate.
	RequiresReplace bool `json:"requires_replace,omitempty" yaml:"requires_replace,omitempty"`

	// SoftDelete replaces DELETE with a PATCH carrying this body. AWX retires a
	// few objects through a state field rather than removing them: DELETE on an
	// instance answers 405, and node_state "deprovisioning" is what takes it out.
	SoftDelete map[string]any `json:"soft_delete,omitempty" yaml:"soft_delete,omitempty"`

	// CredentialRequiresReplace lists input field ids AWX refuses to change
	// after create. A vault credential answers a changed vault_id with "Vault
	// IDs cannot be changed once they have been created."
	CredentialRequiresReplace []string `json:"credential_requires_replace,omitempty" yaml:"credential_requires_replace,omitempty"`

	// CredentialType, when non-empty, marks this item as a typed credential
	// resource generated from resources/api/<VERSION>/payload/credential_type_<value>.json
	// rather than from the regular API actions metadata. The value is the
	// AWX namespace ("aws", "ssh", "vault", ...). The generator routes these
	// items through GenerateCredentialTypeTfDefinition.
	CredentialType string `json:"credential_type,omitempty" yaml:"credential_type,omitempty"`

	// NotificationType, when non-empty, marks this item as a typed notification
	// template generated from the per-type schema AWX embeds in the
	// NotificationTemplate OPTIONS payload under
	// actions.POST.notification_configuration.<value>, rather than from this
	// item's own actions metadata.
	NotificationType string `json:"notification_type,omitempty" yaml:"notification_type,omitempty"`
}

// MetadataUrl returns the URL OPTIONS runs against for this item, and whether an
// id still has to be discovered to fill it.
func (i Item) MetadataUrl() (url string, needsDiscovery bool) {
	if i.MetadataEndpoint == "" {
		return i.Endpoint, false
	}
	return i.MetadataEndpoint, strings.Contains(i.MetadataEndpoint, "%d")
}

// WaitLifecycleConfig opts a resource into post-Create/Update polling. The
// generator emits a Terraform-only bool toggle (WaitAttribute), a timeouts
// block, and the WaitLifecycle wiring on the generated resource so the
// framework polls AWX until the resource reaches a terminal status.
type WaitLifecycleConfig struct {
	// WaitAttribute is the schema attribute name (e.g. "wait_for_sync").
	WaitAttribute string `json:"wait_attribute" yaml:"wait_attribute"`
	// WaitDescription is the schema attribute description.
	WaitDescription string `json:"wait_description" yaml:"wait_description"`
	// EndpointSuffix is appended to the resource's base endpoint with the ID
	// substituted via Sprintf — typically "%d/" for AWX.
	EndpointSuffix string `json:"endpoint_suffix" yaml:"endpoint_suffix"`
	// StatusField is the JSON field on the polled response to inspect.
	StatusField string `json:"status_field" yaml:"status_field"`
	// SuccessValues are terminal field values that mean the wait succeeded.
	SuccessValues []string `json:"success_values" yaml:"success_values"`
	// FailureValues are terminal field values that mean the wait failed.
	FailureValues []string `json:"failure_values" yaml:"failure_values"`
	// DefaultTimeout is the fallback if the user doesn't set timeouts. Go duration string.
	DefaultTimeout string `json:"default_timeout" yaml:"default_timeout"`
	// PollInterval between status reads. Go duration string.
	PollInterval string `json:"poll_interval" yaml:"poll_interval"`
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
		if item.ApiPropertyResourceKey == "" {
			c.Items[idx].ApiPropertyResourceKey = "POST"
		}
		if item.ApiPropertyDataKey == "" {
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
