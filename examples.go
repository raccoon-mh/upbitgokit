package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"upbitgokit/upbitapi"

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

	cred := upbitapi.Credential{
		AccessKey: os.Getenv("UPBIT_OPEN_API_ACCESS_KEY"),
		SecretKey: os.Getenv("UPBIT_OPEN_API_SECRET_KEY"),
	}

	ctx, err := upbitapi.SetCtxCredential(context.Background(), cred)
	if err != nil {
		log.Fatal(err.Error())
	}

	// data, err := upbitapi.CandlesMonthGet(ctx, "KRW-BTC", "2024-10-01 00:00:00", 1)
	// data, err := upbitapi.OrdersChanceGet(ctx, "KRW-SHIB")
	// data, err := upbitapi.OrdersPost(ctx, "KRW-SHIB", "bid", 0.0, 5000, "price", "", "")
	data, err := upbitapi.AccountsGet(ctx)
	// data, err := upbitapi.OrdersPost(ctx, "KRW-SHIB", "bid", 0.0, 5000, "best", "", "ioc")
	// data, err := upbitapi.OrdersPost(ctx, "KRW-SHIB", "ask", 200000., 0.0, "best", "", "ioc")
	// data, err := upbitapi.MarketAllGet(ctx, true)
	// data, err := upbitapi.OrderGet(ctx, "", "")
	// data, err := upbitapi.OrdersClosedGet(ctx, "KRW-SHIB", "", "2024-11-09T18:00:00+09:00", "", 10, "")
	// data, err := upbitapi.OrderUuidsGet(ctx, "", []string{""}, []string{}, "")
	// data, err := upbitapi.OrderOpenGet(ctx, "KRW-SHIB", "wait", 0, 0, "asc")
	// data, err := upbitapi.OrderCancelDelete(ctx, "", "")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	prettyJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Printf("%+v\n", data)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
