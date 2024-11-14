package upbitapi

import (
	"context"
)

type TradesTicksGetResponse struct {
	Market           string  `json:"market"`             // 종목 코드
	TradeDateUTC     string  `json:"trade_date_utc"`     // 체결 일자 (UTC 기준, 포맷: yyyy-MM-dd)
	TradeTimeUTC     string  `json:"trade_time_utc"`     // 체결 시각 (UTC 기준, 포맷: HH:mm:ss)
	Timestamp        int64   `json:"timestamp"`          // 체결 타임스탬프 (Long)
	TradePrice       float64 `json:"trade_price"`        // 체결 가격 (Double)
	TradeVolume      float64 `json:"trade_volume"`       // 체결량 (Double)
	PrevClosingPrice float64 `json:"prev_closing_price"` // 전일 종가 (UTC 0시 기준, Double)
	ChangePrice      float64 `json:"change_price"`       // 변화량 (Double)
	AskBid           string  `json:"ask_bid"`            // 매도/매수
	SequentialID     int64   `json:"sequential_id"`      // 체결 번호 (Unique, Long)
}

type TradesTicksGetResponses []TradesTicksGetResponse

// 최근 체결 내역
// market (string, required):
//   - 종목 코드 (예: KRW-BTC)
//
// to (string, optional):
//   - 마지막 체결 시각
//   - 형식: [HHmmss 또는 HH:mm:ss]
//   - 비워서 요청 시 가장 최근 데이터 반환
//
// count (int32, optional):
//   - 체결 개수
//
// cursor (string, optional):
//   - 페이지네이션 커서 (sequential_id 사용)
//
// days_ago (int32, optional):
//   - 최근 체결 날짜 기준으로 7일 이내의 이전 데이터 조회 가능
//   - 비워서 요청 시 가장 최근 체결 날짜 반환
//   - 유효 범위: 1 ~ 7
//
// https://docs.upbit.com/reference/%EC%B5%9C%EA%B7%BC-%EC%B2%B4%EA%B2%B0-%EB%82%B4%EC%97%AD
func TradesTicksGet(ctx context.Context, market string, to string, count int32, cursor string, daysAgo int32) (*TradesTicksGetResponses, error) {
	if err := validateTimeString(to); err != nil {
		return nil, err
	}

	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"market": market,
		},
	}

	if to != "" {
		reqform.QueryParams["to"] = to
	}
	if count > 0 {
		reqform.QueryParams["count"] = count
	}
	if cursor != "" {
		reqform.QueryParams["cursor"] = cursor
	}
	if daysAgo > 0 && daysAgo <= 7 {
		reqform.QueryParams["days_ago"] = daysAgo
	}

	// API 호출
	return commonAnyCaller(ctx, tradesTicksEndPoint, reqform, &TradesTicksGetResponses{})
}
