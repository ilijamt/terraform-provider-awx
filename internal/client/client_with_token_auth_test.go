package client_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func TestNewClientWithTokenAuth(t *testing.T) {
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

	c := client.NewClientWithTokenAuth("token", server.URL, "test", true, nil)

	for _, tst := range tests {
		t.Run(tst.method, func(t *testing.T) {
			req, err := c.NewRequest(t.Context(), http.MethodGet, "/api/v2/request", nil)
			require.NoError(t, err)
			require.NotNil(t, req)
			data, err := c.Do(t.Context(), req)
			require.ErrorIs(t, err, tst.err)
			require.Empty(t, data)
		})
	}
}

func TestNewClientWithTokenAuthBody(t *testing.T) {
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

	c := client.NewClientWithTokenAuth("token", server.URL, "test", true, nil)
	for _, tst := range tests {
		t.Run(fmt.Sprintf("%s - %s", tst.name, tst.method), func(t *testing.T) {
			req, err := c.NewRequest(t.Context(), http.MethodGet, "/api/v2/request", nil)
			require.NoError(t, err)
			require.NotNil(t, req)
			req.Header.Set("test-x-type", tst.name)
			data, err := c.Do(t.Context(), req)
			require.ErrorIs(t, err, tst.err)
			require.Empty(t, data)
		})
	}

}
