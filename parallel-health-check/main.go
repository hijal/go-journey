package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	servers := []string{"web-01", "web-02", "db-01", "cache-01"}

	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("health check %s: OK\n", server)
		}()
	}

	wg.Wait()
	fmt.Println("all health check finished")
}
