package main

import (
	"errors"
	"fmt"
	"time"
)

var errTimeout = errors.New("timed out")

func sendSMS(provider string, latency time.Duration) <-chan string {
	ack := make(chan string, 1)

	go func() {
		time.Sleep(latency)
		ack <- provider + ": delivered"
	}()

	return ack
}

func deliver(provider string, latency, timeout time.Duration) error {
	select {
	case msg := <-sendSMS(provider, latency):
		fmt.Println(msg)
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("sms via %s: %w", provider, errTimeout)
	}
}

func main() {
	for _, p := range []struct {
		name    string
		latency time.Duration
	}{
		{"fast-sms", 20 * time.Millisecond},
		{"slow-sms", 300 * time.Millisecond},
	} {
		if err := deliver(p.name, p.latency, 300 * time.Millisecond); err != nil {
			fmt.Println("error:", err, "| retry later?", errors.Is(err, errTimeout))
		}
	}
}
