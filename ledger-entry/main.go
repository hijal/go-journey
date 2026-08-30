package main

import "fmt"

type Money float64

type LedgerEntry struct {
	AccountID string
	Debit     Money
	Credit    Money
}

func main() {
	entries := []LedgerEntry{
		{
			AccountID: "cach",
			Debit:     5000,
			Credit:    0,
		},
		{
			AccountID: "revenue",
			Debit:     0,
			Credit:    5000,
		},
	}

	var totalDebit, totalCredit Money

	for _, entry := range entries {
		totalCredit += entry.Credit
		totalDebit += entry.Debit
		fmt.Println(entry.AccountID, "- debit:", entry.Debit, "credit:", entry.Credit)
	}

	fmt.Println("Books balanced:", totalDebit == totalCredit)
}
