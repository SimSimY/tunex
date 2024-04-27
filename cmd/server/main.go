package main

import (
	"flag"
	"github.com/SimSimY/tunex/config"
	request_dumper "github.com/SimSimY/tunex/request-dumper"
	"log"
	"net/http"
	"strconv"
	"sync"

	_ "net/http/pprof"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  256,
	WriteBufferSize: 256,
	WriteBufferPool: &sync.Pool{},
}

func process(c *websocket.Conn) {
	defer c.Close()
	for {
		log.Printf("waiting for msg")

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	// Process connection in a new goroutine
	go process(c)
	// Let the http handler return, the 8k buffer created by it will be garbage collected
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", handler)
	http.Handle("/", request_dumper.RequestDumper{})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.SERVER_PORT), nil))
}
