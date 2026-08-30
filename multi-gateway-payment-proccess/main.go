package main

import "fmt"

type PaymentProcessor interface {
	Charge(amountCents int) error
	Name() string
}

type StripeStyleProcessor struct{}

func (StripeStyleProcessor) Charge(amountCents int) error {
	fmt.Println("charging", amountCents, "cents via card network")
	return nil

}

func (StripeStyleProcessor) Name() string {
	return "card"
}

type MobileWalletProcessor struct{}

func (MobileWalletProcessor) Charge(amountCents int) error {
	fmt.Println("charging", amountCents, "cents via mobile wallet")
	return nil
}

func (MobileWalletProcessor) Name() string {
	return "mobile-wallet"
}

func processPayment(p PaymentProcessor, amountCents int) {
	fmt.Println("using processor:", p.Name())
	if err := p.Charge(amountCents); err != nil {
		fmt.Println("Payment failed", err)
	}
}

func main() {
	processors := []PaymentProcessor{
		StripeStyleProcessor{},
		MobileWalletProcessor{},
	}

	for _, p := range processors {
		processPayment(p, 25000)
	}
}
