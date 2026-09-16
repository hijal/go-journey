package main

import "fmt"

type PaymentMethod interface {
	Pay(amountCents int64) string
}

type Card struct {
	last4 string
}

func (c Card) Pay(amountCents int64) string {
	return fmt.Sprintf("card ****%s charged ৳%.2f", c.last4, float64(amountCents)/100)
}

type MobileWallet struct {
	Phone string
}

func (w MobileWallet) Pay(amountCents int64) string {
	return fmt.Sprintf("wallet %s paid ৳%.2f", w.Phone, float64(amountCents)/100)
}

func checkout(method PaymentMethod, amountCents int64) {
	fmt.Println(method.Pay(amountCents))
}

func main() {
	checkout(Card{last4: "4242"}, 125000)
	checkout(MobileWallet{Phone: "+8801700-000000"}, 59900)

	methods := []PaymentMethod{
		Card{last4: "1111"},
		MobileWallet{Phone: "+8801811-111111"},
	}

	for _, m := range methods {
		checkout(m, 25000)
	}
}
