package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// executionEnvironmentTerraformModel maps the schema for ExecutionEnvironment when using Data Source
type executionEnvironmentTerraformModel struct {
	// Credential ""
	Credential types.Int64 `tfsdk:"credential" json:"credential"`
	// Description "Optional description of this execution environment."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this execution environment."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Image "The full image location, including the container registry, image name, and version tag."
	Image types.String `tfsdk:"image" json:"image"`
	// Managed ""
	Managed types.Bool `tfsdk:"managed" json:"managed"`
	// Name "Name of this execution environment."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "The organization used to determine access to this execution environment."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// Pull "Pull image before running?"
	Pull types.String `tfsdk:"pull" json:"pull"`
}

// Clone the object
func (o *executionEnvironmentTerraformModel) Clone() executionEnvironmentTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for ExecutionEnvironment
func (o *executionEnvironmentTerraformModel) BodyRequest() *executionEnvironmentBodyRequestModel {
	var req executionEnvironmentBodyRequestModel
	req.Credential = o.Credential.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Image = o.Image.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.Pull = o.Pull.ValueString()
	return &req
}

func (o *executionEnvironmentTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Credential, data["credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Image, data["image"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Managed, data["managed"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Pull, data["pull"], false)
		diags.Append(dg...)
	}
	return diags, nil
}

// executionEnvironmentBodyRequestModel maps the schema for ExecutionEnvironment for creating and updating the data
type executionEnvironmentBodyRequestModel struct {
	// Credential ""
	Credential int64 `json:"credential,omitempty"`
	// Description "Optional description of this execution environment."
	Description string `json:"description,omitempty"`
	// Image "The full image location, including the container registry, image name, and version tag."
	Image string `json:"image"`
	// Name "Name of this execution environment."
	Name string `json:"name"`
	// Organization "The organization used to determine access to this execution environment."
	Organization int64 `json:"organization,omitempty"`
	// Pull "Pull image before running?"
	Pull string `json:"pull,omitempty"`
}
