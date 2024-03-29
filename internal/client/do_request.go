package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func doRequest(client *http.Client, ctx context.Context, req *http.Request) (data map[string]any, err error) {
	if client == nil {
		return data, fmt.Errorf("nil http clientWithBasicAuth")
	}

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
