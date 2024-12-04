package upbitws

// func priv(ctx context.Context) {
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
