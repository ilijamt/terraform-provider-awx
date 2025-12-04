package aws_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/awx/credential/aws"
)

func TestNewResource(t *testing.T) {
	obj := aws.NewResource()
	require.NotNil(t, obj)
	validators := obj.(resource.ResourceWithConfigValidators).ConfigValidators(t.Context())
	require.NotEmpty(t, validators)
	resp := &resource.MetadataResponse{}
	obj.Metadata(t.Context(), resource.MetadataRequest{ProviderTypeName: "awx"}, resp)
	require.Equal(t, "awx_credential_aws", resp.TypeName)
}
