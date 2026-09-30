package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const maxConcurrent = 2

func quoteShipping(parcelKg int) int {
	time.Sleep(300 * time.Millisecond)
	return 60 + parcelKg*20
}

func main() {
	parcels := []int{1, 3, 2, 5, 1, 4}
	quotes := make([]int, len(parcels))

	sem := make(chan struct{}, maxConcurrent)

	var inFlight, peak atomic.Int32
	var wg sync.WaitGroup

	for i, kg := range parcels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			n := inFlight.Add(1)

			for {
				p := peak.Load()
				if n <= p || peak.CompareAndSwap(p, n) {
					break
				}
			}
			quotes[i] = quoteShipping(kg)
			inFlight.Add(-1)
		}()
	}

	wg.Wait()

	for i, q := range quotes {
		fmt.Printf("parcel %d (%dkg): %d BDT\n", i+1, parcels[i], q)
	}

	fmt.Println("peak concurrent API calls:", peak.Load())
}
