package client_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func TestUserId(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		user, err := client.UserId(t.Context(), nil)
		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("valid user", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			payload, err := os.ReadFile("../../testdata/me.json")
			require.NoError(t, err)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(payload)
		}))
		t.Cleanup(svr.Close)

		t.Run("basic auth", func(t *testing.T) {
			c := client.NewClientWithBasicAuth("username", "password", svr.URL, "test", true, nil)
			user, err := c.User(t.Context())
			require.NoError(t, err)
			require.NotNil(t, user)
			require.Equal(t, 1, user.ID)
		})

		t.Run("token auth", func(t *testing.T) {
			c := client.NewClientWithTokenAuth("token", svr.URL, "test", true, nil)
			user, err := c.User(t.Context())
			require.NoError(t, err)
			require.NotNil(t, user)
			require.Equal(t, 1, user.ID)
		})
	})

	t.Run("invalid json", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{ invalid json`))
		}))
		t.Cleanup(svr.Close)

		c := client.NewClientWithBasicAuth("username", "password", svr.URL, "test", true, nil)
		user, err := client.UserId(t.Context(), c)
		require.Error(t, err)
		require.Nil(t, user)
	})

}
