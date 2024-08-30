package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type AccountInfo struct {
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	Balance             string `json:"balance"`
	Currency            string `json:"currency"`
	Locked              string `json:"locked"`
	UnitCurrency        string `json:"unit_currency"`
}

type AccountInfos []AccountInfo

func Accounts(serverURL string, tokenString string) (*AccountInfos, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", serverURL+"/v1/accounts", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		panic(ioerr)
	}

	result := &AccountInfos{}
	jsonerr := json.Unmarshal(respBody, result)
	if jsonerr != nil {
		panic(jsonerr)
	}

	return result, nil
}
