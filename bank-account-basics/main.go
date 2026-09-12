package main

import "fmt"

type Account struct {
	Owner   string
	Number  string
	Balance float64
	Active  bool
}

func main() {
	var acc Account
	fmt.Printf("Zero value: %+v\n", acc)

	acc = Account{
		Owner:   "John Doe",
		Number:  "123456789",
		Balance: 5000.0,
		Active:  true,
	}

	acc.Balance -= 1000
	fmt.Printf("After withdrawal: %+v\n", acc)
}
