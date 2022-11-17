package provider

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"awx": providerserver.NewProtocol6WithError(New("test")()),
}

const (
	providerConfig = `
provider "awx" {
  username             = "admin"
  password             = "admin"
  host     			   = "http://localhost"
  insecure_skip_verify = true"
}
`
)

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func TestProvider(t *testing.T) {
	awxProvider := testAccProtoV6ProviderFactories["awx"]
	require.NotNil(t, awxProvider)
	frameworkServer, err := awxProvider()
	require.NoError(t, err)
	require.NotNil(t, frameworkServer)
	schema, err := frameworkServer.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	require.NotNil(t, schema)
}
