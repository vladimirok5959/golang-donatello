package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	v1 "github.com/vladimirok5959/golang-donatello/donatello/v1"
)

var token = ""

func init() {
	flag.StringVar(&token, "token", "", "Token")
	flag.Parse()

	if token == "" {
		if os.Getenv("TOKEN") != "" {
			token = os.Getenv("TOKEN")
		}
	}
}

func main() {
	api := v1.NewClientAPI(time.Second*10, token)
	client := v1.NewClient(api)

	ctx := context.Background()

	fmt.Printf("client.Me:\n")
	respMe, err := client.Me(ctx)
	fmt.Printf("respMe: %#v\n", respMe)
	fmt.Printf("err: %#v\n\n", err)

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("client.Donates:\n")
	respDonates, err := client.Donates(ctx, 0, 20)
	fmt.Printf("respDonates: %#v\n", respDonates)
	fmt.Printf("err: %#v\n\n", err)

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("client.Clients:\n")
	respClients, err := client.Clients(ctx)
	fmt.Printf("respClients: %#v\n", respClients)
	fmt.Printf("err: %#v\n\n", err)

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("client.EachDonate:\n")
	_ = client.EachDonate(ctx, func(donate *v1.ResponseDonatesContent) error {
		fmt.Printf("EachDonate: %#v\n", donate)
		return nil
	})
}
