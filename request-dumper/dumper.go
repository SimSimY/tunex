package request_dumper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpDumpResponse struct {
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
	Method  string      `json:"method"`
	URL     string      `json:"url"`
	Form    url.Values  `json:"form"`
}

type RequestDumper struct {
}

func (d RequestDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
