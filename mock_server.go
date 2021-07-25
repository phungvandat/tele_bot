package main

import (
	"log"
	"net/http"
)

type hdl struct {
}

func (h *hdl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s, Endpoint: %s\n", r.Method, r.RequestURI)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("giauco_bot"))
}
