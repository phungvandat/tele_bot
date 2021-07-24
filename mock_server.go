package main

import (
	"fmt"
	"net/http"
	"os"
)

func HTTPSerer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", port), &hdl{})
}

type hdl struct {
}

func (h *hdl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("giauco_bot"))
}
