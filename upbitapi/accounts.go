package upbitapi

import (
	"context"
)

type AccountsGetResponse struct {
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	Balance             string `json:"balance"`
	Currency            string `json:"currency"`
	Locked              string `json:"locked"`
	UnitCurrency        string `json:"unit_currency"`
	HttpErr             error  `json:"error"`
}

type AccountsGetResponses []AccountsGetResponse

// 전체 계좌 조회
// https://docs.upbit.com/reference/%EC%A0%84%EC%B2%B4-%EA%B3%84%EC%A2%8C-%EC%A1%B0%ED%9A%8C
func AccountsGet(ctx context.Context) (*AccountsGetResponses, error) {
	return commonAnyCaller(ctx, accountsEndPoint, RequestForm{}, &AccountsGetResponses{})
}
