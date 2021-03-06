package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

var (
	base                   float64
	botToken               string
	botID                  string
	symbol                 string
	env                    string
	isProduction           bool
	listCoins              []CoinInfo
	wg                     = sync.WaitGroup{}
	mentions               string
	pingURL                string
	lossAlert              bool
	profitAlert            bool
	profitTargetPercentage uint16 = 110
	exchangePlatform       string = "coinmarketcap"
	binanceWSID            uint
)

const BINANCE = "binance"

func init() {
	godotenv.Load()
	botToken = os.Getenv("TELE_BOT_TOKEN")
	botID = os.Getenv("TELE_BOT_ID")
	symbol = os.Getenv("SYMBOL")
	basePrice := os.Getenv("BASE_PRICE")
	base, _ = strconv.ParseFloat(basePrice, 64)
	env = os.Getenv("ENV")
	isProduction = env == "production"
	mentions = os.Getenv("MENTIONS")
	pingURL = os.Getenv("HEROKU_URL")
	lossAlert = strings.ToLower(os.Getenv("LOSS_ALERT")) == "true"
	profitAlert = strings.ToLower(os.Getenv("PROFIT_ALERT")) == "true"
	profitPercent, _ := strconv.ParseUint(os.Getenv("PROFIT_PERCENTAGE"), 10, 16)
	if profitPercent > 100 {
		profitTargetPercentage = uint16(profitPercent)
	}
	exchangePlatformEnv := os.Getenv("EXCHANE_PLATFOM")
	if exchangePlatformEnv != "" {
		exchangePlatform = strings.ToLower(exchangePlatformEnv)
	}
}

func main() {
	defer func() {
		SendBotMessage(mentions + " Bot stopped. PP!!")
	}()

	listCoins = mapSymbolCoins[symbol]
	if len(listCoins) == 0 {
		SendBotMessage(fmt.Sprintf("Symbol: %s not found", symbol))
		return
	}

	// Setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			SendBotMessage(ErrMsg("time.LoadLocation", err))
			os.Exit(1)
		}
		time.Local = loc
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				DoJob(ctx)
			}
		}
	}()

	SendBotMessage(startMess())

	errChn := make(chan error)
	go func() {
		stopChn := make(chan os.Signal, 1)
		signal.Notify(stopChn, syscall.SIGTERM, syscall.SIGINT)
		log.Printf("exit by signal: %v\n", <-stopChn)
		errChn <- nil
	}()

	srv := http.Server{
		Handler: &hdl{},
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
	}
	go func() {
		errChn <- srv.ListenAndServe()
	}()

	err := <-errChn
	if err != nil {
		SendBotMessage(ErrMsg("err_chn", err))
	}

	err = srv.Shutdown(context.Background())
	if err != nil {
		SendBotMessage(ErrMsg("srv.Shutdown", err))
	}

	cancel()
	wg.Wait()
}

func ErrMsg(what string, err error) string {
	msg := "Error: " + what + " - " + err.Error()
	if mentions != "" {
		msg = mentions + " " + msg
	}
	return msg
}
