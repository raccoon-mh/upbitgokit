package restapi

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

func Accounts(ctx context.Context) (*AccountInfos, error) {
	var accountInfos AccountInfos
	accountInfos, err := commonRestGet(ctx, accountsEndPoint, accountInfos)
	if err != nil {
		return nil, err
	}
	return &accountInfos, nil
}
