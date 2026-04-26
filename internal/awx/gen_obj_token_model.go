package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type tokensTerraformModel struct {
	Application  types.Int64  `tfsdk:"application" json:"application"`
	Description  types.String `tfsdk:"description" json:"description"`
	Expires      types.String `tfsdk:"expires" json:"expires"`
	ID           types.Int64  `tfsdk:"id" json:"id"`
	RefreshToken types.String `tfsdk:"refresh_token" json:"refresh_token"`
	Scope        types.String `tfsdk:"scope" json:"scope"`
	Token        types.String `tfsdk:"token" json:"token"`
	User         types.Int64  `tfsdk:"user" json:"user"`
}

func (o *tokensTerraformModel) Clone() tokensTerraformModel {
	return *o
}

func (o *tokensTerraformModel) BodyRequest() *tokensBodyRequestModel {
	var req tokensBodyRequestModel
	req.Application = o.Application.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Scope = o.Scope.ValueString()
	return &req
}

func (o *tokensTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Application, data["application"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetString(&o.Expires, data["expires"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.RefreshToken, data["refresh_token"], false))
	collect(helpers.AttrValueSetString(&o.Scope, data["scope"], false))
	collect(helpers.AttrValueSetString(&o.Token, data["token"], false))
	collect(helpers.AttrValueSetInt64(&o.User, data["user"]))
	return diags, nil
}

type tokensBodyRequestModel struct {
	Application int64  `json:"application,omitempty"`
	Description string `json:"description,omitempty"`
	Scope       string `json:"scope,omitempty"`
}
