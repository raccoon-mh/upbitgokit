package upbitapi

import (
	"context"
)

type ApikeysGetResponse struct {
	AccessKey string `json:"access_key"`
	ExpireAt  string `json:"expire_at"`
}

type ApikeysGetResponses []ApikeysGetResponse

// API 키 리스트 조회
// https://docs.upbit.com/reference/open-api-%ED%82%A4-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
func ApikeysGet(ctx context.Context) (*ApikeysGetResponses, error) {
	return commonAnyCaller(ctx, apiKeysEndPoint, RequestForm{}, &ApikeysGetResponses{})
}
