package main

import "fmt"

func main() {
	payments := make(chan int)

	go func() {
		for _, amount := range []int{500, 1200, 300} {
			payments <- amount
		}
		close(payments)
	}()

	total := 0

	for amount := range payments {
		total += amount
	}

	fmt.Println("total received:", total)
}
