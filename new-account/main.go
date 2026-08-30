package main

import "fmt"

func main() {
	var accountHolder string
	var balanceInCents int64
	var isActive bool

	fmt.Println("---before setup---")
	fmt.Println("Account holder name:", accountHolder)
	fmt.Println("Account balance(cents):", balanceInCents)
	fmt.Println("Account active?:", isActive)

	accountHolder = "John Doe"
	balanceInCents = 5000000
	isActive = true

	fmt.Println("---after setup---")
	fmt.Println("Account holder name:", accountHolder)
	fmt.Println("Account balance(cents):", balanceInCents)
	fmt.Println("Account active?:", isActive)
}
