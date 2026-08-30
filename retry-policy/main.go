package main

import (
	"fmt"
	"time"
)

const maxRetries = 3
const baseBackoff = 500 * time.Millisecond

func main() {
	attempt := 0

	for attempt <= maxRetries {
		ok := doRequest(attempt)

		if ok {
			fmt.Printf("success after %d attempt(s)\n", attempt+1)
			return
		}
		attempt++
		time.Sleep(baseBackoff * time.Duration(attempt))
	}

	fmt.Println("giving up after", maxRetries, "retries")
}

func doRequest(attempt int) bool {
	fmt.Println("attempt", attempt+1)
	return attempt >= 2
}
