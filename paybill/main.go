package main

import (
	"errors"
	"fmt"
)

func payBill(amount, balance float64) (string, error) {
	if amount <= 0 {
		return "", errors.New("amount must be greater than zero")
	}

	if amount > balance {
		return "", errors.New("insufficiant balance")
	}

	code := fmt.Sprintf("TXN-%.0f-OK", amount)
	return code, nil
}

func main() {
	code, err := payBill(500, 2000)
	if err != nil {
		fmt.Println("payment failed:", err)
		return
	}

	fmt.Println("payment successfull:", code)

	_, err = payBill(5000, 2000)
	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}
}
