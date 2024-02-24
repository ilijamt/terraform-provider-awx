package helpers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestExtractDataIfSearchResult(t *testing.T) {
	t.Run("no count number in the map", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{})
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Empty(t, result)
	})

	for _, val := range []any{
		int(0), int8(0), int16(0), int32(0), int64(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
		json.Number("0"), json.Number("NaN"),
		string("0"), string("NaN"),
	} {
		t.Run(fmt.Sprintf("count is %[1]v as %[1]T", val), func(t *testing.T) {
			result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
				"count": val,
			})
			require.Error(t, err)
			require.True(t, d.HasError())
			require.Equal(t, 1, d.ErrorsCount())
			require.Empty(t, result)
		})
	}

	t.Run("results array is []map[string]any{} should error out", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count":   1,
			"results": []map[string]any{},
		})
		require.ErrorContains(t, err, "[]map[string]interface {} instead of []any")
		require.True(t, d.HasError())
		require.Equal(t, 1, d.ErrorsCount())
		require.Empty(t, result)
	})

	t.Run("results array is []any{} should error out", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count":   1,
			"results": []any{},
		})
		require.ErrorContains(t, err, "expected 1 results, got 0")
		require.True(t, d.HasError())
		require.Equal(t, 1, d.ErrorsCount())
		require.Empty(t, result)
	})

	t.Run("results array has one entry", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count": 1,
			"results": []any{
				map[string]any{"id": 1},
			},
		})
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Equal(t, 0, d.ErrorsCount())
		require.EqualValues(t, map[string]any{"id": 1}, result)
	})

	t.Run("results array has multiple entries", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count": 2,
			"results": []any{
				map[string]any{"id": 1},
				map[string]any{"id": 2},
			},
		})
		require.ErrorContains(t, err, "received 2 entries, expected 1")
		require.True(t, d.HasError())
		require.Equal(t, 1, d.ErrorsCount())
		require.Empty(t, result)
	})

	t.Run("results array has one entry but data is not map[string]any", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count": 1,
			"results": []any{
				string("id"),
			},
		})
		require.ErrorContains(t, err, "received: map[string]interface {} instead of map[string]any")
		require.True(t, d.HasError())
		require.Equal(t, 1, d.ErrorsCount())
		require.Empty(t, result)
	})

}
