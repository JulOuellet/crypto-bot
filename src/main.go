package main

import (
    "log"

    "github.com/joho/godotenv"
    "github.com/JulOuellet/crypto-bot/src/coinbase_api"
)


func main() {
    
    // Load .env file
    err := godotenv.Load(".env")
    if err != nil {
	log.Fatal("Error loading .env file: ", err)
    }

    code, response, err := coinbase_api.GetAccounts()

    println(code)
    println(response)

}
