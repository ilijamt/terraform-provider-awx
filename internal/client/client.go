package client

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
)

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
	ErrJsonDecode        = errors.New("json decode")
)

type Client interface {
	NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Request, error)
	Do(ctx context.Context, req *http.Request) (data map[string]any, err error)
}

func defaultClient(client *http.Client, insecureSkipVerify bool) *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}

	if client == nil {
		client = &http.Client{
			Transport: tr,
		}
	}

	return client
}
