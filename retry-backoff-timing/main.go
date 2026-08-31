package main

import (
	"fmt"
	"time"
)

const (
	InitialBackoff = 200 * time.Millisecond
	MaxBackoff     = 8 * time.Second
	BaseDelaySec   = 2
)

func backoff(attempt int) time.Duration {
	d := time.Duration(attempt) * InitialBackoff
	if d > MaxBackoff {
		return MaxBackoff
	}
	return d
}

func main() {
	for attempt := range 5 {
		fmt.Printf("attempt %d -> backoff %v\n", attempt+1, backoff(attempt))
	}
}
