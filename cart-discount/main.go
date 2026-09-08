package main

import "fmt"

func main() {
	cart := []float64{1290, 349, 89, 2450, 599}

	subtotal := 0.0
	for i, price := range cart {
		subtotal += price
		fmt.Printf("item %d: %8.2f BDT\n", i, price)
	}

	discountRate := 0.0

	switch {
	case subtotal >= 5000:
		discountRate = 0.10
	case subtotal >= 2000:
		discountRate = 0.05
	}

	discount := subtotal * discountRate
	total := subtotal - discount

	fmt.Printf("\nsubtotal: %.2f, discount(%.0f%%): %.2f, payable: %.2f\n", subtotal, discountRate*100, discount, total)
}
