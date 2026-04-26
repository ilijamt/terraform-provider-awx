package framework_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestObjectRolesDataSource_Schema(t *testing.T) {
	tests := []struct {
		name         string
		typeName     string
		displayName  string
		deprecated   bool
		wantIDDesc   string
		wantRoleDesc string
		wantDeprMsg  string
	}{
		{
			name:         "instance group not deprecated",
			typeName:     "instance_group_object_roles",
			displayName:  "InstanceGroup",
			deprecated:   false,
			wantIDDesc:   "InstanceGroup ID",
			wantRoleDesc: "Roles for instance_group",
			wantDeprMsg:  "",
		},
		{
			name:         "credential deprecated",
			typeName:     "credential_object_roles",
			displayName:  "Credential",
			deprecated:   true,
			wantIDDesc:   "Credential ID",
			wantRoleDesc: "Roles for credential",
			wantDeprMsg:  "This data source has been deprecated and will be removed in a future release.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := framework.NewObjectRolesDataSource(tt.typeName, "/api/v2/example/%d/object_roles/", tt.displayName, tt.deprecated)
			schemaProvider, ok := ds.(interface {
				Schema(context.Context, datasource.SchemaRequest, *datasource.SchemaResponse)
			})
			require.True(t, ok, "constructed data source must implement Schema")

			resp := &datasource.SchemaResponse{}
			schemaProvider.Schema(context.Background(), datasource.SchemaRequest{}, resp)

			assert.Equal(t, tt.wantDeprMsg, resp.Schema.DeprecationMessage)

			idAttr, ok := resp.Schema.Attributes["id"]
			require.True(t, ok, "id attribute missing")
			assert.Equal(t, tt.wantIDDesc, idAttr.GetDescription())
			assert.True(t, idAttr.IsRequired())

			rolesAttr, ok := resp.Schema.Attributes["roles"]
			require.True(t, ok, "roles attribute missing")
			assert.Equal(t, tt.wantRoleDesc, rolesAttr.GetDescription())
			assert.True(t, rolesAttr.IsComputed())
		})
	}
}

func TestObjectRolesDataSource_Metadata(t *testing.T) {
	ds := framework.NewObjectRolesDataSource("inventory_object_roles", "/api/v2/inventories/%d/object_roles/", "Inventory", true)
	mdProvider, ok := ds.(interface {
		Metadata(context.Context, datasource.MetadataRequest, *datasource.MetadataResponse)
	})
	require.True(t, ok)

	resp := &datasource.MetadataResponse{}
	mdProvider.Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "awx"}, resp)

	assert.Equal(t, "awx_inventory_object_roles", resp.TypeName)
}
