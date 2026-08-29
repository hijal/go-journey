package main

import "fmt"

type Account struct {
	Code    string
	balance int64
}

func (a *Account) Credit(amountCents int64) {
	a.balance += amountCents
}

func (a *Account) Debit(amountCents int64) error {
	if a.balance < amountCents {
		return fmt.Errorf("insufficient balance in account %s: have %d, need %d", a.Code, a.balance, amountCents)
	}
	a.balance -= amountCents
	return nil
}

func (a *Account) Balance() int64 {
	return a.balance
}

func main() {
	cash := &Account{
		Code:    "1000-CASH",
		balance: 0,
	}
	revenue := &Account{
		Code:    "2000-REV",
		balance: 0,
	}

	cash.Credit(50000)
	if err := cash.Debit(0); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Cash:", cash.Balance())
	fmt.Println("Revenue:", revenue.Balance())

	if err := cash.Debit(60000); err != nil {
		fmt.Println("debit failed", err)
	}
}
