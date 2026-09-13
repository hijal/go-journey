package main

import (
	"fmt"
	"slices"
)

func main() {
	prices := []float64{24.99, 129.00, 7.25, 45.50}

	var total float64
	for _, p := range prices {
		total += p
	}

	expensive := slices.Max(prices)
	avg := total / float64(len(prices))

	fmt.Printf("subtotal: %.2f\n", total)
	fmt.Printf("most expensive item: %.2f\n", expensive)
	fmt.Printf("average price: %.2f\n", avg)
}
