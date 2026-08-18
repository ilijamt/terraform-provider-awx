//go:build integration

package examples

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

// credentialTypeNamespaces holds every managed credential type AWX 24.6.1
// ships. The fixture declares one resource per entry, so a missing namespace
// fails the step.
var credentialTypeNamespaces = []string{
	"aim", "aws", "aws_secretsmanager_credential", "azure_kv", "azure_rm",
	"bitbucket_dc_token", "centrify_vault_kv", "conjur", "controller",
	"galaxy_api_token", "gce", "github_token", "gitlab_token", "gpg_public_key",
	"hashivault_kv", "hashivault_ssh", "insights", "kubernetes_bearer_token",
	"net", "openstack", "registry", "rhv", "satellite6", "scm", "ssh",
	"terraform", "thycotic_dsv", "thycotic_tss", "vault", "vmware",
}

// kind comes back from AWX, so matching it proves the namespace lookup resolved
// to the credential_type the resource claims.
func checkEveryCredentialType() resource.TestCheckFunc {
	checks := make([]resource.TestCheckFunc, 0, len(credentialTypeNamespaces)*2)
	for _, ns := range credentialTypeNamespaces {
		addr := fmt.Sprintf("awx_credential_%s.this", ns)
		checks = append(checks,
			resource.TestCheckResourceAttr(addr, "kind", ns),
			resource.TestCheckResourceAttrSet(addr, "credential_type"),
		)
	}
	return resource.ComposeAggregateTestCheckFunc(checks...)
}

// TestIntegration_CredentialTypes drives every managed credential type through
// create, update and import. The targeted assertions cover what the generator
// emits beyond plain strings.
func TestIntegration_CredentialTypes(t *testing.T) {
	httpClient := NewVCRClient(t, "credential_types")
	cfg := ReadFixture(t, filepath.Join("credential_types", "main.tf"))
	updated := ReadFixture(t, filepath.Join("credential_types", "update.tf"))

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
					checkEveryCredentialType(),

					resource.TestCheckResourceAttr("awx_credential_registry.this", "verify_ssl", "true"),
					resource.TestCheckResourceAttr("awx_credential_registry.this", "host", "quay.io"),
					resource.TestCheckResourceAttr("awx_credential_net.this", "authorize", "true"),
					resource.TestCheckResourceAttr("awx_credential_azure_kv.this", "cloud_name", "AzureCloud"),
					resource.TestCheckResourceAttr("awx_credential_hashivault_kv.this", "api_version", "v1"),
					resource.TestCheckResourceAttr("awx_credential_thycotic_dsv.this", "tld", "com"),
					resource.TestCheckResourceAttr("awx_credential_vault.this", "vault_id", "primary"),

					resource.TestCheckResourceAttrPair("data.awx_credential_registry.by_name", "id", "awx_credential_registry.this", "id"),
					resource.TestCheckResourceAttr("data.awx_credential_registry.by_name", "verify_ssl", "true"),
					resource.TestCheckNoResourceAttr("data.awx_credential_registry.by_name", "password"),
					resource.TestCheckResourceAttr("data.awx_credential_vault.by_id", "vault_id", "primary"),
					resource.TestCheckNoResourceAttr("data.awx_credential_vault.by_id", "vault_password"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkEveryCredentialType(),

					resource.TestCheckResourceAttr("awx_credential_registry.this", "verify_ssl", "false"),
					resource.TestCheckResourceAttr("awx_credential_registry.this", "host", "registry.example.com"),
					resource.TestCheckResourceAttr("awx_credential_aim.this", "verify", "false"),
					resource.TestCheckResourceAttr("awx_credential_kubernetes_bearer_token.this", "verify_ssl", "false"),
					resource.TestCheckResourceAttr("awx_credential_openstack.this", "verify_ssl", "false"),
					resource.TestCheckResourceAttr("awx_credential_controller.this", "verify_ssl", "false"),
					resource.TestCheckResourceAttr("awx_credential_azure_kv.this", "cloud_name", "AzureChinaCloud"),
					resource.TestCheckResourceAttr("awx_credential_hashivault_kv.this", "api_version", "v2"),
					resource.TestCheckResourceAttr("awx_credential_thycotic_dsv.this", "tld", "eu"),
					// AWX refuses a changed vault_id, so this one is recreated.
					resource.TestCheckResourceAttr("awx_credential_vault.this", "vault_id", "secondary"),
				),
			},
			{
				// AWX answers these with $encrypted$, and import has no prior
				// plan state to restore the real values from.
				ResourceName:            "awx_credential_registry.this",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "team", "user"},
			},
			{
				ResourceName:            "awx_credential_hashivault_kv.this",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token", "secret_id", "client_cert_private", "password", "team", "user"},
			},
		},
	})
}
