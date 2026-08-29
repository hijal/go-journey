package main

import "fmt"

func main() {
	quantity := 3
	Quantity := 10

	fmt.Println("Cart requested quantity:", quantity)
	fmt.Println("Warehouse stock quantity:", Quantity)

	if quantity <= Quantity {
		fmt.Println("Order can be fulfilled")
	} else {
		fmt.Println("Not enough stock")
	}
}
