package coinbase_api

import (
	"io"
	"net/http"
	"encoding/json"
)

type Response struct {
   Accounts []Account `json:"accounts"`
   HasNext bool     `json:"has_next"`
   Cursor  string   `json:"cursor"`
   Size    int      `json:"size"`
}

func GetAccounts(apiKey, apiSecret string, client http.Client) (int, Response, error) {
    var response Response

    headers, err := GenerateHeaders(apiKey, apiSecret, "GET", "/api/v3/brokerage/accounts", "")
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

