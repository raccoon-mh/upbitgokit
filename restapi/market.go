package restapi

import (
	"context"
)

type MarketEventCaution struct {
	PriceFluctuations            bool `json:"PRICE_FLUCTUATIONS"`
	TradingVolumeSoaring         bool `json:"TRADING_VOLUME_SOARING"`
	DepositAmountSoaring         bool `json:"DEPOSIT_AMOUNT_SOARING"`
	GlobalPriceDifferences       bool `json:"GLOBAL_PRICE_DIFFERENCES"`
	ConcentrationOfSmallAccounts bool `json:"CONCENTRATION_OF_SMALL_ACCOUNTS"`
}

type MarketEvent struct {
	Warning bool               `json:"warning"`
	Caution MarketEventCaution `json:"caution"`
}

type MarketInfo struct {
	Market        string      `json:"market"`
	KoreanName    string      `json:"korean_name"`
	EnglishName   string      `json:"english_name"`
	MarketWarning string      `json:"market_warning"`
	MarketEvent   MarketEvent `json:"market_event"`
}

type MarketInfos []MarketInfo

func MarketAll(ctx context.Context) (*MarketInfos, error) {
	var accountInfos MarketInfos
	accountInfos, err := commonRestGet(ctx, marketAllEndPoint, accountInfos)
	if err != nil {
		return nil, err
	}
	return &accountInfos, nil
}
