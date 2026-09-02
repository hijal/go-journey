package main

import "fmt"

type Transaction struct {
	ID     string
	Amount float64
	Online bool
}

func main() {
	transactions := []Transaction{
		{ID: "TX-001", Amount: 800, Online: true},
		{ID: "TX-002", Amount: 250000, Online: true},
		{ID: "TX-003", Amount: 1200, Online: false},
		{ID: "TX-004", Amount: 95000, Online: false},
	}

	flagged := 0
	for _, tx := range transactions {
		if tx.Amount < 50000 {
			fmt.Printf("%s approved\n", tx.ID)
			continue
		}

		flagged++

		if tx.Amount > 200000 && tx.Online {
			fmt.Printf("%s: BLOCKED(large online transaction)\n", tx.ID)
		} else {
			fmt.Printf("%s: flagged for manual review\n", tx.ID)
		}
	}

	if flagged == 0 {
		fmt.Println("All transactions passed automatic checks")
	} else {
		fmt.Printf("%d transaction(s) need attention\n", flagged)
	}
}
