package helpers_test

import (
	"errors"
	"testing"

	"github.com/hashicorp/go-multierror"
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

		diags, err := helpers.ApplyFieldMappings(nil, mappings)

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

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
	})

	t.Run("successful field mapping", func(t *testing.T) {
		data := map[string]any{
			"field1": "value1",
			"field2": 42,
		}

		var capturedValues []any
		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					capturedValues = append(capturedValues, value)
					return nil, nil
				},
			},
			{
				APIField: "field2",
				Setter: func(value any) (diag.Diagnostics, error) {
					capturedValues = append(capturedValues, value)
					return nil, nil
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
		assert.Equal(t, []any{"value1", 42}, capturedValues)
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

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.Error(t, err)
		assert.False(t, diags.HasError())
	})

	t.Run("setter with diagnostics", func(t *testing.T) {
		data := map[string]any{
			"field1": "value1",
		}

		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					var d diag.Diagnostics
					d.AddWarning("test warning", "this is a warning")
					return d, nil
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
		assert.Equal(t, 1, len(diags))
		assert.Equal(t, "test warning", diags[0].Summary())
	})

	t.Run("multiple setters with errors", func(t *testing.T) {
		data := map[string]any{
			"field1": "value1",
			"field2": "value2",
		}

		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					return nil, err1
				},
			},
			{
				APIField: "field2",
				Setter: func(value any) (diag.Diagnostics, error) {
					return nil, err2
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.Error(t, err)
		assert.False(t, diags.HasError())

		// Check that both errors are in the multierror
		merr, ok := err.(*multierror.Error)
		assert.True(t, ok)
		assert.Equal(t, 2, len(merr.Errors))
		assert.Contains(t, merr.Error(), err1.Error())
		assert.Contains(t, merr.Error(), err2.Error())
	})

	t.Run("missing field", func(t *testing.T) {
		data := map[string]any{
			"field1": "value1",
		}

		var capturedValue any
		mappings := []helpers.FieldMapping{
			{
				APIField: "missing_field",
				Setter: func(value any) (diag.Diagnostics, error) {
					capturedValue = value
					return nil, nil
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.NoError(t, err)
		assert.False(t, diags.HasError())
		assert.Nil(t, capturedValue) // Should be nil since the field is missing
	})

	t.Run("mixed errors and diagnostics", func(t *testing.T) {
		data := map[string]any{
			"field1": "value1",
			"field2": "value2",
		}

		expectedErr := errors.New("setter error")
		mappings := []helpers.FieldMapping{
			{
				APIField: "field1",
				Setter: func(value any) (diag.Diagnostics, error) {
					var d diag.Diagnostics
					d.AddWarning("test warning", "this is a warning")
					return d, nil
				},
			},
			{
				APIField: "field2",
				Setter: func(value any) (diag.Diagnostics, error) {
					var d diag.Diagnostics
					d.AddError("test error", "this is a diagnostic error")
					return d, expectedErr
				},
			},
		}

		diags, err := helpers.ApplyFieldMappings(data, mappings)

		assert.Error(t, err)
		assert.True(t, diags.HasError())
		assert.GreaterOrEqual(t, 1, diags.WarningsCount())
		assert.Equal(t, 2, len(diags))
		assert.Equal(t, "test warning", diags[0].Summary())
		assert.Equal(t, "test error", diags[1].Summary())
	})
}
