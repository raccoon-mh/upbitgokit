package upbitapi

import (
	"context"
)

type StatusWalletGetResponse struct {
	Currency            string `json:"currency"`
	WalletState         string `json:"wallet_state"`
	BlockState          string `json:"block_state"`
	BlockHeight         int    `json:"block_height,omitempty"` // 기본값 0
	BlockUpdatedAt      string `json:"block_updated_at"`
	BlockElapsedMinutes int    `json:"block_elapsed_minutes,omitempty"` // 기본값 0
	NetType             string `json:"net_type"`
	NetworkName         string `json:"network_name"`
}

type StatusWalletGetResponses []StatusWalletGetResponse

// 입출금 현황
// https://docs.upbit.com/reference/%EC%9E%85%EC%B6%9C%EA%B8%88-%ED%98%84%ED%99%A9
func StatusWalletGet(ctx context.Context) (*StatusWalletGetResponses, error) {
	return commonAnyCaller(ctx, apiKeysEndPoint, RequestForm{}, &StatusWalletGetResponses{})
}
