package upbitapi

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type MarketInfoResponse struct {
	BidFee      string  `json:"bid_fee"`       // 매수 수수료 비율 (NumberString)
	AskFee      string  `json:"ask_fee"`       // 매도 수수료 비율 (NumberString)
	MakerBidFee string  `json:"maker_bid_fee"` // 메이커 매수 수수료 비율 (NumberString)
	MakerAskFee string  `json:"maker_ask_fee"` // 메이커 매도 수수료 비율 (NumberString)
	Market      Market  `json:"market"`        // 마켓에 대한 정보 (Object)
	BidAccount  Account `json:"bid_account"`   // 매수 계좌 상태 (Object)
	AskAccount  Account `json:"ask_account"`   // 매도 계좌 상태 (Object)
}

type Market struct {
	ID         string       `json:"id"`          // 마켓의 유일 키 (String)
	Name       string       `json:"name"`        // 마켓 이름 (String)
	OrderTypes []string     `json:"order_types"` // 지원 주문 방식 (만료, Array[String])
	OrderSides []string     `json:"order_sides"` // 지원 주문 종류 (Array[String])
	BidTypes   []string     `json:"bid_types"`   // 매수 주문 지원 방식 (Array[String])
	AskTypes   []string     `json:"ask_types"`   // 매도 주문 지원 방식 (Array[String])
	Bid        CurrencyInfo `json:"bid"`         // 매수 제약사항 (Object)
	Ask        CurrencyInfo `json:"ask"`         // 매도 제약사항 (Object)
	MaxTotal   string       `json:"max_total"`   // 최대 매도/매수 금액 (NumberString)
	State      string       `json:"state"`       // 마켓 운영 상태 (String)
}

type CurrencyInfo struct {
	Currency  string `json:"currency"`   // 화폐를 의미하는 영문 대문자 코드 (String)
	PriceUnit string `json:"price_unit"` // 주문금액 단위 (String)
	MinTotal  string `json:"min_total"`  // 최소 매도/매수 금액 (NumberString)
}

type Account struct {
	Currency            string `json:"currency"`               // 화폐를 의미하는 영문 대문자 코드 (String)
	Balance             string `json:"balance"`                // 주문가능 금액/수량 (NumberString)
	Locked              string `json:"locked"`                 // 주문 중 묶여있는 금액/수량 (NumberString)
	AvgBuyPrice         string `json:"avg_buy_price"`          // 매수평균가 (NumberString)
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"` // 매수평균가 수정 여부 (Boolean)
	UnitCurrency        string `json:"unit_currency"`          // 평단가 기준 화폐 (String)
}

// 주문 가능 정보
//   - market: 상세 정보를 가져올 마켓의 고유 ID입니다. (타입: String)
//
// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EA%B0%80%EB%8A%A5-%EC%A0%95%EB%B3%B4
func OrdersChanceGet(ctx context.Context, market string) (*MarketInfoResponse, error) {
	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"market": market,
		},
	}
	return commonAnyCaller(ctx, ordersChanceEndPoint, reqform, &MarketInfoResponse{})
}

type OrderGetResponse struct {
	UUID           string  `json:"uuid"`            // 주문 고유 아이디
	Side           string  `json:"side"`            // 주문 방향 (bid/ask)
	OrdType        string  `json:"ord_type"`        // 주문 타입 (limit/price/market 등)
	Price          string  `json:"price"`           // 주문 가격
	State          string  `json:"state"`           // 주문 상태
	Market         string  `json:"market"`          // 마켓 정보
	CreatedAt      string  `json:"created_at"`      // 주문 생성 시간
	ReservedFee    string  `json:"reserved_fee"`    // 예약된 수수료
	RemainingFee   string  `json:"remaining_fee"`   // 남은 수수료
	PaidFee        string  `json:"paid_fee"`        // 지불된 수수료
	Locked         string  `json:"locked"`          // 거래에 사용 중인 비용
	ExecutedVolume string  `json:"executed_volume"` // 체결된 양
	TradesCount    int     `json:"trades_count"`    // 거래 횟수
	Trades         []Trade `json:"trades"`          // 거래 상세 정보 배열
}

type Trade struct {
	Market    string `json:"market"`     // 마켓 정보
	UUID      string `json:"uuid"`       // 거래의 고유 아이디
	Price     string `json:"price"`      // 거래 가격
	Volume    string `json:"volume"`     // 거래량
	Funds     string `json:"funds"`      // 총 자금
	Trend     string `json:"trend"`      // 추세 정보 (예: "up")
	CreatedAt string `json:"created_at"` // 거래 생성 시간
	Side      string `json:"side"`       // 거래 방향 (bid/ask)
}

// 개별 주문 조회
// https://docs.upbit.com/reference/%EA%B0%9C%EB%B3%84-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
func OrderGet(ctx context.Context, uuid, identifier string) (*OrderGetResponse, error) {
	if uuid == "" && identifier == "" {
		return nil, fmt.Errorf("either uuid or identifier must be included")
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{},
	}
	if uuid != "" {
		reqform.QueryParams["uuid"] = uuid
	}
	if identifier != "" {
		reqform.QueryParams["identifier"] = identifier
	}
	return commonAnyCaller(ctx, orderEndPoint, reqform, &OrderGetResponse{})
}

