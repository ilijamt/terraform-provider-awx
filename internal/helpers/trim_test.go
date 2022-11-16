package helpers_test

import (
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrimString(t *testing.T) {
	require.EqualValues(t, "", helpers.TrimAwxString(" \n"))
	require.EqualValues(t, "test\ntes", helpers.TrimAwxString(" test\ntes\n"))
}
