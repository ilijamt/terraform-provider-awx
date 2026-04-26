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

// TestIntegration_CredentialAws exercises the typed awx_credential_aws resource:
//
//   - Create with username/password, verify the credential_type was resolved
//     by the OnConfigure namespace lookup.
//   - Update name + add security_token (proves the typed Sensitive field
//     round-trips through the inputs JSON encoding on the wire).
//   - Read by name through the awx_credential_aws data source.
//   - Import round-trip — `password` and `security_token` are excluded
//     because AWX returns `$encrypted$` placeholders for secrets and
//     the typed-field merge hook only restores them when prior plan state
//     exists (which import lacks).
func TestIntegration_CredentialAws(t *testing.T) {
	httpClient := NewVCRClient(t, "credential_aws")
	cfg := ReadFixture(t, filepath.Join("credential_aws", "main.tf"))
	updated := ReadFixture(t, filepath.Join("credential_aws", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_credential_aws.smoke", "name", "AWS smoke"),
					resource.TestCheckResourceAttr("awx_credential_aws.smoke", "username", "AKIAEXAMPLE"),
					resource.TestCheckResourceAttrSet("awx_credential_aws.smoke", "credential_type"),
					resource.TestCheckResourceAttr("awx_credential_aws.smoke", "kind", "aws"),
					resource.TestCheckResourceAttrPair("awx_credential_aws.smoke", "organization", "awx_organization.org", "id"),
					resource.TestCheckResourceAttrPair("data.awx_credential_aws.smoke_by_name", "id", "awx_credential_aws.smoke", "id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_credential_aws.smoke", "name", "AWS smoke (updated)"),
					resource.TestCheckResourceAttr("awx_credential_aws.smoke", "username", "AKIAUPDATED"),
					resource.TestCheckResourceAttr("awx_credential_aws.smoke", "security_token", "sts-token-value"),
				),
			},
			{
				ResourceName:            "awx_credential_aws.smoke",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "security_token", "team", "user"},
			},
		},
	})
}
