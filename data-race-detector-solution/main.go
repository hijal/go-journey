package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		pageViews int
	)

	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			pageViews++
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews)
}
