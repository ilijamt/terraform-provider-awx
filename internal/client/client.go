package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
	ErrJsonDecode        = errors.New("json decode")
)

type client struct {
	client *http.Client

	username, password, hostname string
	version                      string
}

var _ Client = &client{}

type Client interface {
	NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Request, error)
	Do(ctx context.Context, req *http.Request) (data map[string]any, err error)
}

func NewClient(username, password, hostname string, version string, insecureSkipVerify bool, httpClient *http.Client) Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Transport: tr,
		}
	}

	return &client{
		client:   httpClient,
		hostname: hostname,
		username: username,
		password: password,
		version:  version,
	}
}

func (c *client) NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	req, err = http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.hostname, endpoint), body)
	if err == nil {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", fmt.Sprintf("terraform-provider-awx/%s", c.version))
	}
	return req, err
}

func (c *client) Do(ctx context.Context, req *http.Request) (data map[string]any, err error) {
	return doRequest(c.client, ctx, req)
}
