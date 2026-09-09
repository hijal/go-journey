package main

import "fmt"

func openWallet(owner string) (deposit func(int), withdraw func(int) error, current func() int) {
	balance := 0

	deposit = func(amount int) {
		balance += amount
		fmt.Printf("%s deposit %d, balance %d\n", owner, amount, balance)
	}

	withdraw = func(amount int) error {
		if amount > balance {
			return fmt.Errorf("%s withdraw %d: insufficient funds(balance %d)", owner, amount, balance)
		}

		balance -= amount
		fmt.Printf("%s withdraw %d, balance %d\n", owner, amount, balance)
		return nil
	}
	return deposit, withdraw, func() int { return balance }
}

func main() {
	deposit, withdraw, current := openWallet("john")

	deposit(500)
	deposit(300)

	if err := withdraw(200); err != nil {
		fmt.Println("error:", err)
	}

	if err := withdraw(1000); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("final balance:", current())
}
