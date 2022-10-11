package v1

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ClientAPI struct {
	timeout time.Duration
	token   string
}

func NewClientAPI(timeout time.Duration, token string) *ClientAPI {
	return &ClientAPI{
		timeout: timeout,
		token:   token,
	}
}

func (c *ClientAPI) request(ctx context.Context, url string) (int64, []byte, error) {
	var result []byte

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, result, err
	}
	req.Header = map[string][]string{
		"X-Token": {c.token},
	}

	rcl := &http.Client{}
	var resp *http.Response
	resp, err = rcl.Do(req)
	if err != nil {
		return 0, result, err
	}
	defer resp.Body.Close()

	result, err = io.ReadAll(resp.Body)
	if err != nil {
		return int64(resp.StatusCode), result, err
	}

	return int64(resp.StatusCode), result, nil
}

func (c *ClientAPI) Me(ctx context.Context) (int64, []byte, error) {
	return c.request(ctx, API_URL_ME)
}

func (c *ClientAPI) Donates(ctx context.Context, page, size int64) (int64, []byte, error) {
	params := url.Values{}
	params.Add("page", strconv.FormatInt(page, 10))
	params.Add("size", strconv.FormatInt(size, 10))
	return c.request(ctx, API_URL_DONATES+"?"+params.Encode())
}

func (c *ClientAPI) Clients(ctx context.Context) (int64, []byte, error) {
	return c.request(ctx, API_URL_CLIENTS)
}
