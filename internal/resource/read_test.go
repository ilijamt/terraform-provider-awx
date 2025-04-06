package resource_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

func TestReadResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := NewMockClient(ctrl)
	ctx := t.Context()
	rci := resource.CallInfo{Name: "Name", Endpoint: "/", TypeName: "name"}

	t.Run("fail to create new request", func(t *testing.T) {
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodGet), gomock.Any(), gomock.Eq(nil)).Return(nil, fmt.Errorf("failed to create request")).Times(1)
		updater := &resourceUpdaterTest{}
		d, err := resource.Read(ctx, c, rci, types.Int64Value(1), updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
	})

	t.Run("fail to read resource", func(t *testing.T) {
		var r, _ = http.NewRequest(http.MethodDelete, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodGet), gomock.Any(), gomock.Eq(nil)).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{}, fmt.Errorf("fail to delete")).Times(1)
		updater := &resourceUpdaterTest{}
		d, err := resource.Read(ctx, c, rci, types.Int64Value(1), updater)
		assert.Error(t, err)
		assert.True(t, d.HasError())
	})

	t.Run("success read a resource", func(t *testing.T) {
		var r, _ = http.NewRequest(http.MethodDelete, "http://localhost", nil)
		c.EXPECT().NewRequest(gomock.Eq(ctx), gomock.Eq(http.MethodGet), gomock.Any(), gomock.Eq(nil)).Return(r, nil).Times(1)
		c.EXPECT().Do(gomock.Eq(ctx), gomock.Eq(r)).Return(map[string]any{}, nil).Times(1)
		updater := &resourceUpdaterTest{}
		d, err := resource.Read(ctx, c, rci, types.Int64Value(1), updater)
		assert.NoError(t, err)
		assert.False(t, d.HasError())
	})
}
