package awx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// credentialOpenstackTerraformModel exposes the typed AWX OpenStack
// credential (credential_openstack) inputs as first-class schema attributes rather
// than an opaque JSON blob.
type credentialOpenstackTerraformModel struct {
	ID                types.Int64  `tfsdk:"id" json:"id"`
	Name              types.String `tfsdk:"name" json:"name"`
	Description       types.String `tfsdk:"description" json:"description"`
	Organization      types.Int64  `tfsdk:"organization" json:"organization"`
	Team              types.Int64  `tfsdk:"team" json:"team"`
	User              types.Int64  `tfsdk:"user" json:"user"`
	Kind              types.String `tfsdk:"kind" json:"kind"`
	Managed           types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType    types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Domain            types.String `tfsdk:"domain" json:"-"`
	Host              types.String `tfsdk:"host" json:"-"`
	Password          types.String `tfsdk:"password" json:"-"`
	Project           types.String `tfsdk:"project" json:"-"`
	ProjectDomainName types.String `tfsdk:"project_domain_name" json:"-"`
	Region            types.String `tfsdk:"region" json:"-"`
	Username          types.String `tfsdk:"username" json:"-"`
	VerifySsl         types.Bool   `tfsdk:"verify_ssl" json:"-"`
}

func (o *credentialOpenstackTerraformModel) Clone() credentialOpenstackTerraformModel {
	return *o
}

type credentialOpenstackBodyRequestModel struct {
	CredentialType int64           `json:"credential_type"`
	Description    string          `json:"description,omitempty"`
	Inputs         json.RawMessage `json:"inputs,omitempty"`
	Name           string          `json:"name"`
	Organization   int64           `json:"organization,omitempty"`
	Team           int64           `json:"team,omitempty"`
	User           int64           `json:"user,omitempty"`
}

