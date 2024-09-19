package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"raccoon-upbit-trader/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var commonClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getBaseUrl(ctx context.Context) (string, error) {
	env, err := config.GetCtxEnvConfig(ctx)
	if err != nil {
		return "", err
	}
	return env.UpbitBaseUrl, nil
}

func commonRestGet[AnyStruct any](ctx context.Context, endpoint string, ast AnyStruct) (AnyStruct, error) {
	baseUrl, err := getBaseUrl(ctx)
	if err != nil {
		return ast, fmt.Errorf("getBaseUrl Error : %s", err.Error())
	}
	req, err := http.NewRequest("GET", baseUrl+endpoint, nil)
	if err != nil {
		return ast, fmt.Errorf("newRequest Error : %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+generateNewKeyNoParams(ctx))
	resp, err := commonClient.Do(req)
	if err != nil {
		log.Printf("commonClient Error : %s", err.Error())
		return ast, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ast, fmt.Errorf("io.ReadAll Error : %s", err.Error())
	}
	err = json.Unmarshal(respBody, &ast)
	if err != nil {
		log.Println("err respBody is :", string(respBody))
		return ast, fmt.Errorf("unmarshal Error : %s", err.Error())
	}
	return ast, nil
}

func generateNewKeyNoParams(ctx context.Context) string {
	cfg, err := config.GetCtxCredentialConfig(ctx)
	if err != nil {
		log.Println("error while GetCtxCredentialConfig :", err.Error())
		return ""
	}
	payload := jwt.MapClaims{
		"access_key": cfg.AccessKey,
		"nonce":      uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		log.Println("error while Sign :", err.Error())
		return ""
	}
	return tokenString
}
