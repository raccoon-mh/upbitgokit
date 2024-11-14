package upbitapi

import (
	"context"
	"fmt"
	"time"
)

type CandlesGetResponse struct {
	Market               string  `json:"market"`                  // 시장 정보
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`    // UTC 기준 캔들 시간
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`    // KST 기준 캔들 시간
	OpeningPrice         float64 `json:"opening_price"`           // 시가 (기본값 0)
	HighPrice            float64 `json:"high_price"`              // 고가 (기본값 0)
	LowPrice             float64 `json:"low_price"`               // 저가 (기본값 0)
	TradePrice           float64 `json:"trade_price"`             // 종가 (기본값 0)
	Timestamp            int64   `json:"timestamp"`               // 타임스탬프 (기본값 0), 큰 숫자는 int64로 처리
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  // 누적 거래 금액 (기본값 0)
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` // 누적 거래량 (기본값 0)
	Unit                 int     `json:"unit"`                    // 단위 (기본값 0)
}

type CandlesGetResponses []CandlesGetResponse

func parseTimeString(timeStr string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %s", err.Error())
	}
	return parsedTime, nil
}

// 초(Second) 캔들
// market (string):
//   - 종목 코드 (예: KRW-BTC)
//
// to (string):
//   - 마지막 캔들 시각 (exclusive)
//   - 비워서 요청 시 가장 최근 캔들을 반환
//
// count (int32):
//   - 캔들 개수 (최대 200개까지 요청 가능)
//
// https://docs.upbit.com/reference/%EC%B4%88second-%EC%BA%94%EB%93%A4
func CandlesSecondsGet(ctx context.Context, market string, to string, count int32) (*CandlesGetResponses, error) {
	_, err := parseTimeString(to)
	if err != nil {
		return nil, fmt.Errorf("error parsing time : %s", err.Error())
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"market": market,
			"to":     to,
			"count":  count,
		},
	}
	return commonAnyCaller(ctx, candlesSecondsEndPoint, reqform, &CandlesGetResponses{})
}

// 분(Minute) 캔들
// unit (int32):
//   - 분 단위
//   - 가능한 값: 1, 3, 5, 10, 15, 30, 60, 240
//
// market (string):
//   - 종목 코드 (예: KRW-BTC)
//
// to (string):
//   - 마지막 캔들 시각 (exclusive)
//   - 비워서 요청 시 가장 최근 캔들을 반환
//
// count (int32):
//   - 캔들 개수 (최대 200개까지 요청 가능)
//
// https://docs.upbit.com/reference/%EB%B6%84minute-%EC%BA%94%EB%93%A4-1
func CandlesMinutesUnitGet(ctx context.Context, unit int32, market string, to string, count int32) (*CandlesGetResponses, error) {
	_, err := parseTimeString(to)
	if err != nil {
		return nil, fmt.Errorf("error parsing time : %s", err.Error())
	}
	reqform := RequestForm{
		PathParams: map[string]interface{}{
			"unit": unit,
		},
		QueryParams: map[string]interface{}{
			"market": market,
			"to":     to,
			"count":  count,
		},
	}
	return commonAnyCaller(ctx, candlesMinutesUnitEndPoint, reqform, &CandlesGetResponses{})
}

// 일(Day) 캔들
// market (string):
//   - 종목 코드 (예: KRW-BTC)
//
// to (string):
//   - 마지막 캔들 시각 (exclusive)
//   - 비워서 요청 시 가장 최근 캔들을 반환
//
// count (int32):
//   - 캔들 개수 (최대 200개까지 요청 가능)
//
// https://docs.upbit.com/reference/%EB%B6%84minute-%EC%BA%94%EB%93%A4-1
func CandlesDaysGet(ctx context.Context, market string, to string, count int32) (*CandlesGetResponses, error) {
	_, err := parseTimeString(to)
	if err != nil {
		return nil, fmt.Errorf("error parsing time : %s", err.Error())
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"market": market,
			"to":     to,
			"count":  count,
		},
	}
	return commonAnyCaller(ctx, candlesDaysEndPoint, reqform, &CandlesGetResponses{})
}

// 주(Week) 캔들
// market (string):
//   - 종목 코드 (예: KRW-BTC)
//
// to (string):
//   - 마지막 캔들 시각 (exclusive)
//   - 비워서 요청 시 가장 최근 캔들을 반환
//
// count (int32):
//   - 캔들 개수 (최대 200개까지 요청 가능)
//
// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
func CandlesWeeksGet(ctx context.Context, market string, to string, count int32) (*CandlesGetResponses, error) {
	_, err := parseTimeString(to)
	if err != nil {
		return nil, fmt.Errorf("error parsing time : %s", err.Error())
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"market": market,
			"to":     to,
			"count":  count,
		},
	}
	return commonAnyCaller(ctx, candlesWeeksEndPoint, reqform, &CandlesGetResponses{})
}

// 월(Month) 캔들
// market (string):
//   - 종목 코드 (예: KRW-BTC)
//
// to (string):
//   - 마지막 캔들 시각 (exclusive)
//   - 비워서 요청 시 가장 최근 캔들을 반환
//
// count (int32):
//   - 캔들 개수 (최대 200개까지 요청 가능)
//
// https://docs.upbit.com/reference/%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1
func CandlesMonthGet(ctx context.Context, market string, to string, count int32) (*CandlesGetResponses, error) {
	_, err := parseTimeString(to)
	if err != nil {
		return nil, fmt.Errorf("error parsing time : %s", err.Error())
	}
	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"market": market,
			"to":     to,
			"count":  count,
		},
	}
	return commonAnyCaller(ctx, candlesMonthsEndPoint, reqform, &CandlesGetResponses{})
}
