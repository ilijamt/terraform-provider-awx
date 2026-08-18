package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTfElementType(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
	}{
		{in: "integer", want: "Int64"},
		{in: "id", want: "Int64"},
		{in: "string", want: "String"},
		{in: "choice", want: "String"},
		{in: "boolean", want: "Bool"},
		{in: "decimal", want: "Float64"},
		{in: "", want: "String"},
	} {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.want, tfElementType(tt.in))
		})
	}
}

func TestTfElementTypeIsWhatTheTemplatesUse(t *testing.T) {
	fn, ok := FuncMap["tf_element_type"].(func(string) string)
	assert.True(t, ok, "templates spell element types through tf_element_type")
	assert.Equal(t, tfElementType(""), fn(""))
	assert.Contains(t, awxGoValue("list", ""), "types."+fn("")+"Type")
	assert.Contains(t, awxGoValue("set", ""), "types."+fn("")+"Type")
}
