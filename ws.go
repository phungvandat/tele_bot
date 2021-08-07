package main

import (
	"context"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func initWSConn(
	ctx context.Context,
	u *url.URL,
	subMess interface{},
) (*websocket.Conn, func()) {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		SendBotMessage(ErrMsg("websocket.DefaultDialer.DialContext", err))
		return nil, nil
	}

	err = conn.WriteJSON(subMess)
	if err != nil {
		SendBotMessage(ErrMsg("conn.WriteJSON", err))
		return nil, nil
	}

	return conn, func() {
		conn.Close()
		log.Println("websocket disconnected")
	}
}
