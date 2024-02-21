package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"
)

func TestDoRequest(t *testing.T) {

	t.Run("nil client should error out", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "url", nil)
		require.NoError(t, err)
		require.NotNil(t, req)
		data, err := doRequest(nil, context.Background(), req)
		require.Error(t, err)
		require.ErrorContains(t, err, "nil http client")
		require.Empty(t, data)
	})

	t.Run("io stream error", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer svr.Close()

		req, err := http.NewRequest(http.MethodPost, svr.URL, iotest.ErrReader(fmt.Errorf("io stream error")))
		require.NoError(t, err)
		require.NotNil(t, req)

		data, err := doRequest(http.DefaultClient, context.Background(), req)
		require.Error(t, err)
		require.ErrorContains(t, err, "io stream error")
		require.ErrorContains(t, err, "failed to do request")
		require.Empty(t, data)
	})

	t.Run("invalid payload", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(json.RawMessage(`{ t`))
		}))
		defer svr.Close()

		req, err := http.NewRequest(http.MethodGet, svr.URL, nil)
		require.NoError(t, err)
		require.NotNil(t, req)

		data, err := doRequest(http.DefaultClient, context.Background(), req)
		require.Error(t, err)
		require.ErrorContains(t, err, "failed to decode data")
		require.Empty(t, data)

	})

	t.Run("response io stream error", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Length", "1")
		}))
		defer svr.Close()

		req, err := http.NewRequest(http.MethodPost, svr.URL, nil)
		require.NoError(t, err)
		require.NotNil(t, req)

		data, err := doRequest(http.DefaultClient, context.Background(), req)
		require.Error(t, err)
		require.ErrorContains(t, err, "unexpected EOF")
		require.Empty(t, data)
	})
}
