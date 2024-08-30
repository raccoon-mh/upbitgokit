package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"testResource/handler"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	accessKey := os.Getenv("UPBIT_OPEN_API_ACCESS_KEY")
	secretKey := os.Getenv("UPBIT_OPEN_API_SECRET_KEY")
	serverURL := "https://api.upbit.com"

	payload := jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Upbit Management")
		fmt.Println("0. Quit")
		fmt.Println("1. Accounts")

		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 0:
				return

			case 1:
				result, _ := handler.Accounts(serverURL, tokenString)
				for _, res := range *result {
					PrintStructFields(res)
				}
			}
		}
	}

}

func PrintStructFields(s interface{}) {
	output := fmt.Sprintf("%+v", s)
	output = strings.ReplaceAll(output, " ", "\n-")
	fmt.Println("==================================")
	fmt.Println(output)
	fmt.Println("==================================")
}
