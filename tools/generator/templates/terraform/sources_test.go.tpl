package {{ .PackageName }}_test

import (
    "github.com/ilijamt/terraform-provider-awx/internal/{{ .PackageName }}"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDataSources(t *testing.T) {
    require.NotEmpty(t, {{ .PackageName }}.DataSources()());
}

func TestResources(t *testing.T) {
    require.NotEmpty(t, {{ .PackageName }}.Resources()());
}
