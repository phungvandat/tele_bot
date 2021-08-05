package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	wsConn, wsClose := initWebSocketConn(ctx)
	if wsConn == nil || wsClose == nil {
		return
	}
	defer wsClose()

	Job(ctx, wsConn)
	log.Println("Finish job :)")
}

func Job(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-func() chan string {
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
		}():
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
		if listCoins[idx].ID == data.D.CR.ID {
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
		text    = "<b>"
		percent = price / base
	)

	if lossAlert && percent < float64(95)/float64(100) {
		text += "[XÃ LỖ]"
	} else if percent > float64(110)/float64(100) {
		text += "[CHỐT LỜI]"
	} else {
		return
	}
	text += fmt.Sprintf("</b>\n<b>%s</b> %s: <code>%v</code> USD\n", name, symbol, price)
	text += fmt.Sprintf("Base: <s>%v</s> USD", base)

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
