package upbitapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	cfg "upbitapi/config"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var commonClient = &http.Client{
	Timeout: 10 * time.Second,
}

type RequestForm struct {
	PathParams  map[string]interface{}
	QueryParams map[string]interface{}
	RequestBody map[string]interface{}
}

func commonAnyCaller[AnyStruct any](ctx context.Context, mep string, rf RequestForm, ast AnyStruct) (AnyStruct, error) {
	p := strings.Split(mep, ":")
	if len(p) != 2 {
		return ast, fmt.Errorf("invalid endpoint")
	}
	u, m := p[0], strings.ToUpper(p[1])
	switch m {
	case http.MethodGet:
		return restGet(ctx, u, rf, ast)
	case http.MethodPost:
		return restGet(ctx, u, rf, ast)
	case http.MethodDelete:
		return restGet(ctx, u, rf, ast)
	default:
		return ast, fmt.Errorf("invalid method")
	}
}

func restGet[AnyStruct any](ctx context.Context, ep string, rf RequestForm, ast AnyStruct) (AnyStruct, error) {
	req, err := http.NewRequest(http.MethodGet, serverHost+ep, nil)
	if err != nil {
		return ast, fmt.Errorf("newRequest Error : %s", err.Error())
	}
	jwtToken, err := generateSignedToken(ctx)
	if err != nil {
		return ast, err
	}
	req.Header.Set("Authorization", "Bearer "+jwtToken)
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

func generateSignedToken(ctx context.Context) (string, error) {
	cred, err := cfg.GetCtxCredential(ctx)
	if err != nil {
		log.Println("error while GetCtxCredentialcfg :", err.Error())
		return "", err
	}
	claims := jwt.MapClaims{
		"access_key": cred.AccessKey,
		"nonce":      uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(cred.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
