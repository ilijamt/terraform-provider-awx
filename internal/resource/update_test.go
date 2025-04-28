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

func TestUpdate(t *testing.T) {
	ctx := t.Context()
	rci := resource.CallInfo{Name: "Name", Endpoint: "/", TypeName: "name"}.With(resource.SourceResource, resource.CalleeUpdate)

	t.Run("nil/data client", func(t *testing.T) {
		d, err := resource.Update(ctx, nil, rci, nil)
		require.NotEmpty(t, d)
		require.Error(t, err)
	})

	t.Run("fail to create new request", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodPatch), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("failed to create request")).Times(1)
		updater := &dummyResource{getIdId: "1"}
		d, err := resource.Update(ctx, c, rci, updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
	})

	t.Run("has invalid id", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		updater := &dummyResource{getIdErr: fmt.Errorf("err")}
		d, err := resource.Update(ctx, c, rci, updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
		assert.Empty(t, updater.data)
	})

	t.Run("fail to update resource", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		var r, _ = http.NewRequest(http.MethodPatch, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodPatch), gomock.Any(), gomock.Any()).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{}, fmt.Errorf("fail to create resource")).Times(1)
		updater := &dummyResource{getIdId: "1"}
		d, err := resource.Update(ctx, c, rci, updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
		assert.Empty(t, updater.data)
	})

	t.Run("success update a resource", func(t *testing.T) {
		c := NewMockClient(gomock.NewController(t))
		var r, _ = http.NewRequest(http.MethodPatch, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodPatch), gomock.Any(), gomock.Any()).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{"id": 1}, nil).Times(1)
		updater := &dummyResource{getIdId: "1"}
		d, err := resource.Update(ctx, c, rci, updater)
		assert.NoError(t, err)
		assert.False(t, d.HasError())
		assert.NotEmpty(t, updater.data)
		assert.Equal(t, 1, updater.data["id"])
	})

}
