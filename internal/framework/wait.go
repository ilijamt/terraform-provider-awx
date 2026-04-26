package framework

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"
)

const (
	defaultPollInterval = 5 * time.Second
	maxPollInterval     = 30 * time.Second
)

// WaitForFieldOpts configures WaitForFieldValue.
type WaitForFieldOpts struct {
	// Endpoint is the resource URL to GET. Required.
	Endpoint string
	// Field is the JSON field on the response to inspect. Required, must be a string field.
	Field string
	// SuccessValues are terminal values that mean the wait succeeded.
	SuccessValues []string
	// FailureValues are terminal values that mean the wait failed.
	FailureValues []string
	// PollInterval is how long to sleep between polls. Defaults to 5s; capped at 30s.
	PollInterval time.Duration
}

// WaitTerminalError is returned when the polled field reaches a value listed in FailureValues.
type WaitTerminalError struct {
	// Status is the terminal value seen on the field.
	Status string
}

func (e *WaitTerminalError) Error() string {
	return fmt.Sprintf("reached terminal failure status %q", e.Status)
}

// WaitForFieldValue polls opts.Endpoint until opts.Field reaches a value in
// SuccessValues (returns nil), reaches a value in FailureValues (returns
// *WaitTerminalError), or ctx is canceled. The caller is responsible for
// providing a ctx with a timeout.
func WaitForFieldValue(ctx context.Context, client Requester, opts WaitForFieldOpts) error {
	interval := opts.PollInterval
	if interval <= 0 {
		interval = defaultPollInterval
	}
	if interval > maxPollInterval {
		interval = maxPollInterval
	}

	for {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("waiting for field %q: %w", opts.Field, err)
		}

		req, err := client.NewRequest(ctx, http.MethodGet, opts.Endpoint, nil)
		if err != nil {
			return fmt.Errorf("building poll request for %s: %w", opts.Endpoint, err)
		}
		data, err := client.Do(ctx, req)
		if err != nil {
			return fmt.Errorf("polling %s: %w", opts.Endpoint, err)
		}

		raw, ok := data[opts.Field]
		if !ok || raw == nil {
			return fmt.Errorf("polling response missing field %q", opts.Field)
		}
		status, ok := raw.(string)
		if !ok {
			return fmt.Errorf("polling field %q is not a string (got %T)", opts.Field, raw)
		}

		if slices.Contains(opts.SuccessValues, status) {
			return nil
		}
		if slices.Contains(opts.FailureValues, status) {
			return &WaitTerminalError{Status: status}
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("waiting for field %q: %w", opts.Field, ctx.Err())
		case <-time.After(interval):
		}
	}
}
