package main

import (
	"fmt"
	"time"
)

func fetchExchangeRate(delay time.Duration) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		time.Sleep(delay)
		ch <- 122.40
	}()

	return ch
}

func main() {
	select {
	case rate := <-fetchExchangeRate(50 * time.Millisecond):
		fmt.Println("USD->BDT:", rate)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("rate service timed out!")
	}
}
