package main

import "fmt"

const (
	ChargeRate = 0.10
	VatRate    = 0.15
	People     = 7
)

func main() {
	subtotal := 2450.0
	serviceCharge := subtotal * ChargeRate
	vat := (subtotal + serviceCharge) * VatRate
	total := subtotal + serviceCharge + vat

	perHead := total / float64(People)

	fmt.Printf("subtotal: %.2f\n", subtotal)
	fmt.Printf("service charge: %.2f\n", serviceCharge)
	fmt.Printf("vat: %.2f\n", vat)
	fmt.Printf("total: %.2f\n", total)
	fmt.Printf("per head (%d people): %.2f\n", People, perHead)
}
