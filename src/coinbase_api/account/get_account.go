package account

import (
    "encoding/json"
    "io"
    "net/http"

    "github.com/JulOuellet/crypto-bot/src/coinbase_api"
)

type AccountResponse struct {
   Account Account `json:"account"`
}

func GetAccount(apiKey, apiSecret, accountUuid string, client http.Client) (int, AccountResponse, error) {
    var response AccountResponse

    headers, err := coinbase_api.GenerateHeaders(apiKey, apiSecret, "GET", "/api/v3/brokerage/accounts/" + accountUuid, "")
    if err != nil {
	return 0, response, err
    }

    req, err := http.NewRequest("GET", "https://api.coinbase.com/api/v3/brokerage/accounts/" + accountUuid, nil)
    if err != nil {
	return 0, response, err
    }
    req.Header = headers

    resp, err := client.Do(req)
    if err != nil {
	return 0, response, err
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
	return 0, response, err
    }

    err = json.Unmarshal(respBody, &response)
    if err != nil {
	return 0, response, err
    }

    respCode := resp.StatusCode

    return respCode, response, nil
}
