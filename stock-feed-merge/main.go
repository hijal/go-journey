package main

import (
	"fmt"
	"time"
)

type quote struct {
	exchange string
	symbol   string
	price    float64
}

func feed(exchange string, start time.Duration, prices []float64) <-chan quote {
	out := make(chan quote)

	go func() {
		defer close(out)
		time.Sleep(start)

		for _, p := range prices {
			out <- quote{exchange: exchange, symbol: "GP", price: p}
			time.Sleep(40 * time.Millisecond)
		}
	}()
	return out
}

func main() {
	dse := feed("DSE", 0, []float64{310.5, 311.0, 309.8})
	cse := feed("CSE", 20*time.Millisecond, []float64{310.2, 310.9})

	best := 0.0

	for dse != nil || cse != nil {
		select {
		case q, ok := <-dse:
			if !ok {
				dse = nil
				fmt.Println("DSE feed closed")
				continue
			}
			fmt.Printf("%s %s %.1f\n", q.exchange, q.symbol, q.price)
			best = max(best, q.price)
		case q, ok := <-cse:
			if !ok {
				cse = nil
				fmt.Println("CSE feed closed")
				continue
			}
			fmt.Printf("%s %s %.1f\n", q.exchange, q.symbol, q.price)
			best = max(best, q.price)
		}
	}
	fmt.Printf("best price seen: %.1f\n", best)
}
