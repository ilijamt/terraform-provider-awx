package helpers_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestIsEmptyValue(t *testing.T) {

	var tests = []struct {
		in  attr.Value
		out bool
	}{
		{types.StringNull(), true},
		{types.StringUnknown(), true},
		{types.NumberNull(), true},
		{types.NumberUnknown(), true},
		{types.NumberValue(big.NewFloat(1)), false},
		{types.StringValue("test"), false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s should be %t", test.in, test.out), func(t *testing.T) {
			assert.EqualValues(t, test.out, helpers.IsEmptyValue(test.in))
		})
	}

}
