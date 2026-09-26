package main

import (
	"fmt"
	"sync"
	"time"
)

func notify(channel string, orderID int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("notification sent via %s for order %d\n", channel, orderID)
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for _, channel := range []string{"email", "sms", "push"} {
		wg.Go(func() {
			notify(channel, 5001)
		})
	}

	wg.Wait()

	fmt.Println("checkout finished in about", time.Since(start).Round(50*time.Millisecond))
}
