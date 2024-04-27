package main

import (
	"encoding/json"
	"github.com/SimSimY/tunex/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type DummyServer struct{}

type HttpDumpResponse struct {
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
	Method  string      `json:"method"`
	URL     string      `json:"url"`
	Form    url.Values  `json:"form"`
}

func (h *DummyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf := new(strings.Builder)
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(buf, r.Body)
	dumpResponse := &HttpDumpResponse{
		Headers: r.Header,
		Body:    buf.String(),
		Method:  r.Method,
		URL:     r.URL.String(),
		Form:    r.Form,
	}
	b, _ := json.Marshal(dumpResponse)

	_, _ = w.Write(b)
}

func main() {
	http.Handle("/", &DummyServer{})
	log.Printf("Listening on port %d", config.HTTP_SERVER_PORT)
	err := http.ListenAndServe(":"+strconv.Itoa(config.HTTP_SERVER_PORT), nil)

	if err != nil {
		return
	}
}
