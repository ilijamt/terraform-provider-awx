//go:build integration

package examples

import (
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

// TestIntegration_JobTemplateSurveySpec exercises examples/job_template_survey_spec
// in isolation from the kitchen-sink preload_data fixture:
//
//   - Build the upstream chain (organization → project → inventory → job
//     template) needed for a survey to attach to.
//   - Apply a single-question float-typed survey spec.
//   - Update grows the survey to two questions and bumps the job template's
//     verbosity, exercising the survey's full re-PUT path AND a sibling job
//     template field flip in the same plan.
//   - Import targets the survey's job_template (the survey itself uses the
//     job_template_id as its identifier and is covered by ImportStateVerify
//     on the parent template).
func TestIntegration_JobTemplateSurveySpec(t *testing.T) {
	httpClient := NewVCRClient(t, "job_template_survey_spec")
	cfg := ReadFixture(t, filepath.Join("job_template_survey_spec", "main.tf"))
	updated := ReadFixture(t, filepath.Join("job_template_survey_spec", "update.tf"))

	factories := map[string]func() (tfprotov6.ProviderServer, error){
		"awx": providerserver.NewProtocol6WithError(
			provider.NewFuncProvider(version.Version, httpClient, awx.Resources(), awx.DataSources())(),
		),
	}

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: factories,
		Steps: []resource.TestStep{
			{
				Config: providerHeader(t) + cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_organization.demo_organization", "name", "Job Template Organization"),
					resource.TestCheckResourceAttr("awx_project.demo_project", "scm_type", "git"),
					resource.TestCheckResourceAttr("awx_inventory.demo_inventory", "name", "Job Demo Inventory"),

					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "name", "Job Demo Job Template"),
					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "verbosity", "0"),
					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "allow_simultaneous", "true"),
					resource.TestCheckResourceAttrPair("awx_job_template.demo_job_template", "inventory", "awx_inventory.demo_inventory", "id"),
					resource.TestCheckResourceAttrPair("awx_job_template.demo_job_template", "project", "awx_project.demo_project", "id"),

					resource.TestCheckResourceAttrPair("awx_job_template_survey_spec.demo_job_template", "job_template_id", "awx_job_template.demo_job_template", "id"),
					resource.TestCheckResourceAttrSet("awx_job_template_survey_spec.demo_job_template", "spec"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "verbosity", "1"),
					resource.TestCheckResourceAttrPair("awx_job_template_survey_spec.demo_job_template", "job_template_id", "awx_job_template.demo_job_template", "id"),
					resource.TestCheckResourceAttrSet("awx_job_template_survey_spec.demo_job_template", "spec"),
				),
			},
			{
				ResourceName:      "awx_job_template.demo_job_template",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