type OrdersPostResponse struct {
	Uuid            string `json:"uuid"`             // 주문의 고유 아이디
	Side            string `json:"side"`             // 주문 종류
	OrdType         string `json:"ord_type"`         // 주문 방식
	Price           string `json:"price"`            // 주문 당시 화폐 가격 (NumberString)
	State           string `json:"state"`            // 주문 상태
	Market          string `json:"market"`           // 마켓의 유일키
	CreatedAt       string `json:"created_at"`       // 주문 생성 시간
	Volume          string `json:"volume"`           // 사용자가 입력한 주문 양 (NumberString)
	RemainingVolume string `json:"remaining_volume"` // 체결 후 남은 주문 양 (NumberString)
	ReservedFee     string `json:"reserved_fee"`     // 수수료로 예약된 비용 (NumberString)
	RemainingFee    string `json:"remaining_fee"`    // 남은 수수료 (NumberString)
	PaidFee         string `json:"paid_fee"`         // 사용된 수수료 (NumberString)
	Locked          string `json:"locked"`           // 거래에 사용중인 비용 (NumberString)
	ExecutedVolume  string `json:"executed_volume"`  // 체결된 양 (NumberString)
	TradesCount     int    `json:"trades_count"`     // 해당 주문에 걸린 체결 수
	TimeInForce     string `json:"time_in_force"`    // IOC, FOK 설정
}

// 주문하기
// Request Parameters
// Name          설명 타입
// market *      마켓 ID (필수) String
//   - 예: "KRW-BTC"
//
// side *        주문 종류 (필수) String
//   - bid: 매수
//   - ask: 매도
//
// volume *      주문량 (지정가, 시장가 매도 시 필수) NumberString
// price *       주문 가격. (지정가, 시장가 매수 시 필수) NumberString
//   - 예: KRW-BTC 마켓에서 1BTC당 1,000 KRW로 거래할 경우 값은 "1000"
//   - 예: KRW-BTC 마켓에서 1BTC당 매도 1호가가 500 KRW인 경우,
//     시장가 매수 시 값을 "1000"으로 세팅하면 2BTC 매수 가능 (수수료 영향 있음)
//
// ord_type *    주문 타입 (필수) String
//   - limit: 지정가 주문
//   - price: 시장가 주문 (매수)
//   - market: 시장가 주문 (매도)
//   - best: 최유리 주문 (time_in_force 설정 필수)
//
// identifier    조회용 사용자 지정 값 (선택) String (Uniq 값 사용)
//   - 주문을 조회하기 위한 고유 값
//   - 중복 값이 들어오면 오류 발생
//
// identifier 주의사항:
// - 서비스에서 발급하는 UUID가 아닌, 사용자가 직접 발급하는 키 값이어야 함
// - 중복된 값으로 요청 시 중복 오류 발생
// - 매 요청 시 새로운 값을 생성해야 함
//
// time_in_force IOC, FOK 주문 설정 * String
//   - ioc: Immediate or Cancel
//   - fok: Fill or Kill
//   - ord_type이 best 혹은 limit일 때만 지원
//
// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8%ED%95%98%EA%B8%B0
func OrdersPost(ctx context.Context, market string, side string, volume float64, price int64, orderType, identifier, timeinforce string) (*OrdersPostResponse, error) {
	if market == "" || side == "" || orderType == "" || price == 0 {
		return nil, fmt.Errorf("missing require input")
	}
	if identifier == "" {
		identifier = uuid.New().String()
	}

	var validBidTypes = map[string]bool{"bid": true, "ask": true}
	if !validBidTypes[side] {
		return nil, fmt.Errorf("invalid side type: %s", side)
	}

	var validOrderTypes = map[string]bool{"limit": true, "price": true, "market": true, "best": true}
	var validtimeinforceTypes = map[string]bool{"ioc": true, "fok": true}
	if !validOrderTypes[orderType] {
		return nil, fmt.Errorf("invalid order type: %s", orderType)
	} else if (orderType == "best" || orderType == "limit") && !validtimeinforceTypes[timeinforce] {
		return nil, fmt.Errorf("invalid timeinforce type or missing: %s", timeinforce)
	}

	reqform := RequestForm{
		RequestBody: make(map[string]interface{}),
	}
	reqform.RequestBody["market"] = market
	reqform.RequestBody["side"] = side
	if volume != 0.0 {
		reqform.RequestBody["volume"] = strconv.FormatFloat(volume, 'f', -1, 64)
	}
	reqform.RequestBody["price"] = strconv.FormatInt(price, 10)
	reqform.RequestBody["ord_type"] = orderType
	reqform.RequestBody["identifier"] = identifier
	if timeinforce != "" {
		reqform.RequestBody["time_in_force"] = timeinforce
	}

	return commonAnyCaller(ctx, ordersEndPoint, reqform, &OrdersPostResponse{})
}
