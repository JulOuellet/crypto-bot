package coinbase_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

func GenerateHeaders(apiKey, apiSecret, method, requestPath, body string) (map[string]string, error) {
    timestamp := strconv.FormatInt(time.Now().Unix(), 10)
    prehash := timestamp + method + requestPath + body

    decodedSecret, err := base64.StdEncoding.DecodeString(apiSecret)
    if err != nil {
	return nil, err
    }

    mac := hmac.New(sha256.New, decodedSecret)
    mac.Write([]byte(prehash))

    signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

    headers := map[string]string {
	"CB-ACCESS-KEY": apiKey,
	"CB-ACCESS-SIGN": signature,
	"CB-ACCESS-TIMESTAMP": timestamp,
    }

    return headers, nil
}

