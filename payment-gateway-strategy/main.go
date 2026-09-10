package main

import "fmt"

type PaymentProcessor func(amount float64, currency string) error

func processStripe(amount float64, currency string) error {
	fmt.Printf("Processing $%.2f %s via Stripe API...\n", amount, currency)
	return nil
}

func ProcessPayPal(amount float64, currency string) error {
	fmt.Printf("Processing $%.2f %s via PayPal API...\n", amount, currency)
	return nil
}

func Checkout(processor PaymentProcessor, amount float64, currency string) error {
	fmt.Println("validating cart...")

	if err := processor(amount, currency); err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	fmt.Println("sending recipient email...")

	return nil
}

func main() {
	_ = Checkout(processStripe, 99.99, "usd")

	fmt.Println("----")

	_ = Checkout(ProcessPayPal, 10.11, "BDT")
}
