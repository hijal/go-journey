package main

import (
	"errors"
	"fmt"
)

var errReviewRequired = errors.New("credit needs manual review")
var balances = map[string]int{"acc-1": 5000, "acc-2": 1000}

func transfer(from, to string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("transfer: invalid amount %d", amount)
	}
	if balances[from] < amount {
		return fmt.Errorf("transfer: insufficient funds in %s: have %d", from, balances[from])
	}

	balances[from] -= amount

	failed := false

	defer func() {
		if failed {
			balances[from] += amount
			fmt.Printf("role back %d from %s\n", amount, from)
		}
	}()

	if amount > 2000 {
		failed = true
		return fmt.Errorf("transfer: %w", errReviewRequired)
	}
	balances[to] += amount
	return nil
}

func main() {
	if err := transfer("acc-1", "acc-2", 1500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("after ok transfer:", balances)

	if err := transfer("acc-1", "acc-2", 2500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("after failed transfer:", balances)
}
