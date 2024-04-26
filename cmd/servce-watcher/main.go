package main

import (
	"context"
	"github.com/SimSimY/tunex/config"
	tm "github.com/buger/goterm"
	"github.com/jedib0t/go-pretty/v6/table"
	"math/rand/v2"
	"time"
)

func watchStatus(ctx context.Context, checks []*HttpUpTest) {
	tm.Clear() // Clear current screen

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		tm.MoveCursor(1, 1)
		t := table.NewWriter()
		for _, check := range checks {
			t.AppendRow([]interface{}{check.Name, check.Status, check.State, check.LastCheckTime})
		}

		tm.Println(t.Render())
		tm.Printf("Last update: %v", time.Now())

		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(500*time.Millisecond + time.Millisecond*time.Duration(rand.IntN(1000)))
	}
}

func main() {
	ctx := context.Background()
	checks := []*HttpUpTest{
		{
			Name:          "http-server",
			Port:          config.HTTP_SERVER_PORT,
			Status:        false,
			State:         "",
			LastCheckTime: nil,
		},
		{
			Name:          "server",
			Port:          config.SERVER_PORT,
			Status:        false,
			State:         "",
			LastCheckTime: nil,
		},
		{
			Name:          "api",
			Port:          config.API_PORT,
			Status:        false,
			State:         "",
			LastCheckTime: nil,
		},
	}
	for _, check := range checks {
		go check.Run(ctx)
	}
	go watchStatus(ctx, checks)
	<-ctx.Done()

}
