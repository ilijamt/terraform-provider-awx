package helpers_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestProcessJsonEncryptedValues(t *testing.T) {
	var tests = []struct {
		orig, cur, msg types.String
		dirty          bool
		err            bool
	}{
		{
			orig:  types.String{},
			cur:   types.String{},
			msg:   types.String{},
			dirty: false,
			err:   true,
		},
		{
			orig:  types.StringValue("{}"),
			cur:   types.StringValue("{}"),
			msg:   types.StringValue("{}"),
			dirty: false,
			err:   false,
		},
		{
			orig:  types.StringValue(`{"a":"password","b":"value"}`),
			cur:   types.StringValue(`{"a":"$encrypted$","b":"value"}`),
			msg:   types.StringValue(`{"a":"password","b":"value"}`),
			dirty: true,
			err:   false,
		},
		{
			orig:  types.StringValue(`{"b":"value"}`),
			cur:   types.StringValue(`{"a":"$encrypted$","b":"value"}`),
			msg:   types.StringValue(`{"a":"$encrypted$","b":"value"}`),
			dirty: false,
			err:   true,
		},
		{
			cur:   types.StringValue(`{"b":"value"}`),
			orig:  types.StringValue(`{"a":"$encrypted$","b":"value"}`),
			msg:   types.StringValue(`{"b":"value"}`),
			dirty: false,
			err:   false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%v, %v) should be (%t, %v, %v)", test.orig, test.cur, test.dirty, test.msg, test.err), func(t *testing.T) {
			dirty, msg, err := helpers.ProcessJsonEncryptedValues(test.orig, test.cur)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, test.msg, msg)
			assert.EqualValues(t, test.dirty, dirty)
		})
	}

}
