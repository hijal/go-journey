package main

import (
	"encoding/json"
	"fmt"
)

type Invoice struct {
	ID           string
	CustomerName string
	amountCents  int
}

func main() {
	inv := Invoice{
		ID:           "INV-1000",
		CustomerName: "John Doe",
		amountCents:  5600,
	}

	data, err := json.Marshal(inv)

	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}

	fmt.Println(string(data))
	fmt.Println("Internal amount (only visible inside this package):", inv.amountCents)
}
