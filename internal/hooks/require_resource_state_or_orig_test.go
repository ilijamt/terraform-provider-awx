package hooks_test

import (
	"testing"
	"time"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
	"github.com/stretchr/testify/require"
)

func TestRequireResourceStateOrOrig(t *testing.T) {
	t.Run("orig and state are nil", func(t *testing.T) {
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeCreate, nil, nil))
		require.Error(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeRead, nil, nil))
		require.Error(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeUpdate, nil, nil))
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceData, hooks.CalleeUpdate, nil, nil))
	})
	t.Run("orig is nil and state has data", func(t *testing.T) {
		var obj = time.Now()
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeCreate, nil, &obj))
		require.Error(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeRead, nil, &obj))
		require.Error(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeUpdate, nil, &obj))
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceData, hooks.CalleeUpdate, nil, &obj))
	})
	t.Run("orig has data and state is nil", func(t *testing.T) {
		var obj = time.Now()
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeCreate, &obj, nil))
		require.Error(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeRead, &obj, nil))
		require.Error(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeUpdate, &obj, nil))
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceData, hooks.CalleeUpdate, &obj, nil))
	})
	t.Run("orig and state have data", func(t *testing.T) {
		var obj = time.Now()
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeCreate, &obj, &obj))
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeRead, &obj, &obj))
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceResource, hooks.CalleeUpdate, &obj, &obj))
		require.NoError(t, hooks.RequireResourceStateOrOrig(t.Context(), "v1.0.0", hooks.SourceData, hooks.CalleeUpdate, &obj, &obj))
	})
}
