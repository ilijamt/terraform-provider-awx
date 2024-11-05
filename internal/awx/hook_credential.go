package awx

import (
	"context"
	"fmt"
	"strings"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookCredential(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *credentialTerraformModel) (err error) {
	if source == hooks.SourceResource && (state == nil || orig == nil) && callee != hooks.CalleeDelete {
		return fmt.Errorf("state and orig required for resource")
	} else if source == hooks.SourceData && (state == nil) {
		return fmt.Errorf("state is required for data source")
	}

	if source == hooks.SourceResource && callee == hooks.CalleeCreate {
		state.Inputs = types.StringValue(orig.Inputs.ValueString())
	} else if source == hooks.SourceResource && (callee == hooks.CalleeRead || callee == hooks.CalleeUpdate) {
		if !strings.Contains(state.Inputs.ValueString(), "$encrypted$") {
			return nil
		}
		if dirty, msg, err := helpers.ProcessJsonEncryptedValues(orig.Inputs, state.Inputs); err != nil {
			return err
		} else if dirty {
			state.Inputs = msg
		}
	}

	return nil
}
