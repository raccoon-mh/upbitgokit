package upbitapi

import (
	"context"
	"fmt"
)

type MarketEvent struct {
	Warning bool            `json:"warning"` // 유의종목 지정 여부
	Caution map[string]bool `json:"caution"` // 주의종목 지정 여부 상세 타입 (배열 또는 객체)
}

type MarketAllGetResponse struct {
	Market        string      `json:"market"`         // 업비트에서 제공하는 시장 정보
	KoreanName    string      `json:"korean_name"`    // 거래 대상 디지털 자산 한글명
	EnglishName   string      `json:"english_name"`   // 거래 대상 디지털 자산 영문명
	MarketWarning string      `json:"market_warning"` // 유의 종목 여부 (NONE, CAUTION)
	MarketEvent   MarketEvent `json:"market_event"`   // 시장 경보 정보
}

type MarketAllGetResponses []MarketAllGetResponse

// 종목 코드 조회
// is_details
// boolean : Defaults to false
// 유의종목 필드과 같은 상세 정보 노출 여부 (선택 파라미터)
// https://docs.upbit.com/reference/%EB%A7%88%EC%BC%93-%EC%BD%94%EB%93%9C-%EC%A1%B0%ED%9A%8C
func MarketAllGet(ctx context.Context, isdetail bool) (*MarketAllGetResponses, error) {
	reqform := RequestForm{
		QueryParams: map[string]interface{}{
			"is_details": fmt.Sprintf("%t", isdetail),
		},
	}
	return commonAnyCaller(ctx, marketAllEndPoint, reqform, &MarketAllGetResponses{})
}
