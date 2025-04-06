package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type clientWithTokenAuth struct {
	client *http.Client

	token, hostname string
	version         string
}

var _ Client = &clientWithTokenAuth{}

func NewClientWithTokenAuth(token, hostname string, version string, insecureSkipVerify bool, httpClient *http.Client) Client {
	return &clientWithTokenAuth{
		client:   defaultClient(httpClient, insecureSkipVerify),
		hostname: hostname,
		token:    token,
		version:  version,
	}
}

func (c *clientWithTokenAuth) NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	req, err = http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.hostname, endpoint), body)
	if err == nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", fmt.Sprintf("terraform-provider-awx/%s", c.version))
	}
	return req, err
}

func (c *clientWithTokenAuth) Do(ctx context.Context, req *http.Request) (data map[string]any, err error) {
	return DoRequest(c.client, ctx, req)
}
