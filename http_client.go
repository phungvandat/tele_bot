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
