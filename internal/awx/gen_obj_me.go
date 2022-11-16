package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// meTerraformModel maps the schema for Me when using Data Source
type meTerraformModel struct {
	// Email ""
	Email types.String `tfsdk:"email" json:"email"`
	// ExternalAccount "Set if the account is managed by an external service"
	ExternalAccount types.String `tfsdk:"external_account" json:"external_account"`
	// FirstName ""
	FirstName types.String `tfsdk:"first_name" json:"first_name"`
	// ID "Database ID for this user."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// IsSuperuser "Designates that this user has all permissions without explicitly assigning them."
	IsSuperuser types.Bool `tfsdk:"is_superuser" json:"is_superuser"`
	// IsSystemAuditor ""
	IsSystemAuditor types.Bool `tfsdk:"is_system_auditor" json:"is_system_auditor"`
	// LastLogin ""
	LastLogin types.String `tfsdk:"last_login" json:"last_login"`
	// LastName ""
	LastName types.String `tfsdk:"last_name" json:"last_name"`
	// LdapDn ""
	LdapDn types.String `tfsdk:"ldap_dn" json:"ldap_dn"`
	// Username "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
	Username types.String `tfsdk:"username" json:"username"`
}

// Clone the object
func (o meTerraformModel) Clone() meTerraformModel {
	return meTerraformModel{
		Email:           o.Email,
		ExternalAccount: o.ExternalAccount,
		FirstName:       o.FirstName,
		ID:              o.ID,
		IsSuperuser:     o.IsSuperuser,
		IsSystemAuditor: o.IsSystemAuditor,
		LastLogin:       o.LastLogin,
		LastName:        o.LastName,
		LdapDn:          o.LdapDn,
		Username:        o.Username,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Me
func (o meTerraformModel) BodyRequest() (req meBodyRequestModel) {
	return
}

func (o *meTerraformModel) setEmail(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Email, data, false)
}

func (o *meTerraformModel) setExternalAccount(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ExternalAccount, data, false)
}

func (o *meTerraformModel) setFirstName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.FirstName, data, false)
}

func (o *meTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *meTerraformModel) setIsSuperuser(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsSuperuser, data)
}

func (o *meTerraformModel) setIsSystemAuditor(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsSystemAuditor, data)
}

func (o *meTerraformModel) setLastLogin(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LastLogin, data, false)
}

func (o *meTerraformModel) setLastName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LastName, data, false)
}

func (o *meTerraformModel) setLdapDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LdapDn, data, false)
}

func (o *meTerraformModel) setUsername(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *meTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setEmail(data["email"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExternalAccount(data["external_account"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setFirstName(data["first_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsSuperuser(data["is_superuser"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsSystemAuditor(data["is_system_auditor"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastLogin(data["last_login"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastName(data["last_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLdapDn(data["ldap_dn"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUsername(data["username"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// meBodyRequestModel maps the schema for Me for creating and updating the data
type meBodyRequestModel struct {
}

var (
	_ datasource.DataSource              = &meDataSource{}
	_ datasource.DataSourceWithConfigure = &meDataSource{}
)

// NewMeDataSource is a helper function to instantiate the Me data source.
func NewMeDataSource() datasource.DataSource {
	return &meDataSource{}
}

// meDataSource is the data source implementation.
type meDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *meDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/me/"
}

// Metadata returns the data source type name.
func (o *meDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_me"
}

// GetSchema defines the schema for the data source.
func (o *meDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Me",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"email": {
					Description: "Email address",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"external_account": {
					Description: "Set if the account is managed by an external service",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"first_name": {
					Description: "First name",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this user.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"is_superuser": {
					Description: "Designates that this user has all permissions without explicitly assigning them.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"is_system_auditor": {
					Description: "Is system auditor",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"last_login": {
					Description: "Last login",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"last_name": {
					Description: "Last name",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ldap_dn": {
					Description: "Ldap dn",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"username": {
					Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *meDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state meTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for Me
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Me on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Me
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Me on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
