package client_test

import (
	"context"
	"fmt"
	"github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticationClient(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusUnauthorized)
	}))

	type test struct {
		method string
		err    error
	}

	var tests = []test{
		{method: http.MethodGet, err: client.ErrInvalidStatusCode},
		{method: http.MethodPost, err: client.ErrInvalidStatusCode},
		{method: http.MethodDelete, err: client.ErrInvalidStatusCode},
		{method: http.MethodPatch, err: client.ErrInvalidStatusCode},
	}

	c := client.NewClient("username", "password", server.URL, "test", true, nil)

	for _, tst := range tests {
		t.Run(tst.method, func(t *testing.T) {
			req, err := c.NewRequest(context.Background(), http.MethodGet, "/api/v2/request", nil)
			require.NoError(t, err)
			require.NotNil(t, req)
			data, err := c.Do(context.Background(), req)
			require.ErrorIs(t, err, tst.err)
			require.Empty(t, data)
		})
	}
}

func TestBody(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		method string
		err    error
	}

	var tests = []test{
		{name: "no content", method: http.MethodGet, err: nil},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Header.Get("test-x-type") {
		case "no content":
			rw.WriteHeader(http.StatusNoContent)

		}
	}))

	c := client.NewClient("username", "password", server.URL, "test", true, nil)
	for _, tst := range tests {
		t.Run(fmt.Sprintf("%s - %s", tst.name, tst.method), func(t *testing.T) {
			req, err := c.NewRequest(context.Background(), http.MethodGet, "/api/v2/request", nil)
			require.NoError(t, err)
			require.NotNil(t, req)
			req.Header.Set("test-x-type", tst.name)
			data, err := c.Do(context.Background(), req)
			require.ErrorIs(t, err, tst.err)
			require.Empty(t, data)
		})
	}

}
