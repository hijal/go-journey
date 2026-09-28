package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(30 * time.Millisecond)
	defer ticker.Stop()

	deadline := time.After(100 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("health check")
		case <-deadline:
			fmt.Println("stop checking")
			return
		}
	}
}
