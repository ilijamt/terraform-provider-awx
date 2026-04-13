package framework_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestDataSourceBase_Configure(t *testing.T) {
	tests := []struct {
		name         string
		providerData any
		expectClient bool
		expectPanic  bool
	}{
		{
			name:         "nil provider data is no-op",
			providerData: nil,
			expectClient: false,
		},
		{
			name:         "valid Requester sets Client",
			providerData: successRequester(nil),
			expectClient: true,
		},
		{
			// Documents current behavior: bare type assertion panics
			// on wrong type (provider_base.go:15).
			name:         "wrong type panics",
			providerData: "not a requester",
			expectPanic:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &framework.DataSourceBase{}
			req := datasource.ConfigureRequest{ProviderData: tt.providerData}
			resp := &datasource.ConfigureResponse{}

			if tt.expectPanic {
				assert.Panics(t, func() {
					b.Configure(context.Background(), req, resp)
				})
				return
			}

			b.Configure(context.Background(), req, resp)
			if tt.expectClient {
				assert.NotNil(t, b.Client)
			} else {
				assert.Nil(t, b.Client)
			}
		})
	}
}

func TestDataSourceBase_Metadata(t *testing.T) {
	tests := []struct {
		name             string
		typeName         string
		providerTypeName string
		expected         string
	}{
		{
			name:             "standard",
			typeName:         "job_template",
			providerTypeName: "awx",
			expected:         "awx_job_template",
		},
		{
			name:             "empty provider name",
			typeName:         "job_template",
			providerTypeName: "",
			expected:         "_job_template",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &framework.DataSourceBase{}
			b.TypeName = tt.typeName

			req := datasource.MetadataRequest{ProviderTypeName: tt.providerTypeName}
			resp := &datasource.MetadataResponse{}

			b.Metadata(context.Background(), req, resp)
			assert.Equal(t, tt.expected, resp.TypeName)
		})
	}
}

// mockConfigValidator is a no-op datasource.ConfigValidator for testing.
type mockConfigValidator struct{}

func (mockConfigValidator) Description(context.Context) string          { return "" }
func (mockConfigValidator) MarkdownDescription(context.Context) string  { return "" }
func (mockConfigValidator) ValidateDataSource(context.Context, datasource.ValidateConfigRequest, *datasource.ValidateConfigResponse) {
}

func TestDataSourceBase_ConfigValidators(t *testing.T) {
	tests := []struct {
		name       string
		validators []datasource.ConfigValidator
		expectLen  int
		expectNil  bool
	}{
		{
			name:       "nil validators returns empty slice",
			validators: nil,
			expectLen:  0,
			expectNil:  false,
		},
		{
			name:       "non-nil validators returned as-is",
			validators: []datasource.ConfigValidator{mockConfigValidator{}},
			expectLen:  1,
			expectNil:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &framework.DataSourceBase{Validators: tt.validators}
			got := b.ConfigValidators(context.Background())
			assert.NotNil(t, got)
			assert.Len(t, got, tt.expectLen)
		})
	}
}
