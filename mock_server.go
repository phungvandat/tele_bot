package main

import (
	"net/http"
)

type hdl struct {
}

func (h *hdl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("giauco_bot"))
}
