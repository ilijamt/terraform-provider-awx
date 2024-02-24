package helpers_test

import (
	"testing"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestTrimString(t *testing.T) {
	require.EqualValues(t, "", helpers.TrimAwxString(" \n"))
	require.EqualValues(t, "test\ntes", helpers.TrimAwxString(" test\ntes\n"))
}
