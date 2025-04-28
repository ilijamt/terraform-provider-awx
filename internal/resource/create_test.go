package resource_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

func TestCreate(t *testing.T) {
	ctx := t.Context()
	rci := resource.CallInfo{Name: "Name", Endpoint: "/", TypeName: "name"}

	t.Run("nil/data client", func(t *testing.T) {
		d, err := resource.Create(ctx, nil, rci.With(resource.SourceResource, resource.CalleeCreate), nil)
		require.NotEmpty(t, d)
		require.Error(t, err)
	})

	t.Run("fail to create new request", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodPost), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("failed to create request")).Times(1)
		updater := &dummyResource{}
		d, err := resource.Create(ctx, c, rci, updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
	})

	t.Run("fail to create resource", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		var r, _ = http.NewRequest(http.MethodPost, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodPost), gomock.Any(), gomock.Any()).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{}, fmt.Errorf("fail to create resource")).Times(1)
		updater := &dummyResource{}
		d, err := resource.Create(ctx, c, rci, updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
		assert.Empty(t, updater.data)
	})

	t.Run("success create a resource", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		var r, _ = http.NewRequest(http.MethodPost, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodPost), gomock.Any(), gomock.Any()).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{"id": 1}, nil).Times(1)
		updater := dummyResource{}
		d, err := resource.Create(ctx, c, rci, &updater)
		assert.NoError(t, err)
		assert.False(t, d.HasError())
		assert.NotEmpty(t, updater.data)
		assert.Equal(t, 1, updater.data["id"])
	})
}
