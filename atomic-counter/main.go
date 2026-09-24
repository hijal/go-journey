package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var pageViews atomic.Int64

	for range 1000 {
		wg.Go(func() {
			pageViews.Add(1)
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews.Load())
}
