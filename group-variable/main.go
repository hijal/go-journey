package main

import "fmt"

const (
	StatusPending = iota
	StatusPaid
	StatusShipped
	StatusDelivered
	StatusCancelled
)

func main() {
	var (
		orderID   = "ORD-1023"
		status    = StatusPending
		itemCount int
	)

	itemCount = 3
	fmt.Println("Order ID:", orderID)
	fmt.Println("Status code:", status)
	fmt.Println("Item count:", itemCount)

	status = StatusPaid
	fmt.Println("Updated status code:", status)
}
