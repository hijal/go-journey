package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	committed bool
}

func (t *Transaction) commit() {
	t.committed = true
	fmt.Println("transaction committed")
}

func (t *Transaction) rollback() {
	if !t.committed {
		fmt.Println("transaction rolled back")
	}
}

func transferFunds(amount float64) error {
	tx := &Transaction{}

	defer tx.rollback()

	if amount <= 0 {
		return errors.New("invalid transfer amount")
	}

	fmt.Println("transferring", amount)
	tx.commit()
	return nil
}

func main() {

	if err := transferFunds(100); err != nil {
		fmt.Println("error", err)
	}
	
	fmt.Println("--------------")
	if err := transferFunds(-10); err != nil {
		fmt.Println("error", err)
	}
}
