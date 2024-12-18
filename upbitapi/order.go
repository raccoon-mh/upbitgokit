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

type OrderUuidsGetResponse struct {
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	Volume          string `json:"volume"`
	RemainingVolume string `json:"remaining_volume"`
	ReservedFee     string `json:"reserved_fee"`
	RemainingFee    string `json:"remaining_fee"`
	PaidFee         string `json:"paid_fee"`
	Locked          string `json:"locked"`
	ExecutedVolume  string `json:"executed_volume"`
	ExecutedFunds   string `json:"executed_funds"`
	TradesCount     int    `json:"trades_count"`
	TimeInForce     string `json:"time_in_force"`
}

type OrderUuidsGetResponses []OrderUuidsGetResponse

// id로 주문리스트 조회
// market: 마켓 ID (String)
//
// uuids: 주문 UUID의 목록 (최대 100개) (Array[String])
//
//	uuids 또는 identifiers 중 한 가지 필드는 필수이며, 두 가지 필드를 함께 사용할 수 없습니다.
//
// identifiers: 주문 identifier의 목록 (최대 100개) (Array[String])
//
// order_by: 정렬 방식 (String)
//   - "asc": 오름차순
//   - "desc": 내림차순 (기본값)
//
// https://docs.upbit.com/reference/id%EB%A1%9C-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
func OrderUuidsGet(ctx context.Context, market string, uuid, identifier []string, orderBy string) (*OrderUuidsGetResponses, error) {
	if (len(uuid) == 0 && len(identifier) == 0) || (len(uuid) != 0 && len(identifier) != 0) {
		return nil, fmt.Errorf("either uuid or identifier must be included and can not use both")
	}
	var validOrderByTypes = map[string]bool{"asc": true, "desc": true}
	if orderBy != "" && !validOrderByTypes[orderBy] {
		return nil, fmt.Errorf("invalid orderBy type: %s", orderBy)
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{},
	}
	if market != "" {
		reqform.QueryParams["market"] = market
	}
	if len(uuid) != 0 {
		reqform.QueryParams["uuid"] = uuid
	}
	if len(identifier) != 0 {
		reqform.QueryParams["identifier"] = identifier
	}
	if orderBy != "" {
		reqform.QueryParams["order_by"] = orderBy
	}
	return commonAnyCaller(ctx, orderUuidsEndPoint, reqform, &OrderUuidsGetResponses{})
}

type OrderOpenGetResponse struct {
	UUID            string `json:"uuid"`             // 주문의 고유 아이디
	Side            string `json:"side"`             // 주문 종류
	OrdType         string `json:"ord_type"`         // 주문 방식 (limit, price, market, best)
	Price           string `json:"price"`            // 주문 당시 화폐 가격 (NumberString)
	State           string `json:"state"`            // 주문 상태
	Market          string `json:"market"`           // 마켓 ID
	CreatedAt       string `json:"created_at"`       // 주문 생성 시간 (DateString)
	Volume          string `json:"volume"`           // 사용자가 입력한 주문 양 (NumberString)
	RemainingVolume string `json:"remaining_volume"` // 체결 후 남은 주문 양 (NumberString)
	ReservedFee     string `json:"reserved_fee"`     // 수수료로 예약된 비용 (NumberString)
	RemainingFee    string `json:"remaining_fee"`    // 남은 수수료 (NumberString)
	PaidFee         string `json:"paid_fee"`         // 사용된 수수료 (NumberString)
	Locked          string `json:"locked"`           // 거래에 사용 중인 비용 (NumberString)
	ExecutedVolume  string `json:"executed_volume"`  // 체결된 양 (NumberString)
	ExecutedFunds   string `json:"executed_funds"`   // 현재까지 체결된 금액 (NumberString)
	TradesCount     int    `json:"trades_count"`     // 해당 주문에 걸린 체결 수 (Integer)
	TimeInForce     string `json:"time_in_force"`    // IOC, FOK 설정 (ioc, fok)
}

type OrderOpenGetResponses []OrderOpenGetResponse

