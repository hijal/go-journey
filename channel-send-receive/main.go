package main

import "fmt"

func main() {
	receipts := make(chan string)

	go func() {
		receipts <- "RCPT-9001"
	}()

	r := <-receipts
	fmt.Println("got receipt:", r)
}
