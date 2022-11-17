package helpers_test

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	var obj = helpers.DefaultValue(types.StringValue("default"))
	require.NotEmpty(t, obj.Description(context.Background()))

	t.Run("attribute config is set", func(t *testing.T) {
		var req = tfsdk.ModifyAttributePlanRequest{
			AttributeConfig: types.StringValue("config"),
		}
		require.False(t, req.AttributeConfig.IsNull())
		obj.Modify(context.Background(),
			req,
			&tfsdk.ModifyAttributePlanResponse{},
		)
	})

	t.Run("attribute plan is set", func(t *testing.T) {
		var req = tfsdk.ModifyAttributePlanRequest{
			AttributeConfig: types.StringValue("config"),
			AttributePlan:   types.StringValue("plan"),
		}
		require.False(t, req.AttributeConfig.IsNull())
		require.False(t, req.AttributePlan.IsUnknown())
		require.False(t, req.AttributePlan.IsNull())
		obj.Modify(context.Background(),
			req,
			&tfsdk.ModifyAttributePlanResponse{},
		)
	})

	t.Run("return default value", func(t *testing.T) {
		var res = tfsdk.ModifyAttributePlanResponse{}
		obj.Modify(context.Background(),
			tfsdk.ModifyAttributePlanRequest{
				AttributeConfig: types.StringNull(),
				AttributePlan:   types.StringUnknown(),
			},
			&res,
		)
		require.NotEmpty(t, res.AttributePlan)
		require.EqualValues(t, res.AttributePlan, types.StringValue("default"))
	})
}
