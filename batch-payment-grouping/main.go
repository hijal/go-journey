package main

import "fmt"

type Transaction struct {
	ID     string
	Amount float64
}

func main() {
	incoming := []Transaction{
		{"TX-001", 12.50},
		{"TX-002", 8900.00},
		{"TX-003", 45.00},
		{"TX-004", 15200.75},
		{"TX-005", 3.99},
	}

	const highValueThreshold = 5000.0
	small := make([]Transaction, 0, len(incoming))
	highValue := make([]Transaction, 0)

	for _, tx := range incoming {
		if tx.Amount >= highValueThreshold {
			highValue = append(highValue, tx)
		} else {
			small = append(small, tx)
		}
	}

	fmt.Printf("small transactions: %d\n", len(small))
	for _, tx := range small {
		fmt.Printf("  %s: $%.2f\n", tx.ID, tx.Amount)
	}

	fmt.Printf("high-value(manual review): %d\n", len(highValue))
	for _, tx := range highValue {
		fmt.Printf("  %s: $%.2f\n", tx.ID, tx.Amount)
	}
}
