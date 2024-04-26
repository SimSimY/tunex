package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpUpTest struct {
	Name          string
	Port          int
	Status        bool
	State         string
	LastCheckTime *time.Time
	ctx           context.Context
}

func (receiver *HttpUpTest) Update() {
	// run a http get request to check if the server is up
	// if the server is up, set the status to true
	// if the server is down, set the status to false
	// set the last check time to the current time

	// Do http get request to the port on localhost
	requestURL := fmt.Sprintf("http://localhost:%d", receiver.Port)
	res, err := http.Get(requestURL)
	if err != nil {
		receiver.Status = false
		receiver.State = err.Error()
	} else if res.StatusCode >= 200 && res.StatusCode <= 299 {
		receiver.Status = true
		receiver.State = ""
	} else {
		receiver.Status = false
		receiver.State = res.Status
	}
	now := time.Now()
	receiver.LastCheckTime = &now

}
func (receiver *HttpUpTest) Run(ctx context.Context) {
	receiver.ctx = ctx
	for {
		select {
		case <-receiver.ctx.Done():
			return
		default:
			receiver.Update()
			time.Sleep(1 * time.Second)
		}
	}
}
