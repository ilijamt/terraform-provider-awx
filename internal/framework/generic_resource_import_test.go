package framework_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func importTestResource(t *testing.T, parts []string) (*framework.GenericResource[modelStub, bodyStub, *modelStub], *resource.ImportStateResponse) {
	t.Helper()
	schema := rschema.Schema{Attributes: map[string]rschema.Attribute{
		"id":        rschema.Int64Attribute{Computed: true},
		"parent_id": rschema.Int64Attribute{Required: true},
	}}
	r := &framework.GenericResource[modelStub, bodyStub, *modelStub]{
		Cfg: framework.ResourceCfg[modelStub, bodyStub]{
			Schema:        schema,
			IDKey:         "id",
			ImportIDParts: parts,
			ResourceName:  "Widget",
		},
	}
	ctx := context.Background()
	return r, &resource.ImportStateResponse{State: tfsdk.State{
		Schema: schema,
		Raw:    tftypes.NewValue(schema.Type().TerraformType(ctx), nil),
	}}
}

func TestGenericResource_ImportState_Composite(t *testing.T) {
	for _, tt := range []struct {
		name       string
		id         string
		wantError  bool
		wantParent int64
		wantID     int64
	}{
		{name: "both parts land on their attributes", id: "7/9", wantParent: 7, wantID: 9},
		{name: "too few parts", id: "9", wantError: true},
		{name: "too many parts", id: "1/2/3", wantError: true},
		{name: "non-numeric part", id: "seven/9", wantError: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r, resp := importTestResource(t, []string{"parent_id", "id"})
			r.ImportState(context.Background(), resource.ImportStateRequest{ID: tt.id}, resp)

			if tt.wantError {
				assert.True(t, resp.Diagnostics.HasError())
				return
			}
			require.False(t, resp.Diagnostics.HasError(), resp.Diagnostics.Errors())

			var parent, id int64
			require.False(t, resp.State.GetAttribute(context.Background(), path.Root("parent_id"), &parent).HasError())
			require.False(t, resp.State.GetAttribute(context.Background(), path.Root("id"), &id).HasError())
			assert.Equal(t, tt.wantParent, parent)
			assert.Equal(t, tt.wantID, id)
		})
	}
}
