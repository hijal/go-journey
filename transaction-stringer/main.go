package main

import "fmt"

type Transaction struct {
	ID     string
	Amount int64
	Status string
}

func (t Transaction) String() string {
	return fmt.Sprintf("txn[%s] %-7s ৳%.2f", t.ID, t.Status, float64(t.Amount)/100)
}

func main() {
	txns := []Transaction{
		{ID: "T-1001", Amount: 249900, Status: "SETTLED"},
		{ID: "T-1002", Amount: 15000, Status: "PENDING"},
		{ID: "T-1003", Amount: 8900, Status: "FAILED"},
	}

	for _, t := range txns {
		fmt.Println(t)
	}
}
