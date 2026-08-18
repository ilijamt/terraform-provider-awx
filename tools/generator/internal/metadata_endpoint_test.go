package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemMetadataUrl(t *testing.T) {
	for _, tt := range []struct {
		name          string
		item          Item
		wantUrl       string
		wantDiscovery bool
	}{
		{
			name:    "falls back to the CRUD endpoint",
			item:    Item{Endpoint: "/api/v2/hosts/"},
			wantUrl: "/api/v2/hosts/",
		},
		{
			name:    "static override needs no discovery",
			item:    Item{Endpoint: "/api/v2/hosts/", MetadataEndpoint: "/api/v2/hosts/1/"},
			wantUrl: "/api/v2/hosts/1/",
		},
		{
			name: "a %d asks for discovery",
			item: Item{
				Endpoint:         "/api/v2/workflow_approval_templates/",
				MetadataEndpoint: "/api/v2/workflow_approval_templates/%d/",
			},
			wantUrl:       "/api/v2/workflow_approval_templates/%d/",
			wantDiscovery: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			url, needsDiscovery := tt.item.MetadataUrl()
			assert.Equal(t, tt.wantUrl, url)
			assert.Equal(t, tt.wantDiscovery, needsDiscovery)
		})
	}
}
