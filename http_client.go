package main

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func init() {
	var (
		defaultRoundTripper         = http.DefaultTransport
		defaultTransportPointer, ok = defaultRoundTripper.(*http.Transport)
	)
	if !ok {
		panic("defaultRoundTripper not an *http.Transport")
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	httpClient.Transport = &defaultTransport
}

func pingHeroku() {
	req, err := http.NewRequest(http.MethodGet, pingURL, nil)
	if err != nil {
		SendBotMessage(ErrMsg("http.NewRequest", err))
		return
	}

	res, err := httpClient.Do(req)
	if err != nil {
		SendBotMessage(ErrMsg("httpClient.Do(req)", err))
		return
	}
	defer res.Body.Close()
}