// BodyRequest folds typed input fields back into a single `inputs` JSON object;
// null/unknown values are dropped so the API doesn't receive empty strings for
// unset optionals.
func (o *credentialOpenstackTerraformModel) BodyRequest() *credentialOpenstackBodyRequestModel {
	req := &credentialOpenstackBodyRequestModel{
		CredentialType: o.CredentialType.ValueInt64(),
		Description:    o.Description.ValueString(),
		Name:           o.Name.ValueString(),
		Organization:   o.Organization.ValueInt64(),
	}

	inputs := map[string]any{}
	if !o.Domain.IsNull() && !o.Domain.IsUnknown() {
		inputs["domain"] = o.Domain.ValueString()
	}
	if !o.Host.IsNull() && !o.Host.IsUnknown() {
		inputs["host"] = o.Host.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.Project.IsNull() && !o.Project.IsUnknown() {
		inputs["project"] = o.Project.ValueString()
	}
	if !o.ProjectDomainName.IsNull() && !o.ProjectDomainName.IsUnknown() {
		inputs["project_domain_name"] = o.ProjectDomainName.ValueString()
	}
	if !o.Region.IsNull() && !o.Region.IsUnknown() {
		inputs["region"] = o.Region.ValueString()
	}
	if !o.Username.IsNull() && !o.Username.IsUnknown() {
		inputs["username"] = o.Username.ValueString()
	}
	if !o.VerifySsl.IsNull() && !o.VerifySsl.IsUnknown() {
		inputs["verify_ssl"] = o.VerifySsl.ValueBool()
	}
	if len(inputs) > 0 {
		payload, _ := json.Marshal(inputs)
		req.Inputs = payload
	}
	return req
}

// UpdateFromApiData unfolds the AWX response back into the typed model. Secret
// fields come back as `$encrypted$` placeholders; the per-credential-type
// pre-state-set hook reconciles them against prior plan state.
func (o *credentialOpenstackTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))

	if inputs, ok := data["inputs"].(map[string]any); ok {
		collect(helpers.AttrValueSetString(&o.Domain, inputs["domain"], false))
		collect(helpers.AttrValueSetString(&o.Host, inputs["host"], false))
		collect(helpers.AttrValueSetString(&o.Password, inputs["password"], false))
		collect(helpers.AttrValueSetString(&o.Project, inputs["project"], false))
		collect(helpers.AttrValueSetString(&o.ProjectDomainName, inputs["project_domain_name"], false))
		collect(helpers.AttrValueSetString(&o.Region, inputs["region"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
		collect(helpers.AttrValueSetBool(&o.VerifySsl, inputs["verify_ssl"]))
	}
	return diags, nil
}

// hookCredentialOpenstack reconciles the `$encrypted$` placeholders AWX returns for
// secret fields against the prior plan state, so Terraform doesn't see drift
// every plan.
func hookCredentialOpenstack(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *credentialOpenstackTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// Secrets aren't echoed by AWX in plain form. Carry the planned value
		// forward; force a known null when the user didn't set the field.
		if orig.Password.IsNull() || orig.Password.IsUnknown() {
			state.Password = types.StringNull()
		} else {
			state.Password = orig.Password
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
	}
	return nil
}

// credentialOpenstackDataSourceTerraformModel mirrors the resource
// model without the secret inputs, which AWX only ever answers with the
// literal "$encrypted$".
type credentialOpenstackDataSourceTerraformModel struct {
	ID                types.Int64  `tfsdk:"id" json:"id"`
	Name              types.String `tfsdk:"name" json:"name"`
	Description       types.String `tfsdk:"description" json:"description"`
	Organization      types.Int64  `tfsdk:"organization" json:"organization"`
	Team              types.Int64  `tfsdk:"team" json:"team"`
	User              types.Int64  `tfsdk:"user" json:"user"`
	Kind              types.String `tfsdk:"kind" json:"kind"`
	Managed           types.Bool   `tfsdk:"managed" json:"managed"`
	CredentialType    types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Domain            types.String `tfsdk:"domain" json:"-"`
	Host              types.String `tfsdk:"host" json:"-"`
	Project           types.String `tfsdk:"project" json:"-"`
	ProjectDomainName types.String `tfsdk:"project_domain_name" json:"-"`
	Region            types.String `tfsdk:"region" json:"-"`
	Username          types.String `tfsdk:"username" json:"-"`
	VerifySsl         types.Bool   `tfsdk:"verify_ssl" json:"-"`
}

func (o *credentialOpenstackDataSourceTerraformModel) Clone() credentialOpenstackDataSourceTerraformModel {
	return *o
}

func (o *credentialOpenstackDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))

	if inputs, ok := data["inputs"].(map[string]any); ok {
		collect(helpers.AttrValueSetString(&o.Domain, inputs["domain"], false))
		collect(helpers.AttrValueSetString(&o.Host, inputs["host"], false))
		collect(helpers.AttrValueSetString(&o.Project, inputs["project"], false))
		collect(helpers.AttrValueSetString(&o.ProjectDomainName, inputs["project_domain_name"], false))
		collect(helpers.AttrValueSetString(&o.Region, inputs["region"], false))
		collect(helpers.AttrValueSetString(&o.Username, inputs["username"], false))
		collect(helpers.AttrValueSetBool(&o.VerifySsl, inputs["verify_ssl"]))
	}
	return diags, nil
}

// credentialOpenstackTypeLookup is shared between the resource and
// data source so a single namespace lookup at Configure time covers both.
var credentialOpenstackTypeLookup = framework.NewCredentialTypeLookup()

type credentialOpenstackResource = framework.GenericResource[credentialOpenstackTerraformModel, credentialOpenstackBodyRequestModel, *credentialOpenstackTerraformModel]

