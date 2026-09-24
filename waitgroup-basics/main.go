package main

import (
	"fmt"
	"sync"
	"time"
)

func generateInvoice(orderID int) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("invoice ready for order", orderID)
}

func main() {
	var wg sync.WaitGroup

	// classic pattern
	for _, id := range []int{101, 102, 103} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			generateInvoice(id)
		}()
	}
	wg.Wait()
	fmt.Println("all invoices done with classic pattern")
	fmt.Println()

	// modern pattern
	for _, id := range []int{105, 106, 107} {
		wg.Go(func() {
			generateInvoice(id)
		})
	}
	wg.Wait()

	fmt.Println("all invoices done with modern pattern")
}
