package helpers_test

import (
	"fmt"
	"github.com/ilijamt/envwrap"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFirstSetEnvVar(t *testing.T) {
	var tests = []struct {
		in  []string
		env map[string]string
		out string
	}{
		{in: []string{"VAR1", "VAR2", "VAR3"}, env: make(map[string]string), out: ""},
		{in: []string{"VAR1", "VAR2", "VAR3"}, env: map[string]string{"VAR1": "VAR1"}, out: "VAR1"},
		{in: []string{"VAR1", "VAR2", "VAR3"}, env: map[string]string{"VAR2": "VAR2"}, out: "VAR2"},
		{in: []string{"VAR1", "VAR2", "VAR3"}, env: map[string]string{"VAR3": "VAR3"}, out: "VAR3"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s should be %s", test.in, test.out), func(t *testing.T) {
			env := envwrap.NewStorage()

			for k, v := range test.env {
				_ = env.Store(k, v)
			}

			assert.EqualValues(t, test.out, helpers.GetFirstSetEnvVar(test.in...))
			_ = env.ReleaseAll()
		})
	}

}