// NewCredentialOpenstackResource constructs the typed OpenStack credential resource.
// The credential_type ID is resolved by namespace (openstack) at Configure
// time so the resource works against any AWX instance regardless of how the
// managed credential type is numbered locally.
func NewCredentialOpenstackResource() resource.Resource {
	attrs := framework.CredentialBaseResourceAttrs()
	attrs["domain"] = schema.StringAttribute{
		Description: "OpenStack domains define administrative boundaries. It is only needed for Keystone v3 authentication URLs. Refer to the documentation for common scenarios.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["host"] = schema.StringAttribute{
		Description: "The host to authenticate with.  For example, https://openstack.business.com/v2.0/.",
		Required:    true,
	}
	attrs["password"] = schema.StringAttribute{
		Description: "Password (API Key).",
		Required:    true,
		Sensitive:   true,
	}
	attrs["project"] = schema.StringAttribute{
		Description: "Project (Tenant Name).",
		Required:    true,
	}
	attrs["project_domain_name"] = schema.StringAttribute{
		Description: "Project (Domain Name).",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["region"] = schema.StringAttribute{
		Description: "For some cloud providers, like OVH, region must be specified.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Required:    true,
	}
	attrs["verify_ssl"] = schema.BoolAttribute{
		Description: "Verify SSL. AWX defaults this to \"true\" when unset.",
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	}
	return &credentialOpenstackResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_openstack", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialOpenstackTerraformModel, credentialOpenstackBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages the AWX `OpenStack` (openstack) credential type with first-class typed input attributes. Equivalent to `awx_credential` with `credential_type = data.awx_credential_type.openstack.id`, but with per-field schema validation and sensitivity.",
				Attributes:          attrs,
			},
			IDAccessor:  func(m *credentialOpenstackTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:       "id",
			Hook:        hookCredentialOpenstack,
			OnConfigure: credentialOpenstackTypeLookup.OnConfigure("openstack"),
			MutateBody: func(plan *credentialOpenstackTerraformModel, body *credentialOpenstackBodyRequestModel) {
				body.CredentialType = credentialOpenstackTypeLookup.Load()
			},
			WriteOnlyPlanToBody: func(plan *credentialOpenstackTerraformModel, body *credentialOpenstackBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialOpenstackTerraformModel) {
				// AWX never echoes team/user back. An unset owner has to stay
				// null: an imported credential plans it as null, and writing 0
				// there fails the apply.
				state.Team = helpers.KnownOrNullInt64(plan.Team)
				state.User = helpers.KnownOrNullInt64(plan.User)
				if state.CredentialType.IsNull() || state.CredentialType.IsUnknown() {
					state.CredentialType = types.Int64Value(credentialOpenstackTypeLookup.Load())
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialOpenstack",
		},
	}
}

type credentialOpenstackDataSource = framework.GenericDataSource[credentialOpenstackDataSourceTerraformModel, *credentialOpenstackDataSourceTerraformModel]

// NewCredentialOpenstackDataSource constructs the typed OpenStack credential data source.
func NewCredentialOpenstackDataSource() datasource.DataSource {
	attrs := framework.CredentialBaseDataSourceAttrs()
	attrs["domain"] = dschema.StringAttribute{
		Description: "OpenStack domains define administrative boundaries. It is only needed for Keystone v3 authentication URLs. Refer to the documentation for common scenarios.",
		Computed:    true,
	}
	attrs["host"] = dschema.StringAttribute{
		Description: "The host to authenticate with.  For example, https://openstack.business.com/v2.0/.",
		Computed:    true,
	}
	attrs["project"] = dschema.StringAttribute{
		Description: "Project (Tenant Name).",
		Computed:    true,
	}
	attrs["project_domain_name"] = dschema.StringAttribute{
		Description: "Project (Domain Name).",
		Computed:    true,
	}
	attrs["region"] = dschema.StringAttribute{
		Description: "For some cloud providers, like OVH, region must be specified.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	attrs["verify_ssl"] = dschema.BoolAttribute{
		Description: "Verify SSL. AWX defaults this to \"true\" when unset.",
		Computed:    true,
	}
	return &credentialOpenstackDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_openstack", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialOpenstackDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `OpenStack` (openstack) credential by ID or name.",
				Attributes:          attrs,
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "/?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			OnConfigure:  credentialOpenstackTypeLookup.OnConfigure("openstack"),
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialOpenstack",
		},
	}
}
