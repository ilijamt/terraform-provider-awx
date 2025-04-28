package helpers_test

import (
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestApplyFieldMappings(t *testing.T) {
	t.Run("nil data map", func(t *testing.T) {
		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					return nil, nil
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(nil, mappings...)

		assert.Error(t, err)
		assert.Equal(t, "data must not be nil", err.Error())
		assert.True(t, diags.HasError())
		assert.Equal(t, 1, len(diags))
		assert.Equal(t, "nil pointer", diags[0].Summary())
	})

	t.Run("empty data map", func(t *testing.T) {
		data := map[string]any{}
		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					return nil, nil
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings...)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
	})

	t.Run("no field mappings with data", func(t *testing.T) {
		diags, err := helpers.ApplyFieldMappings(
			map[string]any{
				"field1": "value",
				"field2": "value",
			},
		)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
	})

	t.Run("field mapping with data", func(t *testing.T) {
		var capturedValueField1, capturedValueField2 string
		diags, err := helpers.ApplyFieldMappings(
			map[string]any{
				"field1": "value",
				"field2": "value",
			},
			helpers.FieldMapping{
				APIField: "field1",
				Data: map[string]any{
					"field1": "42",
				},
				Setter: func(value any) (diag.Diagnostics, error) {
					capturedValueField1 = value.(string)
					return nil, nil
				},
			},
			helpers.FieldMapping{
				APIField: "field2",
				Setter: func(value any) (diag.Diagnostics, error) {
					capturedValueField2 = value.(string)
					return nil, nil
				},
			},
		)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
		assert.Equal(t, "42", capturedValueField1)
		assert.Equal(t, "value", capturedValueField2)
	})

	t.Run("setter with error", func(t *testing.T) {
		data := map[string]any{
			"field1": "value1",
		}

		expectedErr := errors.New("setter error")
		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					return nil, expectedErr
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings...)

		assert.Error(t, err)
		assert.False(t, diags.HasError())
	})
}
