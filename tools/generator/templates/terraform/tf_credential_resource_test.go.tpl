package {{ .PackageName }}_test


import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
)

func TestNew{{ $.TypeName | pascalCase }}CredentialResource(t *testing.T) {

    t.Run("config validators should not be empty", func(t *testing.T) {
        obj := awx.New{{ $.TypeName | pascalCase }}CredentialResource()
        require.NotNil(t, obj)
        validators := obj.(resource.ResourceWithConfigValidators).ConfigValidators(t.Context())
        require.NotEmpty(t, validators)
	})
}
