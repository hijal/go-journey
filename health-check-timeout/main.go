package main

import (
	"fmt"
	"time"
)

func pingService(latency time.Duration) <-chan time.Duration {
	out := make(chan time.Duration, 1)

	go func() {
		time.Sleep(latency)
		out <- latency
	}()
	return out
}

func main() {
	const budget = 150 * time.Millisecond

	services := map[string]time.Duration{
		"auth-api":    60 * time.Millisecond,
		"billing-api": 900 * time.Millisecond,
	}

	for _, name := range []string{"auth-api", "billing-api"} {
		select {
		case took := <-pingService(services[name]):
			fmt.Printf("%-12s OK (%v)\n", name, took)
		case <-time.After(budget):
			fmt.Printf("%-12s TIMEOUT (over %v)\n", name, budget)
		}
	}
}
