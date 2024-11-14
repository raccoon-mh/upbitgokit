package upbitapi

import (
	"context"
)

type ApikeysGetResponse struct {
	AccessKey string `json:"access_key"`
	ExpireAt  string `json:"expire_at"`
}

type ApikeysGetResponses []ApikeysGetResponse

func ApikeysGet(ctx context.Context) (*ApikeysGetResponses, error) {
	return commonAnyCaller(ctx, apiKeysEndPoint, RequestForm{}, &ApikeysGetResponses{})
}
