package main

import (
	"errors"
	"fmt"
	"log"
)

type PaymentProcessor interface {
	Charge(customerID string, amountCents int64) (txId string, err error)
	Refund(txID string, amountCents int64) error
	Name() string
}

var ErrInvalidAmount = errors.New("amount must be positive")

type StripeProcessor struct {
	APIKey string
}

type BkashProcessor struct {
	MerchantNumber string
}

var _ PaymentProcessor = (*StripeProcessor)(nil)

func (s *StripeProcessor) Charge(customerID string, amountCents int64) (txId string, err error) {
	if amountCents <= 0 {
		return "", ErrInvalidAmount
	}
	return fmt.Sprintf("ch_%s_%d", customerID, amountCents), nil
}

func (s *StripeProcessor) Refund(txID string, amountCents int64) error {
	if amountCents <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

func (s *StripeProcessor) Name() string {
	return "stripe"
}

var _ PaymentProcessor = (*BkashProcessor)(nil)

func (b *BkashProcessor) Charge(customerID string, amountCents int64) (txId string, err error) {
	if amountCents <= 0 {
		return "", ErrInvalidAmount
	}
	return fmt.Sprintf("BK-%s-%d", customerID, amountCents), nil
}

func (b *BkashProcessor) Refund(txID string, amountCents int64) error {
	if amountCents <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

func (b *BkashProcessor) Name() string {
	return "bkash"
}

func Checkout(p PaymentProcessor, customerID string, amountCents int64) (string, error) {
	txID, err := p.Charge(customerID, amountCents)

	if err != nil {
		return "", fmt.Errorf("checkout via %s: %w", p.Name(), err)
	}
	return txID, nil
}

func main() {
	providers := map[string]PaymentProcessor{
		"US": &StripeProcessor{APIKey: "sk_test_123"},
		"BD": &BkashProcessor{MerchantNumber: "01700000000"},
	}

	for _, country := range []string{"US", "BD"} {
		txID, err := Checkout(providers[country], "c_123", 2599)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: charged via %s, tx=%s\n", country, providers[country].Name(), txID)
	}

	if _, err := Checkout(providers["US"], "cust42", 0); errors.Is(err, ErrInvalidAmount) {
		fmt.Println("rejected:", err)
	}
}
