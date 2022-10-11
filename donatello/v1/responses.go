package v1

import (
	"strconv"
	"time"
)

type ResponseMeDonates struct {
	TotalAmount int64 `json:"totalAmount"`
	TotalCount  int64 `json:"totalCount"`
}

type ResponseMe struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	NickName  string            `json:"nickname"`
	PubID     string            `json:"pubId"`
	Page      string            `json:"page"`
	IsActive  bool              `json:"isActive"`
	IsPublic  bool              `json:"isPublic"`
	Donates   ResponseMeDonates `json:"donates"`
	CreatedAt string            `json:"createdAt"`
}

func (r ResponseMe) CreatedAtTime() time.Time {
	result, err := time.Parse("2006-01-02 15:04:05", r.CreatedAt)
	if err != nil {
		return time.Time{}
	}
	return result
}

// -----------------------------------------------------------------------------

type ResponseDonatesContent struct {
	PubID       string `json:"pubId"`
	ClientName  string `json:"clientName"`
	Message     string `json:"message"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Goal        string `json:"goal"`
	IsPublished bool   `json:"isPublished"`
	CreatedAt   string `json:"createdAt"`
}

func (r ResponseDonatesContent) AmountInt64() int64 {
	var result int64

	var err error
	if result, err = strconv.ParseInt(r.Amount, 10, 64); err != nil {
		result = 0
	}

	return result
}

func (r ResponseDonatesContent) CreatedAtTime() time.Time {
	result, err := time.Parse("2006-01-02 15:04:05", r.CreatedAt)
	if err != nil {
		return time.Time{}
	}
	return result
}

type ResponseDonates struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	Content []ResponseDonatesContent `json:"content"`
	Page    int64                    `json:"page"`
	Size    int64                    `json:"size"`
	Pages   int64                    `json:"pages"`
	First   bool                     `json:"first"`
	Last    bool                     `json:"last"`
	Total   int64                    `json:"total"`
}

// -----------------------------------------------------------------------------

type ResponseClientsClients struct {
	ClientName  string `json:"clientName"`
	TotalAmount int64  `json:"totalAmount"`
}

type ResponseClients struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	Clients []ResponseClientsClients `json:"clients"`
}
