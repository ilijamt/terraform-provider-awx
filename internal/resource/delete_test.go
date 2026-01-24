package resource_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

func TestDeleteResource(t *testing.T) {
	ctx := t.Context()
	rci := resource.CallInfo{Name: "Name", Endpoint: "/", TypeName: "name"}

	t.Run("client is nil", func(t *testing.T) {
		d, err := resource.Delete(ctx, nil, rci, nil)
		assert.Error(t, err)
		assert.NotEmpty(t, d)
	})

	t.Run("has invalid id", func(t *testing.T) {
		c, _ := client.NewTestingClient(t)
		updater := &dummyResource{getIdErr: fmt.Errorf("err")}
		d, err := resource.Delete(ctx, c, rci, updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
		assert.Empty(t, updater.data)
	})

	t.Run("fail to create new request", func(t *testing.T) {
		c, _ := client.NewTestingClient(t)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodDelete), gomock.Any(), gomock.Eq(nil)).Return(nil, fmt.Errorf("failed to create request")).Times(1)
		d, err := resource.Delete(ctx, c, rci, &dummyResource{getIdId: "1"})
		assert.Error(t, err)
		assert.True(t, d.HasError())
	})

	t.Run("fail to delete resource", func(t *testing.T) {
		c, _ := client.NewTestingClient(t)
		var r, _ = http.NewRequest(http.MethodDelete, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodDelete), gomock.Any(), gomock.Eq(nil)).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{}, fmt.Errorf("fail to delete")).Times(1)
		d, err := resource.Delete(ctx, c, rci, &dummyResource{getIdId: "1"})
		assert.Error(t, err)
		assert.True(t, d.HasError())
	})

	t.Run("success deleting a resource", func(t *testing.T) {
		c, _ := client.NewTestingClient(t)
		var r, _ = http.NewRequest(http.MethodDelete, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodDelete), gomock.Any(), gomock.Eq(nil)).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{}, nil).Times(1)
		d, err := resource.Delete(ctx, c, rci, &dummyResource{getIdId: "1"})
		assert.NoError(t, err)
		assert.False(t, d.HasError())
	})
}
