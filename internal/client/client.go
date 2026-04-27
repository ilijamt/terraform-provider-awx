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

// preserveMethodOnRedirect follows redirects but restores the original method,
// body, and key request headers. Go's default policy rewrites 301/302/303 on
// POST/PUT/PATCH/DELETE to GET and drops the body, which silently turned writes
// against an http:// AWX into list reads (308 → POST https://, then Django
// APPEND_SLASH 301 → GET). Cross-scheme redirects also strip Authorization, and
// the implicit POST→GET conversion can drop Content-Type. We re-apply all of
// them so an http:// → https:// upgrade or a missing trailing slash just works.
func preserveMethodOnRedirect(req *http.Request, via []*http.Request) error {
	if len(via) == 0 {
		return nil
	}
	original := via[0]

	for _, key := range []string{"Authorization", "Content-Type", "User-Agent"} {
		if val := original.Header.Get(key); val != "" {
			req.Header.Set(key, val)
		}
	}

	if req.Method == original.Method {
		return nil
	}
	req.Method = original.Method
	if original.GetBody != nil {
		body, err := original.GetBody()
		if err != nil {
			return err
		}
		req.Body = body
		req.ContentLength = original.ContentLength
	}
	return nil
}

func defaultClient(client *http.Client, insecureSkipVerify bool) *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}

	if client == nil {
		client = &http.Client{
			Transport:     tr,
			CheckRedirect: preserveMethodOnRedirect,
		}
	} else if client.CheckRedirect == nil {
		client.CheckRedirect = preserveMethodOnRedirect
	}

	return client
}
