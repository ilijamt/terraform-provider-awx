package framework_test

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// mockRequester implements framework.Requester for testing.
type mockRequester struct {
	newRequestFunc func(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error)
	doFunc         func(ctx context.Context, req *http.Request) (map[string]any, error)
}

var _ framework.Requester = (*mockRequester)(nil)

func (m *mockRequester) NewRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	return m.newRequestFunc(ctx, method, endpoint, body)
}

func (m *mockRequester) Do(ctx context.Context, req *http.Request) (map[string]any, error) {
	return m.doFunc(ctx, req)
}

func successRequester(data map[string]any) *mockRequester {
	return &mockRequester{
		newRequestFunc: func(context.Context, string, string, io.Reader) (*http.Request, error) {
			return &http.Request{}, nil
		},
		doFunc: func(context.Context, *http.Request) (map[string]any, error) { return data, nil },
	}
}

func failNewRequest() *mockRequester {
	return &mockRequester{
		newRequestFunc: func(context.Context, string, string, io.Reader) (*http.Request, error) {
			return nil, fmt.Errorf("request error")
		},
	}
}

func failDo() *mockRequester {
	return &mockRequester{
		newRequestFunc: func(context.Context, string, string, io.Reader) (*http.Request, error) {
			return &http.Request{}, nil
		},
		doFunc: func(context.Context, *http.Request) (map[string]any, error) {
			return nil, fmt.Errorf("do error")
		},
	}
}
