package restapi

// restapi v1.5.1 https://docs.upbit.com/v1.5.1/reference

const (
	// ######## Exchange API ########

	//
	// ******** 자산 ********
	//

	// 전체 계좌 조회
	// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EA%B3%84%EC%A2%8C-%EC%A1%B0%ED%9A%8C
	accountsEndPoint = "/v1/accounts"

	//
	// ******** 주문 ********
	//

	// 주문 가능 정보
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EA%B0%80%EB%8A%A5-%EC%A0%95%EB%B3%B4
	ordersChanceEndPoint = "/v1/orders/chance"

	// 개별 주문 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	orderEndPoint = "/v1/order"

	// id로 주문리스트 조회
	// https://docs.upbit.com/reference/id%EB%A1%9C-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	orderUuidsEndPoint = "/v1/orders/uuids"

	// 체결 대기 주문 (Open Order) 조회
	// https://docs.upbit.com/reference/%EB%8C%80%EA%B8%B0-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	orderOpenEndPoint = "/v1/orders/open"

	// 종료된 주문 (Closed Order) 조회
	// https://docs.upbit.com/reference/%EC%A2%85%EB%A3%8C-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
	ordersClosedEndPoint = "/v1/orders/closed"

	// 주문 취소 접수
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EC%B7%A8%EC%86%8C
	orderCancelEndPoint = "/v1/order"

	// 주문하기
	// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8%ED%95%98%EA%B8%B0
	ordersEndPoint = "/v1/orders"

	//
	// ******** 출금 ********
	//

	// 출금 리스트 조회
	// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EC%B6%9C%EA%B8%88-%EC%A1%B0%ED%9A%8C
	withdrawsEndPoint = "/v1/withdraws"

	// 개별 출금 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%B6%9C%EA%B8%88-%EC%A1%B0%ED%9A%8C
	withdrawEndPoint = "/v1/withdraw"

	// 출금 가능 정보
	// https://docs.upbit.com/reference/%EC%B6%9C%EA%B8%88-%EA%B0%80%EB%8A%A5-%EC%A0%95%EB%B3%B4
	withdrawChanceEndPoint = "/v1/withdraws/chance"

	// 디지털 자산 출금하기
	// https://docs.upbit.com/reference/%EB%94%94%EC%A7%80%ED%84%B8%EC%9E%90%EC%82%B0-%EC%B6%9C%EA%B8%88%ED%95%98%EA%B8%B0
	withdrawCoinEndPoint = "/v1/withdraws/coin"

	// 원화 출금하기
	// https://docs.upbit.com/reference/%EC%9B%90%ED%99%94-%EC%B6%9C%EA%B8%88%ED%95%98%EA%B8%B0
	withdrawKrwEndPoint = "/v1/withdraws/krw"

	// 출금 허용 주소 리스트 조회
	// https://docs.upbit.com/reference/%EC%B6%9C%EA%B8%88-%ED%97%88%EC%9A%A9-%EC%A3%BC%EC%86%8C-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
	withdrawCoinAddressesEndPoint = "/v1/withdraws/coin_addresses"

	//
	// ******** 입금 ********
	//

	// 입금 리스트 조회
	// https://docs.upbit.com/reference/%EC%9E%85%EA%B8%88-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
	depositsEndPoint = "/v1/deposits"

	// 개별 입금 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%9E%85%EA%B8%88-%EC%A1%B0%ED%9A%8C
	depositEndPoint = "/v1/deposit"

	// 입금 주소 생성 요청
	// https://docs.upbit.com/reference/%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%83%9D%EC%84%B1-%EC%9A%94%EC%B2%AD
	depositsGenerateCoinAddressEndPoint = "/v1/deposits/generate_coin_address"

	// 전체 입금 주소 조회
	// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%A1%B0%ED%9A%8C
	depositsCoinAddressesEndPoint = "/v1/deposits/coin_addresses"

	// 개별 입금 주소 조회
	// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%A1%B0%ED%9A%8C
	depositsCoinAddressEndPoint = "/v1/deposits/coin_address"

	// 원화 입금하기
	// https://docs.upbit.com/reference/%EC%9B%90%ED%99%94-%EC%9E%85%EA%B8%88%ED%95%98%EA%B8%B0
	depositsKrwEndPoint = "/v1/deposits/krw"

	// ******** 입금 - 해외거래소 해외거래소 입금 시 계정주 확인 하기(트래블룰 검증) ********
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-%EA%B2%80%EC%A6%9D

	// 계정주 확인(트래블룰 검증)가능 거래소 리스트 조회
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-%EA%B0%80%EB%8A%A5-%EA%B1%B0%EB%9E%98%EC%86%8C
	travelRuleVaspsEndPoint = "/v1/travel_rule/vasps"

	// 입금 UUID로 트래블룰 검증하기
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-uuid
	travelRuleDepositUuidEndPoint = "/v1/travel_rule/deposit/uuid"

	// 입금 TxID로 트래블룰 검증하기
	// https://docs.upbit.com/reference/%ED%8A%B8%EB%9E%98%EB%B8%94%EB%A3%B0-uuid
	travelRuleDepositTxidEndPoint = "/v1/travel_rule/deposit/txid"

	//
	// ******** 서비스 정보 ********
	//

	// 입출금 현황
	// https://docs.upbit.com/reference/%EC%9E%85%EC%B6%9C%EA%B8%88-%ED%98%84%ED%99%A9
	statusWalletEndPoint = "/v1/status/wallet"

	// API 키 리스트 조회
	// https://docs.upbit.com/reference/%EC%9E%85%EC%B6%9C%EA%B8%88-%ED%98%84%ED%99%A9
	apiKeysEndPoint = "/v1/api_keys"

	// ######## Quotation API ########

	//
	// ******** 시세 종목 조회 ********
	//

	// 종목 코드 조회
	// https://docs.upbit.com/reference/%EB%A7%88%EC%BC%93-%EC%BD%94%EB%93%9C-%EC%A1%B0%ED%9A%8C
	marketAllEndPoint = "/v1/market/all"

	//
	// ******** 시세 캔들 조회 ********
	//

	// 분(Minute) 캔들
	// https://docs.upbit.com/reference/%EB%B6%84minute-%EC%BA%94%EB%93%A4-1
	candlesMinutesUnitEndPoint = "/v1/candles/minutes/{unit}"

	// 일(Day) 캔들
	// https://docs.upbit.com/reference/%EB%B6%84minute-%EC%BA%94%EB%93%A4-1
	candlesDaysEndPoint = "/v1/candles/days"

	// 주(Week) 캔들
	// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
	candlesWeeksEndPoint = "/v1/candles/weeks"

	// 월(Month) 캔들
	// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
	candlesMonthsEndPoint = "/v1/candles/months"

	//
	// ******** 시세 체결 조회 ********
	//

	// 최근 체결 내역
	// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
	tradesTicksEndPoint = "/v1/trades/ticks"

	//
	// ******** 시세 현재가(Ticker) 조회 ********
	//

	// 종목 단위 현재가 정보
	// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
	tickerEndPoint = "/v1/ticker"

	// 마켓 단위 현재가 정보
	// https://docs.upbit.com/reference/tickers_by_quote
	tickerAllEndPoint = "/v1/ticker/all"

	//
	// ******** 시세 호가 정보(Orderbook) 조회 ********
	//

	// 호가 정보 조회
	// https://docs.upbit.com/reference/%ED%98%B8%EA%B0%80-%EC%A0%95%EB%B3%B4-%EC%A1%B0%ED%9A%8C
	orderbookEndPoint = "/v1/orderbook"

	// 호가 모아보기 단위 정보 조회
	// https://docs.upbit.com/reference/supported_levels
	orderbookSupportedLevelsEndPoint = "/v1/orderbook/supported_levels"
)
