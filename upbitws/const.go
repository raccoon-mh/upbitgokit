package upbitws

// restapi v1.5.4 stable https://docs.upbit.com/v1.5.4/reference

const (
	serverPublicHost  = "wss://api.upbit.com/websocket/v1"
	serverPrivateHost = "wss://api.upbit.com/websocket/v1/private"

	// ######## Exchange API ########

	//
	// ******** 자산 ********
	//

	// 전체 계좌 조회
	// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EA%B3%84%EC%A2%8C-%EC%A1%B0%ED%9A%8C
	accountsEndPoint = "/v1/accounts:get"

	//
	// ******** 주문 ********
	//

	// 주문 가능 정보
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EA%B0%80%EB%8A%A5-%EC%A0%95%EB%B3%B4
	ordersChanceEndPoint = "/v1/orders/chance:get"

	// 개별 주문 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	orderEndPoint = "/v1/order:get"

	// 주문 리스트 조회
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
	// Deprecated (2024.06~)
	// 	GET v1/orders 를 사용중이실 경우, 용도에 따라 v1/orders/uuids ,v1/orders/open ,v1/orders/closed로 변경해주시기 바랍니다.
	// 	정확한 지원종료 예정일은 추후 공지사항을 통해 다시 안내드리겠습니다.

	// id로 주문리스트 조회
	// https://docs.upbit.com/reference/id%EB%A1%9C-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	orderUuidsEndPoint = "/v1/orders/uuids:get"

	// 체결 대기 주문 (Open Order) 조회
	// https://docs.upbit.com/reference/%EB%8C%80%EA%B8%B0-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	orderOpenEndPoint = "/v1/orders/open:get"

	// 종료된 주문 (Closed Order) 조회
	// https://docs.upbit.com/reference/%EC%A2%85%EB%A3%8C-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	ordersClosedEndPoint = "/v1/orders/closed:get"

	// 주문 취소 접수
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EC%B7%A8%EC%86%8C
	orderCancelEndPoint = "/v1/order:delete"

	// 주문하기
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8%ED%95%98%EA%B8%B0
	ordersEndPoint = "/v1/orders:post"

	//
	// ******** 출금 ********
	//

	// 출금 리스트 조회
	// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EC%B6%9C%EA%B8%88-%EC%A1%B0%ED%9A%8C
	withdrawsEndPoint = "/v1/withdraws:get"

	// 개별 출금 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%B6%9C%EA%B8%88-%EC%A1%B0%ED%9A%8C
	withdrawEndPoint = "/v1/withdraw:get"

	// 출금 가능 정보
	// https://docs.upbit.com/reference/%EC%B6%9C%EA%B8%88-%EA%B0%80%EB%8A%A5-%EC%A0%95%EB%B3%B4
	withdrawChanceEndPoint = "/v1/withdraws/chance:get"

	// 디지털 자산 출금하기
	// https://docs.upbit.com/reference/%EB%94%94%EC%A7%80%ED%84%B8%EC%9E%90%EC%82%B0-%EC%B6%9C%EA%B8%88%ED%95%98%EA%B8%B0
	withdrawCoinEndPoint = "/v1/withdraws/coinl:post"

	// 원화 출금하기
	// https://docs.upbit.com/reference/%EC%9B%90%ED%99%94-%EC%B6%9C%EA%B8%88%ED%95%98%EA%B8%B0
	withdrawKrwEndPoint = "/v1/withdraws/krw:post"

	// 출금 허용 주소 리스트 조회
	// https://docs.upbit.com/reference/%EC%B6%9C%EA%B8%88-%ED%97%88%EC%9A%A9-%EC%A3%BC%EC%86%8C-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
	withdrawCoinAddressesEndPoint = "/v1/withdraws/coin_addresses:get"

	//
	// ******** 입금 ********
	//

	// 입금 리스트 조회
	// https://docs.upbit.com/reference/%EC%9E%85%EA%B8%88-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
	depositsEndPoint = "/v1/deposits:get"

	// 개별 입금 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%9E%85%EA%B8%88-%EC%A1%B0%ED%9A%8C
	depositEndPoint = "/v1/deposit:get"

	// 입금 주소 생성 요청
	// https://docs.upbit.com/reference/%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%83%9D%EC%84%B1-%EC%9A%94%EC%B2%AD
	depositsGenerateCoinAddressEndPoint = "/v1/deposits/generate_coin_address:post"

	// 전체 입금 주소 조회
	// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%A1%B0%ED%9A%8C
	depositsCoinAddressesEndPoint = "/v1/deposits/coin_addresses:get"

	// 개별 입금 주소 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%A1%B0%ED%9A%8C
	depositsCoinAddressEndPoint = "/v1/deposits/coin_address:get"

	// 원화 입금하기
	// https://docs.upbit.com/reference/%EC%9B%90%ED%99%94-%EC%9E%85%EA%B8%88%ED%95%98%EA%B8%B0
	depositsKrwEndPoint = "/v1/deposits/krw:post"

	// ******** 입금 - 해외거래소 해외거래소 입금 시 계정주 확인 하기(트래블룰 검증) ********
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-%EA%B2%80%EC%A6%9D

	// 계정주 확인(트래블룰 검증)가능 거래소 리스트 조회
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-%EA%B0%80%EB%8A%A5-%EA%B1%B0%EB%9E%98%EC%86%8C
	travelRuleVaspsEndPoint = "/v1/travel_rule/vasps:get"

	// 입금 UUID로 트래블룰 검증하기
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-uuid
	travelRuleDepositUuidEndPoint = "/v1/travel_rule/deposit/uuid:post"

	// 입금 TxID로 트래블룰 검증하기
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-uuid
	travelRuleDepositTxidEndPoint = "/v1/travel_rule/deposit/txid:post"

	//
	// ******** 서비스 정보 ********
	//

	// 입출금 현황
	// https://docs.upbit.com/reference/%EC%9E%85%EC%B6%9C%EA%B8%88-%ED%98%84%ED%99%A9
	statusWalletEndPoint = "/v1/status/wallet:get"

	// API 키 리스트 조회
	// https://docs.upbit.com/reference/open-api-%ED%82%A4-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
	apiKeysEndPoint = "/v1/api_keys:get"

	// ######## Quotation API ########

	//
	// ******** 시세 종목 조회 ********
	//

	// 종목 코드 조회
	// https://docs.upbit.com/reference/%EB%A7%88%EC%BC%93-%EC%BD%94%EB%93%9C-%EC%A1%B0%ED%9A%8C
	marketAllEndPoint = "/v1/market/all:get"

	//
	// ******** 시세 캔들 조회 ********
	//

	// 초(Second) 캔들
	// https://docs.upbit.com/reference/%EC%B4%88second-%EC%BA%94%EB%93%A4
	candlesSecondsEndPoint = "/v1/candles/seconds:get"

	// 분(Minute) 캔들
	// https://docs.upbit.com/reference/%EB%B6%84minute-%EC%BA%94%EB%93%A4-1
	candlesMinutesUnitEndPoint = "/v1/candles/minutes/{unit}:get"

	// 일(Day) 캔들
	// https://docs.upbit.com/reference/%EB%B6%84minute-%EC%BA%94%EB%93%A4-1
	candlesDaysEndPoint = "/v1/candles/days:get"

	// 주(Week) 캔들
	// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
	candlesWeeksEndPoint = "/v1/candles/weeks:get"

	// 월(Month) 캔들
	// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
	candlesMonthsEndPoint = "/v1/candles/months:get"

	//
	// ******** 시세 체결 조회 ********
	//

	// 최근 체결 내역
	// https://docs.upbit.com/reference/%EC%B5%9C%EA%B7%BC-%EC%B2%B4%EA%B2%B0-%EB%82%B4%EC%97%AD
	tradesTicksEndPoint = "/v1/trades/ticks:get"

	//
	// ******** 시세 현재가(Ticker) 조회 ********
	//

	// 종목 단위 현재가 정보
	// https://docs.upbit.com/reference/ticker%ED%98%84%EC%9E%AC%EA%B0%80-%EC%A0%95%EB%B3%B4
	tickerEndPoint = "/v1/ticker:get"

	// 마켓 단위 현재가 정보
	// https://docs.upbit.com/reference/tickers_by_quote
	tickerAllEndPoint = "/v1/ticker/all:get"

	//
	// ******** 시세 호가 정보(Orderbook) 조회 ********
	//

	// 호가 정보 조회
	// https://docs.upbit.com/reference/%ED%98%B8%EA%B0%80-%EC%A0%95%EB%B3%B4-%EC%A1%B0%ED%9A%8C
	orderbookEndPoint = "/v1/orderbook:get"

	// 호가 모아보기 단위 정보 조회
	// https://docs.upbit.com/reference/supported_levels
	orderbookSupportedLevelsEndPoint = "/v1/orderbook/supported_levels:get"
)
