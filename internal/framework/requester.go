package framework

import (
	"context"
	"io"
	"net/http"
)

// Requester is the minimal interface for making HTTP requests to the AWX API.
// client.Client satisfies this interface implicitly.
type Requester interface {
	NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Request, error)
	Do(ctx context.Context, req *http.Request) (map[string]any, error)
}
