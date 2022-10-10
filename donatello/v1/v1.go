package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

type API interface {
	Me(ctx context.Context) (int64, []byte, error)
	Donates(ctx context.Context, page, size int64) (int64, []byte, error)
	Clients(ctx context.Context) (int64, []byte, error)
}

type Client struct {
	api API
}

func NewClient(api API) *Client {
	return &Client{
		api: api,
	}
}

func (c *Client) Me(ctx context.Context) (*ResponseMe, error) {
	var response ResponseMe

	status, result, err := c.api.Me(ctx)
	if err != nil {
		return &response, err
	}

	if err := json.Unmarshal(result, &response); err != nil {
		return &response, err
	}

	if status == http.StatusOK {
		response.Success = true
	}

	return &response, nil
}

func (c *Client) Donates(ctx context.Context, page, size int64) (*ResponseDonates, error) {
	var response ResponseDonates

	status, result, err := c.api.Donates(ctx, page, size)
	if err != nil {
		return &response, err
	}

	if err := json.Unmarshal(result, &response); err != nil {
		return &response, err
	}

	if status == http.StatusOK {
		response.Success = true
	}

	return &response, nil
}

func (c *Client) Clients(ctx context.Context) (*ResponseClients, error) {
	var response ResponseClients

	status, result, err := c.api.Clients(ctx)
	if err != nil {
		return &response, err
	}

	if err := json.Unmarshal(result, &response); err != nil {
		return &response, err
	}

	if status == http.StatusOK {
		response.Success = true
	}

	return &response, nil
}
