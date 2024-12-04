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
	"reflect"
	"strings"
	"time"

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

// ctx : setup 완료된 컨텍스트
// mep : method and endpoint, const.go에 정의됨
// rf : requestform - path, query, req
// ast : any response struct
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
		return restPost(ctx, u, rf, ast)
	case http.MethodDelete:
		return restDelete(ctx, u, rf, ast)
	default:
		return ast, fmt.Errorf("invalid method")
	}
}

func restGet[AnyStruct any](ctx context.Context, ep string, rf RequestForm, ast AnyStruct) (AnyStruct, error) {
	if len(rf.PathParams) > 0 {
		for key, value := range rf.PathParams {
			placeholder := fmt.Sprintf("{%s}", key)
			ep = strings.ReplaceAll(ep, placeholder, fmt.Sprintf("%v", value))
		}
	}
	reqURL := serverHost + ep

	queryEncoded, queryStr, _ := generateQueryString(rf)
	jwtToken, err := generateSignedTokenWithRequstQueryString(ctx, queryStr)
	if err != nil {
		return ast, fmt.Errorf("generateSignedToken Error : %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodGet, reqURL+queryEncoded, nil)
	if err != nil {
		return ast, fmt.Errorf("newRequest Error : %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Accept", "application/json")

	resp, err := commonClient.Do(req)
	if err != nil {
		log.Printf("commonClient Error : %s", err.Error())
		return ast, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ast, fmt.Errorf("io.ReadAll Error while %s : %s", resp.Status, err.Error())
		}
		return ast, fmt.Errorf("%s : %s", resp.Status, string(respBody))
	}
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

func restPost[AnyStruct any](ctx context.Context, ep string, rf RequestForm, ast AnyStruct) (AnyStruct, error) {
	if len(rf.PathParams) > 0 {
		for key, value := range rf.PathParams {
			placeholder := fmt.Sprintf("{%s}", key)
			ep = strings.ReplaceAll(ep, placeholder, fmt.Sprintf("%v", value))
		}
	}
	reqURL := serverHost + ep

	jsonData, err := generateRequestBody(rf)
	if err != nil {
		return ast, fmt.Errorf("generateRequestBody Error : %s", err.Error())
	}
	jwtToken, err := generateSignedTokenWithRequestBody(ctx, rf)
	if err != nil {
		return ast, fmt.Errorf("generate Signed Token With RequestBody Error : %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return ast, fmt.Errorf("newRequest Error : %s", err.Error())
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := commonClient.Do(req)
	if err != nil {
		log.Printf("commonClient Error : %s", err.Error())
		return ast, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ast, fmt.Errorf("io.ReadAll Error while %s : %s", resp.Status, err.Error())
		}
		return ast, fmt.Errorf("%s : %s", resp.Status, string(respBody))
	}

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

func restDelete[AnyStruct any](ctx context.Context, ep string, rf RequestForm, ast AnyStruct) (AnyStruct, error) {
	if len(rf.PathParams) > 0 {
		for key, value := range rf.PathParams {
			placeholder := fmt.Sprintf("{%s}", key)
			ep = strings.ReplaceAll(ep, placeholder, fmt.Sprintf("%v", value))
		}
	}
	reqURL := serverHost + ep

	jsonData, err := generateRequestBody(rf)
	if err != nil {
		return ast, fmt.Errorf("generateRequestBody Error : %s", err.Error())
	}
	jwtToken, err := generateSignedTokenWithRequestBody(ctx, rf)
	if err != nil {
		return ast, fmt.Errorf("generate Signed Token With RequestBody Error : %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodDelete, reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return ast, fmt.Errorf("newRequest Error : %s", err.Error())
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := commonClient.Do(req)
	if err != nil {
		log.Printf("commonClient Error : %s", err.Error())
		return ast, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ast, fmt.Errorf("io.ReadAll Error while %s : %s", resp.Status, err.Error())
		}
		return ast, fmt.Errorf("%s : %s", resp.Status, string(respBody))
	}

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

func generateQueryString(rf RequestForm) (string, string, error) {
	query := url.Values{}

	if len(rf.QueryParams) > 0 {
		for key, value := range rf.QueryParams {
			val := reflect.ValueOf(value)
			if val.Kind() == reflect.Slice && val.Type().Elem().Kind() == reflect.String {
				for i := 0; i < val.Len(); i++ {
					query.Add(key+"s[]", val.Index(i).String())
				}
			} else {
				query.Add(key, fmt.Sprintf("%v", value))
			}
		}
	} else {
		return "", "", fmt.Errorf("no QueryParams to generate")
	}

	queryStr := ""
	for key, values := range query {
		for _, value := range values {
			queryStr += key + "=" + value + "&"
		}
	}
	if len(queryStr) > 0 {
		queryStr = queryStr[:len(queryStr)-1]
	}

	return "?" + query.Encode(), queryStr, nil
}

func generateSignedTokenWithRequstQueryString(ctx context.Context, reqString string) (string, error) {
	cred, err := GetCtxCredential(ctx)
	if err != nil {
		return "", fmt.Errorf("error while GetCtxCredentialcfg : %s", err.Error())
	}
	claims := jwt.MapClaims{
		"access_key": cred.AccessKey,
		"nonce":      uuid.New().String(),
	}

	if reqString != "" {
		hasher := sha512.New()
		_, err := hasher.Write([]byte(reqString))
		if err != nil {
			return "", fmt.Errorf("failed to write hash: %v", err)
		}
		queryHash := hex.EncodeToString(hasher.Sum(nil))
		claims["query_hash"] = queryHash
		claims["query_hash_alg"] = "SHA512"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(cred.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func generateRequestBody(rf RequestForm) ([]byte, error) {
	jsonData, err := json.Marshal(rf.RequestBody)
	if err != nil {
		return []byte{}, fmt.Errorf("json.Marshal Error : %s", err.Error())
	}
	return jsonData, nil
}

func generateSignedTokenWithRequestBody(ctx context.Context, rf RequestForm) (string, error) {
	cred, err := GetCtxCredential(ctx)
	if err != nil {
		log.Println("error while GetCtxCredential :", err.Error())
		return "", err
	}
	claims := jwt.MapClaims{
		"access_key": cred.AccessKey,
		"nonce":      uuid.New().String(),
	}
	if len(rf.RequestBody) > 0 {
		queryParams := url.Values{}
		for key, value := range rf.RequestBody {
			queryParams.Add(key, fmt.Sprintf("%v", value))
		}
		queryString := queryParams.Encode()
		hasher := sha512.New()
		_, err := hasher.Write([]byte(queryString))
		if err != nil {
			return "", fmt.Errorf("failed to write hash: %v", err)
		}
		queryHash := hex.EncodeToString(hasher.Sum(nil))
		claims["query_hash"] = queryHash
		claims["query_hash_alg"] = "SHA512"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(cred.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func validateTimeString(to string) error {
	if to != "" {
		_, err := time.Parse("2006-01-02 15:04:05", to)
		if err != nil {
			return fmt.Errorf("error parsing time: %s", err.Error())
		}
	}
	return nil
}

func validateISO8601Format(value string) error {
	if value != "" {
		_, err := time.Parse("2006-01-02T15:04:05Z07:00", value)
		if err != nil {
			return fmt.Errorf("error parsing time: %s", err.Error())
		}
		return nil
	}
	return nil
}

func isWithin7Days(time1, time2 string) (bool, error) {
	t1, err1 := time.Parse("2006-01-02T15:04:05Z07:00", time1)
	t2, err2 := time.Parse("2006-01-02T15:04:05Z07:00", time2)
	if err1 != nil {
		return false, fmt.Errorf("invalid format: %v", err1)
	}
	if err2 != nil {
		return false, fmt.Errorf("invalid format: %v", err2)
	}
	duration := t1.Sub(t2)
	if duration < 0 {
		duration = -duration
	}
	return duration <= 7*24*time.Hour, nil
}
