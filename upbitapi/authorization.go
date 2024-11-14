package upbitapi

import (
	"crypto/sha512"
	"encoding/hex"
	"net/url"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func GenerateNewKey(params string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	accessKey := os.Getenv("UPBIT_OPEN_API_ACCESS_KEY")
	secretKey := os.Getenv("UPBIT_OPEN_API_SECRET_KEY")

	payload := jwt.MapClaims{}
	if params != "" {
		params := url.Values{}
		params.Add("market", "KRW-BTC")
		queryString := params.Encode()

		// UTF-8 인코딩된 queryString 생성
		queryBytes := []byte(queryString)

		// SHA512 해시 생성
		m := sha512.New()
		m.Write(queryBytes)
		queryHash := hex.EncodeToString(m.Sum(nil))

		payload = jwt.MapClaims{
			"access_key":     accessKey,
			"nonce":          uuid.New().String(),
			"query_hash":     queryHash,
			"query_hash_alg": "SHA512",
		}
	} else {
		payload = jwt.MapClaims{
			"access_key": accessKey,
			"nonce":      uuid.New().String(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func StringToPtr(str string) *string {
	return &str
}
