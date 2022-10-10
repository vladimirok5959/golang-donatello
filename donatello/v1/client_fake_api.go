package v1

import (
	"context"
	"net/http"
)

type ClientFakeAPI struct {
	MockMe      func() (int64, []byte, error)
	MockDonates func(page, size int64) (int64, []byte, error)
	MockClients func() (int64, []byte, error)
}

func NewClientFakeAPI() *ClientFakeAPI {
	return &ClientFakeAPI{
		MockMe: func() (int64, []byte, error) {
			return http.StatusUnauthorized, []byte(`{"success":false,"message":"Помилка авторизації"}`), nil
		},
		MockDonates: func(page, size int64) (int64, []byte, error) {
			return http.StatusUnauthorized, []byte(`{"success":false,"message":"Помилка авторизації"}`), nil
		},
		MockClients: func() (int64, []byte, error) {
			return http.StatusUnauthorized, []byte(`{"success":false,"message":"Помилка авторизації"}`), nil
		},
	}
}

func (c *ClientFakeAPI) Me(ctx context.Context) (int64, []byte, error) {
	return c.MockMe()
}

func (c *ClientFakeAPI) Donates(ctx context.Context, page, size int64) (int64, []byte, error) {
	return c.MockDonates(page, size)
}

func (c *ClientFakeAPI) Clients(ctx context.Context) (int64, []byte, error) {
	return c.MockClients()
}
