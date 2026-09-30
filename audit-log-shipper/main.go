package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	maxBatch      = 3
	flushInterval = 50 * time.Millisecond
)

func shipper(events <-chan string, done chan<- struct{}) {
	defer close(done)

	ticker := time.NewTicker(flushInterval)

	defer ticker.Stop()
	batch := make([]string, 0, maxBatch)

	flush := func(reason string) {
		if len(batch) == 0 {
			return
		}
		fmt.Printf("flush %-6s -> [%s]\n", reason, strings.Join(batch, " "))
		batch = batch[:0]
	}

	for {
		select {
		case e, ok := <-events:
			if !ok {
				flush("closed")
				return
			}
			batch = append(batch, e)
			if len(batch) == maxBatch {
				flush("size")
			}
		case <-ticker.C:
			flush("timer")
		}
	}
}

func main() {
	events := make(chan string)
	done := make(chan struct{})

	go shipper(events, done)

	for i := range 5 {
		events <- fmt.Sprintf("login#%d", i+1)
	}

	time.Sleep(80 * time.Millisecond)
	events <- "login#6"
	events <- "login#7"
	close(events)
	<-done
}
