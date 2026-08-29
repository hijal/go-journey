package main

import "fmt"

const feePercentage = 2.5

func calculateFee(amount float64) float64 {
	return amount * feePercentage / 100
}

func main() {
	var transactionAmount float64 = 4500.0
	fee := calculateFee(transactionAmount)
	total := transactionAmount + fee
	fmt.Println("Amount:", transactionAmount)
	fmt.Println("Fee:", fee)
	fmt.Println("Total:", total)
}
