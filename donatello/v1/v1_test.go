package v1_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/vladimirok5959/golang-donatello/donatello/v1"
	"golang.org/x/net/context"
)

const MockMeData = `{
	"nickname": "NickName",
	"pubId": "A1B-A123456",
	"page": "https://donatello.to/nickname",
	"isActive": true,
	"isPublic": true,
	"donates": {
		"totalAmount": 100,
		"totalCount": 1
	},
	"createdAt": "2022-10-20 00:30:50"
}`

const MockDonatesData = `{
	"content": [
		{
			"pubId": "A1B-A123456",
			"clientName": "ClientName",
			"message": "Message",
			"amount": "100",
			"currency": "UAH",
			"isPublished": false,
			"createdAt": "2022-10-20 00:30:50"
		}
	],
	"page": 1,
	"size": 20,
	"pages": 1,
	"first": true,
	"last": true,
	"total": 1
}`

const MockClientsData = `{
	"clients": [
		{
			"clientName": "ClientName",
			"totalAmount": 100
		}
	]
}`

var _ = Describe("Client", func() {
	Context("Func", func() {
		var api *v1.ClientFakeAPI
		var client *v1.Client
		var ctx context.Context

		BeforeEach(func() {
			api = v1.NewClientFakeAPI()
			client = v1.NewClient(api)
			ctx = context.Background()
		})

		Context("Me", func() {
			It("respond as unauthorized", func() {
				resp, err := client.Me(ctx)
				Expect(err).To(Succeed())
				Expect(resp.Success).To(BeFalse())
				Expect(resp.Message).To(Equal("Помилка авторизації"))
			})

			It("respond with correct data", func() {
				api.MockMe = func() (int64, []byte, error) {
					return http.StatusOK, []byte(MockMeData), nil
				}
				resp, err := client.Me(ctx)
				Expect(err).To(Succeed())
				Expect(resp.Success).To(BeTrue())
				Expect(resp.Message).To(Equal(""))
				Expect(resp.NickName).To(Equal("NickName"))
				Expect(resp.PubID).To(Equal("A1B-A123456"))
				Expect(resp.Page).To(Equal("https://donatello.to/nickname"))
				Expect(resp.IsActive).To(BeTrue())
				Expect(resp.IsPublic).To(BeTrue())
				Expect(resp.Donates).To(Equal(v1.ResponseMeDonates{
					TotalAmount: 100,
					TotalCount:  1,
				}))
				Expect(resp.CreatedAt).To(Equal("2022-10-20 00:30:50"))
			})
		})

		Context("Donates", func() {
			It("respond as unauthorized", func() {
				resp, err := client.Donates(ctx, 1, 20)
				Expect(err).To(Succeed())
				Expect(resp.Success).To(BeFalse())
				Expect(resp.Message).To(Equal("Помилка авторизації"))
			})

			It("respond with correct data", func() {
				api.MockDonates = func(page, size int64) (int64, []byte, error) {
					return http.StatusOK, []byte(MockDonatesData), nil
				}
				resp, err := client.Donates(ctx, 1, 20)
				Expect(err).To(Succeed())
				Expect(resp.Success).To(BeTrue())
				Expect(resp.Message).To(Equal(""))
				Expect(resp.Content).To(Equal([]v1.ResponseDonatesContent{
					{
						PubID:       "A1B-A123456",
						ClientName:  "ClientName",
						Message:     "Message",
						Amount:      "100",
						Currency:    "UAH",
						Goal:        "",
						IsPublished: false,
						CreatedAt:   "2022-10-20 00:30:50",
					},
				}))
				Expect(resp.Content[0].AmountInt64()).To(Equal(int64(100)))
				Expect(resp.Page).To(Equal(int64(1)))
				Expect(resp.Size).To(Equal(int64(20)))
				Expect(resp.Pages).To(Equal(int64(1)))
				Expect(resp.First).To(BeTrue())
				Expect(resp.Last).To(BeTrue())
				Expect(resp.Total).To(Equal(int64(1)))
			})
		})

		Context("Clients", func() {
			It("respond as unauthorized", func() {
				resp, err := client.Clients(ctx)
				Expect(err).To(Succeed())
				Expect(resp.Success).To(BeFalse())
				Expect(resp.Message).To(Equal("Помилка авторизації"))
			})

			It("respond with correct data", func() {
				api.MockClients = func() (int64, []byte, error) {
					return http.StatusOK, []byte(MockClientsData), nil
				}
				resp, err := client.Clients(ctx)
				Expect(err).To(Succeed())
				Expect(resp.Success).To(BeTrue())
				Expect(resp.Message).To(Equal(""))
				Expect(resp.Clients).To(Equal([]v1.ResponseClientsClients{
					{
						ClientName:  "ClientName",
						TotalAmount: 100,
					},
				}))
			})
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client")
}
