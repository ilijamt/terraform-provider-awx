package client_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func TestNewClientWithBasicAuth(t *testing.T) {
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

	c := client.NewClientWithBasicAuth("username", "password", server.URL, "test", true, nil)

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

func TestNewClientWithBasicAuthBody(t *testing.T) {
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

	c := client.NewClientWithBasicAuth("username", "password", server.URL, "test", true, nil)
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

func TestStatusErrorIdentifiesNotFound(t *testing.T) {
	notFound := &client.StatusError{StatusCode: http.StatusNotFound, URI: "/api/v2/tokens/1/", Body: "gone"}
	forbidden := &client.StatusError{StatusCode: http.StatusForbidden, URI: "/api/v2/tokens/1/", Body: "nope"}

	assert.True(t, client.IsNotFound(notFound))
	assert.False(t, client.IsNotFound(forbidden))
	assert.False(t, client.IsNotFound(errors.New("boom")))

	// Existing call sites match on the sentinel.
	assert.True(t, errors.Is(notFound, client.ErrInvalidStatusCode))
	assert.Contains(t, notFound.Error(), "invalid status code: 404")

	assert.True(t, client.IsNotFound(fmt.Errorf("wrapped: %w", notFound)))
}

// AWX serves 500s as an HTML page. Decoding before checking the status buried
// that under a JSON parse error, which is what "invalid character '<'" was.
func TestHtmlErrorPageReportsStatusNotDecodeFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<!doctype html>\n<html><head><title>Server Error (500)</title></head></html>"))
	}))
	defer srv.Close()

	cl := client.NewClientWithBasicAuth("u", "p", srv.URL, "test", true, nil)
	req, err := cl.NewRequest(context.Background(), http.MethodPost, "/api/v2/teams/5/roles/", nil)
	require.NoError(t, err)

	_, err = cl.Do(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid status code: 500")
	assert.NotContains(t, err.Error(), "failed to decode data")

	var se *client.StatusError
	require.True(t, errors.As(err, &se))
	assert.Equal(t, http.StatusInternalServerError, se.StatusCode)
}
