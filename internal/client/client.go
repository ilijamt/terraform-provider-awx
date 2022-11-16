package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ilijamt/terraform-provider-awx/config"
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
	insecureSkipVerify           bool
}

var _ Client = &client{}

type Client interface {
	NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Request, error)
	Do(ctx context.Context, req *http.Request) (data map[string]any, err error)
}

func NewClient(username, password, hostname string, insecureSkipVerify bool) Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}

	return &client{
		client: &http.Client{
			Transport: tr,
		},
		hostname: hostname,
		username: username,
		password: password,
	}
}

func (c *client) NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	req, err = http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.hostname, endpoint), body)
	if err == nil {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", fmt.Sprintf("terraform-provider-awx/%s", config.Version))
	}
	return req, err
}

func doRequest(client *http.Client, ctx context.Context, req *http.Request) (data map[string]any, err error) {
	var resp *http.Response
	if resp, err = client.Do(req.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("%w: failed to do request", err)
	}
	defer resp.Body.Close()

	var payload []byte

	if payload, err = io.ReadAll(resp.Body); err != nil {
		return data, err
	}

	dec := json.NewDecoder(bytes.NewReader(payload))
	dec.UseNumber()
	if err = dec.Decode(&data); err != nil && !errors.Is(err, io.EOF) {
		return data, fmt.Errorf("%w: failed to decode data", err)
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return data, fmt.Errorf("%w: %d, on %s with %s", ErrInvalidStatusCode, resp.StatusCode, req.URL.RequestURI(), string(payload))
	}

	return data, nil
}

func (c *client) Do(ctx context.Context, req *http.Request) (data map[string]any, err error) {
	return doRequest(c.client, ctx, req)
}
