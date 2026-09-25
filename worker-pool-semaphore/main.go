package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConcurrent = 100

	sem := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	for imageID := range 1000 {
		wg.Go(func() {
			sem <- struct{}{}
			defer func() { <-sem }()

			fmt.Println("resizing image", imageID)
			time.Sleep(100 * time.Millisecond)
		})
	}

	wg.Wait()

	fmt.Println("all images resized")
}
