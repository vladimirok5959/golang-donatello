# golang-donatello

Go API client for ukrainian donate platform donatello.to

[https://donatello.to](https://donatello.to)

## Demo and testing

Example is here: [https://github.com/vladimirok5959/golang-donatello/blob/main/main.go](https://github.com/vladimirok5959/golang-donatello/blob/main/main.go)

```sh
export TOKEN="YOUR-TOKEN"
go run main.go

client.Me:
respMe: &v1.ResponseMe{
    Success: true,
    Message: "",
    NickName: "NickName",
    PubID: "A1B-A123456",
    Page: "https://donatello.to/nickname",
    IsActive: true,
    IsPublic: true,
    Donates: v1.ResponseMeDonates{
        TotalAmount: 10,
        TotalCount: 1
    },
    CreatedAt: "2022-10-20 00:30:50"
}
err: <nil>

client.Donates:
respDonates: &v1.ResponseDonates{
    Success: true,
    Message: "",
    Content: []v1.ResponseDonatesContent{
        v1.ResponseDonatesContent{
            PubID: "A1B-A123456",
            ClientName: "ClientName",
            Message: "Message",
            Amount: "100",
            Currency: "UAH",
            Goal: "",
            IsPublished: false,
            CreatedAt: "2022-10-20 00:30:50"
        }
    },
    Page: 1,
    Size: 20,
    Pages: 1,
    First: true,
    Last: true,
    Total: 1
}
err: <nil>

client.Clients:
respClients: &v1.ResponseClients{
    Success: true,
    Message: "",
    Clients: []v1.ResponseClientsClients{
        v1.ResponseClientsClients{
            ClientName: "ClientName",
            TotalAmount: 100
        }
    }
}
err: <nil>

client.EachDonate:
EachDonate: &v1.ResponseDonatesContent{
    PubID: "A1B-A123456",
    ClientName: "ClientName",
    Message: "Message",
    Amount: "100",
    Currency: "UAH",
    Goal: "",
    IsPublished: false,
    CreatedAt: "2022-10-20 00:30:50"
}
```

API faker included, see: [https://github.com/vladimirok5959/golang-donatello/blob/main/donatello/v1/client_fake_api.go](https://github.com/vladimirok5959/golang-donatello/blob/main/donatello/v1/client_fake_api.go)
