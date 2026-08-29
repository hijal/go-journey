package main

import "fmt"

func main() {
	transactionID := "TXN-2026-0091"
	amountInCents := 150075
	_isRefunded := false

	fmt.Println("Transaction:", transactionID)
	fmt.Println("Amount (cents):", amountInCents)
	fmt.Println("Refunded:", _isRefunded)
}
