package internal

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoLiteral(t *testing.T) {
	// what json.Unmarshal into map[string]any actually produces
	var decoded map[string]any
	require.NoError(t, json.Unmarshal([]byte(
		`{"s":"deprovisioning","b":false,"whole":5,"frac":2.5,"n":null}`), &decoded))

	for _, tc := range []struct {
		name string
		in   any
		want string
	}{
		{name: "string", in: decoded["s"], want: `"deprovisioning"`},
		{name: "bool", in: decoded["b"], want: "false"},
		{name: "whole number", in: decoded["whole"], want: "5"},
		{name: "fractional number", in: decoded["frac"], want: "2.5"},
		{name: "null", in: decoded["n"], want: "nil"},
		{name: "json.Number", in: json.Number("7"), want: "7"},
		{name: "quotes are escaped", in: `a"b`, want: `"a\"b"`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, goLiteral(tc.in))
		})
	}
}

// The rendered map has to be assignable to map[string]any, which is what a
// quote-everything render broke for anything but a string.
func TestGoLiteralCompilesAsMapValue(t *testing.T) {
	var m map[string]any
	require.NoError(t, json.Unmarshal([]byte(`{"a":"x","b":true,"c":3}`), &m))

	src := "package p\n\nvar _ = map[string]any{\n"
	for _, k := range []string{"a", "b", "c"} {
		src += "\t" + goLiteral(k) + ": " + goLiteral(m[k]) + ",\n"
	}
	src += "}\n"

	assert.Contains(t, src, `"a": "x",`)
	assert.Contains(t, src, `"b": true,`)
	assert.Contains(t, src, `"c": 3,`)
	assertParses(t, src)
}

func assertParses(t *testing.T, src string) {
	t.Helper()
	_, err := parser.ParseFile(token.NewFileSet(), "x.go", src, parser.AllErrors)
	require.NoError(t, err, "rendered source does not parse:\n%s", src)
}
