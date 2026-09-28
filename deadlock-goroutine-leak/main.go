package main

import "fmt"

func main() {
	orders := make(chan string)

	orders <- "ORD-1"

	fmt.Println(<-orders)
}
