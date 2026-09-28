package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	shutdown := make(chan struct{})

	var wg sync.WaitGroup

	for id := range 10 {
		wg.Go(func() {
			<-shutdown

			fmt.Println("worker", id, "stopping")
		})
	}
	time.Sleep(time.Millisecond * 50)
	close(shutdown)
	wg.Wait()
}
