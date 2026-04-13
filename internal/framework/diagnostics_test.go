package framework_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestDiagnosticsHasError(t *testing.T) {
	tests := []struct {
		name           string
		initial        diag.Diagnostics
		append         []diag.Diagnostic
		expectedReturn bool
		expectedLen    int
	}{
		{
			name:           "empty diagnostics no append",
			append:         nil,
			expectedReturn: false,
			expectedLen:    0,
		},
		{
			name:           "append error returns true",
			append:         []diag.Diagnostic{diag.NewErrorDiagnostic("err", "detail")},
			expectedReturn: true,
			expectedLen:    1,
		},
		{
			name:           "append warning returns false",
			append:         []diag.Diagnostic{diag.NewWarningDiagnostic("warn", "detail")},
			expectedReturn: false,
			expectedLen:    1,
		},
		{
			name:           "append error to existing warning",
			initial:        diag.Diagnostics{diag.NewWarningDiagnostic("existing", "warning")},
			append:         []diag.Diagnostic{diag.NewErrorDiagnostic("new", "error")},
			expectedReturn: true,
			expectedLen:    2,
		},
		{
			name:           "existing error append warning still true",
			initial:        diag.Diagnostics{diag.NewErrorDiagnostic("existing", "error")},
			append:         []diag.Diagnostic{diag.NewWarningDiagnostic("new", "warning")},
			expectedReturn: true,
			expectedLen:    2,
		},
		{
			name: "multiple appended one error",
			append: []diag.Diagnostic{
				diag.NewWarningDiagnostic("warn", ""),
				diag.NewErrorDiagnostic("err", ""),
			},
			expectedReturn: true,
			expectedLen:    2,
		},
		{
			name: "multiple appended no errors",
			append: []diag.Diagnostic{
				diag.NewWarningDiagnostic("w1", ""),
				diag.NewWarningDiagnostic("w2", ""),
			},
			expectedReturn: false,
			expectedLen:    2,
		},
		{
			name:           "append nothing to existing error",
			initial:        diag.Diagnostics{diag.NewErrorDiagnostic("existing", "error")},
			append:         nil,
			expectedReturn: true,
			expectedLen:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.initial
			got := framework.DiagnosticsHasError(&d, tt.append...)
			assert.Equal(t, tt.expectedReturn, got)
			assert.Len(t, d, tt.expectedLen)
		})
	}
}

func TestHookError(t *testing.T) {
	tests := []struct {
		name            string
		resourceName    string
		err             error
		expectedReturn  bool
		expectedDiagLen int
		expectedSummary string
		expectedDetail  string
	}{
		{
			name:            "nil error returns false",
			resourceName:    "TestResource",
			err:             nil,
			expectedReturn:  false,
			expectedDiagLen: 0,
		},
		{
			name:            "non-nil error returns true",
			resourceName:    "TestResource",
			err:             errors.New("hook failed"),
			expectedReturn:  true,
			expectedDiagLen: 1,
			expectedSummary: "Unable to process custom hook for the state on TestResource",
			expectedDetail:  "hook failed",
		},
		{
			name:            "empty resource name",
			resourceName:    "",
			err:             errors.New("fail"),
			expectedReturn:  true,
			expectedDiagLen: 1,
			expectedSummary: "Unable to process custom hook for the state on ",
			expectedDetail:  "fail",
		},
		{
			name:            "wrapped error preserves message",
			resourceName:    "Res",
			err:             fmt.Errorf("wrap: %w", errors.New("inner")),
			expectedReturn:  true,
			expectedDiagLen: 1,
			expectedSummary: "Unable to process custom hook for the state on Res",
			expectedDetail:  "wrap: inner",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d diag.Diagnostics
			got := framework.HookError(&d, tt.resourceName, tt.err)
			assert.Equal(t, tt.expectedReturn, got)
			assert.Len(t, d, tt.expectedDiagLen)
			if tt.expectedDiagLen > 0 {
				assert.Equal(t, tt.expectedSummary, d[0].Summary())
				assert.Equal(t, tt.expectedDetail, d[0].Detail())
			}
		})
	}
}
