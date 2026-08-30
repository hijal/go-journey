package main

import "fmt"

// Package-level variable
var totalProcessedCents int64

func processTransaction(amountCents int64) {
	// local variable
	fee := amountCents / 100
	totalProcessedCents += amountCents - fee
	fmt.Printf("processed %d cents (fee %d), running total now %d\n",
		amountCents, fee, totalProcessedCents)

}
func main() {
	processTransaction(10000)
	processTransaction(25000)
	processTransaction(5000)

	fmt.Println("Final ledger total:", totalProcessedCents)
}
