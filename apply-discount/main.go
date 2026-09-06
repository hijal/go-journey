package main

import "fmt"

func applyDiscount(price, percent float64) float64 {
	price = price - (price * percent / 100)
	fmt.Printf(" inside function: %.2f\n", price)
	return price
}

func main() {
	listPrice := 1000.0
	fmt.Printf("before call: %.2f\n", listPrice)

	final := applyDiscount(listPrice, 15)
	fmt.Printf("after call:  %.2f\n", listPrice)
	fmt.Printf("final price: %.2f\n", final)
}
