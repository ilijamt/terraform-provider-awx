package hooks

import (
	"context"
)

type Hook[T any] func(ctx context.Context, apiVersion string, source Source, callee Callee, orig, state T) (err error)
