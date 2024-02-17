package coinbase_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

func BuildHeaders(apiKey, apiSecret, method, requestPath, body string) (http.Header, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	signatureString := timestamp + method + requestPath + body
	hmacObject := hmac.New(sha256.New, []byte(apiSecret))
	_, err := hmacObject.Write([]byte(signatureString))
	if err != nil {
		return nil, err
	}
	signature := hex.EncodeToString(hmacObject.Sum(nil))

	headers := make(http.Header)
	headers.Set("CB-ACCESS-KEY", apiKey)
	headers.Set("CB-ACCESS-SIGN", signature)
	headers.Set("CB-ACCESS-TIMESTAMP", timestamp)

	return headers, nil
}
