package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
)

var (
	base      float64
	botToken  string
	botID     string
	symbol    string
	listCoins []CoinInfo
)

func init() {
	godotenv.Load()
	botToken = os.Getenv("TELE_BOT_TOKEN")
	botID = os.Getenv("TELE_BOT_ID")
	symbol = os.Getenv("SYMBOL")
	basePrice := os.Getenv("BASE_PRICE")
	base, _ = strconv.ParseFloat(basePrice, 64)
}

func main() {
	defer func() {
		SendBotMessage("Bot stopped. PP!!")
	}()

	listCoins = mapSymbolCoins[symbol]
	if len(listCoins) == 0 {
		SendBotMessage(fmt.Sprintf("Symbol: %s not found", symbol))
		return
	}

	go func() {
		for {
			DoJob()
		}
	}()

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
}

func ErrMsg(what string, err error) string {
	return "Error: " + what + " - " + err.Error()
}
