package upbitapi

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"upbitapi/config"

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

func commonCaller[AnyStruct any](ctx context.Context, endpoint string, query string, ast AnyStruct) (AnyStruct, error) {
	p := strings.Split(endpoint, ":")
	if len(p) != 2 {
		return ast, fmt.Errorf("invalid endpoint")
	}
	u, m := p[0], strings.ToUpper(p[1])
	switch m {
	case http.MethodGet:
		return commonRestGet(ctx, u, query, ast)
	case http.MethodPost:
		return commonRestGet(ctx, u, query, ast)
	case http.MethodDelete:
		return commonRestGet(ctx, u, query, ast)
	default:
		return ast, fmt.Errorf("invalid method")
	}
}

func commonRestGet[AnyStruct any](ctx context.Context, endpoint string, query string, ast AnyStruct) (AnyStruct, error) {
	baseUrl, err := getBaseUrl(ctx)
	if err != nil {
		return ast, fmt.Errorf("getBaseUrl Error : %s", err.Error())
	}
	req, err := http.NewRequest(http.MethodGet, baseUrl+endpoint, nil)
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

func commonRestPost[AnyStruct any](ctx context.Context, endpoint string, body interface{}, ast AnyStruct) (AnyStruct, error) {
	baseUrl, err := getBaseUrl(ctx)
	if err != nil {
		return ast, fmt.Errorf("getBaseUrl Error : %s", err.Error())
	}

	bodyData, err := json.Marshal(body)
	if err != nil {
		return ast, fmt.Errorf("marshal Error : %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, baseUrl+endpoint, bytes.NewBuffer(bodyData))
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
	jwtToken, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return ""
	}
	authorizationToken := fmt.Sprintf("Bearer %s", jwtToken)
	return authorizationToken
}

func generateJWTWithQueryHash(ctx context.Context, query interface{}) string {
	cfg, err := config.GetCtxCredentialConfig(ctx)
	if err != nil {
		log.Println("error while GetCtxCredentialConfig :", err.Error())
		return ""
	}
	mapQuery, err := convertToMapStringString(query)
	if err != nil {
		log.Println("error while convertToMapStringString :", err.Error())
		return ""
	}
	queryStr := url.Values{}
	for key, value := range mapQuery {
		queryStr.Add(key, value)
	}
	hasher := sha512.New()
	hasher.Write([]byte(queryStr.Encode()))
	queryHash := hex.EncodeToString(hasher.Sum(nil))
	payload := jwt.MapClaims{
		"access_key":     cfg.AccessKey,
		"nonce":          uuid.New().String(),
		"query_hash":     queryHash,
		"query_hash_alg": "SHA512",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtToken, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return ""
	}
	authorizationToken := fmt.Sprintf("Bearer %s", jwtToken)
	return authorizationToken
}

func convertToMapStringString(i interface{}) (map[string]string, error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("not a map[string]interface{}")
	}
	result := make(map[string]string)
	for key, value := range m {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("value for key '%s' is not a string", key)
		}
		result[key] = strValue
	}
	return result, nil
}
