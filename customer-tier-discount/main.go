package main

import "fmt"

func main() {
	customerTier := "gold"
	listPrice := 4000.0

	var applyDiscount func(float64) float64

	if customerTier == "gold" {
		applyDiscount = func(f float64) float64 {
			return f * 0.85
		}
	} else if customerTier == "silver" {
		applyDiscount = func(f float64) float64 {
			return f * 0.92
		}
	} else {
		applyDiscount = func(f float64) float64 {
			return f
		}
	}

	finalPrice := applyDiscount(listPrice)

	fmt.Printf("tier: %s, final price: %.2f\n", customerTier, finalPrice)
}
