package coinbase_api

import (
    "net/http"
    "io"
)

func GetAccounts(apiKey, apiSecret string, client http.Client) (int, string, error) {

    headers, err := GenerateHeaders(apiKey, apiSecret, "GET", "/api/v3/brokerage/accounts", "")
    if err != nil {
	return 0, "", err
    }

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

