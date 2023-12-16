package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JulOuellet/crypto-bot/src/coinbase_api/account"
	"github.com/joho/godotenv"
)


func main() {
    
    // Load .env file
    err := godotenv.Load(".env")
    if err != nil {
	log.Fatal("Error loading .env file: ", err)
    }

    apiKey := os.Getenv("COINBASE_API_KEY")
    apiSecret := os.Getenv("COINBASE_API_SECRET")

    client := http.Client{}

    code, response, err := account.GetAccounts(apiKey, apiSecret, client)

    for _, account := range response.Accounts {
	if account.HasBalance() {
	    println("Account: ", account.Name + ", balance: ", account.AvailableBalance.Value, account.AvailableBalance.Currency)
	}
    }

    println("Code: ", code)
}

