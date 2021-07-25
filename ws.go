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
		SendBotMessage(ErrMsg("websocket.DefaultDialer.DialContext", err))
		return nil, nil
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
		conn.Close()
		log.Println("websocket disconnected")
	}
}