// 체결 대기 주문 (Open Order) 조회
// market: 마켓 ID (String)
// state: 주문 상태 (String)
//   - "wait": 체결 대기 (기본값)
//   - "watch": 예약주문 대기
//
// states: 주문 상태의 목록 (Array[String])
//   - 기본값은 "wait"이며, "wait"와 "watch"를 함께 조회하려면 둘 다 지정해야 합니다.
//   - state와 states는 동시에 사용할 수 없습니다. 둘 중 하나만 지정해야 합니다.
//
// page: 페이지 수 (Number, 기본값: 1)
//   - 페이지 당 조회 건수는 limit 파라미터로 조절 가능
//
// limit: 요청 개수 (Number, 기본값: 100, 최대값: 100)
// order_by: 정렬 방식 (String)
//   - "asc": 오름차순
//   - "desc": 내림차순 (기본값)
//
// https://docs.upbit.com/reference/%EB%8C%80%EA%B8%B0-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
func OrderOpenGet(ctx context.Context, market, state string, page, limit int64, orderBy string) (*OrderOpenGetResponses, error) {
	var validStateTypes = map[string]bool{"wait": true, "watch": true}
	if state != "" && !validStateTypes[state] {
		return nil, fmt.Errorf("invalid state type: %s", state)
	}
	var validOrderByTypes = map[string]bool{"asc": true, "desc": true}
	if orderBy != "" && !validOrderByTypes[orderBy] {
		return nil, fmt.Errorf("invalid state type: %s", orderBy)
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{},
	}
	if market != "" {
		reqform.QueryParams["market"] = market
	}
	if page != 0 {
		reqform.QueryParams["page"] = page
	}
	if limit != 0 {
		reqform.QueryParams["limit"] = limit
	}
	if orderBy != "" {
		reqform.QueryParams["order_by"] = orderBy
	}
	return commonAnyCaller(ctx, orderOpenEndPoint, reqform, &OrderOpenGetResponses{})
}

type OrdersClosedGetResponse struct {
	UUID            string        `json:"uuid"`                       // 주문 고유 아이디
	Side            string        `json:"side"`                       // 주문 방향 (bid/ask)
	OrdType         string        `json:"ord_type"`                   // 주문 타입 (limit/price/market 등)
	Price           string        `json:"price"`                      // 주문 가격
	State           string        `json:"state"`                      // 주문 상태 (done/cancel 등)
	Market          string        `json:"market"`                     // 마켓 정보
	CreatedAt       string        `json:"created_at"`                 // 주문 생성 시간
	Volume          string        `json:"volume,omitempty"`           // 주문량 (옵션)
	RemainingVolume string        `json:"remaining_volume,omitempty"` // 남은 주문량 (옵션)
	ReservedFee     string        `json:"reserved_fee"`               // 예약된 수수료
	RemainingFee    string        `json:"remaining_fee"`              // 남은 수수료
	PaidFee         string        `json:"paid_fee"`                   // 지불된 수수료
	Locked          string        `json:"locked"`                     // 거래에 사용 중인 비용
	ExecutedVolume  string        `json:"executed_volume"`            // 체결된 양
	ExecutedFunds   string        `json:"executed_funds"`             // 체결된 총 금액
	TradesCount     int           `json:"trades_count"`               // 거래 횟수
	Trades          []TradeDetail `json:"trades,omitempty"`           // 거래 상세 정보 배열 (옵션)
}

type TradeDetail struct {
	Market    string `json:"market"`     // 마켓 정보
	UUID      string `json:"uuid"`       // 거래의 고유 아이디
	Price     string `json:"price"`      // 거래 가격
	Volume    string `json:"volume"`     // 거래량
	Funds     string `json:"funds"`      // 총 자금
	Trend     string `json:"trend"`      // 추세 (예: "up")
	CreatedAt string `json:"created_at"` // 거래 생성 시간
	Side      string `json:"side"`       // 거래 방향 (bid/ask)
}

type OrdersClosedGetResponses []OrdersClosedGetResponse

