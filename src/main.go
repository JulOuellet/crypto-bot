package main

import (
    "log"
    "os"
    "net/http"

    "github.com/joho/godotenv"
    "github.com/JulOuellet/crypto-bot/src/coinbase_api"
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

    code, response, err := coinbase_api.GetAccounts(apiKey, apiSecret, client)

    println(code)
    println(response.Accounts[0].Name)
}

