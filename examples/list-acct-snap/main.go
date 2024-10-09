package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	cato "github.com/cato-networks/cato-go-sdk"
)

func main() {
	token := os.Getenv("CATO_API_KEY")
	accountId := os.Getenv("CATO_ACCOUNT_ID")
	url := "https://api.catonetworks.com/api/v1/graphql2"

	if token == "" {
		fmt.Println("no token provided")
		os.Exit(1)
	}

	if accountId == "" {
		fmt.Println("no account id provided")
		os.Exit(1)
	}

	catoClient, _ := cato.New(url, token, *http.DefaultClient)

	ctx := context.Background()

	// AccountSnapshot(ctx context.Context, siteIDs []string, userIDs []string, accountID *string, interceptors ...clientv2.RequestInterceptor) (*AccountSnapshot, error)
	queryResult, err := catoClient.AccountSnapshot(ctx, nil, nil, &accountId)
	if err != nil {
		fmt.Println("policy query error: ", err)
		return
	}

	queryResultJson, err := json.Marshal(queryResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(queryResultJson))
}
