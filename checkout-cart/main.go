package main

import "fmt"

func main() {
	itemPrice := 249.40
	quantity := 3
	shippingFee := 60.0

	subtotal := itemPrice * float64(quantity)
	total := subtotal + shippingFee

	fmt.Println("Subtotal:", subtotal)
	fmt.Println("Shipping fee:", shippingFee)
	fmt.Println("Total payable:", total)
}
