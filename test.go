package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"raccoon-upbit-trader/config"
	"raccoon-upbit-trader/restapi"
)

func main() {
	serverURL := "https://api.upbit.com"

	ctx := context.Background()

	for {
		fmt.Println("")
		fmt.Println("Upbit Management")
		fmt.Println("0. Quit")
		fmt.Println("1. Accounts")
		fmt.Println("2. OredersChance")
		fmt.Println("3. MarketAll")
		fmt.Println("4. test")

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
				fmt.Println("")
				result, err := restapi.Accounts(ctx)
				if err != nil {
					log.Println(err)
				}
				for _, res := range *result {
					printStructFields(res)
				}

			case 2:
				fmt.Println("")
				restapi.OredersChance(serverURL, "KRW-BTT")
				// for _, res := range *result {
				// 	printStructFields(res)
				// }

			case 3:
				fmt.Println("")
				result, err := restapi.MarketAll(ctx)
				if err != nil {
					log.Println(err)
				}
				for _, res := range *result {
					printStructFields(res)
				}

			case 4:
				fmt.Println("")
				var err error
				ctx, err = config.LoadConfig(ctx, "default")
				if err != nil {
					log.Println(err)
				}

				cfg, err := config.GetCtxAllConfig(ctx)
				if err != nil {
					log.Println(err)
				}

				printStructFields(cfg)

			}
		}
	}

}

func printStructFields(s interface{}) {
	output := fmt.Sprintf("%+v", s)
	output = strings.ReplaceAll(output, " ", "\n-")
	fmt.Println("==================================")
	fmt.Println(output)
	fmt.Println("==================================")
}
