//go:build integration

package examples

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

func TestIntegration_InstanceGroupAssociations(t *testing.T) {
	httpClient := NewVCRClient(t, "instance_group_associations")
	cfg := ReadFixture(t, filepath.Join("instance_group_associations", "main.tf"))
	updated := ReadFixture(t, filepath.Join("instance_group_associations", "update.tf"))

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
					resource.TestCheckResourceAttrPair(
						"awx_organization_associate_instance_group.org_link", "organization_id",
						"awx_organization.org", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_organization_associate_instance_group.org_link", "instance_group_id",
						"awx_instance_group.first", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_job_template_associate_instance_group.jt_link", "job_template_id",
						"awx_job_template.job_template", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_job_template_associate_instance_group.jt_link", "instance_group_id",
						"awx_instance_group.first", "id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"awx_organization_associate_instance_group.org_link", "instance_group_id",
						"awx_instance_group.second", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_job_template_associate_instance_group.jt_link", "instance_group_id",
						"awx_instance_group.second", "id"),
				),
			},
			{
				ResourceName:                         "awx_organization_associate_instance_group.org_link",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "organization_id",
				ImportStateIdFunc: associationImportID(
					"awx_organization_associate_instance_group.org_link",
					"organization_id", "instance_group_id"),
			},
			{
				ResourceName:                         "awx_job_template_associate_instance_group.jt_link",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "job_template_id",
				ImportStateIdFunc: associationImportID(
					"awx_job_template_associate_instance_group.jt_link",
					"job_template_id", "instance_group_id"),
			},
		},
	})
}

func associationImportID(addr, parentAttr, childAttr string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		r, ok := s.RootModule().Resources[addr]
		if !ok {
			return "", fmt.Errorf("%s not found in state", addr)
		}
		return fmt.Sprintf("%s/%s", r.Primary.Attributes[parentAttr], r.Primary.Attributes[childAttr]), nil
	}
}
