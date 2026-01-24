package client

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	tst "testing"

	"go.uber.org/mock/gomock"

	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
	ErrJsonDecode        = errors.New("json decode")
)

type Client interface {
	NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Request, error)
	Do(ctx context.Context, req *http.Request) (data map[string]any, err error)
	User(ctx context.Context) (_ models.User, err error)
}

func NewTestingClient(t *tst.T) (client *MockClient, ctrl *gomock.Controller) {
	ctrl = gomock.NewController(t)
	client = NewMockClient(ctrl)
	return client, ctrl
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
