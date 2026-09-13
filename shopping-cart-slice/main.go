package main

import "fmt"

func main() {
	cart := []float64{24.99, 19.99, 5.49, 12.99}
	fmt.Println("items in cart:", len(cart))
	fmt.Println("first item costs:", cart[0])
	fmt.Println("last item costs:", cart[len(cart)-1])

	total := 0.0
	for _, item := range cart {
		total += item
	}
	fmt.Printf("total cost: %.2f\n", total)
}
