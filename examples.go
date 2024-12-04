package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/raccoon-mh/upbitgokit/upbitapi"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	for {
		// 메뉴 출력
		fmt.Println("\nselect number.")
		fmt.Println("1. wsmain")
		fmt.Println("2. apimain")
		fmt.Println("0. exit")
		fmt.Print("select: ")

		var choice int
		fmt.Scan(&choice)

		// 번호에 따른 함수 실행
		switch choice {
		// case 1:
		// 	wsmain()
		case 2:
			apimain()
		case 0:
			fmt.Println("프로그램 종료")
			return
		default:
			fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
		}
	}
}

// func wsmain() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	cred := upbitapi.Credential{
// 		AccessKey: os.Getenv("UPBIT_OPEN_API_ACCESS_KEY"),
// 		SecretKey: os.Getenv("UPBIT_OPEN_API_SECRET_KEY"),
// 	}

// 	ctx, err := upbitapi.SetCtxCredential(context.Background(), cred)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	token, err := upbitws.GenerateJwtToken(ctx)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	url := "wss://api.upbit.com/websocket/v1/private"
// 	headers := map[string][]string{
// 		"Authorization": {fmt.Sprintf("Bearer %s", token)},
// 	}

// 	conn, _, err := websocket.DefaultDialer.Dial(url, headers)
// 	if err != nil {
// 		log.Fatalf("WebSocket 연결 실패: %v", err)
// 	}
// 	defer conn.Close()

// 	fmt.Println("connected!")

// 	// 메시지 전송
// 	request := []map[string]string{
// 		{"ticket": "test example"},
// 		{"type": "myOrder"},
// 	}
// 	requestBytes, _ := json.Marshal(request)

// 	err = conn.WriteMessage(websocket.TextMessage, requestBytes)
// 	if err != nil {
// 		log.Fatalf("메시지 전송 실패: %v", err)
// 	}

// 	// 메시지 수신
// 	go func() {
// 		for {
// 			_, message, err := conn.ReadMessage()
// 			if err != nil {
// 				log.Println("메시지 수신 중 오류:", err)
// 				return
// 			}
// 			fmt.Println(string(message))
// 		}
// 	}()

// 	// 연결 유지
// 	select {}
// }

func apimain() {
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
