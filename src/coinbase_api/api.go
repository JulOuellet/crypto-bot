package coinbase_api

import (
    "net/http"
    "os"
    "io"
)

func GetAccounts() (int, string, error) {
    apiKey := os.Getenv("COINBASE_API_KEY")
    apiSecret := os.Getenv("COINBASE_API_SECRET")

    headers, err := GenerateHeaders(apiKey, apiSecret, "GET", "/api/v3/brokerage/accounts", "")
    if err != nil {
	return 0, "", err
    }

    client := &http.Client{}

    req, err := http.NewRequest("GET", "https://api.coinbase.com/api/v3/brokerage/accounts", nil)
    if err != nil {
	return 0, "", err
    }

    req.Header = headers

    resp, err := client.Do(req)
    if err != nil {
	return 0, "", err
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
	return 0, "", err
    }

    respCode := resp.StatusCode

    return respCode, string(respBody), nil
} 

