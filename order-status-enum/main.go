package main

import "fmt"

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderPaid
	OrderShipped
	OrderDelivered
	OrderCanceled
)

func (s OrderStatus) String() string {
	names := [...]string{
		"Pending",
		"Paid",
		"Shipped",
		"Delivered",
		"Canceled",
	}

	if int(s) < 0 || int(s) >= len(names) {
		return "Unknown"
	}
	return names[s]
}

func main() {
	current := OrderPaid

	fmt.Println("current status:", current)
	fmt.Println("numeric value", int(current))

	if current == OrderPaid {
		fmt.Println("your order is ready to go!")
	}
}
