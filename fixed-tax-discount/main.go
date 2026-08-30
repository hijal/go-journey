package main

import "fmt"

const vatRate = 0.15
const discount = 0.10
const currency = "BDT"

func main() {
	subtotal := 1100.0

	tax := subtotal * vatRate
	afterDiscount := subtotal * (1 - discount)
	total := afterDiscount + tax

	fmt.Printf("%s %.2f\n", currency, total)
}
