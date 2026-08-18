package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONDocumentRoundTrip(t *testing.T) {
	for _, tt := range []struct {
		name    string
		payload string
	}{
		{
			name:    "keeps the key order it was given",
			payload: "{\n  \"z\": 1,\n  \"a\": 2,\n  \"m\": 3\n}",
		},
		{
			name:    "leaves &, < and > unescaped",
			payload: "{\n  \"url_suffix\": \"?name__exact=%s&organization=%d&a=<b>\"\n}",
		},
		{
			name:    "keeps numbers in their source spelling",
			payload: "{\n  \"a\": 1,\n  \"b\": -2.50,\n  \"c\": 1e3\n}",
		},
		{
			name:    "renders empty containers inline",
			payload: "{\n  \"a\": {},\n  \"b\": []\n}",
		},
		{
			name:    "handles nesting, booleans and null",
			payload: "{\n  \"a\": [\n    {\n      \"b\": true,\n      \"c\": null\n    },\n    [\n      \"d\"\n    ]\n  ]\n}",
		},
		{
			name:    "escapes what JSON has to escape",
			payload: "{\n  \"a\": \"line\\nbreak \\\"quoted\\\" \\\\ tab\\t\"\n}",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var document = mustDecodeJSON(t, tt.payload)
			payload, err := encodeJSONDocument(document)
			require.NoError(t, err)
			require.Equal(t, tt.payload, string(payload))
		})
	}
}

func TestDecodeJSONDocumentRejectsTrailingData(t *testing.T) {
	var _, err = decodeJSONDocument([]byte(`{"a": 1} {"b": 2}`))
	require.ErrorContains(t, err, "unexpected data")
}
