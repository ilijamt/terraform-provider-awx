package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

type clientWithBasicAuth struct {
	client *http.Client

	lock                         *sync.Mutex
	user                         *models.User
	username, password, hostname string
	version                      string
}

func (c *clientWithBasicAuth) User(ctx context.Context) (user models.User, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.user == nil {
		if c.user, err = UserId(ctx, c); err != nil {
			return models.User{}, err
		}
	}

	return *c.user, nil
}

var _ Client = &clientWithBasicAuth{}

func NewClientWithBasicAuth(username, password, hostname string, version string, insecureSkipVerify bool, httpClient *http.Client) Client {
	return &clientWithBasicAuth{
		client:   defaultClient(httpClient, insecureSkipVerify),
		hostname: hostname,
		username: username,
		password: password,
		version:  version,
		lock:     new(sync.Mutex),
	}
}

func (c *clientWithBasicAuth) NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	req, err = http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.hostname, endpoint), body)
	if err == nil {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", fmt.Sprintf("terraform-provider-awx/%s", c.version))
	}
	return req, err
}

func (c *clientWithBasicAuth) Do(ctx context.Context, req *http.Request) (data map[string]any, err error) {
	return DoRequest(ctx, c.client, req)
}
