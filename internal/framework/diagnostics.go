package framework

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// DiagnosticsHasError appends diagnostics and returns true if there are errors.
func DiagnosticsHasError(target *diag.Diagnostics, d ...diag.Diagnostic) bool {
	target.Append(d...)
	return target.HasError()
}

// HookError adds a diagnostic error if err is non-nil and returns true.
func HookError(target *diag.Diagnostics, resourceName string, err error) bool {
	if err == nil {
		return false
	}
	target.AddError(
		fmt.Sprintf("Unable to process custom hook for the state on %s", resourceName),
		err.Error(),
	)
	return true
}
