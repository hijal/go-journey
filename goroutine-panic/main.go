package main

import (
	"fmt"
	"log/slog"
	"sync"
)

func runSafely(wg *sync.WaitGroup, task string, fn func()) {
	wg.Go(func() {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("task panicked", "task", task, "panic", r)
			}
		}()
		fn()
	})
}

func main() {
	var wg sync.WaitGroup

	runSafely(&wg, "sync-orders", func() { fmt.Println("orders synced") })
	runSafely(&wg, "sync-refunds", func() {
		var refunds map[string]int
		refunds["tx-1"] = 100
	})
	wg.Wait()
	fmt.Println("service still alive")
}
