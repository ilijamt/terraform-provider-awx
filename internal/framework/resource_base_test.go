package framework_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestResourceBase_Configure(t *testing.T) {
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
			b := &framework.ResourceBase{}
			req := resource.ConfigureRequest{ProviderData: tt.providerData}
			resp := &resource.ConfigureResponse{}

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

func TestResourceBase_Metadata(t *testing.T) {
	tests := []struct {
		name             string
		typeName         string
		providerTypeName string
		expected         string
	}{
		{
			name:             "standard",
			typeName:         "inventory",
			providerTypeName: "awx",
			expected:         "awx_inventory",
		},
		{
			name:             "empty provider name",
			typeName:         "inventory",
			providerTypeName: "",
			expected:         "_inventory",
		},
		{
			name:             "empty type name",
			typeName:         "",
			providerTypeName: "awx",
			expected:         "awx_",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &framework.ResourceBase{}
			b.TypeName = tt.typeName

			req := resource.MetadataRequest{ProviderTypeName: tt.providerTypeName}
			resp := &resource.MetadataResponse{}

			b.Metadata(context.Background(), req, resp)
			assert.Equal(t, tt.expected, resp.TypeName)
		})
	}
}
