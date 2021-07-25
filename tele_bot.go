package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func SendBotMessage(text string) bool {
	if !isProduction {
		text = "<b>" + env + "</b> - " + text
	}
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.telegram.org/%s/sendMessage", botToken), nil)
	q := url.Values{}
	q.Add("chat_id", botID)
	q.Add("text", text)
	q.Add("parse_mode", "html")
	r.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(r)
	if err != nil {
		log.Printf("SendBotMessage %v\n", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
