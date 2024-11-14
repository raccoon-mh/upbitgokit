package upbitapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func OredersChance(serverURL string, market string) (*AccountInfos, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", serverURL+"/v1/orders/chance", nil)
	if err != nil {
		panic(err)
	}

	params := url.Values{}
	params.Add("market", market)
	paramsEncoded := params.Encode()

	req.Header.Set("Authorization", "Bearer "+GenerateNewKey(paramsEncoded))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		panic(ioerr)
	}

	// result := &AccountInfos{}
	// jsonerr := json.Unmarshal(respBody, result)
	// if jsonerr != nil {
	// 	panic(jsonerr)
	// }
	fmt.Println(string(respBody))
	return nil, nil
}
