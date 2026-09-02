package main

import "fmt"

func main() {
	var balance float64 = 4500.0
	var requestAmount float64 = 6000.0

	if requestAmount <= balance {
		balance -= requestAmount
		fmt.Printf("Withdrawal successful. New balance: %.2f\n", balance)
	} else {
		fmt.Printf("Insufficient funds: requested %.2f, available %.2f\n", requestAmount, balance)
	}
}
