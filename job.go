package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func DoJob(ctx context.Context) {
	wg.Add(1)
	defer wg.Done()

	// ping keep alive heroku
	go pingHeroku()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	log.Println("Start job :)")
	var (
		host, path = "stream.coinmarketcap.com", "/price/latest"
		subMess    = map[string]interface{}{}
	)
	if exchangePlatform == BINANCE {
		host, path = "stream.binance.com", "/stream"
		subMess = map[string]interface{}{
			"method": "SUBSCRIBE",
			"id":     BINANCE_WS_ID,
			"params": []string{fmt.Sprintf("%s@aggTrade", strings.ToLower(listCoins[0].BinanceUSDT))},
		}
	} else {
		host, path = "stream.coinmarketcap.com", "/price/latest"
		var ids = make([]uint16, 0, len(listCoins))
		for idx := range listCoins {
			ids = append(ids, listCoins[idx].ID)
		}

		subMess = map[string]interface{}{
			"method": "subscribe",
			"id":     "price",
			"data": map[string]interface{}{
				"cryptoIds": ids,
				"index":     "detail",
			},
		}
	}

	wsConn, wsClose := initWSConn(ctx, &url.URL{
		Scheme: "wss",
		Host:   host,
		Path:   path,
	}, subMess)
	if wsConn == nil || wsClose == nil {
		return
	}
	log.Println("websocket listening...")
	defer wsClose()

	Job(ctx, wsConn)
	log.Println("Finish job :)")
}

func Job(ctx context.Context, conn *websocket.Conn) {
	var msgFunc = func() chan string {
		errMsg := make(chan string)
		go func() {
			_, message, err := conn.ReadMessage()
			if err != nil {
				errMsg <- ErrMsg("conn.ReadMessage", err)
				return
			}

			data := &Data{}
			err = json.Unmarshal(message, data)
			if err != nil {
				errMsg <- ErrMsg("json.Unmarshal(message, data)", err)
				return
			}
			Alert(data)
			errMsg <- ""
		}()

		return errMsg
	}
	if exchangePlatform == BINANCE {
		msgFunc = func() chan string {
			errMsg := make(chan string)
			go func() {
				_, message, err := conn.ReadMessage()
				if err != nil {
					errMsg <- ErrMsg("conn.ReadMessage", err)
					return
				}

				data := &BinanceData{}
				err = json.Unmarshal(message, data)
				if err != nil {
					errMsg <- ErrMsg("json.Unmarshal(message, data)", err)
					return
				}

				if data.ID == BINANCE_WS_ID && data.Result == nil {
					errMsg <- ""
					return
				}

				var alertData = &Data{
					ID: data.Data.S,
				}
				price, err := strconv.ParseFloat(data.Data.P, 64)
				if err != nil {
					errMsg <- ErrMsg("strconv.ParseFloat(data.Data.P, 64)"+data.Data.P+data.Stream, err)
					return
				}
				alertData.D.CR.Price = price
				Alert(alertData)
				errMsg <- ""
			}()

			return errMsg
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-msgFunc():
			if msg != "" {
				SendBotMessage(msg)
				return
			}
		}
	}
}

func Alert(data *Data) {
	foundIdx := -1
	for idx := range listCoins {
		if listCoins[idx].ID == data.D.CR.ID ||
			listCoins[idx].BinanceUSDT == data.ID {
			foundIdx = idx
			break
		}
	}

	if foundIdx == -1 {
		return
	}

	price := data.D.CR.Price
	if listCoins[foundIdx].LastAlertSuccess &&
		listCoins[foundIdx].LastPrice == price {
		return
	}

	name := listCoins[foundIdx].Name

	var (
		text          = "<b>"
		percent       = price / base
		lostPercent   = float64(95) / float64(100)
		profitPercent = float64(profitTargetPercentage) / float64(100)
	)

	if lossAlert && percent < lostPercent {
		text += fmt.Sprintf("[XÃ LỖ] %d%%", int(100-(percent*100)))
	} else if profitAlert && percent > profitPercent {
		text += fmt.Sprintf("[CHỐT LỜI] %d%%", int((percent*100)-100))
	} else {
		return
	}
	text += fmt.Sprintf("</b>\n<b>%s - %s</b>: <code>%v$</code>\n", name, symbol, price)
	text += fmt.Sprintf("Base: <s>%v$</s>", base)

	sendSuccess := SendBotMessage(text)
	listCoins[foundIdx].LastAlertSuccess = sendSuccess
	if sendSuccess {
		listCoins[foundIdx].LastPrice = price
	}
}

type Data struct {
	ID string `json:"id"`
	D  struct {
		CR struct {
			ID    uint16  `json:"id"`
			Price float64 `json:"p"`
		} `json:"cr"`
	} `json:"d"`
}

type BinanceData struct {
	Data struct {
		// E: 1628335411810 `json:"E"`
		// M: true `json:"M"`
		// T: 1628335411808 `json:"T"`
		// a: 3648214 `json:"a"`
		// e: "aggTrade" `json:"e"`
		// f: 5942568 `json:"f"`
		// l: 5942568 `json:"l"`
		// m: true `json:"m"`
		// p: "20000.98900000" `json:"p"`
		// q: "20000.97700000" `json:"q"`
		// s: "BTCUSDT" `json:"s"`
		P string `json:"p"`
		S string `json:"s"`
	} `json:"data"`
	Stream string      `json:"stream"`
	Result interface{} `json:"result"`
	ID     uint        `json:"id"`
}
