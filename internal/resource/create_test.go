package resource_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := NewMockClient(ctrl)
	ctx := t.Context()
	rci := resource.CallInfo{Name: "Name", Endpoint: "/", TypeName: "name"}
	_ = client

	t.Run("nil/data client", func(t *testing.T) {
		d, err := resource.Create(ctx, nil, rci.With(resource.SourceResource, resource.CalleeCreate), nil)
		require.NotEmpty(t, d)
		require.Error(t, err)
	})
}
