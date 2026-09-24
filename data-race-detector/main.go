package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pageViews := 0
	for range 1000 {
		wg.Go(func() {
			pageViews++
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews)
}
