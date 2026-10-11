package awx

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const importExamplesDir = "../../examples/resources"

// The generator derives import examples from its config rather than from the
// provider, so check every documented ID against the real ImportState.
func TestImportExamplesMatchImportState(t *testing.T) {
	ctx := context.Background()
	registered := map[string]bool{}

	for _, fn := range Resources() {
		r := fn()
		var md resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "awx"}, &md)
		registered[md.TypeName] = true

		t.Run(md.TypeName, func(t *testing.T) {
			id, documented := exampleImportID(t, md.TypeName)
			importer, ok := r.(resource.ResourceWithImportState)
			if !ok {
				assert.False(t, documented, "import example for a resource without ImportState")
				return
			}
			if documented {
				assert.False(t, importHasError(ctx, r, importer, id), "documented import ID %q is rejected", id)
				return
			}
			for _, id := range []string{"1", "1/2"} {
				assert.True(t, importHasError(ctx, r, importer, id), "import of %q succeeds but has no example", id)
			}
		})
	}

	entries, err := os.ReadDir(importExamplesDir)
	require.NoError(t, err)
	for _, e := range entries {
		assert.True(t, registered[e.Name()], "import example for unregistered resource %s", e.Name())
	}
}

func exampleImportID(t *testing.T, typeName string) (string, bool) {
	content, err := os.ReadFile(filepath.Join(importExamplesDir, typeName, "import.sh"))
	if os.IsNotExist(err) {
		return "", false
	}
	require.NoError(t, err)
	for _, line := range strings.Split(string(content), "\n") {
		if fields := strings.Fields(line); len(fields) == 4 && fields[0] == "terraform" && fields[1] == "import" {
			return fields[3], true
		}
	}
	t.Fatalf("no terraform import line in %s example", typeName)
	return "", false
}

func importHasError(ctx context.Context, r resource.Resource, importer resource.ResourceWithImportState, id string) bool {
	var sr resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &sr)
	resp := &resource.ImportStateResponse{State: tfsdk.State{
		Schema: sr.Schema,
		Raw:    tftypes.NewValue(sr.Schema.Type().TerraformType(ctx), nil),
	}}
	importer.ImportState(ctx, resource.ImportStateRequest{ID: id}, resp)
	return resp.Diagnostics.HasError()
}
