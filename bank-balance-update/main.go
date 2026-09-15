package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive, got %.2f", amount)
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive, got %.2f", amount)
	}

	if amount > a.Balance {
		return fmt.Errorf("insufficient funds: balance %.2f, requested %.2f", a.Balance, amount)
	}

	a.Balance -= amount

	return nil
}

func main() {
	acc := &Account{Owner: "Alice", Balance: 100.0}

	if err := acc.Deposit(50); err != nil {
		fmt.Println("deposit error:", err)
	}

	if err := acc.Withdraw(30); err != nil {
		fmt.Println("withdraw error:", err)
	}

	fmt.Printf("%s's balance: %.2f\n", acc.Owner, acc.Balance)
}
