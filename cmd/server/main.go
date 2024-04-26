package main

import (
	"github.com/SimSimY/tunex/config"
	"log"
	"net/http"
	"strconv"
)

type DummyServer struct{}

func (h *DummyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.Handle("/", &DummyServer{})
	log.Printf("Listening on port %d", config.SERVER_PORT)
	err := http.ListenAndServe(":"+strconv.Itoa(config.SERVER_PORT), nil)

	if err != nil {
		return
	}
}
