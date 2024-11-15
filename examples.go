package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"upbitapi/config"
	"upbitapi/upbitapi"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cred := config.Credential{
		AccessKey: os.Getenv("UPBIT_OPEN_API_ACCESS_KEY"),
		SecretKey: os.Getenv("UPBIT_OPEN_API_SECRET_KEY"),
	}

	ctx, err := config.SetCtxCredential(context.Background(), cred)
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := upbitapi.CandlesMonthGet(ctx, "KRW-BTC", "2024-10-01 00:00:00", 1)
	if err != nil {
		log.Fatal(err.Error())
	}

	prettyJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Printf(" %+v\n", data)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
