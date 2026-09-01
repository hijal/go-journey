package main

import "fmt"

func main() {
	cartTotal := 1499.0
	itemCount := 2
	isPremium := false

	freeShipping := cartTotal >= 1500 || (isPremium && itemCount > 0)
	bulkOrder := itemCount >= 10
	upSellCandidate := cartTotal < 1500 && !isPremium

	fmt.Println("free shipping:", freeShipping)
	fmt.Println("bulk order:", bulkOrder)
	fmt.Println("show free-shipping upsell banner:", upSellCandidate)
}
