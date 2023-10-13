package helpers_test

import (
	"fmt"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStr2Bool(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{in: "True", out: true},
		{in: "TrUe", out: true},
		{in: "true", out: true},
		{in: "yes", out: false},
		{in: "no", out: false},
		{in: "False", out: false},
		{in: "FalSe", out: false},
		{in: "false", out: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s should be %t", test.in, test.out), func(t *testing.T) {
			assert.EqualValues(t, test.out, helpers.Str2Bool(test.in))
		})
	}
}
