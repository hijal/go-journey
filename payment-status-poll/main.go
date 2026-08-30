package main

import (
	"fmt"
	"time"
)

func checkPaymentStatus(statusCh chan<- string, delay time.Duration) {
	time.Sleep(delay)
	statusCh <- "confirmed"
}

func pollWithTimeout(delay, timeout time.Duration) string {
	statusCh := make(chan string)

	go checkPaymentStatus(statusCh, delay)

	select {
	case status := <-statusCh:
		return status
	case <-time.After(timeout):
		return "time-out"
	}
}

func main() {
	fast := pollWithTimeout(50*time.Millisecond, 200*time.Millisecond)
	fmt.Println("fast payment result:", fast)

	slow := pollWithTimeout(500*time.Millisecond, 200*time.Millisecond)
	fmt.Println("slow payment result:", slow)
}