// 종료된 주문 (Closed Order) 조회
// market       마켓 ID (필수)                                    String
// state        주문 상태                                         String
//   - done : 전체 체결 완료
//   - cancel : 주문 취소
//   - 시장가 주문이 조회되지 않는 경우
//   - 시장가 매수 주문은 체결 후 주문 상태가 cancel 또는 done이 될 수 있음.
//   - 체결 후 남은 잔량이 있는 경우 잔량이 반환되며 cancel 처리됨.
//   - 주문 잔량이 없이 딱 맞아떨어지게 체결된 경우 done 상태로 표시됨.
//
// start_time   조회 시작 시각 (주문 생성 시각 기준)               String (ISO-8601)
//   - ISO-8601 포맷 필요 (예: 2024-03-13T00:00:00+09:00)
//   - start_time, end_time이 정의되지 않으면 현재 시각 기준 최대 7일 전까지 조회 가능.
//   - start_time만 정의 시, start_time부터 최대 7일 후까지 조회 가능.
//
// end_time     조회 종료 시각 (주문 생성 시각 기준)               String (ISO-8601)
//   - end_time만 정의 시, end_time부터 최대 7일 전까지 조회 가능.
//   - start_time과 end_time 모두 정의 시, 최대 7일 범위까지만 조회 가능.
//
// limit        요청 개수 (기본값: 100, 최대값: 1,000)             Number
//   - 시간 범위 내 주문 개수가 1,000개를 초과하면 시간 범위를 나누어 조회 필요.
//
// order_by     정렬 방식                                         String
//   - asc : 오름차순
//   - desc : 내림차순 (기본값)
//
// https://docs.upbit.com/reference/%EC%A2%85%EB%A3%8C-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C
func OrdersClosedGet(ctx context.Context, market, state, startTime, endTime string, limit int32, orderBy string) (*OrdersClosedGetResponses, error) {
	var validStateTypes = map[string]bool{"done": true, "cancel": true}
	if state != "" && !validStateTypes[state] {
		return nil, fmt.Errorf("invalid state type: %s", state)
	}
	var validOrderByTypes = map[string]bool{"asc": true, "desc": true}
	if orderBy != "" && !validOrderByTypes[orderBy] {
		return nil, fmt.Errorf("invalid state type: %s", orderBy)
	}
	if err := validateISO8601Format(startTime); err != nil {
		return nil, err
	}
	if err := validateISO8601Format(endTime); err != nil {
		return nil, err
	}
	if startTime != "" && endTime != "" {
		isInDay, err := isWithin7Days(startTime, endTime)
		if err != nil {
			return nil, err
		}
		if !isInDay {
			return nil, fmt.Errorf("startTime and endTime is over 7 days")
		}
	}
	reqform := RequestForm{
		QueryParams: make(map[string]interface{}),
	}
	if market != "" {
		reqform.QueryParams["market"] = market
	}
	if state != "" {
		reqform.QueryParams["state"] = state
	}
	if limit > 0 {
		reqform.QueryParams["limit"] = limit
	}
	if orderBy != "" {
		reqform.QueryParams["orderBy"] = orderBy
	}
	return commonAnyCaller(ctx, ordersClosedEndPoint, reqform, &OrdersClosedGetResponses{})
}

type OrderCancelDeleteRespnse struct {
	UUID            string `json:"uuid"`             // 주문의 고유 아이디
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
	Locked          string `json:"locked"`           // 거래에 사용 중인 비용 (NumberString)
	ExecutedVolume  string `json:"executed_volume"`  // 체결된 양 (NumberString)
	TradesCount     int    `json:"trades_count"`     // 해당 주문에 걸린 체결 수
}

// 주문 취소 접수
// https://docs.upbit.com/reference/%EC%A3%BC%EB%AC%B8-%EC%B7%A8%EC%86%8C
func OrderCancelDelete(ctx context.Context, uuid, identifier string) (*OrderCancelDeleteRespnse, error) {
	if uuid == "" && identifier == "" || uuid != "" && identifier != "" {
		return nil, fmt.Errorf("either uuid or identifier must be included and can not use both")

	}
	reqform := RequestForm{
		QueryParams: make(map[string]interface{}),
	}
	if uuid != "" {
		reqform.QueryParams["uuid"] = uuid
	}
	if identifier != "" {
		reqform.QueryParams["identifier"] = identifier
	}
	return commonAnyCaller(ctx, orderCancelEndPoint, reqform, &OrderCancelDeleteRespnse{})
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
	if market == "" || side == "" || orderType == "" {
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
	if price != 0 {
		reqform.RequestBody["price"] = strconv.FormatInt(price, 10)
	}
	reqform.RequestBody["ord_type"] = orderType
	reqform.RequestBody["identifier"] = identifier
	if timeinforce != "" {
		reqform.RequestBody["time_in_force"] = timeinforce
	}
	return commonAnyCaller(ctx, ordersEndPoint, reqform, &OrdersPostResponse{})
}
