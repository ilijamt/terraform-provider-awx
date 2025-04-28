package net_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/awx/credentials/net"
)

func TestNewResource(t *testing.T) {
	obj := net.NewResource()
	require.NotNil(t, obj)
	validators := obj.(resource.ResourceWithConfigValidators).ConfigValidators(t.Context())
	require.NotEmpty(t, validators)
	resp := &resource.MetadataResponse{}
	obj.Metadata(t.Context(), resource.MetadataRequest{ProviderTypeName: "awx"}, resp)
	require.Equal(t, "awx_credential_net", resp.TypeName)
}
