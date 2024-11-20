package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/raccoon-mh/upbitgokit/upbitapi"
)

var ctx context.Context

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cred := upbitapi.Credential{
		AccessKey: os.Getenv("UPBIT_OPEN_API_ACCESS_KEY"),
		SecretKey: os.Getenv("UPBIT_OPEN_API_SECRET_KEY"),
	}

	ctx, err = upbitapi.SetCtxCredential(context.Background(), cred)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "./upbitgokitcli",
		Short: "A sample upbitgokitcli",
	}

	// _________ accounts.go _________

	rootCmd.AddCommand(&cobra.Command{
		Use:   "accountsget",
		Short: "전체 계좌 조회",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.AccountsGet(ctx)
			if err != nil {
				fmt.Printf("err: %s", err.Error())
			}
			printRes(res)
		},
	})

	// _________ apikeys.go _________
	rootCmd.AddCommand(&cobra.Command{
		Use:   "apikeysget",
		Short: "API 키 리스트 조회",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.ApikeysGet(ctx)
			if err != nil {
				fmt.Printf("err: %s", err.Error())
			}
			printRes(res)
		},
	})

	// _________ candles.go _________
	// 초(Second) 캔들 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "candlessecondsget [market] [to] [count]",
		Short: "초(Second) 캔들 조회",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			count, _ := strconv.Atoi(args[2])
			res, err := upbitapi.CandlesSecondsGet(ctx, args[0], args[1], int32(count))
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// 분(Minute) 캔들 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "candlesminutesget [unit] [market] [to] [count]",
		Short: "분(Minute) 캔들 조회",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			unit, _ := strconv.Atoi(args[0])
			count, _ := strconv.Atoi(args[3])
			res, err := upbitapi.CandlesMinutesUnitGet(ctx, int32(unit), args[1], args[2], int32(count))
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// 일(Day) 캔들 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "candlesdaysget [market] [to] [count]",
		Short: "일(Day) 캔들 조회",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			count, _ := strconv.Atoi(args[2])
			res, err := upbitapi.CandlesDaysGet(ctx, args[0], args[1], int32(count))
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// 주(Week) 캔들 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "candlesweeksget [market] [to] [count]",
		Short: "주(Week) 캔들 조회",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			count, _ := strconv.Atoi(args[2])
			res, err := upbitapi.CandlesWeeksGet(ctx, args[0], args[1], int32(count))
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// 월(Month) 캔들 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "candlesmonthsget [market] [to] [count]",
		Short: "월(Month) 캔들 조회",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			count, _ := strconv.Atoi(args[2])
			res, err := upbitapi.CandlesMonthGet(ctx, args[0], args[1], int32(count))
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// _________ market.go _________
	// 종목 코드 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "marketallget [is_details]",
		Short: "종목 코드 조회",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			isDetails, _ := strconv.ParseBool(args[0])
			res, err := upbitapi.MarketAllGet(ctx, isDetails)
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// _________ order.go _________
	rootCmd.AddCommand(&cobra.Command{
		Use:   "orderschanceget [market]",
		Short: "주문 가능 정보 조회",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.OrdersChanceGet(ctx, args[0])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "orderget [uuid] [identifier]",
		Short: "개별 주문 조회",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.OrderGet(ctx, args[0], args[1])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "orderuuidsget [market] [order_by]",
		Short: "주문 리스트 조회",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.OrderUuidsGet(ctx, args[0], nil, nil, args[1])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "orderopenget [market] [page] [limit] [order_by]",
		Short: "체결 대기 주문 조회",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			page, _ := strconv.ParseInt(args[1], 10, 64)
			limit, _ := strconv.ParseInt(args[2], 10, 64)
			res, err := upbitapi.OrderOpenGet(ctx, args[0], "", page, limit, args[3])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "ordersclosedget [market] [state] [start_time] [end_time] [limit] [order_by]",
		Short: "종료된 주문 조회",
		Args:  cobra.ExactArgs(6),
		Run: func(cmd *cobra.Command, args []string) {
			limit, _ := strconv.Atoi(args[4])
			res, err := upbitapi.OrdersClosedGet(ctx, args[0], args[1], args[2], args[3], int32(limit), args[5])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "ordercanceldelete [uuid] [identifier]",
		Short: "주문 취소",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.OrderCancelDelete(ctx, args[0], args[1])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "orderspost [market] [side] [volume] [price] [ord_type] [identifier] [time_in_force]",
		Short: "주문하기",
		Args:  cobra.ExactArgs(7),
		Run: func(cmd *cobra.Command, args []string) {
			volume, _ := strconv.ParseFloat(args[2], 64)
			price, _ := strconv.ParseInt(args[3], 10, 64)
			res, err := upbitapi.OrdersPost(ctx, args[0], args[1], volume, price, args[4], args[5], args[6])
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// _________ status.go _________
	// 입출금 현황 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "statuswalletget",
		Short: "입출금 현황 조회",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := upbitapi.StatusWalletGet(ctx)
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	// _________ trades.go _________
	// 입출금 현황 조회
	// 최근 체결 내역 조회
	rootCmd.AddCommand(&cobra.Command{
		Use:   "tradesticksget [market] [to] [count] [cursor] [days_ago]",
		Short: "최근 체결 내역 조회",
		Args:  cobra.ExactArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			count, _ := strconv.Atoi(args[2])
			daysAgo, _ := strconv.Atoi(args[4])
			res, err := upbitapi.TradesTicksGet(ctx, args[0], args[1], int32(count), args[3], int32(daysAgo))
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			printRes(res)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion scripts",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				rootCmd.GenZshCompletion(os.Stdout)
			case "fish":
				rootCmd.GenFishCompletion(os.Stdout, true)
			case "powershell":
				rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
			default:
				fmt.Println("Invalid shell. Supported: bash, zsh, fish, powershell")
			}
		},
	})

	rootCmd.Execute()
}

func printRes(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Printf("%+v\n", data)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
