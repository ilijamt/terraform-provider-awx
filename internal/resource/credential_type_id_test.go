package resource_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

func TestGetCredentialTypeID(t *testing.T) {
	ctx := t.Context()
	rci := resource.CallInfo{Name: "Name", Endpoint: "/", TypeName: "name"}

	t.Run("nil client and empty name", func(t *testing.T) {
		d, err := resource.GetCredentialTypeID(ctx, nil, rci.With(resource.SourceResource, resource.CalleeCreate), "")
		require.NotEmpty(t, d)
		require.Error(t, err)
	})
}
