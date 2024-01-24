package coinbase_api

import (
    "net/http"

    "github.com/JulOuellet/crypto-bot/src/coinbase_api/account"
)

type CoinbaseApi struct {
    apiKey string
    apiSecret string
    baseUrl string
    client *http.Client
}

func NewCoinbaseApi(apiKey, apiSecret string) *CoinbaseApi {
    return &CoinbaseApi{
	apiKey: apiKey,
	apiSecret: apiSecret,
	baseUrl: "https://api.coinbase.com/api/v3/brokerage/",
	client: &http.Client{},
    }
}

func (c *CoinbaseApi) GetAccounts() (int, account.AccountsResponse, error) {
    headers, err := GenerateHeaders(c.apiKey, c.apiSecret, "GET", "/api/v3/brokerage/accounts", "")
    if err != nil {
	return 0, account.AccountsResponse{}, err
    }
    return account.GetAccounts(c.apiKey, c.apiSecret, *c.client, headers)
}

