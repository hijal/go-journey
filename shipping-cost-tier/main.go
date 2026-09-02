package main

import "fmt"

func main() {
	cartTotal := 1450.0
	var shippingCost float64

	if cartTotal > 2000 {
		shippingCost = 0.0
	} else if cartTotal > 1000 {
		shippingCost = 60.0
	} else {
		shippingCost = 100.0
	}
	fmt.Printf("Cart: %.2f | Shipping: %.2f | Payable: %.2f\n", cartTotal, shippingCost, cartTotal+shippingCost)
}
