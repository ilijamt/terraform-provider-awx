package framework_test

import (
	"context"
	"testing"

	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestSearchField(t *testing.T) {
	t.Run("int64 field", func(t *testing.T) {
		f := framework.SearchField{Name: "id", Type: "int64"}
		assert.Equal(t, "id", f.Name)
		assert.Equal(t, "int64", f.Type)
		assert.False(t, f.URLEscape)
	})

	t.Run("string field with escape", func(t *testing.T) {
		f := framework.SearchField{Name: "name", Type: "string", URLEscape: true}
		assert.Equal(t, "name", f.Name)
		assert.True(t, f.URLEscape)
	})
}

func TestSearchGroup(t *testing.T) {
	t.Run("by_id group", func(t *testing.T) {
		g := framework.SearchGroup{
			Name:      "by_id",
			URLSuffix: "%d/",
			Fields:    []framework.SearchField{{Name: "id", Type: "int64"}},
		}
		assert.Equal(t, "by_id", g.Name)
		assert.Len(t, g.Fields, 1)
	})

	t.Run("multi-field group", func(t *testing.T) {
		g := framework.SearchGroup{
			Name:      "by_name_organization",
			URLSuffix: "?name__exact=%s&organization=%d",
			Fields: []framework.SearchField{
				{Name: "name", Type: "string", URLEscape: true},
				{Name: "organization", Type: "int64"},
			},
		}
		assert.Len(t, g.Fields, 2)
	})
}

// testSchema is reused across EvaluateSearchGroups tests.
var testSchema = dschema.Schema{
	Attributes: map[string]dschema.Attribute{
		"id":           dschema.Int64Attribute{Optional: true},
		"name":         dschema.StringAttribute{Optional: true},
		"organization": dschema.Int64Attribute{Optional: true},
	},
}

var testSchemaType = testSchema.Type().TerraformType(context.Background())

// buildConfig constructs a tfsdk.Config from the test schema and attribute values.
func buildConfig(values map[string]tftypes.Value) tfsdk.Config {
	return tfsdk.Config{
		Raw:    tftypes.NewValue(testSchemaType, values),
		Schema: testSchema,
	}
}

// nullNumber returns a null tftypes.Number value.
func nullNumber() tftypes.Value { return tftypes.NewValue(tftypes.Number, nil) }

// nullString returns a null tftypes.String value.
func nullString() tftypes.Value { return tftypes.NewValue(tftypes.String, nil) }

// unknownNumber returns an unknown tftypes.Number value.
func unknownNumber() tftypes.Value { return tftypes.NewValue(tftypes.Number, tftypes.UnknownValue) }

func TestEvaluateSearchGroups(t *testing.T) {
	ctx := context.Background()
	baseEndpoint := "/api/v2/inventories"

	groups := []framework.SearchGroup{
		{
			Name:      "by_id",
			URLSuffix: "%d/",
			Fields:    []framework.SearchField{{Name: "id", Type: "int64"}},
		},
		{
			Name:      "by_name",
			URLSuffix: "?name__exact=%s",
			Fields:    []framework.SearchField{{Name: "name", Type: "string", URLEscape: true}},
		},
		{
			Name:      "by_name_org",
			URLSuffix: "?name__exact=%s&organization=%d",
			Fields: []framework.SearchField{
				{Name: "name", Type: "string", URLEscape: true},
				{Name: "organization", Type: "int64"},
			},
		},
	}

	tests := []struct {
		name             string
		config           tfsdk.Config
		groups           []framework.SearchGroup
		expectedEndpoint string
		expectError      bool
	}{
		{
			name: "match by_id group",
			config: buildConfig(map[string]tftypes.Value{
				"id":           tftypes.NewValue(tftypes.Number, 42),
				"name":         nullString(),
				"organization": nullNumber(),
			}),
			groups:           groups,
			expectedEndpoint: "/api/v2/inventories/42",
			expectError:      false,
		},
		{
			name: "match by_name group",
			config: buildConfig(map[string]tftypes.Value{
				"id":           nullNumber(),
				"name":         tftypes.NewValue(tftypes.String, "test"),
				"organization": nullNumber(),
			}),
			groups:           groups,
			expectedEndpoint: "/api/v2/inventories/?name__exact=test",
			expectError:      false,
		},
		{
			// Uses a custom group list where by_name_org is the only option,
			// so it can match without by_name intercepting first.
			name: "match multi-field group",
			config: buildConfig(map[string]tftypes.Value{
				"id":           nullNumber(),
				"name":         tftypes.NewValue(tftypes.String, "test"),
				"organization": tftypes.NewValue(tftypes.Number, 1),
			}),
			groups: []framework.SearchGroup{
				{
					Name:      "by_name_org",
					URLSuffix: "?name__exact=%s&organization=%d",
					Fields: []framework.SearchField{
						{Name: "name", Type: "string", URLEscape: true},
						{Name: "organization", Type: "int64"},
					},
				},
			},
			expectedEndpoint: "/api/v2/inventories/?name__exact=test&organization=1",
			expectError:      false,
		},
		{
			name: "first match wins by_id over by_name",
			config: buildConfig(map[string]tftypes.Value{
				"id":           tftypes.NewValue(tftypes.Number, 42),
				"name":         tftypes.NewValue(tftypes.String, "test"),
				"organization": nullNumber(),
			}),
			groups:           groups,
			expectedEndpoint: "/api/v2/inventories/42",
			expectError:      false,
		},
		{
			name: "no group matches all null",
			config: buildConfig(map[string]tftypes.Value{
				"id":           nullNumber(),
				"name":         nullString(),
				"organization": nullNumber(),
			}),
			groups:      groups,
			expectError: true,
		},
		{
			name: "partial match skipped needs both name and org",
			config: buildConfig(map[string]tftypes.Value{
				"id":           nullNumber(),
				"name":         nullString(),
				"organization": tftypes.NewValue(tftypes.Number, 1),
			}),
			groups:      groups,
			expectError: true,
		},
		{
			name: "URL escape on string value",
			config: buildConfig(map[string]tftypes.Value{
				"id":           nullNumber(),
				"name":         tftypes.NewValue(tftypes.String, "hello world"),
				"organization": nullNumber(),
			}),
			groups:           groups,
			expectedEndpoint: "/api/v2/inventories/?name__exact=hello%20world",
			expectError:      false,
		},
		{
			name: "unknown value treated as unset",
			config: buildConfig(map[string]tftypes.Value{
				"id":           unknownNumber(),
				"name":         nullString(),
				"organization": nullNumber(),
			}),
			groups:      groups,
			expectError: true,
		},
		{
			name: "empty groups list",
			config: buildConfig(map[string]tftypes.Value{
				"id":           tftypes.NewValue(tftypes.Number, 42),
				"name":         nullString(),
				"organization": nullNumber(),
			}),
			groups:      []framework.SearchGroup{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint, diags := framework.EvaluateSearchGroups(ctx, tt.config, baseEndpoint, tt.groups)
			if tt.expectError {
				assert.True(t, diags.HasError())
				assert.Empty(t, endpoint)
				return
			}
			require.False(t, diags.HasError())
			assert.Equal(t, tt.expectedEndpoint, endpoint)
		})
	}
}
