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

// TestIntegration_Credentials consolidates examples/credentials and
// examples/credential_with_input_source. It exercises:
//
//   - credentials owned by an organization, a team, and a user
//   - a Vault → Gitlab credential_input_source
//   - update on a credential whose `ssh_key_data` is NOT bound by an input
//     source (avoids an upstream AWX 500 on PATCH-with-input-source — see
//     awx_credential.gitlab below; it is intentionally unchanged in update.tf)
//   - import round-trip on awx_credential.organization
func TestIntegration_Credentials(t *testing.T) {
	httpClient := NewVCRClient(t, "credentials")
	cfg := ReadFixture(t, filepath.Join("credentials", "main.tf"))
	updated := ReadFixture(t, filepath.Join("credentials", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_organization.org", "name", "Credentials"),
					resource.TestCheckResourceAttr("awx_team.team", "name", "Team"),
					resource.TestCheckResourceAttr("awx_user.user", "username", "credentials_user"),

					resource.TestCheckResourceAttr("awx_credential.organization", "name", "Organization"),
					resource.TestCheckResourceAttr("awx_credential.organization", "description", "Assigned to Organization"),
					resource.TestCheckResourceAttrPair("awx_credential.organization", "organization", "awx_organization.org", "id"),
					resource.TestCheckResourceAttrPair("awx_credential.team", "team", "awx_team.team", "id"),
					resource.TestCheckResourceAttrPair("awx_credential.user", "user", "awx_user.user", "id"),

					resource.TestCheckResourceAttr("awx_credential.vault", "name", "Vault"),
					resource.TestCheckResourceAttr("awx_credential.gitlab", "description", "Gitlab all access"),
					resource.TestCheckResourceAttrSet("awx_credential_input_source.gitlab", "id"),
					resource.TestCheckResourceAttrPair("awx_credential_input_source.gitlab", "target_credential", "awx_credential.gitlab", "id"),
					resource.TestCheckResourceAttrPair("awx_credential_input_source.gitlab", "source_credential", "awx_credential.vault", "id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_credential.organization", "description", "Assigned to Organization (updated)"),
					resource.TestCheckResourceAttr("awx_credential.vault", "description", "Vault token lookup source for other credentials (updated)"),
					resource.TestCheckResourceAttr("awx_credential.gitlab", "description", "Gitlab all access"),
				),
			},
			{
				ResourceName:      "awx_credential.organization",
				ImportState:       true,
				ImportStateVerify: true,
				// `inputs` is encrypted by AWX and never returned in plain
				// form. `team` / `user` are mutually-exclusive ownership
				// fields stored as `0` in plan but absent from the imported
				// state for an organization-owned credential.
				ImportStateVerifyIgnore: []string{"inputs", "team", "user"},
			},
		},
	})
}
