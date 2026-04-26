package helpers_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestMergeEncryptedField(t *testing.T) {
	tests := []struct {
		name        string
		orig        types.String
		cur         types.String
		wantValue   types.String
		wantSubbed  bool
	}{
		{
			name:       "plain value passes through",
			orig:       types.StringValue("old-secret"),
			cur:        types.StringValue("new-secret"),
			wantValue:  types.StringValue("new-secret"),
			wantSubbed: false,
		},
		{
			name:       "encrypted placeholder with original substitutes",
			orig:       types.StringValue("real-secret"),
			cur:        types.StringValue("$encrypted$"),
			wantValue:  types.StringValue("real-secret"),
			wantSubbed: true,
		},
		{
			name:       "encrypted placeholder with null orig (import) keeps placeholder",
			orig:       types.StringNull(),
			cur:        types.StringValue("$encrypted$"),
			wantValue:  types.StringValue("$encrypted$"),
			wantSubbed: false,
		},
		{
			name:       "encrypted placeholder with empty orig keeps placeholder",
			orig:       types.StringValue(""),
			cur:        types.StringValue("$encrypted$"),
			wantValue:  types.StringValue("$encrypted$"),
			wantSubbed: false,
		},
		{
			name:       "unknown orig keeps placeholder",
			orig:       types.StringUnknown(),
			cur:        types.StringValue("$encrypted$"),
			wantValue:  types.StringValue("$encrypted$"),
			wantSubbed: false,
		},
		{
			name:       "null current is preserved untouched",
			orig:       types.StringValue("real-secret"),
			cur:        types.StringNull(),
			wantValue:  types.StringNull(),
			wantSubbed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, subbed := helpers.MergeEncryptedField(tt.orig, tt.cur)
			assert.Equal(t, tt.wantSubbed, subbed)
			assert.Equal(t, tt.wantValue, got)
		})
	}
}
