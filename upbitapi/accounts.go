package upbitapi

import (
	"context"
)

type AccountInfo struct {
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	Balance             string `json:"balance"`
	Currency            string `json:"currency"`
	Locked              string `json:"locked"`
	UnitCurrency        string `json:"unit_currency"`
	HttpErr             error  `json:"error"`
}

type AccountInfos []AccountInfo

func AccountsGet(ctx context.Context) (*AccountInfos, error) {
	return commonAnyCaller(ctx, accountsEndPoint, RequestForm{}, &AccountInfos{})
}
