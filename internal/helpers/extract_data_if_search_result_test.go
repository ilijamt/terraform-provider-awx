package helpers_test

import (
	"encoding/json"
	"fmt"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
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

	t.Run("results array is empty should error out", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count":   1,
			"results": []map[string]any{},
		})
		require.Error(t, err)
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

	t.Run("results array has one entry but data is not map[string]any", func(t *testing.T) {
		result, d, err := helpers.ExtractDataIfSearchResult(map[string]any{
			"count": 1,
			"results": []any{
				string("id"),
			},
		})
		require.Error(t, err)
		require.True(t, d.HasError())
		require.Equal(t, 1, d.ErrorsCount())
		require.Empty(t, result)
	})

}
