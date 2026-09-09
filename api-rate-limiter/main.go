package main

import "fmt"

func newRateLimiter(limit int) func(client string) bool {
	used := 0
	return func(client string) bool {
		used++
		fmt.Printf("%s -> request %d/%d\n", client, used, limit)
		return used <= limit
	}
}

func main() {
	serve := newRateLimiter(3)
	serve("client-A")
	serve("client-A")
	serve("client-B")
	fmt.Println("client-B allowed?", serve("client-B"))

	burst := newRateLimiter(1)
	burst("client-C")
	fmt.Println("client-C allowed?", burst("client-C"))
}
