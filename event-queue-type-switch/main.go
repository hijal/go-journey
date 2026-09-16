package main

import "fmt"

type OrderPlaced struct {
	OrderID string
	Amount  int64
}

type PaymentFailed struct {
	OrderID string
	Reason  string
}

type OrderShipped struct {
	OrderID string
	Courier string
}

func handle(ev any) {
	switch e := ev.(type) {
	case OrderPlaced:
		fmt.Printf("reserve stock for %s (৳%.2f)\n", e.OrderID, float64(e.Amount)/100)
	case PaymentFailed:
		fmt.Printf("alert ops: %s failed — %s\n", e.OrderID, e.Reason)
	case OrderShipped:
		fmt.Printf("email tracking link for %s via %s\n", e.OrderID, e.Courier)
	default:
		fmt.Printf("unknown event %T — route to dead-letter queue\n", ev)
	}
}

func main() {
	queue := []any{
		OrderPlaced{OrderID: "ORD-7", Amount: 42000},
		PaymentFailed{OrderID: "ORD-8", Reason: "card declined"},
		OrderShipped{OrderID: "ORD-7", Courier: "Pathao"},
		"bad-message",
	}

	for _, ev := range queue {
		handle(ev)
	}
}
