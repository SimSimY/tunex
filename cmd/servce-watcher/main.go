package main

import (
	"context"
	"github.com/SimSimY/tunex/config"
	tm "github.com/buger/goterm"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/exp/constraints"
	"math/rand/v2"
	"time"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func watchStatus(ctx context.Context, checks []*HttpUpTest) {
	tm.Clear() // Clear current screen

	for {
		sleepTime := 500*time.Millisecond + time.Millisecond*time.Duration(rand.IntN(1000))

		select {
		case <-ctx.Done():
			return

		case <-time.After(sleepTime):
		}
		tm.MoveCursor(1, 1)
		//tm.Flush() // Call it every time at the end of rendering
		//tm.MoveCursor(1, 1)
		// Print 80 clean lines
		//for i := 0; i < 80; i++ {
		//	tm.Println("")
		//}
		tm.Flush()
		tm.MoveCursor(1, 1)
		t := table.NewWriter()
		t.SetStyle(table.StyleLight)
		t.AppendHeader(table.Row{"Name", "Port", "Status", "Last Check Time", "State"})

		for _, check := range checks {
			t.AppendRow([]interface{}{check.Name, check.Port, check.Status, check.LastCheckTime.Format(time.RFC3339), check.State[0:min(len(check.State), 60)]})
		}

		tm.Println(t.Render())
		tm.Printf("Last update: %v", time.Now())

		tm.Flush() // Call it every time at the end of rendering

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
