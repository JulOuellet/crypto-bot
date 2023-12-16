package account

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JulOuellet/crypto-bot/src/coinbase_api"
)

type AccountsResponse struct {
   Accounts []Account `json:"accounts"`
   HasNext bool     `json:"has_next"`
   Cursor  string   `json:"cursor"`
   Size    int      `json:"size"`
}

func GetAccounts(apiKey, apiSecret string, client http.Client) (int, AccountsResponse, error) {
    var response AccountsResponse

    headers, err := coinbase_api.GenerateHeaders(apiKey, apiSecret, "GET", "/api/v3/brokerage/accounts", "")
    if err != nil {
	return 0, response, err
    }

    req, err := http.NewRequest("GET", "https://api.coinbase.com/api/v3/brokerage/accounts", nil)
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

