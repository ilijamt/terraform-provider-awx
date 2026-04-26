package framework_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// scriptedRequester returns a sequence of responses for successive GETs.
type scriptedRequester struct {
	responses []map[string]any
	idx       int32
}

func (s *scriptedRequester) NewRequest(_ context.Context, _, _ string, _ io.Reader) (*http.Request, error) {
	return &http.Request{}, nil
}

func (s *scriptedRequester) Do(_ context.Context, _ *http.Request) (map[string]any, error) {
	i := atomic.AddInt32(&s.idx, 1) - 1
	if int(i) >= len(s.responses) {
		i = int32(len(s.responses) - 1)
	}
	return s.responses[i], nil
}

func TestWaitForFieldValue(t *testing.T) {
	t.Parallel()

	const field = "status"
	successVals := []string{"successful", "ok"}
	failureVals := []string{"failed", "error"}

	tests := []struct {
		name      string
		responses []map[string]any
		ctxFn     func() (context.Context, context.CancelFunc)
		wantErr   bool
		wantTerm  string // terminal failure status if WaitTerminalError expected
	}{
		{
			name:      "immediate success",
			responses: []map[string]any{{"status": "successful"}},
		},
		{
			name: "transition then success",
			responses: []map[string]any{
				{"status": "pending"},
				{"status": "running"},
				{"status": "successful"},
			},
		},
		{
			name:      "ok also succeeds",
			responses: []map[string]any{{"status": "ok"}},
		},
		{
			name:      "terminal failure",
			responses: []map[string]any{{"status": "failed"}},
			wantErr:   true,
			wantTerm:  "failed",
		},
		{
			name:      "terminal error after pending",
			responses: []map[string]any{{"status": "pending"}, {"status": "error"}},
			wantErr:   true,
			wantTerm:  "error",
		},
		{
			name:      "missing field",
			responses: []map[string]any{{"other": "value"}},
			wantErr:   true,
		},
		{
			name:      "null field",
			responses: []map[string]any{{"status": nil}},
			wantErr:   true,
		},
		{
			name:      "non-string field",
			responses: []map[string]any{{"status": 42}},
			wantErr:   true,
		},
		{
			name:      "ctx cancelled while pending",
			responses: []map[string]any{{"status": "pending"}},
			ctxFn: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 25*time.Millisecond)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if tt.ctxFn != nil {
				ctx, cancel = tt.ctxFn()
				defer cancel()
			}

			r := &scriptedRequester{responses: tt.responses}
			err := framework.WaitForFieldValue(ctx, r, framework.WaitForFieldOpts{
				Endpoint:      "/api/v2/projects/1/",
				Field:         field,
				SuccessValues: successVals,
				FailureValues: failureVals,
				PollInterval:  5 * time.Millisecond,
			})

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tt.wantTerm != "" {
					var term *framework.WaitTerminalError
					if !errors.As(err, &term) {
						t.Fatalf("expected *WaitTerminalError, got %T (%v)", err, err)
					}
					if term.Status != tt.wantTerm {
						t.Fatalf("terminal status: got %q, want %q", term.Status, tt.wantTerm)
					}
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestWaitForFieldValue_RequesterError(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r := failDo()
	err := framework.WaitForFieldValue(ctx, r, framework.WaitForFieldOpts{
		Endpoint:      "/api/v2/projects/1/",
		Field:         "status",
		SuccessValues: []string{"successful"},
		PollInterval:  5 * time.Millisecond,
	})
	if err == nil {
		t.Fatalf("expected error from requester")
	}
}
