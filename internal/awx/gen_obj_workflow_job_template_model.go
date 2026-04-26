package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type workflowJobTemplateTerraformModel struct {
	AllowSimultaneous    types.Bool   `tfsdk:"allow_simultaneous" json:"allow_simultaneous"`
	AskInventoryOnLaunch types.Bool   `tfsdk:"ask_inventory_on_launch" json:"ask_inventory_on_launch"`
	AskLabelsOnLaunch    types.Bool   `tfsdk:"ask_labels_on_launch" json:"ask_labels_on_launch"`
	AskLimitOnLaunch     types.Bool   `tfsdk:"ask_limit_on_launch" json:"ask_limit_on_launch"`
	AskScmBranchOnLaunch types.Bool   `tfsdk:"ask_scm_branch_on_launch" json:"ask_scm_branch_on_launch"`
	AskSkipTagsOnLaunch  types.Bool   `tfsdk:"ask_skip_tags_on_launch" json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch      types.Bool   `tfsdk:"ask_tags_on_launch" json:"ask_tags_on_launch"`
	AskVariablesOnLaunch types.Bool   `tfsdk:"ask_variables_on_launch" json:"ask_variables_on_launch"`
	Description          types.String `tfsdk:"description" json:"description"`
	ExtraVars            types.String `tfsdk:"extra_vars" json:"extra_vars"`
	ID                   types.Int64  `tfsdk:"id" json:"id"`
	Inventory            types.Int64  `tfsdk:"inventory" json:"inventory"`
	JobTags              types.String `tfsdk:"job_tags" json:"job_tags"`
	Limit                types.String `tfsdk:"limit" json:"limit"`
	Name                 types.String `tfsdk:"name" json:"name"`
	Organization         types.Int64  `tfsdk:"organization" json:"organization"`
	ScmBranch            types.String `tfsdk:"scm_branch" json:"scm_branch"`
	SkipTags             types.String `tfsdk:"skip_tags" json:"skip_tags"`
	SurveyEnabled        types.Bool   `tfsdk:"survey_enabled" json:"survey_enabled"`
	WebhookCredential    types.Int64  `tfsdk:"webhook_credential" json:"webhook_credential"`
	WebhookService       types.String `tfsdk:"webhook_service" json:"webhook_service"`
}

func (o *workflowJobTemplateTerraformModel) Clone() workflowJobTemplateTerraformModel {
	return *o
}

func (o *workflowJobTemplateTerraformModel) BodyRequest() *workflowJobTemplateBodyRequestModel {
	var req workflowJobTemplateBodyRequestModel
	req.AllowSimultaneous = o.AllowSimultaneous.ValueBool()
	req.AskInventoryOnLaunch = o.AskInventoryOnLaunch.ValueBool()
	req.AskLabelsOnLaunch = o.AskLabelsOnLaunch.ValueBool()
	req.AskLimitOnLaunch = o.AskLimitOnLaunch.ValueBool()
	req.AskScmBranchOnLaunch = o.AskScmBranchOnLaunch.ValueBool()
	req.AskSkipTagsOnLaunch = o.AskSkipTagsOnLaunch.ValueBool()
	req.AskTagsOnLaunch = o.AskTagsOnLaunch.ValueBool()
	req.AskVariablesOnLaunch = o.AskVariablesOnLaunch.ValueBool()
	req.Description = o.Description.ValueString()
	req.ExtraVars = json.RawMessage(o.ExtraVars.String())
	req.Inventory = o.Inventory.ValueInt64()
	req.JobTags = o.JobTags.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.SkipTags = o.SkipTags.ValueString()
	req.SurveyEnabled = o.SurveyEnabled.ValueBool()
	req.WebhookCredential = o.WebhookCredential.ValueInt64()
	req.WebhookService = o.WebhookService.ValueString()
	return &req
}

func (o *workflowJobTemplateTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.AllowSimultaneous, data["allow_simultaneous"]))
	collect(helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data["ask_inventory_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data["ask_labels_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data["ask_limit_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data["ask_scm_branch_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data["ask_skip_tags_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data["ask_tags_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data["ask_variables_on_launch"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetString(&o.SkipTags, data["skip_tags"], false))
	collect(helpers.AttrValueSetBool(&o.SurveyEnabled, data["survey_enabled"]))
	collect(helpers.AttrValueSetInt64(&o.WebhookCredential, data["webhook_credential"]))
	collect(helpers.AttrValueSetString(&o.WebhookService, data["webhook_service"], false))
	return diags, nil
}

type workflowJobTemplateBodyRequestModel struct {
	AllowSimultaneous    bool            `json:"allow_simultaneous"`
	AskInventoryOnLaunch bool            `json:"ask_inventory_on_launch"`
	AskLabelsOnLaunch    bool            `json:"ask_labels_on_launch"`
	AskLimitOnLaunch     bool            `json:"ask_limit_on_launch"`
	AskScmBranchOnLaunch bool            `json:"ask_scm_branch_on_launch"`
	AskSkipTagsOnLaunch  bool            `json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch      bool            `json:"ask_tags_on_launch"`
	AskVariablesOnLaunch bool            `json:"ask_variables_on_launch"`
	Description          string          `json:"description,omitempty"`
	ExtraVars            json.RawMessage `json:"extra_vars,omitempty"`
	Inventory            int64           `json:"inventory,omitempty"`
	JobTags              string          `json:"job_tags,omitempty"`
	Limit                string          `json:"limit,omitempty"`
	Name                 string          `json:"name"`
	Organization         int64           `json:"organization,omitempty"`
	ScmBranch            string          `json:"scm_branch,omitempty"`
	SkipTags             string          `json:"skip_tags,omitempty"`
	SurveyEnabled        bool            `json:"survey_enabled"`
	WebhookCredential    int64           `json:"webhook_credential,omitempty"`
	WebhookService       string          `json:"webhook_service,omitempty"`
}
