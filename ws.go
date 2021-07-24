package main

import (
	"context"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func initWebSocketConn(ctx context.Context) (*websocket.Conn, func()) {
	u := url.URL{
		Scheme: "wss",
		Host:   "stream.coinmarketcap.com",
		Path:   "/price/latest",
	}

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		SendBotMessage(err.Error())
		panic(err)
	}
	log.Println("websocket connected")

	var ids = make([]uint16, 0, len(listCoins))
	for idx := range listCoins {
		ids = append(ids, listCoins[idx].ID)
	}

	err = conn.WriteJSON(map[string]interface{}{
		"method": "subscribe",
		"id":     "price",
		"data": map[string]interface{}{
			"cryptoIds": ids,
			"index":     "detail",
		},
	})
	if err != nil {
		SendBotMessage(ErrMsg("conn.WriteJSON", err))
		return nil, nil
	}

	return conn, func() {
		err := conn.Close()
		if err != nil {
			SendBotMessage(err.Error())
		}
		log.Println("websocket disconnected")
	}
}
