package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type applicationTerraformModel struct {
	AuthorizationGrantType types.String `tfsdk:"authorization_grant_type" json:"authorization_grant_type"`
	ClientId               types.String `tfsdk:"client_id" json:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret" json:"client_secret"`
	ClientType             types.String `tfsdk:"client_type" json:"client_type"`
	Description            types.String `tfsdk:"description" json:"description"`
	ID                     types.Int64  `tfsdk:"id" json:"id"`
	Name                   types.String `tfsdk:"name" json:"name"`
	Organization           types.Int64  `tfsdk:"organization" json:"organization"`
	RedirectUris           types.String `tfsdk:"redirect_uris" json:"redirect_uris"`
	SkipAuthorization      types.Bool   `tfsdk:"skip_authorization" json:"skip_authorization"`
}

func (o *applicationTerraformModel) Clone() applicationTerraformModel {
	return *o
}

func (o *applicationTerraformModel) BodyRequest() *applicationBodyRequestModel {
	var req applicationBodyRequestModel
	req.AuthorizationGrantType = o.AuthorizationGrantType.ValueString()
	req.ClientType = o.ClientType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.RedirectUris = o.RedirectUris.ValueString()
	req.SkipAuthorization = o.SkipAuthorization.ValueBool()
	return &req
}

func (o *applicationTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.AuthorizationGrantType, data["authorization_grant_type"], false))
	collect(helpers.AttrValueSetString(&o.ClientId, data["client_id"], false))
	collect(helpers.AttrValueSetString(&o.ClientSecret, data["client_secret"], false))
	collect(helpers.AttrValueSetString(&o.ClientType, data["client_type"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.RedirectUris, data["redirect_uris"], false))
	collect(helpers.AttrValueSetBool(&o.SkipAuthorization, data["skip_authorization"]))
	return diags, nil
}

type applicationBodyRequestModel struct {
	AuthorizationGrantType string `json:"authorization_grant_type"`
	ClientType             string `json:"client_type"`
	Description            string `json:"description,omitempty"`
	Name                   string `json:"name"`
	Organization           int64  `json:"organization"`
	RedirectUris           string `json:"redirect_uris,omitempty"`
	SkipAuthorization      bool   `json:"skip_authorization"`
}
