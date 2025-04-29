package {{ .PackageName }}_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/awx/credentials/{{ .PackageName }}"
)

func TestModel(t *testing.T) {
    t.Run("GetId", func(t *testing.T) {
        t.Run("no value for id", func(t *testing.T) {
            obj := {{ .PackageName }}.Model{}
            id, err := obj.GetId()
            assert.Error(t, err)
            assert.Empty(t, id)
        })

        t.Run("with value for id", func(t *testing.T) {
            obj := {{ .PackageName }}.Model{
				ID: types.Int64Value(1),
            }
            id, err := obj.GetId()
            assert.NoError(t, err)
            assert.Equal(t, "1", id)
        })
	})

	t.Run("Data", func(t *testing.T) {
		obj := {{ .PackageName }}.Model{
			ID:           types.Int64Value(1),
			Name:         types.StringValue("name"),
			Description:  types.StringValue("description"),
			Organization: types.Int64Value(1),
{{- range $key, $value := .Fields }}
{{- if $value.IsInput }}
           {{ $value.Generated.Name }}: {{ testTfValue $value.Generated.GoType $value.Id }},
{{- end }}
{{- end }}
		}

		data := obj.Data()
		assert.EqualValues(t, obj.Name.ValueString(), data.Name)
		assert.EqualValues(t, obj.Description.ValueString(), data.Description)
		assert.EqualValues(t, obj.Organization.ValueInt64(), *data.Organization)
	})

    t.Run("RequestBody", func(t *testing.T) {
    	obj := {{ .PackageName }}.Model{}
    	payload, err := obj.RequestBody()
    	assert.NoError(t, err)
    	assert.NotEmpty(t, payload)
    	objCloned := obj.Clone()
    	payloadCloned, err := objCloned.RequestBody()
    	assert.NoError(t, err)
    	assert.NotEmpty(t, payloadCloned)
    	assert.EqualValues(t, payload, payloadCloned)
    })
}

