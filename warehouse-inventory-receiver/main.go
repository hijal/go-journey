package main

import (
	"errors"
	"fmt"
)

type Inventory struct {
	SKU   string
	Name  string
	Stock int
}

var ErrorOutOfStock = errors.New("inventory out of stock")

func (i *Inventory) Add(qty int) {
	i.Stock += qty
}

func (i *Inventory) Pick(qty int) error {
	if qty > i.Stock {
		return fmt.Errorf("pick %d of %s: %w", qty, i.SKU, ErrorOutOfStock)
	}

	i.Stock -= qty
	return nil
}

func (i *Inventory) Available() int {
	return i.Stock
}

func main() {
	inv := &Inventory{SKU: "SH-42", Name: "Running Shoes 42", Stock: 10}

	inv.Add(5)
	fmt.Println("after restock:", inv.Available())

	if err := inv.Pick(12); err != nil {
		fmt.Println("order failed:", err)
		fmt.Println("is out of stock error", errors.Is(err, ErrorOutOfStock))
	}

	if err := inv.Pick(7); err != nil {
		fmt.Println("unexpected:", err)
		return
	}

	fmt.Println("after pickup:", inv.Available())
}
